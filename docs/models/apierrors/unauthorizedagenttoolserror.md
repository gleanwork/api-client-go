# UnauthorizedAgentToolsError

Returned when the agent has tools the caller must authorize before the run can start. Each entry in `authenticationSuggestions` names one such tool; POST its `serverId` to `/tool-servers/{serverId}/auth` with `returnUrl` in the request body to obtain an `authorizationUrl` to redirect the end user to, then retry the run once OAuth completes.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Message`                                                                      | `string`                                                                       | :heavy_check_mark:                                                             | Human-readable description of the failure.                                     |
| `AuthenticationSuggestions`                                                    | [][components.ServerToolRequest](../../models/components/servertoolrequest.md) | :heavy_check_mark:                                                             | One entry per tool the caller must authorize.                                  |
| `HTTPMeta`                                                                     | [components.HTTPMetadata](../../models/components/httpmetadata.md)             | :heavy_check_mark:                                                             | N/A                                                                            |