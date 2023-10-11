package docs

import (
	_ "embed"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/schema.json
var jsonSchema []byte

func TestGenerate(t *testing.T) {
	doc, err := Generate(jsonSchema, 0)
	require.NoError(t, err)

	cupaloy.New(
		cupaloy.SnapshotFileExtension(".md"),
		cupaloy.SnapshotSubdirectory("testdata"),
	).SnapshotT(t, doc)
}
