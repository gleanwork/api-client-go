# AnswerLike


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `User`                                                       | [*components.Person](../../models/components/person.md)      | :heavy_minus_sign:                                           | N/A                                                          | {<br/>"name": "George Clooney",<br/>"obfuscatedId": "abc123"<br/>} |
| `CreateTime`                                                 | [*time.Time](https://pkg.go.dev/time#Time)                   | :heavy_minus_sign:                                           | The time the user liked the answer in ISO format (ISO 8601). |                                                              |