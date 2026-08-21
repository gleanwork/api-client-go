# PlatformChatCitationSource

Four-variant citation source union.


## Supported Types

### PlatformChatDocumentSource

```go
platformChatCitationSource := components.CreatePlatformChatCitationSourceDocument(components.PlatformChatDocumentSource{/* values here */})
```

### PlatformChatPersonSource

```go
platformChatCitationSource := components.CreatePlatformChatCitationSourcePerson(components.PlatformChatPersonSource{/* values here */})
```

### PlatformChatFileSource

```go
platformChatCitationSource := components.CreatePlatformChatCitationSourceFile(components.PlatformChatFileSource{/* values here */})
```

### PlatformChatCustomEntitySource

```go
platformChatCitationSource := components.CreatePlatformChatCitationSourceCustomEntity(components.PlatformChatCustomEntitySource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch platformChatCitationSource.Type {
	case components.PlatformChatCitationSourceTypeDocument:
		// platformChatCitationSource.PlatformChatDocumentSource is populated
	case components.PlatformChatCitationSourceTypePerson:
		// platformChatCitationSource.PlatformChatPersonSource is populated
	case components.PlatformChatCitationSourceTypeFile:
		// platformChatCitationSource.PlatformChatFileSource is populated
	case components.PlatformChatCitationSourceTypeCustomEntity:
		// platformChatCitationSource.PlatformChatCustomEntitySource is populated
}
```
