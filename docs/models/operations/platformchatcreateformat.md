# PlatformChatCreateFormat

Output format for the assistant text. TEXT is unconstrained. JSON_SCHEMA constrains the response to the supplied schema.



## Supported Types

### PlatformChatTextFormat

```go
platformChatCreateFormat := operations.CreatePlatformChatCreateFormatPlatformChatTextFormat(components.PlatformChatTextFormat{/* values here */})
```

### PlatformChatJSONSchemaFormat

```go
platformChatCreateFormat := operations.CreatePlatformChatCreateFormatPlatformChatJSONSchemaFormat(components.PlatformChatJSONSchemaFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch platformChatCreateFormat.Type {
	case operations.PlatformChatCreateFormatTypePlatformChatTextFormat:
		// platformChatCreateFormat.PlatformChatTextFormat is populated
	case operations.PlatformChatCreateFormatTypePlatformChatJSONSchemaFormat:
		// platformChatCreateFormat.PlatformChatJSONSchemaFormat is populated
}
```
