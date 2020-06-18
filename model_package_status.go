/*
 * Cloudsmith API
 *
 * The API to the Cloudsmith Service
 *
 * API version: 0.51.52
 * Contact: support@cloudsmith.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudsmith

// PackageStatus struct for PackageStatus
type PackageStatus struct {
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
	//
	SelfUrl string `json:"self_url,omitempty"`
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
	// The datetime the package sync was finished at.
	SyncFinishedAt string `json:"sync_finished_at,omitempty"`
	// Synchronisation progress (from 0-100)
	SyncProgress int64 `json:"sync_progress,omitempty"`
}
