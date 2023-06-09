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

func getTemplateDataFromClientInfos(clientInfos []clientInfo) []serviceTemplateData {
	services := make([]serviceTemplateData, 0)
	serviceMap := make(map[string][]clientInfo)
	for _, clientInfo := range clientInfos {
		serviceMap[clientInfo.PackageName] = append(serviceMap[clientInfo.PackageName], clientInfo)
	}
	for packageName, clientInfos := range serviceMap {
		imports := make(map[string]bool)
		clientsTemplateData := make([]clientTemplateData, 0)
		for _, clientInfo := range clientInfos {
			imports[clientInfo.Import] = true
			for _, extraImport := range clientInfo.ExtraImports {
				imports[extraImport] = true
			}
			clientsTemplateData = append(clientsTemplateData, clientTemplateData{Name: clientInfo.ClientName, Signatures: clientInfo.Signatures})
		}
		services = append(services, serviceTemplateData{
			PackageName: packageName,
			Imports:     maps.Keys(imports),
			Clients:     clientsTemplateData,
		})
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

	services := getTemplateDataFromClientInfos(clientInfos)

	for _, service := range services {
		buff := bytes.Buffer{}
		if err := serviceTpl.Execute(&buff, service); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
		filePath := path.Join(dir, fmt.Sprintf("%s.go", service.PackageName))
		err := formatAndWriteFile(filePath, buff)
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
	PackageName  string
	ClientName   string
	Signatures   []string
	ExtraImports []string
}

type clientTemplateData struct {
	Name       string
	Signatures []string
}

type serviceTemplateData struct {
	PackageName string
	Imports     []string
	Clients     []clientTemplateData
}

func getClientInfo(client any, opts *Options) clientInfo {
	v := reflect.ValueOf(client)
	t := v.Type()
	pkgPath := t.Elem().PkgPath()
	parts := strings.Split(pkgPath, "/")
	versionPattern := regexp.MustCompile(`/v\d+$`)
	pkgName := parts[len(parts)-1]
	if versionPattern.MatchString(pkgPath) {
		pkgName = parts[len(parts)-2]
	}
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
		PackageName:  pkgName,
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
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}
