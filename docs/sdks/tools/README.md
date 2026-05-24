# Tools

## Overview

### Available Operations

* [GetActionPackAuthStatus](#getactionpackauthstatus) - Get end-user authentication status for an action pack.
* [AuthorizeActionPack](#authorizeactionpack) - Start the OAuth authorization flow for an action pack.

## GetActionPackAuthStatus

Reports whether the calling user is already authenticated against the third-party
tool backing the specified action pack. Intended for headless / server-driven clients
that render an "Authorize" prompt when the user has not yet consented to the tool.


### Example Usage

<!-- UsageSnippet language="go" operationID="getActionPackAuthStatus" method="get" path="/rest/api/v1/actions/actionpack/{actionPackId}/auth" -->
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

    res, err := s.Tools.GetActionPackAuthStatus(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ActionPackAuthStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `actionPackID`                                           | `string`                                                 | :heavy_check_mark:                                       | ID of the action pack to query or authorize.             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetActionPackAuthStatusResponse](../../models/operations/getactionpackauthstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AuthorizeActionPack

Starts the third-party OAuth flow for the specified action pack and returns the
redirect URL that the client should navigate the end user to. After the OAuth
callback completes, the user's browser is redirected back to `returnUrl` with a
status query parameter (`?glean_action_auth=success|error&actionPackId=...`).

`returnUrl` must match the tenant's configured return URL allowlist; otherwise the
request is rejected with 400.


### Example Usage

<!-- UsageSnippet language="go" operationID="authorizeActionPack" method="post" path="/rest/api/v1/actions/actionpack/{actionPackId}/auth" -->
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

    res, err := s.Tools.AuthorizeActionPack(ctx, "<id>", components.AuthorizeActionPackRequest{
        ReturnURL: "https://merry-allocation.org/",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthorizeActionPackResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `actionPackID`                                                                                 | `string`                                                                                       | :heavy_check_mark:                                                                             | ID of the action pack to query or authorize.                                                   |
| `authorizeActionPackRequest`                                                                   | [components.AuthorizeActionPackRequest](../../models/components/authorizeactionpackrequest.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.AuthorizeActionPackResponse](../../models/operations/authorizeactionpackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |