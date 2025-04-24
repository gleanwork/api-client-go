# TimePoint


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `EpochSeconds`                                                                        | **int64*                                                                              | :heavy_minus_sign:                                                                    | Epoch seconds. Has precedence over daysFromNow.                                       |
| `DaysFromNow`                                                                         | **int64*                                                                              | :heavy_minus_sign:                                                                    | The number of days from now. Specification relative to current time. Can be negative. |