# PlatformSkillSourcePreviewRequest


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `SourceURL`                                                                | `string`                                                                   | :heavy_check_mark:                                                         | GitHub URL for a skill directory, SKILL.md file, or repository to inspect. |
| `Stream`                                                                   | `*bool`                                                                    | :heavy_minus_sign:                                                         | Whether to stream repository scan progress using server-sent events.       |