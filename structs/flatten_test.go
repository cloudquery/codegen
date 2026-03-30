package structs

import (
	"go/ast"
	"os"
	"path/filepath"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/stretchr/testify/require"
)

func TestFlatten_Basic(t *testing.T) {
	dir := t.TempDir()
	err := Flatten("testdata/input.go", []StructConfig{
		{
			SourceName: "Finding",
			OutputName: "flatFinding",
			OutputFile: "finding_generated.go",
		},
	}, dir, WithPackageName("testpkg"))
	require.NoError(t, err)

	got, err := os.ReadFile(filepath.Join(dir, "finding_generated.go"))
	require.NoError(t, err)
	cupaloy.SnapshotT(t, string(got))
}

func TestFlatten_MultipleStructs(t *testing.T) {
	dir := t.TempDir()
	err := Flatten("testdata/input.go", []StructConfig{
		{
			SourceName: "Finding",
			OutputName: "flatFinding",
			OutputFile: "finding_generated.go",
		},
		{
			SourceName: "Simple",
			OutputName: "flatSimple",
			OutputFile: "simple_generated.go",
		},
	}, dir, WithPackageName("testpkg"))
	require.NoError(t, err)

	got1, err := os.ReadFile(filepath.Join(dir, "finding_generated.go"))
	require.NoError(t, err)
	got2, err := os.ReadFile(filepath.Join(dir, "simple_generated.go"))
	require.NoError(t, err)
	cupaloy.SnapshotT(t, string(got1), string(got2))
}

func TestFlatten_ExtraFields(t *testing.T) {
	dir := t.TempDir()
	err := Flatten("testdata/input.go", []StructConfig{
		{
			SourceName: "Simple",
			OutputName: "extendedSimple",
			OutputFile: "extended_generated.go",
			ExtraFields: []Field{
				{Name: "AssetType", Type: "string", JSONTag: "assetType,omitempty"},
			},
		},
	}, dir, WithPackageName("testpkg"))
	require.NoError(t, err)

	got, err := os.ReadFile(filepath.Join(dir, "extended_generated.go"))
	require.NoError(t, err)
	cupaloy.SnapshotT(t, string(got))
}

func TestFlatten_WithExtraScalars(t *testing.T) {
	dir := t.TempDir()
	// Treat Address as a scalar — it should stay typed instead of map[string]any.
	err := Flatten("testdata/input.go", []StructConfig{
		{
			SourceName: "Finding",
			OutputName: "flatFinding",
			OutputFile: "finding_generated.go",
		},
	}, dir, WithPackageName("testpkg"), WithExtraScalars("Address"))
	require.NoError(t, err)

	got, err := os.ReadFile(filepath.Join(dir, "finding_generated.go"))
	require.NoError(t, err)
	cupaloy.SnapshotT(t, string(got))
}

func TestFlatten_SortFields(t *testing.T) {
	dir := t.TempDir()
	err := Flatten("testdata/input.go", []StructConfig{
		{
			SourceName: "Finding",
			OutputName: "sortedFinding",
			OutputFile: "sorted_generated.go",
		},
	}, dir, WithPackageName("testpkg"), WithSortFields())
	require.NoError(t, err)

	got, err := os.ReadFile(filepath.Join(dir, "sorted_generated.go"))
	require.NoError(t, err)
	cupaloy.SnapshotT(t, string(got))
}

func TestFlatten_SortFieldsWithExtraFields(t *testing.T) {
	dir := t.TempDir()
	err := Flatten("testdata/input.go", []StructConfig{
		{
			SourceName: "Finding",
			OutputName: "sortedFinding",
			OutputFile: "sorted_generated.go",
			ExtraFields: []Field{
				{Name: "AssetType", Type: "string", JSONTag: "assetType,omitempty"},
			},
		},
	}, dir, WithPackageName("testpkg"), WithSortFields())
	require.NoError(t, err)

	got, err := os.ReadFile(filepath.Join(dir, "sorted_generated.go"))
	require.NoError(t, err)
	cupaloy.SnapshotT(t, string(got))
}

