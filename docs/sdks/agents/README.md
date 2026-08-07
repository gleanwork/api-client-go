# Agents

## Overview

### Available Operations

* [Search](#search) - Search agents
* [Get](#get) - Get agent
* [GetSchemas](#getschemas) - Get agent schemas
* [CreateRun](#createrun) - Create agent run

## Search

Search agents available to the authenticated user by agent name.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-agents-search" method="post" path="/api/agents/search" -->
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

    res, err := s.Agents.Search(ctx, components.PlatformAgentsSearchRequest{
        Name: apiclientgo.Pointer("HR Policy Agent"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformAgentsSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.PlatformAgentsSearchRequest](../../models/components/platformagentssearchrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PlatformAgentsSearchResponse](../../models/operations/platformagentssearchresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 413, 429    | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## Get

Retrieve details for an agent available to the authenticated user.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-agents-get" method="get" path="/api/agents/{agent_id}" -->
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

    res, err := s.Agents.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformAgentGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `agentID`                                                | `string`                                                 | :heavy_check_mark:                                       | ID of the agent to retrieve.                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PlatformAgentsGetResponse](../../models/operations/platformagentsgetresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## GetSchemas

Retrieve an agent's input and output JSON schemas.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-agents-get-schemas" method="get" path="/api/agents/{agent_id}/schemas" -->
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

    res, err := s.Agents.GetSchemas(ctx, "<id>", apiclientgo.Pointer(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformAgentSchemasResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `agentID`                                                | `string`                                                 | :heavy_check_mark:                                       | ID of the agent whose schemas should be retrieved.       |
| `includeTools`                                           | `*bool`                                                  | :heavy_minus_sign:                                       | Whether to include tool metadata in the response.        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PlatformAgentsGetSchemasResponse](../../models/operations/platformagentsgetschemasresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## CreateRun

Execute an agent run. Set `stream` to true to receive server-sent events; otherwise the response contains the final agent messages.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-agents-create-run" method="post" path="/api/agents/{agent_id}/runs" -->
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

    res, err := s.Agents.CreateRun(ctx, "<id>", components.PlatformAgentRunCreateRequest{
        Messages: []components.PlatformMessage{
            components.PlatformMessage{
                Role: components.PlatformMessageRoleUser,
                Content: []components.PlatformMessageTextBlock{},
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformAgentRunWaitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `agentID`                                                                                            | `string`                                                                                             | :heavy_check_mark:                                                                                   | ID of the agent to run.                                                                              |
| `platformAgentRunCreateRequest`                                                                      | [components.PlatformAgentRunCreateRequest](../../models/components/platformagentruncreaterequest.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PlatformAgentsCreateRunResponse](../../models/operations/platformagentscreaterunresponse.md), error**

### Errors

| Error Type                                           | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| apierrors.PlatformUnauthorizedAgentToolsProblemError | 422                                                  | application/problem+json                             |
| apierrors.PlatformProblemDetailError                 | 400, 401, 403, 404, 408, 409, 413, 429               | application/problem+json                             |
| apierrors.PlatformProblemDetailError                 | 500, 503                                             | application/problem+json                             |
| apierrors.APIError                                   | 4XX, 5XX                                             | \*/\*                                                |