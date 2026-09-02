# Triggers

## Overview

### Available Operations

* [Create](#create) - Create trigger
* [List](#list) - List triggers
* [Get](#get) - Get trigger
* [Update](#update) - Update trigger
* [Delete](#delete) - Delete trigger
* [SearchEvents](#searchevents) - Search events for a trigger
* [ListPresets](#listpresets) - List trigger presets
* [GetPreset](#getpreset) - Get trigger preset
* [ListPresetInputValues](#listpresetinputvalues) - Search trigger preset input values
* [SearchPresetEvents](#searchpresetevents) - Search events for a trigger preset

## Create

Create a trigger from a preset and return it with its signing secret.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-triggers-create" method="post" path="/api/triggers" -->
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

    res, err := s.Triggers.Create(ctx, components.PlatformTriggerCreateRequest{
        PresetID: "GITHUB_1",
        Inputs: map[string]any{
            "repository": "{repository}",
        },
        Delivery: components.PlatformTriggerDelivery{
            WebhookURL: "https://example.com/webhook",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformTriggerCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.PlatformTriggerCreateRequest](../../models/components/platformtriggercreaterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PlatformTriggersCreateResponse](../../models/operations/platformtriggerscreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.PlatformProblemDetailError   | 400, 401, 403, 404, 408, 409, 413, 429 | application/problem+json               |
| apierrors.PlatformProblemDetailError   | 500, 503                               | application/problem+json               |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## List

List triggers owned by the authenticated caller.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-triggers-list" method="get" path="/api/triggers" -->
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

    res, err := s.Triggers.List(ctx, apiclientgo.Pointer[int64](50), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformTriggerListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `pageSize`                                               | `*int64`                                                 | :heavy_minus_sign:                                       | Maximum number of triggers to return.                    |
| `cursor`                                                 | `*string`                                                | :heavy_minus_sign:                                       | Opaque pagination cursor from a previous response.       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PlatformTriggersListResponse](../../models/operations/platformtriggerslistresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 408, 429              | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## Get

Retrieve a trigger owned by the authenticated caller.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-triggers-get" method="get" path="/api/triggers/{trigger_id}" -->
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

    res, err := s.Triggers.Get(ctx, "{trigger_id}")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformTriggerGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `triggerID`                                              | `string`                                                 | :heavy_check_mark:                                       | ID of the trigger to retrieve.                           | {trigger_id}                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PlatformTriggersGetResponse](../../models/operations/platformtriggersgetresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## Update

Update a trigger.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-triggers-update" method="patch" path="/api/triggers/{trigger_id}" -->
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

    res, err := s.Triggers.Update(ctx, "{trigger_id}", components.PlatformTriggerUpdateRequest{
        Inputs: map[string]any{
            "repository": "{repository}",
        },
        Delivery: &components.PlatformTriggerDelivery{
            WebhookURL: "https://example.com/webhook",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformTriggerGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  | Example                                                                                                      |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |                                                                                                              |
| `triggerID`                                                                                                  | `string`                                                                                                     | :heavy_check_mark:                                                                                           | ID of the trigger to update.                                                                                 | {trigger_id}                                                                                                 |
| `platformTriggerUpdateRequest`                                                                               | [components.PlatformTriggerUpdateRequest](../../models/components/platformtriggerupdaterequest.md)           | :heavy_check_mark:                                                                                           | N/A                                                                                                          | {<br/>"inputs": {<br/>"repository": "{repository}"<br/>},<br/>"delivery": {<br/>"webhook_url": "https://example.com/webhook"<br/>}<br/>} |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |                                                                                                              |

### Response

**[*operations.PlatformTriggersUpdateResponse](../../models/operations/platformtriggersupdateresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 413, 429    | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## Delete

Delete a trigger.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-triggers-delete" method="delete" path="/api/triggers/{trigger_id}" -->
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

    res, err := s.Triggers.Delete(ctx, "{trigger_id}")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `triggerID`                                              | `string`                                                 | :heavy_check_mark:                                       | ID of the trigger to delete.                             | {trigger_id}                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PlatformTriggersDeleteResponse](../../models/operations/platformtriggersdeleteresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## SearchEvents

Search recent content events an existing trigger matches. Read-only — no webhook delivery is made. Covers the last seven days.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-triggers-events-search" method="post" path="/api/triggers/{trigger_id}/events/search" -->
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

    res, err := s.Triggers.SearchEvents(ctx, "{trigger_id}", &components.PlatformTriggerEventSearchRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformTriggerEventSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                     | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   | Example                                                                                                       |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                         | :heavy_check_mark:                                                                                            | The context to use for the request.                                                                           |                                                                                                               |
| `triggerID`                                                                                                   | `string`                                                                                                      | :heavy_check_mark:                                                                                            | ID of the trigger whose events to search.                                                                     | {trigger_id}                                                                                                  |
| `platformTriggerEventSearchRequest`                                                                           | [*components.PlatformTriggerEventSearchRequest](../../models/components/platformtriggereventsearchrequest.md) | :heavy_minus_sign:                                                                                            | N/A                                                                                                           | {<br/>"page_size": 10<br/>}                                                                                   |
| `opts`                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                      | :heavy_minus_sign:                                                                                            | The options for this request.                                                                                 |                                                                                                               |

### Response

**[*operations.PlatformTriggersEventsSearchResponse](../../models/operations/platformtriggerseventssearchresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 413, 429    | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## ListPresets

List the trigger presets available to the caller. A preset is a curated content-trigger template (e.g. a new Jira ticket) which is passed when creating a trigger.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-trigger-presets-list" method="get" path="/api/trigger-presets" -->
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

    res, err := s.Triggers.ListPresets(ctx, nil, apiclientgo.Pointer[int64](50), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformTriggerPresetListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `datasource`                                                       | `*string`                                                          | :heavy_minus_sign:                                                 | Restrict results to presets for a single datasource (e.g. github). |
| `pageSize`                                                         | `*int64`                                                           | :heavy_minus_sign:                                                 | Maximum number of presets to return.                               |
| `cursor`                                                           | `*string`                                                          | :heavy_minus_sign:                                                 | Opaque pagination cursor from a previous response.                 |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.PlatformTriggerPresetsListResponse](../../models/operations/platformtriggerpresetslistresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 408, 429              | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## GetPreset

Retrieve a single trigger preset by id.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-trigger-presets-get" method="get" path="/api/trigger-presets/{preset_id}" -->
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

    res, err := s.Triggers.GetPreset(ctx, "{preset_id}")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformTriggerPresetGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `presetID`                                               | `string`                                                 | :heavy_check_mark:                                       | ID of the preset to retrieve.                            | {preset_id}                                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PlatformTriggerPresetsGetResponse](../../models/operations/platformtriggerpresetsgetresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## ListPresetInputValues

Return up to 300 selectable values for a single picklist input on a preset. Results are intended for typeahead selection and are not cursor-paginated. When `is_truncated` is true, refine `query` to narrow the result set.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-trigger-presets-input-values-list" method="get" path="/api/trigger-presets/{preset_id}/input-values" -->
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

    res, err := s.Triggers.ListPresetInputValues(ctx, "{preset_id}", "{field}", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformTriggerPresetInputValueListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          | Example                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |                                                                                                                      |
| `presetID`                                                                                                           | `string`                                                                                                             | :heavy_check_mark:                                                                                                   | ID of the preset the input belongs to.                                                                               | {preset_id}                                                                                                          |
| `field`                                                                                                              | `string`                                                                                                             | :heavy_check_mark:                                                                                                   | Field identifier of the input whose values to list.                                                                  | {field}                                                                                                              |
| `query`                                                                                                              | `*string`                                                                                                            | :heavy_minus_sign:                                                                                                   | Prefix filter over the input's option values, for typeahead. Matching is on the option value, not its display name.<br/> |                                                                                                                      |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |                                                                                                                      |

### Response

**[*operations.PlatformTriggerPresetsInputValuesListResponse](../../models/operations/platformtriggerpresetsinputvalueslistresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## SearchPresetEvents

Search recent content events an unsaved trigger built from this preset would match, to preview it before creating the trigger. Read-only — no trigger is created and no webhook delivery is made. Covers the last seven days.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-trigger-presets-events-search" method="post" path="/api/trigger-presets/{preset_id}/events/search" -->
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

    res, err := s.Triggers.SearchPresetEvents(ctx, "{preset_id}", &components.PlatformTriggerPresetEventSearchRequest{
        Inputs: map[string]string{
            "repository": "{repository}",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformTriggerEventSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                 | Type                                                                                                                      | Required                                                                                                                  | Description                                                                                                               | Example                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                     | :heavy_check_mark:                                                                                                        | The context to use for the request.                                                                                       |                                                                                                                           |
| `presetID`                                                                                                                | `string`                                                                                                                  | :heavy_check_mark:                                                                                                        | ID of the preset to preview.                                                                                              | {preset_id}                                                                                                               |
| `platformTriggerPresetEventSearchRequest`                                                                                 | [*components.PlatformTriggerPresetEventSearchRequest](../../models/components/platformtriggerpreseteventsearchrequest.md) | :heavy_minus_sign:                                                                                                        | N/A                                                                                                                       | {<br/>"inputs": {<br/>"repository": "{repository}"<br/>},<br/>"page_size": 10<br/>}                                       |
| `opts`                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                  | :heavy_minus_sign:                                                                                                        | The options for this request.                                                                                             |                                                                                                                           |

### Response

**[*operations.PlatformTriggerPresetsEventsSearchResponse](../../models/operations/platformtriggerpresetseventssearchresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 413, 429    | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |