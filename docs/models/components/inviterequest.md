# InviteRequest

A request to send an invite to the specified user[s]


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `Channel`                                                                             | [*components.CommunicationChannel](../../models/components/communicationchannel.md)   | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `Template`                                                                            | [*components.CommunicationTemplate](../../models/components/communicationtemplate.md) | :heavy_minus_sign:                                                                    | The type of email to send                                                             |
| `Recipients`                                                                          | [][components.Person](../../models/components/person.md)                              | :heavy_minus_sign:                                                                    | The people who should receive this invite                                             |
| `RecipientFilters`                                                                    | [*components.PeopleFilters](../../models/components/peoplefilters.md)                 | :heavy_minus_sign:                                                                    | N/A                                                                                   |