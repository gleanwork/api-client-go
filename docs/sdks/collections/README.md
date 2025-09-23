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

<!-- UsageSnippet language="go" operationID="addcollectionitems" method="post" path="/rest/api/v1/addcollectionitems" -->
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
        CollectionID: 7742.68,
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

<!-- UsageSnippet language="go" operationID="createcollection" method="post" path="/rest/api/v1/createcollection" -->
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
                                    Name: apiclientgo.Pointer("Google"),
                                    SearchLinkURLTemplate: apiclientgo.Pointer("https://www.google.com/search?q={query}&hl=en"),
                                },
                                Label: apiclientgo.Pointer("Mortimer's PRs"),
                                Datasource: apiclientgo.Pointer("github"),
                                RequestOptions: &components.SearchRequestOptions{
                                    DatasourceFilter: apiclientgo.Pointer("JIRA"),
                                    DatasourcesFilter: []string{
                                        "JIRA",
                                    },
                                    QueryOverridesFacetFilters: apiclientgo.Pointer(true),
                                    FacetFilters: []components.FacetFilter{
                                        components.FacetFilter{
                                            FieldName: apiclientgo.Pointer("type"),
                                            Values: []components.FacetFilterValue{
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.Pointer("Spreadsheet"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.Pointer("Presentation"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                            },
                                        },
                                    },
                                    FacetFilterSets: []components.FacetFilterSet{
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.Pointer("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.Pointer("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.Pointer("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    FacetBucketSize: 977077,
                                    AuthTokens: []components.AuthToken{
                                        components.AuthToken{
                                            AccessToken: "123abc",
                                            Datasource: "gmail",
                                            Scope: apiclientgo.Pointer("email profile https://www.googleapis.com/auth/gmail.readonly"),
                                            TokenType: apiclientgo.Pointer("Bearer"),
                                            AuthUser: apiclientgo.Pointer("1"),
                                        },
                                    },
                                },
                                Ranges: []components.TextRange{
                                    components.TextRange{
                                        StartIndex: 86650,
                                        Document: &components.Document{
                                            Metadata: &components.DocumentMetadata{
                                                Datasource: apiclientgo.Pointer("datasource"),
                                                ObjectType: apiclientgo.Pointer("Feature Request"),
                                                Container: apiclientgo.Pointer("container"),
                                                ParentID: apiclientgo.Pointer("JIRA_EN-1337"),
                                                MimeType: apiclientgo.Pointer("mimeType"),
                                                DocumentID: apiclientgo.Pointer("documentId"),
                                                CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                                                UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                                                Components: []string{
                                                    "Backend",
                                                    "Networking",
                                                },
                                                Status: apiclientgo.Pointer("[\"Done\"]"),
                                                Pins: []components.PinDocument{
                                                    components.PinDocument{
                                                        AudienceFilters: []components.FacetFilter{
                                                            components.FacetFilter{
                                                                FieldName: apiclientgo.Pointer("type"),
                                                                Values: []components.FacetFilterValue{
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Spreadsheet"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Presentation"),
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
                                                        Description: "meaty dial elegantly while react",
                                                        AudienceFilters: []components.FacetFilter{
                                                            components.FacetFilter{
                                                                FieldName: apiclientgo.Pointer("type"),
                                                                Values: []components.FacetFilterValue{
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Spreadsheet"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Presentation"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        ID: 854591,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 697663,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 697663,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 697663,
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
                                                        components.Reaction{},
                                                    },
                                                    Shares: []components.Share{
                                                        components.Share{
                                                            NumDaysAgo: 365776,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 365776,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 365776,
                                                        },
                                                    },
                                                },
                                                Verification: &components.Verification{
                                                    State: components.StateDeprecated,
                                                    Metadata: &components.VerificationMetadata{
                                                        Reminders: []components.Reminder{
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 268615,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 423482,
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
                                },
                                InputDetails: &components.SearchRequestInputDetails{
                                    HasCopyPaste: apiclientgo.Pointer(true),
                                },
                            },
                            Results: []components.SearchResult{
                                components.SearchResult{
                                    Title: apiclientgo.Pointer("title"),
                                    URL: "https://example.com/foo/bar",
                                    NativeAppURL: apiclientgo.Pointer("slack://foo/bar"),
                                    Snippets: []components.SearchResultSnippet{
                                        components.SearchResultSnippet{
                                            MimeType: apiclientgo.Pointer("mimeType"),
                                            Snippet: "snippet",
                                        },
                                    },
                                },
                            },
                        },
                        components.RelatedDocuments{
                            QuerySuggestion: &components.QuerySuggestion{
                                Query: "app:github type:pull author:mortimer",
                                SearchProviderInfo: &components.SearchProviderInfo{
                                    Name: apiclientgo.Pointer("Google"),
                                    SearchLinkURLTemplate: apiclientgo.Pointer("https://www.google.com/search?q={query}&hl=en"),
                                },
                                Label: apiclientgo.Pointer("Mortimer's PRs"),
                                Datasource: apiclientgo.Pointer("github"),
                                RequestOptions: &components.SearchRequestOptions{
                                    DatasourceFilter: apiclientgo.Pointer("JIRA"),
                                    DatasourcesFilter: []string{
                                        "JIRA",
                                    },
                                    QueryOverridesFacetFilters: apiclientgo.Pointer(true),
                                    FacetFilters: []components.FacetFilter{
                                        components.FacetFilter{
                                            FieldName: apiclientgo.Pointer("type"),
                                            Values: []components.FacetFilterValue{
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.Pointer("Spreadsheet"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.Pointer("Presentation"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                            },
                                        },
                                    },
                                    FacetFilterSets: []components.FacetFilterSet{
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.Pointer("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.Pointer("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.Pointer("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    FacetBucketSize: 977077,
                                    AuthTokens: []components.AuthToken{
                                        components.AuthToken{
                                            AccessToken: "123abc",
                                            Datasource: "gmail",
                                            Scope: apiclientgo.Pointer("email profile https://www.googleapis.com/auth/gmail.readonly"),
                                            TokenType: apiclientgo.Pointer("Bearer"),
                                            AuthUser: apiclientgo.Pointer("1"),
                                        },
                                    },
                                },
                                Ranges: []components.TextRange{
                                    components.TextRange{
                                        StartIndex: 86650,
                                        Document: &components.Document{
                                            Metadata: &components.DocumentMetadata{
                                                Datasource: apiclientgo.Pointer("datasource"),
                                                ObjectType: apiclientgo.Pointer("Feature Request"),
                                                Container: apiclientgo.Pointer("container"),
                                                ParentID: apiclientgo.Pointer("JIRA_EN-1337"),
                                                MimeType: apiclientgo.Pointer("mimeType"),
                                                DocumentID: apiclientgo.Pointer("documentId"),
                                                CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                                                UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                                                Components: []string{
                                                    "Backend",
                                                    "Networking",
                                                },
                                                Status: apiclientgo.Pointer("[\"Done\"]"),
                                                Pins: []components.PinDocument{
                                                    components.PinDocument{
                                                        AudienceFilters: []components.FacetFilter{
                                                            components.FacetFilter{
                                                                FieldName: apiclientgo.Pointer("type"),
                                                                Values: []components.FacetFilterValue{
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Spreadsheet"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Presentation"),
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
                                                        Description: "meaty dial elegantly while react",
                                                        AudienceFilters: []components.FacetFilter{
                                                            components.FacetFilter{
                                                                FieldName: apiclientgo.Pointer("type"),
                                                                Values: []components.FacetFilterValue{
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Spreadsheet"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Presentation"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        ID: 854591,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 697663,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 697663,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 697663,
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
                                                        components.Reaction{},
                                                    },
                                                    Shares: []components.Share{
                                                        components.Share{
                                                            NumDaysAgo: 365776,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 365776,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 365776,
                                                        },
                                                    },
                                                },
                                                Verification: &components.Verification{
                                                    State: components.StateDeprecated,
                                                    Metadata: &components.VerificationMetadata{
                                                        Reminders: []components.Reminder{
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 268615,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 423482,
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
                                },
                                InputDetails: &components.SearchRequestInputDetails{
                                    HasCopyPaste: apiclientgo.Pointer(true),
                                },
                            },
                            Results: []components.SearchResult{
                                components.SearchResult{
                                    Title: apiclientgo.Pointer("title"),
                                    URL: "https://example.com/foo/bar",
                                    NativeAppURL: apiclientgo.Pointer("slack://foo/bar"),
                                    Snippets: []components.SearchResultSnippet{
                                        components.SearchResultSnippet{
                                            MimeType: apiclientgo.Pointer("mimeType"),
                                            Snippet: "snippet",
                                        },
                                    },
                                },
                            },
                        },
                    },
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.Pointer("Actor"),
                        Department: apiclientgo.Pointer("Movies"),
                        Email: apiclientgo.Pointer("george@example.com"),
                        Location: apiclientgo.Pointer("Hollywood, CA"),
                        Phone: apiclientgo.Pointer("6505551234"),
                        PhotoURL: apiclientgo.Pointer("https://example.com/george.jpg"),
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
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.Pointer("Mortimer's PRs"),
                                    Datasource: apiclientgo.Pointer("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
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
                                    components.CreateCustomFieldValueCustomFieldValueStr(
                                        components.CustomFieldValueStr{},
                                    ),
                                },
                            },
                            components.CustomFieldData{
                                Label: "<value>",
                                Values: []components.CustomFieldValue{
                                    components.CreateCustomFieldValueCustomFieldValueStr(
                                        components.CustomFieldValueStr{},
                                    ),
                                    components.CreateCustomFieldValueCustomFieldValueStr(
                                        components.CustomFieldValueStr{},
                                    ),
                                    components.CreateCustomFieldValueCustomFieldValueStr(
                                        components.CustomFieldValueStr{},
                                    ),
                                },
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.Pointer("deployment_name_new_hire"),
                                DisplayName: apiclientgo.Pointer("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.Pointer("#343CED"),
                                    Key: apiclientgo.Pointer("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.Pointer("user"),
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
                        Title: apiclientgo.Pointer("Actor"),
                        Department: apiclientgo.Pointer("Movies"),
                        Email: apiclientgo.Pointer("george@example.com"),
                        Location: apiclientgo.Pointer("Hollywood, CA"),
                        Phone: apiclientgo.Pointer("6505551234"),
                        PhotoURL: apiclientgo.Pointer("https://example.com/george.jpg"),
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
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.Pointer("Mortimer's PRs"),
                                    Datasource: apiclientgo.Pointer("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                                components.ChannelInviteInfo{},
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.Pointer("deployment_name_new_hire"),
                                DisplayName: apiclientgo.Pointer("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.Pointer("#343CED"),
                                    Key: apiclientgo.Pointer("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.Pointer("user"),
                                },
                            },
                        },
                    },
                },
                Role: components.UserRoleViewer,
            },
        },
        AudienceFilters: []components.FacetFilter{
            components.FacetFilter{
                FieldName: apiclientgo.Pointer("type"),
                Values: []components.FacetFilterValue{
                    components.FacetFilterValue{
                        Value: apiclientgo.Pointer("Spreadsheet"),
                        RelationType: components.RelationTypeEquals.ToPointer(),
                    },
                    components.FacetFilterValue{
                        Value: apiclientgo.Pointer("Presentation"),
                        RelationType: components.RelationTypeEquals.ToPointer(),
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
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

<!-- UsageSnippet language="go" operationID="deletecollection" method="post" path="/rest/api/v1/deletecollection" -->
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
            930352,
            156719,
            25102,
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

<!-- UsageSnippet language="go" operationID="deletecollectionitem" method="post" path="/rest/api/v1/deletecollectionitem" -->
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
        CollectionID: 6980.49,
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

<!-- UsageSnippet language="go" operationID="editcollection" method="post" path="/rest/api/v1/editcollection" -->
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
                                    Name: apiclientgo.Pointer("Google"),
                                    SearchLinkURLTemplate: apiclientgo.Pointer("https://www.google.com/search?q={query}&hl=en"),
                                },
                                Label: apiclientgo.Pointer("Mortimer's PRs"),
                                Datasource: apiclientgo.Pointer("github"),
                                RequestOptions: &components.SearchRequestOptions{
                                    DatasourceFilter: apiclientgo.Pointer("JIRA"),
                                    DatasourcesFilter: []string{
                                        "JIRA",
                                    },
                                    QueryOverridesFacetFilters: apiclientgo.Pointer(true),
                                    FacetFilters: []components.FacetFilter{
                                        components.FacetFilter{
                                            FieldName: apiclientgo.Pointer("type"),
                                            Values: []components.FacetFilterValue{
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.Pointer("Spreadsheet"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.Pointer("Presentation"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                            },
                                        },
                                    },
                                    FacetFilterSets: []components.FacetFilterSet{
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.Pointer("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.Pointer("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    FacetBucketSize: 797260,
                                    AuthTokens: []components.AuthToken{
                                        components.AuthToken{
                                            AccessToken: "123abc",
                                            Datasource: "gmail",
                                            Scope: apiclientgo.Pointer("email profile https://www.googleapis.com/auth/gmail.readonly"),
                                            TokenType: apiclientgo.Pointer("Bearer"),
                                            AuthUser: apiclientgo.Pointer("1"),
                                        },
                                    },
                                },
                                Ranges: []components.TextRange{
                                    components.TextRange{
                                        StartIndex: 932928,
                                        Document: &components.Document{
                                            Metadata: &components.DocumentMetadata{
                                                Datasource: apiclientgo.Pointer("datasource"),
                                                ObjectType: apiclientgo.Pointer("Feature Request"),
                                                Container: apiclientgo.Pointer("container"),
                                                ParentID: apiclientgo.Pointer("JIRA_EN-1337"),
                                                MimeType: apiclientgo.Pointer("mimeType"),
                                                DocumentID: apiclientgo.Pointer("documentId"),
                                                CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                                                UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                                                Components: []string{
                                                    "Backend",
                                                    "Networking",
                                                },
                                                Status: apiclientgo.Pointer("[\"Done\"]"),
                                                Pins: []components.PinDocument{
                                                    components.PinDocument{
                                                        AudienceFilters: []components.FacetFilter{
                                                            components.FacetFilter{
                                                                FieldName: apiclientgo.Pointer("type"),
                                                                Values: []components.FacetFilterValue{
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Spreadsheet"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Presentation"),
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
                                                                FieldName: apiclientgo.Pointer("type"),
                                                                Values: []components.FacetFilterValue{
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Spreadsheet"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Presentation"),
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
                                                        Description: "how by extroverted excess kissingly scruple yearningly",
                                                        AudienceFilters: []components.FacetFilter{
                                                            components.FacetFilter{
                                                                FieldName: apiclientgo.Pointer("type"),
                                                                Values: []components.FacetFilterValue{
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Spreadsheet"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Presentation"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        ID: 416110,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                        },
                                                    },
                                                    components.Collection{
                                                        Name: "<value>",
                                                        Description: "how by extroverted excess kissingly scruple yearningly",
                                                        AudienceFilters: []components.FacetFilter{
                                                            components.FacetFilter{
                                                                FieldName: apiclientgo.Pointer("type"),
                                                                Values: []components.FacetFilterValue{
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Spreadsheet"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Presentation"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        ID: 416110,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
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
                                                        components.Reaction{},
                                                        components.Reaction{},
                                                    },
                                                    Shares: []components.Share{
                                                        components.Share{
                                                            NumDaysAgo: 927933,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 927933,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 927933,
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
                                                                RemindAt: 353914,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 353914,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 314497,
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
                                                },
                                                CustomData: map[string]components.CustomDataValue{
                                                    "someCustomField": components.CustomDataValue{},
                                                },
                                            },
                                        },
                                    },
                                },
                                InputDetails: &components.SearchRequestInputDetails{
                                    HasCopyPaste: apiclientgo.Pointer(true),
                                },
                            },
                            Results: []components.SearchResult{
                                components.SearchResult{
                                    Title: apiclientgo.Pointer("title"),
                                    URL: "https://example.com/foo/bar",
                                    NativeAppURL: apiclientgo.Pointer("slack://foo/bar"),
                                    Snippets: []components.SearchResultSnippet{
                                        components.SearchResultSnippet{
                                            MimeType: apiclientgo.Pointer("mimeType"),
                                            Snippet: "snippet",
                                        },
                                    },
                                },
                            },
                        },
                        components.RelatedDocuments{
                            QuerySuggestion: &components.QuerySuggestion{
                                Query: "app:github type:pull author:mortimer",
                                SearchProviderInfo: &components.SearchProviderInfo{
                                    Name: apiclientgo.Pointer("Google"),
                                    SearchLinkURLTemplate: apiclientgo.Pointer("https://www.google.com/search?q={query}&hl=en"),
                                },
                                Label: apiclientgo.Pointer("Mortimer's PRs"),
                                Datasource: apiclientgo.Pointer("github"),
                                RequestOptions: &components.SearchRequestOptions{
                                    DatasourceFilter: apiclientgo.Pointer("JIRA"),
                                    DatasourcesFilter: []string{
                                        "JIRA",
                                    },
                                    QueryOverridesFacetFilters: apiclientgo.Pointer(true),
                                    FacetFilters: []components.FacetFilter{
                                        components.FacetFilter{
                                            FieldName: apiclientgo.Pointer("type"),
                                            Values: []components.FacetFilterValue{
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.Pointer("Spreadsheet"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.Pointer("Presentation"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                            },
                                        },
                                    },
                                    FacetFilterSets: []components.FacetFilterSet{
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.Pointer("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.Pointer("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    FacetBucketSize: 797260,
                                    AuthTokens: []components.AuthToken{
                                        components.AuthToken{
                                            AccessToken: "123abc",
                                            Datasource: "gmail",
                                            Scope: apiclientgo.Pointer("email profile https://www.googleapis.com/auth/gmail.readonly"),
                                            TokenType: apiclientgo.Pointer("Bearer"),
                                            AuthUser: apiclientgo.Pointer("1"),
                                        },
                                    },
                                },
                                Ranges: []components.TextRange{
                                    components.TextRange{
                                        StartIndex: 932928,
                                        Document: &components.Document{
                                            Metadata: &components.DocumentMetadata{
                                                Datasource: apiclientgo.Pointer("datasource"),
                                                ObjectType: apiclientgo.Pointer("Feature Request"),
                                                Container: apiclientgo.Pointer("container"),
                                                ParentID: apiclientgo.Pointer("JIRA_EN-1337"),
                                                MimeType: apiclientgo.Pointer("mimeType"),
                                                DocumentID: apiclientgo.Pointer("documentId"),
                                                CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                                                UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                                                Components: []string{
                                                    "Backend",
                                                    "Networking",
                                                },
                                                Status: apiclientgo.Pointer("[\"Done\"]"),
                                                Pins: []components.PinDocument{
                                                    components.PinDocument{
                                                        AudienceFilters: []components.FacetFilter{
                                                            components.FacetFilter{
                                                                FieldName: apiclientgo.Pointer("type"),
                                                                Values: []components.FacetFilterValue{
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Spreadsheet"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Presentation"),
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
                                                                FieldName: apiclientgo.Pointer("type"),
                                                                Values: []components.FacetFilterValue{
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Spreadsheet"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Presentation"),
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
                                                        Description: "how by extroverted excess kissingly scruple yearningly",
                                                        AudienceFilters: []components.FacetFilter{
                                                            components.FacetFilter{
                                                                FieldName: apiclientgo.Pointer("type"),
                                                                Values: []components.FacetFilterValue{
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Spreadsheet"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Presentation"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        ID: 416110,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                        },
                                                    },
                                                    components.Collection{
                                                        Name: "<value>",
                                                        Description: "how by extroverted excess kissingly scruple yearningly",
                                                        AudienceFilters: []components.FacetFilter{
                                                            components.FacetFilter{
                                                                FieldName: apiclientgo.Pointer("type"),
                                                                Values: []components.FacetFilterValue{
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Spreadsheet"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Presentation"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        ID: 416110,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
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
                                                        components.Reaction{},
                                                        components.Reaction{},
                                                    },
                                                    Shares: []components.Share{
                                                        components.Share{
                                                            NumDaysAgo: 927933,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 927933,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 927933,
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
                                                                RemindAt: 353914,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 353914,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 314497,
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
                                                },
                                                CustomData: map[string]components.CustomDataValue{
                                                    "someCustomField": components.CustomDataValue{},
                                                },
                                            },
                                        },
                                    },
                                },
                                InputDetails: &components.SearchRequestInputDetails{
                                    HasCopyPaste: apiclientgo.Pointer(true),
                                },
                            },
                            Results: []components.SearchResult{
                                components.SearchResult{
                                    Title: apiclientgo.Pointer("title"),
                                    URL: "https://example.com/foo/bar",
                                    NativeAppURL: apiclientgo.Pointer("slack://foo/bar"),
                                    Snippets: []components.SearchResultSnippet{
                                        components.SearchResultSnippet{
                                            MimeType: apiclientgo.Pointer("mimeType"),
                                            Snippet: "snippet",
                                        },
                                    },
                                },
                            },
                        },
                        components.RelatedDocuments{
                            QuerySuggestion: &components.QuerySuggestion{
                                Query: "app:github type:pull author:mortimer",
                                SearchProviderInfo: &components.SearchProviderInfo{
                                    Name: apiclientgo.Pointer("Google"),
                                    SearchLinkURLTemplate: apiclientgo.Pointer("https://www.google.com/search?q={query}&hl=en"),
                                },
                                Label: apiclientgo.Pointer("Mortimer's PRs"),
                                Datasource: apiclientgo.Pointer("github"),
                                RequestOptions: &components.SearchRequestOptions{
                                    DatasourceFilter: apiclientgo.Pointer("JIRA"),
                                    DatasourcesFilter: []string{
                                        "JIRA",
                                    },
                                    QueryOverridesFacetFilters: apiclientgo.Pointer(true),
                                    FacetFilters: []components.FacetFilter{
                                        components.FacetFilter{
                                            FieldName: apiclientgo.Pointer("type"),
                                            Values: []components.FacetFilterValue{
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.Pointer("Spreadsheet"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.Pointer("Presentation"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                            },
                                        },
                                    },
                                    FacetFilterSets: []components.FacetFilterSet{
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.Pointer("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.Pointer("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.Pointer("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    FacetBucketSize: 797260,
                                    AuthTokens: []components.AuthToken{
                                        components.AuthToken{
                                            AccessToken: "123abc",
                                            Datasource: "gmail",
                                            Scope: apiclientgo.Pointer("email profile https://www.googleapis.com/auth/gmail.readonly"),
                                            TokenType: apiclientgo.Pointer("Bearer"),
                                            AuthUser: apiclientgo.Pointer("1"),
                                        },
                                    },
                                },
                                Ranges: []components.TextRange{
                                    components.TextRange{
                                        StartIndex: 932928,
                                        Document: &components.Document{
                                            Metadata: &components.DocumentMetadata{
                                                Datasource: apiclientgo.Pointer("datasource"),
                                                ObjectType: apiclientgo.Pointer("Feature Request"),
                                                Container: apiclientgo.Pointer("container"),
                                                ParentID: apiclientgo.Pointer("JIRA_EN-1337"),
                                                MimeType: apiclientgo.Pointer("mimeType"),
                                                DocumentID: apiclientgo.Pointer("documentId"),
                                                CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                                                UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                                                Components: []string{
                                                    "Backend",
                                                    "Networking",
                                                },
                                                Status: apiclientgo.Pointer("[\"Done\"]"),
                                                Pins: []components.PinDocument{
                                                    components.PinDocument{
                                                        AudienceFilters: []components.FacetFilter{
                                                            components.FacetFilter{
                                                                FieldName: apiclientgo.Pointer("type"),
                                                                Values: []components.FacetFilterValue{
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Spreadsheet"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Presentation"),
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
                                                                FieldName: apiclientgo.Pointer("type"),
                                                                Values: []components.FacetFilterValue{
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Spreadsheet"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Presentation"),
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
                                                        Description: "how by extroverted excess kissingly scruple yearningly",
                                                        AudienceFilters: []components.FacetFilter{
                                                            components.FacetFilter{
                                                                FieldName: apiclientgo.Pointer("type"),
                                                                Values: []components.FacetFilterValue{
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Spreadsheet"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Presentation"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        ID: 416110,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                        },
                                                    },
                                                    components.Collection{
                                                        Name: "<value>",
                                                        Description: "how by extroverted excess kissingly scruple yearningly",
                                                        AudienceFilters: []components.FacetFilter{
                                                            components.FacetFilter{
                                                                FieldName: apiclientgo.Pointer("type"),
                                                                Values: []components.FacetFilterValue{
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Spreadsheet"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                    components.FacetFilterValue{
                                                                        Value: apiclientgo.Pointer("Presentation"),
                                                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                                                    },
                                                                },
                                                            },
                                                        },
                                                        ID: 416110,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 959645,
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
                                                        components.Reaction{},
                                                        components.Reaction{},
                                                    },
                                                    Shares: []components.Share{
                                                        components.Share{
                                                            NumDaysAgo: 927933,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 927933,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 927933,
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
                                                                RemindAt: 353914,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 353914,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 314497,
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
                                                },
                                                CustomData: map[string]components.CustomDataValue{
                                                    "someCustomField": components.CustomDataValue{},
                                                },
                                            },
                                        },
                                    },
                                },
                                InputDetails: &components.SearchRequestInputDetails{
                                    HasCopyPaste: apiclientgo.Pointer(true),
                                },
                            },
                            Results: []components.SearchResult{
                                components.SearchResult{
                                    Title: apiclientgo.Pointer("title"),
                                    URL: "https://example.com/foo/bar",
                                    NativeAppURL: apiclientgo.Pointer("slack://foo/bar"),
                                    Snippets: []components.SearchResultSnippet{
                                        components.SearchResultSnippet{
                                            MimeType: apiclientgo.Pointer("mimeType"),
                                            Snippet: "snippet",
                                        },
                                    },
                                },
                            },
                        },
                    },
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.Pointer("Actor"),
                        Department: apiclientgo.Pointer("Movies"),
                        Email: apiclientgo.Pointer("george@example.com"),
                        Location: apiclientgo.Pointer("Hollywood, CA"),
                        Phone: apiclientgo.Pointer("6505551234"),
                        PhotoURL: apiclientgo.Pointer("https://example.com/george.jpg"),
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
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.Pointer("Mortimer's PRs"),
                                    Datasource: apiclientgo.Pointer("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                                components.ChannelInviteInfo{},
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
                                    components.CreateCustomFieldValueCustomFieldValueStr(
                                        components.CustomFieldValueStr{},
                                    ),
                                },
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.Pointer("deployment_name_new_hire"),
                                DisplayName: apiclientgo.Pointer("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.Pointer("#343CED"),
                                    Key: apiclientgo.Pointer("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.Pointer("user"),
                                },
                            },
                        },
                    },
                },
                Role: components.UserRoleOwner,
            },
        },
        RemovedRoles: []components.UserRoleSpecification{
            components.UserRoleSpecification{
                Person: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.Pointer("Actor"),
                        Department: apiclientgo.Pointer("Movies"),
                        Email: apiclientgo.Pointer("george@example.com"),
                        Location: apiclientgo.Pointer("Hollywood, CA"),
                        Phone: apiclientgo.Pointer("6505551234"),
                        PhotoURL: apiclientgo.Pointer("https://example.com/george.jpg"),
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
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.Pointer("Mortimer's PRs"),
                                    Datasource: apiclientgo.Pointer("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                                components.ChannelInviteInfo{},
                                components.ChannelInviteInfo{},
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.Pointer("deployment_name_new_hire"),
                                DisplayName: apiclientgo.Pointer("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.Pointer("#343CED"),
                                    Key: apiclientgo.Pointer("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.Pointer("user"),
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
                FieldName: apiclientgo.Pointer("type"),
                Values: []components.FacetFilterValue{
                    components.FacetFilterValue{
                        Value: apiclientgo.Pointer("Spreadsheet"),
                        RelationType: components.RelationTypeEquals.ToPointer(),
                    },
                    components.FacetFilterValue{
                        Value: apiclientgo.Pointer("Presentation"),
                        RelationType: components.RelationTypeEquals.ToPointer(),
                    },
                },
            },
        },
        ID: 671264,
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

<!-- UsageSnippet language="go" operationID="editcollectionitem" method="post" path="/rest/api/v1/editcollectionitem" -->
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
        CollectionID: 142375,
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

<!-- UsageSnippet language="go" operationID="getcollection" method="post" path="/rest/api/v1/getcollection" -->
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
        ID: 425335,
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

<!-- UsageSnippet language="go" operationID="listcollections" method="post" path="/rest/api/v1/listcollections" -->
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