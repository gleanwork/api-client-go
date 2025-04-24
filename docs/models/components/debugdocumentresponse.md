# DebugDocumentResponse

Describes the response body of the /debug/{datasource}/document API call


## Fields

| Field                                                                                                 | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `Status`                                                                                              | [*components.DocumentStatusResponse](../../models/components/documentstatusresponse.md)               | :heavy_minus_sign:                                                                                    | Describes the document status response body                                                           |
| `UploadedPermissions`                                                                                 | [*components.DocumentPermissionsDefinition](../../models/components/documentpermissionsdefinition.md) | :heavy_minus_sign:                                                                                    | describes the access control details of the document                                                  |