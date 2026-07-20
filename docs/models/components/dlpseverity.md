# DlpSeverity

Severity levels for DLP findings and analyses. FALSE_POSITIVE ranks below LOW and marks analyses that concluded every flagged entity is a detector false positive.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.DlpSeverityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := components.DlpSeverity("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `DlpSeverityUnspecified`   | UNSPECIFIED                |
| `DlpSeverityLow`           | LOW                        |
| `DlpSeverityMedium`        | MEDIUM                     |
| `DlpSeverityHigh`          | HIGH                       |
| `DlpSeverityFalsePositive` | FALSE_POSITIVE             |