# ServerToolResponseRequestType

The type of request made to the user.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ServerToolResponseRequestTypeExecution

// Open enum: custom values can be created with a direct type cast
custom := components.ServerToolResponseRequestType("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `ServerToolResponseRequestTypeExecution`                | EXECUTION                                               |
| `ServerToolResponseRequestTypeAuthenticationSuggestion` | AUTHENTICATION_SUGGESTION                               |
| `ServerToolResponseRequestTypeVoteSuggestion`           | VOTE_SUGGESTION                                         |