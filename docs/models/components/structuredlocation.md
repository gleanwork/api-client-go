# StructuredLocation

Detailed location with information about country, state, city etc.


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `DeskLocation`                                            | **string*                                                 | :heavy_minus_sign:                                        | Desk number.                                              |
| `Timezone`                                                | **string*                                                 | :heavy_minus_sign:                                        | Location's timezone, e.g. UTC, PST.                       |
| `Address`                                                 | **string*                                                 | :heavy_minus_sign:                                        | Office address or name.                                   |
| `City`                                                    | **string*                                                 | :heavy_minus_sign:                                        | Name of the city.                                         |
| `State`                                                   | **string*                                                 | :heavy_minus_sign:                                        | State code.                                               |
| `Region`                                                  | **string*                                                 | :heavy_minus_sign:                                        | Region information, e.g. NORAM, APAC.                     |
| `ZipCode`                                                 | **string*                                                 | :heavy_minus_sign:                                        | ZIP Code for the address.                                 |
| `Country`                                                 | **string*                                                 | :heavy_minus_sign:                                        | Country name.                                             |
| `CountryCode`                                             | **string*                                                 | :heavy_minus_sign:                                        | Alpha-2 or Alpha-3 ISO 3166 country code, e.g. US or USA. |