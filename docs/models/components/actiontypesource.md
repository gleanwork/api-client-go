# ActionTypeSource

Analytics-only signal (product snapshot) describing WHERE the action's
read/write determination came from. Complementary to the effective
read/write value (the tool's ToolType, which drives HITL): the value says
read-or-write, this says how confident that is. MCP_ANNOTATION = from the
tool's read-only/destructive hints; ADMIN_OVERRIDE = an admin set it;
NONE = no usable hint (the effective value then defaults to write);
NATIVE_TOOL_DEFINITION = from a curated native tool (snapshot-derived).
Does not affect runtime behavior.


## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ActionTypeSourceMcpAnnotation

// Open enum: custom values can be created with a direct type cast
custom := components.ActionTypeSource("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `ActionTypeSourceMcpAnnotation`        | MCP_ANNOTATION                         |
| `ActionTypeSourceAdminOverride`        | ADMIN_OVERRIDE                         |
| `ActionTypeSourceNone`                 | NONE                                   |
| `ActionTypeSourceNativeToolDefinition` | NATIVE_TOOL_DEFINITION                 |