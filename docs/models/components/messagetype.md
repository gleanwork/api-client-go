# MessageType

Semantically groups content of a certain type. It can be used for purposes such as differential UI treatment. USER authored messages should be of type CONTENT and do not need `messageType` specified.


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `MessageTypeUpdate`        | UPDATE                     |
| `MessageTypeContent`       | CONTENT                    |
| `MessageTypeContext`       | CONTEXT                    |
| `MessageTypeDebug`         | DEBUG                      |
| `MessageTypeDebugExternal` | DEBUG_EXTERNAL             |
| `MessageTypeError`         | ERROR                      |
| `MessageTypeHeading`       | HEADING                    |
| `MessageTypeWarning`       | WARNING                    |