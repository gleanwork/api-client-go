package hooks

import (
	"net/http"
	"os"
)

// Header names set by XGleanHook. Kept as consts so the header names live in
// a single place.
const (
	headerExcludeDeprecatedAfter = "X-Glean-Exclude-Deprecated-After"
	headerIncludeExperimental    = "X-Glean-Include-Experimental"
)

// XGleanHook is a beforeRequest hook that sets X-Glean headers for
// experimental features and deprecation testing.
type XGleanHook struct{}

// BeforeRequest sets the X-Glean-Exclude-Deprecated-After and X-Glean-Include-Experimental
// headers based on environment variables or SDK configuration.
// Environment variables take precedence over SDK options.
func (h *XGleanHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	// Get SDK configuration
	sdkConfig := GetXGleanConfig(hookCtx.SDK)

	// Get deprecatedAfter value - environment variable takes precedence
	deprecatedAfter := getFirstValue(
		os.Getenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER"),
		sdkConfig.ExcludeDeprecatedAfter,
	)

	// Get experimental value - environment variable takes precedence
	experimentalEnv := os.Getenv("X_GLEAN_INCLUDE_EXPERIMENTAL")
	experimentalValue := ""
	if experimentalEnv != "" {
		experimentalValue = experimentalEnv
	} else if sdkConfig.IncludeExperimental {
		experimentalValue = "true"
	}

	// Set headers if values are present
	if deprecatedAfter != "" {
		req.Header.Set(headerExcludeDeprecatedAfter, deprecatedAfter)
	}

	if experimentalValue != "" {
		req.Header.Set(headerIncludeExperimental, experimentalValue)
	}

	return req, nil
}

// getFirstValue returns the first non-empty value from the provided arguments.
func getFirstValue(aValue, bValue string) string {
	if aValue != "" {
		return aValue
	}
	return bValue
}
