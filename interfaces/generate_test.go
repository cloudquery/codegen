package interfaces

import (
	"context"
	"io"
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type Client struct{}

func (*Client) ListThings() error {
	return nil
}

func (*Client) ListTables(_ context.Context) error {
	return nil
}

func (*Client) CreateTables(_ context.Context, _ []string) error {
	return nil
}

var wantOutput = `// Code generated by codegen; DO NOT EDIT.
package services

import (
	"github.com/cloudquery/codegen/interfaces"
	"net/http"
)

//go:generate mockgen -package=mocks -destination=../mocks/interfaces.go -source=interfaces.go InterfacesClient
type InterfacesClient interface {
	ListTables(context.Context) error
}
`

func TestGenerate(t *testing.T) {
	dir := t.TempDir()
	err := Generate([]any{&Client{}}, dir, WithIncludeFunc(func(m reflect.Method) bool {
		return MethodHasAnyPrefix(m, []string{"List"}) && MethodHasAnySuffix(m, []string{"Tables"})
	}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	fileName := path.Join(dir, "interfaces.go")
	f, err := os.Open(fileName)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	if diff := cmp.Diff(string(b), wantOutput); diff != "" {
		t.Errorf("unexpected diff (-got +want):\n%s", diff)
	}
}
