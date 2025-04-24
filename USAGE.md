<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/api-client-go/types"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := apiclientgo.New(
		apiclientgo.WithSecurity(os.Getenv("GLEAN_BEARER_AUTH")),
	)

	res, err := s.Client.Activity.Report(ctx, components.Activity{
		Events: []components.ActivityEvent{
			components.ActivityEvent{
				Action:    components.ActivityEventActionHistoricalView,
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionSearch,
				Params: &components.ActivityEventParams{
					Query: apiclientgo.String("query"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/search?q=query",
			},
			components.ActivityEvent{
				Action: components.ActivityEventActionView,
				Params: &components.ActivityEventParams{
					Duration: apiclientgo.Int64(20),
					Referrer: apiclientgo.String("https://example.com/document"),
				},
				Timestamp: types.MustTimeFromString("2000-01-23T04:56:07.000Z"),
				URL:       "https://example.com/",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->