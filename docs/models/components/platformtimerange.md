# PlatformTimeRange

Filter results to those last updated within this range.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Start`                                    | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | Inclusive lower bound in ISO 8601 format.  |
| `End`                                      | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | Exclusive upper bound in ISO 8601 format.  |