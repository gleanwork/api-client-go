# PlatformActionSummary


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `ToolID`                                           | `string`                                           | :heavy_check_mark:                                 | Unique identifier of the action.                   |
| `DisplayName`                                      | `string`                                           | :heavy_check_mark:                                 | Display name of the action.                        |
| `Type`                                             | `*string`                                          | :heavy_minus_sign:                                 | Tool type.                                         |
| `AuthType`                                         | `*string`                                          | :heavy_minus_sign:                                 | Authentication type required by the action.        |
| `WriteActionType`                                  | `*string`                                          | :heavy_minus_sign:                                 | Write-action execution type.                       |
| `IsSetupFinished`                                  | `*bool`                                            | :heavy_minus_sign:                                 | Whether this action has been fully configured.     |
| `DataSource`                                       | `*string`                                          | :heavy_minus_sign:                                 | Kind of knowledge the action accesses or modifies. |