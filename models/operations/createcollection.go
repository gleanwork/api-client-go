// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/gleanwork/api-client-go/models/components"
)

type CreatecollectionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	CreateCollectionResponse *components.CreateCollectionResponse
}

func (o *CreatecollectionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreatecollectionResponse) GetCreateCollectionResponse() *components.CreateCollectionResponse {
	if o == nil {
		return nil
	}
	return o.CreateCollectionResponse
}
