/*
 * Cloudsmith API
 *
 * The API to the Cloudsmith Service
 *
 * API version: 0.51.50
 * Contact: support@cloudsmith.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudsmith

// Package struct for Package
type Package struct {
	//
	Architectures []PackagesOwnerRepoArchitectures `json:"architectures,omitempty"`
	//
	CdnUrl string `json:"cdn_url,omitempty"`
	//
	ChecksumMd5 string `json:"checksum_md5,omitempty"`
	//
	ChecksumSha1 string `json:"checksum_sha1,omitempty"`
	//
	ChecksumSha256 string `json:"checksum_sha256,omitempty"`
	//
	ChecksumSha512 string `json:"checksum_sha512,omitempty"`
	// A textual description of this package.
	Description string `json:"description,omitempty"`
	//
	Distro map[string]interface{} `json:"distro,omitempty"`
	//
	DistroVersion map[string]interface{} `json:"distro_version,omitempty"`
	//
	Downloads int64 `json:"downloads,omitempty"`
	// The epoch of the package version (if any).
	Epoch int64 `json:"epoch,omitempty"`
	//
	Extension string `json:"extension,omitempty"`
	//
	Filename string `json:"filename,omitempty"`
	//
	Files []PackagesOwnerRepoFiles `json:"files,omitempty"`
	//
	Format string `json:"format,omitempty"`
	//
	FormatUrl string `json:"format_url,omitempty"`
	// Unique and permanent identifier for the package.
	IdentifierPerm string `json:"identifier_perm,omitempty"`
	//
	Indexed bool `json:"indexed,omitempty"`
	//
	IsSyncAwaiting bool `json:"is_sync_awaiting,omitempty"`
	//
	IsSyncCompleted bool `json:"is_sync_completed,omitempty"`
	//
	IsSyncFailed bool `json:"is_sync_failed,omitempty"`
	//
	IsSyncInFlight bool `json:"is_sync_in_flight,omitempty"`
	//
	IsSyncInProgress bool `json:"is_sync_in_progress,omitempty"`
	// The license of this package.
	License string `json:"license,omitempty"`
	// The name of this package.
	Name string `json:"name,omitempty"`
	//
	Namespace string `json:"namespace,omitempty"`
	//
	NamespaceUrl string `json:"namespace_url,omitempty"`
	//
	NumFiles int64 `json:"num_files,omitempty"`
	// The type of package contents.
	PackageType string `json:"package_type,omitempty"`
	// The release of the package version (if any).
	Release string `json:"release,omitempty"`
	//
	Repository string `json:"repository,omitempty"`
	//
	RepositoryUrl string `json:"repository_url,omitempty"`
	//
	SelfHtmlUrl string `json:"self_html_url,omitempty"`
	//
	SelfUrl string `json:"self_url,omitempty"`
	// The calculated size of the package.
	Size int64 `json:"size,omitempty"`
	//
	Slug string `json:"slug,omitempty"`
	//
	SlugPerm string `json:"slug_perm,omitempty"`
	// The synchronisation (in progress) stage of the package.
	Stage string `json:"stage,omitempty"`
	//
	StageStr string `json:"stage_str,omitempty"`
	// The datetime the package stage was updated at.
	StageUpdatedAt string `json:"stage_updated_at,omitempty"`
	// The synchronisation status of the package.
	Status string `json:"status,omitempty"`
	// A textual description for the synchronous status reason (if any
	StatusReason string `json:"status_reason,omitempty"`
	//
	StatusStr string `json:"status_str,omitempty"`
	// The datetime the package status was updated at.
	StatusUpdatedAt string `json:"status_updated_at,omitempty"`
	//
	StatusUrl string `json:"status_url,omitempty"`
	//
	Subtype string `json:"subtype,omitempty"`
	// A one-liner synopsis of this package.
	Summary string `json:"summary,omitempty"`
	// The datetime the package sync was finished at.
	SyncFinishedAt string `json:"sync_finished_at,omitempty"`
	// Synchronisation progress (from 0-100)
	SyncProgress int64 `json:"sync_progress,omitempty"`
	//
	TypeDisplay string `json:"type_display,omitempty"`
	// The date this package was uploaded.
	UploadedAt string `json:"uploaded_at,omitempty"`
	//
	Uploader string `json:"uploader,omitempty"`
	//
	UploaderUrl string `json:"uploader_url,omitempty"`
	// The raw version for this package.
	Version string `json:"version,omitempty"`
	//
	VersionOrig string `json:"version_orig,omitempty"`
}
