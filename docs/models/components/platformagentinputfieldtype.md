# PlatformAgentInputFieldType

Type of an agent input field.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformAgentInputFieldTypeText

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformAgentInputFieldType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `PlatformAgentInputFieldTypeText`     | TEXT                                  |
| `PlatformAgentInputFieldTypeSelect`   | SELECT                                |
| `PlatformAgentInputFieldTypeDate`     | DATE                                  |
| `PlatformAgentInputFieldTypeDocument` | DOCUMENT                              |