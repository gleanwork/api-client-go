# PlatformFilterOperator

Supported filter operator.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformFilterOperatorEquals

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformFilterOperator("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PlatformFilterOperatorEquals`    | EQUALS                            |
| `PlatformFilterOperatorNotEquals` | NOT_EQUALS                        |
| `PlatformFilterOperatorGt`        | GT                                |
| `PlatformFilterOperatorGte`       | GTE                               |
| `PlatformFilterOperatorLt`        | LT                                |
| `PlatformFilterOperatorLte`       | LTE                               |