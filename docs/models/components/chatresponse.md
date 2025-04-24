# ChatResponse

A single response from the /chat backend.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Messages`                                                           | [][components.ChatMessage](../../models/components/chatmessage.md)   | :heavy_minus_sign:                                                   | N/A                                                                  |                                                                      |
| `ChatID`                                                             | **string*                                                            | :heavy_minus_sign:                                                   | The id of the associated Chat the messages belong to, if one exists. |                                                                      |
| `FollowUpPrompts`                                                    | []*string*                                                           | :heavy_minus_sign:                                                   | Follow-up prompts for the user to potentially use                    |                                                                      |
| `BackendTimeMillis`                                                  | **int64*                                                             | :heavy_minus_sign:                                                   | Time in milliseconds the backend took to respond to the request.     | 1100                                                                 |
| `ChatSessionTrackingToken`                                           | **string*                                                            | :heavy_minus_sign:                                                   | A token that is used to track the session.                           |                                                                      |