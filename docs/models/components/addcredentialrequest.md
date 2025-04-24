# AddCredentialRequest


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Datasource`                                                     | **string*                                                        | :heavy_minus_sign:                                               | the datasource the credential applies to                         |
| `DatasourceInstance`                                             | **string*                                                        | :heavy_minus_sign:                                               | the datasource instance the credential applies to                |
| `User`                                                           | **string*                                                        | :heavy_minus_sign:                                               | the user info (email or username for example) for the credential |
| `Token`                                                          | **string*                                                        | :heavy_minus_sign:                                               | the token part of the credential (password, apiToken etc)        |
| `Metadata`                                                       | **string*                                                        | :heavy_minus_sign:                                               | any metadata associated with the user credential                 |