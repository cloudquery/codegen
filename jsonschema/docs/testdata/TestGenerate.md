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
          * [`SortCriteria`](#SortCriteria_1)
      * [`SecurityHubFindings`](#SecurityHubFindings)
        * [`CustomSecurityHubGetFindingsInput`](#CustomSecurityHubGetFindingsInput)
          * [`AwsSecurityFindingFilters`](#AwsSecurityFindingFilters)
            * [`StringFilter`](#StringFilter_1)
            * [`NumberFilter`](#NumberFilter_1)
            * [`DateFilter`](#DateFilter_1)
              * [`DateRange`](#DateRange)
            * [`KeywordFilter`](#KeywordFilter)
            * [`IpFilter`](#IpFilter)
            * [`MapFilter`](#MapFilter_1)
            * [`BooleanFilter`](#BooleanFilter)
          * [`SortCriterion`](#SortCriterion)
    * [`EventBasedSync`](#EventBasedSync)
    * [`Strategy`](#Strategy)

##Spec

<a name="Spec"></a>

* `regions` (`[]string`) (nullable)
* `accounts` ([`[]Account`](#Account)) (nullable)
* `org` ([`Org`](#Org)) (nullable)
* `aws_debug` (`boolean`)
* `max_retries` (`integer`) (nullable)
* `max_backoff` (`integer`) (nullable)
* `custom_endpoint_url` (`string`)
* `custom_endpoint_hostname_immutable` (`boolean`) (nullable)
* `custom_endpoint_partition_id` (`string`)
* `custom_endpoint_signing_region` (`string`)
* `initialization_concurrency` (`integer`) (default=`4`)
* `concurrency` (`integer`) (default=`50000`)
* `use_paid_apis` (`boolean`) (default=`false`)
* `table_options` ([`TableOptions`](#TableOptions)) (nullable)
* `event_based_sync` ([`EventBasedSync`](#EventBasedSync)) (nullable)
* `scheduler` ([`Strategy`](#Strategy))

###Account

<a name="Account"></a>

* `id` (`string`) (required)
* `account_name` (`string`)
* `local_profile` (`string`)
* `role_arn` (`string`)
* `role_session_name` (`string`)
* `external_id` (`string`)
* `default_region` (`string`)
* `regions` (`[]string`) (nullable)

###Org

<a name="Org"></a>

* `admin_account` ([`Account`](#Account)) (nullable)
* `member_trusted_principal` ([`Account`](#Account)) (nullable)
* `member_role_name` (`string`) (required)
* `member_role_session_name` (`string`)
* `member_external_id` (`string`)
* `member_regions` (`[]string`) (nullable)
* `organization_units` (`[]string`) (nullable)
* `skip_organization_units` (`[]string`) (nullable)
* `skip_member_accounts` (`[]string`) (nullable)

###TableOptions

<a name="TableOptions"></a>

* `aws_accessanalyzer_analyzer_findings` ([`AccessAnalyzerFindings`](#AccessAnalyzerFindings)) (nullable)
* `aws_cloudtrail_events` ([`CloudtrailEvents`](#CloudtrailEvents)) (nullable)
* `aws_alpha_cloudwatch_metrics` ([`CloudwatchMetrics`](#CloudwatchMetrics)) (nullable)
* `aws_alpha_costexplorer_cost_custom` ([`CostExplorerAPIs`](#CostExplorerAPIs)) (nullable)
* `aws_ecs_cluster_tasks` ([`ECSTasks`](#ECSTasks)) (nullable)
* `aws_inspector2_findings` ([`Inspector2Findings`](#Inspector2Findings)) (nullable)
* `aws_securityhub_findings` ([`SecurityHubFindings`](#SecurityHubFindings)) (nullable)

####AccessAnalyzerFindings

<a name="AccessAnalyzerFindings"></a>

* `list_findings` ([`[]CustomAccessAnalyzerListFindingsInput`](#CustomAccessAnalyzerListFindingsInput)) (nullable)

#####CustomAccessAnalyzerListFindingsInput

<a name="CustomAccessAnalyzerListFindingsInput"></a>

* `Filter` ([`map[string]Criterion`](#Criterion)) (nullable)
* `MaxResults` (`integer`) (nullable)
* `Sort` ([`SortCriteria`](#SortCriteria)) (nullable)

######Criterion

<a name="Criterion"></a>

* `Contains` (`[]string`) (nullable)
* `Eq` (`[]string`) (nullable)
* `Exists` (`boolean`) (nullable)
* `Neq` (`[]string`) (nullable)

######SortCriteria

<a name="SortCriteria"></a>

* `AttributeName` (`string`) (nullable)
* `OrderBy` (`string`)

####CloudtrailEvents

<a name="CloudtrailEvents"></a>

* `lookup_events` ([`[]CustomCloudtrailLookupEventsInput`](#CustomCloudtrailLookupEventsInput)) (nullable)

#####CustomCloudtrailLookupEventsInput

<a name="CustomCloudtrailLookupEventsInput"></a>

* `EndTime` (`string`) (nullable)
* `EventCategory` (`string`)
* `LookupAttributes` ([`[]LookupAttribute`](#LookupAttribute)) (nullable)
* `MaxResults` (`integer`) (nullable)
* `StartTime` (`string`) (nullable)

######LookupAttribute

<a name="LookupAttribute"></a>

* `AttributeKey` (`string`)
* `AttributeValue` (`string`) (nullable)

####CloudwatchMetrics

<a name="CloudwatchMetrics"></a>


####CostExplorerAPIs

<a name="CostExplorerAPIs"></a>

* `get_cost_and_usage` ([`[]CustomGetCostAndUsageInput`](#CustomGetCostAndUsageInput)) (nullable)

#####CustomGetCostAndUsageInput

<a name="CustomGetCostAndUsageInput"></a>

* `Granularity` (`string`)
* `Metrics` (`[]string`) (nullable)
* `TimePeriod` ([`DateInterval`](#DateInterval)) (nullable)
* `Filter` ([`Expression`](#Expression)) (nullable)
* `GroupBy` ([`[]GroupDefinition`](#GroupDefinition)) (nullable)

######DateInterval

<a name="DateInterval"></a>

* `End` (`string`) (nullable)
* `Start` (`string`) (nullable)

######Expression

<a name="Expression"></a>

* `And` ([`[]Expression`](#Expression)) (nullable)
* `CostCategories` ([`CostCategoryValues`](#CostCategoryValues)) (nullable)
* `Dimensions` ([`DimensionValues`](#DimensionValues)) (nullable)
* `Not` ([`Expression`](#Expression)) (nullable)
* `Or` ([`[]Expression`](#Expression)) (nullable)
* `Tags` ([`TagValues`](#TagValues)) (nullable)

######CostCategoryValues

<a name="CostCategoryValues"></a>

* `Key` (`string`) (nullable)
* `MatchOptions` (`[]string`) (nullable)
* `Values` (`[]string`) (nullable)

######DimensionValues

<a name="DimensionValues"></a>

* `Key` (`string`)
* `MatchOptions` (`[]string`) (nullable)
* `Values` (`[]string`) (nullable)

######TagValues

<a name="TagValues"></a>

* `Key` (`string`) (nullable)
* `MatchOptions` (`[]string`) (nullable)
* `Values` (`[]string`) (nullable)

######GroupDefinition

<a name="GroupDefinition"></a>

* `Key` (`string`) (nullable)
* `Type` (`string`)

####ECSTasks

<a name="ECSTasks"></a>

* `list_tasks` ([`[]CustomECSListTasksInput`](#CustomECSListTasksInput)) (nullable)

#####CustomECSListTasksInput

<a name="CustomECSListTasksInput"></a>

* `ContainerInstance` (`string`) (nullable)
* `DesiredStatus` (`string`)
* `Family` (`string`) (nullable)
* `LaunchType` (`string`)
* `MaxResults` (`integer`) (nullable) (default=`100`)
* `ServiceName` (`string`) (nullable)
* `StartedBy` (`string`) (nullable)

####Inspector2Findings

<a name="Inspector2Findings"></a>

* `list_findings` ([`[]CustomInspector2ListFindingsInput`](#CustomInspector2ListFindingsInput)) (nullable)

#####CustomInspector2ListFindingsInput

<a name="CustomInspector2ListFindingsInput"></a>

* `FilterCriteria` ([`FilterCriteria`](#FilterCriteria)) (nullable)
* `MaxResults` (`integer`) (nullable)
* `SortCriteria` ([`SortCriteria`](#SortCriteria_1)) (nullable)

######FilterCriteria

<a name="FilterCriteria"></a>

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

######StringFilter

<a name="StringFilter"></a>

* `Comparison` (`string`)
* `Value` (`string`) (nullable)

######DateFilter

<a name="DateFilter"></a>

* `EndInclusive` (`string`) (nullable)
* `StartInclusive` (`string`) (nullable)

######NumberFilter

<a name="NumberFilter"></a>

* `LowerInclusive` (`number`) (nullable)
* `UpperInclusive` (`number`) (nullable)

######PortRangeFilter

<a name="PortRangeFilter"></a>

* `BeginInclusive` (`integer`) (nullable)
* `EndInclusive` (`integer`) (nullable)

######MapFilter

<a name="MapFilter"></a>

* `Comparison` (`string`)
* `Key` (`string`) (nullable)
* `Value` (`string`) (nullable)

######PackageFilter

<a name="PackageFilter"></a>

* `Architecture` ([`StringFilter`](#StringFilter)) (nullable)
* `Epoch` ([`NumberFilter`](#NumberFilter)) (nullable)
* `Name` ([`StringFilter`](#StringFilter)) (nullable)
* `Release` ([`StringFilter`](#StringFilter)) (nullable)
* `SourceLambdaLayerArn` ([`StringFilter`](#StringFilter)) (nullable)
* `SourceLayerHash` ([`StringFilter`](#StringFilter)) (nullable)
* `Version` ([`StringFilter`](#StringFilter)) (nullable)

######SortCriteria

<a name="SortCriteria_1"></a>

* `Field` (`string`)
* `SortOrder` (`string`)

####SecurityHubFindings

<a name="SecurityHubFindings"></a>

* `get_findings` ([`[]CustomSecurityHubGetFindingsInput`](#CustomSecurityHubGetFindingsInput)) (nullable)

#####CustomSecurityHubGetFindingsInput

<a name="CustomSecurityHubGetFindingsInput"></a>

* `Filters` ([`AwsSecurityFindingFilters`](#AwsSecurityFindingFilters)) (nullable)
* `MaxResults` (`integer`)
* `SortCriteria` ([`[]SortCriterion`](#SortCriterion)) (nullable)

######AwsSecurityFindingFilters

<a name="AwsSecurityFindingFilters"></a>

* `AwsAccountId` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `CompanyName` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ComplianceAssociatedStandardsId` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ComplianceSecurityControlId` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ComplianceStatus` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `Confidence` ([`[]NumberFilter`](#NumberFilter_1)) (nullable)
* `CreatedAt` ([`[]DateFilter`](#DateFilter_1)) (nullable)
* `Criticality` ([`[]NumberFilter`](#NumberFilter_1)) (nullable)
* `Description` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `FindingProviderFieldsConfidence` ([`[]NumberFilter`](#NumberFilter_1)) (nullable)
* `FindingProviderFieldsCriticality` ([`[]NumberFilter`](#NumberFilter_1)) (nullable)
* `FindingProviderFieldsRelatedFindingsId` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `FindingProviderFieldsRelatedFindingsProductArn` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `FindingProviderFieldsSeverityLabel` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `FindingProviderFieldsSeverityOriginal` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `FindingProviderFieldsTypes` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `FirstObservedAt` ([`[]DateFilter`](#DateFilter_1)) (nullable)
* `GeneratorId` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `Id` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `Keyword` ([`[]KeywordFilter`](#KeywordFilter)) (nullable)
* `LastObservedAt` ([`[]DateFilter`](#DateFilter_1)) (nullable)
* `MalwareName` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `MalwarePath` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `MalwareState` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `MalwareType` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `NetworkDestinationDomain` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `NetworkDestinationIpV4` ([`[]IpFilter`](#IpFilter)) (nullable)
* `NetworkDestinationIpV6` ([`[]IpFilter`](#IpFilter)) (nullable)
* `NetworkDestinationPort` ([`[]NumberFilter`](#NumberFilter_1)) (nullable)
* `NetworkDirection` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `NetworkProtocol` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `NetworkSourceDomain` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `NetworkSourceIpV4` ([`[]IpFilter`](#IpFilter)) (nullable)
* `NetworkSourceIpV6` ([`[]IpFilter`](#IpFilter)) (nullable)
* `NetworkSourceMac` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `NetworkSourcePort` ([`[]NumberFilter`](#NumberFilter_1)) (nullable)
* `NoteText` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `NoteUpdatedAt` ([`[]DateFilter`](#DateFilter_1)) (nullable)
* `NoteUpdatedBy` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ProcessLaunchedAt` ([`[]DateFilter`](#DateFilter_1)) (nullable)
* `ProcessName` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ProcessParentPid` ([`[]NumberFilter`](#NumberFilter_1)) (nullable)
* `ProcessPath` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ProcessPid` ([`[]NumberFilter`](#NumberFilter_1)) (nullable)
* `ProcessTerminatedAt` ([`[]DateFilter`](#DateFilter_1)) (nullable)
* `ProductArn` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ProductFields` ([`[]MapFilter`](#MapFilter_1)) (nullable)
* `ProductName` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `RecommendationText` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `RecordState` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `Region` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `RelatedFindingsId` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `RelatedFindingsProductArn` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceAwsEc2InstanceIamInstanceProfileArn` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceAwsEc2InstanceImageId` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceAwsEc2InstanceIpV4Addresses` ([`[]IpFilter`](#IpFilter)) (nullable)
* `ResourceAwsEc2InstanceIpV6Addresses` ([`[]IpFilter`](#IpFilter)) (nullable)
* `ResourceAwsEc2InstanceKeyName` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceAwsEc2InstanceLaunchedAt` ([`[]DateFilter`](#DateFilter_1)) (nullable)
* `ResourceAwsEc2InstanceSubnetId` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceAwsEc2InstanceType` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceAwsEc2InstanceVpcId` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceAwsIamAccessKeyCreatedAt` ([`[]DateFilter`](#DateFilter_1)) (nullable)
* `ResourceAwsIamAccessKeyPrincipalName` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceAwsIamAccessKeyStatus` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceAwsIamAccessKeyUserName` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceAwsIamUserUserName` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceAwsS3BucketOwnerId` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceAwsS3BucketOwnerName` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceContainerImageId` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceContainerImageName` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceContainerLaunchedAt` ([`[]DateFilter`](#DateFilter_1)) (nullable)
* `ResourceContainerName` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceDetailsOther` ([`[]MapFilter`](#MapFilter_1)) (nullable)
* `ResourceId` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourcePartition` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceRegion` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ResourceTags` ([`[]MapFilter`](#MapFilter_1)) (nullable)
* `ResourceType` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `Sample` ([`[]BooleanFilter`](#BooleanFilter)) (nullable)
* `SeverityLabel` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `SeverityNormalized` ([`[]NumberFilter`](#NumberFilter_1)) (nullable)
* `SeverityProduct` ([`[]NumberFilter`](#NumberFilter_1)) (nullable)
* `SourceUrl` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ThreatIntelIndicatorCategory` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ThreatIntelIndicatorLastObservedAt` ([`[]DateFilter`](#DateFilter_1)) (nullable)
* `ThreatIntelIndicatorSource` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ThreatIntelIndicatorSourceUrl` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ThreatIntelIndicatorType` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `ThreatIntelIndicatorValue` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `Title` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `Type` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `UpdatedAt` ([`[]DateFilter`](#DateFilter_1)) (nullable)
* `UserDefinedFields` ([`[]MapFilter`](#MapFilter_1)) (nullable)
* `VerificationState` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `WorkflowState` ([`[]StringFilter`](#StringFilter_1)) (nullable)
* `WorkflowStatus` ([`[]StringFilter`](#StringFilter_1)) (nullable)

######StringFilter

<a name="StringFilter_1"></a>

* `Comparison` (`string`)
* `Value` (`string`) (nullable)

######NumberFilter

<a name="NumberFilter_1"></a>

* `Eq` (`number`)
* `Gte` (`number`)
* `Lte` (`number`)

######DateFilter

<a name="DateFilter_1"></a>

* `DateRange` ([`DateRange`](#DateRange)) (nullable)
* `End` (`string`) (nullable)
* `Start` (`string`) (nullable)

######DateRange

<a name="DateRange"></a>

* `Unit` (`string`)
* `Value` (`integer`)

######KeywordFilter

<a name="KeywordFilter"></a>

* `Value` (`string`) (nullable)

######IpFilter

<a name="IpFilter"></a>

* `Cidr` (`string`) (nullable)

######MapFilter

<a name="MapFilter_1"></a>

* `Comparison` (`string`)
* `Key` (`string`) (nullable)
* `Value` (`string`) (nullable)

######BooleanFilter

<a name="BooleanFilter"></a>

* `Value` (`boolean`)

######SortCriterion

<a name="SortCriterion"></a>

* `Field` (`string`) (nullable)
* `SortOrder` (`string`)

###EventBasedSync

<a name="EventBasedSync"></a>

* `full_sync` (`boolean`) (nullable)
* `account` ([`Account`](#Account))
* `kinesis_stream_arn` (`string`) (required)
* `start_time` (`string`) (nullable)

###Strategy

<a name="Strategy"></a>

