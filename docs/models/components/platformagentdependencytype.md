# PlatformAgentDependencyType

Type of a published artifact referenced as a dependency.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformAgentDependencyTypeSkill

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformAgentDependencyType("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `PlatformAgentDependencyTypeSkill` | SKILL                              |
| `PlatformAgentDependencyTypeAgent` | AGENT                              |