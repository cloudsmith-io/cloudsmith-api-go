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

// PackagesValidateUploadVagrant struct for PackagesValidateUploadVagrant
type PackagesValidateUploadVagrant struct {
	// The name of this package.
	Name string `json:"name"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// The virtual machine provider for the box.
	Provider string `json:"provider"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish bool `json:"republish,omitempty"`
	// The raw version for this package.
	Version string `json:"version"`
}
