# Search

## Overview

### Available Operations

* [Query](#query) - Search

## Query

Execute a search query and retrieve ranked results. This is the data retrieval variant of the search API and returns only results and pagination state.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-search" method="post" path="/api/search" -->
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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Search.Query(ctx, components.PlatformSearchRequest{
        Query: "quarterly planning 2026",
        Datasources: []string{
            "confluence",
            "google_drive",
        },
        DatasourceInstances: []string{
            "slack_acme",
            "slack_eu",
        },
        Filters: []components.PlatformFilter{
            components.PlatformFilter{
                Field: "type",
                Values: []string{
                    "spreadsheet",
                    "presentation",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.PlatformSearchRequest](../../models/components/platformsearchrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PlatformSearchResponse](../../models/operations/platformsearchresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |