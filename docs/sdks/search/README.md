# Search

## Overview

### Available Operations

* [Query](#query) - Search
* [ListFilters](#listfilters) - List search filters

## Query

Search your organization's connected content and return ranked document results with cursor pagination. Use `GET /api/search/filters` to discover datasource identifiers and common filter fields. Built-in filter names are validated; other field names are accepted as custom filters and behavior depends on your Glean configuration and connected sources.
Errors: HTTP 422 `unprocessable_query` returns no `results` or `next_cursor`. See `warnings` on the response for non-blocking issues such as partially available results. Not every query issue produces a warning or error.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-search" method="post" path="/api/search" -->
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

    res, err := s.Search.Query(ctx, components.PlatformSearchRequest{
        Query: "quarterly planning 2026",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.PlatformSearchRequest](../../models/components/platformsearchrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PlatformSearchResponse](../../models/operations/platformsearchresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.PlatformProblemDetailError   | 400, 401, 403, 404, 408, 413, 422, 429 | application/problem+json               |
| apierrors.PlatformProblemDetailError   | 500, 503                               | application/problem+json               |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ListFilters

List datasources and common built-in filter fields visible to the authenticated user. This is a best-effort catalog, not an exhaustive list of every filter search accepts.
Without `query`, returns field metadata only and does not run a search. With a nonblank `query`, provide exactly one `datasources` value to request suggested filter values for that query; no documents are returned and this endpoint does not include warning objects. See `FilterFieldInfo.values` for limitations on suggested values. Rate-limited requests return HTTP 429 with `Retry-After`; temporary backend unavailability returns HTTP 503.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-search-filters" method="get" path="/api/search/filters" -->
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

    res, err := s.Search.ListFilters(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformSearchFiltersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                  | Type                                                                                                                                                                                                                                       | Required                                                                                                                                                                                                                                   | Description                                                                                                                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                                                        |
| `datasources`                                                                                                                                                                                                                              | []`string`                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                         | Restrict metadata to one or more datasource identifiers as returned by this endpoint (for example, `jira`). With a nonblank `query`, exactly one datasource is required. Unknown or inaccessible identifiers return `invalid_datasource`.<br/> |
| `query`                                                                                                                                                                                                                                    | `*string`                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                         | Optional search query that requests suggested filter values for the selected datasource. Must be nonblank when present. Triggers a search for facet values only; does not return documents.<br/>                                           |
| `opts`                                                                                                                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                         | The options for this request.                                                                                                                                                                                                              |

### Response

**[*operations.PlatformSearchFiltersResponse](../../models/operations/platformsearchfiltersresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |