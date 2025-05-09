// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type PromptTemplateResult struct {
	PromptTemplate *PromptTemplate `json:"promptTemplate,omitempty"`
	// An opaque token that represents this prompt template
	TrackingToken *string       `json:"trackingToken,omitempty"`
	FavoriteInfo  *FavoriteInfo `json:"favoriteInfo,omitempty"`
	RunCount      *CountInfo    `json:"runCount,omitempty"`
}

func (o *PromptTemplateResult) GetPromptTemplate() *PromptTemplate {
	if o == nil {
		return nil
	}
	return o.PromptTemplate
}

func (o *PromptTemplateResult) GetTrackingToken() *string {
	if o == nil {
		return nil
	}
	return o.TrackingToken
}

func (o *PromptTemplateResult) GetFavoriteInfo() *FavoriteInfo {
	if o == nil {
		return nil
	}
	return o.FavoriteInfo
}

func (o *PromptTemplateResult) GetRunCount() *CountInfo {
	if o == nil {
		return nil
	}
	return o.RunCount
}
