# GrantScope

Scope of the approval grant. Only applicable when isGranted is true and requestType is EXECUTION.


## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.GrantScopeCurrentRequest

// Open enum: custom values can be created with a direct type cast
custom := components.GrantScope("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `GrantScopeCurrentRequest` | CURRENT_REQUEST            |
| `GrantScopeCurrentSession` | CURRENT_SESSION            |
| `GrantScopeAlways`         | ALWAYS                     |