# User
(*Client.User*)

## Overview

### Available Operations

* [AddCredential](#addcredential) - Create credentials
* [DeleteQueryHistory](#deletequeryhistory) - Delete query history
* [Invite](#invite) - Send invitation
* [GetPublicConfig](#getpublicconfig) - Read public client configuration
* [RemoveCredential](#removecredential) - Delete credentials
* [SendSupportEmail](#sendsupportemail) - Send support email

## AddCredential

API to save a user's credentials. Used for Confluence restricted pages and Tableau per-user auth, for example

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

    res, err := s.Client.User.AddCredential(ctx, components.AddCredentialRequest{}, nil, nil)
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
| `addCredentialRequest`                                                                                                   | [components.AddCredentialRequest](../../models/components/addcredentialrequest.md)                                       | :heavy_check_mark:                                                                                                       | Credential content                                                                                                       |
| `xGleanActAs`                                                                                                            | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.AddcredentialResponse](../../models/operations/addcredentialresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteQueryHistory

Remove one or more queries from a user's query history.

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

    res, err := s.Client.User.DeleteQueryHistory(ctx, components.DeleteQueryHistoryRequest{}, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteQueryHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `deleteQueryHistoryRequest`                                                                                              | [components.DeleteQueryHistoryRequest](../../models/components/deletequeryhistoryrequest.md)                             | :heavy_check_mark:                                                                                                       | Delete query history request                                                                                             |
| `xGleanActAs`                                                                                                            | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.DeletequeryhistoryResponse](../../models/operations/deletequeryhistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Invite

Invites people to Glean via email or Slack

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

    res, err := s.Client.User.Invite(ctx, components.InviteRequest{
        Recipients: []components.Person{
            components.Person{
                Name: "George Clooney",
                ObfuscatedID: "abc123",
            },
        },
        RecipientFilters: &components.PeopleFilters{
            Filter: []components.FacetFilter{
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
| `inviteRequest`                                                                                                          | [components.InviteRequest](../../models/components/inviterequest.md)                                                     | :heavy_check_mark:                                                                                                       | Invite request                                                                                                           |
| `xGleanActAs`                                                                                                            | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.InviteResponse](../../models/operations/inviteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetPublicConfig

Read configuration information such as the company name, logo and etc that is public and is not considered as PII.

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

    res, err := s.Client.User.GetPublicConfig(ctx, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `xGleanActAs`                                                                                                            | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `publicConfigRequest`                                                                                                    | [*components.PublicConfigRequest](../../models/components/publicconfigrequest.md)                                        | :heavy_minus_sign:                                                                                                       | Public Config request                                                                                                    |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.PublicconfigResponse](../../models/operations/publicconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RemoveCredential

Delete a user's credentials. Used for Confluence restricted pages and Tableau per-user auth, for example

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

    res, err := s.Client.User.RemoveCredential(ctx, components.RemoveCredentialRequest{}, nil, nil)
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
| `removeCredentialRequest`                                                                                                | [components.RemoveCredentialRequest](../../models/components/removecredentialrequest.md)                                 | :heavy_check_mark:                                                                                                       | Credential content                                                                                                       |
| `xGleanActAs`                                                                                                            | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.RemovecredentialResponse](../../models/operations/removecredentialresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SendSupportEmail

Sends a support email based on a template to the Glean support team.

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

    res, err := s.Client.User.SendSupportEmail(ctx, components.EmailRequest{
        EmailTemplate: components.CommunicationTemplateOnboardingTips,
        Recipients: []components.Person{
            components.Person{
                Name: "George Clooney",
                ObfuscatedID: "abc123",
            },
        },
        CcRecipients: []components.Person{
            components.Person{
                Name: "George Clooney",
                ObfuscatedID: "abc123",
            },
        },
        RecipientFilters: &components.PeopleFilters{
            Filter: []components.FacetFilter{
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
        Senders: []components.Person{
            components.Person{
                Name: "George Clooney",
                ObfuscatedID: "abc123",
            },
        },
        Documents: []components.Document{
            components.Document{
                Metadata: &components.DocumentMetadata{
                    Datasource: apiclientgo.String("datasource"),
                    ObjectType: apiclientgo.String("Feature Request"),
                    Container: apiclientgo.String("container"),
                    ParentID: apiclientgo.String("JIRA_EN-1337"),
                    MimeType: apiclientgo.String("mimeType"),
                    DocumentID: apiclientgo.String("documentId"),
                    CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                    UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                    Author: &components.Person{
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
                    Owner: &components.Person{
                        Name: "George Clooney",
                        ObfuscatedID: "abc123",
                    },
                    MentionedPeople: []components.Person{},
                    Components: []string{
                        "Backend",
                        "Networking",
                    },
                    Status: apiclientgo.String("[\"Done\"]"),
                    Pins: []components.PinDocument{},
                    AssignedTo: &components.Person{
                        Name: "George Clooney",
                        ObfuscatedID: "abc123",
                    },
                    UpdatedBy: &components.Person{
                        Name: "George Clooney",
                        ObfuscatedID: "abc123",
                    },
                    Collections: []components.Collection{},
                    Interactions: &components.DocumentInteractions{
                        Reacts: []components.Reaction{},
                        Shares: []components.Share{},
                    },
                    Verification: &components.Verification{
                        State: components.StateDeprecated,
                        Metadata: &components.VerificationMetadata{
                            LastVerifier: &components.Person{
                                Name: "George Clooney",
                                ObfuscatedID: "abc123",
                            },
                            Reminders: []components.Reminder{},
                            LastReminder: &components.Reminder{
                                Assignee: components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                },
                                Requestor: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                },
                                RemindAt: 631645,
                            },
                            CandidateVerifiers: []components.Person{},
                        },
                    },
                    CustomData: map[string]components.CustomDataValue{
                        "someCustomField": components.CustomDataValue{},
                    },
                    ContactPerson: &components.Person{
                        Name: "George Clooney",
                        ObfuscatedID: "abc123",
                    },
                },
            },
            components.Document{},
        },
        FeedbackPayload: &components.FeedbackPayload{
            CustomJSON: apiclientgo.String("{\"comment\": \"glean is awesome!\", \"sender\": \"happycustomer@customer.com\"}"),
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
| `emailRequest`                                                                                                           | [components.EmailRequest](../../models/components/emailrequest.md)                                                       | :heavy_check_mark:                                                                                                       | Support request                                                                                                          |
| `xGleanActAs`                                                                                                            | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.SupportEmailResponse](../../models/operations/supportemailresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |