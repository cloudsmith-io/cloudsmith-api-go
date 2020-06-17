/*
 * Cloudsmith API
 *
 * The API to the Cloudsmith Service
 *
 * API version: 0.51.50
 * Contact: support@cloudsmith.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudsmith

// Distribution struct for Distribution
type Distribution struct {
	//
	Format string `json:"format,omitempty"`
	//
	FormatUrl string `json:"format_url,omitempty"`
	//
	Name string `json:"name"`
	//
	SelfUrl string `json:"self_url,omitempty"`
	// The slug identifier for this distribution
	Slug string `json:"slug,omitempty"`
	//
	Variants string `json:"variants,omitempty"`
	// A list of the versions for this distribution
	Versions []DistrosVersions `json:"versions,omitempty"`
}
