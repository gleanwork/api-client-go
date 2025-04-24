# FavoriteInfo


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `UgcType`                                                 | [*components.UgcType](../../models/components/ugctype.md) | :heavy_minus_sign:                                        | N/A                                                       |
| `ID`                                                      | **string*                                                 | :heavy_minus_sign:                                        | Opaque id of the UGC.                                     |
| `Count`                                                   | **int64*                                                  | :heavy_minus_sign:                                        | Number of users this object has been favorited by.        |
| `FavoritedByUser`                                         | **bool*                                                   | :heavy_minus_sign:                                        | If the requesting user has favorited this object.         |