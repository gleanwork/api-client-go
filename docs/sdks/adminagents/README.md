# Admin.Agents

## Overview

### Available Operations

* [List](#list) - List all agents

## List

List every agent in the workspace, ordered by name and then by agent_id. Each agent has the same shape that Get agent returns. Unlike agent search, the list does not depend on who an agent is shared with, so it includes private and disabled agents. Deleted agents are left out. Use the agent_id with other Platform APIs, for example as a usage limit target. The caller must be able to read all agents, administer agents, or read workspace billing.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-admin-agents-list" method="get" path="/api/admin/agents" -->
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

    res, err := s.Admin.Agents.List(ctx, apiclientgo.Pointer[int64](50), nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformListAdminAgentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                               | Type                                                                                                                                                                                                                                                                                                                                                    | Required                                                                                                                                                                                                                                                                                                                                                | Description                                                                                                                                                                                                                                                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                      | The context to use for the request.                                                                                                                                                                                                                                                                                                                     |
| `pageSize`                                                                                                                                                                                                                                                                                                                                              | `*int64`                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                      | Maximum number of agents to return. Defaults to 50. Maximum is 100.                                                                                                                                                                                                                                                                                     |
| `cursor`                                                                                                                                                                                                                                                                                                                                                | `*string`                                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                      | Opaque pagination cursor from a previous response. Send the same filters that the previous request used.<br/>                                                                                                                                                                                                                                           |
| `filters`                                                                                                                                                                                                                                                                                                                                               | [][components.PlatformAdminAgentFilter](../../models/components/platformadminagentfilter.md)                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                      | JSON-encoded filters that choose which agents are returned. The supported fields are name and description, with the operator EQUALS. Multiple values OR within a filter. Multiple filters AND together, and each field can appear in only one filter. For example, to return agents whose name contains "HR", send [{"field":"name","values":["HR"]}].<br/> |
| `opts`                                                                                                                                                                                                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                      | The options for this request.                                                                                                                                                                                                                                                                                                                           |

### Response

**[*operations.PlatformAdminAgentsListResponse](../../models/operations/platformadminagentslistresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |