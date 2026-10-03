# PlatformSkillStatus

The caller's effective activation. The owner sees the skill's stored status. Another caller sees their personal setting, or DISABLED when they have none. DRAFT is the stored draft state and is not set by update. Effective activation may also reflect workspace governance policy.


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