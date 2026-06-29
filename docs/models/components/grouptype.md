# GroupType

The type of user group

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.GroupTypeDepartment

// Open enum: custom values can be created with a direct type cast
custom := components.GroupType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `GroupTypeDepartment`         | DEPARTMENT                    |
| `GroupTypeAll`                | ALL                           |
| `GroupTypeTeam`               | TEAM                          |
| `GroupTypeJobTitle`           | JOB_TITLE                     |
| `GroupTypeRoleType`           | ROLE_TYPE                     |
| `GroupTypeLocation`           | LOCATION                      |
| `GroupTypeRegion`             | REGION                        |
| `GroupTypeExternalGroup`      | EXTERNAL_GROUP                |
| `GroupTypeCollectionAudience` | COLLECTION_AUDIENCE           |