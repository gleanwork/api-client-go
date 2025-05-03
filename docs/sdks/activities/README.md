# Activities
(*Client.Activities*)

## Overview

### Available Operations

* [ReportActivity](#reportactivity) - Report client activity

## ReportActivity

Report events that happen to results within a Glean client UI, such as search result views and clicks.  This signal improves search quality.

### Example Usage

```go
package main

import(
	"context"
	"os"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
    )

    res, err := s.Client.Activities.ReportActivity(ctx, nil, &components.Feedback{
        TrackingTokens: []string{
            "trackingTokens",
        },
        Event: components.EventView,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |                                                                  |
| `feedbackQueryParameter`                                         | **string*                                                        | :heavy_minus_sign:                                               | A URL encoded versions of Feedback. This is useful for requests. |                                                                  |
| `feedback1`                                                      | [*components.Feedback](../../models/components/feedback.md)      | :heavy_minus_sign:                                               | N/A                                                              | {<br/>"trackingTokens": [<br/>"trackingTokens"<br/>],<br/>"event": "VIEW"<br/>} |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |                                                                  |

### Response

**[*operations.FeedbackResponse](../../models/operations/feedbackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |