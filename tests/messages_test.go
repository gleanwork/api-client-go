// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	"context"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/internal/utils"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMessages_Messages(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("messages")

	s := apiclientgo.New(
		apiclientgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		apiclientgo.WithClient(testHTTPClient),
		apiclientgo.WithSecurity(utils.GetEnv("GLEAN_API_TOKEN", "value")),
	)

	res, err := s.Client.Messages.Retrieve(ctx, components.MessagesRequest{
		IDType: components.IDTypeConversationID,
		ID:     "<id>",
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)

}
