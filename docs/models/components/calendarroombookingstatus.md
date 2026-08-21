# CalendarRoomBookingStatus

The current booking status of the room resource associated with an event.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.CalendarRoomBookingStatusNone

// Open enum: custom values can be created with a direct type cast
custom := components.CalendarRoomBookingStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `CalendarRoomBookingStatusNone`     | NONE                                |
| `CalendarRoomBookingStatusAccepted` | ACCEPTED                            |
| `CalendarRoomBookingStatusDeclined` | DECLINED                            |
| `CalendarRoomBookingStatusPending`  | PENDING                             |