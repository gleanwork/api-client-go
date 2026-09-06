# JustificationType

Type of the justification.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.JustificationTypeFrequentlyAccessed

// Open enum: custom values can be created with a direct type cast
custom := components.JustificationType("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `JustificationTypeFrequentlyAccessed`                | FREQUENTLY_ACCESSED                                  |
| `JustificationTypeRecentlyAccessed`                  | RECENTLY_ACCESSED                                    |
| `JustificationTypeTrendingDocument`                  | TRENDING_DOCUMENT                                    |
| `JustificationTypeVerificationReminder`              | VERIFICATION_REMINDER                                |
| `JustificationTypeSuggestedDocument`                 | SUGGESTED_DOCUMENT                                   |
| `JustificationTypeEmptyStateSuggestion`              | EMPTY_STATE_SUGGESTION                               |
| `JustificationTypeFrecencyScored`                    | FRECENCY_SCORED                                      |
| `JustificationTypeServerGenerated`                   | SERVER_GENERATED                                     |
| `JustificationTypeUseCase`                           | USE_CASE                                             |
| `JustificationTypeUpdateSinceLastView`               | UPDATE_SINCE_LAST_VIEW                               |
| `JustificationTypeRecentlyStarted`                   | RECENTLY_STARTED                                     |
| `JustificationTypeEvent`                             | EVENT                                                |
| `JustificationTypeUserMention`                       | USER_MENTION                                         |
| `JustificationTypeAnnouncement`                      | ANNOUNCEMENT                                         |
| `JustificationTypeExternalAnnouncement`              | EXTERNAL_ANNOUNCEMENT                                |
| `JustificationTypePopularityBasedTrending`           | POPULARITY_BASED_TRENDING                            |
| `JustificationTypeCompanyResource`                   | COMPANY_RESOURCE                                     |
| `JustificationTypeEventDocumentFromContent`          | EVENT_DOCUMENT_FROM_CONTENT                          |
| `JustificationTypeEventDocumentFromSearch`           | EVENT_DOCUMENT_FROM_SEARCH                           |
| `JustificationTypeVisitAffinityScored`               | VISIT_AFFINITY_SCORED                                |
| `JustificationTypeSuggestedApp`                      | SUGGESTED_APP                                        |
| `JustificationTypeSuggestedPerson`                   | SUGGESTED_PERSON                                     |
| `JustificationTypeActivityHighlight`                 | ACTIVITY_HIGHLIGHT                                   |
| `JustificationTypeSavedSearch`                       | SAVED_SEARCH                                         |
| `JustificationTypeSuggestedChannel`                  | SUGGESTED_CHANNEL                                    |
| `JustificationTypePeopleCelebrations`                | PEOPLE_CELEBRATIONS                                  |
| `JustificationTypeSocialLink`                        | SOCIAL_LINK                                          |
| `JustificationTypeZeroStateChatSuggestion`           | ZERO_STATE_CHAT_SUGGESTION                           |
| `JustificationTypeZeroStateChatToolSuggestion`       | ZERO_STATE_CHAT_TOOL_SUGGESTION                      |
| `JustificationTypeZeroStatePromptTemplateSuggestion` | ZERO_STATE_PROMPT_TEMPLATE_SUGGESTION                |
| `JustificationTypeZeroStateStaticWorkflowSuggestion` | ZERO_STATE_STATIC_WORKFLOW_SUGGESTION                |
| `JustificationTypeZeroStateAgentSuggestion`          | ZERO_STATE_AGENT_SUGGESTION                          |
| `JustificationTypePersonalizedChatSuggestion`        | PERSONALIZED_CHAT_SUGGESTION                         |
| `JustificationTypeDailyDigest`                       | DAILY_DIGEST                                         |
| `JustificationTypePodcast`                           | PODCAST                                              |
| `JustificationTypeTask`                              | TASK                                                 |
| `JustificationTypePlanMyDay`                         | PLAN_MY_DAY                                          |
| `JustificationTypeEndMyDay`                          | END_MY_DAY                                           |
| `JustificationTypeStarterKitExtension`               | STARTER_KIT_EXTENSION                                |
| `JustificationTypeStarterKitOrgChart`                | STARTER_KIT_ORG_CHART                                |
| `JustificationTypeStarterKitAddDoc`                  | STARTER_KIT_ADD_DOC                                  |
| `JustificationTypeMeetingRecap`                      | MEETING_RECAP                                        |
| `JustificationTypeMeetingRoom`                       | MEETING_ROOM                                         |
| `JustificationTypeMeetingPrepAutomation`             | MEETING_PREP_AUTOMATION                              |
| `JustificationTypeActiveDiscussion`                  | ACTIVE_DISCUSSION                                    |
| `JustificationTypeMidDayCatchUp`                     | MID_DAY_CATCH_UP                                     |
| `JustificationTypeQuerySuggestion`                   | QUERY_SUGGESTION                                     |
| `JustificationTypeCoworkCujPromo`                    | COWORK_CUJ_PROMO                                     |
| `JustificationTypeCardStackPromo`                    | CARD_STACK_PROMO                                     |
| `JustificationTypeWeeklyMeetings`                    | WEEKLY_MEETINGS                                      |
| `JustificationTypeFollowUp`                          | FOLLOW_UP                                            |
| `JustificationTypeMilestoneTimelineCheck`            | MILESTONE_TIMELINE_CHECK                             |
| `JustificationTypeProjectDiscussionDigest`           | PROJECT_DISCUSSION_DIGEST                            |
| `JustificationTypeProjectFocusBlock`                 | PROJECT_FOCUS_BLOCK                                  |
| `JustificationTypeProjectNextStep`                   | PROJECT_NEXT_STEP                                    |
| `JustificationTypeDemoCard`                          | DEMO_CARD                                            |
| `JustificationTypeResolveMeetingConflict`            | RESOLVE_MEETING_CONFLICT                             |
| `JustificationTypeFindingMeetingRoom`                | FINDING_MEETING_ROOM                                 |
| `JustificationTypeOooPlanner`                        | OOO_PLANNER                                          |
| `JustificationTypeOooCatchUp`                        | OOO_CATCH_UP                                         |
| `JustificationTypeAdminHealthCenter`                 | ADMIN_HEALTH_CENTER                                  |
| `JustificationTypeKnowledgeGap`                      | KNOWLEDGE_GAP                                        |
| `JustificationTypeArtifactCreation`                  | ARTIFACT_CREATION                                    |