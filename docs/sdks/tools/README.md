# Client.Tools

## Overview

### Available Operations

* [List](#list) - List available tools
* [Run](#run) - Execute the specified tool
* [RetrieveActionPackAuthStatus](#retrieveactionpackauthstatus) - Get end-user authentication status for an action pack.
* [AuthorizeActionPack](#authorizeactionpack) - Start the OAuth authorization flow for an action pack.
* [RetrieveToolServerAuthStatus](#retrievetoolserverauthstatus) - Get end-user authentication status for a tool server.
* [AuthorizeToolServer](#authorizetoolserver) - Start the OAuth authorization flow for a tool server.

## List

Returns a filtered set of available tools based on optional tool name parameters. If no filters are provided, all available tools are returned.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/rest/api/v1/tools/list" method="get" path="/rest/api/v1/tools/list" -->
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

    res, err := s.Client.Tools.List(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ToolsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `toolNames`                                              | []`string`                                               | :heavy_minus_sign:                                       | Optional array of tool names to filter by                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetRestAPIV1ToolsListResponse](../../models/operations/getrestapiv1toolslistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Run

Execute the specified tool with provided parameters

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/rest/api/v1/tools/call" method="post" path="/rest/api/v1/tools/call" -->
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

    res, err := s.Client.Tools.Run(ctx, components.ToolsCallRequest{
        Name: "<value>",
        Parameters: map[string]components.ToolsCallParameter{
            "key": components.ToolsCallParameter{
                Name: "<value>",
                Value: "<value>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ToolsCallResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.ToolsCallRequest](../../models/components/toolscallrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.PostRestAPIV1ToolsCallResponse](../../models/operations/postrestapiv1toolscallresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RetrieveActionPackAuthStatus

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

    res, err := s.Client.Tools.RetrieveActionPackAuthStatus(ctx, "<id>")
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

    res, err := s.Client.Tools.AuthorizeActionPack(ctx, "<id>", components.AuthorizeActionPackRequest{
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

## RetrieveToolServerAuthStatus

Returns display information and the calling user's current authentication status
for the specified tool server.


### Example Usage

<!-- UsageSnippet language="go" operationID="getToolServerAuthStatus" method="get" path="/rest/api/v1/tool-servers/{serverId}/auth" -->
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

    res, err := s.Client.Tools.RetrieveToolServerAuthStatus(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ToolServerAuthStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serverID`                                               | `string`                                                 | :heavy_check_mark:                                       | Unique identifier of the tool server.                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetToolServerAuthStatusResponse](../../models/operations/gettoolserverauthstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AuthorizeToolServer

Initiates the third-party OAuth flow for the specified tool server and returns the
authorization URL that the client should navigate the end user to. After the OAuth
callback completes, the user's browser is redirected back to `returnUrl` with query
parameters indicating the result.

`returnUrl` must match the tenant's configured return URL allowlist; otherwise the
request is rejected with 400.


### Example Usage

<!-- UsageSnippet language="go" operationID="authorizeToolServer" method="post" path="/rest/api/v1/tool-servers/{serverId}/auth" -->
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

    res, err := s.Client.Tools.AuthorizeToolServer(ctx, "<id>", components.AuthorizeToolServerRequest{
        ReturnURL: "https://lucky-disadvantage.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthorizeToolServerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `serverID`                                                                                     | `string`                                                                                       | :heavy_check_mark:                                                                             | Unique identifier of the tool server.                                                          |
| `authorizeToolServerRequest`                                                                   | [components.AuthorizeToolServerRequest](../../models/components/authorizetoolserverrequest.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.AuthorizeToolServerResponse](../../models/operations/authorizetoolserverresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |