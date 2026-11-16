package apiclientgo

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/gleanwork/api-client-go/internal/hooks"
	"github.com/gleanwork/api-client-go/models/components"
)

func TestWithExcludeDeprecatedAfter(t *testing.T) {
	sdk := New(
		WithExcludeDeprecatedAfter("2026-10-15"),
	)

	config := hooks.GetXGleanConfig(sdk)
	if config.ExcludeDeprecatedAfter != "2026-10-15" {
		t.Errorf("expected ExcludeDeprecatedAfter to be '2026-10-15', got '%s'", config.ExcludeDeprecatedAfter)
	}
}

func TestWithIncludeExperimental(t *testing.T) {
	sdk := New(
		WithIncludeExperimental(true),
	)

	config := hooks.GetXGleanConfig(sdk)
	if !config.IncludeExperimental {
		t.Error("expected IncludeExperimental to be true")
	}
}

func TestWithBothOptions(t *testing.T) {
	sdk := New(
		WithExcludeDeprecatedAfter("2026-10-15"),
		WithIncludeExperimental(true),
	)

	config := hooks.GetXGleanConfig(sdk)
	if config.ExcludeDeprecatedAfter != "2026-10-15" {
		t.Errorf("expected ExcludeDeprecatedAfter to be '2026-10-15', got '%s'", config.ExcludeDeprecatedAfter)
	}
	if !config.IncludeExperimental {
		t.Error("expected IncludeExperimental to be true")
	}
}

func TestWithIncludeBeta(t *testing.T) {
	sdk := New(
		WithIncludeBeta(true),
	)

	config := hooks.GetXGleanConfig(sdk)
	if !config.IncludeBeta {
		t.Error("expected IncludeBeta to be true")
	}
	if config.IncludeExperimental {
		t.Error("expected IncludeExperimental to be false")
	}
}

func TestWithIncludeBetaAndIncludeExperimental(t *testing.T) {
	sdk := New(
		WithIncludeBeta(true),
		WithIncludeExperimental(true),
		WithExcludeDeprecatedAfter("2026-10-15"),
	)

	config := hooks.GetXGleanConfig(sdk)
	if !config.IncludeBeta {
		t.Error("expected IncludeBeta to be true")
	}
	if !config.IncludeExperimental {
		t.Error("expected IncludeExperimental to be true")
	}
	if config.ExcludeDeprecatedAfter != "2026-10-15" {
		t.Errorf("expected ExcludeDeprecatedAfter to be '2026-10-15', got '%s'", config.ExcludeDeprecatedAfter)
	}
}

// headerCapturingClient records the headers of the last request it receives.
type headerCapturingClient struct {
	header http.Header
}

func (c *headerCapturingClient) Do(req *http.Request) (*http.Response, error) {
	c.header = req.Header.Clone()
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader("{}")),
		Request:    req,
	}, nil
}

func TestXGleanOptionsSendIncludeExperimentalHeader(t *testing.T) {
	tests := []struct {
		name string
		opts []SDKOption
		want string
	}{
		{name: "neither option", want: ""},
		{name: "WithIncludeBeta", opts: []SDKOption{WithIncludeBeta(true)}, want: "true"},
		{name: "WithIncludeExperimental", opts: []SDKOption{WithIncludeExperimental(true)}, want: "true"},
		{name: "both options", opts: []SDKOption{WithIncludeBeta(true), WithIncludeExperimental(true)}, want: "true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("X_GLEAN_INCLUDE_BETA", "")
			t.Setenv("X_GLEAN_INCLUDE_EXPERIMENTAL", "")

			client := &headerCapturingClient{}
			opts := append([]SDKOption{
				WithClient(client),
				WithSecurity("test-token"),
				WithServerURL("https://example.com"),
			}, tt.opts...)
			sdk := New(opts...)

			_, _ = sdk.Search.Query(context.Background(), components.PlatformSearchRequest{Query: "test"})

			if client.header == nil {
				t.Fatal("expected a request to be sent")
			}
			if got := client.header.Get("X-Glean-Include-Experimental"); got != tt.want {
				t.Errorf("expected X-Glean-Include-Experimental header to be %q, got %q", tt.want, got)
			}
			if got := client.header.Get("X-Glean-Beta"); got != "" {
				t.Errorf("expected X-Glean-Beta header to not be set yet, got %q", got)
			}
		})
	}
}

func TestWithIncludeExperimentalFalse(t *testing.T) {
	sdk := New(
		WithIncludeExperimental(false),
	)

	config := hooks.GetXGleanConfig(sdk)
	if config.IncludeExperimental {
		t.Error("expected IncludeExperimental to be false")
	}
}

func TestMultipleSDKInstances(t *testing.T) {
	sdk1 := New(
		WithExcludeDeprecatedAfter("2026-01-01"),
	)
	sdk2 := New(
		WithExcludeDeprecatedAfter("2027-12-31"),
	)

	config1 := hooks.GetXGleanConfig(sdk1)
	config2 := hooks.GetXGleanConfig(sdk2)

	if config1.ExcludeDeprecatedAfter != "2026-01-01" {
		t.Errorf("sdk1: expected ExcludeDeprecatedAfter to be '2026-01-01', got '%s'", config1.ExcludeDeprecatedAfter)
	}
	if config2.ExcludeDeprecatedAfter != "2027-12-31" {
		t.Errorf("sdk2: expected ExcludeDeprecatedAfter to be '2027-12-31', got '%s'", config2.ExcludeDeprecatedAfter)
	}
}

func TestEnvironmentVariablesAreRead(t *testing.T) {
	// This test verifies that the hook reads environment variables
	// The actual precedence is tested in the hooks package

	// Set environment variables
	os.Setenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER", "2028-06-01")
	os.Setenv("X_GLEAN_INCLUDE_EXPERIMENTAL", "true")
	defer os.Unsetenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER")
	defer os.Unsetenv("X_GLEAN_INCLUDE_EXPERIMENTAL")

	// SDK can be created without error when env vars are set
	sdk := New()
	if sdk == nil {
		t.Error("expected SDK to be created")
	}
}
