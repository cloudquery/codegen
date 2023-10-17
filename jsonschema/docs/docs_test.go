//go:build !windows
// +build !windows

package docs

import (
	"embed"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/*.json
var schemaFS embed.FS

func genSnapshot(t *testing.T, fileName string) {
	data, err := schemaFS.ReadFile(fileName)
	require.NoError(t, err)

	doc, err := Generate(data, 1)
	require.NoError(t, err)

	cupaloy.New(cupaloy.SnapshotFileExtension(".md")).SnapshotT(t, doc)
}

func TestAWS(t *testing.T) {
	genSnapshot(t, "testdata/aws.json")
}

func TestClickHouse(t *testing.T) {
	genSnapshot(t, "testdata/clickhouse.json")
}
