# PlatformSkillSourcePreviewFailureCode

Stable machine-readable reason a discovered entry was excluded.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformSkillSourcePreviewFailureCodeInvalidSkill

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformSkillSourcePreviewFailureCode("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `PlatformSkillSourcePreviewFailureCodeInvalidSkill`     | INVALID_SKILL                                           |
| `PlatformSkillSourcePreviewFailureCodeSkillFetchFailed` | SKILL_FETCH_FAILED                                      |