# PlatformProblemDetailCode

Stable machine-readable error code.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PlatformProblemDetailCodeInvalidRequest

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformProblemDetailCode("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `PlatformProblemDetailCodeInvalidRequest`          | invalid_request                                    |
| `PlatformProblemDetailCodeMissingRequiredField`    | missing_required_field                             |
| `PlatformProblemDetailCodeInvalidParameter`        | invalid_parameter                                  |
| `PlatformProblemDetailCodeInvalidCursor`           | invalid_cursor                                     |
| `PlatformProblemDetailCodeExpiredCursor`           | expired_cursor                                     |
| `PlatformProblemDetailCodeInvalidFilter`           | invalid_filter                                     |
| `PlatformProblemDetailCodeInvalidDatasource`       | invalid_datasource                                 |
| `PlatformProblemDetailCodeAuthenticationRequired`  | authentication_required                            |
| `PlatformProblemDetailCodeTokenExpired`            | token_expired                                      |
| `PlatformProblemDetailCodeInsufficientPermissions` | insufficient_permissions                           |
| `PlatformProblemDetailCodeResourceNotFound`        | resource_not_found                                 |
| `PlatformProblemDetailCodeMethodNotAllowed`        | method_not_allowed                                 |
| `PlatformProblemDetailCodeRequestTimeout`          | request_timeout                                    |
| `PlatformProblemDetailCodeConflict`                | conflict                                           |
| `PlatformProblemDetailCodeGone`                    | gone                                               |
| `PlatformProblemDetailCodeUnprocessableQuery`      | unprocessable_query                                |
| `PlatformProblemDetailCodeRateLimitExceeded`       | rate_limit_exceeded                                |
| `PlatformProblemDetailCodeInternalError`           | internal_error                                     |
| `PlatformProblemDetailCodeServiceUnavailable`      | service_unavailable                                |