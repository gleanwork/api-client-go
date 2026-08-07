# PlatformChatDocumentSource


## Supported Types

### PlatformChatDocumentSourceDocument1

```go
platformChatDocumentSource := components.CreatePlatformChatDocumentSourcePlatformChatDocumentSourceDocument1(components.PlatformChatDocumentSourceDocument1{/* values here */})
```

### PlatformChatDocumentSourceDocument2

```go
platformChatDocumentSource := components.CreatePlatformChatDocumentSourcePlatformChatDocumentSourceDocument2(components.PlatformChatDocumentSourceDocument2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch platformChatDocumentSource.Type {
	case components.PlatformChatDocumentSourceTypePlatformChatDocumentSourceDocument1:
		// platformChatDocumentSource.PlatformChatDocumentSourceDocument1 is populated
	case components.PlatformChatDocumentSourceTypePlatformChatDocumentSourceDocument2:
		// platformChatDocumentSource.PlatformChatDocumentSourceDocument2 is populated
}
```
