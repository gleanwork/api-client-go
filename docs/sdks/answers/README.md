# Answers
(*Client.Answers*)

## Overview

### Available Operations

* [Create](#create) - Create Answer
* [Delete](#delete) - Delete Answer
* [Edit](#edit) - Update Answer
* [Get](#get) - Read Answer
* [List](#list) - List Answers

## Create

Create a user-generated Answer that contains a question and answer.

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

    res, err := s.Client.Answers.Create(ctx, components.CreateAnswerRequest{
        Data: components.AnswerCreationData{
            Question: apiclientgo.String("Why is the sky blue?"),
            BodyText: apiclientgo.String("From https://en.wikipedia.org/wiki/Diffuse_sky_radiation, the sky is blue because blue light is more strongly scattered than longer-wavelength light."),
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
            Roles: []components.UserRoleSpecification{
                components.UserRoleSpecification{
                    Role: components.UserRoleAnswerModerator,
                },
                components.UserRoleSpecification{
                    Role: components.UserRoleOwner,
                },
                components.UserRoleSpecification{
                    Role: components.UserRoleVerifier,
                },
            },
            CombinedAnswerText: &components.StructuredTextMutableProperties{
                Text: "From https://en.wikipedia.org/wiki/Diffuse_sky_radiation, the sky is blue because blue light is more strongly scattered than longer-wavelength light.",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Answer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.CreateAnswerRequest](../../models/components/createanswerrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.CreateanswerResponse](../../models/operations/createanswerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete an existing user-generated Answer.

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

    res, err := s.Client.Answers.Delete(ctx, components.DeleteAnswerRequest{
        ID: 3,
        DocID: apiclientgo.String("ANSWERS_answer_3"),
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.DeleteAnswerRequest](../../models/components/deleteanswerrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.DeleteanswerResponse](../../models/operations/deleteanswerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Edit

Update an existing user-generated Answer.

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

    res, err := s.Client.Answers.Edit(ctx, components.EditAnswerRequest{
        ID: 3,
        DocID: apiclientgo.String("ANSWERS_answer_3"),
        Question: apiclientgo.String("Why is the sky blue?"),
        BodyText: apiclientgo.String("From https://en.wikipedia.org/wiki/Diffuse_sky_radiation, the sky is blue because blue light is more strongly scattered than longer-wavelength light."),
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
        },
        RemovedRoles: []components.UserRoleSpecification{
            components.UserRoleSpecification{
                Role: components.UserRoleEditor,
            },
        },
        Roles: []components.UserRoleSpecification{
            components.UserRoleSpecification{
                Role: components.UserRoleAnswerModerator,
            },
            components.UserRoleSpecification{
                Role: components.UserRoleOwner,
            },
        },
        CombinedAnswerText: &components.StructuredTextMutableProperties{
            Text: "From https://en.wikipedia.org/wiki/Diffuse_sky_radiation, the sky is blue because blue light is more strongly scattered than longer-wavelength light.",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Answer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.EditAnswerRequest](../../models/components/editanswerrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.EditanswerResponse](../../models/operations/editanswerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Read the details of a particular Answer given its ID.

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

    res, err := s.Client.Answers.Get(ctx, components.GetAnswerRequest{
        ID: apiclientgo.Int64(3),
        DocID: apiclientgo.String("ANSWERS_answer_3"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAnswerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.GetAnswerRequest](../../models/components/getanswerrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.GetanswerResponse](../../models/operations/getanswerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

List Answers created by the current user.

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

    res, err := s.Client.Answers.List(ctx, components.ListAnswersRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAnswersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.ListAnswersRequest](../../models/components/listanswersrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ListanswersResponse](../../models/operations/listanswersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |