// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ChatMessageCitation - Information about the source for a ChatMessage.
type ChatMessageCitation struct {
	// An opaque token that represents this particular result in this particular ChatMessage. To be used for /feedback reporting.
	TrackingToken  *string   `json:"trackingToken,omitempty"`
	SourceDocument *Document `json:"sourceDocument,omitempty"`
	// Structure for file uploaded by a user for Chat.
	SourceFile   *ChatFile `json:"sourceFile,omitempty"`
	SourcePerson *Person   `json:"sourcePerson,omitempty"`
	// Each reference range and its corresponding snippets
	ReferenceRanges []ReferenceRange `json:"referenceRanges,omitempty"`
}

func (o *ChatMessageCitation) GetTrackingToken() *string {
	if o == nil {
		return nil
	}
	return o.TrackingToken
}

func (o *ChatMessageCitation) GetSourceDocument() *Document {
	if o == nil {
		return nil
	}
	return o.SourceDocument
}

func (o *ChatMessageCitation) GetSourceFile() *ChatFile {
	if o == nil {
		return nil
	}
	return o.SourceFile
}

func (o *ChatMessageCitation) GetSourcePerson() *Person {
	if o == nil {
		return nil
	}
	return o.SourcePerson
}

func (o *ChatMessageCitation) GetReferenceRanges() []ReferenceRange {
	if o == nil {
		return nil
	}
	return o.ReferenceRanges
}
