# TimePoint


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `EpochSeconds`                                            | `*int64`                                                  | :heavy_minus_sign:                                        | Epoch seconds. Has precedence over daysFromNow.           |
| `DaysFromNow`                                             | `*int64`                                                  | :heavy_minus_sign:                                        | Number of days in the past, relative to the current date. |