# PlatformDocumentsBatchRequestDocumentSpec


## Supported Types

### PlatformDocumentsBatchSpecByID

```go
platformDocumentsBatchRequestDocumentSpec := components.CreatePlatformDocumentsBatchRequestDocumentSpecPlatformDocumentsBatchSpecByID(components.PlatformDocumentsBatchSpecByID{/* values here */})
```

### PlatformDocumentsBatchSpecByURL

```go
platformDocumentsBatchRequestDocumentSpec := components.CreatePlatformDocumentsBatchRequestDocumentSpecPlatformDocumentsBatchSpecByURL(components.PlatformDocumentsBatchSpecByURL{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch platformDocumentsBatchRequestDocumentSpec.Type {
	case components.PlatformDocumentsBatchRequestDocumentSpecTypePlatformDocumentsBatchSpecByID:
		// platformDocumentsBatchRequestDocumentSpec.PlatformDocumentsBatchSpecByID is populated
	case components.PlatformDocumentsBatchRequestDocumentSpecTypePlatformDocumentsBatchSpecByURL:
		// platformDocumentsBatchRequestDocumentSpec.PlatformDocumentsBatchSpecByURL is populated
}
```
