# PlatformListDepartmentsResponse

One page of departments.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `Results`                                                                         | [][components.PlatformDepartment](../../models/components/platformdepartment.md)  | :heavy_check_mark:                                                                | The departments on this page, ordered by display_name and then by department_id.<br/> |
| `HasMore`                                                                         | `bool`                                                                            | :heavy_check_mark:                                                                | Whether more departments are available after this page.                           |
| `NextCursor`                                                                      | `*string`                                                                         | :heavy_minus_sign:                                                                | Opaque cursor for the next page; null or absent when has_more is false.           |
| `RequestID`                                                                       | `string`                                                                          | :heavy_check_mark:                                                                | Request identifier for correlating this response.                                 |