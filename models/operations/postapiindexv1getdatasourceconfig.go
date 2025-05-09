// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/gleanwork/api-client-go/models/components"
)

type PostAPIIndexV1GetdatasourceconfigResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	CustomDatasourceConfig *components.CustomDatasourceConfig
}

func (o *PostAPIIndexV1GetdatasourceconfigResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostAPIIndexV1GetdatasourceconfigResponse) GetCustomDatasourceConfig() *components.CustomDatasourceConfig {
	if o == nil {
		return nil
	}
	return o.CustomDatasourceConfig
}
