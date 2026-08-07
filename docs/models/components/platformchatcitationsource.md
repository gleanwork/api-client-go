# PlatformChatCitationSource

Four-variant citation source union.


## Supported Types

### PlatformChatDocumentSource

```go
platformChatCitationSource := components.CreatePlatformChatCitationSourcePlatformChatDocumentSource(components.PlatformChatDocumentSource{/* values here */})
```

### PlatformChatPersonSource

```go
platformChatCitationSource := components.CreatePlatformChatCitationSourcePlatformChatPersonSource(components.PlatformChatPersonSource{/* values here */})
```

### PlatformChatFileSource

```go
platformChatCitationSource := components.CreatePlatformChatCitationSourcePlatformChatFileSource(components.PlatformChatFileSource{/* values here */})
```

### PlatformChatCustomEntitySource

```go
platformChatCitationSource := components.CreatePlatformChatCitationSourcePlatformChatCustomEntitySource(components.PlatformChatCustomEntitySource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch platformChatCitationSource.Type {
	case components.PlatformChatCitationSourceTypePlatformChatDocumentSource:
		// platformChatCitationSource.PlatformChatDocumentSource is populated
	case components.PlatformChatCitationSourceTypePlatformChatPersonSource:
		// platformChatCitationSource.PlatformChatPersonSource is populated
	case components.PlatformChatCitationSourceTypePlatformChatFileSource:
		// platformChatCitationSource.PlatformChatFileSource is populated
	case components.PlatformChatCitationSourceTypePlatformChatCustomEntitySource:
		// platformChatCitationSource.PlatformChatCustomEntitySource is populated
}
```
