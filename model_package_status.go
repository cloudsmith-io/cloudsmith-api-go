/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.498.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"time"
)

// checks if the PackageStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageStatus{}

// PackageStatus struct for PackageStatus
type PackageStatus struct {
	IsCancellable       *bool   `json:"is_cancellable,omitempty"`
	IsCopyable          *bool   `json:"is_copyable,omitempty"`
	IsDeleteable        *bool   `json:"is_deleteable,omitempty"`
	IsDownloadable      *bool   `json:"is_downloadable,omitempty"`
	IsMoveable          *bool   `json:"is_moveable,omitempty"`
	IsQuarantinable     *bool   `json:"is_quarantinable,omitempty"`
	IsQuarantined       *bool   `json:"is_quarantined,omitempty"`
	IsResyncable        *bool   `json:"is_resyncable,omitempty"`
	IsSecurityScannable *bool   `json:"is_security_scannable,omitempty"`
	IsSyncAwaiting      *bool   `json:"is_sync_awaiting,omitempty"`
	IsSyncCompleted     *bool   `json:"is_sync_completed,omitempty"`
	IsSyncFailed        *bool   `json:"is_sync_failed,omitempty"`
	IsSyncInFlight      *bool   `json:"is_sync_in_flight,omitempty"`
	IsSyncInProgress    *bool   `json:"is_sync_in_progress,omitempty"`
	SelfUrl             *string `json:"self_url,omitempty"`
	// The synchronisation (in progress) stage of the package.
	Stage    *int64  `json:"stage,omitempty"`
	StageStr *string `json:"stage_str,omitempty"`
	// The datetime the package stage was updated at.
	StageUpdatedAt *time.Time `json:"stage_updated_at,omitempty"`
	// The synchronisation status of the package.
	Status *int64 `json:"status,omitempty"`
	// A textual description for the synchronous status reason (if any
	StatusReason NullableString `json:"status_reason,omitempty"`
	StatusStr    *string        `json:"status_str,omitempty"`
	// The datetime the package status was updated at.
	StatusUpdatedAt *time.Time `json:"status_updated_at,omitempty"`
	// The datetime the package sync was finished at.
	SyncFinishedAt NullableTime `json:"sync_finished_at,omitempty"`
	// Synchronisation progress (from 0-100)
	SyncProgress *int64 `json:"sync_progress,omitempty"`
}

// NewPackageStatus instantiates a new PackageStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageStatus() *PackageStatus {
	this := PackageStatus{}
	return &this
}

// NewPackageStatusWithDefaults instantiates a new PackageStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageStatusWithDefaults() *PackageStatus {
	this := PackageStatus{}
	return &this
}

// GetIsCancellable returns the IsCancellable field value if set, zero value otherwise.
func (o *PackageStatus) GetIsCancellable() bool {
	if o == nil || IsNil(o.IsCancellable) {
		var ret bool
		return ret
	}
	return *o.IsCancellable
}

// GetIsCancellableOk returns a tuple with the IsCancellable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetIsCancellableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCancellable) {
		return nil, false
	}
	return o.IsCancellable, true
}

// HasIsCancellable returns a boolean if a field has been set.
func (o *PackageStatus) HasIsCancellable() bool {
	if o != nil && !IsNil(o.IsCancellable) {
		return true
	}

	return false
}

// SetIsCancellable gets a reference to the given bool and assigns it to the IsCancellable field.
func (o *PackageStatus) SetIsCancellable(v bool) {
	o.IsCancellable = &v
}

// GetIsCopyable returns the IsCopyable field value if set, zero value otherwise.
func (o *PackageStatus) GetIsCopyable() bool {
	if o == nil || IsNil(o.IsCopyable) {
		var ret bool
		return ret
	}
	return *o.IsCopyable
}

// GetIsCopyableOk returns a tuple with the IsCopyable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetIsCopyableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCopyable) {
		return nil, false
	}
	return o.IsCopyable, true
}

// HasIsCopyable returns a boolean if a field has been set.
func (o *PackageStatus) HasIsCopyable() bool {
	if o != nil && !IsNil(o.IsCopyable) {
		return true
	}

	return false
}

// SetIsCopyable gets a reference to the given bool and assigns it to the IsCopyable field.
func (o *PackageStatus) SetIsCopyable(v bool) {
	o.IsCopyable = &v
}

// GetIsDeleteable returns the IsDeleteable field value if set, zero value otherwise.
func (o *PackageStatus) GetIsDeleteable() bool {
	if o == nil || IsNil(o.IsDeleteable) {
		var ret bool
		return ret
	}
	return *o.IsDeleteable
}

// GetIsDeleteableOk returns a tuple with the IsDeleteable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetIsDeleteableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleteable) {
		return nil, false
	}
	return o.IsDeleteable, true
}

// HasIsDeleteable returns a boolean if a field has been set.
func (o *PackageStatus) HasIsDeleteable() bool {
	if o != nil && !IsNil(o.IsDeleteable) {
		return true
	}

	return false
}

// SetIsDeleteable gets a reference to the given bool and assigns it to the IsDeleteable field.
func (o *PackageStatus) SetIsDeleteable(v bool) {
	o.IsDeleteable = &v
}

// GetIsDownloadable returns the IsDownloadable field value if set, zero value otherwise.
func (o *PackageStatus) GetIsDownloadable() bool {
	if o == nil || IsNil(o.IsDownloadable) {
		var ret bool
		return ret
	}
	return *o.IsDownloadable
}

// GetIsDownloadableOk returns a tuple with the IsDownloadable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetIsDownloadableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDownloadable) {
		return nil, false
	}
	return o.IsDownloadable, true
}

// HasIsDownloadable returns a boolean if a field has been set.
func (o *PackageStatus) HasIsDownloadable() bool {
	if o != nil && !IsNil(o.IsDownloadable) {
		return true
	}

	return false
}

// SetIsDownloadable gets a reference to the given bool and assigns it to the IsDownloadable field.
func (o *PackageStatus) SetIsDownloadable(v bool) {
	o.IsDownloadable = &v
}

// GetIsMoveable returns the IsMoveable field value if set, zero value otherwise.
func (o *PackageStatus) GetIsMoveable() bool {
	if o == nil || IsNil(o.IsMoveable) {
		var ret bool
		return ret
	}
	return *o.IsMoveable
}

// GetIsMoveableOk returns a tuple with the IsMoveable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetIsMoveableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMoveable) {
		return nil, false
	}
	return o.IsMoveable, true
}

// HasIsMoveable returns a boolean if a field has been set.
func (o *PackageStatus) HasIsMoveable() bool {
	if o != nil && !IsNil(o.IsMoveable) {
		return true
	}

	return false
}

// SetIsMoveable gets a reference to the given bool and assigns it to the IsMoveable field.
func (o *PackageStatus) SetIsMoveable(v bool) {
	o.IsMoveable = &v
}

// GetIsQuarantinable returns the IsQuarantinable field value if set, zero value otherwise.
func (o *PackageStatus) GetIsQuarantinable() bool {
	if o == nil || IsNil(o.IsQuarantinable) {
		var ret bool
		return ret
	}
	return *o.IsQuarantinable
}

// GetIsQuarantinableOk returns a tuple with the IsQuarantinable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetIsQuarantinableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsQuarantinable) {
		return nil, false
	}
	return o.IsQuarantinable, true
}

// HasIsQuarantinable returns a boolean if a field has been set.
func (o *PackageStatus) HasIsQuarantinable() bool {
	if o != nil && !IsNil(o.IsQuarantinable) {
		return true
	}

	return false
}

// SetIsQuarantinable gets a reference to the given bool and assigns it to the IsQuarantinable field.
func (o *PackageStatus) SetIsQuarantinable(v bool) {
	o.IsQuarantinable = &v
}

// GetIsQuarantined returns the IsQuarantined field value if set, zero value otherwise.
func (o *PackageStatus) GetIsQuarantined() bool {
	if o == nil || IsNil(o.IsQuarantined) {
		var ret bool
		return ret
	}
	return *o.IsQuarantined
}

// GetIsQuarantinedOk returns a tuple with the IsQuarantined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetIsQuarantinedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsQuarantined) {
		return nil, false
	}
	return o.IsQuarantined, true
}

// HasIsQuarantined returns a boolean if a field has been set.
func (o *PackageStatus) HasIsQuarantined() bool {
	if o != nil && !IsNil(o.IsQuarantined) {
		return true
	}

	return false
}

// SetIsQuarantined gets a reference to the given bool and assigns it to the IsQuarantined field.
func (o *PackageStatus) SetIsQuarantined(v bool) {
	o.IsQuarantined = &v
}

// GetIsResyncable returns the IsResyncable field value if set, zero value otherwise.
func (o *PackageStatus) GetIsResyncable() bool {
	if o == nil || IsNil(o.IsResyncable) {
		var ret bool
		return ret
	}
	return *o.IsResyncable
}

// GetIsResyncableOk returns a tuple with the IsResyncable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetIsResyncableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsResyncable) {
		return nil, false
	}
	return o.IsResyncable, true
}

// HasIsResyncable returns a boolean if a field has been set.
func (o *PackageStatus) HasIsResyncable() bool {
	if o != nil && !IsNil(o.IsResyncable) {
		return true
	}

	return false
}

// SetIsResyncable gets a reference to the given bool and assigns it to the IsResyncable field.
func (o *PackageStatus) SetIsResyncable(v bool) {
	o.IsResyncable = &v
}

// GetIsSecurityScannable returns the IsSecurityScannable field value if set, zero value otherwise.
func (o *PackageStatus) GetIsSecurityScannable() bool {
	if o == nil || IsNil(o.IsSecurityScannable) {
		var ret bool
		return ret
	}
	return *o.IsSecurityScannable
}

// GetIsSecurityScannableOk returns a tuple with the IsSecurityScannable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetIsSecurityScannableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSecurityScannable) {
		return nil, false
	}
	return o.IsSecurityScannable, true
}

// HasIsSecurityScannable returns a boolean if a field has been set.
func (o *PackageStatus) HasIsSecurityScannable() bool {
	if o != nil && !IsNil(o.IsSecurityScannable) {
		return true
	}

	return false
}

// SetIsSecurityScannable gets a reference to the given bool and assigns it to the IsSecurityScannable field.
func (o *PackageStatus) SetIsSecurityScannable(v bool) {
	o.IsSecurityScannable = &v
}

// GetIsSyncAwaiting returns the IsSyncAwaiting field value if set, zero value otherwise.
func (o *PackageStatus) GetIsSyncAwaiting() bool {
	if o == nil || IsNil(o.IsSyncAwaiting) {
		var ret bool
		return ret
	}
	return *o.IsSyncAwaiting
}

// GetIsSyncAwaitingOk returns a tuple with the IsSyncAwaiting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetIsSyncAwaitingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSyncAwaiting) {
		return nil, false
	}
	return o.IsSyncAwaiting, true
}

// HasIsSyncAwaiting returns a boolean if a field has been set.
func (o *PackageStatus) HasIsSyncAwaiting() bool {
	if o != nil && !IsNil(o.IsSyncAwaiting) {
		return true
	}

	return false
}

// SetIsSyncAwaiting gets a reference to the given bool and assigns it to the IsSyncAwaiting field.
func (o *PackageStatus) SetIsSyncAwaiting(v bool) {
	o.IsSyncAwaiting = &v
}

// GetIsSyncCompleted returns the IsSyncCompleted field value if set, zero value otherwise.
func (o *PackageStatus) GetIsSyncCompleted() bool {
	if o == nil || IsNil(o.IsSyncCompleted) {
		var ret bool
		return ret
	}
	return *o.IsSyncCompleted
}

// GetIsSyncCompletedOk returns a tuple with the IsSyncCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetIsSyncCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSyncCompleted) {
		return nil, false
	}
	return o.IsSyncCompleted, true
}

// HasIsSyncCompleted returns a boolean if a field has been set.
func (o *PackageStatus) HasIsSyncCompleted() bool {
	if o != nil && !IsNil(o.IsSyncCompleted) {
		return true
	}

	return false
}

// SetIsSyncCompleted gets a reference to the given bool and assigns it to the IsSyncCompleted field.
func (o *PackageStatus) SetIsSyncCompleted(v bool) {
	o.IsSyncCompleted = &v
}

// GetIsSyncFailed returns the IsSyncFailed field value if set, zero value otherwise.
func (o *PackageStatus) GetIsSyncFailed() bool {
	if o == nil || IsNil(o.IsSyncFailed) {
		var ret bool
		return ret
	}
	return *o.IsSyncFailed
}

// GetIsSyncFailedOk returns a tuple with the IsSyncFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetIsSyncFailedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSyncFailed) {
		return nil, false
	}
	return o.IsSyncFailed, true
}

// HasIsSyncFailed returns a boolean if a field has been set.
func (o *PackageStatus) HasIsSyncFailed() bool {
	if o != nil && !IsNil(o.IsSyncFailed) {
		return true
	}

	return false
}

// SetIsSyncFailed gets a reference to the given bool and assigns it to the IsSyncFailed field.
func (o *PackageStatus) SetIsSyncFailed(v bool) {
	o.IsSyncFailed = &v
}

// GetIsSyncInFlight returns the IsSyncInFlight field value if set, zero value otherwise.
func (o *PackageStatus) GetIsSyncInFlight() bool {
	if o == nil || IsNil(o.IsSyncInFlight) {
		var ret bool
		return ret
	}
	return *o.IsSyncInFlight
}

// GetIsSyncInFlightOk returns a tuple with the IsSyncInFlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetIsSyncInFlightOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSyncInFlight) {
		return nil, false
	}
	return o.IsSyncInFlight, true
}

// HasIsSyncInFlight returns a boolean if a field has been set.
func (o *PackageStatus) HasIsSyncInFlight() bool {
	if o != nil && !IsNil(o.IsSyncInFlight) {
		return true
	}

	return false
}

// SetIsSyncInFlight gets a reference to the given bool and assigns it to the IsSyncInFlight field.
func (o *PackageStatus) SetIsSyncInFlight(v bool) {
	o.IsSyncInFlight = &v
}

// GetIsSyncInProgress returns the IsSyncInProgress field value if set, zero value otherwise.
func (o *PackageStatus) GetIsSyncInProgress() bool {
	if o == nil || IsNil(o.IsSyncInProgress) {
		var ret bool
		return ret
	}
	return *o.IsSyncInProgress
}

// GetIsSyncInProgressOk returns a tuple with the IsSyncInProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetIsSyncInProgressOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSyncInProgress) {
		return nil, false
	}
	return o.IsSyncInProgress, true
}

// HasIsSyncInProgress returns a boolean if a field has been set.
func (o *PackageStatus) HasIsSyncInProgress() bool {
	if o != nil && !IsNil(o.IsSyncInProgress) {
		return true
	}

	return false
}

// SetIsSyncInProgress gets a reference to the given bool and assigns it to the IsSyncInProgress field.
func (o *PackageStatus) SetIsSyncInProgress(v bool) {
	o.IsSyncInProgress = &v
}

// GetSelfUrl returns the SelfUrl field value if set, zero value otherwise.
func (o *PackageStatus) GetSelfUrl() string {
	if o == nil || IsNil(o.SelfUrl) {
		var ret string
		return ret
	}
	return *o.SelfUrl
}

// GetSelfUrlOk returns a tuple with the SelfUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetSelfUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SelfUrl) {
		return nil, false
	}
	return o.SelfUrl, true
}

// HasSelfUrl returns a boolean if a field has been set.
func (o *PackageStatus) HasSelfUrl() bool {
	if o != nil && !IsNil(o.SelfUrl) {
		return true
	}

	return false
}

// SetSelfUrl gets a reference to the given string and assigns it to the SelfUrl field.
func (o *PackageStatus) SetSelfUrl(v string) {
	o.SelfUrl = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *PackageStatus) GetStage() int64 {
	if o == nil || IsNil(o.Stage) {
		var ret int64
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetStageOk() (*int64, bool) {
	if o == nil || IsNil(o.Stage) {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *PackageStatus) HasStage() bool {
	if o != nil && !IsNil(o.Stage) {
		return true
	}

	return false
}

// SetStage gets a reference to the given int64 and assigns it to the Stage field.
func (o *PackageStatus) SetStage(v int64) {
	o.Stage = &v
}

// GetStageStr returns the StageStr field value if set, zero value otherwise.
func (o *PackageStatus) GetStageStr() string {
	if o == nil || IsNil(o.StageStr) {
		var ret string
		return ret
	}
	return *o.StageStr
}

// GetStageStrOk returns a tuple with the StageStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetStageStrOk() (*string, bool) {
	if o == nil || IsNil(o.StageStr) {
		return nil, false
	}
	return o.StageStr, true
}

// HasStageStr returns a boolean if a field has been set.
func (o *PackageStatus) HasStageStr() bool {
	if o != nil && !IsNil(o.StageStr) {
		return true
	}

	return false
}

// SetStageStr gets a reference to the given string and assigns it to the StageStr field.
func (o *PackageStatus) SetStageStr(v string) {
	o.StageStr = &v
}

// GetStageUpdatedAt returns the StageUpdatedAt field value if set, zero value otherwise.
func (o *PackageStatus) GetStageUpdatedAt() time.Time {
	if o == nil || IsNil(o.StageUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.StageUpdatedAt
}

// GetStageUpdatedAtOk returns a tuple with the StageUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetStageUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StageUpdatedAt) {
		return nil, false
	}
	return o.StageUpdatedAt, true
}

// HasStageUpdatedAt returns a boolean if a field has been set.
func (o *PackageStatus) HasStageUpdatedAt() bool {
	if o != nil && !IsNil(o.StageUpdatedAt) {
		return true
	}

	return false
}

// SetStageUpdatedAt gets a reference to the given time.Time and assigns it to the StageUpdatedAt field.
func (o *PackageStatus) SetStageUpdatedAt(v time.Time) {
	o.StageUpdatedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PackageStatus) GetStatus() int64 {
	if o == nil || IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PackageStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *PackageStatus) SetStatus(v int64) {
	o.Status = &v
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageStatus) GetStatusReason() string {
	if o == nil || IsNil(o.StatusReason.Get()) {
		var ret string
		return ret
	}
	return *o.StatusReason.Get()
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageStatus) GetStatusReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatusReason.Get(), o.StatusReason.IsSet()
}

// HasStatusReason returns a boolean if a field has been set.
func (o *PackageStatus) HasStatusReason() bool {
	if o != nil && o.StatusReason.IsSet() {
		return true
	}

	return false
}

// SetStatusReason gets a reference to the given NullableString and assigns it to the StatusReason field.
func (o *PackageStatus) SetStatusReason(v string) {
	o.StatusReason.Set(&v)
}

// SetStatusReasonNil sets the value for StatusReason to be an explicit nil
func (o *PackageStatus) SetStatusReasonNil() {
	o.StatusReason.Set(nil)
}

// UnsetStatusReason ensures that no value is present for StatusReason, not even an explicit nil
func (o *PackageStatus) UnsetStatusReason() {
	o.StatusReason.Unset()
}

// GetStatusStr returns the StatusStr field value if set, zero value otherwise.
func (o *PackageStatus) GetStatusStr() string {
	if o == nil || IsNil(o.StatusStr) {
		var ret string
		return ret
	}
	return *o.StatusStr
}

// GetStatusStrOk returns a tuple with the StatusStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetStatusStrOk() (*string, bool) {
	if o == nil || IsNil(o.StatusStr) {
		return nil, false
	}
	return o.StatusStr, true
}

// HasStatusStr returns a boolean if a field has been set.
func (o *PackageStatus) HasStatusStr() bool {
	if o != nil && !IsNil(o.StatusStr) {
		return true
	}

	return false
}

// SetStatusStr gets a reference to the given string and assigns it to the StatusStr field.
func (o *PackageStatus) SetStatusStr(v string) {
	o.StatusStr = &v
}

// GetStatusUpdatedAt returns the StatusUpdatedAt field value if set, zero value otherwise.
func (o *PackageStatus) GetStatusUpdatedAt() time.Time {
	if o == nil || IsNil(o.StatusUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.StatusUpdatedAt
}

// GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetStatusUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StatusUpdatedAt) {
		return nil, false
	}
	return o.StatusUpdatedAt, true
}

// HasStatusUpdatedAt returns a boolean if a field has been set.
func (o *PackageStatus) HasStatusUpdatedAt() bool {
	if o != nil && !IsNil(o.StatusUpdatedAt) {
		return true
	}

	return false
}

// SetStatusUpdatedAt gets a reference to the given time.Time and assigns it to the StatusUpdatedAt field.
func (o *PackageStatus) SetStatusUpdatedAt(v time.Time) {
	o.StatusUpdatedAt = &v
}

// GetSyncFinishedAt returns the SyncFinishedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageStatus) GetSyncFinishedAt() time.Time {
	if o == nil || IsNil(o.SyncFinishedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SyncFinishedAt.Get()
}

// GetSyncFinishedAtOk returns a tuple with the SyncFinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageStatus) GetSyncFinishedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyncFinishedAt.Get(), o.SyncFinishedAt.IsSet()
}

// HasSyncFinishedAt returns a boolean if a field has been set.
func (o *PackageStatus) HasSyncFinishedAt() bool {
	if o != nil && o.SyncFinishedAt.IsSet() {
		return true
	}

	return false
}

// SetSyncFinishedAt gets a reference to the given NullableTime and assigns it to the SyncFinishedAt field.
func (o *PackageStatus) SetSyncFinishedAt(v time.Time) {
	o.SyncFinishedAt.Set(&v)
}

// SetSyncFinishedAtNil sets the value for SyncFinishedAt to be an explicit nil
func (o *PackageStatus) SetSyncFinishedAtNil() {
	o.SyncFinishedAt.Set(nil)
}

