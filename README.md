# api-client-go

The Glean Go SDK provides convenient access to the Glean REST API for Go 1.18+. It offers strongly typed request and response structs, context-based request handling, and uses the standard net/http package.
<!-- No Summary [summary] -->

## Unified SDK Architecture

This SDK combines both the Client and Indexing API namespaces into a single unified package:

- **Client API**: Used for search, retrieval, and end-user interactions with Glean content
- **Indexing API**: Used for indexing content, permissions, and other administrative operations

Each namespace has its own authentication requirements and access patterns. While they serve different purposes, having them in a single SDK provides a consistent developer experience across all Glean API interactions.

```go
// Example of accessing Client namespace
s := apiclientgo.New(
	apiclientgo.WithSecurity("client-token"),
)
res, err := s.Client.Search.Query(ctx, components.SearchRequest{
	Query: "search term",
})

// Example of accessing Indexing namespace 
s := apiclientgo.New(
	apiclientgo.WithSecurity("indexing-token"),
)
res, err := s.Indexing.Documents.Index(ctx, components.DocumentRequest{
	// document data
})
```

Remember that each namespace requires its own authentication token type as described in the [Authentication Methods](#authentication-methods) section.

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [api-client-go](#api-client-go)
  * [Unified SDK Architecture](#unified-sdk-architecture)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Special Types](#special-types)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/gleanwork/api-client-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example 1

```go
package main

import (
	"context"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := apiclientgo.New(
		apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
	)

	res, err := s.Client.Chat.Create(ctx, components.ChatRequest{
		Messages: []components.ChatMessage{
			components.ChatMessage{
				Fragments: []components.ChatMessageFragment{
					components.ChatMessageFragment{
						Text: apiclientgo.String("What are the company holidays this year?"),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ChatResponse != nil {
		// handle response
	}
}

```

### Example 2

```go
package main

import (
	"context"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := apiclientgo.New(
		apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
	)

	res, err := s.Client.Chat.CreateStream(ctx, components.ChatRequest{
		Messages: []components.ChatMessage{
			components.ChatMessage{
				Fragments: []components.ChatMessageFragment{
					components.ChatMessageFragment{
						Text: apiclientgo.String("What are the company holidays this year?"),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ChatRequestStream != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name       | Type | Scheme      | Environment Variable |
| ---------- | ---- | ----------- | -------------------- |
| `APIToken` | http | HTTP Bearer | `GLEAN_API_TOKEN`    |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/api-client-go/types"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := apiclientgo.New(
		apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
	)

	res, err := s.Client.Activity.Report(ctx, components.Activity{
		Events: []components.ActivityEvent{
			components.ActivityEvent{
				Action:    components.ActivityEventActionHistoricalView,
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionSearch,
				Params: &components.ActivityEventParams{
					Query: apiclientgo.String("query"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/search?q=query",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionView,
				Params: &components.ActivityEventParams{
					Duration: apiclientgo.Int64(20),
					Referrer: apiclientgo.String("https://example.com/document"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/",
			},
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
<!-- End Authentication [security] -->

### Authentication Methods

Glean supports different authentication methods depending on which API namespace you're using:

#### Client Namespace

The Client namespace supports two authentication methods:

1. **Manually Provisioned API Tokens**
   - Can be created by an Admin or a user with the API Token Creator role
   - Used for server-to-server integrations

2. **OAuth**
   - Requires OAuth setup to be completed by an Admin
   - Used for user-based authentication flows

#### Indexing Namespace

The Indexing namespace supports only one authentication method:

1. **Manually Provisioned API Tokens**
   - Can be created by an Admin or a user with the API Token Creator role
   - Used for secure document indexing operations

> [!IMPORTANT]
> Client tokens **will not work** for Indexing operations, and Indexing tokens **will not work** for Client operations. You must use the appropriate token type for the namespace you're accessing.

For more information on obtaining the appropriate token type, please contact your Glean administrator.

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Client](docs/sdks/client/README.md)


#### [Client.Activity](docs/sdks/activity/README.md)

* [Report](docs/sdks/activity/README.md#report) - Report document activity
* [Feedback](docs/sdks/activity/README.md#feedback) - Report client activity

#### [Client.Agents](docs/sdks/agents/README.md)

* [Run](docs/sdks/agents/README.md#run) - Runs an Agent.
* [List](docs/sdks/agents/README.md#list) - Lists all agents.
* [RetrieveInputs](docs/sdks/agents/README.md#retrieveinputs) - Gets the inputs to an agent.

#### [Client.Announcements](docs/sdks/announcements/README.md)

* [Create](docs/sdks/announcements/README.md#create) - Create Announcement
* [Delete](docs/sdks/announcements/README.md#delete) - Delete Announcement
* [Update](docs/sdks/announcements/README.md#update) - Update Announcement

#### [Client.Answers](docs/sdks/answers/README.md)

* [Create](docs/sdks/answers/README.md#create) - Create Answer
* [Delete](docs/sdks/answers/README.md#delete) - Delete Answer
* [Update](docs/sdks/answers/README.md#update) - Update Answer
* [Retrieve](docs/sdks/answers/README.md#retrieve) - Read Answer
* [List](docs/sdks/answers/README.md#list) - List Answers

#### [Client.Authentication](docs/sdks/clientauthentication/README.md)

* [CreateToken](docs/sdks/clientauthentication/README.md#createtoken) - Create authentication token

#### [Client.Chat](docs/sdks/chat/README.md)

* [Create](docs/sdks/chat/README.md#create) - Chat
* [DeleteAll](docs/sdks/chat/README.md#deleteall) - Deletes all saved Chats owned by a user
* [Delete](docs/sdks/chat/README.md#delete) - Deletes saved Chats
* [Retrieve](docs/sdks/chat/README.md#retrieve) - Retrieves a Chat
* [List](docs/sdks/chat/README.md#list) - Retrieves all saved Chats
* [RetrieveApplication](docs/sdks/chat/README.md#retrieveapplication) - Gets the metadata for a custom Chat application
* [UploadFiles](docs/sdks/chat/README.md#uploadfiles) - Upload files for Chat.
* [RetrieveFiles](docs/sdks/chat/README.md#retrievefiles) - Get files uploaded by a user for Chat.
* [DeleteFiles](docs/sdks/chat/README.md#deletefiles) - Delete files uploaded by a user for chat.
* [CreateStream](docs/sdks/chat/README.md#createstream) - Chat

#### [Client.Collections](docs/sdks/collections/README.md)

* [AddItems](docs/sdks/collections/README.md#additems) - Add Collection item
* [Create](docs/sdks/collections/README.md#create) - Create Collection
* [Delete](docs/sdks/collections/README.md#delete) - Delete Collection
* [DeleteItem](docs/sdks/collections/README.md#deleteitem) - Delete Collection item
* [Update](docs/sdks/collections/README.md#update) - Update Collection
* [UpdateItem](docs/sdks/collections/README.md#updateitem) - Update Collection item
* [Retrieve](docs/sdks/collections/README.md#retrieve) - Read Collection
* [List](docs/sdks/collections/README.md#list) - List Collections

#### [Client.Documents](docs/sdks/clientdocuments/README.md)

* [RetrievePermissions](docs/sdks/clientdocuments/README.md#retrievepermissions) - Read document permissions
* [Retrieve](docs/sdks/clientdocuments/README.md#retrieve) - Read documents
* [RetrieveByFacets](docs/sdks/clientdocuments/README.md#retrievebyfacets) - Read documents by facets
* [Summarize](docs/sdks/clientdocuments/README.md#summarize) - Summarize documents

#### [Client.Entities](docs/sdks/entities/README.md)

* [List](docs/sdks/entities/README.md#list) - List entities
* [ReadPeople](docs/sdks/entities/README.md#readpeople) - Read people

#### [Client.Insights](docs/sdks/insights/README.md)

* [Retrieve](docs/sdks/insights/README.md#retrieve) - Read insights

#### [Client.Messages](docs/sdks/messages/README.md)

* [Retrieve](docs/sdks/messages/README.md#retrieve) - Read messages

#### [Client.Pins](docs/sdks/pins/README.md)

* [Update](docs/sdks/pins/README.md#update) - Update pin
* [Retrieve](docs/sdks/pins/README.md#retrieve) - Read pin
* [List](docs/sdks/pins/README.md#list) - List pins
* [Create](docs/sdks/pins/README.md#create) - Create pin
* [Remove](docs/sdks/pins/README.md#remove) - Delete pin

#### [Client.Search](docs/sdks/search/README.md)

* [QueryAsAdmin](docs/sdks/search/README.md#queryasadmin) - Search the index (admin)
* [Autocomplete](docs/sdks/search/README.md#autocomplete) - Autocomplete
* [RetrieveFeed](docs/sdks/search/README.md#retrievefeed) - Feed of documents and events
* [Recommendations](docs/sdks/search/README.md#recommendations) - Recommend documents
* [Query](docs/sdks/search/README.md#query) - Search

#### [Client.Shortcuts](docs/sdks/clientshortcuts/README.md)

* [Create](docs/sdks/clientshortcuts/README.md#create) - Create shortcut
* [Delete](docs/sdks/clientshortcuts/README.md#delete) - Delete shortcut
* [Retrieve](docs/sdks/clientshortcuts/README.md#retrieve) - Read shortcut
* [List](docs/sdks/clientshortcuts/README.md#list) - List shortcuts
* [Update](docs/sdks/clientshortcuts/README.md#update) - Update shortcut

#### [Client.Verification](docs/sdks/verification/README.md)

* [AddReminder](docs/sdks/verification/README.md#addreminder) - Create verification
* [List](docs/sdks/verification/README.md#list) - List verifications
* [Verify](docs/sdks/verification/README.md#verify) - Update verification


### [Indexing](docs/sdks/indexing/README.md)


#### [Indexing.Authentication](docs/sdks/indexingauthentication/README.md)

* [RotateToken](docs/sdks/indexingauthentication/README.md#rotatetoken) - Rotate token

#### [Indexing.Datasource](docs/sdks/datasource/README.md)

* [Status](docs/sdks/datasource/README.md#status) - Beta: Get datasource status


#### [Indexing.Datasources](docs/sdks/datasources/README.md)

* [Add](docs/sdks/datasources/README.md#add) - Add or update datasource
* [RetrieveConfig](docs/sdks/datasources/README.md#retrieveconfig) - Get datasource config

#### [Indexing.Documents](docs/sdks/indexingdocuments/README.md)

* [AddOrUpdate](docs/sdks/indexingdocuments/README.md#addorupdate) - Index document
* [Index](docs/sdks/indexingdocuments/README.md#index) - Index documents
* [BulkIndex](docs/sdks/indexingdocuments/README.md#bulkindex) - Bulk index documents
* [ProcessAll](docs/sdks/indexingdocuments/README.md#processall) - Schedules the processing of uploaded documents
* [Delete](docs/sdks/indexingdocuments/README.md#delete) - Delete document
* [Debug](docs/sdks/indexingdocuments/README.md#debug) - Beta: Get document information

* [DebugMany](docs/sdks/indexingdocuments/README.md#debugmany) - Beta: Get information of a batch of documents

* [CheckAccess](docs/sdks/indexingdocuments/README.md#checkaccess) - Check document access
* [~~Status~~](docs/sdks/indexingdocuments/README.md#status) - Get document upload and indexing status :warning: **Deprecated**
* [~~Count~~](docs/sdks/indexingdocuments/README.md#count) - Get document count :warning: **Deprecated**

#### [Indexing.People](docs/sdks/people/README.md)

* [Debug](docs/sdks/people/README.md#debug) - Beta: Get user information

* [~~Count~~](docs/sdks/people/README.md#count) - Get user count :warning: **Deprecated**
* [Index](docs/sdks/people/README.md#index) - Index employee
* [BulkIndex](docs/sdks/people/README.md#bulkindex) - Bulk index employees
* [ProcessAllEmployeesAndTeams](docs/sdks/people/README.md#processallemployeesandteams) - Schedules the processing of uploaded employees and teams
* [Delete](docs/sdks/people/README.md#delete) - Delete employee
* [IndexTeam](docs/sdks/people/README.md#indexteam) - Index team
* [DeleteTeam](docs/sdks/people/README.md#deleteteam) - Delete team
* [BulkIndexTeams](docs/sdks/people/README.md#bulkindexteams) - Bulk index teams

#### [Indexing.Permissions](docs/sdks/permissions/README.md)

* [UpdatePermissions](docs/sdks/permissions/README.md#updatepermissions) - Update document permissions
* [IndexUser](docs/sdks/permissions/README.md#indexuser) - Index user
* [BulkIndexUsers](docs/sdks/permissions/README.md#bulkindexusers) - Bulk index users
* [IndexGroup](docs/sdks/permissions/README.md#indexgroup) - Index group
* [BulkIndexGroups](docs/sdks/permissions/README.md#bulkindexgroups) - Bulk index groups
* [IndexMembership](docs/sdks/permissions/README.md#indexmembership) - Index membership
* [BulkIndexMemberships](docs/sdks/permissions/README.md#bulkindexmemberships) - Bulk index memberships for a group
* [ProcessMemberships](docs/sdks/permissions/README.md#processmemberships) - Schedules the processing of group memberships
* [DeleteUser](docs/sdks/permissions/README.md#deleteuser) - Delete user
* [DeleteGroup](docs/sdks/permissions/README.md#deletegroup) - Delete group
* [DeleteMembership](docs/sdks/permissions/README.md#deletemembership) - Delete membership
* [AuthorizeBetaUsers](docs/sdks/permissions/README.md#authorizebetausers) - Beta users

#### [Indexing.Shortcuts](docs/sdks/indexingshortcuts/README.md)

* [BulkIndex](docs/sdks/indexingshortcuts/README.md#bulkindex) - Bulk index external shortcuts
* [Upload](docs/sdks/indexingshortcuts/README.md#upload) - Upload shortcuts

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/api-client-go/retry"
	"github.com/gleanwork/api-client-go/types"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := apiclientgo.New(
		apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
	)

	res, err := s.Client.Activity.Report(ctx, components.Activity{
		Events: []components.ActivityEvent{
			components.ActivityEvent{
				Action:    components.ActivityEventActionHistoricalView,
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionSearch,
				Params: &components.ActivityEventParams{
					Query: apiclientgo.String("query"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/search?q=query",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionView,
				Params: &components.ActivityEventParams{
					Duration: apiclientgo.Int64(20),
					Referrer: apiclientgo.String("https://example.com/document"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/",
			},
		},
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/api-client-go/retry"
	"github.com/gleanwork/api-client-go/types"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := apiclientgo.New(
		apiclientgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
	)

	res, err := s.Client.Activity.Report(ctx, components.Activity{
		Events: []components.ActivityEvent{
			components.ActivityEvent{
				Action:    components.ActivityEventActionHistoricalView,
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionSearch,
				Params: &components.ActivityEventParams{
					Query: apiclientgo.String("query"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/search?q=query",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionView,
				Params: &components.ActivityEventParams{
					Duration: apiclientgo.Int64(20),
					Referrer: apiclientgo.String("https://example.com/document"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/",
			},
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Create` function may return the following errors:

| Error Type                | Status Code | Content Type     |
| ------------------------- | ----------- | ---------------- |
| apierrors.CollectionError | 422         | application/json |
| apierrors.APIError        | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/apierrors"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/api-client-go/types"
	"log"
	"os"
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
					Name:         "George Clooney",
					ObfuscatedID: "abc123",
					RelatedDocuments: []components.RelatedDocuments{
						components.RelatedDocuments{
							QuerySuggestion: &components.QuerySuggestion{
								Query: "app:github type:pull author:mortimer",
								SearchProviderInfo: &components.SearchProviderInfo{
									Name:                  apiclientgo.String("Google"),
									SearchLinkURLTemplate: apiclientgo.String("https://www.google.com/search?q={query}&hl=en"),
								},
								Label:      apiclientgo.String("Mortimer's PRs"),
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
													Value:        apiclientgo.String("Spreadsheet"),
													RelationType: components.RelationTypeEquals.ToPointer(),
												},
												components.FacetFilterValue{
													Value:        apiclientgo.String("Presentation"),
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
															Value:        apiclientgo.String("Spreadsheet"),
															RelationType: components.RelationTypeEquals.ToPointer(),
														},
														components.FacetFilterValue{
															Value:        apiclientgo.String("Presentation"),
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
															Value:        apiclientgo.String("Spreadsheet"),
															RelationType: components.RelationTypeEquals.ToPointer(),
														},
														components.FacetFilterValue{
															Value:        apiclientgo.String("Presentation"),
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
											Datasource:  "gmail",
											Scope:       apiclientgo.String("email profile https://www.googleapis.com/auth/gmail.readonly"),
											TokenType:   apiclientgo.String("Bearer"),
											AuthUser:    apiclientgo.String("1"),
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
												Container:  apiclientgo.String("container"),
												ParentID:   apiclientgo.String("JIRA_EN-1337"),
												MimeType:   apiclientgo.String("mimeType"),
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
																		Value:        apiclientgo.String("Spreadsheet"),
																		RelationType: components.RelationTypeEquals.ToPointer(),
																	},
																	components.FacetFilterValue{
																		Value:        apiclientgo.String("Presentation"),
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
																		Value:        apiclientgo.String("Spreadsheet"),
																		RelationType: components.RelationTypeEquals.ToPointer(),
																	},
																	components.FacetFilterValue{
																		Value:        apiclientgo.String("Presentation"),
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
																		Value:        apiclientgo.String("Spreadsheet"),
																		RelationType: components.RelationTypeEquals.ToPointer(),
																	},
																	components.FacetFilterValue{
																		Value:        apiclientgo.String("Presentation"),
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
														Name:        "<value>",
														Description: "fumigate convection though zowie",
														AudienceFilters: []components.FacetFilter{
															components.FacetFilter{
																FieldName: apiclientgo.String("type"),
																Values: []components.FacetFilterValue{
																	components.FacetFilterValue{
																		Value:        apiclientgo.String("Spreadsheet"),
																		RelationType: components.RelationTypeEquals.ToPointer(),
																	},
																	components.FacetFilterValue{
																		Value:        apiclientgo.String("Presentation"),
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
																	Name:         "George Clooney",
																	ObfuscatedID: "abc123",
																},
																RemindAt: 491427,
															},
														},
														LastReminder: &components.Reminder{
															Assignee: components.Person{
																Name:         "George Clooney",
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
									Title:        apiclientgo.String("title"),
									URL:          "https://example.com/foo/bar",
									NativeAppURL: apiclientgo.String("slack://foo/bar"),
									Snippets: []components.SearchResultSnippet{
										components.SearchResultSnippet{
											Snippet:  "snippet",
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
									Name:                  apiclientgo.String("Google"),
									SearchLinkURLTemplate: apiclientgo.String("https://www.google.com/search?q={query}&hl=en"),
								},
								Label:      apiclientgo.String("Mortimer's PRs"),
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
													Value:        apiclientgo.String("Spreadsheet"),
													RelationType: components.RelationTypeEquals.ToPointer(),
												},
												components.FacetFilterValue{
													Value:        apiclientgo.String("Presentation"),
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
															Value:        apiclientgo.String("Spreadsheet"),
															RelationType: components.RelationTypeEquals.ToPointer(),
														},
														components.FacetFilterValue{
															Value:        apiclientgo.String("Presentation"),
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
															Value:        apiclientgo.String("Spreadsheet"),
															RelationType: components.RelationTypeEquals.ToPointer(),
														},
														components.FacetFilterValue{
															Value:        apiclientgo.String("Presentation"),
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
															Value:        apiclientgo.String("Spreadsheet"),
															RelationType: components.RelationTypeEquals.ToPointer(),
														},
														components.FacetFilterValue{
															Value:        apiclientgo.String("Presentation"),
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
											Datasource:  "gmail",
											Scope:       apiclientgo.String("email profile https://www.googleapis.com/auth/gmail.readonly"),
											TokenType:   apiclientgo.String("Bearer"),
											AuthUser:    apiclientgo.String("1"),
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
						Type:       components.PersonMetadataTypeFullTime.ToPointer(),
						Title:      apiclientgo.String("Actor"),
						Department: apiclientgo.String("Movies"),
						Email:      apiclientgo.String("george@example.com"),
						Location:   apiclientgo.String("Hollywood, CA"),
						Phone:      apiclientgo.String("6505551234"),
						PhotoURL:   apiclientgo.String("https://example.com/george.jpg"),
						StartDate:  types.MustNewDateFromString("2000-01-23"),
						DatasourceProfile: []components.DatasourceProfile{
							components.DatasourceProfile{
								Datasource: "github",
								Handle:     "<value>",
							},
						},
						QuerySuggestions: &components.QuerySuggestionList{
							Suggestions: []components.QuerySuggestion{
								components.QuerySuggestion{
									Query:      "app:github type:pull author:mortimer",
									Label:      apiclientgo.String("Mortimer's PRs"),
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
								Label:  "<value>",
								Values: []components.CustomFieldValue{},
							},
						},
						Badges: []components.Badge{
							components.Badge{
								Key:         apiclientgo.String("deployment_name_new_hire"),
								DisplayName: apiclientgo.String("New hire"),
								IconConfig: &components.IconConfig{
									Color:    apiclientgo.String("#343CED"),
									Key:      apiclientgo.String("person_icon"),
									IconType: components.IconTypeGlyph.ToPointer(),
									Name:     apiclientgo.String("user"),
								},
							},
						},
					},
				},
				Role: components.UserRoleOwner,
			},
			components.UserRoleSpecification{
				Person: &components.Person{
					Name:         "George Clooney",
					ObfuscatedID: "abc123",
					Metadata: &components.PersonMetadata{
						Type:       components.PersonMetadataTypeFullTime.ToPointer(),
						Title:      apiclientgo.String("Actor"),
						Department: apiclientgo.String("Movies"),
						Email:      apiclientgo.String("george@example.com"),
						Location:   apiclientgo.String("Hollywood, CA"),
						Phone:      apiclientgo.String("6505551234"),
						PhotoURL:   apiclientgo.String("https://example.com/george.jpg"),
						StartDate:  types.MustNewDateFromString("2000-01-23"),
						DatasourceProfile: []components.DatasourceProfile{
							components.DatasourceProfile{
								Datasource: "github",
								Handle:     "<value>",
							},
						},
						QuerySuggestions: &components.QuerySuggestionList{},
						InviteInfo:       &components.InviteInfo{},
						Badges: []components.Badge{
							components.Badge{
								Key:         apiclientgo.String("deployment_name_new_hire"),
								DisplayName: apiclientgo.String("New hire"),
								IconConfig: &components.IconConfig{
									Color:    apiclientgo.String("#343CED"),
									Key:      apiclientgo.String("person_icon"),
									IconType: components.IconTypeGlyph.ToPointer(),
									Name:     apiclientgo.String("user"),
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
					Name:         "George Clooney",
					ObfuscatedID: "abc123",
					Metadata: &components.PersonMetadata{
						Type:       components.PersonMetadataTypeFullTime.ToPointer(),
						Title:      apiclientgo.String("Actor"),
						Department: apiclientgo.String("Movies"),
						Email:      apiclientgo.String("george@example.com"),
						Location:   apiclientgo.String("Hollywood, CA"),
						Phone:      apiclientgo.String("6505551234"),
						PhotoURL:   apiclientgo.String("https://example.com/george.jpg"),
						StartDate:  types.MustNewDateFromString("2000-01-23"),
						DatasourceProfile: []components.DatasourceProfile{
							components.DatasourceProfile{
								Datasource: "github",
								Handle:     "<value>",
							},
						},
						QuerySuggestions: &components.QuerySuggestionList{},
						InviteInfo:       &components.InviteInfo{},
						Badges: []components.Badge{
							components.Badge{
								Key:         apiclientgo.String("deployment_name_new_hire"),
								DisplayName: apiclientgo.String("New hire"),
								IconConfig: &components.IconConfig{
									Color:    apiclientgo.String("#343CED"),
									Key:      apiclientgo.String("person_icon"),
									IconType: components.IconTypeGlyph.ToPointer(),
									Name:     apiclientgo.String("user"),
								},
							},
						},
					},
				},
				Role: components.UserRoleVerifier,
			},
			components.UserRoleSpecification{
				Person: &components.Person{
					Name:         "George Clooney",
					ObfuscatedID: "abc123",
					Metadata: &components.PersonMetadata{
						Type:       components.PersonMetadataTypeFullTime.ToPointer(),
						Title:      apiclientgo.String("Actor"),
						Department: apiclientgo.String("Movies"),
						Email:      apiclientgo.String("george@example.com"),
						Location:   apiclientgo.String("Hollywood, CA"),
						Phone:      apiclientgo.String("6505551234"),
						PhotoURL:   apiclientgo.String("https://example.com/george.jpg"),
						StartDate:  types.MustNewDateFromString("2000-01-23"),
						DatasourceProfile: []components.DatasourceProfile{
							components.DatasourceProfile{
								Datasource: "github",
								Handle:     "<value>",
							},
							components.DatasourceProfile{
								Datasource: "github",
								Handle:     "<value>",
							},
							components.DatasourceProfile{
								Datasource: "github",
								Handle:     "<value>",
							},
						},
						QuerySuggestions: &components.QuerySuggestionList{},
						InviteInfo:       &components.InviteInfo{},
						Badges: []components.Badge{
							components.Badge{
								Key:         apiclientgo.String("deployment_name_new_hire"),
								DisplayName: apiclientgo.String("New hire"),
								IconConfig: &components.IconConfig{
									Color:    apiclientgo.String("#343CED"),
									Key:      apiclientgo.String("person_icon"),
									IconType: components.IconTypeGlyph.ToPointer(),
									Name:     apiclientgo.String("user"),
								},
							},
						},
					},
				},
				Role: components.UserRoleAnswerModerator,
			},
			components.UserRoleSpecification{
				Person: &components.Person{
					Name:         "George Clooney",
					ObfuscatedID: "abc123",
					Metadata: &components.PersonMetadata{
						Type:       components.PersonMetadataTypeFullTime.ToPointer(),
						Title:      apiclientgo.String("Actor"),
						Department: apiclientgo.String("Movies"),
						Email:      apiclientgo.String("george@example.com"),
						Location:   apiclientgo.String("Hollywood, CA"),
						Phone:      apiclientgo.String("6505551234"),
						PhotoURL:   apiclientgo.String("https://example.com/george.jpg"),
						StartDate:  types.MustNewDateFromString("2000-01-23"),
						DatasourceProfile: []components.DatasourceProfile{
							components.DatasourceProfile{
								Datasource: "github",
								Handle:     "<value>",
							},
						},
						QuerySuggestions: &components.QuerySuggestionList{},
						InviteInfo:       &components.InviteInfo{},
						Badges: []components.Badge{
							components.Badge{
								Key:         apiclientgo.String("deployment_name_new_hire"),
								DisplayName: apiclientgo.String("New hire"),
								IconConfig: &components.IconConfig{
									Color:    apiclientgo.String("#343CED"),
									Key:      apiclientgo.String("person_icon"),
									IconType: components.IconTypeGlyph.ToPointer(),
									Name:     apiclientgo.String("user"),
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
						Value:        apiclientgo.String("Spreadsheet"),
						RelationType: components.RelationTypeEquals.ToPointer(),
					},
					components.FacetFilterValue{
						Value:        apiclientgo.String("Presentation"),
						RelationType: components.RelationTypeEquals.ToPointer(),
					},
				},
			},
		},
	})
	if err != nil {

		var e *apierrors.CollectionError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Server Variables

The default server `https://{instance}-be.glean.com` contains variables and is set to `https://instance-name-be.glean.com` by default. To override default values, the following options are available when initializing the SDK client instance:

| Variable   | Option                          | Default           | Description                                                                                                  |
| ---------- | ------------------------------- | ----------------- | ------------------------------------------------------------------------------------------------------------ |
| `instance` | `WithInstance(instance string)` | `"instance-name"` | The instance name (typically the email domain without the extension) that determines the deployment backend. |

#### Example

```go
package main

import (
	"context"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/api-client-go/types"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := apiclientgo.New(
		apiclientgo.WithInstance("<value>"),
		apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
	)

	res, err := s.Client.Activity.Report(ctx, components.Activity{
		Events: []components.ActivityEvent{
			components.ActivityEvent{
				Action:    components.ActivityEventActionHistoricalView,
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionSearch,
				Params: &components.ActivityEventParams{
					Query: apiclientgo.String("query"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/search?q=query",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionView,
				Params: &components.ActivityEventParams{
					Duration: apiclientgo.Int64(20),
					Referrer: apiclientgo.String("https://example.com/document"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/",
			},
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

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/api-client-go/types"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := apiclientgo.New(
		apiclientgo.WithServerURL("https://instance-name-be.glean.com"),
		apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
	)

	res, err := s.Client.Activity.Report(ctx, components.Activity{
		Events: []components.ActivityEvent{
			components.ActivityEvent{
				Action:    components.ActivityEventActionHistoricalView,
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionSearch,
				Params: &components.ActivityEventParams{
					Query: apiclientgo.String("query"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/search?q=query",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionView,
				Params: &components.ActivityEventParams{
					Duration: apiclientgo.Int64(20),
					Referrer: apiclientgo.String("https://example.com/document"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/",
			},
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/gleanwork/api-client-go&utm_campaign=go)
