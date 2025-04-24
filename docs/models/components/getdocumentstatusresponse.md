# GetDocumentStatusResponse

Describes the response body of the /getdocumentstatus API call


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `UploadStatus`                                                | **string*                                                     | :heavy_minus_sign:                                            | Upload status, enum of NOT_UPLOADED, UPLOADED, STATUS_UNKNOWN |
| `LastUploadedAt`                                              | **int64*                                                      | :heavy_minus_sign:                                            | Time of last successful upload, in epoch seconds              |
| `IndexingStatus`                                              | **string*                                                     | :heavy_minus_sign:                                            | Indexing status, enum of NOT_INDEXED, INDEXED, STATUS_UNKNOWN |
| `LastIndexedAt`                                               | **int64*                                                      | :heavy_minus_sign:                                            | Time of last successful indexing, in epoch seconds            |