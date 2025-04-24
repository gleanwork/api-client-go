# PublicConfigRequest

Will only send back publicly available config and will ignore other keys


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ThemeKeys`                                             | []*string*                                              | :heavy_minus_sign:                                      | A list of theme keys to include in the response.        |
| `BoolKeys`                                              | []*string*                                              | :heavy_minus_sign:                                      | A list of boolean flag keys to include in the response. |
| `IntegerKeys`                                           | []*string*                                              | :heavy_minus_sign:                                      | A list of integer flag keys to include in the response. |