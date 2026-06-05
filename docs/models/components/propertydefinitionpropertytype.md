# PropertyDefinitionPropertyType

The type of custom property - this governs the search and faceting behavior. Note that MULTIPICKLIST is not yet supported.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PropertyDefinitionPropertyTypeText

// Open enum: custom values can be created with a direct type cast
custom := components.PropertyDefinitionPropertyType("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `PropertyDefinitionPropertyTypeText`          | TEXT                                          |
| `PropertyDefinitionPropertyTypeDate`          | DATE                                          |
| `PropertyDefinitionPropertyTypeInt`           | INT                                           |
| `PropertyDefinitionPropertyTypeUserid`        | USERID                                        |
| `PropertyDefinitionPropertyTypePicklist`      | PICKLIST                                      |
| `PropertyDefinitionPropertyTypeTextlist`      | TEXTLIST                                      |
| `PropertyDefinitionPropertyTypeMultipicklist` | MULTIPICKLIST                                 |