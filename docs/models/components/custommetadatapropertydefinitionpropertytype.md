# CustomMetadataPropertyDefinitionPropertyType

The type of metadata key. This governs the search and faceting behavior.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.CustomMetadataPropertyDefinitionPropertyTypeText

// Open enum: custom values can be created with a direct type cast
custom := components.CustomMetadataPropertyDefinitionPropertyType("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `CustomMetadataPropertyDefinitionPropertyTypeText`          | TEXT                                                        |
| `CustomMetadataPropertyDefinitionPropertyTypePicklist`      | PICKLIST                                                    |
| `CustomMetadataPropertyDefinitionPropertyTypeTextlist`      | TEXTLIST                                                    |
| `CustomMetadataPropertyDefinitionPropertyTypeMultipicklist` | MULTIPICKLIST                                               |