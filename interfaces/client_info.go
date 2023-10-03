package interfaces

import (
	"path"
	"reflect"
	"regexp"
	"strings"

	"github.com/jpillora/longestcommon"
)

type clientInfo struct {
	Import       string
	ClientName   string
	Signatures   []string
	ExtraImports []string
}

func getClientInfo(client any, opts *Options) clientInfo {
	v := reflect.ValueOf(client)
	t := v.Type()
	pkgPath := t.Elem().PkgPath()
	clientName := t.Elem().Name()
	signatures := make([]string, 0)
	extraImports := make([]string, 0)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		if opts.ShouldInclude(method) {
			sig := signature(method.Name, v.Method(i).Interface())
			signatures = append(signatures, sig)
		}
		extraImports = append(extraImports, opts.ExtraImports(method)...)
	}
	return clientInfo{
		Import:       pkgPath,
		ClientName:   clientName,
		Signatures:   signatures,
		ExtraImports: extraImports,
	}
}

func getPackageNames(clientInfos []clientInfo) []string {
	versionPattern := regexp.MustCompile(`/v\d+$`)
	allImports := make([]string, len(clientInfos))
	for i, clientInfo := range clientInfos {
		allImports[i] = clientInfo.Import
	}
	// To get the shortest possible package name without collisions, we need to find the longest common prefix
	importPrefix := longestcommon.Prefix(allImports)
	packageNames := make([]string, len(clientInfos))
	for i, clientInfo := range clientInfos {
		var pkgName string
		if clientInfo.Import == importPrefix {
			pkgName = versionPattern.ReplaceAllString(clientInfo.Import, "")
			pkgName = path.Base(pkgName)
		} else {
			pkgName = strings.TrimPrefix(clientInfo.Import, importPrefix)
			pkgName = strings.ReplaceAll(versionPattern.ReplaceAllString(pkgName, ""), "/", "_")
		}

		packageNames[i] = strings.ReplaceAll(pkgName, "-", "")
	}
	return packageNames
}

func (c clientInfo) templateData(singlePackageMode bool) clientTemplateData {
	var packageName string
	if singlePackageMode {
		packageName = getPackageNames([]clientInfo{c})[0]
	}
	return clientTemplateData{
		packageName: packageName,
		Name:        c.ClientName,
		Signatures:  c.Signatures,
	}
}
