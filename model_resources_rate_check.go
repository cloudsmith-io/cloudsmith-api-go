/*
 * Cloudsmith API
 *
 * The API to the Cloudsmith Service
 *
 * API version: 0.51.48
 * Contact: support@cloudsmith.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudsmith

// ResourcesRateCheck struct for ResourcesRateCheck
type ResourcesRateCheck struct {
	// Rate limit values per resource
	Resources map[string]interface{} `json:"resources,omitempty"`
}
