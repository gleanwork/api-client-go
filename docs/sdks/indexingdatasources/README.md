# Indexing.Datasources

## Overview

### Available Operations

* [Add](#add) - Add or update datasource
* [RetrieveConfig](#retrieveconfig) - Get datasource config
* [Submit](#submit) - Submit datasource data

## Add

Add or update a custom datasource and its schema.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/api/index/v1/adddatasource" method="post" path="/api/index/v1/adddatasource" -->
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

    res, err := s.Indexing.Datasources.Add(ctx, components.CustomDatasourceConfig{
        Name: "<value>",
        URLRegex: apiclientgo.Pointer("https://example-company.datasource.com/.*"),
        Quicklinks: []components.Quicklink{
            components.Quicklink{
                IconConfig: &components.IconConfig{
                    Color: apiclientgo.Pointer("#343CED"),
                    Key: apiclientgo.Pointer("person_icon"),
                    IconType: components.IconTypeGlyph.ToPointer(),
                    Name: apiclientgo.Pointer("user"),
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.CustomDatasourceConfig](../../models/components/customdatasourceconfig.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.PostAPIIndexV1AdddatasourceResponse](../../models/operations/postapiindexv1adddatasourceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RetrieveConfig

Fetches the datasource config for the specified custom datasource.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/api/index/v1/getdatasourceconfig" method="post" path="/api/index/v1/getdatasourceconfig" -->
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

    res, err := s.Indexing.Datasources.RetrieveConfig(ctx, components.GetDatasourceConfigRequest{
        Datasource: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomDatasourceConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.GetDatasourceConfigRequest](../../models/components/getdatasourceconfigrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PostAPIIndexV1GetdatasourceconfigResponse](../../models/operations/postapiindexv1getdatasourceconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Submit

Validates and asynchronously processes a datasource-specific submission.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/rest/api/index/submissions/{datasourceInstance}/{type}" method="post" path="/rest/api/index/submissions/{datasourceInstance}/{type}" -->
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

    res, err := s.Indexing.Datasources.Submit(ctx, "<value>", "<value>", map[string]any{
        "key": "<value>",
        "key1": "<value>",
        "key2": "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `datasourceInstance`                                     | `string`                                                 | :heavy_check_mark:                                       | Datasource instance that should process the submission   |
| `type_`                                                  | `string`                                                 | :heavy_check_mark:                                       | Submission type registered for the datasource            |
| `requestBody`                                            | map[string]`any`                                         | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PostRestAPIIndexSubmissionsDatasourceInstanceTypeResponse](../../models/operations/postrestapiindexsubmissionsdatasourceinstancetyperesponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| apierrors.ErrorInfoResponse | 400, 401, 404               | application/json            |
| apierrors.ErrorInfoResponse | 500                         | application/json            |
| apierrors.APIError          | 4XX, 5XX                    | \*/\*                       |