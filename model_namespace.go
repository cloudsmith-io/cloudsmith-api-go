/*
 * Cloudsmith API
 *
 * The API to the Cloudsmith Service
 *
 * API version: 0.51.38
 * Contact: support@cloudsmith.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudsmith

// Namespace struct for Namespace
type Namespace struct {
	//
	Name string `json:"name,omitempty"`
	//
	Slug string `json:"slug,omitempty"`
	//
	SlugPerm string `json:"slug_perm,omitempty"`
	//
	TypeName string `json:"type_name,omitempty"`
}
