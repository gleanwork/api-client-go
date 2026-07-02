package hooks

import (
	"context"
	"net/http"
	"os"
	"testing"
)

// mockSDK is a simple struct to use as an SDK instance for testing.
type mockSDK struct{}

func createMockRequest() *http.Request {
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "https://example.com/api/test", nil)
	return req
}

func createMockContext(sdk any) BeforeRequestContext {
	return BeforeRequestContext{
		HookContext: HookContext{
			SDK:         sdk,
			BaseURL:     "https://example.com",
			OperationID: "test-operation",
		},
	}
}

func TestXGleanHook_NoConfiguration(t *testing.T) {
	// Clear environment variables
	os.Unsetenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER")
	os.Unsetenv("X_GLEAN_INCLUDE_EXPERIMENTAL")

	hook := &XGleanHook{}
	sdk := &mockSDK{}
	req := createMockRequest()
	ctx := createMockContext(sdk)

	result, err := hook.BeforeRequest(ctx, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Header.Get(headerExcludeDeprecatedAfter) != "" {
		t.Error("expected X-Glean-Exclude-Deprecated-After header to not be set")
	}

	if result.Header.Get(headerIncludeExperimental) != "" {
		t.Error("expected X-Glean-Include-Experimental header to not be set")
	}
}

func TestXGleanHook_SDKOptions_ExcludeDeprecatedAfter(t *testing.T) {
	// Clear environment variables
	os.Unsetenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER")
	os.Unsetenv("X_GLEAN_INCLUDE_EXPERIMENTAL")

	hook := &XGleanHook{}
	sdk := &mockSDK{}

	// Set SDK configuration
	SetXGleanConfig(sdk, XGleanConfig{
		ExcludeDeprecatedAfter: "2026-10-15",
	})

	req := createMockRequest()
	ctx := createMockContext(sdk)

	result, err := hook.BeforeRequest(ctx, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got := result.Header.Get(headerExcludeDeprecatedAfter); got != "2026-10-15" {
		t.Errorf("expected X-Glean-Exclude-Deprecated-After header to be '2026-10-15', got '%s'", got)
	}
}

func TestXGleanHook_SDKOptions_IncludeExperimental(t *testing.T) {
	// Clear environment variables
	os.Unsetenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER")
	os.Unsetenv("X_GLEAN_INCLUDE_EXPERIMENTAL")

	hook := &XGleanHook{}
	sdk := &mockSDK{}

	// Set SDK configuration
	SetXGleanConfig(sdk, XGleanConfig{
		IncludeExperimental: true,
	})

	req := createMockRequest()
	ctx := createMockContext(sdk)

	result, err := hook.BeforeRequest(ctx, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got := result.Header.Get(headerIncludeExperimental); got != "true" {
		t.Errorf("expected X-Glean-Include-Experimental header to be 'true', got '%s'", got)
	}
}

func TestXGleanHook_SDKOptions_IncludeExperimentalFalse(t *testing.T) {
	// Clear environment variables
	os.Unsetenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER")
	os.Unsetenv("X_GLEAN_INCLUDE_EXPERIMENTAL")

	hook := &XGleanHook{}
	sdk := &mockSDK{}

	// Set SDK configuration with IncludeExperimental = false
	SetXGleanConfig(sdk, XGleanConfig{
		IncludeExperimental: false,
	})

	req := createMockRequest()
	ctx := createMockContext(sdk)

	result, err := hook.BeforeRequest(ctx, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Header.Get(headerIncludeExperimental) != "" {
		t.Error("expected X-Glean-Include-Experimental header to not be set when IncludeExperimental is false")
	}
}

func TestXGleanHook_SDKOptions_BothOptions(t *testing.T) {
	// Clear environment variables
	os.Unsetenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER")
	os.Unsetenv("X_GLEAN_INCLUDE_EXPERIMENTAL")

	hook := &XGleanHook{}
	sdk := &mockSDK{}

	// Set SDK configuration
	SetXGleanConfig(sdk, XGleanConfig{
		ExcludeDeprecatedAfter: "2026-10-15",
		IncludeExperimental:    true,
	})

	req := createMockRequest()
	ctx := createMockContext(sdk)

	result, err := hook.BeforeRequest(ctx, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got := result.Header.Get(headerExcludeDeprecatedAfter); got != "2026-10-15" {
		t.Errorf("expected X-Glean-Exclude-Deprecated-After header to be '2026-10-15', got '%s'", got)
	}

	if got := result.Header.Get(headerIncludeExperimental); got != "true" {
		t.Errorf("expected X-Glean-Include-Experimental header to be 'true', got '%s'", got)
	}
}

func TestXGleanHook_EnvVars_ExcludeDeprecatedAfter(t *testing.T) {
	// Set environment variable
	os.Setenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER", "2027-01-01")
	defer os.Unsetenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER")
	os.Unsetenv("X_GLEAN_INCLUDE_EXPERIMENTAL")

	hook := &XGleanHook{}
	sdk := &mockSDK{}
	req := createMockRequest()
	ctx := createMockContext(sdk)

	result, err := hook.BeforeRequest(ctx, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got := result.Header.Get(headerExcludeDeprecatedAfter); got != "2027-01-01" {
		t.Errorf("expected X-Glean-Exclude-Deprecated-After header to be '2027-01-01', got '%s'", got)
	}
}

func TestXGleanHook_EnvVars_IncludeExperimental(t *testing.T) {
	// Set environment variable
	os.Unsetenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER")
	os.Setenv("X_GLEAN_INCLUDE_EXPERIMENTAL", "true")
	defer os.Unsetenv("X_GLEAN_INCLUDE_EXPERIMENTAL")

	hook := &XGleanHook{}
	sdk := &mockSDK{}
	req := createMockRequest()
	ctx := createMockContext(sdk)

	result, err := hook.BeforeRequest(ctx, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got := result.Header.Get(headerIncludeExperimental); got != "true" {
		t.Errorf("expected X-Glean-Include-Experimental header to be 'true', got '%s'", got)
	}
}

func TestXGleanHook_EnvVars_BothOptions(t *testing.T) {
	// Set environment variables
	os.Setenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER", "2027-06-15")
	os.Setenv("X_GLEAN_INCLUDE_EXPERIMENTAL", "true")
	defer os.Unsetenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER")
	defer os.Unsetenv("X_GLEAN_INCLUDE_EXPERIMENTAL")

	hook := &XGleanHook{}
	sdk := &mockSDK{}
	req := createMockRequest()
	ctx := createMockContext(sdk)

	result, err := hook.BeforeRequest(ctx, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got := result.Header.Get(headerExcludeDeprecatedAfter); got != "2027-06-15" {
		t.Errorf("expected X-Glean-Exclude-Deprecated-After header to be '2027-06-15', got '%s'", got)
	}

	if got := result.Header.Get(headerIncludeExperimental); got != "true" {
		t.Errorf("expected X-Glean-Include-Experimental header to be 'true', got '%s'", got)
	}
}

func TestXGleanHook_EnvVarsPrecedence_ExcludeDeprecatedAfter(t *testing.T) {
	// Set environment variable
	os.Setenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER", "2027-12-31")
	defer os.Unsetenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER")
	os.Unsetenv("X_GLEAN_INCLUDE_EXPERIMENTAL")

	hook := &XGleanHook{}
	sdk := &mockSDK{}

	// Set SDK configuration with different value
	SetXGleanConfig(sdk, XGleanConfig{
		ExcludeDeprecatedAfter: "2026-01-01",
	})

	req := createMockRequest()
	ctx := createMockContext(sdk)

	result, err := hook.BeforeRequest(ctx, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Environment variable should take precedence
	if got := result.Header.Get(headerExcludeDeprecatedAfter); got != "2027-12-31" {
		t.Errorf("expected X-Glean-Exclude-Deprecated-After header to be '2027-12-31' (from env var), got '%s'", got)
	}
}

func TestXGleanHook_EnvVarsPrecedence_IncludeExperimental(t *testing.T) {
	// Set environment variable
	os.Unsetenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER")
	os.Setenv("X_GLEAN_INCLUDE_EXPERIMENTAL", "true")
	defer os.Unsetenv("X_GLEAN_INCLUDE_EXPERIMENTAL")

	hook := &XGleanHook{}
	sdk := &mockSDK{}

	// Set SDK configuration with IncludeExperimental = false
	SetXGleanConfig(sdk, XGleanConfig{
		IncludeExperimental: false,
	})

	req := createMockRequest()
	ctx := createMockContext(sdk)

	result, err := hook.BeforeRequest(ctx, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Environment variable should take precedence
	if got := result.Header.Get(headerIncludeExperimental); got != "true" {
		t.Errorf("expected X-Glean-Include-Experimental header to be 'true' (from env var), got '%s'", got)
	}
}

func TestXGleanHook_EnvVarsPrecedence_BothOptions(t *testing.T) {
	// Set environment variables
	os.Setenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER", "2028-01-01")
	os.Setenv("X_GLEAN_INCLUDE_EXPERIMENTAL", "true")
	defer os.Unsetenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER")
	defer os.Unsetenv("X_GLEAN_INCLUDE_EXPERIMENTAL")

	hook := &XGleanHook{}
	sdk := &mockSDK{}

	// Set SDK configuration with different values
	SetXGleanConfig(sdk, XGleanConfig{
		ExcludeDeprecatedAfter: "2026-06-01",
		IncludeExperimental:    false,
	})

	req := createMockRequest()
	ctx := createMockContext(sdk)

	result, err := hook.BeforeRequest(ctx, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Environment variables should take precedence
	if got := result.Header.Get(headerExcludeDeprecatedAfter); got != "2028-01-01" {
		t.Errorf("expected X-Glean-Exclude-Deprecated-After header to be '2028-01-01' (from env var), got '%s'", got)
	}

	if got := result.Header.Get(headerIncludeExperimental); got != "true" {
		t.Errorf("expected X-Glean-Include-Experimental header to be 'true' (from env var), got '%s'", got)
	}
}
