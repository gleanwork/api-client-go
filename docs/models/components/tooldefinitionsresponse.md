# ToolDefinitionsResponse


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Tools`                                                                  | [][components.ToolDefinition](../../models/components/tooldefinition.md) | :heavy_check_mark:                                                       | Definitions for the requested tools that exist on this server.           |
| `NotFound`                                                               | []`string`                                                               | :heavy_minus_sign:                                                       | Requested names that do not exist on this server.                        |