# ChatFileMetadata

Metadata of a file uploaded by a user for Chat.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `Status`                                                                              | [*components.ChatFileStatus](../../models/components/chatfilestatus.md)               | :heavy_minus_sign:                                                                    | Current status of the file.                                                           |
| `UploadTime`                                                                          | **int64*                                                                              | :heavy_minus_sign:                                                                    | Upload time, in epoch seconds.                                                        |
| `ProcessedSize`                                                                       | **int64*                                                                              | :heavy_minus_sign:                                                                    | Size of the processed file in bytes.                                                  |
| `FailureReason`                                                                       | [*components.ChatFileFailureReason](../../models/components/chatfilefailurereason.md) | :heavy_minus_sign:                                                                    | Reason for failed status.                                                             |
| `MimeType`                                                                            | **string*                                                                             | :heavy_minus_sign:                                                                    | MIME type of the file.                                                                |