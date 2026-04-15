# ExportInfoExportType

The type of export to perform

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ExportInfoExportTypeFindings

// Open enum: custom values can be created with a direct type cast
custom := components.ExportInfoExportType("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ExportInfoExportTypeFindings`  | FINDINGS                        |
| `ExportInfoExportTypeDocuments` | DOCUMENTS                       |
| `ExportInfoExportTypeIssues`    | ISSUES                          |