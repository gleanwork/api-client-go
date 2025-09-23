# Answers
(*Client.Answers*)

## Overview

### Available Operations

* [Create](#create) - Create Answer
* [Delete](#delete) - Delete Answer
* [Update](#update) - Update Answer
* [Retrieve](#retrieve) - Read Answer
* [List](#list) - List Answers

## Create

Create a user-generated Answer that contains a question and answer.

### Example Usage

<!-- UsageSnippet language="go" operationID="createanswer" method="post" path="/rest/api/v1/createanswer" -->
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

    res, err := s.Client.Answers.Create(ctx, components.CreateAnswerRequest{
        Data: components.AnswerCreationData{
            Question: apiclientgo.Pointer("Why is the sky blue?"),
            BodyText: apiclientgo.Pointer("From https://en.wikipedia.org/wiki/Diffuse_sky_radiation, the sky is blue because blue light is more strongly scattered than longer-wavelength light."),
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
                                        },
                                        FacetBucketSize: 708564,
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
                                            StartIndex: 945729,
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
                                                            Description: "via vain astride question",
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
                                                            ID: 51416,
                                                            Items: []components.CollectionItem{
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                            },
                                                        },
                                                        components.Collection{
                                                            Name: "<value>",
                                                            Description: "via vain astride question",
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
                                                            ID: 51416,
                                                            Items: []components.CollectionItem{
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
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
                                                                NumDaysAgo: 85387,
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
                                                                    RemindAt: 161639,
                                                                },
                                                            },
                                                            LastReminder: &components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 613051,
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
                                            StartIndex: 945729,
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
                                                            Description: "via vain astride question",
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
                                                            ID: 51416,
                                                            Items: []components.CollectionItem{
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                            },
                                                        },
                                                        components.Collection{
                                                            Name: "<value>",
                                                            Description: "via vain astride question",
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
                                                            ID: 51416,
                                                            Items: []components.CollectionItem{
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
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
                                                                NumDaysAgo: 85387,
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
                                                                    RemindAt: 161639,
                                                                },
                                                            },
                                                            LastReminder: &components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 613051,
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
                                        },
                                        FacetBucketSize: 708564,
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
                                            StartIndex: 945729,
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
                                                            Description: "via vain astride question",
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
                                                            ID: 51416,
                                                            Items: []components.CollectionItem{
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                            },
                                                        },
                                                        components.Collection{
                                                            Name: "<value>",
                                                            Description: "via vain astride question",
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
                                                            ID: 51416,
                                                            Items: []components.CollectionItem{
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
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
                                                                NumDaysAgo: 85387,
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
                                                                    RemindAt: 161639,
                                                                },
                                                            },
                                                            LastReminder: &components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 613051,
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
                                            StartIndex: 945729,
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
                                                            Description: "via vain astride question",
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
                                                            ID: 51416,
                                                            Items: []components.CollectionItem{
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                            },
                                                        },
                                                        components.Collection{
                                                            Name: "<value>",
                                                            Description: "via vain astride question",
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
                                                            ID: 51416,
                                                            Items: []components.CollectionItem{
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
                                                                    Shortcut: &components.Shortcut{
                                                                        InputAlias: "<value>",
                                                                    },
                                                                    ItemType: components.CollectionItemItemTypeText,
                                                                },
                                                                components.CollectionItem{
                                                                    CollectionID: 33920,
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
                                                                NumDaysAgo: 85387,
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
                                                                    RemindAt: 161639,
                                                                },
                                                            },
                                                            LastReminder: &components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 613051,
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
                    Role: components.UserRoleEditor,
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
                    Role: components.UserRoleOwner,
                },
            },
            Roles: []components.UserRoleSpecification{
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

<!-- UsageSnippet language="go" operationID="deleteanswer" method="post" path="/rest/api/v1/deleteanswer" -->
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

    res, err := s.Client.Answers.Delete(ctx, components.DeleteAnswerRequest{
        ID: 3,
        DocID: apiclientgo.Pointer("ANSWERS_answer_3"),
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

## Update

Update an existing user-generated Answer.

### Example Usage

<!-- UsageSnippet language="go" operationID="editanswer" method="post" path="/rest/api/v1/editanswer" -->
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

    res, err := s.Client.Answers.Update(ctx, components.EditAnswerRequest{
        ID: 3,
        DocID: apiclientgo.Pointer("ANSWERS_answer_3"),
        Question: apiclientgo.Pointer("Why is the sky blue?"),
        BodyText: apiclientgo.Pointer("From https://en.wikipedia.org/wiki/Diffuse_sky_radiation, the sky is blue because blue light is more strongly scattered than longer-wavelength light."),
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
                                    FacetBucketSize: 552856,
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
                                        StartIndex: 23264,
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
                                                        Description: "boohoo hunger energetic who whoa grimy vibrant wisely",
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
                                                        ID: 919335,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 972120,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 972120,
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
                                                            NumDaysAgo: 439797,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 439797,
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
                                                                RemindAt: 996442,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 996442,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 886538,
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
                                Values: []components.CustomFieldValue{},
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
                                    FacetBucketSize: 552856,
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
                                        StartIndex: 23264,
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
                                                        Description: "boohoo hunger energetic who whoa grimy vibrant wisely",
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
                                                        ID: 919335,
                                                        Items: []components.CollectionItem{
                                                            components.CollectionItem{
                                                                CollectionID: 972120,
                                                                Shortcut: &components.Shortcut{
                                                                    InputAlias: "<value>",
                                                                },
                                                                ItemType: components.CollectionItemItemTypeText,
                                                            },
                                                            components.CollectionItem{
                                                                CollectionID: 972120,
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
                                                            NumDaysAgo: 439797,
                                                        },
                                                        components.Share{
                                                            NumDaysAgo: 439797,
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
                                                                RemindAt: 996442,
                                                            },
                                                            components.Reminder{
                                                                Assignee: components.Person{
                                                                    Name: "George Clooney",
                                                                    ObfuscatedID: "abc123",
                                                                },
                                                                RemindAt: 996442,
                                                            },
                                                        },
                                                        LastReminder: &components.Reminder{
                                                            Assignee: components.Person{
                                                                Name: "George Clooney",
                                                                ObfuscatedID: "abc123",
                                                            },
                                                            RemindAt: 886538,
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
                                Values: []components.CustomFieldValue{},
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
                Role: components.UserRoleViewer,
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
                Role: components.UserRoleViewer,
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
                Role: components.UserRoleViewer,
            },
        },
        Roles: []components.UserRoleSpecification{
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

## Retrieve

Read the details of a particular Answer given its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getanswer" method="post" path="/rest/api/v1/getanswer" -->
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

    res, err := s.Client.Answers.Retrieve(ctx, components.GetAnswerRequest{
        ID: apiclientgo.Pointer[int64](3),
        DocID: apiclientgo.Pointer("ANSWERS_answer_3"),
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

<!-- UsageSnippet language="go" operationID="listanswers" method="post" path="/rest/api/v1/listanswers" -->
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