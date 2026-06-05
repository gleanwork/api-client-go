# ChatSuggestion


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Query`                                                                 | `*string`                                                               | :heavy_minus_sign:                                                      | The actionable chat query to run when the user selects this suggestion. |
| `Cta`                                                                   | `*string`                                                               | :heavy_minus_sign:                                                      | Button text to show for the suggestion action.                          |
| `Feature`                                                               | `*string`                                                               | :heavy_minus_sign:                                                      | Targeted Glean Chat feature for the suggestion.                         |
| `SourceDocumentIds`                                                     | []`string`                                                              | :heavy_minus_sign:                                                      | Document IDs that grounded the suggestion.                              |