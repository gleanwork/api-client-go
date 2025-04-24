# ContentDefinition

Describes text content or base64 encoded binary content


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `MimeType`                                                                                | *string*                                                                                  | :heavy_check_mark:                                                                        | N/A                                                                                       |
| `TextContent`                                                                             | **string*                                                                                 | :heavy_minus_sign:                                                                        | text content. Only one of textContent or binary content can be specified                  |
| `BinaryContent`                                                                           | **string*                                                                                 | :heavy_minus_sign:                                                                        | base64 encoded binary content. Only one of textContent or binary content can be specified |