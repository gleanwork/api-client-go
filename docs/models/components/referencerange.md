# ReferenceRange

Each text range from the response can correspond to an array of snippets from the citation source.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `TextRange`                                                                        | [*components.TextRange](../../models/components/textrange.md)                      | :heavy_minus_sign:                                                                 | A subsection of a given string to which some special formatting should be applied. |
| `Snippets`                                                                         | [][components.SearchResultSnippet](../../models/components/searchresultsnippet.md) | :heavy_minus_sign:                                                                 | N/A                                                                                |