# PlatformListAdminAgentsResponse

One page of agents.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Results`                                                                        | [][components.PlatformAdminAgent](../../models/components/platformadminagent.md) | :heavy_check_mark:                                                               | The agents on this page, ordered by display_name and then by agent_id.           |
| `HasMore`                                                                        | `bool`                                                                           | :heavy_check_mark:                                                               | Whether more agents are available after this page.                               |
| `NextCursor`                                                                     | `*string`                                                                        | :heavy_minus_sign:                                                               | Opaque cursor for the next page; null or absent when has_more is false.          |
| `RequestID`                                                                      | `string`                                                                         | :heavy_check_mark:                                                               | Request identifier for correlating this response.                                |