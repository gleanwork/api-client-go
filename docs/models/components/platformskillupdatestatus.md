# PlatformSkillUpdateStatus

Activation to apply for the authenticated caller. For the owner, this updates the skill's stored status. For any other caller, it updates only that caller's setting.


## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformSkillUpdateStatusEnabled
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `PlatformSkillUpdateStatusEnabled`  | ENABLED                             |
| `PlatformSkillUpdateStatusDisabled` | DISABLED                            |