// UnsetSyncFinishedAt ensures that no value is present for SyncFinishedAt, not even an explicit nil
func (o *PackageStatus) UnsetSyncFinishedAt() {
	o.SyncFinishedAt.Unset()
}

// GetSyncProgress returns the SyncProgress field value if set, zero value otherwise.
func (o *PackageStatus) GetSyncProgress() int64 {
	if o == nil || IsNil(o.SyncProgress) {
		var ret int64
		return ret
	}
	return *o.SyncProgress
}

// GetSyncProgressOk returns a tuple with the SyncProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageStatus) GetSyncProgressOk() (*int64, bool) {
	if o == nil || IsNil(o.SyncProgress) {
		return nil, false
	}
	return o.SyncProgress, true
}

// HasSyncProgress returns a boolean if a field has been set.
func (o *PackageStatus) HasSyncProgress() bool {
	if o != nil && !IsNil(o.SyncProgress) {
		return true
	}

	return false
}

// SetSyncProgress gets a reference to the given int64 and assigns it to the SyncProgress field.
func (o *PackageStatus) SetSyncProgress(v int64) {
	o.SyncProgress = &v
}

func (o PackageStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsCancellable) {
		toSerialize["is_cancellable"] = o.IsCancellable
	}
	if !IsNil(o.IsCopyable) {
		toSerialize["is_copyable"] = o.IsCopyable
	}
	if !IsNil(o.IsDeleteable) {
		toSerialize["is_deleteable"] = o.IsDeleteable
	}
	if !IsNil(o.IsDownloadable) {
		toSerialize["is_downloadable"] = o.IsDownloadable
	}
	if !IsNil(o.IsMoveable) {
		toSerialize["is_moveable"] = o.IsMoveable
	}
	if !IsNil(o.IsQuarantinable) {
		toSerialize["is_quarantinable"] = o.IsQuarantinable
	}
	if !IsNil(o.IsQuarantined) {
		toSerialize["is_quarantined"] = o.IsQuarantined
	}
	if !IsNil(o.IsResyncable) {
		toSerialize["is_resyncable"] = o.IsResyncable
	}
	if !IsNil(o.IsSecurityScannable) {
		toSerialize["is_security_scannable"] = o.IsSecurityScannable
	}
	if !IsNil(o.IsSyncAwaiting) {
		toSerialize["is_sync_awaiting"] = o.IsSyncAwaiting
	}
	if !IsNil(o.IsSyncCompleted) {
		toSerialize["is_sync_completed"] = o.IsSyncCompleted
	}
	if !IsNil(o.IsSyncFailed) {
		toSerialize["is_sync_failed"] = o.IsSyncFailed
	}
	if !IsNil(o.IsSyncInFlight) {
		toSerialize["is_sync_in_flight"] = o.IsSyncInFlight
	}
	if !IsNil(o.IsSyncInProgress) {
		toSerialize["is_sync_in_progress"] = o.IsSyncInProgress
	}
	if !IsNil(o.SelfUrl) {
		toSerialize["self_url"] = o.SelfUrl
	}
	if !IsNil(o.Stage) {
		toSerialize["stage"] = o.Stage
	}
	if !IsNil(o.StageStr) {
		toSerialize["stage_str"] = o.StageStr
	}
	if !IsNil(o.StageUpdatedAt) {
		toSerialize["stage_updated_at"] = o.StageUpdatedAt
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.StatusReason.IsSet() {
		toSerialize["status_reason"] = o.StatusReason.Get()
	}
	if !IsNil(o.StatusStr) {
		toSerialize["status_str"] = o.StatusStr
	}
	if !IsNil(o.StatusUpdatedAt) {
		toSerialize["status_updated_at"] = o.StatusUpdatedAt
	}
	if o.SyncFinishedAt.IsSet() {
		toSerialize["sync_finished_at"] = o.SyncFinishedAt.Get()
	}
	if !IsNil(o.SyncProgress) {
		toSerialize["sync_progress"] = o.SyncProgress
	}
	return toSerialize, nil
}

type NullablePackageStatus struct {
	value *PackageStatus
	isSet bool
}

func (v NullablePackageStatus) Get() *PackageStatus {
	return v.value
}

func (v *NullablePackageStatus) Set(val *PackageStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageStatus(val *PackageStatus) *NullablePackageStatus {
	return &NullablePackageStatus{value: val, isSet: true}
}

func (v NullablePackageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
