# EditCollectionResponseErrorCode

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.EditCollectionResponseErrorCodeNameExists

// Open enum: custom values can be created with a direct type cast
custom := components.EditCollectionResponseErrorCode("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `EditCollectionResponseErrorCodeNameExists`              | NAME_EXISTS                                              |
| `EditCollectionResponseErrorCodeNotFound`                | NOT_FOUND                                                |
| `EditCollectionResponseErrorCodeCollectionPinned`        | COLLECTION_PINNED                                        |
| `EditCollectionResponseErrorCodeConcurrentHierarchyEdit` | CONCURRENT_HIERARCHY_EDIT                                |
| `EditCollectionResponseErrorCodeHeightViolation`         | HEIGHT_VIOLATION                                         |
| `EditCollectionResponseErrorCodeWidthViolation`          | WIDTH_VIOLATION                                          |
| `EditCollectionResponseErrorCodeNoPermissions`           | NO_PERMISSIONS                                           |
| `EditCollectionResponseErrorCodeCorruptItem`             | CORRUPT_ITEM                                             |