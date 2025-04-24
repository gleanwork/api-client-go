# TeamsResponse


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Results`                                                            | [][components.Team](../../models/components/team.md)                 | :heavy_minus_sign:                                                   | A Team and a deep copy of all its members for each ID in the request |
| `Errors`                                                             | []*string*                                                           | :heavy_minus_sign:                                                   | A list of IDs that could not be found.                               |