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

// EntitlementsPartialUpdate struct for EntitlementsPartialUpdate
type EntitlementsPartialUpdate struct {
	// If enabled, the token will allow downloads based on configured restrictions (if any).
	IsActive bool `json:"is_active,omitempty"`
	// The starting date/time the token is allowed to be used from.
	LimitDateRangeFrom string `json:"limit_date_range_from,omitempty"`
	// The ending date/time the token is allowed to be used until.
	LimitDateRangeTo string `json:"limit_date_range_to,omitempty"`
	// The maximum number of unique clients allowed for the token. Please note that since clients are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.
	LimitNumClients int64 `json:"limit_num_clients,omitempty"`
	// The maximum number of downloads allowed for the token. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.
	LimitNumDownloads int64 `json:"limit_num_downloads,omitempty"`
	// The package-based search query to apply to restrict downloads to. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. This will still allow access to non-package files, such as metadata. For package formats that support dynamic metadata indexes, the contents of the metadata will also be filtered.
	LimitPackageQuery string `json:"limit_package_query,omitempty"`
	// The path-based search query to apply to restrict downloads to. This supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. The path evaluated does not include the domain name, the namespace, the entitlement code used, the package format, etc. and it always starts with a forward slash.
	LimitPathQuery string `json:"limit_path_query,omitempty"`
	// None
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// None
	Name string `json:"name,omitempty"`
	// None
	Token string `json:"token,omitempty"`
}
