# PlatformUsageSettingsUpdateRequestUnion

Request body for updating workspace-wide usage settings.



## Supported Types

### PlatformUsageSettingsUpdateRequest1

```go
platformUsageSettingsUpdateRequestUnion := components.CreatePlatformUsageSettingsUpdateRequestUnionPlatformUsageSettingsUpdateRequest1(components.PlatformUsageSettingsUpdateRequest1{/* values here */})
```

### PlatformUsageSettingsUpdateRequest2

```go
platformUsageSettingsUpdateRequestUnion := components.CreatePlatformUsageSettingsUpdateRequestUnionPlatformUsageSettingsUpdateRequest2(components.PlatformUsageSettingsUpdateRequest2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch platformUsageSettingsUpdateRequestUnion.Type {
	case components.PlatformUsageSettingsUpdateRequestUnionTypePlatformUsageSettingsUpdateRequest1:
		// platformUsageSettingsUpdateRequestUnion.PlatformUsageSettingsUpdateRequest1 is populated
	case components.PlatformUsageSettingsUpdateRequestUnionTypePlatformUsageSettingsUpdateRequest2:
		// platformUsageSettingsUpdateRequestUnion.PlatformUsageSettingsUpdateRequest2 is populated
}
```
