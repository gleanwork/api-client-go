# Reaction


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Type`                                                              | **string*                                                           | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Count`                                                             | **int64*                                                            | :heavy_minus_sign:                                                  | The count of the reaction type on the document.                     |
| `Reactors`                                                          | [][components.Person](../../models/components/person.md)            | :heavy_minus_sign:                                                  | N/A                                                                 |
| `ReactedByViewer`                                                   | **bool*                                                             | :heavy_minus_sign:                                                  | Whether the user in context reacted with this type to the document. |