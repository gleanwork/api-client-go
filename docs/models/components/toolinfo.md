# ToolInfo


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `Metadata`                                                                                    | [*components.ToolMetadata](../../models/components/toolmetadata.md)                           | :heavy_minus_sign:                                                                            | The manifest for a tool that can be used to augment Glean Assistant.                          |
| `Parameters`                                                                                  | map[string][components.WriteActionParameter](../../models/components/writeactionparameter.md) | :heavy_minus_sign:                                                                            | Parameters supported by the tool.                                                             |