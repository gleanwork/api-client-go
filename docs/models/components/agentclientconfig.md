# AgentClientConfig

Describes the configurations that GleanChat has based on an AgentConfig.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `AgentConfig`                                                                 | [*components.AgentConfig](../../models/components/agentconfig.md)             | :heavy_minus_sign:                                                            | Describes the agent that executes the request.                                |
| `InputCharLimit`                                                              | **int64*                                                                      | :heavy_minus_sign:                                                            | The character limit of an input to GleanChat under the specified AgentConfig. |