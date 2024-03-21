/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.372.0
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
	IsDownloadable   *bool   `json:"is_downloadable,omitempty"`
	IsQuarantined    *bool   `json:"is_quarantined,omitempty"`
	IsSyncAwaiting   *bool   `json:"is_sync_awaiting,omitempty"`
	IsSyncCompleted  *bool   `json:"is_sync_completed,omitempty"`
	IsSyncFailed     *bool   `json:"is_sync_failed,omitempty"`
	IsSyncInFlight   *bool   `json:"is_sync_in_flight,omitempty"`
	IsSyncInProgress *bool   `json:"is_sync_in_progress,omitempty"`
	SelfUrl          *string `json:"self_url,omitempty"`
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
	if !IsNil(o.IsDownloadable) {
		toSerialize["is_downloadable"] = o.IsDownloadable
	}
	if !IsNil(o.IsQuarantined) {
		toSerialize["is_quarantined"] = o.IsQuarantined
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
