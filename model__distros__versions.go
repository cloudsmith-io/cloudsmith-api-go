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

// DistrosVersions struct for DistrosVersions
type DistrosVersions struct {
	// The textual name for this version.
	Name string `json:"name,omitempty"`
	// The slug identifier for this version
	Slug string `json:"slug,omitempty"`
}
