# IndexingShortcut


## Fields

| Field                                                                                                                                                                | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `InputAlias`                                                                                                                                                         | *string*                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                   | link text following the viewPrefix as entered by the user. For example, if the view prefix is `go/` and the shortened URL is `go/abc`, then `abc` is the inputAlias. |
| `Description`                                                                                                                                                        | **string*                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                   | A short, plain text blurb to help people understand the intent of the shortcut.                                                                                      |
| `DestinationURL`                                                                                                                                                     | *string*                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                   | destination URL for the shortcut.                                                                                                                                    |
| `CreatedBy`                                                                                                                                                          | *string*                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                   | Email of the user who created this shortcut.                                                                                                                         |
| `CreateTime`                                                                                                                                                         | **int64*                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                   | The time the shortcut was created in epoch seconds.                                                                                                                  |
| `UpdatedBy`                                                                                                                                                          | **string*                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                   | Email of the user who last updated this shortcut.                                                                                                                    |
| `UpdateTime`                                                                                                                                                         | **int64*                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                   | The time the shortcut was updated in epoch seconds.                                                                                                                  |
| `Unlisted`                                                                                                                                                           | **bool*                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                   | Whether this shortcut is unlisted or not. Unlisted shortcuts are visible to author and admins only.                                                                  |
| `URLTemplate`                                                                                                                                                        | **string*                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                   | For variable shortcuts, contains the URL template; note, `destinationUrl` contains default URL.                                                                      |