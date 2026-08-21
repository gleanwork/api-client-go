# PlatformTriggerPresetsListRequest


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Datasource`                                                       | `*string`                                                          | :heavy_minus_sign:                                                 | Restrict results to presets for a single datasource (e.g. github). |
| `PageSize`                                                         | `*int64`                                                           | :heavy_minus_sign:                                                 | Maximum number of presets to return.                               |
| `Cursor`                                                           | `*string`                                                          | :heavy_minus_sign:                                                 | Opaque pagination cursor from a previous response.                 |