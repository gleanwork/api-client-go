# PlatformOAuthClientFilterOperator

The only supported operator is EQUALS. Defaults to EQUALS.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformOAuthClientFilterOperatorEquals
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `PlatformOAuthClientFilterOperatorEquals`    | EQUALS                                       |
| `PlatformOAuthClientFilterOperatorNotEquals` | NOT_EQUALS                                   |
| `PlatformOAuthClientFilterOperatorGt`        | GT                                           |
| `PlatformOAuthClientFilterOperatorGte`       | GTE                                          |
| `PlatformOAuthClientFilterOperatorLt`        | LT                                           |
| `PlatformOAuthClientFilterOperatorLte`       | LTE                                          |