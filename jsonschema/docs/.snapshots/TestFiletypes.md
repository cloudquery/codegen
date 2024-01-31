# Table of contents

* [`FileSpec`](#FileSpec)
  * [`Spec`](#Spec)
  * [`Spec`](#Spec-1)
  * [`Spec`](#Spec-2)

## <a name="FileSpec"></a>FileSpec

* `format` (`string`) (required) (possible values: `csv`, `json`, `parquet`)

  Output format.

* `format_spec` ([`Spec`](#Spec), [`Spec`](#Spec-1) or [`Spec`](#Spec-2)) (nullable)

* `compression` (`string`) (possible values: ` `, `gzip`)

  Compression type.
  Empty or missing stands for no compression.

### <a name="Spec"></a>Spec

  CloudQuery CSV file output spec.

* `skip_header` (`boolean`) (default: `false`)

  Specifies if the first line of a file should be the header.

* `delimiter` (`string`) ([pattern](https://json-schema.org/draft/2020-12/json-schema-validation#section-6.3.3): `^.$`) (default: `,`)

  Character that will be used as the delimiter.

### <a name="Spec-1"></a>Spec

  CloudQuery JSON file output spec.

(`object`)

### <a name="Spec-2"></a>Spec

  CloudQuery Parquet file output spec.

(`object`)
