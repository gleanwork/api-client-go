# AuthHeaderType

Defines the header structure for sending the API key or token to the server. Defaults to AUTHORIZATION_BEARER. Select the specific header format the server expects for transmitting the key.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.AuthHeaderTypeAuthorizationBearer

// Open enum: custom values can be created with a direct type cast
custom := components.AuthHeaderType("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `AuthHeaderTypeAuthorizationBearer` | AUTHORIZATION_BEARER                |
| `AuthHeaderTypeAuthorizationToken`  | AUTHORIZATION_TOKEN                 |
| `AuthHeaderTypeAuthorizationAPIKey` | AUTHORIZATION_API_KEY               |
| `AuthHeaderTypeXAPIKey`             | X_API_KEY                           |