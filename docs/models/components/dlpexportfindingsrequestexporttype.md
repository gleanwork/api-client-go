# DlpExportFindingsRequestExportType

The type of export to perform. FINDINGS, DOCUMENTS and ISSUES produce JSONL; FINDINGS_CSV produces one CSV row per finding.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.DlpExportFindingsRequestExportTypeFindings
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `DlpExportFindingsRequestExportTypeFindings`    | FINDINGS                                        |
| `DlpExportFindingsRequestExportTypeDocuments`   | DOCUMENTS                                       |
| `DlpExportFindingsRequestExportTypeIssues`      | ISSUES                                          |
| `DlpExportFindingsRequestExportTypeFindingsCsv` | FINDINGS_CSV                                    |