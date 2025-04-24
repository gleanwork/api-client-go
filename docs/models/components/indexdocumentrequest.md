# IndexDocumentRequest

Describes the request body of the /indexdocument API call


## Fields

| Field                                                                                                           | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `Version`                                                                                                       | **int64*                                                                                                        | :heavy_minus_sign:                                                                                              | Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done. |
| `Document`                                                                                                      | [components.DocumentDefinition](../../models/components/documentdefinition.md)                                  | :heavy_check_mark:                                                                                              | Indexable document structure                                                                                    |