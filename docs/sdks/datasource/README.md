# Indexing.Datasource

## Overview

### Available Operations

* [Status](#status) - Get datasource status


## Status

Gather information about the datasource's overall status.

Tip: Refer to the [Troubleshooting tutorial](https://developers.glean.com/indexing/debugging/datasource-config) for more information.


### Example Usage

<!-- UsageSnippet language="go" operationID="post_/api/index/v1/debug/{datasource}/status" method="post" path="/api/index/v1/debug/{datasource}/status" -->
```go
package main

import(
	"context"
	"os"
	apiclientgo "github.com/gleanwork/api-client-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Indexing.Datasource.Status(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DebugDatasourceStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `datasource`                                             | `string`                                                 | :heavy_check_mark:                                       | The datasource to get debug status for.                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PostAPIIndexV1DebugDatasourceStatusResponse](../../models/operations/postapiindexv1debugdatasourcestatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |