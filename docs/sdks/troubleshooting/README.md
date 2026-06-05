# Troubleshooting

## Overview

### Available Operations

* [PostAPIIndexV1DebugDatasourceDocumentEvents](#postapiindexv1debugdatasourcedocumentevents) - Beta: Get document lifecycle events


## PostAPIIndexV1DebugDatasourceDocumentEvents

Retrieves lifecycle events for a specific document including upload time, index times and deletions. Rate limited to 1 request per minute per datasource. Currently in beta, might undergo breaking changes without prior notice.


### Example Usage

<!-- UsageSnippet language="go" operationID="post_/api/index/v1/debug/{datasource}/document/events" method="post" path="/api/index/v1/debug/{datasource}/document/events" -->
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

    res, err := s.Troubleshooting.PostAPIIndexV1DebugDatasourceDocumentEvents(ctx, "<value>", components.DebugDocumentLifecycleRequest{
        ObjectType: "Article",
        DocID: "art123",
        StartDate: apiclientgo.Pointer("2025-05-01"),
        MaxEvents: apiclientgo.Pointer[int64](50),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DebugDocumentLifecycleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `datasource`                                                                                         | `string`                                                                                             | :heavy_check_mark:                                                                                   | The datasource to which the document belongs                                                         |
| `debugDocumentLifecycleRequest`                                                                      | [components.DebugDocumentLifecycleRequest](../../models/components/debugdocumentlifecyclerequest.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PostAPIIndexV1DebugDatasourceDocumentEventsResponse](../../models/operations/postapiindexv1debugdatasourcedocumenteventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |