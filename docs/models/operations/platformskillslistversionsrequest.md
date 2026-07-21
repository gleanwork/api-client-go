# PlatformSkillsListVersionsRequest


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `SkillID`                                          | `string`                                           | :heavy_check_mark:                                 | Glean skill ID.                                    |
| `PageSize`                                         | `*int64`                                           | :heavy_minus_sign:                                 | Maximum number of versions to return.              |
| `Cursor`                                           | `*string`                                          | :heavy_minus_sign:                                 | Opaque pagination cursor from a previous response. |