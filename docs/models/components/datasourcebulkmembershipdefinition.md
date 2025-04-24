# DatasourceBulkMembershipDefinition

describes the membership row of a group in the bulk uploaded. Only one of memberUserId and memberGroupName can be specified.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `MemberUserID`                                                        | **string*                                                             | :heavy_minus_sign:                                                    | If the member is a user, then the email or datasource id for the user |
| `MemberGroupName`                                                     | **string*                                                             | :heavy_minus_sign:                                                    | If the member is a group, then the name of the member group           |