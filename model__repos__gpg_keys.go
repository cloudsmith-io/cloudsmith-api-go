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

// ReposGpgKeys struct for ReposGpgKeys
type ReposGpgKeys struct {
	// If selected this is the active key for this repository.
	Active bool `json:"active,omitempty"`
	//
	Comment string `json:"comment,omitempty"`
	//
	CreatedAt string `json:"created_at,omitempty"`
	// If selected this is the default key for this repository.
	Default bool `json:"default,omitempty"`
	// The long identifier used by GPG for this key.
	Fingerprint string `json:"fingerprint,omitempty"`
	//
	FingerprintShort string `json:"fingerprint_short,omitempty"`
	// The public key given to repository users.
	PublicKey string `json:"public_key,omitempty"`
}
