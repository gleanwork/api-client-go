# GetPersonPhotoResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `TwoHundredImagePngResponseStream`                                 | `io.ReadCloser`                                                    | :heavy_minus_sign:                                                 | Photo bytes returned successfully.                                 |
| `TwoHundredImageJpegResponseStream`                                | `io.ReadCloser`                                                    | :heavy_minus_sign:                                                 | Photo bytes returned successfully.                                 |
| `Headers`                                                          | map[string][]`string`                                              | :heavy_check_mark:                                                 | N/A                                                                |