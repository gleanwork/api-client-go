# Calendar
(*Client.Calendar*)

## Overview

### Available Operations

* [GetEvents](#getevents) - Read events

## GetEvents

Read detailed information about the given event ids.

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

    res, err := s.Client.Calendar.GetEvents(ctx, components.GetEventsRequest{
        Ids: []string{},
        AuthTokens: []components.AuthToken{
            components.AuthToken{
                AccessToken: "123abc",
                Datasource: "gmail",
                Scope: apiclientgo.String("email profile https://www.googleapis.com/auth/gmail.readonly"),
                TokenType: apiclientgo.String("Bearer"),
                AuthUser: apiclientgo.String("1"),
            },
        },
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `getEventsRequest`                                                                                                       | [components.GetEventsRequest](../../models/components/geteventsrequest.md)                                               | :heavy_check_mark:                                                                                                       | GetEvents request                                                                                                        |
| `xGleanActAs`                                                                                                            | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.GeteventsResponse](../../models/operations/geteventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |