# PlatformMessageRole

Role of the message author.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformMessageRoleUser

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformMessageRole("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `PlatformMessageRoleUser`    | USER                         |
| `PlatformMessageRoleGleanAi` | GLEAN_AI                     |