# PlatformSkillStatus

Current skill status.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformSkillStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformSkillStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `PlatformSkillStatusDraft`    | DRAFT                         |
| `PlatformSkillStatusEnabled`  | ENABLED                       |
| `PlatformSkillStatusDisabled` | DISABLED                      |