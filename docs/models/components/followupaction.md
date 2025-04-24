# FollowupAction

A follow-up action that can be invoked by the user after a response. The action parameters are not included and need to be predicted/filled separately.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ActionRunID`                                                                  | **string*                                                                      | :heavy_minus_sign:                                                             | Unique identifier for this actionRun recommendation event.                     |
| `ActionInstanceID`                                                             | **string*                                                                      | :heavy_minus_sign:                                                             | The ID of the action instance that will be invoked.                            |
| `ActionID`                                                                     | **string*                                                                      | :heavy_minus_sign:                                                             | The ID of the associated action.                                               |
| `RecommendationText`                                                           | **string*                                                                      | :heavy_minus_sign:                                                             | Text to be displayed to the user when recommending the action instance.        |
| `ActionLabel`                                                                  | **string*                                                                      | :heavy_minus_sign:                                                             | The label to be used when displaying a button to execute this action instance. |
| `UserConfirmationRequired`                                                     | **bool*                                                                        | :heavy_minus_sign:                                                             | Whether user confirmation is needed before executing this action instance.     |