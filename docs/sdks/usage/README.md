# Admin.Usage

## Overview

### Available Operations

* [GetSettings](#getsettings) - Get usage settings
* [UpdateSettings](#updatesettings) - Update usage settings

## GetSettings

Retrieve workspace-wide policy from the same configuration used by Billing. Group resolution uses the configured policy and its backend defaults; business justification is currently optional. No per-user settings are read, and updated_at is omitted because no policy-update timestamp is available. Requires admin.usage.settings:read and permission to read workspace billing; a global administrator role is not required.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-admin-usage-settings-get" method="get" path="/api/admin/usage/settings" -->
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

    res, err := s.Admin.Usage.GetSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformUsageSettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PlatformAdminUsageSettingsGetResponse](../../models/operations/platformadminusagesettingsgetresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## UpdateSettings

Assign workspace-wide usage policy. Assignment is currently unavailable until a writer to authoritative Billing configuration is integrated. The backend config platform.usageSettingsWritesEnabled is disabled by default; enabling it alone does not implement assignment, and an authorized request still returns 503 without changing configuration or persisting a second copy of policy. Requires admin.usage.settings:write and permission to edit workspace billing; a global administrator role is not required.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-admin-usage-settings-update" method="patch" path="/api/admin/usage/settings" -->
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

    res, err := s.Admin.Usage.UpdateSettings(ctx, components.CreatePlatformUsageSettingsUpdateRequestUnionPlatformUsageSettingsUpdateRequest1(
        components.PlatformUsageSettingsUpdateRequest1{
            MultipleMembershipResolution: components.PlatformMultipleMembershipResolutionSettings{
                PerMember: components.PlatformPerMemberSettings{
                    IdpGroups: components.IdpGroupsHighest,
                },
            },
            LimitIncreaseRequests: &components.PlatformLimitIncreaseRequestSettings{
                BusinessJustification: components.BusinessJustificationRequired,
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformUsageSettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [components.PlatformUsageSettingsUpdateRequestUnion](../../models/components/platformusagesettingsupdaterequestunion.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.PlatformAdminUsageSettingsUpdateResponse](../../models/operations/platformadminusagesettingsupdateresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 413, 429    | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |