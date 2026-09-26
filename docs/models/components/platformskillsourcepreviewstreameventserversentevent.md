# PlatformSkillSourcePreviewStreamEventServerSentEvent

A typed server-sent event.


## Supported Types

### PlatformSkillSourcePreviewStreamScanServerSentEvent

```go
platformSkillSourcePreviewStreamEventServerSentEvent := components.CreatePlatformSkillSourcePreviewStreamEventServerSentEventScan(components.PlatformSkillSourcePreviewStreamScanServerSentEvent{/* values here */})
```

### PlatformSkillSourcePreviewStreamProgressServerSentEvent

```go
platformSkillSourcePreviewStreamEventServerSentEvent := components.CreatePlatformSkillSourcePreviewStreamEventServerSentEventProgress(components.PlatformSkillSourcePreviewStreamProgressServerSentEvent{/* values here */})
```

### PlatformSkillSourcePreviewStreamSkillServerSentEvent

```go
platformSkillSourcePreviewStreamEventServerSentEvent := components.CreatePlatformSkillSourcePreviewStreamEventServerSentEventSkill(components.PlatformSkillSourcePreviewStreamSkillServerSentEvent{/* values here */})
```

### PlatformSkillSourcePreviewStreamResultServerSentEvent

```go
platformSkillSourcePreviewStreamEventServerSentEvent := components.CreatePlatformSkillSourcePreviewStreamEventServerSentEventResult(components.PlatformSkillSourcePreviewStreamResultServerSentEvent{/* values here */})
```

### PlatformSkillSourcePreviewStreamErrorServerSentEvent

```go
platformSkillSourcePreviewStreamEventServerSentEvent := components.CreatePlatformSkillSourcePreviewStreamEventServerSentEventError(components.PlatformSkillSourcePreviewStreamErrorServerSentEvent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch platformSkillSourcePreviewStreamEventServerSentEvent.Type {
	case components.PlatformSkillSourcePreviewStreamEventServerSentEventTypeScan:
		// platformSkillSourcePreviewStreamEventServerSentEvent.PlatformSkillSourcePreviewStreamScanServerSentEvent is populated
	case components.PlatformSkillSourcePreviewStreamEventServerSentEventTypeProgress:
		// platformSkillSourcePreviewStreamEventServerSentEvent.PlatformSkillSourcePreviewStreamProgressServerSentEvent is populated
	case components.PlatformSkillSourcePreviewStreamEventServerSentEventTypeSkill:
		// platformSkillSourcePreviewStreamEventServerSentEvent.PlatformSkillSourcePreviewStreamSkillServerSentEvent is populated
	case components.PlatformSkillSourcePreviewStreamEventServerSentEventTypeResult:
		// platformSkillSourcePreviewStreamEventServerSentEvent.PlatformSkillSourcePreviewStreamResultServerSentEvent is populated
	case components.PlatformSkillSourcePreviewStreamEventServerSentEventTypeError:
		// platformSkillSourcePreviewStreamEventServerSentEvent.PlatformSkillSourcePreviewStreamErrorServerSentEvent is populated
}
```
