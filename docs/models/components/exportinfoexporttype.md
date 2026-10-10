# ExportInfoExportType

The type of export to perform. FINDINGS, DOCUMENTS and ISSUES produce JSONL; FINDINGS_CSV produces one CSV row per finding.

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

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `ExportInfoExportTypeFindings`    | FINDINGS                          |
| `ExportInfoExportTypeDocuments`   | DOCUMENTS                         |
| `ExportInfoExportTypeIssues`      | ISSUES                            |
| `ExportInfoExportTypeFindingsCsv` | FINDINGS_CSV                      |