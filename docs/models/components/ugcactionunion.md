# UgcActionUnion

An action to perform on user-generated content. This may be accompanied by `text` on the ChatMessageFragment, which acts as the display name content of the pill.


## Supported Types

### UgcAction

```go
ugcActionUnion := components.CreateUgcActionUnionUgcAction(components.UgcAction{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ugcActionUnion.Type {
	case components.UgcActionUnionTypeUgcAction:
		// ugcActionUnion.UgcAction is populated
}
```
