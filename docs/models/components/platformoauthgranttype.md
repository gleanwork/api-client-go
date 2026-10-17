# PlatformOAuthGrantType

An OAuth grant type.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformOAuthGrantTypeAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformOAuthGrantType("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `PlatformOAuthGrantTypeAuthorizationCode` | AUTHORIZATION_CODE                        |
| `PlatformOAuthGrantTypeClientCredentials` | CLIENT_CREDENTIALS                        |
| `PlatformOAuthGrantTypeRefreshToken`      | REFRESH_TOKEN                             |
| `PlatformOAuthGrantTypeDeviceCode`        | DEVICE_CODE                               |
| `PlatformOAuthGrantTypeTokenExchange`     | TOKEN_EXCHANGE                            |
| `PlatformOAuthGrantTypeJwtBearer`         | JWT_BEARER                                |