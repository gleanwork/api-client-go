# PlatformTriggerPresetEventSearchRequest


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Inputs`                                                    | map[string]`string`                                         | :heavy_minus_sign:                                          | Values for the preset's input fields, keyed by field name.<br/> | {<br/>"repository": "acme/payments-api"<br/>}               |
| `PageSize`                                                  | `*int64`                                                    | :heavy_minus_sign:                                          | Maximum number of events to return.                         |                                                             |