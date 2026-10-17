# Departments

## Overview

### Available Operations

* [List](#list) - List departments
* [Get](#get) - Get a department

## List

List the departments in the Glean directory, ordered by display_name and then by department_id. A department is active when at least one active user belongs to it. A department is inactive when only inactive users, such as former employees, belong to it, or when only a usage limit names it. Inactive departments are left out unless an is_active filter includes the value "false". Department names are exact, so names that differ only by case are different departments.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-departments-list" method="get" path="/api/departments" -->
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

    res, err := s.Departments.List(ctx, apiclientgo.Pointer[int64](50), nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformListDepartmentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                     | Type                                                                                                                                                                                                                                                                                                                                                                                                          | Required                                                                                                                                                                                                                                                                                                                                                                                                      | Description                                                                                                                                                                                                                                                                                                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                            | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                           |
| `pageSize`                                                                                                                                                                                                                                                                                                                                                                                                    | `*int64`                                                                                                                                                                                                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                            | Maximum number of departments to return. Defaults to 50. Maximum is 100.                                                                                                                                                                                                                                                                                                                                      |
| `cursor`                                                                                                                                                                                                                                                                                                                                                                                                      | `*string`                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                            | Opaque pagination cursor from a previous response. Send the same filters that the previous request used.<br/>                                                                                                                                                                                                                                                                                                 |
| `filters`                                                                                                                                                                                                                                                                                                                                                                                                     | [][components.PlatformDepartmentFilter](../../models/components/platformdepartmentfilter.md)                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                            | JSON-encoded filters that choose which departments are returned. The only supported field is is_active, with the values "true" and "false" and the operator EQUALS. Multiple values OR within a filter. Multiple filters AND together. Without an is_active filter, only active departments are returned. To return active and inactive departments, send [{"field":"is_active","values":["true","false"]}].<br/> |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                            | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                 |

### Response

**[*operations.PlatformDepartmentsListResponse](../../models/operations/platformdepartmentslistresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## Get

Get one department by its department_id. Inactive departments are returned, so a caller can show the name of a department that a usage limit still names. A department is active when at least one active user belongs to it. A department is inactive when only inactive users, such as former employees, belong to it, or when only a usage limit names it. A department_id that is not valid, or that no user and no usage limit uses, returns 404 resource_not_found. A department_id comes from the exact department name, so a renamed department gets a new ID, and a name that is used again gets its old ID back.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-departments-get" method="get" path="/api/departments/{department_id}" -->
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

    res, err := s.Departments.Get(ctx, "{department_id}")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformGetDepartmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  | Example                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |                                                                                                                              |
| `departmentID`                                                                                                               | `string`                                                                                                                     | :heavy_check_mark:                                                                                                           | The department_id from a list departments response or from a user's department. IDs are opaque. Do not parse or build them.<br/> | {department_id}                                                                                                              |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |                                                                                                                              |

### Response

**[*operations.PlatformDepartmentsGetResponse](../../models/operations/platformdepartmentsgetresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |