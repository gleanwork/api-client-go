# PlatformAgentRunState

State of the persisted workflow execution. REQUIRES_INPUT is nonterminal.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformAgentRunStateQueued

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformAgentRunState("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `PlatformAgentRunStateQueued`        | QUEUED                               |
| `PlatformAgentRunStateRunning`       | RUNNING                              |
| `PlatformAgentRunStateRequiresInput` | REQUIRES_INPUT                       |
| `PlatformAgentRunStateSucceeded`     | SUCCEEDED                            |
| `PlatformAgentRunStateFailed`        | FAILED                               |
| `PlatformAgentRunStateCancelling`    | CANCELLING                           |
| `PlatformAgentRunStateCancelled`     | CANCELLED                            |
| `PlatformAgentRunStateExpired`       | EXPIRED                              |