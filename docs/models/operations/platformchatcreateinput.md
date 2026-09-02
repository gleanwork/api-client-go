# PlatformChatCreateInput

Either a plain string (single user turn) or a chronological array of `USER`/`ASSISTANT` messages. The final array message must be `USER`.



## Supported Types

### 

```go
platformChatCreateInput := operations.CreatePlatformChatCreateInputStr(string{/* values here */})
```

### 

```go
platformChatCreateInput := operations.CreatePlatformChatCreateInputArrayOfPlatformChatInputMessage([]components.PlatformChatInputMessage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch platformChatCreateInput.Type {
	case operations.PlatformChatCreateInputTypeStr:
		// platformChatCreateInput.Str is populated
	case operations.PlatformChatCreateInputTypeArrayOfPlatformChatInputMessage:
		// platformChatCreateInput.ArrayOfPlatformChatInputMessage is populated
}
```
