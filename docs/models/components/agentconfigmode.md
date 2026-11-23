# AgentConfigMode

Top level modes to run GleanChat in.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.AgentConfigModeDefault

// Open enum: custom values can be created with a direct type cast
custom := components.AgentConfigMode("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `AgentConfigModeDefault` | DEFAULT                  |
| `AgentConfigModeQuick`   | QUICK                    |