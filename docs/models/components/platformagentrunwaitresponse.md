# PlatformAgentRunWaitResponse


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Run`                                                                       | [*components.PlatformAgentRun](../../models/components/platformagentrun.md) | :heavy_minus_sign:                                                          | N/A                                                                         |
| `Messages`                                                                  | [][components.PlatformMessage](../../models/components/platformmessage.md)  | :heavy_minus_sign:                                                          | Messages returned by the completed run.                                     |
| `RequestID`                                                                 | `string`                                                                    | :heavy_check_mark:                                                          | Platform-generated request ID for support correlation.                      |