/*
 * Cloudsmith API
 *
 * The API to the Cloudsmith Service
 *
 * API version: 0.51.68
 * Contact: support@cloudsmith.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudsmith

// StatusBasic struct for StatusBasic
type StatusBasic struct {
	// The message describing the state of the API.
	Detail string `json:"detail,omitempty"`
	// The current version for the Cloudsmith service.
	Version string `json:"version,omitempty"`
}
