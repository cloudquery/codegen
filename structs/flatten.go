// Package structs provides code generation utilities for optimizing Go struct
// types. The primary use case is flattening deeply nested API model structs
// into decode-efficient versions where nested struct fields are replaced with
// map[string]any, eliminating the decode -> allocate -> re-encode cycle that occurs
// when the CloudQuery SDK serializes nested fields as JSON columns.
package structs

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

var tmpl = template.Must(template.ParseFS(templatesFS, "templates/*.go.tpl"))

// StructConfig defines how to transform one source struct into an optimized version.
type StructConfig struct {
	// SourceName is the struct name in the source file (e.g. "VulnerabilityFinding").
	SourceName string
	// OutputName is the generated struct name (e.g. "customVulnerabilityFinding").
	OutputName string
	// OutputFile is the filename for the generated file (e.g. "vulnerability_finding_generated.go").
	OutputFile string
	// ExtraFields are additional fields prepended to the generated struct.
	ExtraFields []Field
}

// Field represents a struct field in the generated output.
type Field struct {
	Name    string
	Type    string
	JSONTag string
}

type templateData struct {
	HeaderComment string
	PackageName   string
	SourceName    string
	OutputName    string
	Fields        []Field
}

// Flatten parses the Go source file at srcPath, finds the structs described by
// configs, replaces deeply nested struct fields with map[string]any, and writes
// generated files into outputDir.
//
// Scalar types and string-based enum types
// are kept typed. All other types (structs, interfaces, cross-package types,
// maps) are flattened to map[string]any. Pointer wrappers are preserved for
// scalar types but collapsed for complex types.
func Flatten(srcPath string, configs []StructConfig, outputDir string, opts ...Option) error {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	o.setDefaults(outputDir)

	f, err := parseFile(srcPath)
	if err != nil {
		return err
	}

	typeKinds := buildTypeKinds(f)
	scalars := o.scalars()

	var errs []error
	for _, cfg := range configs {
		if err := flattenOne(f, cfg, typeKinds, scalars, outputDir, &o); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", cfg.SourceName, err))
		}
	}
	return errors.Join(errs...)
}

func flattenOne(f *ast.File, cfg StructConfig, typeKinds map[string]string, scalars map[string]bool, outputDir string, o *options) error {
	st, err := findStruct(f, cfg.SourceName)
	if err != nil {
		return err
	}

	fields := make([]Field, 0, len(cfg.ExtraFields)+len(st.Fields.List))
	fields = append(fields, cfg.ExtraFields...)
	sourceFields := extractFields(st, typeKinds, scalars)
	if o.sortFields {
		sortFields(sourceFields)
	}
	fields = append(fields, sourceFields...)

	data := templateData{
		HeaderComment: o.headerComment,
		PackageName:   o.packageName,
		SourceName:    cfg.SourceName,
		OutputName:    cfg.OutputName,
		Fields:        fields,
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "struct.go.tpl", data); err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		// Write unformatted so the developer can see what went wrong.
		formatted = buf.Bytes()
	}

	outputPath := filepath.Join(outputDir, cfg.OutputFile)
	if err := os.WriteFile(outputPath, formatted, 0o644); err != nil {
		return fmt.Errorf("writing %s: %w", outputPath, err)
	}

	fmt.Printf("Generated %s with %d fields\n", outputPath, len(fields))
	return nil
}

// sortFields sorts fields with "ID" first, then alphabetically by name (case-insensitive).
func sortFields(fields []Field) {
	slices.SortStableFunc(fields, func(a, b Field) int {
		aIsID := strings.EqualFold(a.Name, "ID")
		bIsID := strings.EqualFold(b.Name, "ID")
		switch {
		case aIsID && !bIsID:
			return -1
		case !aIsID && bIsID:
			return 1
		default:
			return strings.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
		}
	})
}
