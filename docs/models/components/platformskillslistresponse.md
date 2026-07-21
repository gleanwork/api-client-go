# PlatformSkillsListResponse


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Skills`                                                               | [][components.PlatformSkill](../../models/components/platformskill.md) | :heavy_check_mark:                                                     | Skills available to the user.                                          |
| `HasMore`                                                              | `bool`                                                                 | :heavy_check_mark:                                                     | Whether additional results are available.                              |
| `NextCursor`                                                           | `*string`                                                              | :heavy_check_mark:                                                     | Cursor for the next page, or null when no more results are available.  |
| `RequestID`                                                            | `string`                                                               | :heavy_check_mark:                                                     | Platform-generated request ID for support correlation.                 |