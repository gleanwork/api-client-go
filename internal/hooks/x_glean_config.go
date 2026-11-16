package hooks

import "sync"

// XGleanConfig holds configuration for X-Glean headers.
// This is used to pass configuration from SDK options to the hook.
type XGleanConfig struct {
	// ExcludeDeprecatedAfter excludes API endpoints that will be deprecated after this date.
	// Format: YYYY-MM-DD (e.g., "2026-10-15")
	ExcludeDeprecatedAfter string

	// IncludeBeta enables beta API features when true.
	IncludeBeta bool

	// IncludeExperimental is an alias of IncludeBeta. Either one enables beta
	// API features.
	IncludeExperimental bool
}

// xGleanConfigs maps SDK instances to their X-Glean configuration.
// Uses a sync.Map for thread-safe access across multiple SDK instances.
var xGleanConfigs sync.Map

// SetXGleanConfig stores the X-Glean configuration for a given SDK instance.
func SetXGleanConfig(sdk any, config XGleanConfig) {
	xGleanConfigs.Store(sdk, config)
}

// GetXGleanConfig retrieves the X-Glean configuration for a given SDK instance.
// Returns an empty config if none is set.
func GetXGleanConfig(sdk any) XGleanConfig {
	if config, ok := xGleanConfigs.Load(sdk); ok {
		return config.(XGleanConfig)
	}
	return XGleanConfig{}
}
