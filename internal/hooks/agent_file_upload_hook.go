package hooks

import (
	"net/http"
	"strings"

	"github.com/gleanwork/api-client-go/internal/utils"
	"github.com/gleanwork/api-client-go/models/apierrors"
)

type AgentFileUploadHook struct{}

func (h *AgentFileUploadHook) AfterError(hookCtx AfterErrorContext, res *http.Response, err error) (*http.Response, error) {
	if res == nil {
		return res, err
	}

	operationID := hookCtx.OperationID
	if operationID != "createAndWaitRun" && operationID != "createAndStreamRun" {
		return res, err
	}

	if res.StatusCode != 400 {
		return res, err
	}

	bodyBytes, readErr := utils.ConsumeRawBody(res)
	if readErr != nil {
		return res, err
	}

	bodyStr := string(bodyBytes)
	bodyLower := strings.ToLower(bodyStr)

	if !strings.Contains(bodyLower, "permission") {
		return res, err
	}

	enhancedMessage := `Agent file upload error: Cannot pass file objects directly to agents.Run().

You must:
1. Upload the file first using client.Chat.UploadFiles()
2. Extract the file ID from the response
3. Pass the file ID string (not the file object) in the input map

Example:
  uploadResp, err := client.Chat.UploadFiles(ctx, components.UploadChatFilesRequest{
      Files: []components.UploadChatFilesRequestFile{
          {FileName: "data.csv", Content: fileBytes},
      },
  }, nil)
  if err != nil {
      return err
  }

  fileID := *uploadResp.UploadChatFilesResponse.Files[0].ID

  runResp, err := client.Agents.Run(ctx, components.AgentRunCreate{
      AgentID: "agent-123",
      Input: map[string]any{"file_id": fileID},
  })
`

	return res, apierrors.NewAPIError(enhancedMessage, res.StatusCode, bodyStr, res)
}
