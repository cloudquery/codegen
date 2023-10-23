# Table of contents

* [`FileSpec`](#FileSpec)
  * [`CSVSpec`](#CSVSpec)
  * [`JSONSpec`](#JSONSpec)
  * [`ParquetSpec`](#ParquetSpec)

## <a name="FileSpec"></a>FileSpec

* `format` (`string`) (required) (possible values: `csv`, `json`, `parquet`)

  Output format.

* `format_spec` ([`CSVSpec`](#CSVSpec), [`JSONSpec`](#JSONSpec) or [`ParquetSpec`](#ParquetSpec)) (nullable)

* `compression` (`string`) (possible values: ` `, `gzip`)

  Compression type.
  Empty or missing stands for no compression.

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
