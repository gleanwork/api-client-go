# Announcements
(*Client.Announcements*)

## Overview

### Available Operations

* [Create](#create) - Create Announcement
* [Delete](#delete) - Delete Announcement
* [Update](#update) - Update Announcement

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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Client.Announcements.Create(ctx, components.CreateAnnouncementRequest{
        StartTime: types.MustTimeFromString("2024-06-17T07:14:55.338Z"),
        EndTime: types.MustTimeFromString("2024-11-30T17:06:07.804Z"),
        Title: "<value>",
        Body: &components.StructuredText{
            Text: "From https://en.wikipedia.org/wiki/Diffuse_sky_radiation, the sky is blue because blue light is more strongly scattered than longer-wavelength light.",
            StructuredList: []components.StructuredTextItem{
                components.StructuredTextItem{
                    Link: apiclientgo.String("https://en.wikipedia.org/wiki/Diffuse_sky_radiation"),
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
                            Author: &components.Person{
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
                                                },
                                                FacetBucketSize: 796474,
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
                                                    StartIndex: 86,
                                                },
                                                components.TextRange{
                                                    StartIndex: 169727,
                                                },
                                                components.TextRange{
                                                    StartIndex: 89964,
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
                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                Datasource: apiclientgo.String("github"),
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
                            Owner: &components.Person{
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
                            MentionedPeople: []components.Person{
                                components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                },
                            },
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
                                    Attribution: &components.Person{
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
                                    UpdatedBy: &components.Person{
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
                                },
                            },
                            AssignedTo: &components.Person{
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
                            UpdatedBy: &components.Person{
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
                            Collections: []components.Collection{
                                components.Collection{
                                    Name: "<value>",
                                    Description: "yuck mortally round",
                                    AddedRoles: []components.UserRoleSpecification{
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
                                            Role: components.UserRoleViewer,
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
                                    ID: 924484,
                                    Creator: &components.Person{
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
                                    UpdatedBy: &components.Person{
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
                                    Items: []components.CollectionItem{
                                        components.CollectionItem{
                                            CollectionID: 583353,
                                            CreatedBy: &components.Person{
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
                                            Shortcut: &components.Shortcut{
                                                InputAlias: "<value>",
                                                CreatedBy: &components.Person{
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
                                                UpdatedBy: &components.Person{
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
                                                Roles: []components.UserRoleSpecification{
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
                                                        Role: components.UserRoleVerifier,
                                                    },
                                                },
                                            },
                                            ItemType: components.CollectionItemItemTypeText,
                                        },
                                    },
                                },
                            },
                            Interactions: &components.DocumentInteractions{
                                Reacts: []components.Reaction{
                                    components.Reaction{
                                        Reactors: []components.Person{
                                            components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                            },
                                        },
                                    },
                                    components.Reaction{},
                                    components.Reaction{},
                                },
                                Shares: []components.Share{
                                    components.Share{
                                        NumDaysAgo: 911742,
                                        Sharer: &components.Person{
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
                                    },
                                },
                            },
                            Verification: &components.Verification{
                                State: components.StateVerified,
                                Metadata: &components.VerificationMetadata{
                                    LastVerifier: &components.Person{
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
                                    Reminders: []components.Reminder{
                                        components.Reminder{
                                            Assignee: components.Person{
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
                                            Requestor: &components.Person{
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
                                            RemindAt: 428745,
                                        },
                                    },
                                    LastReminder: &components.Reminder{
                                        Assignee: components.Person{
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
                                        Requestor: &components.Person{
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
                                        RemindAt: 860420,
                                    },
                                    CandidateVerifiers: []components.Person{
                                        components.Person{
                                            Name: "George Clooney",
                                            ObfuscatedID: "abc123",
                                        },
                                    },
                                },
                            },
                            Shortcuts: []components.Shortcut{
                                components.Shortcut{
                                    InputAlias: "<value>",
                                    CreatedBy: &components.Person{
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
                                    UpdatedBy: &components.Person{
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
                                },
                            },
                            CustomData: map[string]components.CustomDataValue{
                                "someCustomField": components.CustomDataValue{},
                            },
                            ContactPerson: &components.Person{
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
                        },
                    },
                    Text: apiclientgo.String("Because its wavelengths are shorter, blue light is more strongly scattered than the longer-wavelength lights, red or green. Hence the result that when looking at the sky away from the direct incident sunlight, the human eye perceives the sky to be blue."),
                },
                components.StructuredTextItem{
                    Link: apiclientgo.String("https://en.wikipedia.org/wiki/Diffuse_sky_radiation"),
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
                            Author: &components.Person{
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
                            Owner: &components.Person{
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
                            Components: []string{
                                "Backend",
                                "Networking",
                            },
                            Status: apiclientgo.String("[\"Done\"]"),
                            AssignedTo: &components.Person{
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
                            UpdatedBy: &components.Person{
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
                            Interactions: &components.DocumentInteractions{},
                            Verification: &components.Verification{
                                State: components.StateDeprecated,
                                Metadata: &components.VerificationMetadata{
                                    LastVerifier: &components.Person{
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
                                    LastReminder: &components.Reminder{
                                        Assignee: components.Person{
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
                                        Requestor: &components.Person{
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
                                        RemindAt: 284120,
                                    },
                                },
                            },
                            CustomData: map[string]components.CustomDataValue{
                                "someCustomField": components.CustomDataValue{},
                            },
                            ContactPerson: &components.Person{
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
                        },
                    },
                    Text: apiclientgo.String("Because its wavelengths are shorter, blue light is more strongly scattered than the longer-wavelength lights, red or green. Hence the result that when looking at the sky away from the direct incident sunlight, the human eye perceives the sky to be blue."),
                },
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
    if res.Announcement != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.CreateAnnouncementRequest](../../models/components/createannouncementrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateannouncementResponse](../../models/operations/createannouncementresponse.md), error**

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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Client.Announcements.Delete(ctx, components.DeleteAnnouncementRequest{
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.DeleteAnnouncementRequest](../../models/components/deleteannouncementrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.DeleteannouncementResponse](../../models/operations/deleteannouncementresponse.md), error**

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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Client.Announcements.Update(ctx, components.UpdateAnnouncementRequest{
        StartTime: types.MustTimeFromString("2025-07-28T19:04:48.565Z"),
        EndTime: types.MustTimeFromString("2024-10-16T10:52:42.015Z"),
        Title: "<value>",
        Body: &components.StructuredText{
            Text: "From https://en.wikipedia.org/wiki/Diffuse_sky_radiation, the sky is blue because blue light is more strongly scattered than longer-wavelength light.",
            StructuredList: []components.StructuredTextItem{
                components.StructuredTextItem{
                    Link: apiclientgo.String("https://en.wikipedia.org/wiki/Diffuse_sky_radiation"),
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
                            Author: &components.Person{
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
                                                FacetBucketSize: 488852,
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
                                                    StartIndex: 54062,
                                                },
                                                components.TextRange{
                                                    StartIndex: 896501,
                                                },
                                                components.TextRange{
                                                    StartIndex: 446863,
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
                                                FacetBucketSize: 249440,
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
                                                FacetBucketSize: 789275,
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
                            Owner: &components.Person{
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
                            MentionedPeople: []components.Person{
                                components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                },
                            },
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
                                    Attribution: &components.Person{
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
                                    UpdatedBy: &components.Person{
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
                                    Attribution: &components.Person{
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
                                    UpdatedBy: &components.Person{
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
                                    Attribution: &components.Person{
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
                                    UpdatedBy: &components.Person{
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
                                },
                            },
                            AssignedTo: &components.Person{
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
                            UpdatedBy: &components.Person{
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
                            Collections: []components.Collection{
                                components.Collection{
                                    Name: "<value>",
                                    Description: "daintily certainly yak surprised beyond blah intensely",
                                    AddedRoles: []components.UserRoleSpecification{
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
                                            Role: components.UserRoleEditor,
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
                                    ID: 249026,
                                    Creator: &components.Person{
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
                                    UpdatedBy: &components.Person{
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
                                    Items: []components.CollectionItem{
                                        components.CollectionItem{
                                            CollectionID: 784089,
                                            CreatedBy: &components.Person{
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
                                            Shortcut: &components.Shortcut{
                                                InputAlias: "<value>",
                                                CreatedBy: &components.Person{
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
                                                UpdatedBy: &components.Person{
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
                                                Roles: []components.UserRoleSpecification{
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
                                                        Role: components.UserRoleVerifier,
                                                    },
                                                },
                                            },
                                            ItemType: components.CollectionItemItemTypeText,
                                        },
                                        components.CollectionItem{
                                            CollectionID: 416023,
                                            CreatedBy: &components.Person{
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
                                            Shortcut: &components.Shortcut{
                                                InputAlias: "<value>",
                                                CreatedBy: &components.Person{
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
                                                UpdatedBy: &components.Person{
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
                                            },
                                            ItemType: components.CollectionItemItemTypeCollection,
                                        },
                                    },
                                },
                                components.Collection{
                                    Name: "<value>",
                                    Description: "misjudge scare cinema ouch weary euphonium",
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
                                    ID: 553539,
                                    Creator: &components.Person{
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
                                    UpdatedBy: &components.Person{
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
                                },
                                components.Collection{
                                    Name: "<value>",
                                    Description: "exotic fussy shadowy",
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
                                    ID: 890948,
                                    Creator: &components.Person{
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
                                    UpdatedBy: &components.Person{
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
                                },
                            },
                            Interactions: &components.DocumentInteractions{
                                Reacts: []components.Reaction{
                                    components.Reaction{
                                        Reactors: []components.Person{
                                            components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                            },
                                        },
                                    },
                                    components.Reaction{},
                                    components.Reaction{},
                                },
                                Shares: []components.Share{
                                    components.Share{
                                        NumDaysAgo: 178177,
                                        Sharer: &components.Person{
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
                                    },
                                },
                            },
                            Verification: &components.Verification{
                                State: components.StateUnverified,
                                Metadata: &components.VerificationMetadata{
                                    LastVerifier: &components.Person{
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
                                    Reminders: []components.Reminder{
                                        components.Reminder{
                                            Assignee: components.Person{
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
                                            Requestor: &components.Person{
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
                                            RemindAt: 691669,
                                        },
                                    },
                                    LastReminder: &components.Reminder{
                                        Assignee: components.Person{
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
                                        Requestor: &components.Person{
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
                                        RemindAt: 219050,
                                    },
                                    CandidateVerifiers: []components.Person{
                                        components.Person{
                                            Name: "George Clooney",
                                            ObfuscatedID: "abc123",
                                        },
                                    },
                                },
                            },
                            Shortcuts: []components.Shortcut{
                                components.Shortcut{
                                    InputAlias: "<value>",
                                    CreatedBy: &components.Person{
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
                                    UpdatedBy: &components.Person{
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
                                },
                            },
                            CustomData: map[string]components.CustomDataValue{
                                "someCustomField": components.CustomDataValue{},
                            },
                            ContactPerson: &components.Person{
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
                        },
                    },
                    Text: apiclientgo.String("Because its wavelengths are shorter, blue light is more strongly scattered than the longer-wavelength lights, red or green. Hence the result that when looking at the sky away from the direct incident sunlight, the human eye perceives the sky to be blue."),
                },
                components.StructuredTextItem{
                    Link: apiclientgo.String("https://en.wikipedia.org/wiki/Diffuse_sky_radiation"),
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
                            Author: &components.Person{
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
                            Owner: &components.Person{
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
                            Components: []string{
                                "Backend",
                                "Networking",
                            },
                            Status: apiclientgo.String("[\"Done\"]"),
                            AssignedTo: &components.Person{
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
                            UpdatedBy: &components.Person{
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
                            Interactions: &components.DocumentInteractions{},
                            Verification: &components.Verification{
                                State: components.StateVerified,
                                Metadata: &components.VerificationMetadata{
                                    LastVerifier: &components.Person{
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
                                    LastReminder: &components.Reminder{
                                        Assignee: components.Person{
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
                                        Requestor: &components.Person{
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
                                        RemindAt: 358043,
                                    },
                                },
                            },
                            CustomData: map[string]components.CustomDataValue{
                                "someCustomField": components.CustomDataValue{},
                            },
                            ContactPerson: &components.Person{
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
                        },
                    },
                    Text: apiclientgo.String("Because its wavelengths are shorter, blue light is more strongly scattered than the longer-wavelength lights, red or green. Hence the result that when looking at the sky away from the direct incident sunlight, the human eye perceives the sky to be blue."),
                },
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
        ID: 761625,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Announcement != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.UpdateAnnouncementRequest](../../models/components/updateannouncementrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateannouncementResponse](../../models/operations/updateannouncementresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |