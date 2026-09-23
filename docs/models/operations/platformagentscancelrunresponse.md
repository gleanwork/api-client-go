# PlatformAgentsCancelRunResponse


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `HTTPMeta`                                                                                  | [components.HTTPMetadata](../../models/components/httpmetadata.md)                          | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `PlatformAgentRunResponse`                                                                  | [*components.PlatformAgentRunResponse](../../models/components/platformagentrunresponse.md) | :heavy_minus_sign:                                                                          | Current run snapshot after requesting cancellation, or an unchanged terminal run.           |