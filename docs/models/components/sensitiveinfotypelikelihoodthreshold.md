# ~~SensitiveInfoTypeLikelihoodThreshold~~

> :warning: **DEPRECATED**: Deprecated on 2026-02-05, removal scheduled for 2027-01-15: Field is deprecated.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.SensitiveInfoTypeLikelihoodThresholdLikely

// Open enum: custom values can be created with a direct type cast
custom := components.SensitiveInfoTypeLikelihoodThreshold("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `SensitiveInfoTypeLikelihoodThresholdLikely`       | LIKELY                                             |
| `SensitiveInfoTypeLikelihoodThresholdVeryLikely`   | VERY_LIKELY                                        |
| `SensitiveInfoTypeLikelihoodThresholdPossible`     | POSSIBLE                                           |
| `SensitiveInfoTypeLikelihoodThresholdUnlikely`     | UNLIKELY                                           |
| `SensitiveInfoTypeLikelihoodThresholdVeryUnlikely` | VERY_UNLIKELY                                      |