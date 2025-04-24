# IndexDocumentsRequest

Describes the request body of the /indexdocuments API call


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `UploadID`                                                                       | **string*                                                                        | :heavy_minus_sign:                                                               | Optional id parameter to identify and track a batch of documents.                |
| `Datasource`                                                                     | *string*                                                                         | :heavy_check_mark:                                                               | Datasource of the documents                                                      |
| `Documents`                                                                      | [][components.DocumentDefinition](../../models/components/documentdefinition.md) | :heavy_check_mark:                                                               | Batch of documents being added/updated                                           |