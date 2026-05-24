# ErrorInfoResponse

Error response for custom metadata operations


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Error`                                                            | `string`                                                           | :heavy_check_mark:                                                 | Error message describing what went wrong                           |
| `Message`                                                          | `*string`                                                          | :heavy_minus_sign:                                                 | Additional details about the error                                 |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |