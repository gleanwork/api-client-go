# DebugDocumentsResponseItem

Describes the response body of a single document in the /debug/{datasource}/documents API call


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `DocID`                                                                               | **string*                                                                             | :heavy_minus_sign:                                                                    | Id of the document                                                                    |
| `ObjectType`                                                                          | **string*                                                                             | :heavy_minus_sign:                                                                    | objectType of the document                                                            |
| `DebugInfo`                                                                           | [*components.DebugDocumentResponse](../../models/components/debugdocumentresponse.md) | :heavy_minus_sign:                                                                    | Describes the response body of the /debug/{datasource}/document API call              |