# McpServerBreakdown


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Server`                                                                   | `*string`                                                                  | :heavy_minus_sign:                                                         | MCP server name.                                                           |
| `TotalCalls`                                                               | `*int64`                                                                   | :heavy_minus_sign:                                                         | Total number of MCP calls for this server in the specified time period.    |
| `ActiveUsers`                                                              | `*int64`                                                                   | :heavy_minus_sign:                                                         | Total number of active users for this server in the specified time period. |
| `HostApplications`                                                         | []`string`                                                                 | :heavy_minus_sign:                                                         | Host applications using this server in the specified time period.          |