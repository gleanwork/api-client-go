# Chat

## Overview

### Available Operations

* [Create](#create) - Create a chat response
* [Update](#update) - Update a chat
* [Resume](#resume) - Resume a chat stream
* [Cancel](#cancel) - Cancel a chat
* [CreateStream](#createstream) - Create a streaming chat response

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
            "Summarize our parental leave policy.",
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

## Update

Update the name of a chat. Only the chat owner can update it.

### Example Usage

<!-- UsageSnippet language="go" operationID="platform-chat-update" method="patch" path="/api/chats/{chat_id}" -->
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

    res, err := s.Chat.Update(ctx, "{chat_id}", components.PlatformChatUpdateRequest{
        Name: "Quarterly planning notes",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformChatUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `chatID`                                                                                     | `string`                                                                                     | :heavy_check_mark:                                                                           | The ID of the chat to update.                                                                | {chat_id}                                                                                    |
| `platformChatUpdateRequest`                                                                  | [components.PlatformChatUpdateRequest](../../models/components/platformchatupdaterequest.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.PlatformChatUpdateResponse](../../models/operations/platformchatupdateresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 405, 413, 429    | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500                                  | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## Resume

Reconnect to a chat that is still running, or replay one that just finished. The response is server-sent events in the same format as the streaming response of the create endpoint. Pass the `sequence_id` of the last event you received as `last_seen_sequence_id`, or omit it to start from the beginning of the stream. The final event is `RESPONSE_COMPLETED` or `RESPONSE_FAILED`, and its text covers only the events sent on this connection. Streams are kept for a limited time, so a chat that finished long ago returns 404.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-chat-resume" method="get" path="/api/chats/{chat_id}/resume" -->
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

    res, err := s.Chat.Resume(ctx, "{chat_id}", apiclientgo.Pointer[int64](0))
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

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      | Example                                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |                                                                                                                                                  |
| `chatID`                                                                                                                                         | `string`                                                                                                                                         | :heavy_check_mark:                                                                                                                               | The ID of the chat to resume.                                                                                                                    | {chat_id}                                                                                                                                        |
| `lastSeenSequenceID`                                                                                                                             | `*int64`                                                                                                                                         | :heavy_minus_sign:                                                                                                                               | The `sequence_id` of the last event already received. The stream continues with the next event. Defaults to 0, which starts from the beginning.<br/> | 0                                                                                                                                                |
| `opts`                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |                                                                                                                                                  |

### Response

**[*operations.PlatformChatResumeResponse](../../models/operations/platformchatresumeresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 429              | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## Cancel

Stop a running chat. Only users who can edit the chat can cancel it. Cancelling a chat that is not running succeeds without effect.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-chat-cancel" method="post" path="/api/chats/{chat_id}/cancellations" -->
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

    res, err := s.Chat.Cancel(ctx, "{chat_id}")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformChatCancelResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `chatID`                                                 | `string`                                                 | :heavy_check_mark:                                       | The ID of the chat to cancel.                            | {chat_id}                                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PlatformChatCancelResponse](../../models/operations/platformchatcancelresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 429              | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## CreateStream

SDK-only logical operation. HTTP clients must call the base path; the URL fragment is not sent. Run an assistant turn as server-sent events. HTTP clients request this mode by setting `stream` to true in the JSON body.


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
            "Summarize our parental leave policy.",
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