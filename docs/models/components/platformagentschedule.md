# PlatformAgentSchedule

Wall-clock schedule for INPUT_FORM and CHAT_MESSAGE agents.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Frequency`                                                                          | `string`                                                                             | :heavy_check_mark:                                                                   | Five-field Unix cron expression, for example `0 9 * * 1-5`.                          |
| `Timezone`                                                                           | `string`                                                                             | :heavy_check_mark:                                                                   | IANA time zone name, for example `America/Los_Angeles`.                              |
| `StartAt`                                                                            | [*time.Time](https://pkg.go.dev/time#Time)                                           | :heavy_minus_sign:                                                                   | ISO 8601 timestamp with an explicit offset before which the schedule stays dormant.<br/> |