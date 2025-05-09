# Collections
(*Client.Collections*)

## Overview

### Available Operations

* [AddItems](#additems) - Add Collection item
* [Create](#create) - Create Collection
* [Delete](#delete) - Delete Collection
* [DeleteItem](#deleteitem) - Delete Collection item
* [Update](#update) - Update Collection
* [UpdateItem](#updateitem) - Update Collection item
* [Retrieve](#retrieve) - Read Collection
* [List](#list) - List Collections

## AddItems

Add items to a Collection.

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

    res, err := s.Client.Collections.AddItems(ctx, components.AddCollectionItemsRequest{
        CollectionID: 6460.15,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AddCollectionItemsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.AddCollectionItemsRequest](../../models/components/addcollectionitemsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.AddcollectionitemsResponse](../../models/operations/addcollectionitemsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a publicly visible (empty) Collection of documents.

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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Client.Collections.Create(ctx, components.CreateCollectionRequest{
        Name: "<value>",
        AddedRoles: []components.UserRoleSpecification{
            components.UserRoleSpecification{
                Person: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    RelatedDocuments: []components.RelatedDocuments{
                        components.RelatedDocuments{
                            QuerySuggestion: &components.QuerySuggestion{
                                Query: "app:github type:pull author:mortimer",
                                SearchProviderInfo: &components.SearchProviderInfo{
                                    Name: apiclientgo.String("Google"),
                                    SearchLinkURLTemplate: apiclientgo.String("https://www.google.com/search?q={query}&hl=en"),
                                },
                                Label: apiclientgo.String("Mortimer's PRs"),
                                Datasource: apiclientgo.String("github"),
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
                                    FacetBucketSize: 134365,
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
                                Ranges: []components.TextRange{
                                    components.TextRange{
                                        StartIndex: 796474,
                                        Document: &components.Document{
                                            Metadata: &components.DocumentMetadata{
                                                Datasource: apiclientgo.String("datasource"),
                                                ObjectType: apiclientgo.String("Feature Request"),
                                                Container: apiclientgo.String("container"),
                                                ParentID: apiclientgo.String("JIRA_EN-1337"),
                                                MimeType: apiclientgo.String("mimeType"),
                                                DocumentID: apiclientgo.String("documentId"),
                                                CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                                                UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                                                Components: []string{
                                                    "Backend",
                                                    "Networking",
                                                },
                                                Status: apiclientgo.String("[\"Done\"]"),
                                                Pins: []components.PinDocument{
                                                    components.PinDocument{
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
                                                        DocumentID: "<id>",
                                                    },
                                                    components.PinDocument{
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
                                                        DocumentID: "<id>",
                                                    },
                                                    components.PinDocument{
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
                                                        DocumentID: "<id>",
                                                    },
                                                },
                                                Collections: []components.Collection{
                                                    components.Collection{
                                                        Name: "<value>",
                                                        Description: "fumigate convection though zowie",
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
                                                        ID: 496323,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 782367,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeDocument,
                                                            },
                                                        },
                                                    },
                                                },
                                                Interactions: &components.DocumentInteractions{
                                                    Reacts: []components.Reaction{
                                                        components.Reaction{},
                                                        components.Reaction{},
                                                        components.Reaction{},
                                                    },
                                                    Shares: []components.Share{
                                                        components.Share{
                                                            NumDaysAgo: 219974,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 449221,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 427887,
                                                        },
                                                    },
                                                },
                                                Verification: &components.Verification{
                                                    State: components.StateVerified,
                                                    Metadata: &components.VerificationMetadata{
                                                        Reminders: []components.Reminder{
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 491427,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 490420,
                                                        },
                                                    },
                                                },
                                                Shortcuts: []components.Shortcut{
                                                    components.Shortcut{
                                                        InputAlias: "<value>",
                                                    },
                                                },
                                                CustomData: map[string]components.CustomDataValue{
                                                    "someCustomField": components.CustomDataValue{},
                                                },
                                            },
                                        },
                                    },
                                },
                                InputDetails: &components.SearchRequestInputDetails{
                                    HasCopyPaste: apiclientgo.Bool(true),
                                },
                            },
                            Results: []components.SearchResult{
                                components.SearchResult{
                                    Title: apiclientgo.String("title"),
                                    URL: "https://example.com/foo/bar",
                                    NativeAppURL: apiclientgo.String("slack://foo/bar"),
                                    Snippets: []components.SearchResultSnippet{
                                        components.SearchResultSnippet{
                                            Snippet: "snippet",
                                            MimeType: apiclientgo.String("mimeType"),
                                        },
                                    },
                                },
                            },
                        },
                        components.RelatedDocuments{
                            QuerySuggestion: &components.QuerySuggestion{
                                Query: "app:github type:pull author:mortimer",
                                SearchProviderInfo: &components.SearchProviderInfo{
                                    Name: apiclientgo.String("Google"),
                                    SearchLinkURLTemplate: apiclientgo.String("https://www.google.com/search?q={query}&hl=en"),
                                },
                                Label: apiclientgo.String("Mortimer's PRs"),
                                Datasource: apiclientgo.String("github"),
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
                                    FacetBucketSize: 45416,
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
                                InputDetails: &components.SearchRequestInputDetails{
                                    HasCopyPaste: apiclientgo.Bool(true),
                                },
                            },
                        },
                    },
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
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                            },
                        },
                        CustomFields: []components.CustomFieldData{
                            components.CustomFieldData{
                                Label: "<value>",
                                Values: []components.CustomFieldValue{
                                    components.CreateCustomFieldValueCustomFieldValueStr(
                                        components.CustomFieldValueStr{},
                                    ),
                                    components.CreateCustomFieldValueCustomFieldValueStr(
                                        components.CustomFieldValueStr{},
                                    ),
                                },
                            },
                            components.CustomFieldData{
                                Label: "<value>",
                                Values: []components.CustomFieldValue{},
                            },
                        },
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
                Person: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
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
                        QuerySuggestions: &components.QuerySuggestionList{},
                        InviteInfo: &components.InviteInfo{},
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
                Role: components.UserRoleVerifier,
            },
        },
        RemovedRoles: []components.UserRoleSpecification{
            components.UserRoleSpecification{
                Person: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
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
                        QuerySuggestions: &components.QuerySuggestionList{},
                        InviteInfo: &components.InviteInfo{},
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
                Role: components.UserRoleVerifier,
            },
            components.UserRoleSpecification{
                Person: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
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
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{},
                        InviteInfo: &components.InviteInfo{},
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
            components.UserRoleSpecification{
                Person: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
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
                        QuerySuggestions: &components.QuerySuggestionList{},
                        InviteInfo: &components.InviteInfo{},
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
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateCollectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.CreateCollectionRequest](../../models/components/createcollectionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreatecollectionResponse](../../models/operations/createcollectionresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| apierrors.CollectionError | 422                       | application/json          |
| apierrors.APIError        | 4XX, 5XX                  | \*/\*                     |

## Delete

Delete a Collection given the Collection's ID.

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

    res, err := s.Client.Collections.Delete(ctx, components.DeleteCollectionRequest{
        Ids: []int64{
            698486,
            386564,
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.DeleteCollectionRequest](../../models/components/deletecollectionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.DeletecollectionResponse](../../models/operations/deletecollectionresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| apierrors.CollectionError | 422                       | application/json          |
| apierrors.APIError        | 4XX, 5XX                  | \*/\*                     |

## DeleteItem

Delete a single item from a Collection.

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

    res, err := s.Client.Collections.DeleteItem(ctx, components.DeleteCollectionItemRequest{
        CollectionID: 1357.59,
        ItemID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteCollectionItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.DeleteCollectionItemRequest](../../models/components/deletecollectionitemrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.DeletecollectionitemResponse](../../models/operations/deletecollectionitemresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the properties of an existing Collection.

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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Client.Collections.Update(ctx, components.EditCollectionRequest{
        Name: "<value>",
        AddedRoles: []components.UserRoleSpecification{
            components.UserRoleSpecification{
                Person: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    RelatedDocuments: []components.RelatedDocuments{
                        components.RelatedDocuments{
                            QuerySuggestion: &components.QuerySuggestion{
                                Query: "app:github type:pull author:mortimer",
                                SearchProviderInfo: &components.SearchProviderInfo{
                                    Name: apiclientgo.String("Google"),
                                    SearchLinkURLTemplate: apiclientgo.String("https://www.google.com/search?q={query}&hl=en"),
                                },
                                Label: apiclientgo.String("Mortimer's PRs"),
                                Datasource: apiclientgo.String("github"),
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
                                    FacetBucketSize: 991464,
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
                                Ranges: []components.TextRange{
                                    components.TextRange{
                                        StartIndex: 488852,
                                        Document: &components.Document{
                                            Metadata: &components.DocumentMetadata{
                                                Datasource: apiclientgo.String("datasource"),
                                                ObjectType: apiclientgo.String("Feature Request"),
                                                Container: apiclientgo.String("container"),
                                                ParentID: apiclientgo.String("JIRA_EN-1337"),
                                                MimeType: apiclientgo.String("mimeType"),
                                                DocumentID: apiclientgo.String("documentId"),
                                                CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                                                UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                                                Components: []string{
                                                    "Backend",
                                                    "Networking",
                                                },
                                                Status: apiclientgo.String("[\"Done\"]"),
                                                Pins: []components.PinDocument{
                                                    components.PinDocument{
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
                                                        DocumentID: "<id>",
                                                    },
                                                    components.PinDocument{
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
                                                        DocumentID: "<id>",
                                                    },
                                                    components.PinDocument{
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
                                                        DocumentID: "<id>",
                                                    },
                                                },
                                                Collections: []components.Collection{
                                                    components.Collection{
                                                        Name: "<value>",
                                                        Description: "eulogise whereas till mild than during meanwhile disapprove finer ha",
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
                                                        ID: 2984,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 477967,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 813577,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                        },
                                                    },
                                                },
                                                Interactions: &components.DocumentInteractions{
                                                    Reacts: []components.Reaction{
                                                        components.Reaction{},
                                                    },
                                                    Shares: []components.Share{
                                                        components.Share{
                                                            NumDaysAgo: 301848,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 657278,
                                                        },
                                                    },
                                                },
                                                Verification: &components.Verification{
                                                    State: components.StateUnverified,
                                                    Metadata: &components.VerificationMetadata{
                                                        Reminders: []components.Reminder{
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 335191,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 532806,
                                                        },
                                                    },
                                                },
                                                Shortcuts: []components.Shortcut{
                                                    components.Shortcut{
                                                        InputAlias: "<value>",
                                                    },
                                                    components.Shortcut{
                                                        InputAlias: "<value>",
                                                    },
                                                    components.Shortcut{
                                                        InputAlias: "<value>",
                                                    },
                                                },
                                                CustomData: map[string]components.CustomDataValue{
                                                    "someCustomField": components.CustomDataValue{},
                                                },
                                            },
                                        },
                                    },
                                    components.TextRange{
                                        StartIndex: 556706,
                                        Document: &components.Document{
                                            Metadata: &components.DocumentMetadata{
                                                Datasource: apiclientgo.String("datasource"),
                                                ObjectType: apiclientgo.String("Feature Request"),
                                                Container: apiclientgo.String("container"),
                                                ParentID: apiclientgo.String("JIRA_EN-1337"),
                                                MimeType: apiclientgo.String("mimeType"),
                                                DocumentID: apiclientgo.String("documentId"),
                                                CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                                                UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                                                Components: []string{
                                                    "Backend",
                                                    "Networking",
                                                },
                                                Status: apiclientgo.String("[\"Done\"]"),
                                                Interactions: &components.DocumentInteractions{},
                                                Verification: &components.Verification{
                                                    State: components.StateVerified,
                                                    Metadata: &components.VerificationMetadata{
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 416023,
                                                        },
                                                    },
                                                },
                                                CustomData: map[string]components.CustomDataValue{
                                                    "someCustomField": components.CustomDataValue{},
                                                },
                                            },
                                        },
                                    },
                                },
                                InputDetails: &components.SearchRequestInputDetails{
                                    HasCopyPaste: apiclientgo.Bool(true),
                                },
                            },
                            Results: []components.SearchResult{
                                components.SearchResult{
                                    Title: apiclientgo.String("title"),
                                    URL: "https://example.com/foo/bar",
                                    NativeAppURL: apiclientgo.String("slack://foo/bar"),
                                    Snippets: []components.SearchResultSnippet{
                                        components.SearchResultSnippet{
                                            Snippet: "snippet",
                                            MimeType: apiclientgo.String("mimeType"),
                                        },
                                    },
                                },
                            },
                        },
                        components.RelatedDocuments{
                            QuerySuggestion: &components.QuerySuggestion{
                                Query: "app:github type:pull author:mortimer",
                                SearchProviderInfo: &components.SearchProviderInfo{
                                    Name: apiclientgo.String("Google"),
                                    SearchLinkURLTemplate: apiclientgo.String("https://www.google.com/search?q={query}&hl=en"),
                                },
                                Label: apiclientgo.String("Mortimer's PRs"),
                                Datasource: apiclientgo.String("github"),
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
                                    FacetBucketSize: 967252,
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
                                InputDetails: &components.SearchRequestInputDetails{
                                    HasCopyPaste: apiclientgo.Bool(true),
                                },
                            },
                        },
                    },
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
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                            },
                        },
                        CustomFields: []components.CustomFieldData{
                            components.CustomFieldData{
                                Label: "<value>",
                                Values: []components.CustomFieldValue{},
                            },
                            components.CustomFieldData{
                                Label: "<value>",
                                Values: []components.CustomFieldValue{},
                            },
                            components.CustomFieldData{
                                Label: "<value>",
                                Values: []components.CustomFieldValue{},
                            },
                        },
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
            components.UserRoleSpecification{
                Person: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
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
                        QuerySuggestions: &components.QuerySuggestionList{},
                        InviteInfo: &components.InviteInfo{},
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
                Person: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
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
                        QuerySuggestions: &components.QuerySuggestionList{},
                        InviteInfo: &components.InviteInfo{},
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
                Person: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
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
                        QuerySuggestions: &components.QuerySuggestionList{},
                        InviteInfo: &components.InviteInfo{},
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
        ID: 720396,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EditCollectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.EditCollectionRequest](../../models/components/editcollectionrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.EditcollectionResponse](../../models/operations/editcollectionresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| apierrors.CollectionError | 422                       | application/json          |
| apierrors.APIError        | 4XX, 5XX                  | \*/\*                     |

## UpdateItem

Update the URL, Glean Document ID, description of an item within a Collection given its ID.

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

    res, err := s.Client.Collections.UpdateItem(ctx, components.EditCollectionItemRequest{
        CollectionID: 795203,
        ItemID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EditCollectionItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.EditCollectionItemRequest](../../models/components/editcollectionitemrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.EditcollectionitemResponse](../../models/operations/editcollectionitemresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Retrieve

Read the details of a Collection given its ID. Does not fetch items in this Collection.

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

    res, err := s.Client.Collections.Retrieve(ctx, components.GetCollectionRequest{
        ID: 700347,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetCollectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.GetCollectionRequest](../../models/components/getcollectionrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetcollectionResponse](../../models/operations/getcollectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

List all existing Collections.

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

    res, err := s.Client.Collections.List(ctx, components.ListCollectionsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCollectionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.ListCollectionsRequest](../../models/components/listcollectionsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListcollectionsResponse](../../models/operations/listcollectionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |