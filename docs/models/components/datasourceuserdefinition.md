# DatasourceUserDefinition

describes a user in the datasource


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Email`                                                          | *string*                                                         | :heavy_check_mark:                                               | N/A                                                              |
| `UserID`                                                         | **string*                                                        | :heavy_minus_sign:                                               | To be supplied if the user id in the datasource is not the email |
| `Name`                                                           | *string*                                                         | :heavy_check_mark:                                               | N/A                                                              |
| `IsActive`                                                       | **bool*                                                          | :heavy_minus_sign:                                               | set to false if the user is a former employee or a bot           |