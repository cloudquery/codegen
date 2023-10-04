module github.com/cloudquery/codegen

go 1.21.1

require (
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v2 v2.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications v1.1.1
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications v1.1.1
	github.com/cloudquery/plugin-sdk/v4 v4.12.3
	github.com/google/go-cmp v0.5.9
	github.com/invopop/jsonschema v0.11.0
	github.com/jpillora/longestcommon v0.0.0-20161227235612-adb9d91ee629
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9
)

// https://github.com/invopop/jsonschema/pull/106
replace github.com/invopop/jsonschema => github.com/cloudquery/jsonschema v0.0.0-20231004102900-26eed64ef87a

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.7.2 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.3.0 // indirect
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/wk8/go-ordered-map/v2 v2.1.8 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
