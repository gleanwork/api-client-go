# FileUploadConfig

Configuration settings for the chat file upload feature


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Enabled`                                                     | **bool*                                                       | :heavy_minus_sign:                                            | Whether file upload for Chat is enabled for the deployment    |
| `MaxFileCount`                                                | **int64*                                                      | :heavy_minus_sign:                                            | Maximum number of files that can be uploaded in a single turn |
| `MaxFileSize`                                                 | **int64*                                                      | :heavy_minus_sign:                                            | Maximum file size, in bytes, that is allowed for upload       |
| `UploadTimeoutSeconds`                                        | **int64*                                                      | :heavy_minus_sign:                                            | Timeout in seconds for polling the file upload status         |