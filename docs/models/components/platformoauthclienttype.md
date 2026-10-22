# PlatformOAuthClientType

Whether an OAuth client can keep a secret.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformOAuthClientTypeConfidential

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformOAuthClientType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `PlatformOAuthClientTypeConfidential` | CONFIDENTIAL                          |
| `PlatformOAuthClientTypePublic`       | PUBLIC                                |