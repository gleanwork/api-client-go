# CollectionErrorErrorCode

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.CollectionErrorErrorCodeNameExists

// Open enum: custom values can be created with a direct type cast
custom := components.CollectionErrorErrorCode("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `CollectionErrorErrorCodeNameExists`              | NAME_EXISTS                                       |
| `CollectionErrorErrorCodeNotFound`                | NOT_FOUND                                         |
| `CollectionErrorErrorCodeCollectionPinned`        | COLLECTION_PINNED                                 |
| `CollectionErrorErrorCodeConcurrentHierarchyEdit` | CONCURRENT_HIERARCHY_EDIT                         |
| `CollectionErrorErrorCodeHeightViolation`         | HEIGHT_VIOLATION                                  |
| `CollectionErrorErrorCodeWidthViolation`          | WIDTH_VIOLATION                                   |
| `CollectionErrorErrorCodeNoPermissions`           | NO_PERMISSIONS                                    |
| `CollectionErrorErrorCodeCorruptItem`             | CORRUPT_ITEM                                      |