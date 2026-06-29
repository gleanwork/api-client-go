# CreateWorkflowRequest


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `Name`                                            | `*string`                                         | :heavy_minus_sign:                                | The name of the workflow.                         |
| `Transient`                                       | `*bool`                                           | :heavy_minus_sign:                                | Used to create a transient workflow.              |
| `ParentWorkflowID`                                | `*string`                                         | :heavy_minus_sign:                                | id of the parent workflow for transient workflows |