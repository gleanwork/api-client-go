# Chat

## Overview

### Available Operations

* [Create](#create) - Create a chat response
* [CreateStream](#createstream) - SDK-only logical operation. HTTP clients must call the base path; the URL fragment is not sent. Create a chat response

## Create

Run an assistant turn. The default response is JSON. HTTP clients request server-sent events by setting `stream` to true in the JSON body. An `Accept: text/event-stream` header does not replace `stream`.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-chat-create" method="post" path="/api/chat" -->
```go
package main

import(
	"context"
	"os"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Chat.Create(ctx, operations.PlatformChatCreateRequest{
        Input: operations.CreatePlatformChatCreateInputStr(
            "What is our parental leave policy?",
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformChatCompletedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PlatformChatCreateRequest](../../models/operations/platformchatcreaterequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PlatformChatCreateResponse](../../models/operations/platformchatcreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.PlatformProblemDetailError   | 400, 401, 403, 404, 408, 413, 422, 429 | application/problem+json               |
| apierrors.PlatformProblemDetailError   | 500, 503                               | application/problem+json               |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateStream

SDK-only logical operation. HTTP clients must call the base path; the URL fragment is not sent. Run an assistant turn. The default response is JSON. HTTP clients request server-sent events by setting `stream` to true in the JSON body. An `Accept: text/event-stream` header does not replace `stream`.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-chat-create-stream" method="post" path="/api/chat#stream" -->
```go
package main

import(
	"context"
	"os"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Chat.CreateStream(ctx, operations.PlatformChatCreateStreamRequest{
        Input: operations.CreatePlatformChatCreateStreamInputStr(
            "What is our parental leave policy?",
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformChatStreamEventServerSentEvent != nil {
        defer res.PlatformChatStreamEventServerSentEvent.Close()

        for res.PlatformChatStreamEventServerSentEvent.Next() {
            event := res.PlatformChatStreamEventServerSentEvent.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PlatformChatCreateStreamRequest](../../models/operations/platformchatcreatestreamrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PlatformChatCreateStreamResponse](../../models/operations/platformchatcreatestreamresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.PlatformProblemDetailError   | 400, 401, 403, 404, 408, 413, 422, 429 | application/problem+json               |
| apierrors.PlatformProblemDetailError   | 500, 503                               | application/problem+json               |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |