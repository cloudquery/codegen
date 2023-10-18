# Table of contents

* [`Spec`](#Spec)
  * [`Account`](#Account)
  * [`Org`](#Org)
  * [`TableOptions`](#TableOptions)
    * [`AccessAnalyzerFindings`](#AccessAnalyzerFindings)
      * [`CustomAccessAnalyzerListFindingsInput`](#CustomAccessAnalyzerListFindingsInput)
        * [`Criterion`](#Criterion)
        * [`SortCriteria`](#SortCriteria)
    * [`CloudtrailEvents`](#CloudtrailEvents)
      * [`CustomCloudtrailLookupEventsInput`](#CustomCloudtrailLookupEventsInput)
        * [`LookupAttribute`](#LookupAttribute)
    * [`CloudwatchMetrics`](#CloudwatchMetrics)
      * [`CloudwatchMetric`](#CloudwatchMetric)
        * [`CloudwatchListMetricsInput`](#CloudwatchListMetricsInput)
          * [`DimensionFilter`](#DimensionFilter)
        * [`CloudwatchGetMetricStatisticsInput`](#CloudwatchGetMetricStatisticsInput)
    * [`CostExplorerAPIs`](#CostExplorerAPIs)
      * [`CustomGetCostAndUsageInput`](#CustomGetCostAndUsageInput)
        * [`DateInterval`](#DateInterval)
        * [`Expression`](#Expression)
          * [`CostCategoryValues`](#CostCategoryValues)
          * [`DimensionValues`](#DimensionValues)
          * [`TagValues`](#TagValues)
        * [`GroupDefinition`](#GroupDefinition)
    * [`ECSTasks`](#ECSTasks)
      * [`CustomECSListTasksInput`](#CustomECSListTasksInput)
    * [`Inspector2Findings`](#Inspector2Findings)
      * [`CustomInspector2ListFindingsInput`](#CustomInspector2ListFindingsInput)
        * [`FilterCriteria`](#FilterCriteria)
          * [`StringFilter`](#StringFilter)
          * [`DateFilter`](#DateFilter)
          * [`NumberFilter`](#NumberFilter)
          * [`PortRangeFilter`](#PortRangeFilter)
          * [`MapFilter`](#MapFilter)
          * [`PackageFilter`](#PackageFilter)
        * [`SortCriteria`](#SortCriteria-1)
    * [`SecurityHubFindings`](#SecurityHubFindings)
      * [`CustomSecurityHubGetFindingsInput`](#CustomSecurityHubGetFindingsInput)
        * [`AwsSecurityFindingFilters`](#AwsSecurityFindingFilters)
          * [`StringFilter`](#StringFilter-1)
          * [`NumberFilter`](#NumberFilter-1)
          * [`DateFilter`](#DateFilter-1)
            * [`DateRange`](#DateRange)
          * [`KeywordFilter`](#KeywordFilter)
          * [`IpFilter`](#IpFilter)
          * [`MapFilter`](#MapFilter-1)
          * [`BooleanFilter`](#BooleanFilter)
        * [`SortCriterion`](#SortCriterion)
  * [`EventBasedSync`](#EventBasedSync)
  * [`Strategy`](#Strategy)

## <a name="Spec"></a>Spec

* `regions` (`[]string`) (nullable)

* `accounts` ([`[]Account`](#Account)) (nullable)

* `org` ([`Org`](#Org)) (nullable)

* `aws_debug` (`boolean`)

* `max_retries` (`integer`) (nullable) (default: `10`)

* `max_backoff` (`integer`) (nullable) (default: `30`)

* `custom_endpoint_url` (`string`)

* `custom_endpoint_hostname_immutable` (`boolean`) (nullable)

* `custom_endpoint_partition_id` (`string`)

* `custom_endpoint_signing_region` (`string`)

* `initialization_concurrency` (`integer`) (range: `[1,+∞)`) (default: `4`)

* `concurrency` (`integer`) (range: `[1,+∞)`) (default: `50000`)

* `use_paid_apis` (`boolean`) (default: `false`)

* `table_options` ([`TableOptions`](#TableOptions)) (nullable)

* `event_based_sync` ([`EventBasedSync`](#EventBasedSync)) (nullable)

* `scheduler` ([`Strategy`](#Strategy))

### <a name="Account"></a>Account

* `id` (`string`) (required)

* `account_name` (`string`)

* `local_profile` (`string`)

* `role_arn` (`string`) ([pattern](https://json-schema.org/draft/2020-12/json-schema-validation#section-6.3.3): `^arn(:[^:\n]*){5}([:/].*)?$`)

* `role_session_name` (`string`)

* `external_id` (`string`)

* `default_region` (`string`)

* `regions` (`[]string`) (nullable)

### <a name="Org"></a>Org

* `admin_account` ([`Account`](#Account)) (nullable)

* `member_trusted_principal` ([`Account`](#Account)) (nullable)

* `member_role_name` (`string`) (required)

* `member_role_session_name` (`string`)

* `member_external_id` (`string`)

* `member_regions` (`[]string`) (nullable)

* `organization_units` (`[]string`) (nullable)

* `skip_organization_units` (`[]string`) (nullable)

* `skip_member_accounts` (`[]string`) (nullable)

### <a name="TableOptions"></a>TableOptions

* `aws_accessanalyzer_analyzer_findings` ([`AccessAnalyzerFindings`](#AccessAnalyzerFindings)) (nullable)

* `aws_cloudtrail_events` ([`CloudtrailEvents`](#CloudtrailEvents)) (nullable)

* `aws_alpha_cloudwatch_metrics` ([`CloudwatchMetrics`](#CloudwatchMetrics)) (nullable)

* `aws_alpha_costexplorer_cost_custom` ([`CostExplorerAPIs`](#CostExplorerAPIs)) (nullable)

* `aws_ecs_cluster_tasks` ([`ECSTasks`](#ECSTasks)) (nullable)

* `aws_inspector2_findings` ([`Inspector2Findings`](#Inspector2Findings)) (nullable)

* `aws_securityhub_findings` ([`SecurityHubFindings`](#SecurityHubFindings)) (nullable)

#### <a name="AccessAnalyzerFindings"></a>AccessAnalyzerFindings

* `list_findings` ([`[]CustomAccessAnalyzerListFindingsInput`](#CustomAccessAnalyzerListFindingsInput)) (nullable)

##### <a name="CustomAccessAnalyzerListFindingsInput"></a>CustomAccessAnalyzerListFindingsInput

* `Filter` ([`map[string]Criterion`](#Criterion)) (nullable)

* `MaxResults` (`integer`) (nullable)

* `Sort` ([`SortCriteria`](#SortCriteria)) (nullable)

###### <a name="Criterion"></a>Criterion

* `Contains` (`[]string`) (nullable)

* `Eq` (`[]string`) (nullable)

* `Exists` (`boolean`) (nullable)

* `Neq` (`[]string`) (nullable)

###### <a name="SortCriteria"></a>SortCriteria

* `AttributeName` (`string`) (nullable)

* `OrderBy` (`string`)

#### <a name="CloudtrailEvents"></a>CloudtrailEvents

* `lookup_events` ([`[]CustomCloudtrailLookupEventsInput`](#CustomCloudtrailLookupEventsInput)) (nullable)

##### <a name="CustomCloudtrailLookupEventsInput"></a>CustomCloudtrailLookupEventsInput

* `EndTime` (`string`) (nullable) ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `date-time`)

* `EventCategory` (`string`)

* `LookupAttributes` ([`[]LookupAttribute`](#LookupAttribute)) (nullable)

* `MaxResults` (`integer`) (nullable)

* `StartTime` (`string`) (nullable) ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `date-time`)

###### <a name="LookupAttribute"></a>LookupAttribute

* `AttributeKey` (`string`)

* `AttributeValue` (`string`) (nullable)

#### <a name="CloudwatchMetrics"></a>CloudwatchMetrics

([`[]CloudwatchMetric`](#CloudwatchMetric))

##### <a name="CloudwatchMetric"></a>CloudwatchMetric

* `list_metrics` ([`CloudwatchListMetricsInput`](#CloudwatchListMetricsInput))

* `get_metric_statistics` ([`[]CloudwatchGetMetricStatisticsInput`](#CloudwatchGetMetricStatisticsInput)) (nullable)

###### <a name="CloudwatchListMetricsInput"></a>CloudwatchListMetricsInput

* `Dimensions` ([`[]DimensionFilter`](#DimensionFilter)) (nullable)

* `IncludeLinkedAccounts` (`boolean`)

* `MetricName` (`string`) (nullable)

* `Namespace` (`string`) (nullable)

* `OwningAccount` (`string`) (nullable)

* `RecentlyActive` (`string`)

###### <a name="DimensionFilter"></a>DimensionFilter

* `Name` (`string`) (nullable)

* `Value` (`string`) (nullable)

###### <a name="CloudwatchGetMetricStatisticsInput"></a>CloudwatchGetMetricStatisticsInput

* `EndTime` (`string`) (nullable) ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `date-time`)

* `Period` (`integer`) (nullable)

* `StartTime` (`string`) (nullable) ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `date-time`)

* `ExtendedStatistics` (`[]string`) (nullable)

* `Statistics` (`[]string`) (nullable)

* `Unit` (`string`)

#### <a name="CostExplorerAPIs"></a>CostExplorerAPIs

* `get_cost_and_usage` ([`[]CustomGetCostAndUsageInput`](#CustomGetCostAndUsageInput)) (nullable)

##### <a name="CustomGetCostAndUsageInput"></a>CustomGetCostAndUsageInput

* `Granularity` (`string`)

* `Metrics` (`[]string`) (nullable)

* `TimePeriod` ([`DateInterval`](#DateInterval)) (nullable)

* `Filter` ([`Expression`](#Expression)) (nullable)

* `GroupBy` ([`[]GroupDefinition`](#GroupDefinition)) (nullable)

###### <a name="DateInterval"></a>DateInterval

* `End` (`string`) (nullable)

* `Start` (`string`) (nullable)

###### <a name="Expression"></a>Expression

* `And` ([`[]Expression`](#Expression)) (nullable)

* `CostCategories` ([`CostCategoryValues`](#CostCategoryValues)) (nullable)

* `Dimensions` ([`DimensionValues`](#DimensionValues)) (nullable)

* `Not` ([`Expression`](#Expression)) (nullable)

* `Or` ([`[]Expression`](#Expression)) (nullable)

* `Tags` ([`TagValues`](#TagValues)) (nullable)

###### <a name="CostCategoryValues"></a>CostCategoryValues

* `Key` (`string`) (nullable)

* `MatchOptions` (`[]string`) (nullable)

* `Values` (`[]string`) (nullable)

###### <a name="DimensionValues"></a>DimensionValues

* `Key` (`string`)

* `MatchOptions` (`[]string`) (nullable)

* `Values` (`[]string`) (nullable)

###### <a name="TagValues"></a>TagValues

* `Key` (`string`) (nullable)

* `MatchOptions` (`[]string`) (nullable)

* `Values` (`[]string`) (nullable)

###### <a name="GroupDefinition"></a>GroupDefinition

* `Key` (`string`) (nullable)

* `Type` (`string`)

#### <a name="ECSTasks"></a>ECSTasks

* `list_tasks` ([`[]CustomECSListTasksInput`](#CustomECSListTasksInput)) (nullable)

##### <a name="CustomECSListTasksInput"></a>CustomECSListTasksInput

* `ContainerInstance` (`string`) (nullable)

* `DesiredStatus` (`string`)

* `Family` (`string`) (nullable)

* `LaunchType` (`string`)

* `MaxResults` (`integer`) (nullable) (range: `[1,100]`) (default: `100`)

* `ServiceName` (`string`) (nullable)

* `StartedBy` (`string`) (nullable)

#### <a name="Inspector2Findings"></a>Inspector2Findings

* `list_findings` ([`[]CustomInspector2ListFindingsInput`](#CustomInspector2ListFindingsInput)) (nullable)

##### <a name="CustomInspector2ListFindingsInput"></a>CustomInspector2ListFindingsInput

* `FilterCriteria` ([`FilterCriteria`](#FilterCriteria)) (nullable)

* `MaxResults` (`integer`) (nullable)

* `SortCriteria` ([`SortCriteria`](#SortCriteria-1)) (nullable)

###### <a name="FilterCriteria"></a>FilterCriteria

* `AwsAccountId` ([`[]StringFilter`](#StringFilter)) (nullable)

* `CodeVulnerabilityDetectorName` ([`[]StringFilter`](#StringFilter)) (nullable)

* `CodeVulnerabilityDetectorTags` ([`[]StringFilter`](#StringFilter)) (nullable)

* `CodeVulnerabilityFilePath` ([`[]StringFilter`](#StringFilter)) (nullable)

* `ComponentId` ([`[]StringFilter`](#StringFilter)) (nullable)

* `ComponentType` ([`[]StringFilter`](#StringFilter)) (nullable)

* `Ec2InstanceImageId` ([`[]StringFilter`](#StringFilter)) (nullable)

* `Ec2InstanceSubnetId` ([`[]StringFilter`](#StringFilter)) (nullable)

* `Ec2InstanceVpcId` ([`[]StringFilter`](#StringFilter)) (nullable)

* `EcrImageArchitecture` ([`[]StringFilter`](#StringFilter)) (nullable)

* `EcrImageHash` ([`[]StringFilter`](#StringFilter)) (nullable)

* `EcrImagePushedAt` ([`[]DateFilter`](#DateFilter)) (nullable)

* `EcrImageRegistry` ([`[]StringFilter`](#StringFilter)) (nullable)

* `EcrImageRepositoryName` ([`[]StringFilter`](#StringFilter)) (nullable)

* `EcrImageTags` ([`[]StringFilter`](#StringFilter)) (nullable)

* `EpssScore` ([`[]NumberFilter`](#NumberFilter)) (nullable)

* `ExploitAvailable` ([`[]StringFilter`](#StringFilter)) (nullable)

* `FindingArn` ([`[]StringFilter`](#StringFilter)) (nullable)

* `FindingStatus` ([`[]StringFilter`](#StringFilter)) (nullable)

* `FindingType` ([`[]StringFilter`](#StringFilter)) (nullable)

* `FirstObservedAt` ([`[]DateFilter`](#DateFilter)) (nullable)

* `FixAvailable` ([`[]StringFilter`](#StringFilter)) (nullable)

* `InspectorScore` ([`[]NumberFilter`](#NumberFilter)) (nullable)

* `LambdaFunctionExecutionRoleArn` ([`[]StringFilter`](#StringFilter)) (nullable)

* `LambdaFunctionLastModifiedAt` ([`[]DateFilter`](#DateFilter)) (nullable)

* `LambdaFunctionLayers` ([`[]StringFilter`](#StringFilter)) (nullable)

* `LambdaFunctionName` ([`[]StringFilter`](#StringFilter)) (nullable)

* `LambdaFunctionRuntime` ([`[]StringFilter`](#StringFilter)) (nullable)

* `LastObservedAt` ([`[]DateFilter`](#DateFilter)) (nullable)

* `NetworkProtocol` ([`[]StringFilter`](#StringFilter)) (nullable)

* `PortRange` ([`[]PortRangeFilter`](#PortRangeFilter)) (nullable)

* `RelatedVulnerabilities` ([`[]StringFilter`](#StringFilter)) (nullable)

* `ResourceId` ([`[]StringFilter`](#StringFilter)) (nullable)

* `ResourceTags` ([`[]MapFilter`](#MapFilter)) (nullable)

* `ResourceType` ([`[]StringFilter`](#StringFilter)) (nullable)

* `Severity` ([`[]StringFilter`](#StringFilter)) (nullable)

* `Title` ([`[]StringFilter`](#StringFilter)) (nullable)

* `UpdatedAt` ([`[]DateFilter`](#DateFilter)) (nullable)

* `VendorSeverity` ([`[]StringFilter`](#StringFilter)) (nullable)

* `VulnerabilityId` ([`[]StringFilter`](#StringFilter)) (nullable)

* `VulnerabilitySource` ([`[]StringFilter`](#StringFilter)) (nullable)

* `VulnerablePackages` ([`[]PackageFilter`](#PackageFilter)) (nullable)

###### <a name="StringFilter"></a>StringFilter

* `Comparison` (`string`)

* `Value` (`string`) (nullable)

###### <a name="DateFilter"></a>DateFilter

* `EndInclusive` (`string`) (nullable) ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `date-time`)

* `StartInclusive` (`string`) (nullable) ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `date-time`)

###### <a name="NumberFilter"></a>NumberFilter

* `LowerInclusive` (`number`) (nullable)

* `UpperInclusive` (`number`) (nullable)

###### <a name="PortRangeFilter"></a>PortRangeFilter

* `BeginInclusive` (`integer`) (nullable)

* `EndInclusive` (`integer`) (nullable)

###### <a name="MapFilter"></a>MapFilter

* `Comparison` (`string`)

* `Key` (`string`) (nullable)

* `Value` (`string`) (nullable)

###### <a name="PackageFilter"></a>PackageFilter

* `Architecture` ([`StringFilter`](#StringFilter)) (nullable)

* `Epoch` ([`NumberFilter`](#NumberFilter)) (nullable)

* `Name` ([`StringFilter`](#StringFilter)) (nullable)

* `Release` ([`StringFilter`](#StringFilter)) (nullable)

* `SourceLambdaLayerArn` ([`StringFilter`](#StringFilter)) (nullable)

* `SourceLayerHash` ([`StringFilter`](#StringFilter)) (nullable)

* `Version` ([`StringFilter`](#StringFilter)) (nullable)

###### <a name="SortCriteria-1"></a>SortCriteria

* `Field` (`string`)

* `SortOrder` (`string`)

#### <a name="SecurityHubFindings"></a>SecurityHubFindings

* `get_findings` ([`[]CustomSecurityHubGetFindingsInput`](#CustomSecurityHubGetFindingsInput)) (nullable)

##### <a name="CustomSecurityHubGetFindingsInput"></a>CustomSecurityHubGetFindingsInput

* `Filters` ([`AwsSecurityFindingFilters`](#AwsSecurityFindingFilters)) (nullable)

* `MaxResults` (`integer`) (range: `[1,100]`) (default: `100`)

* `SortCriteria` ([`[]SortCriterion`](#SortCriterion)) (nullable)

###### <a name="AwsSecurityFindingFilters"></a>AwsSecurityFindingFilters

* `AwsAccountId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `CompanyName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ComplianceAssociatedStandardsId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ComplianceSecurityControlId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ComplianceStatus` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `Confidence` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

* `CreatedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

* `Criticality` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

* `Description` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `FindingProviderFieldsConfidence` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

* `FindingProviderFieldsCriticality` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

* `FindingProviderFieldsRelatedFindingsId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `FindingProviderFieldsRelatedFindingsProductArn` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `FindingProviderFieldsSeverityLabel` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `FindingProviderFieldsSeverityOriginal` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `FindingProviderFieldsTypes` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `FirstObservedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

* `GeneratorId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `Id` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `Keyword` ([`[]KeywordFilter`](#KeywordFilter)) (nullable)

* `LastObservedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

* `MalwareName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `MalwarePath` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `MalwareState` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `MalwareType` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `NetworkDestinationDomain` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `NetworkDestinationIpV4` ([`[]IpFilter`](#IpFilter)) (nullable)

* `NetworkDestinationIpV6` ([`[]IpFilter`](#IpFilter)) (nullable)

* `NetworkDestinationPort` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

* `NetworkDirection` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `NetworkProtocol` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `NetworkSourceDomain` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `NetworkSourceIpV4` ([`[]IpFilter`](#IpFilter)) (nullable)

* `NetworkSourceIpV6` ([`[]IpFilter`](#IpFilter)) (nullable)

* `NetworkSourceMac` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `NetworkSourcePort` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

* `NoteText` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `NoteUpdatedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

* `NoteUpdatedBy` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ProcessLaunchedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

* `ProcessName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ProcessParentPid` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

* `ProcessPath` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ProcessPid` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

* `ProcessTerminatedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

* `ProductArn` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ProductFields` ([`[]MapFilter`](#MapFilter-1)) (nullable)

* `ProductName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `RecommendationText` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `RecordState` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `Region` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `RelatedFindingsId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `RelatedFindingsProductArn` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceAwsEc2InstanceIamInstanceProfileArn` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceAwsEc2InstanceImageId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceAwsEc2InstanceIpV4Addresses` ([`[]IpFilter`](#IpFilter)) (nullable)

* `ResourceAwsEc2InstanceIpV6Addresses` ([`[]IpFilter`](#IpFilter)) (nullable)

* `ResourceAwsEc2InstanceKeyName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceAwsEc2InstanceLaunchedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

* `ResourceAwsEc2InstanceSubnetId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceAwsEc2InstanceType` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceAwsEc2InstanceVpcId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceAwsIamAccessKeyCreatedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

* `ResourceAwsIamAccessKeyPrincipalName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceAwsIamAccessKeyStatus` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceAwsIamAccessKeyUserName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceAwsIamUserUserName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceAwsS3BucketOwnerId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceAwsS3BucketOwnerName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceContainerImageId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceContainerImageName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceContainerLaunchedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

* `ResourceContainerName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceDetailsOther` ([`[]MapFilter`](#MapFilter-1)) (nullable)

* `ResourceId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourcePartition` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceRegion` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ResourceTags` ([`[]MapFilter`](#MapFilter-1)) (nullable)

* `ResourceType` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `Sample` ([`[]BooleanFilter`](#BooleanFilter)) (nullable)

* `SeverityLabel` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `SeverityNormalized` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

* `SeverityProduct` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

* `SourceUrl` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ThreatIntelIndicatorCategory` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ThreatIntelIndicatorLastObservedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

* `ThreatIntelIndicatorSource` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ThreatIntelIndicatorSourceUrl` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ThreatIntelIndicatorType` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `ThreatIntelIndicatorValue` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `Title` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `Type` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `UpdatedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

* `UserDefinedFields` ([`[]MapFilter`](#MapFilter-1)) (nullable)

* `VerificationState` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `WorkflowState` ([`[]StringFilter`](#StringFilter-1)) (nullable)

* `WorkflowStatus` ([`[]StringFilter`](#StringFilter-1)) (nullable)

###### <a name="StringFilter-1"></a>StringFilter

* `Comparison` (`string`)

* `Value` (`string`) (nullable)

###### <a name="NumberFilter-1"></a>NumberFilter

* `Eq` (`number`)

* `Gte` (`number`)

* `Lte` (`number`)

###### <a name="DateFilter-1"></a>DateFilter

* `DateRange` ([`DateRange`](#DateRange)) (nullable)

* `End` (`string`) (nullable)

* `Start` (`string`) (nullable)

###### <a name="DateRange"></a>DateRange

* `Unit` (`string`)

* `Value` (`integer`)

###### <a name="KeywordFilter"></a>KeywordFilter

* `Value` (`string`) (nullable)

###### <a name="IpFilter"></a>IpFilter

* `Cidr` (`string`) (nullable)

###### <a name="MapFilter-1"></a>MapFilter

* `Comparison` (`string`)

* `Key` (`string`) (nullable)

* `Value` (`string`) (nullable)

###### <a name="BooleanFilter"></a>BooleanFilter

* `Value` (`boolean`)

###### <a name="SortCriterion"></a>SortCriterion

* `Field` (`string`) (nullable)

* `SortOrder` (`string`)

### <a name="EventBasedSync"></a>EventBasedSync

* `full_sync` (`boolean`) (nullable) (default: `true`)

* `account` ([`Account`](#Account))

* `kinesis_stream_arn` (`string`) (required) ([pattern](https://json-schema.org/draft/2020-12/json-schema-validation#section-6.3.3): `^arn(:[^:\n]*){5}([:/].*)?$`)

* `start_time` (`string`) (nullable) ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `date-time`) (default: `now`)

### <a name="Strategy"></a>Strategy

CloudQuery scheduling strategy

(`string`) (possible values: `dfs`, `round-robin`, `shuffle`) (default: `dfs`)
