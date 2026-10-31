# Admin.Oauth.Clients

## Overview

### Available Operations

* [List](#list) - List OAuth clients
* [Get](#get) - Get an OAuth client

## List

List the static OAuth clients that admins created for this workspace, ordered by display_name and then by client_id. Dynamically registered clients and Glean's own clients, such as Glean Desktop, are not listed. Disabled clients are left out unless an is_enabled filter includes the value "false". Use filters to find clients by name, client_id, usage_client_id, LLM Gateway access, or supported app. To set a usage limit for a client, send its ai_gateway.usage_client_id, not its client_id, to the usage APIs. The caller needs admin read or billing read permission.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-admin-oauth-clients-list" method="get" path="/api/admin/oauth/clients" -->
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

    res, err := s.Admin.Oauth.Clients.List(ctx, apiclientgo.Pointer[int64](50), nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformListOAuthClientsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `pageSize`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | `*int64`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Maximum number of clients to return. Defaults to 50. Maximum is 100.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| `cursor`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | `*string`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Opaque pagination cursor from a previous response. Send the same filters that the previous request used.<br/>                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| `filters`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [][components.PlatformOAuthClientFilter](../../models/components/platformoauthclientfilter.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | JSON-encoded filters that choose which clients are returned. The supported fields are is_enabled, name, client_id, usage_client_id, has_ai_gateway, and supported_app, with the operator EQUALS. Multiple values OR within a filter. Multiple filters AND together. Each field can appear in only one filter. Without an is_enabled filter, only enabled clients are returned. To return enabled and disabled clients, send [{"field":"is_enabled","values":["true","false"]}]. To find the Claude Code clients, send [{"field":"supported_app","values":["CLAUDE_CODE"]}].<br/> |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |

### Response

**[*operations.PlatformAdminOauthClientsListResponse](../../models/operations/platformadminoauthclientslistresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## Get

Get one static OAuth client that an admin created for this workspace. A client that an admin disabled is returned with is_enabled false. A deleted client, a dynamically registered client, one of Glean's own clients, such as Glean Desktop, and an unknown client_id all return 404. The path takes the OAuth client_id, not ai_gateway.usage_client_id. The caller needs admin read or billing read permission.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-admin-oauth-clients-get" method="get" path="/api/admin/oauth/clients/{client_id}" -->
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

    res, err := s.Admin.Oauth.Clients.Get(ctx, "{client_id}")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformGetOAuthClientResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `clientID`                                               | `string`                                                 | :heavy_check_mark:                                       | The OAuth client_id of the client.                       | {client_id}                                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PlatformAdminOauthClientsGetResponse](../../models/operations/platformadminoauthclientsgetresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |