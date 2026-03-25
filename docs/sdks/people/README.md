# Indexing.People

## Overview

### Available Operations

* [Debug](#debug) - Beta: Get user information

* [~~Count~~](#count) - Get user count :warning: **Deprecated**
* [Index](#index) - Index employee
* [~~BulkIndex~~](#bulkindex) - Bulk index employees :warning: **Deprecated**
* [ProcessAllEmployeesAndTeams](#processallemployeesandteams) - Schedules the processing of uploaded employees and teams
* [Delete](#delete) - Delete employee
* [IndexTeam](#indexteam) - Index team
* [DeleteTeam](#deleteteam) - Delete team
* [BulkIndexTeams](#bulkindexteams) - Bulk index teams

## Debug

Gives various information that would help in debugging related to a particular user. Currently in beta, might undergo breaking changes without prior notice.

Tip: Refer to the [Troubleshooting tutorial](https://developers.glean.com/indexing/debugging/datasource-config) for more information.


### Example Usage

<!-- UsageSnippet language="go" operationID="post_/api/index/v1/debug/{datasource}/user" method="post" path="/api/index/v1/debug/{datasource}/user" -->
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

    res, err := s.Indexing.People.Debug(ctx, "<value>", components.DebugUserRequest{
        Email: "u1@foo.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DebugUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `datasource`                                                               | `string`                                                                   | :heavy_check_mark:                                                         | The datasource to which the user belongs                                   |
| `debugUserRequest`                                                         | [components.DebugUserRequest](../../models/components/debuguserrequest.md) | :heavy_check_mark:                                                         | N/A                                                                        |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.PostAPIIndexV1DebugDatasourceUserResponse](../../models/operations/postapiindexv1debugdatasourceuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ~~Count~~

Fetches user count for the specified custom datasource.

Tip: Use [/debug/{datasource}/status](https://developers.glean.com/indexing/debugging/datasource-status) for richer information.


> :warning: **DEPRECATED**: Deprecated on 2026-02-03, removal scheduled for 2026-10-15: Endpoint is deprecated.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/api/index/v1/getusercount" method="post" path="/api/index/v1/getusercount" -->
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

    res, err := s.Indexing.People.Count(ctx, components.GetUserCountRequest{
        Datasource: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetUserCountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.GetUserCountRequest](../../models/components/getusercountrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.PostAPIIndexV1GetusercountResponse](../../models/operations/postapiindexv1getusercountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Index

Adds an employee or replaces the existing information about an employee.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/api/index/v1/indexemployee" method="post" path="/api/index/v1/indexemployee" -->
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

    res, err := s.Indexing.People.Index(ctx, components.IndexEmployeeRequest{
        Employee: components.EmployeeInfoDefinition{
            Email: "Jerrold_Hermann@hotmail.com",
            Department: "<value>",
            DatasourceProfiles: []components.DatasourceProfile{
                components.DatasourceProfile{
                    Datasource: "github",
                    Handle: "<value>",
                },
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

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.IndexEmployeeRequest](../../models/components/indexemployeerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.PostAPIIndexV1IndexemployeeResponse](../../models/operations/postapiindexv1indexemployeeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ~~BulkIndex~~

Replaces all the currently indexed employees using paginated batch API calls. Please refer to the [bulk indexing](https://developers.glean.com/indexing/documents/bulk-upload-model) documentation for an explanation of how to use bulk endpoints.

> :warning: **DEPRECATED**: Deprecated on 2026-02-03, removal scheduled for 2026-10-15: Endpoint is deprecated.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/api/index/v1/bulkindexemployees" method="post" path="/api/index/v1/bulkindexemployees" -->
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

    res, err := s.Indexing.People.BulkIndex(ctx, components.BulkIndexEmployeesRequest{
        UploadID: "<id>",
        Employees: []components.EmployeeInfoDefinition{
            components.EmployeeInfoDefinition{
                Email: "Robin.Stoltenberg@yahoo.com",
                Department: "<value>",
                DatasourceProfiles: []components.DatasourceProfile{
                    components.DatasourceProfile{
                        Datasource: "github",
                        Handle: "<value>",
                    },
                },
            },
            components.EmployeeInfoDefinition{
                Email: "Robin.Stoltenberg@yahoo.com",
                Department: "<value>",
                DatasourceProfiles: []components.DatasourceProfile{
                    components.DatasourceProfile{
                        Datasource: "github",
                        Handle: "<value>",
                    },
                },
            },
            components.EmployeeInfoDefinition{
                Email: "Robin.Stoltenberg@yahoo.com",
                Department: "<value>",
                DatasourceProfiles: []components.DatasourceProfile{
                    components.DatasourceProfile{
                        Datasource: "github",
                        Handle: "<value>",
                    },
                },
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

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.BulkIndexEmployeesRequest](../../models/components/bulkindexemployeesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PostAPIIndexV1BulkindexemployeesResponse](../../models/operations/postapiindexv1bulkindexemployeesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ProcessAllEmployeesAndTeams

Schedules the immediate processing of employees and teams uploaded through the indexing API. By default all uploaded people data will be processed asynchronously but this API can be used to schedule its processing on demand.


### Example Usage

<!-- UsageSnippet language="go" operationID="post_/api/index/v1/processallemployeesandteams" method="post" path="/api/index/v1/processallemployeesandteams" -->
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
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Indexing.People.ProcessAllEmployeesAndTeams(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PostAPIIndexV1ProcessallemployeesandteamsResponse](../../models/operations/postapiindexv1processallemployeesandteamsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete an employee. Silently succeeds if employee is not present.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/api/index/v1/deleteemployee" method="post" path="/api/index/v1/deleteemployee" -->
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

    res, err := s.Indexing.People.Delete(ctx, components.DeleteEmployeeRequest{
        EmployeeEmail: "<value>",
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
| `request`                                                                            | [components.DeleteEmployeeRequest](../../models/components/deleteemployeerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PostAPIIndexV1DeleteemployeeResponse](../../models/operations/postapiindexv1deleteemployeeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## IndexTeam

Adds a team or updates information about a team

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/api/index/v1/indexteam" method="post" path="/api/index/v1/indexteam" -->
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

    res, err := s.Indexing.People.IndexTeam(ctx, components.IndexTeamRequest{
        Team: components.TeamInfoDefinition{
            ID: "<id>",
            Name: "<value>",
            DatasourceProfiles: []components.DatasourceProfile{
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
            Members: []components.TeamMember{
                components.TeamMember{
                    Email: "Nasir.Hilll73@hotmail.com",
                },
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

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.IndexTeamRequest](../../models/components/indexteamrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.PostAPIIndexV1IndexteamResponse](../../models/operations/postapiindexv1indexteamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteTeam

Delete a team based on provided id.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/api/index/v1/deleteteam" method="post" path="/api/index/v1/deleteteam" -->
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

    res, err := s.Indexing.People.DeleteTeam(ctx, components.DeleteTeamRequest{
        ID: "<id>",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.DeleteTeamRequest](../../models/components/deleteteamrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.PostAPIIndexV1DeleteteamResponse](../../models/operations/postapiindexv1deleteteamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## BulkIndexTeams

Replaces all the currently indexed teams using paginated batch API calls. Please refer to the [bulk indexing](https://developers.glean.com/indexing/documents/bulk-upload-model) documentation for an explanation of how to use bulk endpoints.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/api/index/v1/bulkindexteams" method="post" path="/api/index/v1/bulkindexteams" -->
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

    res, err := s.Indexing.People.BulkIndexTeams(ctx, components.BulkIndexTeamsRequest{
        UploadID: "<id>",
        Teams: []components.TeamInfoDefinition{
            components.TeamInfoDefinition{
                ID: "<id>",
                Name: "<value>",
                DatasourceProfiles: []components.DatasourceProfile{
                    components.DatasourceProfile{
                        Datasource: "github",
                        Handle: "<value>",
                    },
                    components.DatasourceProfile{
                        Datasource: "github",
                        Handle: "<value>",
                    },
                },
                Members: []components.TeamMember{},
            },
            components.TeamInfoDefinition{
                ID: "<id>",
                Name: "<value>",
                DatasourceProfiles: []components.DatasourceProfile{
                    components.DatasourceProfile{
                        Datasource: "github",
                        Handle: "<value>",
                    },
                    components.DatasourceProfile{
                        Datasource: "github",
                        Handle: "<value>",
                    },
                },
                Members: []components.TeamMember{},
            },
            components.TeamInfoDefinition{
                ID: "<id>",
                Name: "<value>",
                DatasourceProfiles: []components.DatasourceProfile{
                    components.DatasourceProfile{
                        Datasource: "github",
                        Handle: "<value>",
                    },
                    components.DatasourceProfile{
                        Datasource: "github",
                        Handle: "<value>",
                    },
                },
                Members: []components.TeamMember{},
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

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.BulkIndexTeamsRequest](../../models/components/bulkindexteamsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PostAPIIndexV1BulkindexteamsResponse](../../models/operations/postapiindexv1bulkindexteamsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |