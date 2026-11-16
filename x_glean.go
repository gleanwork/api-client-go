package apiclientgo

import (
	"github.com/gleanwork/api-client-go/internal/hooks"
)

// xGleanOptions holds the X-Glean configuration options that will be applied after SDK initialization.
type xGleanOptions struct {
	excludeDeprecatedAfter string
	includeBeta            bool
	includeExperimental    bool
}

// pendingXGleanOptions stores options that need to be applied after SDK initialization.
var pendingXGleanOptions = make(map[*Glean]*xGleanOptions)

// WithExcludeDeprecatedAfter configures the SDK to exclude API endpoints that will be
// deprecated after the specified date. This is useful for testing your integration
// against upcoming deprecations.
//
// Format: YYYY-MM-DD (e.g., "2026-10-15")
//
// This can also be set via the X_GLEAN_EXCLUDE_DEPRECATED_AFTER environment variable.
// Environment variables take precedence over SDK options.
//
// More information: https://developers.glean.com/deprecations/overview
func WithExcludeDeprecatedAfter(date string) SDKOption {
	return func(sdk *Glean) {
		opts := getOrCreatePendingOptions(sdk)
		opts.excludeDeprecatedAfter = date
		applyXGleanConfig(sdk, opts)
	}
}

// WithIncludeBeta configures the SDK to enable beta API features that are not
// yet generally available. Use this to preview and test new functionality.
//
// Warning: Beta features may change or be removed without notice.
// Do not rely on beta features in production environments.
//
// This can also be set via the X_GLEAN_INCLUDE_BETA environment variable.
// Environment variables take precedence over SDK options.
//
// WithIncludeExperimental is an alias. If either one is enabled, the SDK
// enables beta API features.
func WithIncludeBeta(enabled bool) SDKOption {
	return func(sdk *Glean) {
		opts := getOrCreatePendingOptions(sdk)
		opts.includeBeta = enabled
		applyXGleanConfig(sdk, opts)
	}
}

// WithIncludeExperimental is an alias of WithIncludeBeta. If either one is
// enabled, the SDK enables beta API features.
//
// This can also be set via the X_GLEAN_INCLUDE_EXPERIMENTAL environment variable.
// Environment variables take precedence over SDK options.
func WithIncludeExperimental(enabled bool) SDKOption {
	return func(sdk *Glean) {
		opts := getOrCreatePendingOptions(sdk)
		opts.includeExperimental = enabled
		applyXGleanConfig(sdk, opts)
	}
}

// getOrCreatePendingOptions retrieves or creates the pending options for an SDK instance.
func getOrCreatePendingOptions(sdk *Glean) *xGleanOptions {
	if opts, ok := pendingXGleanOptions[sdk]; ok {
		return opts
	}
	opts := &xGleanOptions{}
	pendingXGleanOptions[sdk] = opts
	return opts
}

// applyXGleanConfig applies the X-Glean configuration to the hooks package.
func applyXGleanConfig(sdk *Glean, opts *xGleanOptions) {
	hooks.SetXGleanConfig(sdk, hooks.XGleanConfig{
		ExcludeDeprecatedAfter: opts.excludeDeprecatedAfter,
		IncludeBeta:            opts.includeBeta,
		IncludeExperimental:    opts.includeExperimental,
	})
}
