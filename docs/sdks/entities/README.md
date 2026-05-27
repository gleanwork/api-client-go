# Entities

## Overview

### Available Operations

* [GetPersonPhoto](#getpersonphoto) - Get person photo

## GetPersonPhoto

Returns the profile photo bytes for a person whose photo is stored in Glean (crawled from an identity source or user-uploaded via admin console). Photos hosted externally (e.g. Slack CDN) are not served by this endpoint; callers should follow the photoUrl from /people or /listentities directly. Responses include a Cache-Control header (max-age=3600) to reduce redundant fetches.


### Example Usage

<!-- UsageSnippet language="go" operationID="getPersonPhoto" method="get" path="/rest/api/v1/people/{person_id}/photo" -->
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

    res, err := s.Entities.GetPersonPhoto(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredImagePngResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                              | Type                                                                                                                                                                                                   | Required                                                                                                                                                                                               | Description                                                                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                    |
| `personID`                                                                                                                                                                                             | `string`                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                     | The obfuscated ID of the person whose photo to retrieve.                                                                                                                                               |
| `ds`                                                                                                                                                                                                   | `*string`                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                     | Optional datasource override for crawled photos (e.g. AZURE, GDRIVE, OKTA). When omitted, the datasource is derived from the person's stored photo URL or the deployment's primary person datasource.<br/> |
| `opts`                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                     | The options for this request.                                                                                                                                                                          |

### Response

**[*operations.GetPersonPhotoResponse](../../models/operations/getpersonphotoresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |