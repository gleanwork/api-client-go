# ChatStatus

The status of the Chat.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ChatStatusSuccess

// Open enum: custom values can be created with a direct type cast
custom := components.ChatStatus("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `ChatStatusSuccess`    | SUCCESS                |
| `ChatStatusError`      | ERROR                  |
| `ChatStatusInProgress` | IN_PROGRESS            |
| `ChatStatusCancelled`  | CANCELLED              |