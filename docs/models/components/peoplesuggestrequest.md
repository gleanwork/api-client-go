# PeopleSuggestRequest

Includes request params for type of suggestions.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Categories`                                                                                 | [][components.PeopleSuggestionCategory](../../models/components/peoplesuggestioncategory.md) | :heavy_check_mark:                                                                           | Categories of data requested. Request can include single or multiple categories.             |
| `Departments`                                                                                | []*string*                                                                                   | :heavy_minus_sign:                                                                           | Departments that the data is requested for. If empty, corresponds to whole company.          |