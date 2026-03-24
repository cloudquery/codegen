// flatten generates an optimized Go struct from a source file.
//
// Usage:
//
//	//go:generate go run github.com/cloudquery/codegen/cmd/flatten -source=../api/models.go -from=Finding -to=flatFinding -output=finding_generated.go -package=services -sort
//	//go:generate go run github.com/cloudquery/codegen/cmd/flatten -source=../api/models.go -from=Finding -to=flatFinding -output=finding_generated.go -package=services -sort -extra=AssetType:string:assetType,omitempty
package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/cloudquery/codegen/structs"
)

func main() {
	source := flag.String("source", "", "path to Go source file containing the struct")
	from := flag.String("from", "", "name of the source struct")
	to := flag.String("to", "", "name of the generated struct")
	output := flag.String("output", "", "output filename")
	pkg := flag.String("package", "", "Go package name for the generated file")
	sort := flag.Bool("sort", false, "sort fields: ID first, then alphabetically")
	extra := flag.String("extra", "", "extra fields to prepend (Name:Type:JSONTag)")
	flag.Parse()

	if *source == "" || *from == "" || *to == "" || *output == "" {
		flag.Usage()
		log.Fatal("required flags: -source, -from, -to, -output")
	}

	cfg := structs.StructConfig{
		SourceName:  *from,
		OutputName:  *to,
		OutputFile:  filepath.Base(*output),
		ExtraFields: parseExtraFields(*extra),
	}

	outputDir := filepath.Dir(*output)
	if outputDir == "" || outputDir == "." {
		outputDir = "."
	}

	var opts []structs.Option
	if *pkg != "" {
		opts = append(opts, structs.WithPackageName(*pkg))
	}
	if *sort {
		opts = append(opts, structs.WithSortFields())
	}

	if err := structs.Flatten(*source, []structs.StructConfig{cfg}, outputDir, opts...); err != nil {
		log.Fatal(err)
	}
}

func parseExtraFields(s string) []structs.Field {
	if s == "" {
		return nil
	}
	parts := strings.SplitN(s, ":", 3)
	if len(parts) != 3 {
		log.Fatalf("invalid -extra %q: expected Name:Type:JSONTag", s)
	}
	return []structs.Field{{
		Name:    strings.TrimSpace(parts[0]),
		Type:    strings.TrimSpace(parts[1]),
		JSONTag: strings.TrimSpace(parts[2]),
	}}
}

func init() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage: flatten [flags]\n\nGenerates an optimized Go struct with nested fields replaced by map[string]any.\n\nFlags:\n")
		flag.PrintDefaults()
	}
}
