# DisplayableLists
(*Client.DisplayableLists*)

## Overview

### Available Operations

* [Create](#create) - Create displayable lists
* [Delete](#delete) - Delete displayable lists
* [Get](#get) - Read displayable lists
* [Update](#update) - Update displayable lists

## Create

Create one or more lists that can be display on the home page.

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

    res, err := s.Client.DisplayableLists.Create(ctx, components.CreateDisplayableListsRequest{
        Items: []components.DisplayableList{
            components.DisplayableList{
                Config: &components.DisplayableListConfig{
                    AudienceFilters: []components.FacetFilter{
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
        },
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDisplayableListsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `createDisplayableListsRequest`                                                                                          | [components.CreateDisplayableListsRequest](../../models/components/createdisplayablelistsrequest.md)                     | :heavy_check_mark:                                                                                                       | Create new displayable lists                                                                                             |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.CreatedisplayablelistsResponse](../../models/operations/createdisplayablelistsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete one or more displayable lists.

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

    res, err := s.Client.DisplayableLists.Delete(ctx, components.DeleteDisplayableListsRequest{
        Ids: []int{
            698486,
            386564,
        },
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `deleteDisplayableListsRequest`                                                                                          | [components.DeleteDisplayableListsRequest](../../models/components/deletedisplayablelistsrequest.md)                     | :heavy_check_mark:                                                                                                       | Updated version of the displayable list configs.                                                                         |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.DeletedisplayablelistsResponse](../../models/operations/deletedisplayablelistsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Read the details of the displayable lists with the given IDs.

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

    res, err := s.Client.DisplayableLists.Get(ctx, components.GetDisplayableListsRequest{
        Ids: []int{
            558834,
            544221,
        },
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetDisplayableListsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `getDisplayableListsRequest`                                                                                             | [components.GetDisplayableListsRequest](../../models/components/getdisplayablelistsrequest.md)                           | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.GetdisplayablelistsResponse](../../models/operations/getdisplayablelistsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update one or more displayable lists with all fields from request fields.

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

    res, err := s.Client.DisplayableLists.Update(ctx, components.UpdateDisplayableListsRequest{
        Items: []components.DisplayableList{
            components.DisplayableList{
                Config: &components.DisplayableListConfig{
                    AudienceFilters: []components.FacetFilter{
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
            components.DisplayableList{
                Config: &components.DisplayableListConfig{
                    AudienceFilters: []components.FacetFilter{
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
            components.DisplayableList{
                Config: &components.DisplayableListConfig{
                    AudienceFilters: []components.FacetFilter{
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
        },
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateDisplayableListsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `updateDisplayableListsRequest`                                                                                          | [components.UpdateDisplayableListsRequest](../../models/components/updatedisplayablelistsrequest.md)                     | :heavy_check_mark:                                                                                                       | Updated version of the displayable list configs.                                                                         |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UpdatedisplayablelistsResponse](../../models/operations/updatedisplayablelistsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |