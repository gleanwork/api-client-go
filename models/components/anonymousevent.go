// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// AnonymousEventEventType - The nature of the event, for example "out of office".
type AnonymousEventEventType string

const (
	AnonymousEventEventTypeDefault     AnonymousEventEventType = "DEFAULT"
	AnonymousEventEventTypeOutOfOffice AnonymousEventEventType = "OUT_OF_OFFICE"
)

func (e AnonymousEventEventType) ToPointer() *AnonymousEventEventType {
	return &e
}
func (e *AnonymousEventEventType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DEFAULT":
		fallthrough
	case "OUT_OF_OFFICE":
		*e = AnonymousEventEventType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AnonymousEventEventType: %v", v)
	}
}

// AnonymousEvent - A generic, light-weight calendar event.
type AnonymousEvent struct {
	Time *TimeInterval `json:"time,omitempty"`
	// The nature of the event, for example "out of office".
	EventType *AnonymousEventEventType `json:"eventType,omitempty"`
}

func (o *AnonymousEvent) GetTime() *TimeInterval {
	if o == nil {
		return nil
	}
	return o.Time
}

func (o *AnonymousEvent) GetEventType() *AnonymousEventEventType {
	if o == nil {
		return nil
	}
	return o.EventType
}
