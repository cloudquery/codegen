# Table of contents

* [`Spec`](#Spec)
  * [`Engine`](#Engine)
  * [`Duration`](#Duration)

## <a name="Spec"></a>Spec

  CloudQuery ClickHouse destination plugin spec.

* `connection_string` (`string`) (required)

  Connection string to connect to the database.
  See [SDK documentation](https://github.com/ClickHouse/clickhouse-go#dsn) for more details.

* `cluster` (`string`)

  Cluster name to be used for [distributed DDL](https://clickhouse.com/docs/en/sql-reference/distributed-ddl).
  If the value is empty, DDL operations will affect only the server the plugin is connected to.

* `engine` ([`Engine`](#Engine)) (nullable)

  Engine to be used for tables.
  Only [`*MergeTree` family](https://clickhouse.com/docs/en/engines/table-engines/mergetree-family) is supported at the moment.

* `ca_cert` (`string`)

  PEM-encoded certificate authorities.
  When set, a certificate pool will be created by appending the certificates to the system pool.
  
  See [file variable substitution](/docs/advanced-topics/environment-variable-substitution#file-variable-substitution-example)
  for how to read this value from a file.

* `batch_size` (`integer`) (range: `[1,+∞)`) (default: `10000`)

  This parameter controls the maximum amount of items may be grouped together to be written as a single write.

* `batch_size_bytes` (`integer`) (range: `[1,+∞)`) (default: `5242880`)

  This parameter controls the maximum size of items that may be grouped together to be written as a single write.

* `batch_timeout` ([`Duration`](#Duration)) (nullable) (default: `20s`)

  This parameter controls the maximum interval between batch writes.

### <a name="Engine"></a>Engine

  Engine allows to specify a custom table engine to be used.

* `name` (`string`) ([pattern](https://json-schema.org/draft/2020-12/json-schema-validation#section-6.3.3): `^.*MergeTree$`) (default: `MergeTree`)

  Name of the table engine.
  Only [`*MergeTree` family](https://clickhouse.com/docs/en/engines/table-engines/mergetree-family) is supported at the moment.

* `parameters` (`[]anything`) (nullable)

  Engine parameters.
  Currently, no restrictions are imposed on the parameter types.

### <a name="Duration"></a>Duration

CloudQuery configtype.Duration

(`string`) ([pattern](https://json-schema.org/draft/2020-12/json-schema-validation#section-6.3.3): `^[-+]?([0-9]*(\\.[0-9]*)?[a-z]+)+$`)
