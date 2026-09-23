# PlatformAgentsCreateRunResponsesResponse


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `HTTPMeta`                                                                                  | [components.HTTPMetadata](../../models/components/httpmetadata.md)                          | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `PlatformAgentRunResponse`                                                                  | [*components.PlatformAgentRunResponse](../../models/components/platformagentrunresponse.md) | :heavy_minus_sign:                                                                          | Decisions accepted, or an identical retry. Poll GET run for continued execution.            |