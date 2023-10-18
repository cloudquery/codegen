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

  Regions to use.

* `accounts` ([`[]Account`](#Account)) (nullable)

  List of all accounts to fetch information from.

* `org` ([`Org`](#Org)) (nullable)

  In AWS organization mode, CloudQuery will source all accounts underneath automatically.

* `aws_debug` (`boolean`)

  If `true`, will log AWS debug logs, including retries and other request/response metadata.

* `max_retries` (`integer`) (nullable) (default: `10`)

  Defines the maximum number of times an API request will be retried.

* `max_backoff` (`integer`) (nullable) (default: `30`)

  Defines the duration between retry attempts.

* `custom_endpoint_url` (`string`)

  The base URL endpoint the SDK API clients will use to make API calls to.
  The SDK will suffix URI path and query elements to this endpoint.

* `custom_endpoint_hostname_immutable` (`boolean`) (nullable)

  Specifies if the endpoint's hostname can be modified by the SDK's API client.
  When using something like LocalStack make sure to set it equal to `true`.

* `custom_endpoint_partition_id` (`string`)

  The AWS partition the endpoint belongs to.

* `custom_endpoint_signing_region` (`string`)

  The region that should be used for signing the request to the endpoint.

* `initialization_concurrency` (`integer`) (range: `[1,+∞)`) (default: `4`)

  During initialization the AWS source plugin fetches information about each account and region.
  This setting controls how many accounts can be initialized concurrently.
  Only configurations with many accounts (either hardcoded or discovered via Organizations)
  should require modifying this setting, to either lower it to avoid rate limit errors, or to increase it to speed up the initialization process.

* `concurrency` (`integer`) (range: `[1,+∞)`) (default: `50000`)

  The best effort maximum number of Go routines to use. Lower this number to reduce memory usage.

* `use_paid_apis` (`boolean`) (default: `false`)

  When set to `true` plugin will sync data from APIs that incur a fee.
  Currently only `aws_costexplorer*` and `aws_alpha_cloudwatch_metric*` tables require this flag to be set to `true`.

* `table_options` ([`TableOptions`](#TableOptions)) (nullable)

  This is a preview feature (for more information about `preview` features look at [plugin versioning](/docs/plugins/sources/aws/versioning))
  that enables users to override the default options for specific tables.
  The root of the object takes a table name, and the next level takes an API method name.
  The final level is the actual input object as defined by the API.

* `event_based_sync` ([`EventBasedSync`](#EventBasedSync)) (nullable)

  This feature is available only in premium version of the plugin.

* `scheduler` ([`Strategy`](#Strategy))

  The scheduler to use when determining the priority of resources to sync.
  
  For more information about this, see [performance tuning](/docs/advanced-topics/performance-tuning).

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

  A filter to match for the findings to return.

* `MaxResults` (`integer`) (nullable)

  The maximum number of results to return in the response.

* `Sort` ([`SortCriteria`](#SortCriteria)) (nullable)

  The sort order for the findings returned.

###### <a name="Criterion"></a>Criterion

  The criteria to use in the filter that defines the archive rule.

* `Contains` (`[]string`) (nullable)

  A "contains" operator to match for the filter used to create the rule.

* `Eq` (`[]string`) (nullable)

  An "equals" operator to match for the filter used to create the rule.

* `Exists` (`boolean`) (nullable)

  An "exists" operator to match for the filter used to create the rule.

* `Neq` (`[]string`) (nullable)

  A "not equals" operator to match for the filter used to create the rule.

###### <a name="SortCriteria"></a>SortCriteria

  The criteria used to sort.

* `AttributeName` (`string`) (nullable)

  The name of the attribute to sort on.

* `OrderBy` (`string`)

  The sort order, ascending or descending.

#### <a name="CloudtrailEvents"></a>CloudtrailEvents

* `lookup_events` ([`[]CustomCloudtrailLookupEventsInput`](#CustomCloudtrailLookupEventsInput)) (nullable)

##### <a name="CustomCloudtrailLookupEventsInput"></a>CustomCloudtrailLookupEventsInput

* `EndTime` (`string`) (nullable) ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `date-time`)

  Specifies that only events that occur before or at the specified time are
  returned. If the specified end time is before the specified start time, an error
  is returned.

* `EventCategory` (`string`)

  Specifies the event category. If you do not specify an event category, events
  of the category are not returned in the response. For example, if you do not
  specify insight as the value of EventCategory , no Insights events are returned.

* `LookupAttributes` ([`[]LookupAttribute`](#LookupAttribute)) (nullable)

  Contains a list of lookup attributes. Currently the list can contain only one
  item.

* `MaxResults` (`integer`) (nullable)

  The number of events to return. Possible values are 1 through 50. The default
  is 50.

* `StartTime` (`string`) (nullable) ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `date-time`)

  Specifies that only events that occur after or at the specified time are
  returned. If the specified start time is after the specified end time, an error
  is returned.

###### <a name="LookupAttribute"></a>LookupAttribute

  Specifies an attribute and value that filter the events returned.

* `AttributeKey` (`string`)

  Specifies an attribute on which to filter the events returned.
  
  This member is required.

* `AttributeValue` (`string`) (nullable)

  Specifies a value for the specified AttributeKey.
  
  This member is required.

#### <a name="CloudwatchMetrics"></a>CloudwatchMetrics

([`[]CloudwatchMetric`](#CloudwatchMetric))

##### <a name="CloudwatchMetric"></a>CloudwatchMetric

* `list_metrics` ([`CloudwatchListMetricsInput`](#CloudwatchListMetricsInput))

* `get_metric_statistics` ([`[]CloudwatchGetMetricStatisticsInput`](#CloudwatchGetMetricStatisticsInput)) (nullable)

###### <a name="CloudwatchListMetricsInput"></a>CloudwatchListMetricsInput

* `Dimensions` ([`[]DimensionFilter`](#DimensionFilter)) (nullable)

  The dimensions to filter against. Only the dimensions that match exactly will
  be returned.

* `IncludeLinkedAccounts` (`boolean`)

  If you are using this operation in a monitoring account, specify true to
  include metrics from source accounts in the returned data. The default is false .

* `MetricName` (`string`) (nullable)

  The name of the metric to filter against. Only the metrics with names that
  match exactly will be returned.

* `Namespace` (`string`) (nullable)

  The metric namespace to filter against. Only the namespace that matches exactly
  will be returned.

* `OwningAccount` (`string`) (nullable)

  When you use this operation in a monitoring account, use this field to return
  metrics only from one source account. To do so, specify that source account ID
  in this field, and also specify true for IncludeLinkedAccounts .

* `RecentlyActive` (`string`)

  To filter the results to show only metrics that have had data points published
  in the past three hours, specify this parameter with a value of PT3H . This is
  the only valid value for this parameter. The results that are returned are an
  approximation of the value you specify. There is a low probability that the
  returned results include metrics with last published data as much as 40 minutes
  more than the specified time interval.

###### <a name="DimensionFilter"></a>DimensionFilter

  Represents filters for a dimension.

* `Name` (`string`) (nullable)

  The dimension name to be matched.
  
  This member is required.

* `Value` (`string`) (nullable)

  The value of the dimension to be matched.

###### <a name="CloudwatchGetMetricStatisticsInput"></a>CloudwatchGetMetricStatisticsInput

* `EndTime` (`string`) (nullable) ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `date-time`)

  The time stamp that determines the last data point to return. The value
  specified is exclusive; results include data points up to the specified time
  stamp. In a raw HTTP query, the time stamp must be in ISO 8601 UTC format (for
  example, 2016-10-10T23:00:00Z).
  
  This member is required.

* `Period` (`integer`) (nullable)

  The granularity, in seconds, of the returned data points. For metrics with
  regular resolution, a period can be as short as one minute (60 seconds) and must
  be a multiple of 60. For high-resolution metrics that are collected at intervals
  of less than one minute, the period can be 1, 5, 10, 30, 60, or any multiple of
  60. High-resolution metrics are those metrics stored by a PutMetricData call
  that includes a StorageResolution of 1 second. If the StartTime parameter
  specifies a time stamp that is greater than 3 hours ago, you must specify the
  period as follows or no data points in that time range is returned:
    - Start time between 3 hours and 15 days ago - Use a multiple of 60 seconds
    (1 minute).
    - Start time between 15 and 63 days ago - Use a multiple of 300 seconds (5
    minutes).
    - Start time greater than 63 days ago - Use a multiple of 3600 seconds (1
    hour).
  
  This member is required.

* `StartTime` (`string`) (nullable) ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `date-time`)

  The time stamp that determines the first data point to return. Start times are
  evaluated relative to the time that CloudWatch receives the request. The value
  specified is inclusive; results include data points with the specified time
  stamp. In a raw HTTP query, the time stamp must be in ISO 8601 UTC format (for
  example, 2016-10-03T23:00:00Z). CloudWatch rounds the specified time stamp as
  follows:
    - Start time less than 15 days ago - Round down to the nearest whole minute.
    For example, 12:32:34 is rounded down to 12:32:00.
    - Start time between 15 and 63 days ago - Round down to the nearest 5-minute
    clock interval. For example, 12:32:34 is rounded down to 12:30:00.
    - Start time greater than 63 days ago - Round down to the nearest 1-hour
    clock interval. For example, 12:32:34 is rounded down to 12:00:00.
  If you set Period to 5, 10, or 30, the start time of your request is rounded
  down to the nearest time that corresponds to even 5-, 10-, or 30-second
  divisions of a minute. For example, if you make a query at (HH:mm:ss) 01:05:23
  for the previous 10-second period, the start time of your request is rounded
  down and you receive data from 01:05:10 to 01:05:20. If you make a query at
  15:07:17 for the previous 5 minutes of data, using a period of 5 seconds, you
  receive data timestamped between 15:02:15 and 15:07:15.
  
  This member is required.

* `ExtendedStatistics` (`[]string`) (nullable)

  The percentile statistics. Specify values between p0.0 and p100. When calling
  GetMetricStatistics , you must specify either Statistics or ExtendedStatistics ,
  but not both. Percentile statistics are not available for metrics when any of
  the metric values are negative numbers.

* `Statistics` (`[]string`) (nullable)

  The metric statistics, other than percentile. For percentile statistics, use
  ExtendedStatistics . When calling GetMetricStatistics , you must specify either
  Statistics or ExtendedStatistics , but not both.

* `Unit` (`string`)

  The unit for a given metric. If you omit Unit , all data that was collected with
  any unit is returned, along with the corresponding units that were specified
  when the data was reported to CloudWatch. If you specify a unit, the operation
  returns only data that was collected with that unit specified. If you specify a
  unit that does not match the data collected, the results of the operation are
  null. CloudWatch does not perform unit conversions.

#### <a name="CostExplorerAPIs"></a>CostExplorerAPIs

* `get_cost_and_usage` ([`[]CustomGetCostAndUsageInput`](#CustomGetCostAndUsageInput)) (nullable)

##### <a name="CustomGetCostAndUsageInput"></a>CustomGetCostAndUsageInput

* `Granularity` (`string`)

  Sets the Amazon Web Services cost granularity to MONTHLY or DAILY , or HOURLY .
  If Granularity isn't set, the response object doesn't include the Granularity ,
  either MONTHLY or DAILY , or HOURLY .
  
  This member is required.

* `Metrics` (`[]string`) (nullable)

  Which metrics are returned in the query. For more information about blended and
  unblended rates, see Why does the "blended" annotation appear on some line
  items in my bill? (http://aws.amazon.com/premiumsupport/knowledge-center/blended-rates-intro/)
  . Valid values are AmortizedCost , BlendedCost , NetAmortizedCost ,
  NetUnblendedCost , NormalizedUsageAmount , UnblendedCost , and UsageQuantity .
  If you return the UsageQuantity metric, the service aggregates all usage
  numbers without taking into account the units. For example, if you aggregate
  usageQuantity across all of Amazon EC2, the results aren't meaningful because
  Amazon EC2 compute hours and data transfer are measured in different units (for
  example, hours and GB). To get more meaningful UsageQuantity metrics, filter by
  UsageType or UsageTypeGroups . Metrics is required for GetCostAndUsage requests.
  
  This member is required.

* `TimePeriod` ([`DateInterval`](#DateInterval)) (nullable)

  Sets the start date and end date for retrieving Amazon Web Services costs. The
  start date is inclusive, but the end date is exclusive. For example, if start
  is 2017-01-01 and end is 2017-05-01 , then the cost and usage data is retrieved
  from 2017-01-01 up to and including 2017-04-30 but not including 2017-05-01 .
  
  This member is required.

* `Filter` ([`Expression`](#Expression)) (nullable)

  Filters Amazon Web Services costs by different dimensions. For example, you can
  specify SERVICE and LINKED_ACCOUNT and get the costs that are associated with
  that account's usage of that service. You can nest Expression objects to define
  any combination of dimension filters. For more information, see Expression (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
  . Valid values for MatchOptions for Dimensions are EQUALS and CASE_SENSITIVE .
  Valid values for MatchOptions for CostCategories and Tags are EQUALS , ABSENT ,
  and CASE_SENSITIVE . Default values are EQUALS and CASE_SENSITIVE .

* `GroupBy` ([`[]GroupDefinition`](#GroupDefinition)) (nullable)

  You can group Amazon Web Services costs using up to two different groups,
  either dimensions, tag keys, cost categories, or any two group by types. Valid
  values for the DIMENSION type are AZ , INSTANCE_TYPE , LEGAL_ENTITY_NAME ,
  INVOICING_ENTITY , LINKED_ACCOUNT , OPERATION , PLATFORM , PURCHASE_TYPE ,
  SERVICE , TENANCY , RECORD_TYPE , and USAGE_TYPE . When you group by the TAG
  type and include a valid tag key, you get all tag values, including empty
  strings.

###### <a name="DateInterval"></a>DateInterval

  The time period of the request.

* `End` (`string`) (nullable)

  The end of the time period. The end date is exclusive. For example, if end is
  2017-05-01 , Amazon Web Services retrieves cost and usage data from the start
  date up to, but not including, 2017-05-01 .
  
  This member is required.

* `Start` (`string`) (nullable)

  The beginning of the time period. The start date is inclusive. For example, if
  start is 2017-01-01 , Amazon Web Services retrieves cost and usage data starting
  at 2017-01-01 up to the end date. The start date must be equal to or no later
  than the current date to avoid a validation error.
  
  This member is required.

###### <a name="Expression"></a>Expression

  Use Expression to filter in various Cost Explorer APIs.

* `And` ([`[]Expression`](#Expression)) (nullable)

  Return results that match both Dimension objects.

* `CostCategories` ([`CostCategoryValues`](#CostCategoryValues)) (nullable)

  The filter that's based on CostCategory values.

* `Dimensions` ([`DimensionValues`](#DimensionValues)) (nullable)

  The specific Dimension to use for Expression .

* `Not` ([`Expression`](#Expression)) (nullable)

  Return results that don't match a Dimension object.

* `Or` ([`[]Expression`](#Expression)) (nullable)

  Return results that match either Dimension object.

* `Tags` ([`TagValues`](#TagValues)) (nullable)

  The specific Tag to use for Expression .

###### <a name="CostCategoryValues"></a>CostCategoryValues

  The Cost Categories values used for filtering the costs.

* `Key` (`string`) (nullable)

  The unique name of the Cost Category.

* `MatchOptions` (`[]string`) (nullable)

  The match options that you can use to filter your results. MatchOptions is only
  applicable for actions related to cost category. The default values for
  MatchOptions is EQUALS and CASE_SENSITIVE .

* `Values` (`[]string`) (nullable)

  The specific value of the Cost Category.

###### <a name="DimensionValues"></a>DimensionValues

  The metadata that you can use to filter and group your results.

* `Key` (`string`)

  The names of the metadata types that you can use to filter and group your
  results. For example, AZ returns a list of Availability Zones. Not all
  dimensions are supported in each API. Refer to the documentation for each
  specific API to see what is supported. LINK_ACCOUNT_NAME and SERVICE_CODE can
  only be used in CostCategoryRule (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_CostCategoryRule.html)
  . ANOMALY_TOTAL_IMPACT_ABSOLUTE and ANOMALY_TOTAL_IMPACT_PERCENTAGE can only be
  used in AnomalySubscriptions (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_AnomalySubscription.html)
  .

* `MatchOptions` (`[]string`) (nullable)

  The match options that you can use to filter your results. MatchOptions is only
  applicable for actions related to Cost Category and Anomaly Subscriptions. Refer
  to the documentation for each specific API to see what is supported. The default
  values for MatchOptions are EQUALS and CASE_SENSITIVE .

* `Values` (`[]string`) (nullable)

  The metadata values that you can use to filter and group your results. You can
  use GetDimensionValues to find specific values.

###### <a name="TagValues"></a>TagValues

  The values that are available for a tag.

* `Key` (`string`) (nullable)

  The key for the tag.

* `MatchOptions` (`[]string`) (nullable)

  The match options that you can use to filter your results. MatchOptions is only
  applicable for actions related to Cost Category. The default values for
  MatchOptions are EQUALS and CASE_SENSITIVE .

* `Values` (`[]string`) (nullable)

  The specific value of the tag.

###### <a name="GroupDefinition"></a>GroupDefinition

  Represents a group when you specify a group by criteria or in the response to a query with a specific grouping.

* `Key` (`string`) (nullable)

  The string that represents a key for a specified group.

* `Type` (`string`)

  The string that represents the type of group.

#### <a name="ECSTasks"></a>ECSTasks

* `list_tasks` ([`[]CustomECSListTasksInput`](#CustomECSListTasksInput)) (nullable)

##### <a name="CustomECSListTasksInput"></a>CustomECSListTasksInput

* `ContainerInstance` (`string`) (nullable)

  The container instance ID or full ARN of the container instance to use when
  filtering the ListTasks results. Specifying a containerInstance limits the
  results to tasks that belong to that container instance.

* `DesiredStatus` (`string`)

  The task desired status to use when filtering the ListTasks results. Specifying
  a desiredStatus of STOPPED limits the results to tasks that Amazon ECS has set
  the desired status to STOPPED . This can be useful for debugging tasks that
  aren't starting properly or have died or finished. The default status filter is
  RUNNING , which shows tasks that Amazon ECS has set the desired status to
  RUNNING . Although you can filter results based on a desired status of PENDING ,
  this doesn't return any results. Amazon ECS never sets the desired status of a
  task to that value (only a task's lastStatus may have a value of PENDING ).

* `Family` (`string`) (nullable)

  The name of the task definition family to use when filtering the ListTasks
  results. Specifying a family limits the results to tasks that belong to that
  family.

* `LaunchType` (`string`)

  The launch type to use when filtering the ListTasks results.

* `MaxResults` (`integer`) (nullable) (range: `[1,100]`) (default: `100`)

  The maximum number of task results that ListTasks returned in paginated output.
  When this parameter is used, ListTasks only returns maxResults results in a
  single page along with a nextToken response element. The remaining results of
  the initial request can be seen by sending another ListTasks request with the
  returned nextToken value. This value can be between 1 and 100. If this
  parameter isn't used, then ListTasks returns up to 100 results and a nextToken
  value if applicable.

* `ServiceName` (`string`) (nullable)

  The name of the service to use when filtering the ListTasks results. Specifying
  a serviceName limits the results to tasks that belong to that service.

* `StartedBy` (`string`) (nullable)

  The startedBy value to filter the task results with. Specifying a startedBy
  value limits the results to tasks that were started with that value. When you
  specify startedBy as the filter, it must be the only filter that you use.

#### <a name="Inspector2Findings"></a>Inspector2Findings

* `list_findings` ([`[]CustomInspector2ListFindingsInput`](#CustomInspector2ListFindingsInput)) (nullable)

##### <a name="CustomInspector2ListFindingsInput"></a>CustomInspector2ListFindingsInput

* `FilterCriteria` ([`FilterCriteria`](#FilterCriteria)) (nullable)

  Details on the filters to apply to your finding results.

* `MaxResults` (`integer`) (nullable)

  The maximum number of results to return in the response.

* `SortCriteria` ([`SortCriteria`](#SortCriteria-1)) (nullable)

  Details on the sort criteria to apply to your finding results.

###### <a name="FilterCriteria"></a>FilterCriteria

  Details on the criteria used to define the filter.

* `AwsAccountId` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details of the Amazon Web Services account IDs used to filter findings.

* `CodeVulnerabilityDetectorName` ([`[]StringFilter`](#StringFilter)) (nullable)

  The name of the detector used to identify a code vulnerability in a Lambda
  function used to filter findings.

* `CodeVulnerabilityDetectorTags` ([`[]StringFilter`](#StringFilter)) (nullable)

  The detector type tag associated with the vulnerability used to filter
  findings. Detector tags group related vulnerabilities by common themes or
  tactics. For a list of available tags by programming language, see Java tags (https://docs.aws.amazon.com/codeguru/detector-library/java/tags/)
  , or Python tags (https://docs.aws.amazon.com/codeguru/detector-library/python/tags/)
  .

* `CodeVulnerabilityFilePath` ([`[]StringFilter`](#StringFilter)) (nullable)

  The file path to the file in a Lambda function that contains a code
  vulnerability used to filter findings.

* `ComponentId` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details of the component IDs used to filter findings.

* `ComponentType` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details of the component types used to filter findings.

* `Ec2InstanceImageId` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details of the Amazon EC2 instance image IDs used to filter findings.

* `Ec2InstanceSubnetId` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details of the Amazon EC2 instance subnet IDs used to filter findings.

* `Ec2InstanceVpcId` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details of the Amazon EC2 instance VPC IDs used to filter findings.

* `EcrImageArchitecture` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details of the Amazon ECR image architecture types used to filter findings.

* `EcrImageHash` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details of the Amazon ECR image hashes used to filter findings.

* `EcrImagePushedAt` ([`[]DateFilter`](#DateFilter)) (nullable)

  Details on the Amazon ECR image push date and time used to filter findings.

* `EcrImageRegistry` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details on the Amazon ECR registry used to filter findings.

* `EcrImageRepositoryName` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details on the name of the Amazon ECR repository used to filter findings.

* `EcrImageTags` ([`[]StringFilter`](#StringFilter)) (nullable)

  The tags attached to the Amazon ECR container image.

* `EpssScore` ([`[]NumberFilter`](#NumberFilter)) (nullable)

  The EPSS score used to filter findings.

* `ExploitAvailable` ([`[]StringFilter`](#StringFilter)) (nullable)

  Filters the list of AWS Lambda findings by the availability of exploits.

* `FindingArn` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details on the finding ARNs used to filter findings.

* `FindingStatus` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details on the finding status types used to filter findings.

* `FindingType` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details on the finding types used to filter findings.

* `FirstObservedAt` ([`[]DateFilter`](#DateFilter)) (nullable)

  Details on the date and time a finding was first seen used to filter findings.

* `FixAvailable` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details on whether a fix is available through a version update. This value can
  be YES , NO , or PARTIAL . A PARTIAL fix means that some, but not all, of the
  packages identified in the finding have fixes available through updated
  versions.

* `InspectorScore` ([`[]NumberFilter`](#NumberFilter)) (nullable)

  The Amazon Inspector score to filter on.

* `LambdaFunctionExecutionRoleArn` ([`[]StringFilter`](#StringFilter)) (nullable)

  Filters the list of AWS Lambda functions by execution role.

* `LambdaFunctionLastModifiedAt` ([`[]DateFilter`](#DateFilter)) (nullable)

  Filters the list of AWS Lambda functions by the date and time that a user last
  updated the configuration, in ISO 8601 format (https://www.iso.org/iso-8601-date-and-time-format.html)

* `LambdaFunctionLayers` ([`[]StringFilter`](#StringFilter)) (nullable)

  Filters the list of AWS Lambda functions by the function's  layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
  . A Lambda function can have up to five layers.

* `LambdaFunctionName` ([`[]StringFilter`](#StringFilter)) (nullable)

  Filters the list of AWS Lambda functions by the name of the function.

* `LambdaFunctionRuntime` ([`[]StringFilter`](#StringFilter)) (nullable)

  Filters the list of AWS Lambda functions by the runtime environment for the
  Lambda function.

* `LastObservedAt` ([`[]DateFilter`](#DateFilter)) (nullable)

  Details on the date and time a finding was last seen used to filter findings.

* `NetworkProtocol` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details on network protocol used to filter findings.

* `PortRange` ([`[]PortRangeFilter`](#PortRangeFilter)) (nullable)

  Details on the port ranges used to filter findings.

* `RelatedVulnerabilities` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details on the related vulnerabilities used to filter findings.

* `ResourceId` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details on the resource IDs used to filter findings.

* `ResourceTags` ([`[]MapFilter`](#MapFilter)) (nullable)

  Details on the resource tags used to filter findings.

* `ResourceType` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details on the resource types used to filter findings.

* `Severity` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details on the severity used to filter findings.

* `Title` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details on the finding title used to filter findings.

* `UpdatedAt` ([`[]DateFilter`](#DateFilter)) (nullable)

  Details on the date and time a finding was last updated at used to filter
  findings.

* `VendorSeverity` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details on the vendor severity used to filter findings.

* `VulnerabilityId` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details on the vulnerability ID used to filter findings.

* `VulnerabilitySource` ([`[]StringFilter`](#StringFilter)) (nullable)

  Details on the vulnerability type used to filter findings.

* `VulnerablePackages` ([`[]PackageFilter`](#PackageFilter)) (nullable)

  Details on the vulnerable packages used to filter findings.

###### <a name="StringFilter"></a>StringFilter

  An object that describes the details of a string filter.

* `Comparison` (`string`)

  The operator to use when comparing values in the filter.
  
  This member is required.

* `Value` (`string`) (nullable)

  The value to filter on.
  
  This member is required.

###### <a name="DateFilter"></a>DateFilter

  Contains details on the time range used to filter findings.

* `EndInclusive` (`string`) (nullable) ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `date-time`)

  A timestamp representing the end of the time period filtered on.

* `StartInclusive` (`string`) (nullable) ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `date-time`)

  A timestamp representing the start of the time period filtered on.

###### <a name="NumberFilter"></a>NumberFilter

  An object that describes the details of a number filter.

* `LowerInclusive` (`number`) (nullable)

  The lowest number to be included in the filter.

* `UpperInclusive` (`number`) (nullable)

  The highest number to be included in the filter.

###### <a name="PortRangeFilter"></a>PortRangeFilter

  An object that describes the details of a port range filter.

* `BeginInclusive` (`integer`) (nullable)

  The port number the port range begins at.

* `EndInclusive` (`integer`) (nullable)

  The port number the port range ends at.

###### <a name="MapFilter"></a>MapFilter

  An object that describes details of a map filter.

* `Comparison` (`string`)

  The operator to use when comparing values in the filter.
  
  This member is required.

* `Key` (`string`) (nullable)

  The tag key used in the filter.
  
  This member is required.

* `Value` (`string`) (nullable)

  The tag value used in the filter.

###### <a name="PackageFilter"></a>PackageFilter

  Contains information on the details of a package filter.

* `Architecture` ([`StringFilter`](#StringFilter)) (nullable)

  An object that contains details on the package architecture type to filter on.

* `Epoch` ([`NumberFilter`](#NumberFilter)) (nullable)

  An object that contains details on the package epoch to filter on.

* `Name` ([`StringFilter`](#StringFilter)) (nullable)

  An object that contains details on the name of the package to filter on.

* `Release` ([`StringFilter`](#StringFilter)) (nullable)

  An object that contains details on the package release to filter on.

* `SourceLambdaLayerArn` ([`StringFilter`](#StringFilter)) (nullable)

  An object that describes the details of a string filter.

* `SourceLayerHash` ([`StringFilter`](#StringFilter)) (nullable)

  An object that contains details on the source layer hash to filter on.

* `Version` ([`StringFilter`](#StringFilter)) (nullable)

  The package version to filter on.

###### <a name="SortCriteria-1"></a>SortCriteria

  Details about the criteria used to sort finding results.

* `Field` (`string`)

  The finding detail field by which results are sorted.
  
  This member is required.

* `SortOrder` (`string`)

  The order by which findings are sorted.
  
  This member is required.

#### <a name="SecurityHubFindings"></a>SecurityHubFindings

* `get_findings` ([`[]CustomSecurityHubGetFindingsInput`](#CustomSecurityHubGetFindingsInput)) (nullable)

##### <a name="CustomSecurityHubGetFindingsInput"></a>CustomSecurityHubGetFindingsInput

* `Filters` ([`AwsSecurityFindingFilters`](#AwsSecurityFindingFilters)) (nullable)

  The finding attributes used to define a condition to filter the returned
  findings. You can filter by up to 10 finding attributes. For each attribute, you
  can provide up to 20 filter values. Note that in the available filter fields,
  WorkflowState is deprecated. To search for a finding based on its workflow
  status, use WorkflowStatus .

* `MaxResults` (`integer`) (range: `[1,100]`) (default: `100`)

  The maximum number of findings to return.

* `SortCriteria` ([`[]SortCriterion`](#SortCriterion)) (nullable)

  The finding attributes used to sort the list of returned findings.

###### <a name="AwsSecurityFindingFilters"></a>AwsSecurityFindingFilters

  A collection of attributes that are applied to all active Security Hub-aggregated findings and that result in a subset of findings that are included in this insight.

* `AwsAccountId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The Amazon Web Services account ID that a finding is generated in.

* `CompanyName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The name of the findings provider (company) that owns the solution (product)
  that generates findings.

* `ComplianceAssociatedStandardsId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The unique identifier of a standard in which a control is enabled. This field
  consists of the resource portion of the Amazon Resource Name (ARN) returned for
  a standard in the DescribeStandards (https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeStandards.html)
  API response.

* `ComplianceSecurityControlId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The unique identifier of a control across standards. Values for this field
  typically consist of an Amazon Web Service and a number, such as APIGateway.5.

* `ComplianceStatus` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  Exclusive to findings that are generated as the result of a check run against a
  specific rule in a supported standard, such as CIS Amazon Web Services
  Foundations. Contains security standard-related finding details.

* `Confidence` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

  A finding's confidence. Confidence is defined as the likelihood that a finding
  accurately identifies the behavior or issue that it was intended to identify.
  Confidence is scored on a 0-100 basis using a ratio scale, where 0 means zero
  percent confidence and 100 means 100 percent confidence.

* `CreatedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

  An ISO8601-formatted timestamp that indicates when the security findings
  provider captured the potential security issue that a finding captured. A
  correctly formatted example is 2020-05-21T20:16:34.724Z . The value cannot
  contain spaces, and date and time should be separated by T . For more
  information, see RFC 3339 section 5.6, Internet Date/Time Format (https://www.rfc-editor.org/rfc/rfc3339#section-5.6)
  .

* `Criticality` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

  The level of importance assigned to the resources associated with the finding.
  A score of 0 means that the underlying resources have no criticality, and a
  score of 100 is reserved for the most critical resources.

* `Description` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  A finding's description.

* `FindingProviderFieldsConfidence` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

  The finding provider value for the finding confidence. Confidence is defined as
  the likelihood that a finding accurately identifies the behavior or issue that
  it was intended to identify. Confidence is scored on a 0-100 basis using a ratio
  scale, where 0 means zero percent confidence and 100 means 100 percent
  confidence.

* `FindingProviderFieldsCriticality` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

  The finding provider value for the level of importance assigned to the
  resources associated with the findings. A score of 0 means that the underlying
  resources have no criticality, and a score of 100 is reserved for the most
  critical resources.

* `FindingProviderFieldsRelatedFindingsId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The finding identifier of a related finding that is identified by the finding
  provider.

* `FindingProviderFieldsRelatedFindingsProductArn` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The ARN of the solution that generated a related finding that is identified by
  the finding provider.

* `FindingProviderFieldsSeverityLabel` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The finding provider value for the severity label.

* `FindingProviderFieldsSeverityOriginal` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The finding provider's original value for the severity.

* `FindingProviderFieldsTypes` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  One or more finding types that the finding provider assigned to the finding.
  Uses the format of namespace/category/classifier that classify a finding. Valid
  namespace values are: Software and Configuration Checks | TTPs | Effects |
  Unusual Behaviors | Sensitive Data Identifications

* `FirstObservedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

  An ISO8601-formatted timestamp that indicates when the security findings
  provider first observed the potential security issue that a finding captured. A
  correctly formatted example is 2020-05-21T20:16:34.724Z . The value cannot
  contain spaces, and date and time should be separated by T . For more
  information, see RFC 3339 section 5.6, Internet Date/Time Format (https://www.rfc-editor.org/rfc/rfc3339#section-5.6)
  .

* `GeneratorId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The identifier for the solution-specific component (a discrete unit of logic)
  that generated a finding. In various security findings providers' solutions,
  this generator can be called a rule, a check, a detector, a plugin, etc.

* `Id` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The security findings provider-specific identifier for a finding.

* `Keyword` ([`[]KeywordFilter`](#KeywordFilter)) (nullable)

  A keyword for a finding.
  
  Deprecated: The Keyword property is deprecated.

* `LastObservedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

  An ISO8601-formatted timestamp that indicates when the security findings
  provider most recently observed the potential security issue that a finding
  captured. A correctly formatted example is 2020-05-21T20:16:34.724Z . The value
  cannot contain spaces, and date and time should be separated by T . For more
  information, see RFC 3339 section 5.6, Internet Date/Time Format (https://www.rfc-editor.org/rfc/rfc3339#section-5.6)
  .

* `MalwareName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The name of the malware that was observed.

* `MalwarePath` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The filesystem path of the malware that was observed.

* `MalwareState` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The state of the malware that was observed.

* `MalwareType` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The type of the malware that was observed.

* `NetworkDestinationDomain` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The destination domain of network-related information about a finding.

* `NetworkDestinationIpV4` ([`[]IpFilter`](#IpFilter)) (nullable)

  The destination IPv4 address of network-related information about a finding.

* `NetworkDestinationIpV6` ([`[]IpFilter`](#IpFilter)) (nullable)

  The destination IPv6 address of network-related information about a finding.

* `NetworkDestinationPort` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

  The destination port of network-related information about a finding.

* `NetworkDirection` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  Indicates the direction of network traffic associated with a finding.

* `NetworkProtocol` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The protocol of network-related information about a finding.

* `NetworkSourceDomain` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The source domain of network-related information about a finding.

* `NetworkSourceIpV4` ([`[]IpFilter`](#IpFilter)) (nullable)

  The source IPv4 address of network-related information about a finding.

* `NetworkSourceIpV6` ([`[]IpFilter`](#IpFilter)) (nullable)

  The source IPv6 address of network-related information about a finding.

* `NetworkSourceMac` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The source media access control (MAC) address of network-related information
  about a finding.

* `NetworkSourcePort` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

  The source port of network-related information about a finding.

* `NoteText` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The text of a note.

* `NoteUpdatedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

  The timestamp of when the note was updated.

* `NoteUpdatedBy` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The principal that created a note.

* `ProcessLaunchedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

  A timestamp that identifies when the process was launched. A correctly
  formatted example is 2020-05-21T20:16:34.724Z . The value cannot contain spaces,
  and date and time should be separated by T . For more information, see RFC 3339
  section 5.6, Internet Date/Time Format (https://www.rfc-editor.org/rfc/rfc3339#section-5.6)
  .

* `ProcessName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The name of the process.

* `ProcessParentPid` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

  The parent process ID. This field accepts positive integers between O and
  2147483647 .

* `ProcessPath` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The path to the process executable.

* `ProcessPid` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

  The process ID.

* `ProcessTerminatedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

  A timestamp that identifies when the process was terminated. A correctly
  formatted example is 2020-05-21T20:16:34.724Z . The value cannot contain spaces,
  and date and time should be separated by T . For more information, see RFC 3339
  section 5.6, Internet Date/Time Format (https://www.rfc-editor.org/rfc/rfc3339#section-5.6)
  .

* `ProductArn` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The ARN generated by Security Hub that uniquely identifies a third-party
  company (security findings provider) after this provider's product (solution
  that generates findings) is registered with Security Hub.

* `ProductFields` ([`[]MapFilter`](#MapFilter-1)) (nullable)

  A data type where security findings providers can include additional
  solution-specific details that aren't part of the defined AwsSecurityFinding
  format.

* `ProductName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The name of the solution (product) that generates findings.

* `RecommendationText` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The recommendation of what to do about the issue described in a finding.

* `RecordState` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The updated record state for the finding.

* `Region` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The Region from which the finding was generated.

* `RelatedFindingsId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The solution-generated identifier for a related finding.

* `RelatedFindingsProductArn` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The ARN of the solution that generated a related finding.

* `ResourceAwsEc2InstanceIamInstanceProfileArn` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The IAM profile ARN of the instance.

* `ResourceAwsEc2InstanceImageId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The Amazon Machine Image (AMI) ID of the instance.

* `ResourceAwsEc2InstanceIpV4Addresses` ([`[]IpFilter`](#IpFilter)) (nullable)

  The IPv4 addresses associated with the instance.

* `ResourceAwsEc2InstanceIpV6Addresses` ([`[]IpFilter`](#IpFilter)) (nullable)

  The IPv6 addresses associated with the instance.

* `ResourceAwsEc2InstanceKeyName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The key name associated with the instance.

* `ResourceAwsEc2InstanceLaunchedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

  The date and time the instance was launched.

* `ResourceAwsEc2InstanceSubnetId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The identifier of the subnet that the instance was launched in.

* `ResourceAwsEc2InstanceType` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The instance type of the instance.

* `ResourceAwsEc2InstanceVpcId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The identifier of the VPC that the instance was launched in.

* `ResourceAwsIamAccessKeyCreatedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

  The creation date/time of the IAM access key related to a finding.

* `ResourceAwsIamAccessKeyPrincipalName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The name of the principal that is associated with an IAM access key.

* `ResourceAwsIamAccessKeyStatus` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The status of the IAM access key related to a finding.

* `ResourceAwsIamAccessKeyUserName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The user associated with the IAM access key related to a finding.
  
  Deprecated: This filter is deprecated. Instead, use
  ResourceAwsIamAccessKeyPrincipalName.

* `ResourceAwsIamUserUserName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The name of an IAM user.

* `ResourceAwsS3BucketOwnerId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The canonical user ID of the owner of the S3 bucket.

* `ResourceAwsS3BucketOwnerName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The display name of the owner of the S3 bucket.

* `ResourceContainerImageId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The identifier of the image related to a finding.

* `ResourceContainerImageName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The name of the image related to a finding.

* `ResourceContainerLaunchedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

  A timestamp that identifies when the container was started. A correctly
  formatted example is 2020-05-21T20:16:34.724Z . The value cannot contain spaces,
  and date and time should be separated by T . For more information, see RFC 3339
  section 5.6, Internet Date/Time Format (https://www.rfc-editor.org/rfc/rfc3339#section-5.6)
  .

* `ResourceContainerName` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The name of the container related to a finding.

* `ResourceDetailsOther` ([`[]MapFilter`](#MapFilter-1)) (nullable)

  The details of a resource that doesn't have a specific subfield for the
  resource type defined.

* `ResourceId` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The canonical identifier for the given resource type.

* `ResourcePartition` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The canonical Amazon Web Services partition name that the Region is assigned to.

* `ResourceRegion` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The canonical Amazon Web Services external Region name where this resource is
  located.

* `ResourceTags` ([`[]MapFilter`](#MapFilter-1)) (nullable)

  A list of Amazon Web Services tags associated with a resource at the time the
  finding was processed.

* `ResourceType` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  Specifies the type of the resource that details are provided for.

* `Sample` ([`[]BooleanFilter`](#BooleanFilter)) (nullable)

  Indicates whether or not sample findings are included in the filter results.

* `SeverityLabel` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The label of a finding's severity.

* `SeverityNormalized` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

  The normalized severity of a finding.
  
  Deprecated: This filter is deprecated. Instead, use SeverityLabel or
  FindingProviderFieldsSeverityLabel.

* `SeverityProduct` ([`[]NumberFilter`](#NumberFilter-1)) (nullable)

  The native severity as defined by the security findings provider's solution
  that generated the finding.
  
  Deprecated: This filter is deprecated. Instead, use
  FindingProviderSeverityOriginal.

* `SourceUrl` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  A URL that links to a page about the current finding in the security findings
  provider's solution.

* `ThreatIntelIndicatorCategory` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The category of a threat intelligence indicator.

* `ThreatIntelIndicatorLastObservedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

  A timestamp that identifies the last observation of a threat intelligence
  indicator.

* `ThreatIntelIndicatorSource` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The source of the threat intelligence.

* `ThreatIntelIndicatorSourceUrl` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The URL for more details from the source of the threat intelligence.

* `ThreatIntelIndicatorType` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The type of a threat intelligence indicator.

* `ThreatIntelIndicatorValue` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The value of a threat intelligence indicator.

* `Title` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  A finding's title.

* `Type` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  A finding type in the format of namespace/category/classifier that classifies a
  finding.

* `UpdatedAt` ([`[]DateFilter`](#DateFilter-1)) (nullable)

  An ISO8601-formatted timestamp that indicates when the security findings
  provider last updated the finding record. A correctly formatted example is
  2020-05-21T20:16:34.724Z . The value cannot contain spaces, and date and time
  should be separated by T . For more information, see RFC 3339 section 5.6,
  Internet Date/Time Format (https://www.rfc-editor.org/rfc/rfc3339#section-5.6) .

* `UserDefinedFields` ([`[]MapFilter`](#MapFilter-1)) (nullable)

  A list of name/value string pairs associated with the finding. These are
  custom, user-defined fields added to a finding.

* `VerificationState` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The veracity of a finding.

* `WorkflowState` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The workflow state of a finding. Note that this field is deprecated. To search
  for a finding based on its workflow status, use WorkflowStatus .

* `WorkflowStatus` ([`[]StringFilter`](#StringFilter-1)) (nullable)

  The status of the investigation into a finding. Allowed values are the
  following.
    - NEW - The initial state of a finding, before it is reviewed. Security Hub
    also resets the workflow status from NOTIFIED or RESOLVED to NEW in the
    following cases:
    - RecordState changes from ARCHIVED to ACTIVE .
    - Compliance.Status changes from PASSED to either WARNING , FAILED , or
    NOT_AVAILABLE .
    - NOTIFIED - Indicates that the resource owner has been notified about the
    security issue. Used when the initial reviewer is not the resource owner, and
    needs intervention from the resource owner. If one of the following occurs, the
    workflow status is changed automatically from NOTIFIED to NEW :
    - RecordState changes from ARCHIVED to ACTIVE .
    - Compliance.Status changes from PASSED to FAILED , WARNING , or NOT_AVAILABLE
    .
    - SUPPRESSED - Indicates that you reviewed the finding and do not believe that
    any action is needed. The workflow status of a SUPPRESSED finding does not
    change if RecordState changes from ARCHIVED to ACTIVE .
    - RESOLVED - The finding was reviewed and remediated and is now considered
    resolved. The finding remains RESOLVED unless one of the following occurs:
    - RecordState changes from ARCHIVED to ACTIVE .
    - Compliance.Status changes from PASSED to FAILED , WARNING , or NOT_AVAILABLE
    . In those cases, the workflow status is automatically reset to NEW . For
    findings from controls, if Compliance.Status is PASSED , then Security Hub
    automatically sets the workflow status to RESOLVED .

###### <a name="StringFilter-1"></a>StringFilter

  A string filter for filtering Security Hub findings.

* `Comparison` (`string`)

  The condition to apply to a string value when filtering Security Hub findings.
  To search for values that have the filter value, use one of the following
  comparison operators:
    - To search for values that include the filter value, use CONTAINS . For
    example, the filter Title CONTAINS CloudFront matches findings that have a
    Title that includes the string CloudFront.
    - To search for values that exactly match the filter value, use EQUALS . For
    example, the filter AwsAccountId EQUALS 123456789012 only matches findings
    that have an account ID of 123456789012 .
    - To search for values that start with the filter value, use PREFIX . For
    example, the filter ResourceRegion PREFIX us matches findings that have a
    ResourceRegion that starts with us . A ResourceRegion that starts with a
    different value, such as af , ap , or ca , doesn't match.
  CONTAINS , EQUALS , and PREFIX filters on the same field are joined by OR . A
  finding matches if it matches any one of those filters. For example, the filters
  Title CONTAINS CloudFront OR Title CONTAINS CloudWatch match a finding that
  includes either CloudFront , CloudWatch , or both strings in the title. To
  search for values that don’t have the filter value, use one of the following
  comparison operators:
    - To search for values that exclude the filter value, use NOT_CONTAINS . For
    example, the filter Title NOT_CONTAINS CloudFront matches findings that have a
    Title that excludes the string CloudFront.
    - To search for values other than the filter value, use NOT_EQUALS . For
    example, the filter AwsAccountId NOT_EQUALS 123456789012 only matches findings
    that have an account ID other than 123456789012 .
    - To search for values that don't start with the filter value, use
    PREFIX_NOT_EQUALS . For example, the filter ResourceRegion PREFIX_NOT_EQUALS
    us matches findings with a ResourceRegion that starts with a value other than
    us .
  NOT_CONTAINS , NOT_EQUALS , and PREFIX_NOT_EQUALS filters on the same field are
  joined by AND . A finding matches only if it matches all of those filters. For
  example, the filters Title NOT_CONTAINS CloudFront AND Title NOT_CONTAINS
  CloudWatch match a finding that excludes both CloudFront and CloudWatch in the
  title. You can’t have both a CONTAINS filter and a NOT_CONTAINS filter on the
  same field. Similarly, you can't provide both an EQUALS filter and a NOT_EQUALS
  or PREFIX_NOT_EQUALS filter on the same field. Combining filters in this way
  returns an error. CONTAINS filters can only be used with other CONTAINS
  filters. NOT_CONTAINS filters can only be used with other NOT_CONTAINS filters.
  You can combine PREFIX filters with NOT_EQUALS or PREFIX_NOT_EQUALS filters for
  the same field. Security Hub first processes the PREFIX filters, and then the
  NOT_EQUALS or PREFIX_NOT_EQUALS filters. For example, for the following
  filters, Security Hub first identifies findings that have resource types that
  start with either AwsIam or AwsEc2 . It then excludes findings that have a
  resource type of AwsIamPolicy and findings that have a resource type of
  AwsEc2NetworkInterface .
    - ResourceType PREFIX AwsIam
    - ResourceType PREFIX AwsEc2
    - ResourceType NOT_EQUALS AwsIamPolicy
    - ResourceType NOT_EQUALS AwsEc2NetworkInterface
  CONTAINS and NOT_CONTAINS operators can be used only with automation rules. For
  more information, see Automation rules (https://docs.aws.amazon.com/securityhub/latest/userguide/automation-rules.html)
  in the Security Hub User Guide.

* `Value` (`string`) (nullable)

  The string filter value. Filter values are case sensitive. For example, the
  product name for control-based findings is Security Hub . If you provide
  security hub as the filter value, there's no match.

###### <a name="NumberFilter-1"></a>NumberFilter

  A number filter for querying findings.

* `Eq` (`number`)

  The equal-to condition to be applied to a single field when querying for
  findings.

* `Gte` (`number`)

  The greater-than-equal condition to be applied to a single field when querying
  for findings.

* `Lte` (`number`)

  The less-than-equal condition to be applied to a single field when querying for
  findings.

###### <a name="DateFilter-1"></a>DateFilter

  A date filter for querying findings.

* `DateRange` ([`DateRange`](#DateRange)) (nullable)

  A date range for the date filter.

* `End` (`string`) (nullable)

  A timestamp that provides the end date for the date filter. A correctly
  formatted example is 2020-05-21T20:16:34.724Z . The value cannot contain spaces,
  and date and time should be separated by T . For more information, see RFC 3339
  section 5.6, Internet Date/Time Format (https://www.rfc-editor.org/rfc/rfc3339#section-5.6)
  .

* `Start` (`string`) (nullable)

  A timestamp that provides the start date for the date filter. A correctly
  formatted example is 2020-05-21T20:16:34.724Z . The value cannot contain spaces,
  and date and time should be separated by T . For more information, see RFC 3339
  section 5.6, Internet Date/Time Format (https://www.rfc-editor.org/rfc/rfc3339#section-5.6)
  .

###### <a name="DateRange"></a>DateRange

  A date range for the date filter.

* `Unit` (`string`)

  A date range unit for the date filter.

* `Value` (`integer`)

  A date range value for the date filter.

###### <a name="KeywordFilter"></a>KeywordFilter

  A keyword filter for querying findings.

* `Value` (`string`) (nullable)

  A value for the keyword.

###### <a name="IpFilter"></a>IpFilter

  The IP filter for querying findings.

* `Cidr` (`string`) (nullable)

  A finding's CIDR value.

###### <a name="MapFilter-1"></a>MapFilter

  A map filter for filtering Security Hub findings.

* `Comparison` (`string`)

  The condition to apply to the key value when filtering Security Hub findings
  with a map filter. To search for values that have the filter value, use one of
  the following comparison operators:
    - To search for values that include the filter value, use CONTAINS . For
    example, for the ResourceTags field, the filter Department CONTAINS Security
    matches findings that include the value Security for the Department tag. In
    the same example, a finding with a value of Security team for the Department
    tag is a match.
    - To search for values that exactly match the filter value, use EQUALS . For
    example, for the ResourceTags field, the filter Department EQUALS Security
    matches findings that have the value Security for the Department tag.
  CONTAINS and EQUALS filters on the same field are joined by OR . A finding
  matches if it matches any one of those filters. For example, the filters
  Department CONTAINS Security OR Department CONTAINS Finance match a finding that
  includes either Security , Finance , or both values. To search for values that
  don't have the filter value, use one of the following comparison operators:
    - To search for values that exclude the filter value, use NOT_CONTAINS . For
    example, for the ResourceTags field, the filter Department NOT_CONTAINS
    Finance matches findings that exclude the value Finance for the Department
    tag.
    - To search for values other than the filter value, use NOT_EQUALS . For
    example, for the ResourceTags field, the filter Department NOT_EQUALS Finance
    matches findings that don’t have the value Finance for the Department tag.
  NOT_CONTAINS and NOT_EQUALS filters on the same field are joined by AND . A
  finding matches only if it matches all of those filters. For example, the
  filters Department NOT_CONTAINS Security AND Department NOT_CONTAINS Finance
  match a finding that excludes both the Security and Finance values. CONTAINS
  filters can only be used with other CONTAINS filters. NOT_CONTAINS filters can
  only be used with other NOT_CONTAINS filters. You can’t have both a CONTAINS
  filter and a NOT_CONTAINS filter on the same field. Similarly, you can’t have
  both an EQUALS filter and a NOT_EQUALS filter on the same field. Combining
  filters in this way returns an error. CONTAINS and NOT_CONTAINS operators can
  be used only with automation rules. For more information, see Automation rules (https://docs.aws.amazon.com/securityhub/latest/userguide/automation-rules.html)
  in the Security Hub User Guide.

* `Key` (`string`) (nullable)

  The key of the map filter. For example, for ResourceTags , Key identifies the
  name of the tag. For UserDefinedFields , Key is the name of the field.

* `Value` (`string`) (nullable)

  The value for the key in the map filter. Filter values are case sensitive. For
  example, one of the values for a tag called Department might be Security . If
  you provide security as the filter value, then there's no match.

###### <a name="BooleanFilter"></a>BooleanFilter

  Boolean filter for querying findings.

* `Value` (`boolean`)

  The value of the boolean.

###### <a name="SortCriterion"></a>SortCriterion

  A collection of finding attributes used to sort findings.

* `Field` (`string`) (nullable)

  The finding attribute used to sort findings.

* `SortOrder` (`string`)

  The order used to sort findings.

### <a name="EventBasedSync"></a>EventBasedSync

* `full_sync` (`boolean`) (nullable) (default: `true`)

* `account` ([`Account`](#Account))

* `kinesis_stream_arn` (`string`) (required) ([pattern](https://json-schema.org/draft/2020-12/json-schema-validation#section-6.3.3): `^arn(:[^:\n]*){5}([:/].*)?$`)

* `start_time` (`string`) (nullable) ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `date-time`) (default: `now`)

### <a name="Strategy"></a>Strategy

CloudQuery scheduling strategy

(`string`) (possible values: `dfs`, `round-robin`, `shuffle`) (default: `dfs`)
