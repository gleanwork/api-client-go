# DatasourceCredentialStatus

Lifecycle state of the credentials installed for a datasource instance. Mirrors the internal admin Status enum so the handler can surface the same health signals already tracked today. EXPIRING_SOON is represented as VALID_WITH_WARNINGS (with detail in `message`); EXPIRED is surfaced as INVALID plus a non-null `expiresAt` in the past.


## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.DatasourceCredentialStatusValid

// Open enum: custom values can be created with a direct type cast
custom := components.DatasourceCredentialStatus("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `DatasourceCredentialStatusValid`             | VALID                                         |
| `DatasourceCredentialStatusValidWithWarnings` | VALID_WITH_WARNINGS                           |
| `DatasourceCredentialStatusValidating`        | VALIDATING                                    |
| `DatasourceCredentialStatusInvalid`           | INVALID                                       |
| `DatasourceCredentialStatusMissing`           | MISSING                                       |