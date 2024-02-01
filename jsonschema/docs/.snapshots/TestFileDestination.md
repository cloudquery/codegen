# Table of contents

* [`Spec`](#Spec)
  * [`CSVSpec`](#CSVSpec)
  * [`JSONSpec`](#JSONSpec)
  * [`ParquetSpec`](#ParquetSpec)
  * [`Duration`](#Duration)

## <a name="Spec"></a>Spec

* `format` (`string`) (required) (possible values: `csv`, `json`, `parquet`)

  Output format.

* `format_spec` ([`CSVSpec`](#CSVSpec), [`JSONSpec`](#JSONSpec) or [`ParquetSpec`](#ParquetSpec)) (nullable)

* `compression` (`string`) (possible values: ` `, `gzip`)

  Compression type.
  Empty or missing stands for no compression.

* `path` (`string`) (required)

  Path template string that determines where files will be written.
  
  The path supports the following placeholder variables:
  - `{{TABLE}}` will be replaced with the table name
  - `{{FORMAT}}` will be replaced with the file format, such as `csv`, `json` or `parquet`. If compression is enabled, the format will be `csv.gz`, `json.gz` etc.
  - `{{UUID}}` will be replaced with a random UUID to uniquely identify each file
  - `{{YEAR}}` will be replaced with the current year in `YYYY` format
  - `{{MONTH}}` will be replaced with the current month in `MM` format
  - `{{DAY}}` will be replaced with the current day in `DD` format
  - `{{HOUR}}` will be replaced with the current hour in `HH` format
  - `{{MINUTE}}` will be replaced with the current minute in `mm` format
  
   **Note** that timestamps are in `UTC` and will be the current time at the time the file is written, not when the sync started.

* `no_rotate` (`boolean`) (default: `false`)

  If set to `true`, the plugin will write to one file per table.
  Otherwise, for every batch a new file will be created with a different `.<UUID>` suffix.

* `batch_size` (`integer`) (nullable) (range: `[1,+∞)`) (default: `10000`)

  This parameter controls the maximum amount of items may be grouped together to be written in a single write.
  
  Defaults to `10000` unless `no_rotate` is `true` (will be `0` then).

* `batch_size_bytes` (`integer`) (nullable) (range: `[1,+∞)`) (default: `52428800`)

  This parameter controls the maximum size of items that may be grouped together to be written in a single write.
  
  Defaults to `52428800` (50 MiB) unless `no_rotate` is `true` (will be `0` then).

* `batch_timeout` ([`Duration`](#Duration)) (nullable) (default: `30s`)

  This parameter controls the maximum interval between batch writes.
  
  Defaults to `30s` unless `no_rotate` is `true` (will be `0s` then).

### <a name="CSVSpec"></a>CSVSpec

  CloudQuery CSV file output spec.

* `skip_header` (`boolean`) (default: `false`)

  Specifies if the first line of a file should be the header.

* `delimiter` (`string`) ([pattern](https://json-schema.org/draft/2020-12/json-schema-validation#section-6.3.3): `^.$`) (default: `,`)

  Character that will be used as the delimiter.

### <a name="JSONSpec"></a>JSONSpec

  CloudQuery JSON file output spec.

(`object`)

### <a name="ParquetSpec"></a>ParquetSpec

  CloudQuery Parquet file output spec.

(`object`)

### <a name="Duration"></a>Duration

CloudQuery configtype.Duration

(`string`) ([pattern](https://json-schema.org/draft/2020-12/json-schema-validation#section-6.3.3): `^[-+]?([0-9]*(\\.[0-9]*)?[a-z]+)+$`)
