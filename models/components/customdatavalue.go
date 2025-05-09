// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CustomDataValue struct {
	DisplayLabel *string `json:"displayLabel,omitempty"`
	StringValue  *string `json:"stringValue,omitempty"`
	// list of strings for multi-value properties
	StringListValue []string `json:"stringListValue,omitempty"`
	NumberValue     *float64 `json:"numberValue,omitempty"`
	BooleanValue    *bool    `json:"booleanValue,omitempty"`
}

func (o *CustomDataValue) GetDisplayLabel() *string {
	if o == nil {
		return nil
	}
	return o.DisplayLabel
}

func (o *CustomDataValue) GetStringValue() *string {
	if o == nil {
		return nil
	}
	return o.StringValue
}

func (o *CustomDataValue) GetStringListValue() []string {
	if o == nil {
		return nil
	}
	return o.StringListValue
}

func (o *CustomDataValue) GetNumberValue() *float64 {
	if o == nil {
		return nil
	}
	return o.NumberValue
}

func (o *CustomDataValue) GetBooleanValue() *bool {
	if o == nil {
		return nil
	}
	return o.BooleanValue
}
