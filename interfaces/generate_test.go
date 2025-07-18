package interfaces

import (
	"context"
	"io"
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v3"
	resourcesarmmanagedapplications "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications"
	solutionsarmmanagedapplications "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications/v2"
	"github.com/google/go-cmp/cmp"
)

type Response struct {
}
type Pager[T any] struct {
}

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

func (*Client) NewPager(_ context.Context) *Pager[Response] {
	return nil
}

var wantOutput = `// Code generated by codegen; DO NOT EDIT.
package interfaces

import (
	"github.com/cloudquery/codegen/interfaces"
	"net/http"
)

//go:generate mockgen -package=mocks -destination=../mocks/interfaces.go -source=interfaces.go Client
type Client interface {
	ListTables(context.Context) error
	NewPager(context.Context) *interfaces.Pager[interfaces.Response]
}
`

type Client2 struct{}

func (*Client2) ListTables(_ context.Context) error {
	return nil
}

var wantOutputMultipleClients = `// Code generated by codegen; DO NOT EDIT.
package interfaces

import (
	"github.com/cloudquery/codegen/interfaces"
	"net/http"
)

//go:generate mockgen -package=mocks -destination=../mocks/interfaces.go -source=interfaces.go Client
type Client interface {
	ListTables(context.Context) error
	NewPager(context.Context) *interfaces.Pager[interfaces.Response]
}

//go:generate mockgen -package=mocks -destination=../mocks/interfaces.go -source=interfaces.go Client2
type Client2 interface {
	ListTables(context.Context) error
}
`

var wantOutputNonV1 = `// Code generated by codegen; DO NOT EDIT.
package armappconfiguration

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v3"
)

//go:generate mockgen -package=mocks -destination=../mocks/armappconfiguration.go -source=armappconfiguration.go ConfigurationStoresClient
type ConfigurationStoresClient interface {
	NewListPager(*armappconfiguration.ConfigurationStoresClientListOptions) *runtime.Pager[armappconfiguration.ConfigurationStoresClientListResponse]
}
`

var wantOutputOverlappingPackages = map[string]string{
	"solutions_armmanagedapplications/solutions_armmanagedapplications.go": `// Code generated by codegen; DO NOT EDIT.
package solutions_armmanagedapplications

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications/v2"
)

//go:generate mockgen -package=mocks -destination=../mocks/solutions_armmanagedapplications.go -source=solutions_armmanagedapplications.go ApplicationClient
type ApplicationClient interface {
}
`,

	"resources_armmanagedapplications/resources_armmanagedapplications.go": `// Code generated by codegen; DO NOT EDIT.
package resources_armmanagedapplications

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications"
)

//go:generate mockgen -package=mocks -destination=../mocks/resources_armmanagedapplications.go -source=resources_armmanagedapplications.go ApplicationClient
type ApplicationClient interface {
}
`,
	"appconfiguration_armappconfiguration/appconfiguration_armappconfiguration.go": `// Code generated by codegen; DO NOT EDIT.
package appconfiguration_armappconfiguration

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v3"
)

//go:generate mockgen -package=mocks -destination=../mocks/appconfiguration_armappconfiguration.go -source=appconfiguration_armappconfiguration.go ConfigurationStoresClient
type ConfigurationStoresClient interface {
	NewListPager(*armappconfiguration.ConfigurationStoresClientListOptions) *runtime.Pager[armappconfiguration.ConfigurationStoresClientListResponse]
}
`,
}

func TestGenerate(t *testing.T) {
	dir := t.TempDir()
	err := Generate([]any{&Client{}}, dir,
		WithIncludeFunc(func(m reflect.Method) bool {
			return MethodHasAnyPrefix(m, []string{"List"}) && MethodHasAnySuffix(m, []string{"Tables"}) || MethodHasAnyPrefix(m, []string{"NewPager"})
		}),
		WithExtraImports(func(m reflect.Method) []string { return []string{"net/http"} }))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	fileName := path.Join(dir, "interfaces", "interfaces.go")
	f, err := os.Open(fileName)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			t.Fatalf("failed to close file: %v", err)
		}
	}()
	b, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	if diff := cmp.Diff(string(b), wantOutput); diff != "" {
		t.Errorf("unexpected diff (-got +want):\n%s", diff)
	}
}

func TestGenerateMultipleClientsSamePackage(t *testing.T) {
	dir := t.TempDir()
	err := Generate([]any{&Client{}, &Client2{}}, dir,
		WithIncludeFunc(func(m reflect.Method) bool {
			return MethodHasAnyPrefix(m, []string{"ListTables"}) || MethodHasAnyPrefix(m, []string{"NewPager"})
		}),
		WithExtraImports(func(m reflect.Method) []string { return []string{"net/http"} }))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	fileName := path.Join(dir, "interfaces", "interfaces.go")
	f, err := os.Open(fileName)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			t.Fatalf("failed to close file: %v", err)
		}
	}()
	b, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	if diff := cmp.Diff(string(b), wantOutputMultipleClients); diff != "" {
		t.Errorf("unexpected diff (-got +want):\n%s", diff)
	}
}

func TestGenerateNonV1(t *testing.T) {
	dir := t.TempDir()
	err := Generate([]any{&armappconfiguration.ConfigurationStoresClient{}}, dir,
		WithIncludeFunc(func(m reflect.Method) bool {
			return MethodHasAnyPrefix(m, []string{"NewListPager"})
		}),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	fileName := path.Join(dir, "armappconfiguration", "armappconfiguration.go")
	f, err := os.Open(fileName)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			t.Fatalf("failed to close file: %v", err)
		}
	}()
	b, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	if diff := cmp.Diff(string(b), wantOutputNonV1); diff != "" {
		t.Errorf("unexpected diff (-got +want):\n%s", diff)
	}
}

func TestGenerateOverlappingPackageNames(t *testing.T) {
	dir := t.TempDir()
	err := Generate([]any{
		&armappconfiguration.ConfigurationStoresClient{},
		&resourcesarmmanagedapplications.ApplicationClient{},
		&solutionsarmmanagedapplications.ApplicationClient{},
	}, dir,
		WithIncludeFunc(func(m reflect.Method) bool {
			return MethodHasAnyPrefix(m, []string{"NewListPager"})
		}),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for filename, want := range wantOutputOverlappingPackages {
		fileName := path.Join(dir, filename)
		f, err := os.Open(fileName)
		if err != nil {
			t.Fatalf("failed to open file: %v", err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				t.Fatalf("failed to close file: %v", err)
			}
		}()
		b, err := io.ReadAll(f)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if diff := cmp.Diff(string(b), want); diff != "" {
			t.Errorf("unexpected diff (-got +want):\n%s", diff)
		}
	}
}
