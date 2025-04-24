# StructuredLink

The display configuration for a link.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       | Example                                                                           |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `Name`                                                                            | **string*                                                                         | :heavy_minus_sign:                                                                | The display name for the link                                                     |                                                                                   |
| `URL`                                                                             | **string*                                                                         | :heavy_minus_sign:                                                                | The URL for the link.                                                             |                                                                                   |
| `IconConfig`                                                                      | [*components.IconConfig](../../models/components/iconconfig.md)                   | :heavy_minus_sign:                                                                | Defines how to render an icon                                                     | {<br/>"color": "#343CED",<br/>"key": "person_icon",<br/>"iconType": "GLYPH",<br/>"name": "user"<br/>} |