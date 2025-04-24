# PersonDistance


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Name`                                                                               | *string*                                                                             | :heavy_check_mark:                                                                   | The display name.                                                                    |
| `ObfuscatedID`                                                                       | *string*                                                                             | :heavy_check_mark:                                                                   | An opaque identifier that can be used to request metadata for a Person.              |
| `Distance`                                                                           | *float32*                                                                            | :heavy_check_mark:                                                                   | Distance to person, refer to PeopleDistance pipeline on interpretation of the value. |