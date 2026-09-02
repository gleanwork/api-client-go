# PlatformChatStreamEventServerSentEvent

A typed server-sent event.


## Supported Types

### PlatformChatStreamResponseCreatedServerSentEvent

```go
platformChatStreamEventServerSentEvent := components.CreatePlatformChatStreamEventServerSentEventResponseCreated(components.PlatformChatStreamResponseCreatedServerSentEvent{/* values here */})
```

### PlatformChatStreamProgressServerSentEvent

```go
platformChatStreamEventServerSentEvent := components.CreatePlatformChatStreamEventServerSentEventResponseProgress(components.PlatformChatStreamProgressServerSentEvent{/* values here */})
```

### PlatformChatStreamOutputTextDeltaServerSentEvent

```go
platformChatStreamEventServerSentEvent := components.CreatePlatformChatStreamEventServerSentEventResponseOutputTextDelta(components.PlatformChatStreamOutputTextDeltaServerSentEvent{/* values here */})
```

### PlatformChatStreamOutputTextDoneServerSentEvent

```go
platformChatStreamEventServerSentEvent := components.CreatePlatformChatStreamEventServerSentEventResponseOutputTextDone(components.PlatformChatStreamOutputTextDoneServerSentEvent{/* values here */})
```

### PlatformChatStreamResponseCompletedServerSentEvent

```go
platformChatStreamEventServerSentEvent := components.CreatePlatformChatStreamEventServerSentEventResponseCompleted(components.PlatformChatStreamResponseCompletedServerSentEvent{/* values here */})
```

### PlatformChatStreamResponseFailedServerSentEvent

```go
platformChatStreamEventServerSentEvent := components.CreatePlatformChatStreamEventServerSentEventResponseFailed(components.PlatformChatStreamResponseFailedServerSentEvent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch platformChatStreamEventServerSentEvent.Type {
	case components.PlatformChatStreamEventServerSentEventTypeResponseCreated:
		// platformChatStreamEventServerSentEvent.PlatformChatStreamResponseCreatedServerSentEvent is populated
	case components.PlatformChatStreamEventServerSentEventTypeResponseProgress:
		// platformChatStreamEventServerSentEvent.PlatformChatStreamProgressServerSentEvent is populated
	case components.PlatformChatStreamEventServerSentEventTypeResponseOutputTextDelta:
		// platformChatStreamEventServerSentEvent.PlatformChatStreamOutputTextDeltaServerSentEvent is populated
	case components.PlatformChatStreamEventServerSentEventTypeResponseOutputTextDone:
		// platformChatStreamEventServerSentEvent.PlatformChatStreamOutputTextDoneServerSentEvent is populated
	case components.PlatformChatStreamEventServerSentEventTypeResponseCompleted:
		// platformChatStreamEventServerSentEvent.PlatformChatStreamResponseCompletedServerSentEvent is populated
	case components.PlatformChatStreamEventServerSentEventTypeResponseFailed:
		// platformChatStreamEventServerSentEvent.PlatformChatStreamResponseFailedServerSentEvent is populated
}
```
