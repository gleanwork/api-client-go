# PlatformChatResumeResponse


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `HTTPMeta`                                                               | [components.HTTPMetadata](../../models/components/httpmetadata.md)       | :heavy_check_mark:                                                       | N/A                                                                      |
| `PlatformChatStreamEventServerSentEvent`                                 | `*stream.EventStream[components.PlatformChatStreamEventServerSentEvent]` | :heavy_minus_sign:                                                       | Server-sent events for the chat.                                         |