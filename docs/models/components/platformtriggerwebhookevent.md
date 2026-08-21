# PlatformTriggerWebhookEvent

Every delivered webhook is one of these two variants. A trigger with a cron schedule carries no document and so cannot be delivered over a webhook.


## Supported Types

### PlatformDocumentChangeWebhookEvent

```go
platformTriggerWebhookEvent := components.CreatePlatformTriggerWebhookEventDocumentChange(components.PlatformDocumentChangeWebhookEvent{/* values here */})
```

### PlatformContentScheduleWebhookEvent

```go
platformTriggerWebhookEvent := components.CreatePlatformTriggerWebhookEventContentSchedule(components.PlatformContentScheduleWebhookEvent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch platformTriggerWebhookEvent.Type {
	case components.PlatformTriggerWebhookEventTypeDocumentChange:
		// platformTriggerWebhookEvent.PlatformDocumentChangeWebhookEvent is populated
	case components.PlatformTriggerWebhookEventTypeContentSchedule:
		// platformTriggerWebhookEvent.PlatformContentScheduleWebhookEvent is populated
}
```
