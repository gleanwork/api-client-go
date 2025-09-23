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

<!-- UsageSnippet language="go" operationID="createshortcut" method="post" path="/rest/api/v1/createshortcut" -->
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

    res, err := s.Client.Shortcuts.Create(ctx, components.CreateShortcutRequest{
        Data: components.ShortcutMutableProperties{
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
                                        FacetBucketSize: 523271,
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
                                            StartIndex: 365818,
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
                                                            Description: "terribly before ferociously how preclude aw quarterly definite schlep",
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
                                                            ID: 100369,
                                                            Items: []components.CollectionItem{
                                                                components.CollectionItem{
                                                                    CollectionID: 239768,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 239768,
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
                                                                NumDaysAgo: 771560,
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
                                                                    RemindAt: 919708,
                                                                },
                                                                components.Reminder{
                                                                    Assignee: components.Person{
                                                                        Name: "George Clooney",
                                                                        ObfuscatedID: "abc123",
                                                                    },
                                                                    RemindAt: 919708,
                                                                },
                                                            },
                                                            LastReminder: &components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 627228,
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
                                        components.TextRange{
                                            StartIndex: 365818,
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
                                                            Description: "terribly before ferociously how preclude aw quarterly definite schlep",
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
                                                            ID: 100369,
                                                            Items: []components.CollectionItem{
                                                                components.CollectionItem{
                                                                    CollectionID: 239768,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 239768,
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
                                                                NumDaysAgo: 771560,
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
                                                                    RemindAt: 919708,
                                                                },
                                                                components.Reminder{
                                                                    Assignee: components.Person{
                                                                        Name: "George Clooney",
                                                                        ObfuscatedID: "abc123",
                                                                    },
                                                                    RemindAt: 919708,
                                                                },
                                                            },
                                                            LastReminder: &components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 627228,
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
                    Role: components.UserRoleViewer,
                },
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
                                        FacetBucketSize: 523271,
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
                                            StartIndex: 365818,
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
                                                            Description: "terribly before ferociously how preclude aw quarterly definite schlep",
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
                                                            ID: 100369,
                                                            Items: []components.CollectionItem{
                                                                components.CollectionItem{
                                                                    CollectionID: 239768,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 239768,
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
                                                                NumDaysAgo: 771560,
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
                                                                    RemindAt: 919708,
                                                                },
                                                                components.Reminder{
                                                                    Assignee: components.Person{
                                                                        Name: "George Clooney",
                                                                        ObfuscatedID: "abc123",
                                                                    },
                                                                    RemindAt: 919708,
                                                                },
                                                            },
                                                            LastReminder: &components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 627228,
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
                                        components.TextRange{
                                            StartIndex: 365818,
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
                                                            Description: "terribly before ferociously how preclude aw quarterly definite schlep",
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
                                                            ID: 100369,
                                                            Items: []components.CollectionItem{
                                                                components.CollectionItem{
                                                                    CollectionID: 239768,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 239768,
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
                                                                NumDaysAgo: 771560,
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
                                                                    RemindAt: 919708,
                                                                },
                                                                components.Reminder{
                                                                    Assignee: components.Person{
                                                                        Name: "George Clooney",
                                                                        ObfuscatedID: "abc123",
                                                                    },
                                                                    RemindAt: 919708,
                                                                },
                                                            },
                                                            LastReminder: &components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 627228,
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
                    Role: components.UserRoleViewer,
                },
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
                                        FacetBucketSize: 523271,
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
                                            StartIndex: 365818,
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
                                                            Description: "terribly before ferociously how preclude aw quarterly definite schlep",
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
                                                            ID: 100369,
                                                            Items: []components.CollectionItem{
                                                                components.CollectionItem{
                                                                    CollectionID: 239768,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 239768,
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
                                                                NumDaysAgo: 771560,
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
                                                                    RemindAt: 919708,
                                                                },
                                                                components.Reminder{
                                                                    Assignee: components.Person{
                                                                        Name: "George Clooney",
                                                                        ObfuscatedID: "abc123",
                                                                    },
                                                                    RemindAt: 919708,
                                                                },
                                                            },
                                                            LastReminder: &components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 627228,
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
                                        components.TextRange{
                                            StartIndex: 365818,
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
                                                            Description: "terribly before ferociously how preclude aw quarterly definite schlep",
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
                                                            ID: 100369,
                                                            Items: []components.CollectionItem{
                                                                components.CollectionItem{
                                                                    CollectionID: 239768,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 239768,
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
                                                                NumDaysAgo: 771560,
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
                                                                    RemindAt: 919708,
                                                                },
                                                                components.Reminder{
                                                                    Assignee: components.Person{
                                                                        Name: "George Clooney",
                                                                        ObfuscatedID: "abc123",
                                                                    },
                                                                    RemindAt: 919708,
                                                                },
                                                            },
                                                            LastReminder: &components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 627228,
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
                    Role: components.UserRoleViewer,
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
                    Role: components.UserRoleOwner,
                },
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
                    Role: components.UserRoleOwner,
                },
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

<!-- UsageSnippet language="go" operationID="deleteshortcut" method="post" path="/rest/api/v1/deleteshortcut" -->
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

    res, err := s.Client.Shortcuts.Delete(ctx, components.DeleteShortcutRequest{
        ID: 975862,
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

<!-- UsageSnippet language="go" operationID="getshortcut" method="post" path="/rest/api/v1/getshortcut" -->
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

<!-- UsageSnippet language="go" operationID="listshortcuts" method="post" path="/rest/api/v1/listshortcuts" -->
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

    res, err := s.Client.Shortcuts.List(ctx, components.ListShortcutsPaginatedRequest{
        PageSize: 10,
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

<!-- UsageSnippet language="go" operationID="updateshortcut" method="post" path="/rest/api/v1/updateshortcut" -->
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

    res, err := s.Client.Shortcuts.Update(ctx, components.UpdateShortcutRequest{
        ID: 268238,
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
                                    FacetBucketSize: 801764,
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
                                        StartIndex: 352334,
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
                                                        Description: "scare following gadzooks woot translation",
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
                                                        ID: 78594,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
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
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
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
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 234150,
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
                                        StartIndex: 352334,
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
                                                        Description: "scare following gadzooks woot translation",
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
                                                        ID: 78594,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
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
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
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
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 234150,
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
                                        StartIndex: 352334,
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
                                                        Description: "scare following gadzooks woot translation",
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
                                                        ID: 78594,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
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
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
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
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 234150,
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
                                    },
                                    FacetBucketSize: 801764,
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
                                        StartIndex: 352334,
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
                                                        Description: "scare following gadzooks woot translation",
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
                                                        ID: 78594,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
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
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
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
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 234150,
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
                                        StartIndex: 352334,
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
                                                        Description: "scare following gadzooks woot translation",
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
                                                        ID: 78594,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
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
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
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
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 234150,
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
                                        StartIndex: 352334,
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
                                                        Description: "scare following gadzooks woot translation",
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
                                                        ID: 78594,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
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
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
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
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 234150,
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
                Role: components.UserRoleAnswerModerator,
            },
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
                                    FacetBucketSize: 801764,
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
                                        StartIndex: 352334,
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
                                                        Description: "scare following gadzooks woot translation",
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
                                                        ID: 78594,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
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
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
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
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 234150,
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
                                        StartIndex: 352334,
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
                                                        Description: "scare following gadzooks woot translation",
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
                                                        ID: 78594,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
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
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
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
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 234150,
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
                                        StartIndex: 352334,
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
                                                        Description: "scare following gadzooks woot translation",
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
                                                        ID: 78594,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
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
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
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
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 234150,
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
                                    },
                                    FacetBucketSize: 801764,
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
                                        StartIndex: 352334,
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
                                                        Description: "scare following gadzooks woot translation",
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
                                                        ID: 78594,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
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
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
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
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 234150,
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
                                        StartIndex: 352334,
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
                                                        Description: "scare following gadzooks woot translation",
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
                                                        ID: 78594,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
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
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
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
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 234150,
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
                                        StartIndex: 352334,
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
                                                        Description: "scare following gadzooks woot translation",
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
                                                        ID: 78594,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 56865,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeURL,
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
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 776330,
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
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 84322,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 234150,
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
                Role: components.UserRoleEditor,
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