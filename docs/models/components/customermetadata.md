# CustomerMetadata


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `DatasourceID`                                                                      | **string*                                                                           | :heavy_minus_sign:                                                                  | The user visible id of the salesforce customer account.                             |
| `CustomData`                                                                        | map[string][components.CustomDataValue](../../models/components/customdatavalue.md) | :heavy_minus_sign:                                                                  | Custom fields specific to individual datasources                                    |