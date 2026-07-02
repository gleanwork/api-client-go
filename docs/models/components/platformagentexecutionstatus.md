# PlatformAgentExecutionStatus

Status of the agent run.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformAgentExecutionStatusError

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformAgentExecutionStatus("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `PlatformAgentExecutionStatusError`   | error                                 |
| `PlatformAgentExecutionStatusSuccess` | success                               |