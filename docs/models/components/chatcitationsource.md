# ChatCitationSource

Four-variant citation source union.


## Supported Types

### ChatDocumentSource

```go
chatCitationSource := components.CreateChatCitationSourceDocument(components.ChatDocumentSource{/* values here */})
```

### ChatPersonSource

```go
chatCitationSource := components.CreateChatCitationSourcePerson(components.ChatPersonSource{/* values here */})
```

### ChatFileSource

```go
chatCitationSource := components.CreateChatCitationSourceFile(components.ChatFileSource{/* values here */})
```

### ChatCustomEntitySource

```go
chatCitationSource := components.CreateChatCitationSourceCustomEntity(components.ChatCustomEntitySource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch chatCitationSource.Type {
	case components.ChatCitationSourceTypeDocument:
		// chatCitationSource.ChatDocumentSource is populated
	case components.ChatCitationSourceTypePerson:
		// chatCitationSource.ChatPersonSource is populated
	case components.ChatCitationSourceTypeFile:
		// chatCitationSource.ChatFileSource is populated
	case components.ChatCitationSourceTypeCustomEntity:
		// chatCitationSource.ChatCustomEntitySource is populated
}
```
