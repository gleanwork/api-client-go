# InsightsAssistantRequest


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Departments`                                                                      | []`string`                                                                         | :heavy_minus_sign:                                                                 | Departments for which Insights are requested.                                      |
| `ManagerEmails`                                                                    | []`string`                                                                         | :heavy_minus_sign:                                                                 | Manager emails whose teams should be filtered for. Empty array means no filtering. |
| `DayRange`                                                                         | [*components.Period](../../models/components/period.md)                            | :heavy_minus_sign:                                                                 | N/A                                                                                |