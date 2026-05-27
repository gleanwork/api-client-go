# Indexing.CustomMetadata

## Overview

### Available Operations

* [Upsert](#upsert) - Add or update custom metadata
* [Delete](#delete) - Remove custom metadata
* [GetSchema](#getschema) - Retrieve metadata schema
* [UpsertSchema](#upsertschema) - Create or update metadata schema
* [DeleteSchema](#deleteschema) - Remove metadata schema

## Upsert

Associates custom metadata with a specific document. Custom metadata enables you to enrich documents with additional structured information that can be used for search, filtering, and faceting.

### Example Usage

<!-- UsageSnippet language="go" operationID="put_/rest/api/index/document/{docId}/custom-metadata/{groupName}" method="put" path="/rest/api/index/document/{docId}/custom-metadata/{groupName}" -->
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

    res, err := s.Indexing.CustomMetadata.Upsert(ctx, "<id>", "<value>", components.CustomMetadataPutRequest{
        CustomMetadata: []components.CustomProperty{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `docID`                                                                                    | `string`                                                                                   | :heavy_check_mark:                                                                         | Unique Glean identifier of the document                                                    |
| `groupName`                                                                                | `string`                                                                                   | :heavy_check_mark:                                                                         | Name of the metadata group as specified while adding schema                                |
| `customMetadataPutRequest`                                                                 | [components.CustomMetadataPutRequest](../../models/components/custommetadataputrequest.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PutRestAPIIndexDocumentDocIDCustomMetadataGroupNameResponse](../../models/operations/putrestapiindexdocumentdocidcustommetadatagroupnameresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| apierrors.ErrorInfoResponse | 400, 401, 404, 429          | application/json            |
| apierrors.ErrorInfoResponse | 500                         | application/json            |
| apierrors.APIError          | 4XX, 5XX                    | \*/\*                       |

## Delete

Removes all custom metadata for the specified metadata group from a document.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_/rest/api/index/document/{docId}/custom-metadata/{groupName}" method="delete" path="/rest/api/index/document/{docId}/custom-metadata/{groupName}" -->
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

    res, err := s.Indexing.CustomMetadata.Delete(ctx, "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |
| `docID`                                                     | `string`                                                    | :heavy_check_mark:                                          | Unique Glean identifier of the document                     |
| `groupName`                                                 | `string`                                                    | :heavy_check_mark:                                          | Name of the metadata group as specified while adding schema |
| `opts`                                                      | [][operations.Option](../../models/operations/option.md)    | :heavy_minus_sign:                                          | The options for this request.                               |

### Response

**[*operations.DeleteRestAPIIndexDocumentDocIDCustomMetadataGroupNameResponse](../../models/operations/deleterestapiindexdocumentdocidcustommetadatagroupnameresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| apierrors.ErrorInfoResponse | 400, 401, 404, 429          | application/json            |
| apierrors.ErrorInfoResponse | 500                         | application/json            |
| apierrors.APIError          | 4XX, 5XX                    | \*/\*                       |

## GetSchema

Retrieves the current schema definition for a metadata group.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/rest/api/index/custom-metadata/schema/{groupName}" method="get" path="/rest/api/index/custom-metadata/schema/{groupName}" -->
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

    res, err := s.Indexing.CustomMetadata.GetSchema(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomMetadataSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `groupName`                                              | `string`                                                 | :heavy_check_mark:                                       | Name of the metadata group schema                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetRestAPIIndexCustomMetadataSchemaGroupNameResponse](../../models/operations/getrestapiindexcustommetadataschemagroupnameresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| apierrors.ErrorInfoResponse | 401, 404, 429               | application/json            |
| apierrors.ErrorInfoResponse | 500                         | application/json            |
| apierrors.APIError          | 4XX, 5XX                    | \*/\*                       |

## UpsertSchema

Defines or updates the schema for a metadata group. Schemas should be defined before indexing metadata.

### Example Usage

<!-- UsageSnippet language="go" operationID="put_/rest/api/index/custom-metadata/schema/{groupName}" method="put" path="/rest/api/index/custom-metadata/schema/{groupName}" -->
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

    res, err := s.Indexing.CustomMetadata.UpsertSchema(ctx, "<value>", components.CustomMetadataSchema{
        MetadataKeys: []components.CustomMetadataPropertyDefinition{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `groupName`                                                                        | `string`                                                                           | :heavy_check_mark:                                                                 | Name of the metadata group schema                                                  |
| `customMetadataSchema`                                                             | [components.CustomMetadataSchema](../../models/components/custommetadataschema.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.PutRestAPIIndexCustomMetadataSchemaGroupNameResponse](../../models/operations/putrestapiindexcustommetadataschemagroupnameresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| apierrors.ErrorInfoResponse | 400, 401, 409, 429          | application/json            |
| apierrors.ErrorInfoResponse | 500                         | application/json            |
| apierrors.APIError          | 4XX, 5XX                    | \*/\*                       |

## DeleteSchema

Deletes the schema definition for a metadata group. This does not delete existing metadata values on documents.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_/rest/api/index/custom-metadata/schema/{groupName}" method="delete" path="/rest/api/index/custom-metadata/schema/{groupName}" -->
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

    res, err := s.Indexing.CustomMetadata.DeleteSchema(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `groupName`                                              | `string`                                                 | :heavy_check_mark:                                       | Name of the metadata group schema                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteRestAPIIndexCustomMetadataSchemaGroupNameResponse](../../models/operations/deleterestapiindexcustommetadataschemagroupnameresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| apierrors.ErrorInfoResponse | 400, 401, 404, 429          | application/json            |
| apierrors.ErrorInfoResponse | 500                         | application/json            |
| apierrors.APIError          | 4XX, 5XX                    | \*/\*                       |