// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/gleanwork/api-client-go/models/components"
)

type EditcollectionitemResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	EditCollectionItemResponse *components.EditCollectionItemResponse
}

func (o *EditcollectionitemResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *EditcollectionitemResponse) GetEditCollectionItemResponse() *components.EditCollectionItemResponse {
	if o == nil {
		return nil
	}
	return o.EditCollectionItemResponse
}
