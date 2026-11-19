package hooks

import (
	"net/http"
	"os"
	"strings"
)

// Header names set by XGleanHook. Kept as consts so the header names live in
// a single place.
const (
	headerExcludeDeprecatedAfter = "X-Glean-Exclude-Deprecated-After"
	headerIncludeExperimental    = "X-Glean-Include-Experimental"
)

// Environment variable names read by XGleanHook.
const (
	envExcludeDeprecatedAfter = "X_GLEAN_EXCLUDE_DEPRECATED_AFTER"
	envIncludeExperimental    = "X_GLEAN_INCLUDE_EXPERIMENTAL"
	envIncludeBeta            = "X_GLEAN_INCLUDE_BETA"
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
		os.Getenv(envExcludeDeprecatedAfter),
		sdkConfig.ExcludeDeprecatedAfter,
	)

	// Get experimental value - environment variable takes precedence
	experimentalEnv := os.Getenv(envIncludeExperimental)
	experimentalValue := ""
	if experimentalEnv != "" {
		experimentalValue = experimentalEnv
	} else if sdkConfig.IncludeExperimental {
		experimentalValue = "true"
	}

	// IncludeExperimental is an alias of IncludeBeta. Each one resolves on its
	// own, and either one enables the header. Disabling one never disables the
	// other. Keep sending X-Glean-Include-Experimental for both: older
	// deployments do not allow X-Glean-Beta in their CORS preflight.
	if resolveBooleanOption(os.Getenv(envIncludeBeta), sdkConfig.IncludeBeta) {
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

// resolveBooleanOption resolves a boolean opt-in from its environment variable
// and its SDK option. The environment variable takes precedence: "true"
// (case-insensitive) enables it, any other non-empty value disables it, and an
// unset or blank value falls back to the SDK option.
func resolveBooleanOption(envValue string, optionValue bool) bool {
	if v := strings.TrimSpace(envValue); v != "" {
		return strings.EqualFold(v, "true")
	}
	return optionValue
}

// getFirstValue returns the first non-empty value from the provided arguments.
func getFirstValue(aValue, bValue string) string {
	if aValue != "" {
		return aValue
	}
	return bValue
}
