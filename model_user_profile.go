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

// UserProfile struct for UserProfile
type UserProfile struct {
	//
	Company string `json:"company,omitempty"`
	//
	FirstName string `json:"first_name"`
	//
	JobTitle string `json:"job_title,omitempty"`
	//
	JoinedAt string `json:"joined_at,omitempty"`
	//
	LastName string `json:"last_name"`
	//
	Name string `json:"name,omitempty"`
	//
	Slug string `json:"slug,omitempty"`
	//
	SlugPerm string `json:"slug_perm,omitempty"`
	// Your tagline is a sentence about you. Make it funny. Make it professional. Either way, it's public and it represents who you are.
	Tagline string `json:"tagline,omitempty"`
	//
	Url string `json:"url,omitempty"`
}
