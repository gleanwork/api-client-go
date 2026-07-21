package main

import (
	"context"
	"os"

	"github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
)

func main() {
	ctx := context.Background()
	client := apiclientgo.New()

	csvContent := []byte(`name,age,department
Alice,30,Engineering
Bob,25,Marketing`)

	uploadResp, _ := client.Client.Chat.UploadFiles(ctx, components.UploadChatFilesRequest{
		Files: []components.UploadChatFilesRequestFile{
			{FileName: "data.csv", Content: csvContent},
		},
	}, nil, nil)

	fileID := uploadResp.UploadChatFilesResponse.Files[0].ID

	client.Client.Agents.Run(ctx, components.AgentRunCreate{
		AgentID: os.Getenv("GLEAN_AGENT_ID"),
		Input: map[string]any{
			"file_id": *fileID,
		},
	}, nil)
}
