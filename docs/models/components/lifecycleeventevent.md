# LifeCycleEventEvent

Type of event

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.LifeCycleEventEventUploaded

// Open enum: custom values can be created with a direct type cast
custom := components.LifeCycleEventEvent("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `LifeCycleEventEventUploaded`          | UPLOADED                               |
| `LifeCycleEventEventIndexed`           | INDEXED                                |
| `LifeCycleEventEventDeletionRequested` | DELETION_REQUESTED                     |
| `LifeCycleEventEventDeleted`           | DELETED                                |