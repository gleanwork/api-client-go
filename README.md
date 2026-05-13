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
  * [Experimental Features and Deprecation Testing](#experimental-features-and-deprecation-testing)
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
						Text: apiclientgo.Pointer("What are the company holidays this year?"),
					},
				},
			},
		},
	}, nil, nil)
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
						Text: apiclientgo.Pointer("What are the company holidays this year?"),
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
					Query: apiclientgo.Pointer("query"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/search?q=query",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionView,
				Params: &components.ActivityEventParams{
					Duration: apiclientgo.Pointer[int64](20),
					Referrer: apiclientgo.Pointer("https://example.com/document"),
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

### [Agents](docs/sdks/agents/README.md)

* [EditAgent](docs/sdks/agents/README.md#editagent) - Edit an agent

### [Authentication](docs/sdks/authentication/README.md)

* [Checkdatasourceauth](docs/sdks/authentication/README.md#checkdatasourceauth) - Check datasource authorization

### [Chat](docs/sdks/chat/README.md)

* [GetChatFile](docs/sdks/chat/README.md#getchatfile) - Download a chat file

### [Client.Activity](docs/sdks/activity/README.md)

* [Report](docs/sdks/activity/README.md#report) - Report document activity
* [Feedback](docs/sdks/activity/README.md#feedback) - Report client activity

### [Client.Agents](docs/sdks/clientagents/README.md)

* [Retrieve](docs/sdks/clientagents/README.md#retrieve) - Retrieve an agent
* [RetrieveSchemas](docs/sdks/clientagents/README.md#retrieveschemas) - List an agent's schemas
* [List](docs/sdks/clientagents/README.md#list) - Search agents
* [RunStream](docs/sdks/clientagents/README.md#runstream) - Create an agent run and stream the response
* [Run](docs/sdks/clientagents/README.md#run) - Create an agent run and wait for the response

### [Client.Announcements](docs/sdks/announcements/README.md)

* [Create](docs/sdks/announcements/README.md#create) - Create Announcement
* [Delete](docs/sdks/announcements/README.md#delete) - Delete Announcement
* [Update](docs/sdks/announcements/README.md#update) - Update Announcement

### [Client.Answers](docs/sdks/answers/README.md)

* [Create](docs/sdks/answers/README.md#create) - Create Answer
* [Delete](docs/sdks/answers/README.md#delete) - Delete Answer
* [Update](docs/sdks/answers/README.md#update) - Update Answer
* [Retrieve](docs/sdks/answers/README.md#retrieve) - Read Answer
* [~~List~~](docs/sdks/answers/README.md#list) - List Answers :warning: **Deprecated**

### [Client.Authentication](docs/sdks/clientauthentication/README.md)

* [CreateToken](docs/sdks/clientauthentication/README.md#createtoken) - Create authentication token

### [Client.Chat](docs/sdks/clientchat/README.md)

* [Create](docs/sdks/clientchat/README.md#create) - Chat
* [DeleteAll](docs/sdks/clientchat/README.md#deleteall) - Deletes all saved Chats owned by a user
* [Delete](docs/sdks/clientchat/README.md#delete) - Deletes saved Chats
* [Retrieve](docs/sdks/clientchat/README.md#retrieve) - Retrieves a Chat
* [List](docs/sdks/clientchat/README.md#list) - Retrieves all saved Chats
* [RetrieveApplication](docs/sdks/clientchat/README.md#retrieveapplication) - Gets the metadata for a custom Chat application
* [UploadFiles](docs/sdks/clientchat/README.md#uploadfiles) - Upload files for Chat
* [RetrieveFiles](docs/sdks/clientchat/README.md#retrievefiles) - Get files uploaded by a user for Chat
* [DeleteFiles](docs/sdks/clientchat/README.md#deletefiles) - Delete files uploaded by a user for chat
* [CreateStream](docs/sdks/clientchat/README.md#createstream) - Chat

### [Client.Collections](docs/sdks/collections/README.md)

* [AddItems](docs/sdks/collections/README.md#additems) - Add Collection item
* [Create](docs/sdks/collections/README.md#create) - Create Collection
* [Delete](docs/sdks/collections/README.md#delete) - Delete Collection
* [DeleteItem](docs/sdks/collections/README.md#deleteitem) - Delete Collection item
* [Update](docs/sdks/collections/README.md#update) - Update Collection
* [UpdateItem](docs/sdks/collections/README.md#updateitem) - Update Collection item
* [Retrieve](docs/sdks/collections/README.md#retrieve) - Read Collection
* [List](docs/sdks/collections/README.md#list) - List Collections

### [Client.Documents](docs/sdks/clientdocuments/README.md)

* [RetrievePermissions](docs/sdks/clientdocuments/README.md#retrievepermissions) - Read document permissions
* [Retrieve](docs/sdks/clientdocuments/README.md#retrieve) - Read documents
* [RetrieveByFacets](docs/sdks/clientdocuments/README.md#retrievebyfacets) - Read documents by facets
* [Summarize](docs/sdks/clientdocuments/README.md#summarize) - Summarize documents

### [Client.Entities](docs/sdks/cliententities/README.md)

* [List](docs/sdks/cliententities/README.md#list) - List entities
* [ReadPeople](docs/sdks/cliententities/README.md#readpeople) - Read people

### [Client.Governance.Data.Policies](docs/sdks/policies/README.md)

* [Retrieve](docs/sdks/policies/README.md#retrieve) - Gets specified policy
* [Update](docs/sdks/policies/README.md#update) - Updates an existing policy
* [List](docs/sdks/policies/README.md#list) - Lists policies
* [Create](docs/sdks/policies/README.md#create) - Creates new policy
* [Download](docs/sdks/policies/README.md#download) - Downloads violations CSV for policy

### [Client.Governance.Data.Reports](docs/sdks/reports/README.md)

* [Create](docs/sdks/reports/README.md#create) - Creates new one-time report
* [Download](docs/sdks/reports/README.md#download) - Downloads violations CSV for report
* [Status](docs/sdks/reports/README.md#status) - Fetches report run status

### [Client.Governance.Documents.Visibilityoverrides](docs/sdks/visibilityoverrides/README.md)

* [List](docs/sdks/visibilityoverrides/README.md#list) - Fetches documents visibility
* [Create](docs/sdks/visibilityoverrides/README.md#create) - Hide or unhide docs

### [Client.Insights](docs/sdks/insights/README.md)

* [Retrieve](docs/sdks/insights/README.md#retrieve) - Get insights

### [Client.Messages](docs/sdks/messages/README.md)

* [Retrieve](docs/sdks/messages/README.md#retrieve) - Read messages

### [Client.Pins](docs/sdks/pins/README.md)

* [Update](docs/sdks/pins/README.md#update) - Update pin
* [Retrieve](docs/sdks/pins/README.md#retrieve) - Read pin
* [List](docs/sdks/pins/README.md#list) - List pins
* [Create](docs/sdks/pins/README.md#create) - Create pin
* [Remove](docs/sdks/pins/README.md#remove) - Delete pin

### [Client.Search](docs/sdks/search/README.md)

* [QueryAsAdmin](docs/sdks/search/README.md#queryasadmin) - Search the index (admin)
* [Autocomplete](docs/sdks/search/README.md#autocomplete) - Autocomplete
* [RetrieveFeed](docs/sdks/search/README.md#retrievefeed) - Feed of documents and events
* [Recommendations](docs/sdks/search/README.md#recommendations) - Recommend documents
* [Query](docs/sdks/search/README.md#query) - Search

### [Client.Shortcuts](docs/sdks/clientshortcuts/README.md)

* [Create](docs/sdks/clientshortcuts/README.md#create) - Create shortcut
* [Delete](docs/sdks/clientshortcuts/README.md#delete) - Delete shortcut
* [Retrieve](docs/sdks/clientshortcuts/README.md#retrieve) - Read shortcut
* [List](docs/sdks/clientshortcuts/README.md#list) - List shortcuts
* [Update](docs/sdks/clientshortcuts/README.md#update) - Update shortcut

### [Client.Tools](docs/sdks/clienttools/README.md)

* [List](docs/sdks/clienttools/README.md#list) - List available tools
* [Run](docs/sdks/clienttools/README.md#run) - Execute the specified tool

### [Client.Verification](docs/sdks/verification/README.md)

* [AddReminder](docs/sdks/verification/README.md#addreminder) - Create verification
* [List](docs/sdks/verification/README.md#list) - List verifications
* [Verify](docs/sdks/verification/README.md#verify) - Update verification

### [Datasources](docs/sdks/datasources/README.md)

* [GetDatasourceInstanceConfiguration](docs/sdks/datasources/README.md#getdatasourceinstanceconfiguration) - Get datasource instance configuration
* [UpdateDatasourceInstanceConfiguration](docs/sdks/datasources/README.md#updatedatasourceinstanceconfiguration) - Update datasource instance configuration
* [GetDatasourceCredentialStatus](docs/sdks/datasources/README.md#getdatasourcecredentialstatus) - Get datasource instance credential status
* [RotateDatasourceCredentials](docs/sdks/datasources/README.md#rotatedatasourcecredentials) - Rotate datasource instance credentials

### [Entities](docs/sdks/entities/README.md)

* [GetPersonPhoto](docs/sdks/entities/README.md#getpersonphoto) - Get person photo

### [Governance](docs/sdks/governance/README.md)

* [Createfindingsexport](docs/sdks/governance/README.md#createfindingsexport) - Creates findings export
* [Listfindingsexports](docs/sdks/governance/README.md#listfindingsexports) - Lists findings exports
* [Downloadfindingsexport](docs/sdks/governance/README.md#downloadfindingsexport) - Downloads findings export
* [Deletefindingsexport](docs/sdks/governance/README.md#deletefindingsexport) - Deletes findings export

### [Indexing.Authentication](docs/sdks/indexingauthentication/README.md)

* [RotateToken](docs/sdks/indexingauthentication/README.md#rotatetoken) - Rotate token

### [Indexing.Datasource](docs/sdks/datasource/README.md)

* [Status](docs/sdks/datasource/README.md#status) - Beta: Get datasource status


### [Indexing.Datasources](docs/sdks/indexingdatasources/README.md)

* [Add](docs/sdks/indexingdatasources/README.md#add) - Add or update datasource
* [RetrieveConfig](docs/sdks/indexingdatasources/README.md#retrieveconfig) - Get datasource config

### [Indexing.Documents](docs/sdks/indexingdocuments/README.md)

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

### [Indexing.People](docs/sdks/people/README.md)

* [Debug](docs/sdks/people/README.md#debug) - Beta: Get user information

* [~~Count~~](docs/sdks/people/README.md#count) - Get user count :warning: **Deprecated**
* [Index](docs/sdks/people/README.md#index) - Index employee
* [~~BulkIndex~~](docs/sdks/people/README.md#bulkindex) - Bulk index employees :warning: **Deprecated**
* [ProcessAllEmployeesAndTeams](docs/sdks/people/README.md#processallemployeesandteams) - Schedules the processing of uploaded employees and teams
* [Delete](docs/sdks/people/README.md#delete) - Delete employee
* [IndexTeam](docs/sdks/people/README.md#indexteam) - Index team
* [DeleteTeam](docs/sdks/people/README.md#deleteteam) - Delete team
* [BulkIndexTeams](docs/sdks/people/README.md#bulkindexteams) - Bulk index teams

### [Indexing.Permissions](docs/sdks/permissions/README.md)

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

### [Indexing.Shortcuts](docs/sdks/indexingshortcuts/README.md)

* [BulkIndex](docs/sdks/indexingshortcuts/README.md#bulkindex) - Bulk index external shortcuts
* [Upload](docs/sdks/indexingshortcuts/README.md#upload) - Upload shortcuts

### [Tools](docs/sdks/tools/README.md)

* [GetActionPackAuthStatus](docs/sdks/tools/README.md#getactionpackauthstatus) - Get end-user authentication status for an action pack.
* [AuthorizeActionPack](docs/sdks/tools/README.md#authorizeactionpack) - Start the OAuth authorization flow for an action pack.

### [Troubleshooting](docs/sdks/troubleshooting/README.md)

* [PostAPIIndexV1DebugDatasourceDocumentEvents](docs/sdks/troubleshooting/README.md#postapiindexv1debugdatasourcedocumentevents) - Beta: Get document lifecycle events


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
					Query: apiclientgo.Pointer("query"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/search?q=query",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionView,
				Params: &components.ActivityEventParams{
					Duration: apiclientgo.Pointer[int64](20),
					Referrer: apiclientgo.Pointer("https://example.com/document"),
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
					Query: apiclientgo.Pointer("query"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/search?q=query",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionView,
				Params: &components.ActivityEventParams{
					Duration: apiclientgo.Pointer[int64](20),
					Referrer: apiclientgo.Pointer("https://example.com/document"),
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

For example, the `Retrieve` function may return the following errors:

| Error Type              | Status Code | Content Type     |
| ----------------------- | ----------- | ---------------- |
| apierrors.ErrorResponse | 404         | application/json |
| apierrors.APIError      | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/apierrors"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := apiclientgo.New(
		apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
	)

	res, err := s.Client.Agents.Retrieve(ctx, "<id>", nil, nil)
	if err != nil {

		var e *apierrors.ErrorResponse
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

| Variable   | Option                          | Default           | Description                                                                                            |
| ---------- | ------------------------------- | ----------------- | ------------------------------------------------------------------------------------------------------ |
| `instance` | `WithInstance(instance string)` | `"instance-name"` | The instance name (typically the email domain without the TLD) that determines the deployment backend. |

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
		apiclientgo.WithServerIndex(0),
		apiclientgo.WithInstance("instance-name"),
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
					Query: apiclientgo.Pointer("query"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/search?q=query",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionView,
				Params: &components.ActivityEventParams{
					Duration: apiclientgo.Pointer[int64](20),
					Referrer: apiclientgo.Pointer("https://example.com/document"),
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
					Query: apiclientgo.Pointer("query"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/search?q=query",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionView,
				Params: &components.ActivityEventParams{
					Duration: apiclientgo.Pointer[int64](20),
					Referrer: apiclientgo.Pointer("https://example.com/document"),
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

	"github.com/gleanwork/api-client-go"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = apiclientgo.New(apiclientgo.WithClient(httpClient))
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

## Experimental Features and Deprecation Testing

The SDK provides options to test upcoming API changes before they become the default behavior. This is useful for:

- **Testing experimental features** before they are generally available
- **Preparing for deprecations** by excluding deprecated endpoints ahead of their removal

### Configuration Options

You can configure these options either via environment variables or SDK constructor options:

#### Using Environment Variables

```bash
export X_GLEAN_EXCLUDE_DEPRECATED_AFTER="2026-10-15"
export X_GLEAN_INCLUDE_EXPERIMENTAL="true"
```

```go
package main

import (
	"context"
	"log"
	"os"

	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
)

func main() {
	ctx := context.Background()

	s := apiclientgo.New(
		apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
		apiclientgo.WithServerURL("https://mycompany-be.glean.com"),
	)

	res, err := s.Client.Search.Query(ctx, components.SearchRequest{
		Query: "test",
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Headers are automatically set based on environment variables
	log.Println(res)
}
```

#### Using SDK Constructor Options

```go
package main

import (
	"context"
	"log"
	"os"

	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
)

func main() {
	ctx := context.Background()

	s := apiclientgo.New(
		apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
		apiclientgo.WithServerURL("https://mycompany-be.glean.com"),
		apiclientgo.WithExcludeDeprecatedAfter("2026-10-15"),
		apiclientgo.WithIncludeExperimental(true),
	)

	res, err := s.Client.Search.Query(ctx, components.SearchRequest{
		Query: "test",
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
```

### Option Reference

| Option | Environment Variable | Type | Description |
| ------ | -------------------- | ---- | ----------- |
| `WithExcludeDeprecatedAfter` | `X_GLEAN_EXCLUDE_DEPRECATED_AFTER` | `string` (date) | Exclude API endpoints that will be deprecated after this date (format: `YYYY-MM-DD`). Use this to test your integration against upcoming deprecations. |
| `WithIncludeExperimental` | `X_GLEAN_INCLUDE_EXPERIMENTAL` | `bool` | When `true`, enables experimental API features that are not yet generally available. Use this to preview and test new functionality. |

> **Note:** Environment variables take precedence over SDK constructor options when both are set.

> **Warning:** Experimental features may change or be removed without notice. Do not rely on experimental features in production environments.

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
