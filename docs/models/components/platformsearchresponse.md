# PlatformSearchResponse


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Results`                                                                | [][components.PlatformResult](../../models/components/platformresult.md) | :heavy_check_mark:                                                       | Ordered list of search results.                                          |
| `HasMore`                                                                | `bool`                                                                   | :heavy_check_mark:                                                       | Indicates whether additional pages of results are available.             |
| `NextCursor`                                                             | `*string`                                                                | :heavy_check_mark:                                                       | Opaque token to pass as `cursor` in the next request.                    |
| `RequestID`                                                              | `string`                                                                 | :heavy_check_mark:                                                       | Platform-generated request ID for support correlation.                   |