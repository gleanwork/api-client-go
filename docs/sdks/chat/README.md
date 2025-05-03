# Chat
(*Chat*)

## Overview

### Available Operations

* [ChatStream](#chatstream) - Chat

## ChatStream

Have a conversation with Glean AI.

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

    res, err := s.Chat.ChatStream(ctx, components.ChatRequest{
        Messages: []components.ChatMessage{
            components.ChatMessage{
                Fragments: []components.ChatMessageFragment{
                    components.ChatMessageFragment{
                        Text: apiclientgo.String("What are the company holidays this year?"),
                    },
                },
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ChatRequestStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `chatRequest`                                                                                              | [components.ChatRequest](../../models/components/chatrequest.md)                                           | :heavy_check_mark:                                                                                         | Includes chat history for Glean AI to respond to.                                                          |
| `timezoneOffset`                                                                                           | **int64*                                                                                                   | :heavy_minus_sign:                                                                                         | The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC. |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ChatStreamResponse](../../models/operations/chatstreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |