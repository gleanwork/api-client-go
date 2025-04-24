# QueryInsight


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Query`                                                              | *string*                                                             | :heavy_check_mark:                                                   | The query string the information is about.                           |
| `SearchCount`                                                        | [*components.CountInfo](../../models/components/countinfo.md)        | :heavy_minus_sign:                                                   | N/A                                                                  |
| `SearchorCount`                                                      | [*components.CountInfo](../../models/components/countinfo.md)        | :heavy_minus_sign:                                                   | N/A                                                                  |
| `SearchWithClickCount`                                               | [*components.CountInfo](../../models/components/countinfo.md)        | :heavy_minus_sign:                                                   | N/A                                                                  |
| `ClickCount`                                                         | [*components.CountInfo](../../models/components/countinfo.md)        | :heavy_minus_sign:                                                   | N/A                                                                  |
| `SimilarQueries`                                                     | [][components.QueryInsight](../../models/components/queryinsight.md) | :heavy_minus_sign:                                                   | list of similar queries to current one.                              |