<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := apiclientgo.New(
		apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
	)

	res, err := s.Client.Chat.Start(ctx, components.ChatRequest{
		Messages: []components.ChatMessage{
			components.ChatMessage{
				Fragments: []components.ChatMessageFragment{
					components.ChatMessageFragment{
						Text: apiclientgo.String("What are the company holidays this year?"),
					},
				},
			},
		},
	}, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ChatResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->