# ClientDocuments
(*Client.Documents*)

## Overview

### Available Operations

* [RetrievePermissions](#retrievepermissions) - Read document permissions
* [Retrieve](#retrieve) - Read documents
* [RetrieveByFacets](#retrievebyfacets) - Read documents by facets
* [Summarize](#summarize) - Summarize documents

## RetrievePermissions

Read the emails of all users who have access to the given document.

### Example Usage

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

    res, err := s.Client.Documents.RetrievePermissions(ctx, components.GetDocPermissionsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.GetDocPermissionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.GetDocPermissionsRequest](../../models/components/getdocpermissionsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetdocpermissionsResponse](../../models/operations/getdocpermissionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Retrieve

Read the documents including metadata (does not include enhanced metadata via `/documentmetadata`) for the given list of Glean Document IDs or URLs specified in the request.

### Example Usage

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

    res, err := s.Client.Documents.Retrieve(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetDocumentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.GetDocumentsRequest](../../models/components/getdocumentsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetdocumentsResponse](../../models/operations/getdocumentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RetrieveByFacets

Read the documents including metadata (does not include enhanced metadata via `/documentmetadata`) macthing the given facet conditions.

### Example Usage

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

    res, err := s.Client.Documents.RetrieveByFacets(ctx, &components.GetDocumentsByFacetsRequest{
        FilterSets: []components.FacetFilterSet{
            components.FacetFilterSet{
                Filters: []components.FacetFilter{
                    components.FacetFilter{
                        FieldName: apiclientgo.String("type"),
                        Values: []components.FacetFilterValue{
                            components.FacetFilterValue{
                                Value: apiclientgo.String("Spreadsheet"),
                                RelationType: components.RelationTypeEquals.ToPointer(),
                            },
                            components.FacetFilterValue{
                                Value: apiclientgo.String("Presentation"),
                                RelationType: components.RelationTypeEquals.ToPointer(),
                            },
                        },
                    },
                },
            },
            components.FacetFilterSet{
                Filters: []components.FacetFilter{
                    components.FacetFilter{
                        FieldName: apiclientgo.String("type"),
                        Values: []components.FacetFilterValue{
                            components.FacetFilterValue{
                                Value: apiclientgo.String("Spreadsheet"),
                                RelationType: components.RelationTypeEquals.ToPointer(),
                            },
                            components.FacetFilterValue{
                                Value: apiclientgo.String("Presentation"),
                                RelationType: components.RelationTypeEquals.ToPointer(),
                            },
                        },
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetDocumentsByFacetsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.GetDocumentsByFacetsRequest](../../models/components/getdocumentsbyfacetsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetdocumentsbyfacetsResponse](../../models/operations/getdocumentsbyfacetsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Summarize

Generate an AI summary of the requested documents.

### Example Usage

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

    res, err := s.Client.Documents.Summarize(ctx, components.SummarizeRequest{
        DocumentSpecs: []components.DocumentSpecUnion{
            components.CreateDocumentSpecUnionDocumentSpec1(
                components.DocumentSpec1{},
            ),
            components.CreateDocumentSpecUnionDocumentSpec1(
                components.DocumentSpec1{},
            ),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SummarizeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.SummarizeRequest](../../models/components/summarizerequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.SummarizeResponse](../../models/operations/summarizeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |