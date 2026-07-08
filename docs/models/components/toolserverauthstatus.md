# ToolServerAuthStatus

Authentication status for the calling user.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ToolServerAuthStatusAwaitingAuth

// Open enum: custom values can be created with a direct type cast
custom := components.ToolServerAuthStatus("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `ToolServerAuthStatusAwaitingAuth` | AWAITING_AUTH                      |
| `ToolServerAuthStatusAuthorized`   | AUTHORIZED                         |