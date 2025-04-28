# Insights
(*Client.Insights*)

## Overview

### Available Operations

* [Get](#get) - Read insights

## Get

Reads the aggregate information for each user, query, and content.

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

    res, err := s.Client.Insights.Get(ctx, components.InsightsRequest{
        Categories: []components.InsightsRequestCategory{
            components.InsightsRequestCategoryContent,
            components.InsightsRequestCategoryContent,
        },
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.InsightsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `insightsRequest`                                                                                                        | [components.InsightsRequest](../../models/components/insightsrequest.md)                                                 | :heavy_check_mark:                                                                                                       | Includes request params for insights dashboard data.                                                                     |
| `xGleanActAs`                                                                                                            | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.InsightsResponse](../../models/operations/insightsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |