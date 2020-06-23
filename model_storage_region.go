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

// StorageRegion struct for StorageRegion
type StorageRegion struct {
	// Name of the storage region
	Label string `json:"label"`
	// Slug for the storage region
	Slug string `json:"slug"`
}
