package interfaces

import (
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"golang.org/x/exp/maps"
)

var (
	csr = caser.New()
)

type serviceTemplateData struct {
	PackageName string
	FileName    string
	Imports     []string
	Clients     []clientTemplateData
}

func getTemplateDataFromClientInfos(clientInfos []clientInfo, options *Options) []serviceTemplateData {
	packageNames := getPackageNames(clientInfos)
	services := make([]serviceTemplateData, 0)
	serviceMap := make(map[string][]clientInfo)
	for i, clientInfo := range clientInfos {
		serviceMap[packageNames[i]] = append(serviceMap[packageNames[i]], clientInfo)
	}
	for packageName, infos := range serviceMap {
		imports := make(map[string]bool)
		clientsTemplateData := make([]clientTemplateData, 0)
		for _, clientInfo := range infos {
			imports[clientInfo.Import] = true
			for _, extraImport := range clientInfo.ExtraImports {
				imports[extraImport] = true
			}
			clientsTemplateData = append(clientsTemplateData, clientInfo.templateData(len(options.SinglePackage) > 0))
		}
		svc := serviceTemplateData{
			PackageName: packageName,
			FileName:    packageName,
			Imports:     maps.Keys(imports),
			Clients:     clientsTemplateData,
		}
		if len(options.SinglePackage) > 0 {
			svc.PackageName = options.SinglePackage
		}
		services = append(services, svc)
	}
	return services
}

type clientTemplateData struct {
	packageName string
	Name        string
	Signatures  []string
}

func (c clientTemplateData) ClientName() string {
	return csr.ToPascal(c.packageName) + c.Name
}
