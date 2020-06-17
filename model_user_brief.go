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

// UserBrief struct for UserBrief
type UserBrief struct {
	// If true then you're logged in as a user.
	Authenticated bool `json:"authenticated,omitempty"`
	// Your email address that we use to contact you. This is only visible to you.
	Email string `json:"email,omitempty"`
	// The full name of the user (if any).
	Name string `json:"name,omitempty"`
	// The URL for the full profile of the user.
	ProfileUrl string `json:"profile_url,omitempty"`
	//
	SelfUrl string `json:"self_url,omitempty"`
	//
	Slug string `json:"slug,omitempty"`
	//
	SlugPerm string `json:"slug_perm,omitempty"`
}
