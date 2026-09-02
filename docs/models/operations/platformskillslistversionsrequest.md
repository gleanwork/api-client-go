# PlatformSkillsListVersionsRequest


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `SkillID`                                          | `string`                                           | :heavy_check_mark:                                 | Glean skill ID.                                    | {skill_id}                                         |
| `PageSize`                                         | `*int64`                                           | :heavy_minus_sign:                                 | Maximum number of versions to return.              |                                                    |
| `Cursor`                                           | `*string`                                          | :heavy_minus_sign:                                 | Opaque pagination cursor from a previous response. |                                                    |