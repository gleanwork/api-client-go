# PlacementReason

Placement source for ranked feed results. ORGANIC means the card was emitted by normal feed ranking. PROMO means the card was inserted by the homepage cards promo framework.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlacementReasonOrganic

// Open enum: custom values can be created with a direct type cast
custom := components.PlacementReason("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `PlacementReasonOrganic` | ORGANIC                  |
| `PlacementReasonPromo`   | PROMO                    |