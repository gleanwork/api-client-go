# github.com/gleanwork/api-client-go

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/gleanwork/api-client-go* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/gleanwork/api-client-go&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/glean-el2/sdk). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Glean API: # Introduction
In addition to the data sources that Glean has built-in support for, Glean also provides a REST API that enables customers to put arbitrary content in the search index. This is useful, for example, for doing permissions-aware search over content in internal tools that reside on-prem as well as for searching over applications that Glean does not currently support first class. In addition these APIs allow the customer to push organization data (people info, organization structure etc) into Glean.

# Usage guidelines
This API is evolving fast. Glean will provide advance notice of any planned backwards incompatible changes along
with a 6-month sunset period for anything that requires developers to adopt the new versions.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/gleanwork/api-client-go](#githubcomgleanworkapi-client-go)
* [Usage guidelines](#usage-guidelines)
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

### Example

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
		apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
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
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type | Scheme      | Environment Variable |
| ------------ | ---- | ----------- | -------------------- |
| `BearerAuth` | http | HTTP Bearer | `GLEAN_BEARER_AUTH`  |

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
		apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Agents](docs/sdks/agents/README.md)

* [Runagent](docs/sdks/agents/README.md#runagent) - Runs an Agent.
* [Listagents](docs/sdks/agents/README.md#listagents) - Lists all agents.
* [Getagentinputs](docs/sdks/agents/README.md#getagentinputs) - Gets the inputs to an agent.

### [Client](docs/sdks/client/README.md)


#### [Client.Activities](docs/sdks/activities/README.md)

* [ReportActivity](docs/sdks/activities/README.md#reportactivity) - Report client activity

#### [Client.Activity](docs/sdks/activity/README.md)

* [Report](docs/sdks/activity/README.md#report) - Report document activity

#### [Client.Announcements](docs/sdks/announcements/README.md)

* [Create](docs/sdks/announcements/README.md#create) - Create Announcement
* [CreateDraft](docs/sdks/announcements/README.md#createdraft) - Create draft Announcement
* [Delete](docs/sdks/announcements/README.md#delete) - Delete Announcement
* [DeleteDraft](docs/sdks/announcements/README.md#deletedraft) - Delete draft Announcement
* [Get](docs/sdks/announcements/README.md#get) - Read Announcement
* [GetDraft](docs/sdks/announcements/README.md#getdraft) - Read draft Announcement
* [List](docs/sdks/announcements/README.md#list) - List Announcements
* [Preview](docs/sdks/announcements/README.md#preview) - Preview Announcement
* [PreviewDraft](docs/sdks/announcements/README.md#previewdraft) - Preview draft Announcement
* [Publish](docs/sdks/announcements/README.md#publish) - Publish draft Announcement
* [Unpublish](docs/sdks/announcements/README.md#unpublish) - Unpublish Announcement
* [Update](docs/sdks/announcements/README.md#update) - Update Announcement
* [UpdateDraft](docs/sdks/announcements/README.md#updatedraft) - Update draft Announcement

#### [Client.Answers](docs/sdks/answers/README.md)

* [Create](docs/sdks/answers/README.md#create) - Create Answer
* [Delete](docs/sdks/answers/README.md#delete) - Delete Answer
* [Edit](docs/sdks/answers/README.md#edit) - Update Answer
* [Get](docs/sdks/answers/README.md#get) - Read Answer
* [List](docs/sdks/answers/README.md#list) - List Answers
* [Preview](docs/sdks/answers/README.md#preview) - Preview Answer
* [PreviewDraft](docs/sdks/answers/README.md#previewdraft) - Preview draft Answer
* [UpdateLikes](docs/sdks/answers/README.md#updatelikes) - Update Answer likes
* [~~CreateBoard~~](docs/sdks/answers/README.md#createboard) - Create Answer Board :warning: **Deprecated**
* [~~DeleteBoard~~](docs/sdks/answers/README.md#deleteboard) - Delete Answer Board :warning: **Deprecated**
* [~~UpdateBoard~~](docs/sdks/answers/README.md#updateboard) - Update Answer Board :warning: **Deprecated**
* [~~GetBoard~~](docs/sdks/answers/README.md#getboard) - Read Answer Board :warning: **Deprecated**
* [~~ListBoards~~](docs/sdks/answers/README.md#listboards) - List Answer Boards :warning: **Deprecated**

#### [Client.Authentication](docs/sdks/clientauthentication/README.md)

* [CreateAnonymousToken](docs/sdks/clientauthentication/README.md#createanonymoustoken) - Create anonymous token
* [CreateToken](docs/sdks/clientauthentication/README.md#createtoken) - Create authentication token

#### [Client.Calendar](docs/sdks/calendar/README.md)

* [GetEvents](docs/sdks/calendar/README.md#getevents) - Read events

#### [Client.Chat](docs/sdks/chat/README.md)

* [Ask](docs/sdks/chat/README.md#ask) - Detect and answer questions
* [Start](docs/sdks/chat/README.md#start) - Chat
* [DeleteAll](docs/sdks/chat/README.md#deleteall) - Deletes all saved Chats owned by a user
* [Delete](docs/sdks/chat/README.md#delete) - Deletes saved Chats
* [Get](docs/sdks/chat/README.md#get) - Retrieves a Chat
* [List](docs/sdks/chat/README.md#list) - Retrieves all saved Chats
* [GetApplication](docs/sdks/chat/README.md#getapplication) - Gets the metadata for a custom Chat application
* [UploadFiles](docs/sdks/chat/README.md#uploadfiles) - Upload files for Chat.
* [GetFiles](docs/sdks/chat/README.md#getfiles) - Get files uploaded by a user for Chat.
* [DeleteFiles](docs/sdks/chat/README.md#deletefiles) - Delete files uploaded by a user for chat.

#### [Client.Collections](docs/sdks/collections/README.md)

* [AddItems](docs/sdks/collections/README.md#additems) - Add Collection item
* [Create](docs/sdks/collections/README.md#create) - Create Collection
* [Delete](docs/sdks/collections/README.md#delete) - Delete Collection
* [DeleteItem](docs/sdks/collections/README.md#deleteitem) - Delete Collection item
* [Update](docs/sdks/collections/README.md#update) - Update Collection
* [EditItem](docs/sdks/collections/README.md#edititem) - Update Collection item
* [Edit](docs/sdks/collections/README.md#edit) - Update document Collections
* [Get](docs/sdks/collections/README.md#get) - Read Collection
* [List](docs/sdks/collections/README.md#list) - List Collections
* [MoveItem](docs/sdks/collections/README.md#moveitem) - Move Collection item
* [Pin](docs/sdks/collections/README.md#pin) - Pin Collection

#### [Client.DisplayableLists](docs/sdks/displayablelists/README.md)

* [Create](docs/sdks/displayablelists/README.md#create) - Create displayable lists
* [Delete](docs/sdks/displayablelists/README.md#delete) - Delete displayable lists
* [Get](docs/sdks/displayablelists/README.md#get) - Read displayable lists
* [Update](docs/sdks/displayablelists/README.md#update) - Update displayable lists

#### [Client.Documents](docs/sdks/clientdocuments/README.md)

* [GetPermissions](docs/sdks/clientdocuments/README.md#getpermissions) - Read document permissions
* [Get](docs/sdks/clientdocuments/README.md#get) - Read documents
* [GetByFacets](docs/sdks/clientdocuments/README.md#getbyfacets) - Read documents by facets
* [GetAnalytics](docs/sdks/clientdocuments/README.md#getanalytics) - Read document analytics

#### [Client.Entities](docs/sdks/entities/README.md)

* [List](docs/sdks/entities/README.md#list) - List entities
* [ReadPeople](docs/sdks/entities/README.md#readpeople) - Read people
* [GetTeams](docs/sdks/entities/README.md#getteams) - Read teams

#### [Client.Images](docs/sdks/images/README.md)

* [Get](docs/sdks/images/README.md#get) - Get image
* [Upload](docs/sdks/images/README.md#upload) - Upload images

#### [Client.Insights](docs/sdks/insights/README.md)

* [Get](docs/sdks/insights/README.md#get) - Read insights

#### [Client.Messages](docs/sdks/messages/README.md)

* [Get](docs/sdks/messages/README.md#get) - Read messages

#### [Client.Pins](docs/sdks/pins/README.md)

* [Edit](docs/sdks/pins/README.md#edit) - Update pin
* [Get](docs/sdks/pins/README.md#get) - Read pin
* [List](docs/sdks/pins/README.md#list) - List pins
* [Create](docs/sdks/pins/README.md#create) - Create pin
* [Remove](docs/sdks/pins/README.md#remove) - Delete pin

#### [Client.Search](docs/sdks/search/README.md)

* [Admin](docs/sdks/search/README.md#admin) - Search the index (admin)
* [Autocomplete](docs/sdks/search/README.md#autocomplete) - Autocomplete
* [GetFeed](docs/sdks/search/README.md#getfeed) - Feed of documents and events
* [SuggestPeople](docs/sdks/search/README.md#suggestpeople) - Suggest people
* [SuggestPeopleAdmin](docs/sdks/search/README.md#suggestpeopleadmin) - Suggest people (admin)
* [Recommendations](docs/sdks/search/README.md#recommendations) - Recommend documents
* [Execute](docs/sdks/search/README.md#execute) - Search

#### [Client.Shortcuts](docs/sdks/clientshortcuts/README.md)

* [Create](docs/sdks/clientshortcuts/README.md#create) - Create shortcut
* [Delete](docs/sdks/clientshortcuts/README.md#delete) - Delete shortcut
* [Get](docs/sdks/clientshortcuts/README.md#get) - Read shortcut
* [GetSimilar](docs/sdks/clientshortcuts/README.md#getsimilar) - Get similar shortcuts
* [List](docs/sdks/clientshortcuts/README.md#list) - List shortcuts
* [Preview](docs/sdks/clientshortcuts/README.md#preview) - Preview shortcut
* [Update](docs/sdks/clientshortcuts/README.md#update) - Update shortcut
* [Upload](docs/sdks/clientshortcuts/README.md#upload) - Upload shortcuts

#### [Client.Summarize](docs/sdks/summarize/README.md)

* [Generate](docs/sdks/summarize/README.md#generate) - Summarize documents

#### [Client.Tools](docs/sdks/tools/README.md)

* [ExecuteAction](docs/sdks/tools/README.md#executeaction) - Execute Action Tool

#### [Client.User](docs/sdks/user/README.md)

* [AddCredential](docs/sdks/user/README.md#addcredential) - Create credentials
* [DeleteQueryHistory](docs/sdks/user/README.md#deletequeryhistory) - Delete query history
* [Invite](docs/sdks/user/README.md#invite) - Send invitation
* [GetPublicConfig](docs/sdks/user/README.md#getpublicconfig) - Read public client configuration
* [RemoveCredential](docs/sdks/user/README.md#removecredential) - Delete credentials
* [SendSupportEmail](docs/sdks/user/README.md#sendsupportemail) - Send support email

#### [Client.Verification](docs/sdks/verification/README.md)

* [AddReminder](docs/sdks/verification/README.md#addreminder) - Create verification
* [List](docs/sdks/verification/README.md#list) - List verifications
* [Verify](docs/sdks/verification/README.md#verify) - Update verification


### [Indexing](docs/sdks/indexing/README.md)


#### [Indexing.Authentication](docs/sdks/indexingauthentication/README.md)

* [RotateToken](docs/sdks/indexingauthentication/README.md#rotatetoken) - Rotate token

#### [Indexing.Datasources](docs/sdks/datasources/README.md)

* [Add](docs/sdks/datasources/README.md#add) - Add or update datasource
* [GetConfig](docs/sdks/datasources/README.md#getconfig) - Get datasource config

#### [Indexing.Documents](docs/sdks/indexingdocuments/README.md)

* [AddOrUpdate](docs/sdks/indexingdocuments/README.md#addorupdate) - Index document
* [Index](docs/sdks/indexingdocuments/README.md#index) - Index documents
* [BulkIndex](docs/sdks/indexingdocuments/README.md#bulkindex) - Bulk index documents
* [ProcessAll](docs/sdks/indexingdocuments/README.md#processall) - Schedules the processing of uploaded documents
* [Delete](docs/sdks/indexingdocuments/README.md#delete) - Delete document

#### [Indexing.People](docs/sdks/people/README.md)

* [Index](docs/sdks/people/README.md#index) - Index employee
* [BulkIndexEmployees](docs/sdks/people/README.md#bulkindexemployees) - Bulk index employees
* [~~BulkIndex~~](docs/sdks/people/README.md#bulkindex) - Bulk index employees :warning: **Deprecated**
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

#### [Indexing.Troubleshooting](docs/sdks/troubleshooting/README.md)

* [GetDatasourceStatus](docs/sdks/troubleshooting/README.md#getdatasourcestatus) - Beta: Get datasource status

* [PostDocumentDebug](docs/sdks/troubleshooting/README.md#postdocumentdebug) - Beta: Get document information

* [PostDocumentsDebug](docs/sdks/troubleshooting/README.md#postdocumentsdebug) - Beta: Get information of a batch of documents

* [DebugUser](docs/sdks/troubleshooting/README.md#debuguser) - Beta: Get user information

* [CheckAccess](docs/sdks/troubleshooting/README.md#checkaccess) - Check document access
* [~~GetStatus~~](docs/sdks/troubleshooting/README.md#getstatus) - Get document upload and indexing status :warning: **Deprecated**
* [~~GetDocumentCount~~](docs/sdks/troubleshooting/README.md#getdocumentcount) - Get document count :warning: **Deprecated**
* [~~GetUserCount~~](docs/sdks/troubleshooting/README.md#getusercount) - Get user count :warning: **Deprecated**

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
		apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
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
		apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
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

For example, the `Ask` function may return the following errors:

| Error Type          | Status Code | Content Type     |
| ------------------- | ----------- | ---------------- |
| apierrors.ErrorInfo | 403, 422    | application/json |
| apierrors.APIError  | 4XX, 5XX    | \*/\*            |

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
		apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
	)

	res, err := s.Client.Chat.Ask(ctx, nil, nil, &components.AskRequest{
		SearchRequest: components.SearchRequest{
			TrackingToken: apiclientgo.String("trackingToken"),
			SourceDocument: &components.Document{
				Metadata: &components.DocumentMetadata{
					Datasource: apiclientgo.String("datasource"),
					ObjectType: apiclientgo.String("Feature Request"),
					Container:  apiclientgo.String("container"),
					ParentID:   apiclientgo.String("JIRA_EN-1337"),
					MimeType:   apiclientgo.String("mimeType"),
					DocumentID: apiclientgo.String("documentId"),
					CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
					UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
					Author: &components.Person{
						Name:             "George Clooney",
						ObfuscatedID:     "abc123",
						RelatedDocuments: []components.RelatedDocuments{},
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
								Suggestions: []components.QuerySuggestion{},
							},
							InviteInfo: &components.InviteInfo{
								Invites: []components.ChannelInviteInfo{},
							},
							CustomFields: []components.CustomFieldData{},
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
					Owner: &components.Person{
						Name:         "George Clooney",
						ObfuscatedID: "abc123",
					},
					MentionedPeople: []components.Person{},
					Components: []string{
						"Backend",
						"Networking",
					},
					Status: apiclientgo.String("[\"Done\"]"),
					Pins:   []components.PinDocument{},
					AssignedTo: &components.Person{
						Name:         "George Clooney",
						ObfuscatedID: "abc123",
					},
					UpdatedBy: &components.Person{
						Name:         "George Clooney",
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
								Name:         "George Clooney",
								ObfuscatedID: "abc123",
							},
							Reminders: []components.Reminder{},
							LastReminder: &components.Reminder{
								Assignee: components.Person{
									Name:         "George Clooney",
									ObfuscatedID: "abc123",
								},
								Requestor: &components.Person{
									Name:         "George Clooney",
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
						Name:         "George Clooney",
						ObfuscatedID: "abc123",
					},
				},
			},
			PageSize:       apiclientgo.Int64(100),
			MaxSnippetSize: apiclientgo.Int64(400),
			Query:          "vacation policy",
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
				},
				FacetBucketSize: 132988,
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
			TimeoutMillis: apiclientgo.Int64(5000),
			People: []components.Person{
				components.Person{
					Name:         "George Clooney",
					ObfuscatedID: "abc123",
				},
			},
		},
	})
	if err != nil {

		var e *apierrors.ErrorInfo
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

The default server `https://{domain}-be.glean.com` contains variables and is set to `https://domain-be.glean.com` by default. To override default values, the following options are available when initializing the SDK client instance:

| Variable | Option                      | Default    | Description                                                              |
| -------- | --------------------------- | ---------- | ------------------------------------------------------------------------ |
| `domain` | `WithDomain(domain string)` | `"domain"` | Email domain (without extension) that determines the deployment backend. |

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
		apiclientgo.WithDomain("scared-pearl.biz"),
		apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
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
		apiclientgo.WithServerURL("https://domain-be.glean.com"),
		apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
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
