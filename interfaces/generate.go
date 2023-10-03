package interfaces

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"os"
	"path"
	"reflect"
	"regexp"
	"strings"
	"text/template"

	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/jpillora/longestcommon"
	"golang.org/x/exp/maps"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

type Options struct {
	// ShouldInclude tests whether a method should be included in the generated interfaces. If it returns true,
	// the method will be included. MethodHasPrefix and MethodHasSuffix can be used inside a custom function here
	// to customize the behavior.
	ShouldInclude func(reflect.Method) bool

	// ExtraImports can add extra imports for a method
	ExtraImports func(reflect.Method) []string

	// SinglePackage allows to generate all passed clients into a single package.
	// The clients will get their package name as prefix to the interface name (e.g., s3.Client -> S3Client)
	SinglePackage string
}

func (o *Options) SetDefaults() {
	if o.ShouldInclude == nil {
		o.ShouldInclude = func(reflect.Method) bool { return true }
	}
	if o.ExtraImports == nil {
		o.ExtraImports = func(reflect.Method) []string { return nil }
	}
}

type Option func(*Options)

func WithIncludeFunc(f func(reflect.Method) bool) Option {
	return func(o *Options) {
		o.ShouldInclude = f
	}
}

func WithExtraImports(f func(reflect.Method) []string) Option {
	return func(o *Options) {
		o.ExtraImports = f
	}
}

func WithSinglePackage(name string) Option {
	return func(o *Options) {
		o.SinglePackage = name
	}
}

func getPackageNames(clientInfos []clientInfo) []string {
	versionPattern := regexp.MustCompile(`/v\d+$`)
	allImports := make([]string, len(clientInfos))
	for i, clientInfo := range clientInfos {
		allImports[i] = clientInfo.Import
	}
	// To get the shortest possible package name without collisions, we need to find the longest common prefix
	importPrefix := longestcommon.Prefix(allImports)
	packageNames := make([]string, len(clientInfos))
	for i, clientInfo := range clientInfos {
		var pkgName string
		if clientInfo.Import == importPrefix {
			pkgName = versionPattern.ReplaceAllString(clientInfo.Import, "")
			pkgName = path.Base(pkgName)
		} else {
			pkgName = strings.TrimPrefix(clientInfo.Import, importPrefix)
			pkgName = strings.ReplaceAll(versionPattern.ReplaceAllString(pkgName, ""), "/", "_")
		}

		packageNames[i] = strings.ReplaceAll(pkgName, "-", "")
	}
	return packageNames
}

func getTemplateDataFromClientInfos(clientInfos []clientInfo, options *Options) []serviceTemplateData {
	packageNames := getPackageNames(clientInfos)
	services := make([]serviceTemplateData, 0)
	serviceMap := make(map[string][]clientInfo)
	for i, clientInfo := range clientInfos {
		serviceMap[packageNames[i]] = append(serviceMap[packageNames[i]], clientInfo)
	}
	for packageName, infos := range serviceMap {
		imports := make(map[string]bool)
		clientsTemplateData := make([]clientTemplateData, 0)
		for _, clientInfo := range infos {
			imports[clientInfo.Import] = true
			for _, extraImport := range clientInfo.ExtraImports {
				imports[extraImport] = true
			}
			clientsTemplateData = append(clientsTemplateData, clientInfo.templateData(len(options.SinglePackage) > 0))
		}
		svc := serviceTemplateData{
			PackageName: packageName,
			FileName:    packageName,
			Imports:     maps.Keys(imports),
			Clients:     clientsTemplateData,
		}
		if len(options.SinglePackage) > 0 {
			svc.PackageName = options.SinglePackage
		}
		services = append(services, svc)
	}
	return services
}

// Generate generates service interfaces to be used for generating
// mocks. The clients passed in as the first argument should be structs that will be used to
// generate the service interfaces. The second argument, dir, is the path to the output
// directory where the service interface files will be created.
func Generate(clients []any, dir string, opts ...Option) error {
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}
	options.SetDefaults()

	clientInfos := make([]clientInfo, 0)
	for _, client := range clients {
		clientInfos = append(clientInfos, getClientInfo(client, options))
	}

	// write individual service files
	serviceTpl, err := template.New("service.go.tpl").ParseFS(templatesFS, "templates/service.go.tpl")
	if err != nil {
		return err
	}

	services := getTemplateDataFromClientInfos(clientInfos, options)

	for _, service := range services {
		buff := bytes.Buffer{}
		if err := serviceTpl.Execute(&buff, service); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
		err := formatAndWriteFile(service.getFilePath(dir), buff)
		if err != nil {
			return fmt.Errorf("failed to format and write file for service %v: %w", service, err)
		}
	}

	return nil
}

