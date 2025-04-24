# CheckDocumentAccessRequest

Describes the request body of the /checkdocumentaccess API call


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `Datasource`                                 | *string*                                     | :heavy_check_mark:                           | Datasource of document to check access for.  |
| `ObjectType`                                 | *string*                                     | :heavy_check_mark:                           | Object type of document to check access for. |
| `DocID`                                      | *string*                                     | :heavy_check_mark:                           | Glean Document ID to check access for.       |
| `UserEmail`                                  | *string*                                     | :heavy_check_mark:                           | Email of user to check access for.           |