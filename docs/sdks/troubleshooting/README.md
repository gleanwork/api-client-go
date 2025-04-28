# Troubleshooting
(*Indexing.Troubleshooting*)

## Overview

### Available Operations

* [GetDatasourceStatus](#getdatasourcestatus) - Beta: Get datasource status

* [PostDocumentDebug](#postdocumentdebug) - Beta: Get document information

* [PostDocumentsDebug](#postdocumentsdebug) - Beta: Get information of a batch of documents

* [DebugUser](#debuguser) - Beta: Get user information

* [CheckAccess](#checkaccess) - Check document access
* [~~GetStatus~~](#getstatus) - Get document upload and indexing status :warning: **Deprecated**
* [~~GetDocumentCount~~](#getdocumentcount) - Get document count :warning: **Deprecated**
* [~~GetUserCount~~](#getusercount) - Get user count :warning: **Deprecated**

## GetDatasourceStatus

Gather information about the datasource's overall status. Currently in beta, might undergo breaking changes without prior notice.

Tip: Refer to the [Troubleshooting tutorial](https://developers.glean.com/docs/indexing_api/indexing_api_troubleshooting/) for more information.


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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
    )

    res, err := s.Indexing.Troubleshooting.GetDatasourceStatus(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DebugDatasourceStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `datasource`                                             | *string*                                                 | :heavy_check_mark:                                       | The datasource to get debug status for.                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PostAPIIndexV1DebugDatasourceStatusResponse](../../models/operations/postapiindexv1debugdatasourcestatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PostDocumentDebug

Gives various information that would help in debugging related to a particular document. Currently in beta, might undergo breaking changes without prior notice.

Tip: Refer to the [Troubleshooting tutorial](https://developers.glean.com/docs/indexing_api/indexing_api_troubleshooting/) for more information.


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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
    )

    res, err := s.Indexing.Troubleshooting.PostDocumentDebug(ctx, "<value>", components.DebugDocumentRequest{
        ObjectType: "Article",
        DocID: "art123",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DebugDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `datasource`                                                                       | *string*                                                                           | :heavy_check_mark:                                                                 | The datasource to which the document belongs                                       |
| `debugDocumentRequest`                                                             | [components.DebugDocumentRequest](../../models/components/debugdocumentrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.PostAPIIndexV1DebugDatasourceDocumentResponse](../../models/operations/postapiindexv1debugdatasourcedocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PostDocumentsDebug

Gives various information that would help in debugging related to a batch of documents. Currently in beta, might undergo breaking changes without prior notice.

Tip: Refer to the [Troubleshooting tutorial](https://developers.glean.com/docs/indexing_api/indexing_api_troubleshooting/) for more information.


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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
    )

    res, err := s.Indexing.Troubleshooting.PostDocumentsDebug(ctx, "<value>", components.DebugDocumentsRequest{
        DebugDocuments: []components.DebugDocumentRequest{
            components.DebugDocumentRequest{
                ObjectType: "Article",
                DocID: "art123",
            },
            components.DebugDocumentRequest{
                ObjectType: "Article",
                DocID: "art123",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DebugDocumentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `datasource`                                                                         | *string*                                                                             | :heavy_check_mark:                                                                   | The datasource to which the document belongs                                         |
| `debugDocumentsRequest`                                                              | [components.DebugDocumentsRequest](../../models/components/debugdocumentsrequest.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PostAPIIndexV1DebugDatasourceDocumentsResponse](../../models/operations/postapiindexv1debugdatasourcedocumentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DebugUser

Gives various information that would help in debugging related to a particular user. Currently in beta, might undergo breaking changes without prior notice.

Tip: Refer to the [Troubleshooting tutorial](https://developers.glean.com/docs/indexing_api/indexing_api_troubleshooting/) for more information.


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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
    )

    res, err := s.Indexing.Troubleshooting.DebugUser(ctx, "<value>", components.DebugUserRequest{
        Email: "u1@foo.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DebugUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `datasource`                                                               | *string*                                                                   | :heavy_check_mark:                                                         | The datasource to which the user belongs                                   |
| `debugUserRequest`                                                         | [components.DebugUserRequest](../../models/components/debuguserrequest.md) | :heavy_check_mark:                                                         | N/A                                                                        |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.PostAPIIndexV1DebugDatasourceUserResponse](../../models/operations/postapiindexv1debugdatasourceuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CheckAccess

Check if a given user has access to access a document in a custom datasource

Tip: Refer to the [Troubleshooting tutorial](https://developers.glean.com/docs/indexing_api/indexing_api_troubleshooting/) for more information.


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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
    )

    res, err := s.Indexing.Troubleshooting.CheckAccess(ctx, components.CheckDocumentAccessRequest{
        Datasource: "<value>",
        ObjectType: "<value>",
        DocID: "<id>",
        UserEmail: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CheckDocumentAccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.CheckDocumentAccessRequest](../../models/components/checkdocumentaccessrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PostAPIIndexV1CheckdocumentaccessResponse](../../models/operations/postapiindexv1checkdocumentaccessresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ~~GetStatus~~

Intended for debugging/validation. Fetches the current upload and indexing status of documents.

Tip: Use [/debug/{datasource}/document](https://developers.glean.com/docs/indexing_api/indexing_api_troubleshooting/#debug-datasource-document) for richer information.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
    )

    res, err := s.Indexing.Troubleshooting.GetStatus(ctx, components.GetDocumentStatusRequest{
        Datasource: "<value>",
        ObjectType: "<value>",
        DocID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetDocumentStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.GetDocumentStatusRequest](../../models/components/getdocumentstatusrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PostAPIIndexV1GetdocumentstatusResponse](../../models/operations/postapiindexv1getdocumentstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ~~GetDocumentCount~~

Fetches document count for the specified custom datasource.

Tip: Use [/debug/{datasource}/status](https://developers.glean.com/docs/indexing_api/indexing_api_troubleshooting/#debug-datasource-status) for richer information.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
    )

    res, err := s.Indexing.Troubleshooting.GetDocumentCount(ctx, components.GetDocumentCountRequest{
        Datasource: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetDocumentCountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.GetDocumentCountRequest](../../models/components/getdocumentcountrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PostAPIIndexV1GetdocumentcountResponse](../../models/operations/postapiindexv1getdocumentcountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ~~GetUserCount~~

Fetches user count for the specified custom datasource.

Tip: Use [/debug/{datasource}/status](https://developers.glean.com/docs/indexing_api/indexing_api_troubleshooting/#debug-datasource-status) for richer information.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
    )

    res, err := s.Indexing.Troubleshooting.GetUserCount(ctx, components.GetUserCountRequest{
        Datasource: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetUserCountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.GetUserCountRequest](../../models/components/getusercountrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.PostAPIIndexV1GetusercountResponse](../../models/operations/postapiindexv1getusercountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |