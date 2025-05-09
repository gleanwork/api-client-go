// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ContentDefinition - Describes text content or base64 encoded binary content
type ContentDefinition struct {
	MimeType string `json:"mimeType"`
	// text content. Only one of textContent or binary content can be specified
	TextContent *string `json:"textContent,omitempty"`
	// base64 encoded binary content. Only one of textContent or binary content can be specified
	BinaryContent *string `json:"binaryContent,omitempty"`
}

func (o *ContentDefinition) GetMimeType() string {
	if o == nil {
		return ""
	}
	return o.MimeType
}

func (o *ContentDefinition) GetTextContent() *string {
	if o == nil {
		return nil
	}
	return o.TextContent
}

func (o *ContentDefinition) GetBinaryContent() *string {
	if o == nil {
		return nil
	}
	return o.BinaryContent
}
