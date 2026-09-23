# Agents

## Overview

### Available Operations

* [Search](#search) - Search agents
* [Get](#get) - Get agent
* [GetSchemas](#getschemas) - Get agent schemas
* [CreateRun](#createrun) - Create agent run
* [GetRun](#getrun) - Get agent run
* [CancelRun](#cancelrun) - Cancel an agent run
* [RespondToRun](#respondtorun) - Respond to agent run approvals

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

    res, err := s.Agents.Get(ctx, "{agent_id}")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformAgentGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `agentID`                                                | `string`                                                 | :heavy_check_mark:                                       | ID of the agent to retrieve.                             | {agent_id}                                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

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

    res, err := s.Agents.GetSchemas(ctx, "{agent_id}", apiclientgo.Pointer(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformAgentSchemasResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `agentID`                                                | `string`                                                 | :heavy_check_mark:                                       | ID of the agent whose schemas should be retrieved.       | {agent_id}                                               |
| `includeTools`                                           | `*bool`                                                  | :heavy_minus_sign:                                       | Whether to include tool metadata in the response.        |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PlatformAgentsGetSchemasResponse](../../models/operations/platformagentsgetschemasresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## CreateRun

Execute an agent run. By default, set `stream` to true to receive server-sent events; otherwise the response contains the final agent messages. Set `execution_mode` to `DURABLE` to persist a new run and return its initial snapshot with HTTP 201 without waiting for execution. Poll the agent-scoped GET run endpoint for progress. Durable execution continues after an HTTP disconnect, but is not automatically resumed after a QE restart or crash. An active turn becomes overdue more than 40 minutes after acceptance (a 30-minute execution timeout plus 10 minutes of grace). The next GET of the run marks the overdue turn FAILED without replay; there is no periodic sweep. Without a GET, the stored run can remain RUNNING. Failure does not prove that external tool work has stopped. Paused runs are not expired; an accepted approval continuation starts a fresh deadline. Each POST creates a new run; retrying a POST can create another execution. Submit pending approval decisions through the run responses endpoint, and cancellation can be requested through the run cancellations endpoint. A run tracks one workflow execution; automatic background-subagent wake turns are separate executions, not continuations tracked by this run ID.


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

    res, err := s.Agents.CreateRun(ctx, "{agent_id}", components.PlatformAgentRunCreateRequest{
        Messages: []components.PlatformMessageInput{
            components.PlatformMessageInput{
                Role: components.PlatformMessageRoleUser,
                Content: []components.PlatformMessageTextBlockInput{
                    components.PlatformMessageTextBlockInput{
                        Text: "What is our parental leave policy?",
                        Type: components.PlatformContentTypeText,
                    },
                },
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

| Parameter                                                                                                               | Type                                                                                                                    | Required                                                                                                                | Description                                                                                                             | Example                                                                                                                 |
| ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                   | :heavy_check_mark:                                                                                                      | The context to use for the request.                                                                                     |                                                                                                                         |
| `agentID`                                                                                                               | `string`                                                                                                                | :heavy_check_mark:                                                                                                      | ID of the agent to run.                                                                                                 | {agent_id}                                                                                                              |
| `platformAgentRunCreateRequest`                                                                                         | [components.PlatformAgentRunCreateRequest](../../models/components/platformagentruncreaterequest.md)                    | :heavy_check_mark:                                                                                                      | N/A                                                                                                                     | {<br/>"messages": [<br/>{<br/>"role": "USER",<br/>"content": [<br/>{<br/>"text": "What is our parental leave policy?",<br/>"type": "text"<br/>}<br/>]<br/>}<br/>]<br/>} |
| `opts`                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                | :heavy_minus_sign:                                                                                                      | The options for this request.                                                                                           |                                                                                                                         |

### Response

**[*operations.PlatformAgentsCreateRunResponse](../../models/operations/platformagentscreaterunresponse.md), error**

### Errors

| Error Type                                           | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| apierrors.PlatformUnauthorizedAgentToolsProblemError | 422                                                  | application/problem+json                             |
| apierrors.PlatformProblemDetailError                 | 400, 401, 403, 404, 408, 409, 413, 429               | application/problem+json                             |
| apierrors.PlatformProblemDetailError                 | 500, 503                                             | application/problem+json                             |
| apierrors.APIError                                   | 4XX, 5XX                                             | \*/\*                                                |

## GetRun

Retrieve a persisted workflow execution owned by the authenticated user. The run must belong to the specified agent, and the user must still have access to that agent. Unknown runs, runs owned by another user, and mismatched agent/run identifiers return 404. Requires the agents.run scope. Executions without a persisted workflow record are not available through this endpoint.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-agents-get-run" method="get" path="/api/agents/{agent_id}/runs/{run_id}" -->
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

    res, err := s.Agents.GetRun(ctx, "{agent_id}", "{run_id}")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformAgentRunResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `agentID`                                                | `string`                                                 | :heavy_check_mark:                                       | ID of the agent that owns the run.                       | {agent_id}                                               |
| `runID`                                                  | `string`                                                 | :heavy_check_mark:                                       | ID of the durable run to retrieve.                       | {run_id}                                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PlatformAgentsGetRunResponse](../../models/operations/platformagentsgetrunresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## CancelRun

Request cooperative cancellation of the durable agent run identified by `run_id` in the JSON body. Requires ownership, current agent access, and the agents.run scope. Sending a cancellation signal does not itself change an active run from RUNNING; poll GET run for the final state. Paused runs become CANCELLED without resuming execution. Repeated requests and requests for terminal runs return the current snapshot. Completion may win a race with cancellation. Completed tool side effects cannot be undone, and external work may continue if a tool does not support cancellation. Cancellation targets this run, not separate background-subagent executions. An active run without a cancellation registration returns 409. Cancellation signaling requires Redis. An interrupted active run can instead become FAILED through deadline cleanup; this does not verify that external tool work has stopped.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-agents-cancel-run" method="post" path="/api/agents/{agent_id}/cancellations" -->
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

    res, err := s.Agents.CancelRun(ctx, "{agent_id}", components.PlatformAgentRunCancellationRequest{
        RunID: "{run_id}",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformAgentRunResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      | Example                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |                                                                                                                  |
| `agentID`                                                                                                        | `string`                                                                                                         | :heavy_check_mark:                                                                                               | ID of the agent that owns the run.                                                                               | {agent_id}                                                                                                       |
| `platformAgentRunCancellationRequest`                                                                            | [components.PlatformAgentRunCancellationRequest](../../models/components/platformagentruncancellationrequest.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              | {<br/>"run_id": "{run_id}"<br/>}                                                                                 |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |                                                                                                                  |

### Response

**[*operations.PlatformAgentsCancelRunResponse](../../models/operations/platformagentscancelrunresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.PlatformProblemDetailError   | 400, 401, 403, 404, 408, 409, 413, 429 | application/problem+json               |
| apierrors.PlatformProblemDetailError   | 500, 503                               | application/problem+json               |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RespondToRun

Submit decisions for every pending tool approval in the paused run's current batch. The run is identified by `run_id` in the JSON body. Decisions apply only to the stored invocations and arguments; argument edits, authentication responses, and session-wide grants are not supported. The caller must own the run, still have agent access, and have the agents.run scope. Acceptance persists the decisions before resuming the same run and chat session. Identical accepted decisions return the current snapshot without another continuation. Conflicting, stale, incomplete, or non-pending decisions return 409. Cancellation registration failure returns 503 without accepting the decisions; retry the same approval batch. This retry guarantee does not cover an indeterminate database commit outcome. Go workflow approval resumes currently support one tool invocation and one approval response. Unsupported multi-tool or multi-decision Go resumes fail without executing tools. The resumed action must resolve to the tool identified by the stored approval request and paused checkpoint. Missing or inconsistent identity fails without executing tools. Execution continues after HTTP disconnects, but is not automatically resumed after a QE crash. Each accepted continuation starts a fresh 30-minute execution timeout and 40-minute cleanup deadline. Identical retries do not extend that deadline. Waiting for approval does not expire a run. The next GET marks an overdue active turn FAILED without replaying execution.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-agents-create-run-responses" method="post" path="/api/agents/{agent_id}/responses" -->
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

    res, err := s.Agents.RespondToRun(ctx, "{agent_id}", components.PlatformAgentRunResponsesRequest{
        RunID: "{run_id}",
        Responses: []components.PlatformAgentRunApprovalDecision{
            components.PlatformAgentRunApprovalDecision{
                InteractionID: "{interaction_id}",
                Decision: components.DecisionApprove,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformAgentRunResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |                                                                                                            |
| `agentID`                                                                                                  | `string`                                                                                                   | :heavy_check_mark:                                                                                         | ID of the agent that owns the run.                                                                         | {agent_id}                                                                                                 |
| `platformAgentRunResponsesRequest`                                                                         | [components.PlatformAgentRunResponsesRequest](../../models/components/platformagentrunresponsesrequest.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        | {<br/>"run_id": "{run_id}",<br/>"responses": [<br/>{<br/>"interaction_id": "{interaction_id}",<br/>"decision": "APPROVE"<br/>}<br/>]<br/>} |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |                                                                                                            |

### Response

**[*operations.PlatformAgentsCreateRunResponsesResponse](../../models/operations/platformagentscreaterunresponsesresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.PlatformProblemDetailError   | 400, 401, 403, 404, 408, 409, 413, 429 | application/problem+json               |
| apierrors.PlatformProblemDetailError   | 500, 503                               | application/problem+json               |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |