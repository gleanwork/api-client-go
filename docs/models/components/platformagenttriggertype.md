# PlatformAgentTriggerType

Invocation mode of the agent.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformAgentTriggerTypeChatMessage

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformAgentTriggerType("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `PlatformAgentTriggerTypeChatMessage`    | CHAT_MESSAGE                             |
| `PlatformAgentTriggerTypeInputForm`      | INPUT_FORM                               |
| `PlatformAgentTriggerTypeContentTrigger` | CONTENT_TRIGGER                          |