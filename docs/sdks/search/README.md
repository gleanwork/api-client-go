# Search
(*Client.Search*)

## Overview

### Available Operations

* [QueryAsAdmin](#queryasadmin) - Search the index (admin)
* [Autocomplete](#autocomplete) - Autocomplete
* [RetrieveFeed](#retrievefeed) - Feed of documents and events
* [Recommendations](#recommendations) - Recommend documents
* [Query](#query) - Search

## QueryAsAdmin

Retrieves results for search query without respect for permissions. This is available only to privileged users.

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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Client.Search.QueryAsAdmin(ctx, &components.SearchRequest{
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
                    State: components.StateVerified,
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
                            RemindAt: 333878,
                        },
                        CandidateVerifiers: []components.Person{},
                    },
                },
                Shortcuts: []components.Shortcut{},
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
            FacetBucketSize: 226218,
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
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [components.SearchRequest](../../models/components/searchrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.AdminsearchResponse](../../models/operations/adminsearchresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| apierrors.GleanDataError | 403, 422                 | application/json         |
| apierrors.APIError       | 4XX, 5XX                 | \*/\*                    |

## Autocomplete

Retrieve query suggestions, operators and documents for the given partially typed query.

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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Client.Search.Autocomplete(ctx, components.AutocompleteRequest{
        TrackingToken: apiclientgo.String("trackingToken"),
        Query: apiclientgo.String("San Fra"),
        Datasource: apiclientgo.String("GDRIVE"),
        ResultSize: apiclientgo.Int64(10),
        AuthTokens: []components.AuthToken{
            components.AuthToken{
                AccessToken: "123abc",
                Datasource: "gmail",
                Scope: apiclientgo.String("email profile https://www.googleapis.com/auth/gmail.readonly"),
                TokenType: apiclientgo.String("Bearer"),
                AuthUser: apiclientgo.String("1"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AutocompleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.AutocompleteRequest](../../models/components/autocompleterequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.AutocompleteResponse](../../models/operations/autocompleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RetrieveFeed

The personalized feed/home includes different types of contents including suggestions, recents, calendar events and many more.

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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Client.Search.RetrieveFeed(ctx, components.FeedRequest{
        TimeoutMillis: apiclientgo.Int64(5000),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FeedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [components.FeedRequest](../../models/components/feedrequest.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.FeedResponse](../../models/operations/feedresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Recommendations

Retrieve recommended documents for the given URL or Glean Document ID.

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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Client.Search.Recommendations(ctx, &components.RecommendationsRequest{
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
                    State: components.StateVerified,
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
                            RemindAt: 986764,
                        },
                        CandidateVerifiers: []components.Person{},
                    },
                },
                Shortcuts: []components.Shortcut{},
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
        RequestOptions: &components.RecommendationsRequestOptions{
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
            Context: &components.Document{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResultsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.RecommendationsRequest](../../models/components/recommendationsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.RecommendationsResponse](../../models/operations/recommendationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Query

Retrieve results from the index for the given query and filters.

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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Client.Search.Query(ctx, &components.SearchRequest{
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
                            RemindAt: 675445,
                        },
                        CandidateVerifiers: []components.Person{},
                    },
                },
                Shortcuts: []components.Shortcut{},
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
            FacetBucketSize: 653619,
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
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [components.SearchRequest](../../models/components/searchrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.SearchResponse](../../models/operations/searchresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| apierrors.GleanDataError | 403, 422                 | application/json         |
| apierrors.APIError       | 4XX, 5XX                 | \*/\*                    |