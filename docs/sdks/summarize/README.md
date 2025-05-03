# Summarize
(*Client.Summarize*)

## Overview

### Available Operations

* [Generate](#generate) - Summarize documents

## Generate

Generate an AI summary of the requested documents.

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

    res, err := s.Client.Summarize.Generate(ctx, components.SummarizeRequest{
        DocumentSpecs: []components.DocumentSpecUnion{
            components.CreateDocumentSpecUnionDocumentSpec1(
                components.DocumentSpec1{},
            ),
            components.CreateDocumentSpecUnionDocumentSpec1(
                components.DocumentSpec1{},
            ),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SummarizeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.SummarizeRequest](../../models/components/summarizerequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.SummarizeResponse](../../models/operations/summarizeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |