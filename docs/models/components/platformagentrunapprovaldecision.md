# PlatformAgentRunApprovalDecision

Approve or reject the invocation identified by a pending interaction ID.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `InteractionID`                                                                    | `string`                                                                           | :heavy_check_mark:                                                                 | The interaction_id returned in the run's pending_interactions.                     |
| `Decision`                                                                         | [components.Decision](../../models/components/decision.md)                         | :heavy_check_mark:                                                                 | The decision for this invocation only. Rejection follows normal workflow behavior. |