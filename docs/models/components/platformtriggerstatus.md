# PlatformTriggerStatus

Current trigger lifecycle state.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformTriggerStatusEnabled

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformTriggerStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `PlatformTriggerStatusEnabled`  | ENABLED                         |
| `PlatformTriggerStatusDisabled` | DISABLED                        |