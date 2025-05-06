# ClientShortcuts
(*Client.Shortcuts*)

## Overview

### Available Operations

* [Create](#create) - Create shortcut
* [Delete](#delete) - Delete shortcut
* [Retrieve](#retrieve) - Read shortcut
* [List](#list) - List shortcuts
* [Update](#update) - Update shortcut

## Create

Create a user-generated shortcut that contains an alias and destination URL.

### Example Usage

```go
package main

import(
	"context"
	"os"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/api-client-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
    )

    res, err := s.Client.Shortcuts.Create(ctx, components.CreateShortcutRequest{
        Data: components.ShortcutMutableProperties{
            AddedRoles: []components.UserRoleSpecification{
                components.UserRoleSpecification{
                    Person: &components.Person{
                        Name: "George Clooney",
                        ObfuscatedID: "abc123",
                        RelatedDocuments: []components.RelatedDocuments{},
                        Metadata: &components.PersonMetadata{
                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                            Title: apiclientgo.String("Actor"),
                            Department: apiclientgo.String("Movies"),
                            Email: apiclientgo.String("george@example.com"),
                            Location: apiclientgo.String("Hollywood, CA"),
                            Phone: apiclientgo.String("6505551234"),
                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                            StartDate: types.MustNewDateFromString("2000-01-23"),
                            DatasourceProfile: []components.DatasourceProfile{
                                components.DatasourceProfile{
                                    Datasource: "github",
                                    Handle: "<value>",
                                },
                            },
                            QuerySuggestions: &components.QuerySuggestionList{
                                Suggestions: []components.QuerySuggestion{},
                            },
                            InviteInfo: &components.InviteInfo{
                                Invites: []components.ChannelInviteInfo{},
                            },
                            CustomFields: []components.CustomFieldData{},
                            Badges: []components.Badge{
                                components.Badge{
                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                    DisplayName: apiclientgo.String("New hire"),
                                    IconConfig: &components.IconConfig{
                                        Color: apiclientgo.String("#343CED"),
                                        Key: apiclientgo.String("person_icon"),
                                        IconType: components.IconTypeGlyph.ToPointer(),
                                        Name: apiclientgo.String("user"),
                                    },
                                },
                            },
                        },
                    },
                    Role: components.UserRoleOwner,
                },
                components.UserRoleSpecification{
                    Role: components.UserRoleVerifier,
                },
            },
            RemovedRoles: []components.UserRoleSpecification{
                components.UserRoleSpecification{
                    Role: components.UserRoleVerifier,
                },
                components.UserRoleSpecification{
                    Role: components.UserRoleAnswerModerator,
                },
                components.UserRoleSpecification{
                    Role: components.UserRoleOwner,
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateShortcutResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.CreateShortcutRequest](../../models/components/createshortcutrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateshortcutResponse](../../models/operations/createshortcutresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete an existing user-generated shortcut.

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

    res, err := s.Client.Shortcuts.Delete(ctx, components.DeleteShortcutRequest{
        ID: 545907,
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.DeleteShortcutRequest](../../models/components/deleteshortcutrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.DeleteshortcutResponse](../../models/operations/deleteshortcutresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Retrieve

Read a particular shortcut's details given its ID.

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

    res, err := s.Client.Shortcuts.Retrieve(ctx, components.CreateGetShortcutRequestUnionGetShortcutRequest(
        components.GetShortcutRequest{
            Alias: "<value>",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.GetShortcutResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.GetShortcutRequestUnion](../../models/components/getshortcutrequestunion.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetshortcutResponse](../../models/operations/getshortcutresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

List shortcuts editable/owned by the currently authenticated user.

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

    res, err := s.Client.Shortcuts.List(ctx, components.ListShortcutsPaginatedRequest{
        PageSize: 10,
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
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListShortcutsPaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [components.ListShortcutsPaginatedRequest](../../models/components/listshortcutspaginatedrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListshortcutsResponse](../../models/operations/listshortcutsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Updates the shortcut with the given ID.

### Example Usage

```go
package main

import(
	"context"
	"os"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/api-client-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
    )

    res, err := s.Client.Shortcuts.Update(ctx, components.UpdateShortcutRequest{
        ID: 857478,
        AddedRoles: []components.UserRoleSpecification{
            components.UserRoleSpecification{
                Person: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    RelatedDocuments: []components.RelatedDocuments{},
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.String("Actor"),
                        Department: apiclientgo.String("Movies"),
                        Email: apiclientgo.String("george@example.com"),
                        Location: apiclientgo.String("Hollywood, CA"),
                        Phone: apiclientgo.String("6505551234"),
                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                        StartDate: types.MustNewDateFromString("2000-01-23"),
                        DatasourceProfile: []components.DatasourceProfile{
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{},
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{},
                        },
                        CustomFields: []components.CustomFieldData{},
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.String("deployment_name_new_hire"),
                                DisplayName: apiclientgo.String("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.String("#343CED"),
                                    Key: apiclientgo.String("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.String("user"),
                                },
                            },
                        },
                    },
                },
                Role: components.UserRoleEditor,
            },
            components.UserRoleSpecification{
                Role: components.UserRoleAnswerModerator,
            },
        },
        RemovedRoles: []components.UserRoleSpecification{
            components.UserRoleSpecification{
                Role: components.UserRoleEditor,
            },
            components.UserRoleSpecification{
                Role: components.UserRoleAnswerModerator,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateShortcutResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.UpdateShortcutRequest](../../models/components/updateshortcutrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateshortcutResponse](../../models/operations/updateshortcutresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |