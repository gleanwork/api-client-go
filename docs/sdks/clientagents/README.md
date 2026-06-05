# Client.Agents

## Overview

### Available Operations

* [Retrieve](#retrieve) - Retrieve an agent
* [RetrieveSchemas](#retrieveschemas) - List an agent's schemas
* [List](#list) - Search agents
* [RunStream](#runstream) - Create an agent run and stream the response
* [Run](#run) - Create an agent run and wait for the response

## Retrieve

Returns details of an [agent](https://developers.glean.com/agents/agents-api) created in the Agent Builder.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAgent" method="get" path="/rest/api/v1/agents/{agent_id}" -->
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

    res, err := s.Client.Agents.Retrieve(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Agent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                           | Type                                                                                                                                                                                                | Required                                                                                                                                                                                            | Description                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                 |
| `agentID`                                                                                                                                                                                           | `string`                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                  | The ID of the agent.                                                                                                                                                                                |
| `locale`                                                                                                                                                                                            | `*string`                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                  | The client's preferred locale in rfc5646 format (e.g. `en`, `ja`, `pt-BR`). If omitted, the `Accept-Language` will be used. If not present or not supported, defaults to the closest match or `en`. |
| `timezoneOffset`                                                                                                                                                                                    | `*int64`                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC.                                                                                          |
| `opts`                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | The options for this request.                                                                                                                                                                       |

### Response

**[*operations.GetAgentResponse](../../models/operations/getagentresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## RetrieveSchemas

Return [agent](https://developers.glean.com/agents/agents-api)'s input and output schemas. You can use these schemas to detect changes to an agent's input or output structure.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAgentSchemas" method="get" path="/rest/api/v1/agents/{agent_id}/schemas" -->
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

    res, err := s.Client.Agents.RetrieveSchemas(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentSchemas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                           | Type                                                                                                                                                                                                | Required                                                                                                                                                                                            | Description                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                 |
| `agentID`                                                                                                                                                                                           | `string`                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                  | The ID of the agent.                                                                                                                                                                                |
| `locale`                                                                                                                                                                                            | `*string`                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                  | The client's preferred locale in rfc5646 format (e.g. `en`, `ja`, `pt-BR`). If omitted, the `Accept-Language` will be used. If not present or not supported, defaults to the closest match or `en`. |
| `timezoneOffset`                                                                                                                                                                                    | `*int64`                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC.                                                                                          |
| `opts`                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | The options for this request.                                                                                                                                                                       |

### Response

**[*operations.GetAgentSchemasResponse](../../models/operations/getagentschemasresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## List

Search for [agents](https://developers.glean.com/agents/agents-api) by agent name.

### Example Usage

<!-- UsageSnippet language="go" operationID="searchAgents" method="post" path="/rest/api/v1/agents/search" -->
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

    res, err := s.Client.Agents.List(ctx, components.SearchAgentsRequest{
        Name: apiclientgo.Pointer("HR Policy Agent"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAgentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.SearchAgentsRequest](../../models/components/searchagentsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.SearchAgentsResponse](../../models/operations/searchagentsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## RunStream

Executes an [agent](https://developers.glean.com/agents/agents-api) run and returns the result as a stream of server-sent events (SSE). **Note**: If the agent uses an input form trigger, all form fields (including optional fields) must be included in the `input` object.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAndStreamRun" method="post" path="/rest/api/v1/agents/runs/stream" -->
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

    res, err := s.Client.Agents.RunStream(ctx, components.AgentRunCreate{
        AgentID: "<id>",
        Messages: []components.Message{
            components.Message{
                Role: apiclientgo.Pointer("USER"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.AgentRunCreate](../../models/components/agentruncreate.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.CreateAndStreamRunResponse](../../models/operations/createandstreamrunresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 409, 422           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Run

Executes an [agent](https://developers.glean.com/agents/agents-api) run and returns the final response. **Note**: If the agent uses an input form trigger, all form fields (including optional fields) must be included in the `input` object.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAndWaitRun" method="post" path="/rest/api/v1/agents/runs/wait" -->
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

    res, err := s.Client.Agents.Run(ctx, components.AgentRunCreate{
        AgentID: "<id>",
        Messages: []components.Message{
            components.Message{
                Role: apiclientgo.Pointer("USER"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentRunWaitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.AgentRunCreate](../../models/components/agentruncreate.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.CreateAndWaitRunResponse](../../models/operations/createandwaitrunresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |