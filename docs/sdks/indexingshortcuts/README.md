# IndexingShortcuts
(*Indexing.Shortcuts*)

## Overview

### Available Operations

* [BulkIndex](#bulkindex) - Bulk index external shortcuts

## BulkIndex

Replaces all the currently indexed shortcuts using paginated batch API calls. Note that this endpoint is used for indexing shortcuts not hosted by Glean. If you want to upload shortcuts that would be hosted by Glean, please use the `/uploadshortcuts` endpoint. For information on what you can do with Golinks, which are Glean-hosted shortcuts, please refer to [this](https://help.glean.com/en/articles/5628838-how-go-links-work) page.

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

    res, err := s.Indexing.Shortcuts.BulkIndex(ctx, components.BulkIndexShortcutsRequest{
        UploadID: "<id>",
        Shortcuts: []components.ExternalShortcut{
            components.ExternalShortcut{
                InputAlias: "<value>",
                DestinationURL: "https://only-juggernaut.com/",
                CreatedBy: "<value>",
                IntermediateURL: "https://descriptive-electronics.name",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.BulkIndexShortcutsRequest](../../models/components/bulkindexshortcutsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PostAPIIndexV1BulkindexshortcutsResponse](../../models/operations/postapiindexv1bulkindexshortcutsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |