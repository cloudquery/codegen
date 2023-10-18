# Table of contents

* [`Spec`](#Spec)
  * [`Strategy`](#Strategy)
  * [`CredentialsConfig`](#CredentialsConfig)

## <a name="Spec"></a>Spec

  Spec defines GCP source plugin Spec

* `project_ids` (`[]string`) (nullable)

  Specify projects to connect to.
  If either `folder_ids` or `project_filter` is specified,
  these projects will be synced in addition to the projects from the folder/filter.
  
  Empty or `null` value will use all projects available to the current authenticated account.

* `folder_ids` (`[]string`) (nullable) ([pattern](https://json-schema.org/draft/2020-12/json-schema-validation#section-6.3.3): `^(folders|organizations)/(.)+$`)

  CloudQuery will sync from all the projects in the specified folders, recursively.
  `folder_ids` must be of the format `folders/<folder_id>` or `organizations/<organization_id>`.
  This feature requires the `resourcemanager.folders.list` permission.
  
  By default, CloudQuery will also sync from sub-folders recursively (up to depth `100`).
  To reduce this, set `folder_recursion_depth` to a lower value (or to `0` to disable recursion completely).
  
  Mutually exclusive with `project_filter`.

* `folder_recursion_depth` (`integer`) (nullable) (range: `[0,+∞)`) (default: `100`)

  The maximum depth to recurse into sub-folders.
  `0` means no recursion (only the top-level projects in folders will be used for sync).

* `organization_ids` (`[]string`) (nullable)

  Specify organizations to use when syncing organization level resources (e.g.
  [folders](https://github.com/cloudquery/cloudquery/blob/0e384a84d1c9545b24c2eda9af00f111bab79c36/plugins/source/gcp/resources/services/resourcemanager/folders_fetch.go#L23)
  or
  [security findings](https://github.com/cloudquery/cloudquery/blob/0e384a84d1c9545b24c2eda9af00f111bab79c36/plugins/source/gcp/resources/services/securitycenter/organization_findings.go#L43)).
  
  If `organization_filter` is specified, these organizations will be used in addition to the organizations from the filter.
  
  Empty or `null` value will use all organizations available to the current authenticated account).

* `project_filter` (`string`)

  A filter to determine the projects that are synced, mutually exclusive with `folder_ids`.
  
  For instance, to only sync projects where the name starts with `how-`, set `project_filter` to `name:how-*`.
  
  More examples:
  
  - `"name:how-* OR name:test-*"` matches projects starting with `how-` or `test-`
  - `"NOT name:test-*"` matches all projects _not_ starting with `test-`
  
  For syntax and example queries refer to API References
  [here](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list#google.cloudresourcemanager.v1.Projects.ListProjects)
  and
  [here](https://cloud.google.com/sdk/gcloud/reference/topic/filters).

* `organization_filter` (`string`)

  A filter to determine the organizations to use when syncing organization level resources (e.g.
  [folders](https://github.com/cloudquery/cloudquery/blob/0e384a84d1c9545b24c2eda9af00f111bab79c36/plugins/source/gcp/resources/services/resourcemanager/folders_fetch.go#L23)
  or
  [security findings](https://github.com/cloudquery/cloudquery/blob/0e384a84d1c9545b24c2eda9af00f111bab79c36/plugins/source/gcp/resources/services/securitycenter/organization_findings.go#L43)).
  
  For instance, to use only organizations from the `cloudquery.io` domain, set `organization_filter` to `domain:cloudquery.io`.
  
  For syntax and example queries refer to API Reference [here](https://cloud.google.com/resource-manager/reference/rest/v1/organizations/search#google.cloudresourcemanager.v1.SearchOrganizationsRequest).

* `service_account_key_json` (`string`)

  GCP service account key content.
  
  Using service accounts is not recommended, but if it is used it is better to use
  [environment or file variable substitution](/docs/advanced-topics/environment-variable-substitution).

* `backoff_delay` (`integer`) (range: `[0,+∞)`) (default: `30`)

  If specified APIs will be retried with exponential backoff if they are rate limited.
  This is the max delay (in seconds) between retries.

* `backoff_retries` (`integer`) (range: `[0,+∞)`) (default: `0`)

  If specified APIs will be retried with exponential backoff if they are rate limited.
  This is the max number of retries.

* `enabled_services_only` (`boolean`)

  If enabled CloudQuery will skip any resources that belong to a service that has been disabled or not been enabled.
  
  If you use this option on a large organization (with more than `500` projects)
  you should also set the `backoff_retries` to a value greater than `0`, otherwise you may hit the API rate limits.
  
  In `>=v9.0.0` if an error is returned then CloudQuery will assume that all services are enabled
  and will continue to attempt to sync all specified tables rather than just ending the sync.

* `concurrency` (`integer`) (range: `[1,+∞)`) (default: `50000`)

  The best effort maximum number of Go routines to use.
  Lower this number to reduce memory usage.

* `discovery_concurrency` (`integer`) (range: `[1,+∞)`) (default: `100`)

  The number of concurrent requests that CloudQuery will make to resolve enabled services.
  This is only used when `enabled_services_only` is set to `true`.

* `scheduler` ([`Strategy`](#Strategy))

  The scheduler to use when determining the priority of resources to sync.
  
  For more information about this, see [performance tuning](/docs/advanced-topics/performance-tuning).

* `service_account_impersonation` ([`CredentialsConfig`](#CredentialsConfig)) (nullable)

  Service Account impersonation configuration.

### <a name="Strategy"></a>Strategy

CloudQuery scheduling strategy

(`string`) (possible values: `dfs`, `round-robin`, `shuffle`) (default: `dfs`)

### <a name="CredentialsConfig"></a>CredentialsConfig

* `target_principal` (`string`) (required) ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `email`)

  The email address of the service account to impersonate.

* `scopes` (`[]string`) (nullable) ([pattern](https://json-schema.org/draft/2020-12/json-schema-validation#section-6.3.3): `^https://www.googleapis.com/auth/(.)+$`) (default: `[https://www.googleapis.com/auth/cloud-platform]`)

  Scopes that the impersonated credential should have.
  
  See available scopes in the [documentation](https://developers.google.com/identity/protocols/oauth2/scopes).

* `delegates` (`[]string`) (nullable) ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `email`)

  Delegates are the service account email addresses in a delegation chain.
  Each service account must be granted `roles/iam.serviceAccountTokenCreator` on the next service account in the chain.

* `subject` (`string`)

  The subject field of a JWT (`sub`).
  This field should only be set if you wish to impersonate a user.
  This feature is useful when using domain wide delegation.
