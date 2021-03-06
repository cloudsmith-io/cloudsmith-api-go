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

// UserTokenCreate struct for UserTokenCreate
type UserTokenCreate struct {
	// Email address to authenticate with
	Email string `json:"email,omitempty"`
	// Password to authenticate with
	Password string `json:"password,omitempty"`
}
