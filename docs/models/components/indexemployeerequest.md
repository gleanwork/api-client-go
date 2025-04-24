# IndexEmployeeRequest

Info about an employee and optional version for that info


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `Employee`                                                                             | [components.EmployeeInfoDefinition](../../models/components/employeeinfodefinition.md) | :heavy_check_mark:                                                                     | Describes employee info                                                                |
| `Version`                                                                              | **int64*                                                                               | :heavy_minus_sign:                                                                     | Version number for the employee object. If absent or 0 then no version checks are done |