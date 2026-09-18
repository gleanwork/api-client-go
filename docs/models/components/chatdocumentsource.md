# ChatDocumentSource


## Supported Types

### ChatDocumentSourceDocument1

```go
chatDocumentSource := components.CreateChatDocumentSourceChatDocumentSourceDocument1(components.ChatDocumentSourceDocument1{/* values here */})
```

### ChatDocumentSourceDocument2

```go
chatDocumentSource := components.CreateChatDocumentSourceChatDocumentSourceDocument2(components.ChatDocumentSourceDocument2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatDocumentSource.Type {
	case components.ChatDocumentSourceTypeChatDocumentSourceDocument1:
		// chatDocumentSource.ChatDocumentSourceDocument1 is populated
	case components.ChatDocumentSourceTypeChatDocumentSourceDocument2:
		// chatDocumentSource.ChatDocumentSourceDocument2 is populated
}
```
