# BusinessJustification

Whether a business justification is required when requesting a usage limit increase.


## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.BusinessJustificationOptional

// Open enum: custom values can be created with a direct type cast
custom := components.BusinessJustification("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `BusinessJustificationOptional` | OPTIONAL                        |
| `BusinessJustificationRequired` | REQUIRED                        |