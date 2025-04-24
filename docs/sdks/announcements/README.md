# Announcements
(*Client.Announcements*)

## Overview

### Available Operations

* [Create](#create) - Create Announcement
* [CreateDraft](#createdraft) - Create draft Announcement
* [Delete](#delete) - Delete Announcement
* [DeleteDraft](#deletedraft) - Delete draft Announcement
* [Get](#get) - Read Announcement
* [GetDraft](#getdraft) - Read draft Announcement
* [List](#list) - List Announcements
* [Preview](#preview) - Preview Announcement
* [PreviewDraft](#previewdraft) - Preview draft Announcement
* [Publish](#publish) - Publish draft Announcement
* [Unpublish](#unpublish) - Unpublish Announcement
* [Update](#update) - Update Announcement
* [UpdateDraft](#updatedraft) - Update draft Announcement

## Create

Create a textual announcement visible to some set of users based on department and location.

### Example Usage

```go
package main

import(
	"context"
	"os"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/types"
	"github.com/gleanwork/api-client-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
    )

    res, err := s.Client.Announcements.Create(ctx, components.CreateAnnouncementRequest{
        StartTime: types.MustTimeFromString("2024-06-17T07:14:55.338Z"),
        EndTime: types.MustTimeFromString("2024-11-30T17:06:07.804Z"),
        Title: "<value>",
        Body: &components.StructuredText{
            Text: "From https://en.wikipedia.org/wiki/Diffuse_sky_radiation, the sky is blue because blue light is more strongly scattered than longer-wavelength light.",
            StructuredList: []components.StructuredTextItem{},
        },
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
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Announcement != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `createAnnouncementRequest`                                                                                              | [components.CreateAnnouncementRequest](../../models/components/createannouncementrequest.md)                             | :heavy_check_mark:                                                                                                       | Announcement content                                                                                                     |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.CreateannouncementResponse](../../models/operations/createannouncementresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateDraft

Create a draft of a textual announcement visible to some set of users based on department and location.

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

    res, err := s.Client.Announcements.CreateDraft(ctx, components.CreateDraftAnnouncementRequest{
        Body: &components.StructuredText{
            Text: "From https://en.wikipedia.org/wiki/Diffuse_sky_radiation, the sky is blue because blue light is more strongly scattered than longer-wavelength light.",
            StructuredList: []components.StructuredTextItem{},
        },
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
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Announcement != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `createDraftAnnouncementRequest`                                                                                         | [components.CreateDraftAnnouncementRequest](../../models/components/createdraftannouncementrequest.md)                   | :heavy_check_mark:                                                                                                       | Draft announcement content                                                                                               |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.CreatedraftannouncementResponse](../../models/operations/createdraftannouncementresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete an existing user-generated announcement.

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

    res, err := s.Client.Announcements.Delete(ctx, components.DeleteAnnouncementRequest{
        ID: 545907,
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
| `deleteAnnouncementRequest`                                                                                              | [components.DeleteAnnouncementRequest](../../models/components/deleteannouncementrequest.md)                             | :heavy_check_mark:                                                                                                       | Delete announcement request                                                                                              |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.DeleteannouncementResponse](../../models/operations/deleteannouncementresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteDraft

Delete an existing user-generated draft Announcement.

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

    res, err := s.Client.Announcements.DeleteDraft(ctx, components.DeleteAnnouncementRequest{
        ID: 171443,
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
| `deleteAnnouncementRequest`                                                                                              | [components.DeleteAnnouncementRequest](../../models/components/deleteannouncementrequest.md)                             | :heavy_check_mark:                                                                                                       | Delete draft announcement request                                                                                        |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.DeletedraftannouncementResponse](../../models/operations/deletedraftannouncementresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Read the details of an Announcement given its ID.

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

    res, err := s.Client.Announcements.Get(ctx, components.GetAnnouncementRequest{
        ID: 700347,
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAnnouncementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `getAnnouncementRequest`                                                                                                 | [components.GetAnnouncementRequest](../../models/components/getannouncementrequest.md)                                   | :heavy_check_mark:                                                                                                       | GetAnnouncement request                                                                                                  |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.GetannouncementResponse](../../models/operations/getannouncementresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetDraft

Read the details of an existing user-generated draft Announcement.

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

    res, err := s.Client.Announcements.GetDraft(ctx, components.GetAnnouncementRequest{
        ID: 476509,
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetDraftAnnouncementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `getAnnouncementRequest`                                                                                                 | [components.GetAnnouncementRequest](../../models/components/getannouncementrequest.md)                                   | :heavy_check_mark:                                                                                                       | Get draft announcement request                                                                                           |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.GetdraftannouncementResponse](../../models/operations/getdraftannouncementresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

List Announcement details for all Announcements matching the given criteria.

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

    res, err := s.Client.Announcements.List(ctx, components.ListAnnouncementsRequest{}, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAnnouncementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `listAnnouncementsRequest`                                                                                               | [components.ListAnnouncementsRequest](../../models/components/listannouncementsrequest.md)                               | :heavy_check_mark:                                                                                                       | Includes request params for querying announcements.                                                                      |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.ListannouncementsResponse](../../models/operations/listannouncementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Preview

Generate a preview for a user-generated Announcement from structured text.

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

    res, err := s.Client.Announcements.Preview(ctx, components.StructuredTextMutableProperties{
        Text: "From https://en.wikipedia.org/wiki/Diffuse_sky_radiation, the sky is blue because blue light is more strongly scattered than longer-wavelength light.",
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PreviewStructuredTextResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `structuredTextMutableProperties`                                                                                        | [components.StructuredTextMutableProperties](../../models/components/structuredtextmutableproperties.md)                 | :heavy_check_mark:                                                                                                       | preview structured text request                                                                                          |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.PreviewannouncementResponse](../../models/operations/previewannouncementresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PreviewDraft

Generates a preview for a user-generated Announcement from a draft.

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

    res, err := s.Client.Announcements.PreviewDraft(ctx, components.PreviewUgcRequest{
        Draft: &components.UgcDraft{
            Announcement: &components.AnnouncementMutableProperties{
                Body: &components.StructuredText{
                    Text: "From https://en.wikipedia.org/wiki/Diffuse_sky_radiation, the sky is blue because blue light is more strongly scattered than longer-wavelength light.",
                    StructuredList: []components.StructuredTextItem{},
                },
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
            Answer: &components.AnswerMutableProperties{
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
                                ManagementChain: []components.Person{},
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                Reports: []components.Person{},
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
                        Role: components.UserRoleAnswerModerator,
                    },
                },
                RemovedRoles: []components.UserRoleSpecification{
                    components.UserRoleSpecification{
                        Role: components.UserRoleOwner,
                    },
                    components.UserRoleSpecification{
                        Role: components.UserRoleViewer,
                    },
                    components.UserRoleSpecification{
                        Role: components.UserRoleOwner,
                    },
                },
                Roles: []components.UserRoleSpecification{
                    components.UserRoleSpecification{
                        Role: components.UserRoleViewer,
                    },
                    components.UserRoleSpecification{
                        Role: components.UserRoleViewer,
                    },
                    components.UserRoleSpecification{
                        Role: components.UserRoleOwner,
                    },
                },
            },
        },
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PreviewUgcResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `previewUgcRequest`                                                                                                      | [components.PreviewUgcRequest](../../models/components/previewugcrequest.md)                                             | :heavy_check_mark:                                                                                                       | preview announcement request                                                                                             |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.PreviewannouncementdraftResponse](../../models/operations/previewannouncementdraftresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Publish

Promote a draft Announcement to be visible to others.

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

    res, err := s.Client.Announcements.Publish(ctx, components.PublishDraftAnnouncementRequest{
        ID: 876134,
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
| `publishDraftAnnouncementRequest`                                                                                        | [components.PublishDraftAnnouncementRequest](../../models/components/publishdraftannouncementrequest.md)                 | :heavy_check_mark:                                                                                                       | Publish draft announcement content.                                                                                      |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.PublishdraftannouncementResponse](../../models/operations/publishdraftannouncementresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Unpublish

Unpublish an Announcement to hide it from users.

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

    res, err := s.Client.Announcements.Unpublish(ctx, components.UnpublishAnnouncementRequest{
        ID: 195182,
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Announcement != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `unpublishAnnouncementRequest`                                                                                           | [components.UnpublishAnnouncementRequest](../../models/components/unpublishannouncementrequest.md)                       | :heavy_check_mark:                                                                                                       | Unpublish announcement content.                                                                                          |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UnpublishannouncementResponse](../../models/operations/unpublishannouncementresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update a textual announcement visible to some set of users based on department and location.

### Example Usage

```go
package main

import(
	"context"
	"os"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/types"
	"github.com/gleanwork/api-client-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
    )

    res, err := s.Client.Announcements.Update(ctx, components.UpdateAnnouncementRequest{
        StartTime: types.MustTimeFromString("2025-07-28T19:04:48.565Z"),
        EndTime: types.MustTimeFromString("2024-10-16T10:52:42.015Z"),
        Title: "<value>",
        Body: &components.StructuredText{
            Text: "From https://en.wikipedia.org/wiki/Diffuse_sky_radiation, the sky is blue because blue light is more strongly scattered than longer-wavelength light.",
            StructuredList: []components.StructuredTextItem{},
        },
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
        ID: 761625,
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Announcement != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `updateAnnouncementRequest`                                                                                              | [components.UpdateAnnouncementRequest](../../models/components/updateannouncementrequest.md)                             | :heavy_check_mark:                                                                                                       | Announcement content. Id need to be specified for the announcement.                                                      |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UpdateannouncementResponse](../../models/operations/updateannouncementresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateDraft

Update a textual Announcement visible to some set of users based on department and location.

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

    res, err := s.Client.Announcements.UpdateDraft(ctx, components.UpdateDraftAnnouncementRequest{
        Body: &components.StructuredText{
            Text: "From https://en.wikipedia.org/wiki/Diffuse_sky_radiation, the sky is blue because blue light is more strongly scattered than longer-wavelength light.",
            StructuredList: []components.StructuredTextItem{},
        },
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
        DraftID: 758103,
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Announcement != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `updateDraftAnnouncementRequest`                                                                                         | [components.UpdateDraftAnnouncementRequest](../../models/components/updatedraftannouncementrequest.md)                   | :heavy_check_mark:                                                                                                       | Draft announcement content. DraftId needs to be specified.                                                               |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UpdatedraftannouncementResponse](../../models/operations/updatedraftannouncementresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |