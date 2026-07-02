# Client.Search

## Overview

### Available Operations

* [QueryAsAdmin](#queryasadmin) - Search the index (admin)
* [Autocomplete](#autocomplete) - Autocomplete
* [RetrieveFeed](#retrievefeed) - Feed of documents and events
* [Recommendations](#recommendations) - Recommend documents
* [Query](#query) - Search

## QueryAsAdmin

Retrieves results for search query without respect for permissions. This is available only to privileged users.

### Example Usage

<!-- UsageSnippet language="go" operationID="adminsearch" method="post" path="/rest/api/v1/adminsearch" -->
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

    res, err := s.Client.Search.QueryAsAdmin(ctx, components.SearchRequest{
        TrackingToken: apiclientgo.Pointer("trackingToken"),
        PageSize: apiclientgo.Pointer[int64](10),
        Query: "vacation policy",
        RequestOptions: &components.SearchRequestOptions{
            FacetFilters: []components.FacetFilter{
                components.FacetFilter{
                    FieldName: apiclientgo.Pointer("type"),
                    Values: []components.FacetFilterValue{
                        components.FacetFilterValue{
                            Value: apiclientgo.Pointer("article"),
                            RelationType: components.RelationTypeEquals.ToPointer(),
                        },
                        components.FacetFilterValue{
                            Value: apiclientgo.Pointer("document"),
                            RelationType: components.RelationTypeEquals.ToPointer(),
                        },
                    },
                },
                components.FacetFilter{
                    FieldName: apiclientgo.Pointer("department"),
                    Values: []components.FacetFilterValue{
                        components.FacetFilterValue{
                            Value: apiclientgo.Pointer("engineering"),
                            RelationType: components.RelationTypeEquals.ToPointer(),
                        },
                    },
                },
            },
            FacetBucketSize: 421489,
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                                                                                                                                                                                                                                                                                      | Required                                                                                                                                                                                                                                                                                                                                                                  | Description                                                                                                                                                                                                                                                                                                                                                               | Example                                                                                                                                                                                                                                                                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                        | The context to use for the request.                                                                                                                                                                                                                                                                                                                                       |                                                                                                                                                                                                                                                                                                                                                                           |
| `searchRequest`                                                                                                                                                                                                                                                                                                                                                           | [components.SearchRequest](../../models/components/searchrequest.md)                                                                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                        | Admin search request                                                                                                                                                                                                                                                                                                                                                      | {<br/>"trackingToken": "trackingToken",<br/>"query": "vacation policy",<br/>"pageSize": 10,<br/>"requestOptions": {<br/>"facetFilters": [<br/>{<br/>"fieldName": "type",<br/>"values": [<br/>{<br/>"value": "article",<br/>"relationType": "EQUALS"<br/>},<br/>{<br/>"value": "document",<br/>"relationType": "EQUALS"<br/>}<br/>]<br/>},<br/>{<br/>"fieldName": "department",<br/>"values": [<br/>{<br/>"value": "engineering",<br/>"relationType": "EQUALS"<br/>}<br/>]<br/>}<br/>]<br/>}<br/>} |
| `locale`                                                                                                                                                                                                                                                                                                                                                                  | `*string`                                                                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                        | The client's preferred locale in rfc5646 format (e.g. `en`, `ja`, `pt-BR`). If omitted, the `Accept-Language` will be used. If not present or not supported, defaults to the closest match or `en`.                                                                                                                                                                       |                                                                                                                                                                                                                                                                                                                                                                           |
| `opts`                                                                                                                                                                                                                                                                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                        | The options for this request.                                                                                                                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                                                                                                                                           |

### Response

**[*operations.AdminsearchResponse](../../models/operations/adminsearchresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| apierrors.GleanDataError | 403, 422                 | application/json         |
| apierrors.APIError       | 4XX, 5XX                 | \*/\*                    |

## Autocomplete

Retrieve query suggestions, operators and documents for the given partially typed query.

### Example Usage

<!-- UsageSnippet language="go" operationID="autocomplete" method="post" path="/rest/api/v1/autocomplete" -->
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

    res, err := s.Client.Search.Autocomplete(ctx, components.AutocompleteRequest{
        TrackingToken: apiclientgo.Pointer("trackingToken"),
        Query: apiclientgo.Pointer("what is a que"),
        Datasource: apiclientgo.Pointer("GDRIVE"),
        ResultSize: apiclientgo.Pointer[int64](10),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AutocompleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                           | Type                                                                                                                                                                                                | Required                                                                                                                                                                                            | Description                                                                                                                                                                                         | Example                                                                                                                                                                                             |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                 |                                                                                                                                                                                                     |
| `autocompleteRequest`                                                                                                                                                                               | [components.AutocompleteRequest](../../models/components/autocompleterequest.md)                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                  | Autocomplete request                                                                                                                                                                                | {<br/>"trackingToken": "trackingToken",<br/>"query": "what is a que",<br/>"datasource": "GDRIVE",<br/>"resultSize": 10<br/>}                                                                        |
| `locale`                                                                                                                                                                                            | `*string`                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                  | The client's preferred locale in rfc5646 format (e.g. `en`, `ja`, `pt-BR`). If omitted, the `Accept-Language` will be used. If not present or not supported, defaults to the closest match or `en`. |                                                                                                                                                                                                     |
| `opts`                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | The options for this request.                                                                                                                                                                       |                                                                                                                                                                                                     |

### Response

**[*operations.AutocompleteResponse](../../models/operations/autocompleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RetrieveFeed

The personalized feed/home includes different types of contents including suggestions, recents, calendar events and many more.

### Example Usage

<!-- UsageSnippet language="go" operationID="feed" method="post" path="/rest/api/v1/feed" -->
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

    res, err := s.Client.Search.RetrieveFeed(ctx, components.FeedRequest{
        TimeoutMillis: apiclientgo.Pointer[int64](5000),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FeedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                           | Type                                                                                                                                                                                                | Required                                                                                                                                                                                            | Description                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                 |
| `feedRequest`                                                                                                                                                                                       | [components.FeedRequest](../../models/components/feedrequest.md)                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                  | Includes request params, client data and more for making user's feed.                                                                                                                               |
| `locale`                                                                                                                                                                                            | `*string`                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                  | The client's preferred locale in rfc5646 format (e.g. `en`, `ja`, `pt-BR`). If omitted, the `Accept-Language` will be used. If not present or not supported, defaults to the closest match or `en`. |
| `opts`                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | The options for this request.                                                                                                                                                                       |

### Response

**[*operations.FeedResponse](../../models/operations/feedresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Recommendations

Retrieve recommended documents for the given URL or Glean Document ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="recommendations" method="post" path="/rest/api/v1/recommendations" -->
```go
package main

import(
	"context"
	"os"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/types"
	"github.com/gleanwork/api-client-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Client.Search.Recommendations(ctx, components.RecommendationsRequest{
        SourceDocument: &components.Document{
            ContainerDocument: &components.Document{
                Metadata: &components.DocumentMetadata{
                    Datasource: apiclientgo.Pointer("datasource"),
                    ObjectType: apiclientgo.Pointer("Feature Request"),
                    Container: apiclientgo.Pointer("container"),
                    ParentID: apiclientgo.Pointer("JIRA_EN-1337"),
                    MimeType: apiclientgo.Pointer("mimeType"),
                    DocumentID: apiclientgo.Pointer("documentId"),
                    CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                    UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                    Author: &components.Person{
                        Name: "name",
                        ObfuscatedID: "<id>",
                    },
                    Components: []string{
                        "Backend",
                        "Networking",
                    },
                    Status: apiclientgo.Pointer("[\"Done\"]"),
                    CustomData: map[string]components.CustomDataValue{
                        "someCustomField": components.CustomDataValue{},
                    },
                },
            },
            ParentDocument: &components.Document{
                Metadata: &components.DocumentMetadata{
                    Datasource: apiclientgo.Pointer("datasource"),
                    ObjectType: apiclientgo.Pointer("Feature Request"),
                    Container: apiclientgo.Pointer("container"),
                    ParentID: apiclientgo.Pointer("JIRA_EN-1337"),
                    MimeType: apiclientgo.Pointer("mimeType"),
                    DocumentID: apiclientgo.Pointer("documentId"),
                    CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                    UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                    Author: &components.Person{
                        Name: "name",
                        ObfuscatedID: "<id>",
                    },
                    Components: []string{
                        "Backend",
                        "Networking",
                    },
                    Status: apiclientgo.Pointer("[\"Done\"]"),
                    CustomData: map[string]components.CustomDataValue{
                        "someCustomField": components.CustomDataValue{},
                    },
                },
            },
            Metadata: &components.DocumentMetadata{
                Datasource: apiclientgo.Pointer("datasource"),
                ObjectType: apiclientgo.Pointer("Feature Request"),
                Container: apiclientgo.Pointer("container"),
                ParentID: apiclientgo.Pointer("JIRA_EN-1337"),
                MimeType: apiclientgo.Pointer("mimeType"),
                DocumentID: apiclientgo.Pointer("documentId"),
                CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                Author: &components.Person{
                    Name: "name",
                    ObfuscatedID: "abc123",
                },
                Components: []string{
                    "Backend",
                    "Networking",
                },
                Status: apiclientgo.Pointer("[\"Done\"]"),
                CustomData: map[string]components.CustomDataValue{
                    "someCustomField": components.CustomDataValue{},
                },
            },
        },
        PageSize: apiclientgo.Pointer[int64](100),
        MaxSnippetSize: apiclientgo.Pointer[int64](400),
        RequestOptions: &components.RecommendationsRequestOptions{
            FacetFilterSets: []components.FacetFilterSet{
                components.FacetFilterSet{
                    Filters: []components.FacetFilter{
                        components.FacetFilter{
                            FieldName: apiclientgo.Pointer("type"),
                            Values: []components.FacetFilterValue{
                                components.FacetFilterValue{
                                    Value: apiclientgo.Pointer("Spreadsheet"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                                components.FacetFilterValue{
                                    Value: apiclientgo.Pointer("Presentation"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                            },
                        },
                    },
                },
                components.FacetFilterSet{
                    Filters: []components.FacetFilter{
                        components.FacetFilter{
                            FieldName: apiclientgo.Pointer("type"),
                            Values: []components.FacetFilterValue{
                                components.FacetFilterValue{
                                    Value: apiclientgo.Pointer("Spreadsheet"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                                components.FacetFilterValue{
                                    Value: apiclientgo.Pointer("Presentation"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                            },
                        },
                    },
                },
                components.FacetFilterSet{
                    Filters: []components.FacetFilter{
                        components.FacetFilter{
                            FieldName: apiclientgo.Pointer("type"),
                            Values: []components.FacetFilterValue{
                                components.FacetFilterValue{
                                    Value: apiclientgo.Pointer("Spreadsheet"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                                components.FacetFilterValue{
                                    Value: apiclientgo.Pointer("Presentation"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                            },
                        },
                    },
                },
            },
            Context: &components.Document{
                ContainerDocument: &components.Document{
                    Metadata: &components.DocumentMetadata{
                        Datasource: apiclientgo.Pointer("datasource"),
                        ObjectType: apiclientgo.Pointer("Feature Request"),
                        Container: apiclientgo.Pointer("container"),
                        ParentID: apiclientgo.Pointer("JIRA_EN-1337"),
                        MimeType: apiclientgo.Pointer("mimeType"),
                        DocumentID: apiclientgo.Pointer("documentId"),
                        CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                        UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                        Author: &components.Person{
                            Name: "name",
                            ObfuscatedID: "<id>",
                        },
                        Components: []string{
                            "Backend",
                            "Networking",
                        },
                        Status: apiclientgo.Pointer("[\"Done\"]"),
                        CustomData: map[string]components.CustomDataValue{
                            "someCustomField": components.CustomDataValue{},
                        },
                    },
                },
                ParentDocument: &components.Document{
                    Metadata: &components.DocumentMetadata{
                        Datasource: apiclientgo.Pointer("datasource"),
                        ObjectType: apiclientgo.Pointer("Feature Request"),
                        Container: apiclientgo.Pointer("container"),
                        ParentID: apiclientgo.Pointer("JIRA_EN-1337"),
                        MimeType: apiclientgo.Pointer("mimeType"),
                        DocumentID: apiclientgo.Pointer("documentId"),
                        CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                        UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                        Author: &components.Person{
                            Name: "name",
                            ObfuscatedID: "<id>",
                        },
                        Components: []string{
                            "Backend",
                            "Networking",
                        },
                        Status: apiclientgo.Pointer("[\"Done\"]"),
                        CustomData: map[string]components.CustomDataValue{
                            "someCustomField": components.CustomDataValue{},
                        },
                    },
                },
                Metadata: &components.DocumentMetadata{
                    Datasource: apiclientgo.Pointer("datasource"),
                    ObjectType: apiclientgo.Pointer("Feature Request"),
                    Container: apiclientgo.Pointer("container"),
                    ParentID: apiclientgo.Pointer("JIRA_EN-1337"),
                    MimeType: apiclientgo.Pointer("mimeType"),
                    DocumentID: apiclientgo.Pointer("documentId"),
                    CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                    UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                    Author: &components.Person{
                        Name: "name",
                        ObfuscatedID: "abc123",
                    },
                    Components: []string{
                        "Backend",
                        "Networking",
                    },
                    Status: apiclientgo.Pointer("[\"Done\"]"),
                    CustomData: map[string]components.CustomDataValue{
                        "someCustomField": components.CustomDataValue{},
                    },
                },
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResultsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                           | Type                                                                                                                                                                                                | Required                                                                                                                                                                                            | Description                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                 |
| `recommendationsRequest`                                                                                                                                                                            | [components.RecommendationsRequest](../../models/components/recommendationsrequest.md)                                                                                                              | :heavy_check_mark:                                                                                                                                                                                  | Recommendations request                                                                                                                                                                             |
| `locale`                                                                                                                                                                                            | `*string`                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                  | The client's preferred locale in rfc5646 format (e.g. `en`, `ja`, `pt-BR`). If omitted, the `Accept-Language` will be used. If not present or not supported, defaults to the closest match or `en`. |
| `opts`                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | The options for this request.                                                                                                                                                                       |

### Response

**[*operations.RecommendationsResponse](../../models/operations/recommendationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Query

Retrieve results from the index for the given query and filters.

### Example Usage

<!-- UsageSnippet language="go" operationID="search" method="post" path="/rest/api/v1/search" -->
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

    res, err := s.Client.Search.Query(ctx, components.SearchRequest{
        TrackingToken: apiclientgo.Pointer("trackingToken"),
        PageSize: apiclientgo.Pointer[int64](10),
        Query: "vacation policy",
        RequestOptions: &components.SearchRequestOptions{
            FacetFilters: []components.FacetFilter{
                components.FacetFilter{
                    FieldName: apiclientgo.Pointer("type"),
                    Values: []components.FacetFilterValue{
                        components.FacetFilterValue{
                            Value: apiclientgo.Pointer("article"),
                            RelationType: components.RelationTypeEquals.ToPointer(),
                        },
                        components.FacetFilterValue{
                            Value: apiclientgo.Pointer("document"),
                            RelationType: components.RelationTypeEquals.ToPointer(),
                        },
                    },
                },
                components.FacetFilter{
                    FieldName: apiclientgo.Pointer("department"),
                    Values: []components.FacetFilterValue{
                        components.FacetFilterValue{
                            Value: apiclientgo.Pointer("engineering"),
                            RelationType: components.RelationTypeEquals.ToPointer(),
                        },
                    },
                },
            },
            FacetBucketSize: 400611,
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                                                                                                                                                                                                                                                                                      | Required                                                                                                                                                                                                                                                                                                                                                                  | Description                                                                                                                                                                                                                                                                                                                                                               | Example                                                                                                                                                                                                                                                                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                        | The context to use for the request.                                                                                                                                                                                                                                                                                                                                       |                                                                                                                                                                                                                                                                                                                                                                           |
| `searchRequest`                                                                                                                                                                                                                                                                                                                                                           | [components.SearchRequest](../../models/components/searchrequest.md)                                                                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                        | Search request                                                                                                                                                                                                                                                                                                                                                            | {<br/>"trackingToken": "trackingToken",<br/>"query": "vacation policy",<br/>"pageSize": 10,<br/>"requestOptions": {<br/>"facetFilters": [<br/>{<br/>"fieldName": "type",<br/>"values": [<br/>{<br/>"value": "article",<br/>"relationType": "EQUALS"<br/>},<br/>{<br/>"value": "document",<br/>"relationType": "EQUALS"<br/>}<br/>]<br/>},<br/>{<br/>"fieldName": "department",<br/>"values": [<br/>{<br/>"value": "engineering",<br/>"relationType": "EQUALS"<br/>}<br/>]<br/>}<br/>]<br/>}<br/>} |
| `locale`                                                                                                                                                                                                                                                                                                                                                                  | `*string`                                                                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                        | The client's preferred locale in rfc5646 format (e.g. `en`, `ja`, `pt-BR`). If omitted, the `Accept-Language` will be used. If not present or not supported, defaults to the closest match or `en`.                                                                                                                                                                       |                                                                                                                                                                                                                                                                                                                                                                           |
| `opts`                                                                                                                                                                                                                                                                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                        | The options for this request.                                                                                                                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                                                                                                                                           |

### Response

**[*operations.SearchResponse](../../models/operations/searchresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| apierrors.GleanDataError | 403, 422                 | application/json         |
| apierrors.APIError       | 4XX, 5XX                 | \*/\*                    |