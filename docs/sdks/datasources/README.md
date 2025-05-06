# Datasources
(*Indexing.Datasources*)

## Overview

### Available Operations

* [Add](#add) - Add or update datasource
* [RetrieveConfig](#retrieveconfig) - Get datasource config

## Add

Add or update a custom datasource and its schema.

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

    res, err := s.Indexing.Datasources.Add(ctx, components.CustomDatasourceConfig{
        Name: "<value>",
        URLRegex: apiclientgo.String("https://example-company.datasource.com/.*"),
        Quicklinks: []components.Quicklink{
            components.Quicklink{
                IconConfig: &components.IconConfig{
                    Color: apiclientgo.String("#343CED"),
                    Key: apiclientgo.String("person_icon"),
                    IconType: components.IconTypeGlyph.ToPointer(),
                    Name: apiclientgo.String("user"),
                },
            },
            components.Quicklink{
                IconConfig: &components.IconConfig{
                    Color: apiclientgo.String("#343CED"),
                    Key: apiclientgo.String("person_icon"),
                    IconType: components.IconTypeGlyph.ToPointer(),
                    Name: apiclientgo.String("user"),
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