# IdpGroups

How to resolve effective usage limits when a user belongs to multiple IdP groups.


## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.IdpGroupsHighest

// Open enum: custom values can be created with a direct type cast
custom := components.IdpGroups("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `IdpGroupsHighest` | HIGHEST            |
| `IdpGroupsLowest`  | LOWEST             |