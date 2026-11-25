# MessagesRequestDirection

The direction of the results asked with respect to the reference timestamp. Missing field defaults to OLDER. Only applicable when using a message_id.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.MessagesRequestDirectionOlder
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `MessagesRequestDirectionOlder` | OLDER                           |
| `MessagesRequestDirectionNewer` | NEWER                           |