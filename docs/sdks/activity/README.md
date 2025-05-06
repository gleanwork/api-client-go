# Activity
(*Client.Activity*)

## Overview

### Available Operations

* [Report](#report) - Report document activity
* [Feedback](#feedback) - Report client activity

## Report

Report user activity that occurs on indexed documents such as viewing or editing. This signal improves search quality.

### Example Usage

```go
package main

import(
	"context"
	"os"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/api-client-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
    )

    res, err := s.Client.Activity.Report(ctx, components.Activity{
        Events: []components.ActivityEvent{
            components.ActivityEvent{
                Action: components.ActivityEventActionHistoricalView,
                Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
                URL: "https://example.com/",
            },
            components.ActivityEvent{
                Action: components.ActivityEventActionSearch,
                Params: &components.ActivityEventParams{
                    Query: apiclientgo.String("query"),
                },
                Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
                URL: "https://example.com/search?q=query",
            },
            components.ActivityEvent{
                Action: components.ActivityEventActionView,
                Params: &components.ActivityEventParams{
                    Duration: apiclientgo.Int64(20),
                    Referrer: apiclientgo.String("https://example.com/document"),
                },
                Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
                URL: "https://example.com/",
            },
        },
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

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `request`                                                  | [components.Activity](../../models/components/activity.md) | :heavy_check_mark:                                         | The request object to use for the request.                 |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |

### Response

**[*operations.ActivityResponse](../../models/operations/activityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Feedback

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

    res, err := s.Client.Activity.Feedback(ctx, nil, &components.Feedback{
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