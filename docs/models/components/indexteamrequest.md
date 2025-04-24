# IndexTeamRequest

Info about a team and optional version for that info


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Team`                                                                             | [components.TeamInfoDefinition](../../models/components/teaminfodefinition.md)     | :heavy_check_mark:                                                                 | Information about an employee's team                                               |
| `Version`                                                                          | **int64*                                                                           | :heavy_minus_sign:                                                                 | Version number for the team object. If absent or 0 then no version checks are done |