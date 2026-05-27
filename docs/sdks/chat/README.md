# Chat

## Overview

### Available Operations

* [GetChatFile](#getchatfile) - Download a chat file

## GetChatFile

Download the raw content of a file generated or uploaded during a chat session (for example, an image produced by the assistant). Returns the file bytes with a Content-Type header matching the file's MIME type.


### Example Usage

<!-- UsageSnippet language="go" operationID="getChatFile" method="get" path="/rest/api/v1/chat-files/{fileId}" -->
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

    res, err := s.Chat.GetChatFile(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `fileID`                                                                                                                   | `string`                                                                                                                   | :heavy_check_mark:                                                                                                         | Identifier of the chat file to download.                                                                                   |
| `preview`                                                                                                                  | `*bool`                                                                                                                    | :heavy_minus_sign:                                                                                                         | When true and the file is a PDF, the response is served inline (Content-Disposition: inline) instead of as an attachment.<br/> |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.GetChatFileResponse](../../models/operations/getchatfileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |