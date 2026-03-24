package structs

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

// buildTypeKinds walks the AST and builds a map of type name to the underlying kind.
// For example, `type Severity string` maps "Severity" → "string",
// and `type Finding struct{...}` maps "Finding" → "struct".
func buildTypeKinds(f *ast.File) map[string]string {
	kinds := map[string]string{}
	ast.Inspect(f, func(n ast.Node) bool {
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}
		switch t := ts.Type.(type) {
		case *ast.Ident:
			kinds[ts.Name.Name] = t.Name
		case *ast.StructType:
			kinds[ts.Name.Name] = "struct"
		case *ast.InterfaceType:
			kinds[ts.Name.Name] = "interface"
		}
		return true
	})
	return kinds
}

// findStruct locates a named struct type in the AST.
func findStruct(f *ast.File, name string) (*ast.StructType, error) {
	var found *ast.StructType
	ast.Inspect(f, func(n ast.Node) bool {
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}
		if ts.Name.Name == name {
			if st, ok := ts.Type.(*ast.StructType); ok {
				found = st
			}
		}
		return true
	})
	if found == nil {
		return nil, fmt.Errorf("struct %q not found", name)
	}
	return found, nil
}

// resolveType determines the optimized Go type for a field expression.
// Scalar types and string-based enums stay typed; nested structs become map[string]any.
func resolveType(expr ast.Expr, typeKinds map[string]string, scalars map[string]bool) string {
	switch t := expr.(type) {
	case *ast.Ident:
		if scalars[t.Name] {
			return t.Name
		}
		if kind, ok := typeKinds[t.Name]; ok && scalars[kind] {
			return kind
		}
		return "map[string]any"

	case *ast.StarExpr:
		inner := resolveType(t.X, typeKinds, scalars)
		if inner == "map[string]any" || inner == "[]map[string]any" {
			return inner
		}
		return "*" + inner

	case *ast.ArrayType:
		inner := resolveType(t.Elt, typeKinds, scalars)
		if inner == "map[string]any" {
			return "[]map[string]any"
		}
		return "[]" + inner

	default:
		// SelectorExpr (cross-package types), MapType, InterfaceType, etc.
		return "map[string]any"
	}
}

// extractFields walks a struct's field list and returns flattened Field values.
// Embedded fields and fields without JSON tags are skipped.
func extractFields(st *ast.StructType, typeKinds map[string]string, scalars map[string]bool) []Field {
	var fields []Field
	for _, f := range st.Fields.List {
		if len(f.Names) == 0 {
			continue // skip embedded fields
		}
		name := f.Names[0].Name
		jsonTag := extractJSONTag(f.Tag)
		if jsonTag == "" || jsonTag == "-" {
			continue
		}
		fields = append(fields, Field{
			Name:    name,
			Type:    resolveType(f.Type, typeKinds, scalars),
			JSONTag: jsonTag,
		})
	}
	return fields
}

// extractJSONTag extracts the JSON tag value from a struct field tag.
func extractJSONTag(tag *ast.BasicLit) string {
	if tag == nil {
		return ""
	}
	raw := strings.Trim(tag.Value, "`")
	for _, part := range strings.Fields(raw) {
		if strings.HasPrefix(part, `json:"`) {
			val := strings.TrimPrefix(part, `json:"`)
			val = strings.TrimSuffix(val, `"`)
			return val
		}
	}
	return ""
}

// parseFile parses a Go source file and returns the AST.
func parseFile(srcPath string) (*ast.File, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, srcPath, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parsing %s: %w", srcPath, err)
	}
	return f, nil
}

// baseName returns the last element of a path, used for the default package name.
func baseName(dir string) string {
	return filepath.Base(dir)
}
