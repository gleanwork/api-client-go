# PlatformAgentsGetRunResponse


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `HTTPMeta`                                                                                  | [components.HTTPMetadata](../../models/components/httpmetadata.md)                          | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `PlatformAgentRunResponse`                                                                  | [*components.PlatformAgentRunResponse](../../models/components/platformagentrunresponse.md) | :heavy_minus_sign:                                                                          | Current persisted run snapshot. A failed execution is still a successful retrieval.         |