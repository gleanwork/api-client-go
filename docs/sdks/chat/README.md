# Chat
(*Client.Chat*)

## Overview

### Available Operations

* [Ask](#ask) - Detect and answer questions
* [Start](#start) - Chat
* [DeleteAll](#deleteall) - Deletes all saved Chats owned by a user
* [Delete](#delete) - Deletes saved Chats
* [Get](#get) - Retrieves a Chat
* [List](#list) - Retrieves all saved Chats
* [GetApplication](#getapplication) - Gets the metadata for a custom Chat application
* [UploadFiles](#uploadfiles) - Upload files for Chat.
* [GetFiles](#getfiles) - Get files uploaded by a user for Chat.
* [DeleteFiles](#deletefiles) - Delete files uploaded by a user for chat.

## Ask

Classify a query as information seeking or not. If so, generate an AI answer and/or provide relevant documents. Useful for integrating into existing chat interfaces.

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

    res, err := s.Client.Chat.Ask(ctx, nil, nil, &components.AskRequest{
        SearchRequest: components.SearchRequest{
            TrackingToken: apiclientgo.String("trackingToken"),
            SourceDocument: &components.Document{
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
                        State: components.StateUnverified,
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
                                RemindAt: 892381,
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
            PageSize: apiclientgo.Int64(100),
            MaxSnippetSize: apiclientgo.Int64(400),
            Query: "vacation policy",
            InputDetails: &components.SearchRequestInputDetails{
                HasCopyPaste: apiclientgo.Bool(true),
            },
            RequestOptions: &components.SearchRequestOptions{
                DatasourceFilter: apiclientgo.String("JIRA"),
                DatasourcesFilter: []string{
                    "JIRA",
                },
                QueryOverridesFacetFilters: apiclientgo.Bool(true),
                FacetFilters: []components.FacetFilter{
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
                FacetFilterSets: []components.FacetFilterSet{
                    components.FacetFilterSet{
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
                    },
                },
                FacetBucketSize: 132988,
                AuthTokens: []components.AuthToken{
                    components.AuthToken{
                        AccessToken: "123abc",
                        Datasource: "gmail",
                        Scope: apiclientgo.String("email profile https://www.googleapis.com/auth/gmail.readonly"),
                        TokenType: apiclientgo.String("Bearer"),
                        AuthUser: apiclientgo.String("1"),
                    },
                },
            },
            TimeoutMillis: apiclientgo.Int64(5000),
            People: []components.Person{
                components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AskResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `askRequest`                                                                                                             | [*components.AskRequest](../../models/components/askrequest.md)                                                          | :heavy_minus_sign:                                                                                                       | Ask request                                                                                                              |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.AskResponse](../../models/operations/askresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| apierrors.ErrorInfo | 403, 422            | application/json    |
| apierrors.APIError  | 4XX, 5XX            | \*/\*               |

## Start

Have a conversation with Glean AI.

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

    res, err := s.Client.Chat.Start(ctx, components.ChatRequest{
        Messages: []components.ChatMessage{
            components.ChatMessage{
                Fragments: []components.ChatMessageFragment{
                    components.ChatMessageFragment{
                        Text: apiclientgo.String("What are the company holidays this year?"),
                    },
                },
            },
        },
    }, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ChatResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `chatRequest`                                                                                                            | [components.ChatRequest](../../models/components/chatrequest.md)                                                         | :heavy_check_mark:                                                                                                       | Includes chat history for Glean AI to respond to.                                                                        |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `timezoneOffset`                                                                                                         | **int64*                                                                                                                 | :heavy_minus_sign:                                                                                                       | The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC.               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.ChatResponse](../../models/operations/chatresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteAll

Deletes all saved Chats a user has had and all their contained conversational content.

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

    res, err := s.Client.Chat.DeleteAll(ctx, nil, nil, nil)
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
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `timezoneOffset`                                                                                                         | **int64*                                                                                                                 | :heavy_minus_sign:                                                                                                       | The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC.               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.DeleteallchatsResponse](../../models/operations/deleteallchatsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Deletes saved Chats and all their contained conversational content.

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

    res, err := s.Client.Chat.Delete(ctx, components.DeleteChatsRequest{
        Ids: []string{
            "<value>",
            "<value>",
        },
    }, nil, nil, nil)
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
| `deleteChatsRequest`                                                                                                     | [components.DeleteChatsRequest](../../models/components/deletechatsrequest.md)                                           | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `timezoneOffset`                                                                                                         | **int64*                                                                                                                 | :heavy_minus_sign:                                                                                                       | The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC.               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.DeletechatsResponse](../../models/operations/deletechatsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieves the chat history between Glean Assistant and the user for a given Chat.

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

    res, err := s.Client.Chat.Get(ctx, components.GetChatRequest{
        ID: "<id>",
    }, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetChatResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `getChatRequest`                                                                                                         | [components.GetChatRequest](../../models/components/getchatrequest.md)                                                   | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `timezoneOffset`                                                                                                         | **int64*                                                                                                                 | :heavy_minus_sign:                                                                                                       | The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC.               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.GetchatResponse](../../models/operations/getchatresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Retrieves all the saved Chats between Glean Assistant and the user. The returned Chats contain only metadata and no conversational content.

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

    res, err := s.Client.Chat.List(ctx, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListChatsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `timezoneOffset`                                                                                                         | **int64*                                                                                                                 | :heavy_minus_sign:                                                                                                       | The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC.               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.ListchatsResponse](../../models/operations/listchatsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetApplication

Gets the Chat application details for the specified application ID.

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

    res, err := s.Client.Chat.GetApplication(ctx, components.GetChatApplicationRequest{
        ID: "<id>",
    }, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetChatApplicationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `getChatApplicationRequest`                                                                                              | [components.GetChatApplicationRequest](../../models/components/getchatapplicationrequest.md)                             | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `timezoneOffset`                                                                                                         | **int64*                                                                                                                 | :heavy_minus_sign:                                                                                                       | The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC.               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.GetchatapplicationResponse](../../models/operations/getchatapplicationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UploadFiles

Upload files for Chat.

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

    res, err := s.Client.Chat.UploadFiles(ctx, components.UploadChatFilesRequest{
        Files: []components.File{},
    }, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UploadChatFilesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `uploadChatFilesRequest`                                                                                                 | [components.UploadChatFilesRequest](../../models/components/uploadchatfilesrequest.md)                                   | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `timezoneOffset`                                                                                                         | **int64*                                                                                                                 | :heavy_minus_sign:                                                                                                       | The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC.               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UploadchatfilesResponse](../../models/operations/uploadchatfilesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetFiles

Get files uploaded by a user for Chat.

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

    res, err := s.Client.Chat.GetFiles(ctx, components.GetChatFilesRequest{
        FileIds: []string{
            "<value>",
            "<value>",
        },
    }, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetChatFilesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `getChatFilesRequest`                                                                                                    | [components.GetChatFilesRequest](../../models/components/getchatfilesrequest.md)                                         | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `timezoneOffset`                                                                                                         | **int64*                                                                                                                 | :heavy_minus_sign:                                                                                                       | The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC.               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.GetchatfilesResponse](../../models/operations/getchatfilesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteFiles

Delete files uploaded by a user for Chat.

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

    res, err := s.Client.Chat.DeleteFiles(ctx, components.DeleteChatFilesRequest{
        FileIds: []string{
            "<value>",
        },
    }, nil, nil, nil)
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
| `deleteChatFilesRequest`                                                                                                 | [components.DeleteChatFilesRequest](../../models/components/deletechatfilesrequest.md)                                   | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `xScioActas`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). |
| `xGleanAuthType`                                                                                                         | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | Auth type being used to access the endpoint (should be non-empty only for global tokens).                                |
| `timezoneOffset`                                                                                                         | **int64*                                                                                                                 | :heavy_minus_sign:                                                                                                       | The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC.               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.DeletechatfilesResponse](../../models/operations/deletechatfilesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |