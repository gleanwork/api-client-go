# AddCollectionItemsRequest


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `CollectionID`                                                                               | *float64*                                                                                    | :heavy_check_mark:                                                                           | The ID of the Collection to add items to.                                                    |
| `AddedCollectionItemDescriptors`                                                             | [][components.CollectionItemDescriptor](../../models/components/collectionitemdescriptor.md) | :heavy_minus_sign:                                                                           | The CollectionItemDescriptors of the items being added.                                      |