# Skills

## Overview

### Available Operations

* [Create](#create) - Create skill
* [List](#list) - List skills
* [Import](#import) - Import skills from GitHub
* [Validate](#validate) - Validate skill bundle
* [PreviewSource](#previewsource) - Preview a GitHub skill source
* [Update](#update) - Update skill
* [Delete](#delete) - Delete skill
* [Retrieve](#retrieve) - Retrieve skill
* [RetrieveContent](#retrievecontent) - Download skill content
* [Sync](#sync) - Sync a GitHub-imported skill
* [CreateVersion](#createversion) - Create skill version
* [ListVersions](#listversions) - List skill versions
* [RetrieveVersion](#retrieveversion) - Retrieve skill version
* [RetrieveVersionContent](#retrieveversioncontent) - Download skill version content

## Create

Create a skill from an uploaded SKILL.md, .zip, or .skill bundle. If the authenticated user already has a skill with the same name, the existing skill is superseded with a new version.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-skills-create" method="post" path="/api/skills" -->
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

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.Skills.Create(ctx, components.PlatformSkillCreateRequest{
        File: components.PlatformSkillCreateRequestFile{
            FileName: "example.file",
            Content: example,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformSkillCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.PlatformSkillCreateRequest](../../models/components/platformskillcreaterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PlatformSkillsCreateResponse](../../models/operations/platformskillscreateresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 413, 429    | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## List

List skills available to the authenticated user.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-skills-list" method="get" path="/api/skills" -->
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

    res, err := s.Skills.List(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformSkillsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `pageSize`                                               | `*int64`                                                 | :heavy_minus_sign:                                       | Maximum number of skills to return.                      |
| `cursor`                                                 | `*string`                                                | :heavy_minus_sign:                                       | Opaque pagination cursor from a previous response.       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PlatformSkillsListResponse](../../models/operations/platformskillslistresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## Import

Import one or more skills selected from a GitHub source preview. Each source URL is fetched and persisted as an independent skill with source provenance. This operation does not create a durable source resource. The import is atomic: if any source cannot be fetched, validated, or persisted, no skills are created.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-skills-import" method="post" path="/api/skills/import" -->
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

    res, err := s.Skills.Import(ctx, components.PlatformSkillImportRequest{
        SourceUrls: []string{
            "https://github.com/anthropics/skills/tree/main/skills/skill-creator",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformSkillImportResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.PlatformSkillImportRequest](../../models/components/platformskillimportrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PlatformSkillsImportResponse](../../models/operations/platformskillsimportresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 408, 409, 413, 429    | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## Validate

Validate a skill bundle without persisting it. Accepts a SKILL.md, .zip, or .skill upload and returns parsed metadata plus the normalized file layout.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-skills-validate" method="post" path="/api/skills/validation" -->
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

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.Skills.Validate(ctx, components.PlatformSkillValidationRequest{
        File: components.PlatformSkillValidationRequestFile{
            FileName: "example.file",
            Content: example,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformSkillValidationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [components.PlatformSkillValidationRequest](../../models/components/platformskillvalidationrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.PlatformSkillsValidateResponse](../../models/operations/platformskillsvalidateresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 413, 429    | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## PreviewSource

Inspect a GitHub URL without persisting a source or any discovered skills. Set stream to true to receive repository scan progress as server-sent events; otherwise the response contains the completed preview.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-skills-preview-source" method="post" path="/api/skills/sources/preview" -->
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

    res, err := s.Skills.PreviewSource(ctx, components.PlatformSkillSourcePreviewRequest{
        SourceURL: "https://github.com/anthropics/skills",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformSkillSourcePreviewResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [components.PlatformSkillSourcePreviewRequest](../../models/components/platformskillsourcepreviewrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.PlatformSkillsPreviewSourceResponse](../../models/operations/platformskillspreviewsourceresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 408, 413, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## Update

Update mutable metadata for a skill. V1 supports enabling or disabling a skill without changing its content.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-skills-update" method="patch" path="/api/skills/{skill_id}" -->
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

    res, err := s.Skills.Update(ctx, "{skill_id}", components.PlatformSkillUpdateRequest{
        Status: components.PlatformSkillUpdateStatusEnabled,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformSkillUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    | Example                                                                                        |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |                                                                                                |
| `skillID`                                                                                      | `string`                                                                                       | :heavy_check_mark:                                                                             | Glean skill ID.                                                                                | {skill_id}                                                                                     |
| `platformSkillUpdateRequest`                                                                   | [components.PlatformSkillUpdateRequest](../../models/components/platformskillupdaterequest.md) | :heavy_check_mark:                                                                             | N/A                                                                                            | {<br/>"status": "ENABLED"<br/>}                                                                |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |                                                                                                |

### Response

**[*operations.PlatformSkillsUpdateResponse](../../models/operations/platformskillsupdateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.PlatformProblemDetailError   | 400, 401, 403, 404, 408, 409, 413, 429 | application/problem+json               |
| apierrors.PlatformProblemDetailError   | 500, 503                               | application/problem+json               |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Delete

Delete a skill the authenticated caller is allowed to manage. This operation permanently removes all versions of the skill.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-skills-delete" method="delete" path="/api/skills/{skill_id}" -->
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

    res, err := s.Skills.Delete(ctx, "{skill_id}")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `skillID`                                                | `string`                                                 | :heavy_check_mark:                                       | Glean skill ID.                                          | {skill_id}                                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PlatformSkillsDeleteResponse](../../models/operations/platformskillsdeleteresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## Retrieve

Retrieve metadata for a skill available to the authenticated user.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-skills-get" method="get" path="/api/skills/{skill_id}" -->
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

    res, err := s.Skills.Retrieve(ctx, "{skill_id}")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformSkillGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `skillID`                                                | `string`                                                 | :heavy_check_mark:                                       | Glean skill ID.                                          | {skill_id}                                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PlatformSkillsGetResponse](../../models/operations/platformskillsgetresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## RetrieveContent

Download the latest installable bundle for a skill available to the authenticated user.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-skills-get-content" method="get" path="/api/skills/{skill_id}/content" -->
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

    res, err := s.Skills.RetrieveContent(ctx, "{skill_id}")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `skillID`                                                | `string`                                                 | :heavy_check_mark:                                       | Glean skill ID.                                          | {skill_id}                                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PlatformSkillsGetContentResponse](../../models/operations/platformskillsgetcontentresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## Sync

Refresh one GitHub-imported skill from its stored source URL. If the skill content has changed, this operation creates a new skill version. If the skill is no longer present upstream, the stored skill is left unchanged and must be deleted explicitly.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-skills-sync" method="post" path="/api/skills/{skill_id}/sync" -->
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

    res, err := s.Skills.Sync(ctx, "{skill_id}")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformSkillSyncResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `skillID`                                                | `string`                                                 | :heavy_check_mark:                                       | ID of the GitHub-imported skill to sync.                 | {skill_id}                                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PlatformSkillsSyncResponse](../../models/operations/platformskillssyncresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.PlatformProblemDetailError   | 400, 401, 403, 404, 408, 409, 413, 429 | application/problem+json               |
| apierrors.PlatformProblemDetailError   | 500, 503                               | application/problem+json               |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateVersion

Create a new immutable version for an existing caller-managed skill from an uploaded SKILL.md, .zip, or .skill bundle.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-skills-create-version" method="post" path="/api/skills/{skill_id}/versions" -->
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

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.Skills.CreateVersion(ctx, "{skill_id}", components.PlatformSkillVersionCreateRequest{
        File: components.PlatformSkillVersionCreateRequestFile{
            FileName: "example.file",
            Content: example,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformSkillVersionCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  | Example                                                                                                      |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |                                                                                                              |
| `skillID`                                                                                                    | `string`                                                                                                     | :heavy_check_mark:                                                                                           | Glean skill ID.                                                                                              | {skill_id}                                                                                                   |
| `platformSkillVersionCreateRequest`                                                                          | [components.PlatformSkillVersionCreateRequest](../../models/components/platformskillversioncreaterequest.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |                                                                                                              |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |                                                                                                              |

### Response

**[*operations.PlatformSkillsCreateVersionResponse](../../models/operations/platformskillscreateversionresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.PlatformProblemDetailError   | 400, 401, 403, 404, 408, 409, 413, 429 | application/problem+json               |
| apierrors.PlatformProblemDetailError   | 500, 503                               | application/problem+json               |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ListVersions

List versions for a skill available to the authenticated user.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-skills-list-versions" method="get" path="/api/skills/{skill_id}/versions" -->
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

    res, err := s.Skills.ListVersions(ctx, "{skill_id}", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformSkillVersionsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `skillID`                                                | `string`                                                 | :heavy_check_mark:                                       | Glean skill ID.                                          | {skill_id}                                               |
| `pageSize`                                               | `*int64`                                                 | :heavy_minus_sign:                                       | Maximum number of versions to return.                    |                                                          |
| `cursor`                                                 | `*string`                                                | :heavy_minus_sign:                                       | Opaque pagination cursor from a previous response.       |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PlatformSkillsListVersionsResponse](../../models/operations/platformskillslistversionsresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## RetrieveVersion

Retrieve metadata for a skill version available to the authenticated user.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-skills-get-version" method="get" path="/api/skills/{skill_id}/versions/{version}" -->
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

    res, err := s.Skills.RetrieveVersion(ctx, "{skill_id}", 1)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlatformSkillVersionGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `skillID`                                                | `string`                                                 | :heavy_check_mark:                                       | Glean skill ID.                                          | {skill_id}                                               |
| `version`                                                | `int64`                                                  | :heavy_check_mark:                                       | Major version number.                                    | 1                                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PlatformSkillsGetVersionResponse](../../models/operations/platformskillsgetversionresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## RetrieveVersionContent

Download the installable bundle for a skill version available to the authenticated user.


### Example Usage

<!-- UsageSnippet language="go" operationID="platform-skills-get-version-content" method="get" path="/api/skills/{skill_id}/versions/{version}/content" -->
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

    res, err := s.Skills.RetrieveVersionContent(ctx, "{skill_id}", 1)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `skillID`                                                | `string`                                                 | :heavy_check_mark:                                       | Glean skill ID.                                          | {skill_id}                                               |
| `version`                                                | `int64`                                                  | :heavy_check_mark:                                       | Major version number.                                    | 1                                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PlatformSkillsGetVersionContentResponse](../../models/operations/platformskillsgetversioncontentresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.PlatformProblemDetailError | 400, 401, 403, 404, 408, 429         | application/problem+json             |
| apierrors.PlatformProblemDetailError | 500, 503                             | application/problem+json             |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |