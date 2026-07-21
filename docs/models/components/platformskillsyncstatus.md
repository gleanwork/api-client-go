# PlatformSkillSyncStatus

Current external-source sync status.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformSkillSyncStatusUpToDate

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformSkillSyncStatus("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `PlatformSkillSyncStatusUpToDate`        | UP_TO_DATE                               |
| `PlatformSkillSyncStatusUpdateAvailable` | UPDATE_AVAILABLE                         |
| `PlatformSkillSyncStatusSyncFailed`      | SYNC_FAILED                              |