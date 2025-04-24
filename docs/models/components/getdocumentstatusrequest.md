# GetDocumentStatusRequest

Describes the request body for /getdocumentstatus API call


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Datasource`                                                   | *string*                                                       | :heavy_check_mark:                                             | Datasource to get fetch document status for                    |
| `ObjectType`                                                   | *string*                                                       | :heavy_check_mark:                                             | Object type of the document to get the status for              |
| `DocID`                                                        | *string*                                                       | :heavy_check_mark:                                             | Glean Document ID within the datasource to get the status for. |