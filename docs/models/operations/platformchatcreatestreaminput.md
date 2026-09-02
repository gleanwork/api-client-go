# PlatformChatCreateStreamInput

Either a plain string (single user turn) or a chronological array of `USER`/`ASSISTANT` messages. The final array message must be `USER`.



## Supported Types

### 

```go
platformChatCreateStreamInput := operations.CreatePlatformChatCreateStreamInputStr(string{/* values here */})
```

### 

```go
platformChatCreateStreamInput := operations.CreatePlatformChatCreateStreamInputArrayOfPlatformChatInputMessage([]components.PlatformChatInputMessage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch platformChatCreateStreamInput.Type {
	case operations.PlatformChatCreateStreamInputTypeStr:
		// platformChatCreateStreamInput.Str is populated
	case operations.PlatformChatCreateStreamInputTypeArrayOfPlatformChatInputMessage:
		// platformChatCreateStreamInput.ArrayOfPlatformChatInputMessage is populated
}
```
