# DebugUserResponse

Describes the response body of the /debug/{datasource}/user API call


## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `Status`                                                                                       | [*components.UserStatusResponse](../../models/components/userstatusresponse.md)                | :heavy_minus_sign:                                                                             | Describes the user status response body                                                        |
| `UploadedGroups`                                                                               | [][components.DatasourceGroupDefinition](../../models/components/datasourcegroupdefinition.md) | :heavy_minus_sign:                                                                             | List of groups the user is a member of, as uploaded via permissions API.                       |