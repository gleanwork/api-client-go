# PlatformPersonType

Person employment status or account type.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformPersonTypeFullTime

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformPersonType("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `PlatformPersonTypeFullTime`       | FULL_TIME                          |
| `PlatformPersonTypeContractor`     | CONTRACTOR                         |
| `PlatformPersonTypeNonEmployee`    | NON_EMPLOYEE                       |
| `PlatformPersonTypeFormerEmployee` | FORMER_EMPLOYEE                    |