func normalizedGenericTypeName(str string) string {
	// Generic output types have the full import path in the string value, so we need to normalize it
	pattern := regexp.MustCompile(`\[(.*?)\]`)
	groups := pattern.FindStringSubmatch((str))
	if len(groups) < 2 {
		return str
	}

	typeName := groups[1]
	normalizedGenericTypeName := strings.Split(typeName, "/")
	importName := normalizedGenericTypeName[len(normalizedGenericTypeName)-1]
	versionPattern := regexp.MustCompile(`/v\d+\.`)
	if versionPattern.MatchString(typeName) {
		// Example typeName: github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v2.ConfigurationStoresClientCreateResponse
		importName = normalizedGenericTypeName[len(normalizedGenericTypeName)-2] + "." + strings.Split(normalizedGenericTypeName[len(normalizedGenericTypeName)-1], ".")[1]
	}
	return pattern.ReplaceAllString(str, "["+importName+"]")
}

// Adapted from https://stackoverflow.com/a/54129236
func signature(name string, f any) string {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		return "<not a function>"
	}

	buf := strings.Builder{}
	buf.WriteString(name + "(")
	for i := 0; i < t.NumIn(); i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		if t.IsVariadic() && i == t.NumIn()-1 {
			buf.WriteString("..." + strings.TrimPrefix(t.In(i).String(), "[]"))
		} else {
			buf.WriteString(t.In(i).String())
		}
	}
	buf.WriteString(")")
	if numOut := t.NumOut(); numOut > 0 {
		if numOut > 1 {
			buf.WriteString(" (")
		} else {
			buf.WriteString(" ")
		}
		for i := 0; i < t.NumOut(); i++ {
			if i > 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(normalizedGenericTypeName(t.Out(i).String()))
		}
		if numOut > 1 {
			buf.WriteString(")")
		}
	}

	return buf.String()
}

type clientInfo struct {
	Import       string
	ClientName   string
	Signatures   []string
	ExtraImports []string
}

func (c clientInfo) templateData(singlePackageMode bool) clientTemplateData {
	var packageName string
	if singlePackageMode {
		packageName = caser.New().ToPascal(getPackageNames([]clientInfo{c})[0])
	}
	return clientTemplateData{
		Name:       packageName + c.ClientName,
		Signatures: c.Signatures,
	}
}

type clientTemplateData struct {
	Name       string
	Signatures []string
}

type serviceTemplateData struct {
	PackageName string
	FileName    string
	Imports     []string
	Clients     []clientTemplateData
}

func (s serviceTemplateData) getFilePath(baseDir string) string {
	if s.FileName == s.PackageName {
		return path.Join(baseDir, s.PackageName, fmt.Sprintf("%s.go", s.PackageName))
	}
	return path.Join(baseDir, fmt.Sprintf("%s.go", s.FileName))
}

func getClientInfo(client any, opts *Options) clientInfo {
	v := reflect.ValueOf(client)
	t := v.Type()
	pkgPath := t.Elem().PkgPath()
	clientName := t.Elem().Name()
	signatures := make([]string, 0)
	extraImports := make([]string, 0)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		if opts.ShouldInclude(method) {
			sig := signature(method.Name, v.Method(i).Interface())
			signatures = append(signatures, sig)
		}
		extraImports = append(extraImports, opts.ExtraImports(method)...)
	}
	return clientInfo{
		Import:       pkgPath,
		ClientName:   clientName,
		Signatures:   signatures,
		ExtraImports: extraImports,
	}
}

func formatAndWriteFile(filePath string, buff bytes.Buffer) error {
	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Printf("failed to format source: %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}
	if err := os.MkdirAll(path.Dir(filePath), 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path.Dir(filePath), err)
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}
