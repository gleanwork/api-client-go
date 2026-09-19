# PlatformChatCreateStreamFormat

Output format for the assistant text. TEXT is unconstrained. JSON_SCHEMA constrains the response to the supplied schema.



## Supported Types

### PlatformChatTextFormat

```go
platformChatCreateStreamFormat := operations.CreatePlatformChatCreateStreamFormatPlatformChatTextFormat(components.PlatformChatTextFormat{/* values here */})
```

### PlatformChatJSONSchemaFormat

```go
platformChatCreateStreamFormat := operations.CreatePlatformChatCreateStreamFormatPlatformChatJSONSchemaFormat(components.PlatformChatJSONSchemaFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch platformChatCreateStreamFormat.Type {
	case operations.PlatformChatCreateStreamFormatTypePlatformChatTextFormat:
		// platformChatCreateStreamFormat.PlatformChatTextFormat is populated
	case operations.PlatformChatCreateStreamFormatTypePlatformChatJSONSchemaFormat:
		// platformChatCreateStreamFormat.PlatformChatJSONSchemaFormat is populated
}
```
