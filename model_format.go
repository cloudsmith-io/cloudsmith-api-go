/*
 * Cloudsmith API
 *
 * The API to the Cloudsmith Service
 *
 * API version: 0.51.52
 * Contact: support@cloudsmith.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudsmith

// Format struct for Format
type Format struct {
	// Description of the package format
	Description string `json:"description"`
	// The distributions supported by this package format
	Distributions []FormatsDistributions `json:"distributions,omitempty"`
	// A non-exhaustive list of extensions supported
	Extensions []string `json:"extensions"`
	// Name for the package format
	Name string `json:"name"`
	// If true the package format is a premium-only feature
	Premium bool `json:"premium"`
	// The minimum plan id required for this package format
	PremiumPlanId string `json:"premium_plan_id,omitempty"`
	// The minimum plan name required for this package format
	PremiumPlanName string `json:"premium_plan_name,omitempty"`
	// Slug for the package format
	Slug string `json:"slug"`
	// A set of what the package format supports
	Supports map[string]interface{} `json:"supports"`
}
