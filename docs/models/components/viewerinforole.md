# ~~ViewerInfoRole~~

DEPRECATED - use permissions instead. Viewer's role on the specific document.

> :warning: **DEPRECATED**: Deprecated on 2026-02-05, removal scheduled for 2026-10-15: Use permissions instead.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ViewerInfoRoleAnswerModerator

// Open enum: custom values can be created with a direct type cast
custom := components.ViewerInfoRole("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ViewerInfoRoleAnswerModerator` | ANSWER_MODERATOR                |
| `ViewerInfoRoleOwner`           | OWNER                           |
| `ViewerInfoRoleViewer`          | VIEWER                          |