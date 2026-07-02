# Client.Datasources

## Overview

### Available Operations

* [RetrieveConfiguration](#retrieveconfiguration) - Get datasource instance configuration
* [UpdateConfiguration](#updateconfiguration) - Update datasource instance configuration
* [RetrieveCredentialStatus](#retrievecredentialstatus) - Get datasource instance credential status
* [RotateCredentials](#rotatecredentials) - Rotate datasource instance credentials

## RetrieveConfiguration

Gets the greenlisted configuration values for a datasource instance. Returns only configuration keys that are exposed via the public API greenlist.


### Example Usage

<!-- UsageSnippet language="go" operationID="getDatasourceInstanceConfiguration" method="get" path="/rest/api/v1/configure/datasources/{datasourceId}/instances/{instanceId}" -->
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

    res, err := s.Client.Datasources.RetrieveConfiguration(ctx, "o365sharepoint", "o365sharepoint_abc123")
    if err != nil {
        log.Fatal(err)
    }
    if res.DatasourceConfigurationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `datasourceID`                                           | `string`                                                 | :heavy_check_mark:                                       | The datasource type identifier (e.g. o365sharepoint)     | o365sharepoint                                           |
| `instanceID`                                             | `string`                                                 | :heavy_check_mark:                                       | The datasource instance identifier                       | o365sharepoint_abc123                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetDatasourceInstanceConfigurationResponse](../../models/operations/getdatasourceinstanceconfigurationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 403, 404           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UpdateConfiguration

Updates the greenlisted configuration values for a datasource instance. Only configuration keys that are exposed via the public API greenlist may be set. Returns the full greenlisted configuration after the update is applied.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateDatasourceInstanceConfiguration" method="patch" path="/rest/api/v1/configure/datasources/{datasourceId}/instances/{instanceId}" -->
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

    res, err := s.Client.Datasources.UpdateConfiguration(ctx, "o365sharepoint", "o365sharepoint_abc123", components.UpdateDatasourceConfigurationRequest{
        Configuration: components.DatasourceInstanceConfiguration{
            Values: map[string]components.ConfigurationValue{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatasourceConfigurationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        | Example                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |                                                                                                                    |
| `datasourceID`                                                                                                     | `string`                                                                                                           | :heavy_check_mark:                                                                                                 | The datasource type identifier (e.g. o365sharepoint)                                                               | o365sharepoint                                                                                                     |
| `instanceID`                                                                                                       | `string`                                                                                                           | :heavy_check_mark:                                                                                                 | The datasource instance identifier                                                                                 | o365sharepoint_abc123                                                                                              |
| `updateDatasourceConfigurationRequest`                                                                             | [components.UpdateDatasourceConfigurationRequest](../../models/components/updatedatasourceconfigurationrequest.md) | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |                                                                                                                    |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |                                                                                                                    |

### Response

**[*operations.UpdateDatasourceInstanceConfigurationResponse](../../models/operations/updatedatasourceinstanceconfigurationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 403, 404           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## RetrieveCredentialStatus

Returns the current credential status for a datasource instance. Access is limited to callers with the ADMIN scope; the handler enforces this check.


### Example Usage

<!-- UsageSnippet language="go" operationID="getDatasourceCredentialStatus" method="get" path="/rest/api/v1/datasource/{datasourceInstanceId}/credentialstatus" -->
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

    res, err := s.Client.Datasources.RetrieveCredentialStatus(ctx, "o365sharepoint_abc123")
    if err != nil {
        log.Fatal(err)
    }
    if res.DatasourceCredentialStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |                                                                      |
| `datasourceInstanceID`                                               | `string`                                                             | :heavy_check_mark:                                                   | The full datasource instance identifier (e.g. o365sharepoint_abc123) | o365sharepoint_abc123                                                |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |                                                                      |

### Response

**[*operations.GetDatasourceCredentialStatusResponse](../../models/operations/getdatasourcecredentialstatusresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 403, 404           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## RotateCredentials

Rotates the credentials that a datasource instance uses to connect to its upstream system. Replaces the active credential material with the supplied values and returns the credential status after rotation. Access is limited to callers with the ADMIN scope; the handler enforces this check.
Only keys recognized as credential material for the datasource type may be set in `credentials.values` (e.g. `clientSecret`, `apiToken`, `privateKey`, depending on the configured auth method). Unrecognized keys, or keys that correspond to non-credential configuration, cause a 400; other instance configuration must be updated via PATCH /configure/datasources/{datasourceId}/instances/{instanceId}.


### Example Usage

<!-- UsageSnippet language="go" operationID="rotateDatasourceCredentials" method="post" path="/rest/api/v1/datasource/{datasourceInstanceId}/credentials" -->
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

    res, err := s.Client.Datasources.RotateCredentials(ctx, "o365sharepoint_abc123", components.RotateDatasourceCredentialsRequest{
        Credentials: components.DatasourceInstanceConfiguration{
            Values: map[string]components.ConfigurationValue{
                "key": components.ConfigurationValue{},
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatasourceCredentialStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `datasourceInstanceID`                                                                                         | `string`                                                                                                       | :heavy_check_mark:                                                                                             | The full datasource instance identifier (e.g. o365sharepoint_abc123)                                           | o365sharepoint_abc123                                                                                          |
| `rotateDatasourceCredentialsRequest`                                                                           | [components.RotateDatasourceCredentialsRequest](../../models/components/rotatedatasourcecredentialsrequest.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |                                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.RotateDatasourceCredentialsResponse](../../models/operations/rotatedatasourcecredentialsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 403, 404           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |