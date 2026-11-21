# PlatformAgentSkill

A skill defined inline in the agent. Reference a published skill through `dependencies` instead.



## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `DisplayName`                                                                                            | `string`                                                                                                 | :heavy_check_mark:                                                                                       | Human-readable skill name.                                                                               |
| `MainContent`                                                                                            | `string`                                                                                                 | :heavy_check_mark:                                                                                       | Main skill content, inlined from `SKILL.md`.                                                             |
| `Files`                                                                                                  | [][components.PlatformSkillSourcePreviewFile](../../models/components/platformskillsourcepreviewfile.md) | :heavy_minus_sign:                                                                                       | Additional files bundled in the skill folder.                                                            |