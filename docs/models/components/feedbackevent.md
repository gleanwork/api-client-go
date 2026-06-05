# FeedbackEvent

The action the user took within a Glean client with respect to the object referred to by the given `trackingToken`.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.FeedbackEventClick
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `FeedbackEventClick`                    | CLICK                                   |
| `FeedbackEventContainerClick`           | CONTAINER_CLICK                         |
| `FeedbackEventCopyLink`                 | COPY_LINK                               |
| `FeedbackEventCreate`                   | CREATE                                  |
| `FeedbackEventDismiss`                  | DISMISS                                 |
| `FeedbackEventDownvote`                 | DOWNVOTE                                |
| `FeedbackEventEmail`                    | EMAIL                                   |
| `FeedbackEventExecute`                  | EXECUTE                                 |
| `FeedbackEventFilter`                   | FILTER                                  |
| `FeedbackEventFirstToken`               | FIRST_TOKEN                             |
| `FeedbackEventFocusIn`                  | FOCUS_IN                                |
| `FeedbackEventLastToken`                | LAST_TOKEN                              |
| `FeedbackEventManualFeedback`           | MANUAL_FEEDBACK                         |
| `FeedbackEventManualFeedbackSideBySide` | MANUAL_FEEDBACK_SIDE_BY_SIDE            |
| `FeedbackEventFeedbackTimeSaved`        | FEEDBACK_TIME_SAVED                     |
| `FeedbackEventMarkAsRead`               | MARK_AS_READ                            |
| `FeedbackEventMessage`                  | MESSAGE                                 |
| `FeedbackEventMiddleClick`              | MIDDLE_CLICK                            |
| `FeedbackEventPageBlur`                 | PAGE_BLUR                               |
| `FeedbackEventPageFocus`                | PAGE_FOCUS                              |
| `FeedbackEventPageLeave`                | PAGE_LEAVE                              |
| `FeedbackEventPreview`                  | PREVIEW                                 |
| `FeedbackEventRelatedClick`             | RELATED_CLICK                           |
| `FeedbackEventRightClick`               | RIGHT_CLICK                             |
| `FeedbackEventSectionClick`             | SECTION_CLICK                           |
| `FeedbackEventSeen`                     | SEEN                                    |
| `FeedbackEventSelect`                   | SELECT                                  |
| `FeedbackEventShare`                    | SHARE                                   |
| `FeedbackEventShowMore`                 | SHOW_MORE                               |
| `FeedbackEventUpvote`                   | UPVOTE                                  |
| `FeedbackEventView`                     | VIEW                                    |
| `FeedbackEventVisible`                  | VISIBLE                                 |