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
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

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
