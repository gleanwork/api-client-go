# Client.Documents

## Overview

### Available Operations

* [RetrievePermissions](#retrievepermissions) - Read document permissions
* [Retrieve](#retrieve) - Read documents
* [RetrieveByFacets](#retrievebyfacets) - Read documents by facets
* [Summarize](#summarize) - Summarize documents

## RetrievePermissions

Read the emails of all users who have access to the given document.

### Example Usage

<!-- UsageSnippet language="go" operationID="getdocpermissions" method="post" path="/rest/api/v1/getdocpermissions" -->
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

    res, err := s.Client.Documents.RetrievePermissions(ctx, components.GetDocPermissionsRequest{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetDocPermissionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                           | Type                                                                                                                                                                                                | Required                                                                                                                                                                                            | Description                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                 |
| `getDocPermissionsRequest`                                                                                                                                                                          | [components.GetDocPermissionsRequest](../../models/components/getdocpermissionsrequest.md)                                                                                                          | :heavy_check_mark:                                                                                                                                                                                  | Document permissions request                                                                                                                                                                        |
| `locale`                                                                                                                                                                                            | `*string`                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                  | The client's preferred locale in rfc5646 format (e.g. `en`, `ja`, `pt-BR`). If omitted, the `Accept-Language` will be used. If not present or not supported, defaults to the closest match or `en`. |
| `opts`                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | The options for this request.                                                                                                                                                                       |

### Response

**[*operations.GetdocpermissionsResponse](../../models/operations/getdocpermissionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Retrieve

Read the documents including metadata (does not include enhanced metadata via `/documentmetadata`) for the given list of Glean Document IDs or URLs specified in the request.

### Example Usage

<!-- UsageSnippet language="go" operationID="getdocuments" method="post" path="/rest/api/v1/getdocuments" -->
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

    res, err := s.Client.Documents.Retrieve(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetDocumentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                           | Type                                                                                                                                                                                                | Required                                                                                                                                                                                            | Description                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                 |
| `locale`                                                                                                                                                                                            | `*string`                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                  | The client's preferred locale in rfc5646 format (e.g. `en`, `ja`, `pt-BR`). If omitted, the `Accept-Language` will be used. If not present or not supported, defaults to the closest match or `en`. |
| `getDocumentsRequest`                                                                                                                                                                               | [*components.GetDocumentsRequest](../../models/components/getdocumentsrequest.md)                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                  | Information about documents requested.                                                                                                                                                              |
| `opts`                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | The options for this request.                                                                                                                                                                       |

### Response

**[*operations.GetdocumentsResponse](../../models/operations/getdocumentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RetrieveByFacets

Read the documents including metadata (does not include enhanced metadata via `/documentmetadata`) macthing the given facet conditions.

### Example Usage

<!-- UsageSnippet language="go" operationID="getdocumentsbyfacets" method="post" path="/rest/api/v1/getdocumentsbyfacets" -->
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

    res, err := s.Client.Documents.RetrieveByFacets(ctx, nil, &components.GetDocumentsByFacetsRequest{
        FilterSets: []components.FacetFilterSet{
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
        Cursor: apiclientgo.Pointer(""),
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

| Parameter                                                                                                                                                                                           | Type                                                                                                                                                                                                | Required                                                                                                                                                                                            | Description                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                 |
| `locale`                                                                                                                                                                                            | `*string`                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                  | The client's preferred locale in rfc5646 format (e.g. `en`, `ja`, `pt-BR`). If omitted, the `Accept-Language` will be used. If not present or not supported, defaults to the closest match or `en`. |
| `getDocumentsByFacetsRequest`                                                                                                                                                                       | [*components.GetDocumentsByFacetsRequest](../../models/components/getdocumentsbyfacetsrequest.md)                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                  | Information about facet conditions for documents to be retrieved.                                                                                                                                   |
| `opts`                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | The options for this request.                                                                                                                                                                       |

### Response

**[*operations.GetdocumentsbyfacetsResponse](../../models/operations/getdocumentsbyfacetsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Summarize

Generate an AI summary of the requested documents.

### Example Usage

<!-- UsageSnippet language="go" operationID="summarize" method="post" path="/rest/api/v1/summarize" -->
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
            components.CreateDocumentSpecUnionDocumentSpec4(
                components.DocumentSpec4{
                    UgcType: components.DocumentSpecUgcType2Chats,
                    UgcID: "<id>",
                },
            ),
            components.CreateDocumentSpecUnionDocumentSpec4(
                components.DocumentSpec4{
                    UgcType: components.DocumentSpecUgcType2Chats,
                    UgcID: "<id>",
                },
            ),
            components.CreateDocumentSpecUnionDocumentSpec2(
                components.DocumentSpec2{
                    ID: "<id>",
                },
            ),
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SummarizeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                           | Type                                                                                                                                                                                                | Required                                                                                                                                                                                            | Description                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                 |
| `summarizeRequest`                                                                                                                                                                                  | [components.SummarizeRequest](../../models/components/summarizerequest.md)                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                  | Includes request params such as the query and specs of the documents to summarize.                                                                                                                  |
| `locale`                                                                                                                                                                                            | `*string`                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                  | The client's preferred locale in rfc5646 format (e.g. `en`, `ja`, `pt-BR`). If omitted, the `Accept-Language` will be used. If not present or not supported, defaults to the closest match or `en`. |
| `opts`                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | The options for this request.                                                                                                                                                                       |

### Response

**[*operations.SummarizeResponse](../../models/operations/summarizeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |