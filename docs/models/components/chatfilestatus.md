# ChatFileStatus

Current status of the file.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ChatFileStatusProcessing

// Open enum: custom values can be created with a direct type cast
custom := components.ChatFileStatus("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `ChatFileStatusProcessing`         | PROCESSING                         |
| `ChatFileStatusProcessed`          | PROCESSED                          |
| `ChatFileStatusPartiallyProcessed` | PARTIALLY_PROCESSED                |
| `ChatFileStatusFailed`             | FAILED                             |
| `ChatFileStatusDeleted`            | DELETED                            |