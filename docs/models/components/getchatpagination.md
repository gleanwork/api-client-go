# GetChatPagination

Pagination cursors and page flags for a getchat message page.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `StartCursor`                                                             | `*string`                                                                 | :heavy_minus_sign:                                                        | Cursor of the first (oldest) returned message. Omitted for an empty page. |
| `EndCursor`                                                               | `*string`                                                                 | :heavy_minus_sign:                                                        | Cursor of the last (newest) returned message. Omitted for an empty page.  |
| `HasNextPage`                                                             | `bool`                                                                    | :heavy_check_mark:                                                        | Whether visible messages exist after endCursor in chronological order.    |
| `HasPreviousPage`                                                         | `bool`                                                                    | :heavy_check_mark:                                                        | Whether visible messages exist before startCursor in chronological order. |