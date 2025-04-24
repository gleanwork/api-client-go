# AuthConfigType

The type of authentication being used.
Use 'OAUTH_*' when Glean calls an external API (e.g., Jira) on behalf of a user to obtain an OAuth token.
'OAUTH_ADMIN' utilizes an admin token for external API calls on behalf all users.
'OAUTH_USER' uses individual user tokens for external API calls.
'DWD' refers to domain wide delegation.



## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `AuthConfigTypeNone`       | NONE                       |
| `AuthConfigTypeOauthUser`  | OAUTH_USER                 |
| `AuthConfigTypeOauthAdmin` | OAUTH_ADMIN                |
| `AuthConfigTypeAPIKey`     | API_KEY                    |
| `AuthConfigTypeBasicAuth`  | BASIC_AUTH                 |
| `AuthConfigTypeDwd`        | DWD                        |