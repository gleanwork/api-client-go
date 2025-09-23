# Group


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Type`                                                                | [components.GroupType](../../models/components/grouptype.md)          | :heavy_check_mark:                                                    | The type of user group                                                |
| `ID`                                                                  | *string*                                                              | :heavy_check_mark:                                                    | A unique identifier for the group. May be the same as name.           |
| `Name`                                                                | **string*                                                             | :heavy_minus_sign:                                                    | Name of the group.                                                    |
| `DatasourceInstance`                                                  | **string*                                                             | :heavy_minus_sign:                                                    | Datasource instance if the group belongs to one e.g. external groups. |
| `ProvisioningID`                                                      | **string*                                                             | :heavy_minus_sign:                                                    | identifier for greenlist provisioning, aka sciokey                    |