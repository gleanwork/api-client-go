# DatasourceCategory

The type of this datasource. It is an important signal for relevance and must be specified and cannot be UNCATEGORIZED. Please refer to [this](https://developers.glean.com/docs/indexing_api_datasource_category/) for more details.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.DatasourceCategoryUncategorized

// Open enum: custom values can be created with a direct type cast
custom := components.DatasourceCategory("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `DatasourceCategoryUncategorized`        | UNCATEGORIZED                            |
| `DatasourceCategoryTickets`              | TICKETS                                  |
| `DatasourceCategoryCrm`                  | CRM                                      |
| `DatasourceCategoryPublishedContent`     | PUBLISHED_CONTENT                        |
| `DatasourceCategoryCollaborativeContent` | COLLABORATIVE_CONTENT                    |
| `DatasourceCategoryQuestionAnswer`       | QUESTION_ANSWER                          |
| `DatasourceCategoryMessaging`            | MESSAGING                                |
| `DatasourceCategoryCodeRepository`       | CODE_REPOSITORY                          |
| `DatasourceCategoryChangeManagement`     | CHANGE_MANAGEMENT                        |
| `DatasourceCategoryPeople`               | PEOPLE                                   |
| `DatasourceCategoryEmail`                | EMAIL                                    |
| `DatasourceCategorySso`                  | SSO                                      |
| `DatasourceCategoryAts`                  | ATS                                      |
| `DatasourceCategoryKnowledgeHub`         | KNOWLEDGE_HUB                            |
| `DatasourceCategoryExternalShortcut`     | EXTERNAL_SHORTCUT                        |
| `DatasourceCategoryEntity`               | ENTITY                                   |
| `DatasourceCategoryCalendar`             | CALENDAR                                 |
| `DatasourceCategoryAgents`               | AGENTS                                   |
| `DatasourceCategoryAiConversation`       | AI_CONVERSATION                          |
| `DatasourceCategoryAiArtifact`           | AI_ARTIFACT                              |