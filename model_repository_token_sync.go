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

// RepositoryTokenSync struct for RepositoryTokenSync
type RepositoryTokenSync struct {
	// The entitlements that have been synchronised.
	Tokens []RepositoryTokenSyncTokens `json:"tokens,omitempty"`
}
