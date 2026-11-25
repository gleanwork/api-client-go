# PlatformAgentRunMessageEventData

Payload for the Agents SSE event `message`.


## Fields

| Field                                                                                                  | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `Messages`                                                                                             | [][components.PlatformAgentRunStreamMessage](../../models/components/platformagentrunstreammessage.md) | :heavy_check_mark:                                                                                     | Messages emitted by the running agent in this event.                                                   |