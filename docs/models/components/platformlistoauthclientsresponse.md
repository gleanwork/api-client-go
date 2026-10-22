# PlatformListOAuthClientsResponse

One page of OAuth clients.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Results`                                                                          | [][components.PlatformOAuthClient](../../models/components/platformoauthclient.md) | :heavy_check_mark:                                                                 | The clients on this page, ordered by display_name and then by client_id.           |
| `HasMore`                                                                          | `bool`                                                                             | :heavy_check_mark:                                                                 | Whether more clients are available after this page.                                |
| `NextCursor`                                                                       | `*string`                                                                          | :heavy_minus_sign:                                                                 | Opaque cursor for the next page; null or absent when has_more is false.            |
| `RequestID`                                                                        | `string`                                                                           | :heavy_check_mark:                                                                 | Request identifier for correlating this response.                                  |