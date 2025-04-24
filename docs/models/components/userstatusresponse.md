# UserStatusResponse

Describes the user status response body


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 | Example                                                                     |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `IsActiveUser`                                                              | **bool*                                                                     | :heavy_minus_sign:                                                          | Whether the user is active or not                                           | true                                                                        |
| `UploadStatus`                                                              | [*components.UploadStatusEnum](../../models/components/uploadstatusenum.md) | :heavy_minus_sign:                                                          | Upload status, enum of NOT_UPLOADED, UPLOADED, STATUS_UNKNOWN               | UPLOADED                                                                    |
| `LastUploadedAt`                                                            | **string*                                                                   | :heavy_minus_sign:                                                          | Time of last successful upload for the user, in ISO 8601 format             | 2021-08-06T17:58:01.000Z                                                    |