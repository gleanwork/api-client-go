# DlpReportData

Dlp report metadata which is used to construct report email


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Frequency`                                                               | [*components.Frequency](../../models/components/frequency.md)             | :heavy_minus_sign:                                                        | The frequency of the report                                               |
| `RequestTime`                                                             | [*time.Time](https://pkg.go.dev/time#Time)                                | :heavy_minus_sign:                                                        | The time the report was requested, applicable only for one time reports   |
| `ReportName`                                                              | **string*                                                                 | :heavy_minus_sign:                                                        | N/A                                                                       |
| `Status`                                                                  | [*components.DlpSimpleResult](../../models/components/dlpsimpleresult.md) | :heavy_minus_sign:                                                        | N/A                                                                       |