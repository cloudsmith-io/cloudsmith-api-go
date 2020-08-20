/*
 * Cloudsmith API
 *
 * The API to the Cloudsmith Service
 *
 * API version: 0.52.62
 * Contact: support@cloudsmith.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudsmith

// PackagesValidateUploadNpm struct for PackagesValidateUploadNpm
type PackagesValidateUploadNpm struct {
	// The default npm dist-tag for this package/version - This will replace any other package/version if they are using the same tag.
	NpmDistTag string `json:"npm_dist_tag,omitempty"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish bool `json:"republish,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags string `json:"tags,omitempty"`
}
