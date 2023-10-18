# Table of contents

* [`Spec`](#Spec)
  * [`Engine`](#Engine)
  * [`Duration`](#Duration)

## <a name="Spec"></a>Spec

* `connection_string` (`string`) (required)

* `cluster` (`string`)

* `engine` ([`Engine`](#Engine)) (nullable)

* `ca_cert` (`string`)

* `batch_size` (`integer`) (range: `[1,+∞)`) (default: `10000`)

* `batch_size_bytes` (`integer`) (range: `[1,+∞)`) (default: `5242880`)

* `batch_timeout` ([`Duration`](#Duration)) (nullable) (default: `20s`)

### <a name="Engine"></a>Engine

* `name` (`string`) ([pattern](https://json-schema.org/draft/2020-12/json-schema-validation#section-6.3.3): `^.*MergeTree$`) (default: `MergeTree`)

* `parameters` (`[]anything`) (nullable)

### <a name="Duration"></a>Duration

CloudQuery configtype.Duration

(`string`) ([pattern](https://json-schema.org/draft/2020-12/json-schema-validation#section-6.3.3): `^[-+]?([0-9]*(\\.[0-9]*)?[a-z]+)+$`)
