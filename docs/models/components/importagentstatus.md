# ImportAgentStatus

Outcome of the import for the target agent.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ImportAgentStatusCreated

// Open enum: custom values can be created with a direct type cast
custom := components.ImportAgentStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ImportAgentStatusCreated`      | CREATED                         |
| `ImportAgentStatusUpdated`      | UPDATED                         |
| `ImportAgentStatusDraftPreview` | DRAFT_PREVIEW                   |