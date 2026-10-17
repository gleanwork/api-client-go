# Admin.Idpgroups

## Overview

### Available Operations

* [List](#list) - List IdP groups
* [Get](#get) - Get an IdP group

## List

List the identity provider (IdP) groups that Glean loads from the directory, ordered by display_name and then by idp_group_id. Use idp_group_id to name a group in other Platform APIs. An IdP group is active when Glean loaded it from the directory. An IdP group is inactive when only a usage limit names it, for example after the group was removed from the identity provider. Inactive IdP groups are left out unless an is_active filter includes the value "false". Usage limits apply only to IdP groups from the datasource instance that the usage settings select, so use the provider_instance filter to list only those groups. Internal O365 ownership groups and SharePoint groups are left out. The caller must be able to read the workspace configuration or the workspace billing.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-admin-idp-groups-list" method="get" path="/api/admin/idp-groups" -->
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

    res, err := s.Admin.Idpgroups.List(ctx, apiclientgo.Pointer[int64](50), nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformListIdpGroupsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| `pageSize`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | `*int64`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Maximum number of IdP groups to return. Defaults to 50. Maximum is 100.                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| `cursor`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | `*string`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Opaque pagination cursor from a previous response. Send the same filters that the previous request used.<br/>                                                                                                                                                                                                                                                                                                                                                                                                                             |
| `filters`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [][components.PlatformIdpGroupFilter](../../models/components/platformidpgroupfilter.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | JSON-encoded filters that choose which IdP groups are returned. The supported fields are is_active, with the values "true" and "false", and provider_instance, with datasource instance names such as "okta". provider_instance values match without case. The only supported operator is EQUALS. Multiple values OR within a filter. Multiple filters AND together. Without an is_active filter, only active IdP groups are returned. To return active and inactive IdP groups, send [{"field":"is_active","values":["true","false"]}].<br/> |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |

### Response

**[*operations.PlatformAdminIdpGroupsListResponse](../../models/operations/platformadminidpgroupslistresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## Get

Get one identity provider (IdP) group by its idp_group_id. Inactive IdP groups are returned, so a caller can still show a group after it was removed from the identity provider. An IdP group is active when Glean loaded it from the directory. An IdP group is inactive when only a usage limit names it, for example after the group was removed from the identity provider. An idp_group_id that is not valid, or that names no group that the list returns, returns 404 resource_not_found. The caller must be able to read the workspace configuration or the workspace billing.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-admin-idp-groups-get" method="get" path="/api/admin/idp-groups/{idp_group_id}" -->
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

    res, err := s.Admin.Idpgroups.Get(ctx, "{idp_group_id}")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformGetIdpGroupResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  | Example                                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |                                                                                                                                                                              |
| `idpGroupID`                                                                                                                                                                 | `string`                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                           | The idp_group_id from a list IdP groups response. IDs are opaque. Do not parse or build them. The group_id in usage limit responses is not an idp_group_id and returns 404.<br/> | {idp_group_id}                                                                                                                                                               |
| `opts`                                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                                     | :heavy_minus_sign:                                                                                                                                                           | The options for this request.                                                                                                                                                |                                                                                                                                                                              |

### Response

**[*operations.PlatformAdminIdpGroupsGetResponse](../../models/operations/platformadminidpgroupsgetresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |