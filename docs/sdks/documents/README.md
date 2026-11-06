# Documents

## Overview

### Available Operations

* [GetByURL](#getbyurl) - Get document by URL
* [GetByID](#getbyid) - Get document by ID
* [Batch](#batch) - Batch get documents

## GetByURL

Retrieve one permission-visible indexed document by URL. The url is the lookup key that Platform search returns. Many URLs can resolve to one document. The response is a single document, not a list. The returned url is the permalink and may differ from the submitted URL. Gmail and Outlook search results and chat citations carry caller-specific IDs that get document by ID cannot resolve, so retrieve them by URL here. A Gmail or Outlook URL that Platform search returned can still return 404, as it can with legacy getdocuments. Retrieving search results is best-effort parity with legacy getdocuments. Federated results and go-link shortcut documents are not retrievable in this release.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-documents-get-by-url" method="get" path="/api/documents" -->
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

    res, err := s.Documents.GetByURL(ctx, "{url}", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformDocumentGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                           | Type                                                                                                                                | Required                                                                                                                            | Description                                                                                                                         | Example                                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                               | :heavy_check_mark:                                                                                                                  | The context to use for the request.                                                                                                 |                                                                                                                                     |
| `url_`                                                                                                                              | `string`                                                                                                                            | :heavy_check_mark:                                                                                                                  | URL of the document to retrieve.                                                                                                    | {url}                                                                                                                               |
| `include`                                                                                                                           | [][components.PlatformDocumentInclude](../../models/components/platformdocumentinclude.md)                                          | :heavy_minus_sign:                                                                                                                  | Optional field groups to add to each found document. Each value only adds fields; it does not change which documents are returned.<br/> |                                                                                                                                     |
| `opts`                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                            | :heavy_minus_sign:                                                                                                                  | The options for this request.                                                                                                       |                                                                                                                                     |

### Response

**[*operations.PlatformDocumentsGetByURLResponse](../../models/operations/platformdocumentsgetbyurlresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## GetByID

Retrieve one permission-visible indexed document by its Glean store ID. Gmail and Outlook search results and chat citations can carry caller-specific IDs that this operation cannot resolve. Try get document by URL for those documents. URL lookup does not resolve every Gmail or Outlook result.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-documents-get" method="get" path="/api/documents/{document_id}" -->
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

    res, err := s.Documents.GetByID(ctx, "{document_id}", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformDocumentGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                           | Type                                                                                                                                | Required                                                                                                                            | Description                                                                                                                         | Example                                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                               | :heavy_check_mark:                                                                                                                  | The context to use for the request.                                                                                                 |                                                                                                                                     |
| `documentID`                                                                                                                        | `string`                                                                                                                            | :heavy_check_mark:                                                                                                                  | Indexed Glean store ID. For a Gmail or Outlook citation with a caller-specific ID, try get document by URL instead.<br/>            | {document_id}                                                                                                                       |
| `include`                                                                                                                           | [][components.PlatformDocumentInclude](../../models/components/platformdocumentinclude.md)                                          | :heavy_minus_sign:                                                                                                                  | Optional field groups to add to each found document. Each value only adds fields; it does not change which documents are returned.<br/> |                                                                                                                                     |
| `opts`                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                            | :heavy_minus_sign:                                                                                                                  | The options for this request.                                                                                                       |                                                                                                                                     |

### Response

**[*operations.PlatformDocumentsGetResponse](../../models/operations/platformdocumentsgetresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## Batch

Retrieve up to 100 permission-visible indexed documents by URL or document_id. Found items are in results. Complete misses are warnings correlated by index. Incomplete store work fails the whole request. Retrieving search results is best-effort parity with legacy getdocuments. Gmail and Outlook caller-specific citation IDs are not accepted as document_id selectors. Use URL selectors for them. A Gmail or Outlook URL can still return a resource_not_found warning. Federated results and go-link shortcut documents are not retrievable in this release.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-documents-batch" method="post" path="/api/documents/batch" -->
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

    res, err := s.Documents.Batch(ctx, components.PlatformDocumentsBatchRequest{
        DocumentSpecs: []components.PlatformDocumentsBatchRequestDocumentSpec{
            components.CreatePlatformDocumentsBatchRequestDocumentSpecPlatformDocumentsBatchSpecByID(
                components.PlatformDocumentsBatchSpecByID{
                    DocumentID: "{document_id}",
                },
            ),
            components.CreatePlatformDocumentsBatchRequestDocumentSpecPlatformDocumentsBatchSpecByURL(
                components.PlatformDocumentsBatchSpecByURL{
                    URL: "https://company.atlassian.net/wiki/spaces/ENG/pages/12345",
                },
            ),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformDocumentsBatchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [components.PlatformDocumentsBatchRequest](../../models/components/platformdocumentsbatchrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PlatformDocumentsBatchResponse](../../models/operations/platformdocumentsbatchresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 413, 429    | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |