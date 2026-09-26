# PlatformSkillSyncResponse


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `CommitSha`                                            | `string`                                               | :heavy_check_mark:                                     | Git commit SHA now associated with the skill.          |
| `IsUpdated`                                            | `bool`                                                 | :heavy_check_mark:                                     | Whether this request created a new skill version.      |
| `RequestID`                                            | `string`                                               | :heavy_check_mark:                                     | Platform-generated request ID for support correlation. |