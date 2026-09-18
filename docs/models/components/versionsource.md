# VersionSource

Provenance recorded on the staged commit or published version this import produces. Doesn't change the agent's management mode (workflowSource). GIT: synced from a Git repository. USER: uploaded by a user. Defaults to USER when omitted. Ignored for transient imports.


## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.VersionSourceGit
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `VersionSourceGit`  | GIT                 |
| `VersionSourceUser` | USER                |