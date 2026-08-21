# PlatformTriggerEventReason

Why the event fired.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformTriggerEventReasonCreated

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformTriggerEventReason("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `PlatformTriggerEventReasonCreated`              | CREATED                                          |
| `PlatformTriggerEventReasonUpdated`              | UPDATED                                          |
| `PlatformTriggerEventReasonDeleted`              | DELETED                                          |
| `PlatformTriggerEventReasonMeetsCondition`       | MEETS_CONDITION                                  |
| `PlatformTriggerEventReasonAssigned`             | ASSIGNED                                         |
| `PlatformTriggerEventReasonUnassigned`           | UNASSIGNED                                       |
| `PlatformTriggerEventReasonLabeled`              | LABELED                                          |
| `PlatformTriggerEventReasonUnlabeled`            | UNLABELED                                        |
| `PlatformTriggerEventReasonReviewRequested`      | REVIEW_REQUESTED                                 |
| `PlatformTriggerEventReasonReviewRequestRemoved` | REVIEW_REQUEST_REMOVED                           |
| `PlatformTriggerEventReasonReadyForReview`       | READY_FOR_REVIEW                                 |
| `PlatformTriggerEventReasonConvertedToDraft`     | CONVERTED_TO_DRAFT                               |
| `PlatformTriggerEventReasonWebhookUpdated`       | WEBHOOK_UPDATED                                  |