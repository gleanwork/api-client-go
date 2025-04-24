# AiAppActionCounts

Map from action to frequency.


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `TotalSlackbotResponses`                                                | **int64*                                                                | :heavy_minus_sign:                                                      | Total number of Slackbot responses, both proactive and reactive.        |
| `TotalSlackbotResponsesShared`                                          | **int64*                                                                | :heavy_minus_sign:                                                      | Total number of Slackbot responses shared publicly (upvoted).           |
| `TotalSlackbotResponsesNotHelpful`                                      | **int64*                                                                | :heavy_minus_sign:                                                      | Total number of Slackbot responses rejected as not helpful (downvoted). |
| `TotalChatMessages`                                                     | **int64*                                                                | :heavy_minus_sign:                                                      | Total number of Chat messages sent in requested period.                 |
| `TotalUpvotes`                                                          | **int64*                                                                | :heavy_minus_sign:                                                      | Total number of Chat messages which received upvotes by the user.       |
| `TotalDownvotes`                                                        | **int64*                                                                | :heavy_minus_sign:                                                      | Total number of Chat messages which received downvotes by the user.     |
| `AdditionalProperties`                                                  | map[string]*int64*                                                      | :heavy_minus_sign:                                                      | N/A                                                                     |