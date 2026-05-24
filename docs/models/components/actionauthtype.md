# ActionAuthType

Authentication mechanism used by an action pack.
  - `AUTH_USER_OAUTH`: Requires per-user OAuth consent to the third-party tool.
  - `AUTH_ADMIN`: Uses a service-account / admin-owned credential. End users do not authorize individually.
  - `AUTH_NONE`: Action pack requires no authentication.


## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ActionAuthTypeAuthUserOauth

// Open enum: custom values can be created with a direct type cast
custom := components.ActionAuthType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `ActionAuthTypeAuthUserOauth` | AUTH_USER_OAUTH               |
| `ActionAuthTypeAuthAdmin`     | AUTH_ADMIN                    |
| `ActionAuthTypeAuthNone`      | AUTH_NONE                     |