# PlatformListUsersResponse

One page of users.


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Results`                                                               | [][components.PlatformUser](../../models/components/platformuser.md)    | :heavy_check_mark:                                                      | The users on this page, ordered by display_name and then by user_id.    |
| `HasMore`                                                               | `bool`                                                                  | :heavy_check_mark:                                                      | Whether more users are available after this page.                       |
| `NextCursor`                                                            | `*string`                                                               | :heavy_minus_sign:                                                      | Opaque cursor for the next page; null or absent when has_more is false. |
| `RequestID`                                                             | `string`                                                                | :heavy_check_mark:                                                      | Request identifier for correlating this response.                       |