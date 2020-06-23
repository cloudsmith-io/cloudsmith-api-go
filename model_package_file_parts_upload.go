/*
 * Cloudsmith API
 *
 * The API to the Cloudsmith Service
 *
 * API version: 0.51.65
 * Contact: support@cloudsmith.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudsmith

// PackageFilePartsUpload struct for PackageFilePartsUpload
type PackageFilePartsUpload struct {
	// The identifier for the file to use uploading parts.
	Identifier string `json:"identifier,omitempty"`
	// The querystring to use for the next-step PUT upload.
	UploadQuerystring string `json:"upload_querystring,omitempty"`
	// The URL to use for the next-step PUT upload
	UploadUrl string `json:"upload_url,omitempty"`
}