func TestFlatten_StructNotFound(t *testing.T) {
	dir := t.TempDir()
	err := Flatten("testdata/input.go", []StructConfig{
		{
			SourceName: "NonExistent",
			OutputName: "flat",
			OutputFile: "flat.go",
		},
	}, dir, WithPackageName("testpkg"))
	require.Error(t, err)
	require.Contains(t, err.Error(), `struct "NonExistent" not found`)
}

func TestFlatten_FileNotFound(t *testing.T) {
	dir := t.TempDir()
	err := Flatten("testdata/nonexistent.go", []StructConfig{
		{SourceName: "X", OutputName: "Y", OutputFile: "y.go"},
	}, dir)
	require.Error(t, err)
	require.Contains(t, err.Error(), "nonexistent.go")
}

func TestResolveType(t *testing.T) {
	typeKinds := map[string]string{
		"Severity": "string",
		"Address":  "struct",
		"Tag":      "struct",
		"Entity":   "interface",
	}
	scalars := defaultScalarKinds

	tests := []struct {
		name string
		expr ast.Expr
		want string
	}{
		{
			name: "scalar string",
			expr: &ast.Ident{Name: "string"},
			want: "string",
		},
		{
			name: "scalar bool",
			expr: &ast.Ident{Name: "bool"},
			want: "bool",
		},
		{
			name: "string-based enum",
			expr: &ast.Ident{Name: "Severity"},
			want: "string",
		},
		{
			name: "struct type",
			expr: &ast.Ident{Name: "Address"},
			want: "map[string]any",
		},
		{
			name: "pointer to scalar",
			expr: &ast.StarExpr{X: &ast.Ident{Name: "float64"}},
			want: "*float64",
		},
		{
			name: "pointer to struct",
			expr: &ast.StarExpr{X: &ast.Ident{Name: "Address"}},
			want: "map[string]any",
		},
		{
			name: "slice of scalars",
			expr: &ast.ArrayType{Elt: &ast.Ident{Name: "string"}},
			want: "[]string",
		},
		{
			name: "slice of structs",
			expr: &ast.ArrayType{Elt: &ast.Ident{Name: "Tag"}},
			want: "[]map[string]any",
		},
		{
			name: "slice of pointer to struct",
			expr: &ast.ArrayType{Elt: &ast.StarExpr{X: &ast.Ident{Name: "Tag"}}},
			want: "[]map[string]any",
		},
		{
			name: "cross-package scalar (time.Time)",
			expr: &ast.SelectorExpr{X: &ast.Ident{Name: "time"}, Sel: &ast.Ident{Name: "Time"}},
			want: "time.Time",
		},
		{
			name: "cross-package struct",
			expr: &ast.SelectorExpr{X: &ast.Ident{Name: "api"}, Sel: &ast.Ident{Name: "Finding"}},
			want: "map[string]any",
		},
		{
			name: "interface type",
			expr: &ast.Ident{Name: "Entity"},
			want: "map[string]any",
		},
		{
			name: "map type",
			expr: &ast.MapType{Key: &ast.Ident{Name: "string"}, Value: &ast.Ident{Name: "string"}},
			want: "map[string]any",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveType(tt.expr, typeKinds, scalars)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestExtractJSONTag(t *testing.T) {
	tests := []struct {
		name string
		tag  *ast.BasicLit
		want string
	}{
		{name: "nil tag", tag: nil, want: ""},
		{name: "simple", tag: &ast.BasicLit{Value: "`json:\"id\"`"}, want: "id"},
		{name: "with omitempty", tag: &ast.BasicLit{Value: "`json:\"name,omitempty\"`"}, want: "name,omitempty"},
		{name: "skip tag", tag: &ast.BasicLit{Value: "`json:\"-\"`"}, want: "-"},
		{name: "multiple tags", tag: &ast.BasicLit{Value: "`csv:\"id\" json:\"id\"`"}, want: "id"},
		{name: "no json tag", tag: &ast.BasicLit{Value: "`csv:\"id\"`"}, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractJSONTag(tt.tag)
			require.Equal(t, tt.want, got)
		})
	}
}
