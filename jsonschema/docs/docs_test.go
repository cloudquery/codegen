package docs

import (
	"embed"
	"encoding/json"
	"strings"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/invopop/jsonschema"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/*.json
var schemaFS embed.FS

func normalizeContent(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "\n")
	return s
}

func genSnapshot(t *testing.T, fileName string) {
	data, err := schemaFS.ReadFile(fileName)
	require.NoError(t, err)

	doc, err := Generate(data, 1)
	require.NoError(t, err)

	cupaloy.New(cupaloy.SnapshotFileExtension(".md")).SnapshotT(t, normalizeContent(doc))
}

func genSnapshotStruct(t *testing.T, fileName string) {
	data, err := schemaFS.ReadFile(fileName)
	require.NoError(t, err)

	var root jsonschema.Schema
	err = json.Unmarshal(data, &root)

	require.NoError(t, err)

	doc, err := GenerateFromSchema(root, 1)
	require.NoError(t, err)

	cupaloy.New(cupaloy.SnapshotFileExtension(".md")).SnapshotT(t, normalizeContent(doc))
}

func TestAWS(t *testing.T) {
	genSnapshot(t, "testdata/aws.json")
	genSnapshotStruct(t, "testdata/aws.json")
}

func TestGCP(t *testing.T) {
	genSnapshot(t, "testdata/gcp.json")
	genSnapshotStruct(t, "testdata/gcp.json")
}

func TestClickHouse(t *testing.T) {
	genSnapshot(t, "testdata/clickhouse.json")
	genSnapshotStruct(t, "testdata/clickhouse.json")

}

func TestFiletypes(t *testing.T) {
	genSnapshot(t, "testdata/filetypes.json")
	genSnapshotStruct(t, "testdata/filetypes.json")

}

func TestFileDestination(t *testing.T) {
	genSnapshot(t, "testdata/file-destination.json")
	genSnapshotStruct(t, "testdata/file-destination.json")

}
