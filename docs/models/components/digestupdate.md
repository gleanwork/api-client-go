# DigestUpdate


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `URL`                                                                  | **string*                                                              | :heavy_minus_sign:                                                     | URL link to the content or document.                                   |
| `Title`                                                                | **string*                                                              | :heavy_minus_sign:                                                     | Title or headline of the update.                                       |
| `Datasource`                                                           | **string*                                                              | :heavy_minus_sign:                                                     | Name or identifier of the data source (e.g., slack, confluence, etc.). |
| `Summary`                                                              | **string*                                                              | :heavy_minus_sign:                                                     | Brief summary or description of the update content.                    |
| `Type`                                                                 | [*components.UpdateType](../../models/components/updatetype.md)        | :heavy_minus_sign:                                                     | Optional type classification for the update.                           |