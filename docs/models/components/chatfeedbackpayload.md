# ChatFeedbackPayload

Payload for chat feedback reporting. Required when template is `CHAT_FEEDBACK`.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Rating`                                                                             | **string*                                                                            | :heavy_minus_sign:                                                                   | Rating given to the conversation (currently either "upvoted" or "downvoted").        |
| `Issues`                                                                             | []*string*                                                                           | :heavy_minus_sign:                                                                   | The type(s) of issue being reported.                                                 |
| `Comments`                                                                           | **string*                                                                            | :heavy_minus_sign:                                                                   | Additional freeform comments provided by the reporter.                               |
| `PreviousMessages`                                                                   | []*string*                                                                           | :heavy_minus_sign:                                                                   | Previous messages in this conversation.                                              |
| `ChatTranscript`                                                                     | [][components.FeedbackChatExchange](../../models/components/feedbackchatexchange.md) | :heavy_minus_sign:                                                                   | N/A                                                                                  |