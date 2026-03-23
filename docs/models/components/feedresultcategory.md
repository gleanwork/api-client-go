# FeedResultCategory

Category of the result, one of the requested categories in incoming request.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.FeedResultCategoryDocumentSuggestion

// Open enum: custom values can be created with a direct type cast
custom := components.FeedResultCategory("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `FeedResultCategoryDocumentSuggestion`           | DOCUMENT_SUGGESTION                              |
| `FeedResultCategoryDocumentSuggestionScenario`   | DOCUMENT_SUGGESTION_SCENARIO                     |
| `FeedResultCategoryTrendingDocument`             | TRENDING_DOCUMENT                                |
| `FeedResultCategoryUseCase`                      | USE_CASE                                         |
| `FeedResultCategoryVerificationReminder`         | VERIFICATION_REMINDER                            |
| `FeedResultCategoryEvent`                        | EVENT                                            |
| `FeedResultCategoryAnnouncement`                 | ANNOUNCEMENT                                     |
| `FeedResultCategoryMention`                      | MENTION                                          |
| `FeedResultCategoryDatasourceAffinity`           | DATASOURCE_AFFINITY                              |
| `FeedResultCategoryRecent`                       | RECENT                                           |
| `FeedResultCategoryCompanyResource`              | COMPANY_RESOURCE                                 |
| `FeedResultCategoryExperimental`                 | EXPERIMENTAL                                     |
| `FeedResultCategoryPeopleCelebrations`           | PEOPLE_CELEBRATIONS                              |
| `FeedResultCategorySocialLink`                   | SOCIAL_LINK                                      |
| `FeedResultCategoryExternalTasks`                | EXTERNAL_TASKS                                   |
| `FeedResultCategoryDisplayableList`              | DISPLAYABLE_LIST                                 |
| `FeedResultCategoryZeroStateChatSuggestion`      | ZERO_STATE_CHAT_SUGGESTION                       |
| `FeedResultCategoryZeroStateChatToolSuggestion`  | ZERO_STATE_CHAT_TOOL_SUGGESTION                  |
| `FeedResultCategoryZeroStateWorkflowCreatedByMe` | ZERO_STATE_WORKFLOW_CREATED_BY_ME                |
| `FeedResultCategoryZeroStateWorkflowFavorites`   | ZERO_STATE_WORKFLOW_FAVORITES                    |
| `FeedResultCategoryZeroStateWorkflowPopular`     | ZERO_STATE_WORKFLOW_POPULAR                      |
| `FeedResultCategoryZeroStateWorkflowRecent`      | ZERO_STATE_WORKFLOW_RECENT                       |
| `FeedResultCategoryZeroStateWorkflowSuggestion`  | ZERO_STATE_WORKFLOW_SUGGESTION                   |
| `FeedResultCategoryPersonalizedChatSuggestion`   | PERSONALIZED_CHAT_SUGGESTION                     |
| `FeedResultCategoryDailyDigest`                  | DAILY_DIGEST                                     |
| `FeedResultCategoryTask`                         | TASK                                             |
| `FeedResultCategoryPlanMyDay`                    | PLAN_MY_DAY                                      |
| `FeedResultCategoryEndMyDay`                     | END_MY_DAY                                       |
| `FeedResultCategoryStarterKit`                   | STARTER_KIT                                      |
| `FeedResultCategoryMidDayCatchUp`                | MID_DAY_CATCH_UP                                 |
| `FeedResultCategoryQuerySuggestion`              | QUERY_SUGGESTION                                 |
| `FeedResultCategoryWeeklyMeetings`               | WEEKLY_MEETINGS                                  |