# PlatformAgentRunStreamMessage

A message emitted during a streaming agent run. Unlike `Message`, `content` is optional: the server omits it when an event has no text and no citation annotations.



## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    | Example                                                                                        |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `Role`                                                                                         | [components.PlatformMessageRole](../../models/components/platformmessagerole.md)               | :heavy_check_mark:                                                                             | Role of the message author.                                                                    | USER                                                                                           |
| `Content`                                                                                      | [][components.PlatformMessageTextBlock](../../models/components/platformmessagetextblock.md)   | :heavy_minus_sign:                                                                             | Content blocks in the message. Omitted when the event has no text and no citation annotations. |                                                                                                |