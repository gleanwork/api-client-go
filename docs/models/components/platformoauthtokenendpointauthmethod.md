# PlatformOAuthTokenEndpointAuthMethod

How an OAuth client authenticates to the token endpoint.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformOAuthTokenEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformOAuthTokenEndpointAuthMethod("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `PlatformOAuthTokenEndpointAuthMethodClientSecretBasic` | CLIENT_SECRET_BASIC                                     |
| `PlatformOAuthTokenEndpointAuthMethodClientSecretPost`  | CLIENT_SECRET_POST                                      |
| `PlatformOAuthTokenEndpointAuthMethodNoAuthentication`  | NO_AUTHENTICATION                                       |
| `PlatformOAuthTokenEndpointAuthMethodPrivateKeyJwt`     | PRIVATE_KEY_JWT                                         |
| `PlatformOAuthTokenEndpointAuthMethodTLSClientAuth`     | TLS_CLIENT_AUTH                                         |