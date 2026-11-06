# PlatformDocumentsBatchWarningContext


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Index`                                                                          | `int64`                                                                          | :heavy_check_mark:                                                               | Zero-based index into document_specs.                                            |
| `DocumentID`                                                                     | `*string`                                                                        | :heavy_minus_sign:                                                               | Submitted document_id when the spec was an ID.                                   |
| `URL`                                                                            | `*string`                                                                        | :heavy_minus_sign:                                                               | Submitted URL when the spec was a URL. Empty when the caller sent an empty url.<br/> |