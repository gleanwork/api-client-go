# DlpIssueStatus

Status of a DLP issue.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.DlpIssueStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := components.DlpIssueStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `DlpIssueStatusOpen`       | OPEN                       |
| `DlpIssueStatusClosed`     | CLOSED                     |
| `DlpIssueStatusInProgress` | IN_PROGRESS                |
| `DlpIssueStatusResolved`   | RESOLVED                   |