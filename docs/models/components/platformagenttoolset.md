# PlatformAgentToolset

A tool provider and the subset of its tools available to the agent.


## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ToolProviderID`                                                                               | `string`                                                                                       | :heavy_check_mark:                                                                             | ID of the tool provider, or `native` for built-in Glean tools.                                 |
| `SelectedTools`                                                                                | [][components.PlatformAgentSelectedTool](../../models/components/platformagentselectedtool.md) | :heavy_minus_sign:                                                                             | Tools from the provider that are exposed to the agent.                                         |
| `CustomisationData`                                                                            | [*components.CustomisationData](../../models/components/customisationdata.md)                  | :heavy_minus_sign:                                                                             | Settings applied to every tool in this provider entry.                                         |