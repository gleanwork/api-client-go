# FeedbackCategory

The feature category to which the feedback applies. These should be broad product areas such as Announcements, Answers, Search, etc. rather than specific components or UI treatments within those areas.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.FeedbackCategoryAnnouncement
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `FeedbackCategoryAnnouncement` | ANNOUNCEMENT                   |
| `FeedbackCategoryAnswers`      | ANSWERS                        |
| `FeedbackCategoryArtifacts`    | ARTIFACTS                      |
| `FeedbackCategoryAutocomplete` | AUTOCOMPLETE                   |
| `FeedbackCategoryCollections`  | COLLECTIONS                    |
| `FeedbackCategoryFeed`         | FEED                           |
| `FeedbackCategorySearch`       | SEARCH                         |
| `FeedbackCategoryChat`         | CHAT                           |
| `FeedbackCategoryNtp`          | NTP                            |
| `FeedbackCategoryWorkflows`    | WORKFLOWS                      |
| `FeedbackCategorySummary`      | SUMMARY                        |
| `FeedbackCategoryGeneral`      | GENERAL                        |
| `FeedbackCategoryPrism`        | PRISM                          |
| `FeedbackCategoryPrompts`      | PROMPTS                        |