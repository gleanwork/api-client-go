# WritePermission

Describes the write permissions levels that a user has for a specific feature


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ScopeType`                                                                          | [*components.ScopeType](../../models/components/scopetype.md)                        | :heavy_minus_sign:                                                                   | Describes the scope for a ReadPermission, WritePermission, or GrantPermission object |
| `Create`                                                                             | **bool*                                                                              | :heavy_minus_sign:                                                                   | True if user has create permission for this feature and scope                        |
| `Update`                                                                             | **bool*                                                                              | :heavy_minus_sign:                                                                   | True if user has update permission for this feature and scope                        |
| `Delete`                                                                             | **bool*                                                                              | :heavy_minus_sign:                                                                   | True if user has delete permission for this feature and scope                        |