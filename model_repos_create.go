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

// ReposCreate struct for ReposCreate
type ReposCreate struct {
	// None
	Description string `json:"description,omitempty"`
	// If checked, files contained in packages will be indexed, which increase the synchronisation time required for packages. Note that it is recommended you keep this enabled unless the synchronisation time is significantly impacted.
	IndexFiles bool `json:"index_files,omitempty"`
	// None
	IsOpenSource bool `json:"is_open_source,omitempty"`
	// None
	IsPrivate bool `json:"is_private,omitempty"`
	// None
	IsPublic bool `json:"is_public,omitempty"`
	// A descriptive name for the repository.
	Name string `json:"name"`
	// None
	Namespace string `json:"namespace"`
	// None
	RepositoryTypeStr string `json:"repository_type_str,omitempty"`
	// None
	Slug string `json:"slug"`
}
