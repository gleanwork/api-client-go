# ServerToolRequestRequestType

The type of request made to the user.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ServerToolRequestRequestTypeExecution

// Open enum: custom values can be created with a direct type cast
custom := components.ServerToolRequestRequestType("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `ServerToolRequestRequestTypeExecution`                | EXECUTION                                              |
| `ServerToolRequestRequestTypeAuthenticationSuggestion` | AUTHENTICATION_SUGGESTION                              |
| `ServerToolRequestRequestTypeVoteSuggestion`           | VOTE_SUGGESTION                                        |
| `ServerToolRequestRequestTypeSandboxEgress`            | SANDBOX_EGRESS                                         |