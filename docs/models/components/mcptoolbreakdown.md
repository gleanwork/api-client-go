# McpToolBreakdown


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Tool`                                                                   | `*string`                                                                | :heavy_minus_sign:                                                       | MCP tool name.                                                           |
| `TotalCalls`                                                             | `*int64`                                                                 | :heavy_minus_sign:                                                       | Total number of MCP calls for this tool in the specified time period.    |
| `ActiveUsers`                                                            | `*int64`                                                                 | :heavy_minus_sign:                                                       | Total number of active users for this tool in the specified time period. |
| `HostApplications`                                                       | []`string`                                                               | :heavy_minus_sign:                                                       | Host applications using this tool in the specified time period.          |