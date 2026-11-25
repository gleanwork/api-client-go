# PlatformAgentContentTriggerFilterOperator

Operator of a content trigger filter.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformAgentContentTriggerFilterOperatorIsOneOf

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformAgentContentTriggerFilterOperator("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `PlatformAgentContentTriggerFilterOperatorIsOneOf`      | IS_ONE_OF                                               |
| `PlatformAgentContentTriggerFilterOperatorIsNotOneOf`   | IS_NOT_ONE_OF                                           |
| `PlatformAgentContentTriggerFilterOperatorContains`     | CONTAINS                                                |
| `PlatformAgentContentTriggerFilterOperatorFieldChanged` | FIELD_CHANGED                                           |