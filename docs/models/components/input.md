# Input

Either a plain string (single user turn) or a chronological array of `USER`/`ASSISTANT` messages. The final array message must be `USER`.



## Supported Types

### 

```go
input := components.CreateInputStr(string{/* values here */})
```

### 

```go
input := components.CreateInputArrayOfPlatformChatInputMessage([]components.PlatformChatInputMessage{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch input.Type {
	case components.InputTypeStr:
		// input.Str is populated
	case components.InputTypeArrayOfPlatformChatInputMessage:
		// input.ArrayOfPlatformChatInputMessage is populated
}
```
