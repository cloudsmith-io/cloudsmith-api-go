/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.736.13
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"time"
)

// checks if the RepositoryTokenRequestPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryTokenRequestPatch{}

// RepositoryTokenRequestPatch struct for RepositoryTokenRequestPatch
type RepositoryTokenRequestPatch struct {
	// If checked, a EULA acceptance is required for this token.
	EulaRequired *bool `json:"eula_required,omitempty"`
	// If enabled, the token will allow downloads based on configured restrictions (if any).
	IsActive *bool `json:"is_active,omitempty"`
	// The maximum download bandwidth allowed for the token. Values are expressed as the selected unit of bandwidth. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.
	LimitBandwidth     NullableInt64  `json:"limit_bandwidth,omitempty"`
	LimitBandwidthUnit NullableString `json:"limit_bandwidth_unit,omitempty"`
	// The starting date/time the token is allowed to be used from.
	LimitDateRangeFrom NullableTime `json:"limit_date_range_from,omitempty"`
	// The ending date/time the token is allowed to be used until.
	LimitDateRangeTo NullableTime `json:"limit_date_range_to,omitempty"`
	// The maximum number of unique clients allowed for the token. Please note that since clients are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.
	LimitNumClients NullableInt64 `json:"limit_num_clients,omitempty"`
	// The maximum number of downloads allowed for the token. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.
	LimitNumDownloads NullableInt64 `json:"limit_num_downloads,omitempty"`
	// The package-based search query to apply to restrict downloads to. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. This will still allow access to non-package files, such as metadata.
	LimitPackageQuery NullableString `json:"limit_package_query,omitempty"`
	// THIS WILL SOON BE DEPRECATED, please use limit_package_query instead. The path-based search query to apply to restrict downloads to. This supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. The path evaluated does not include the domain name, the namespace, the entitlement code used, the package format, etc. and it always starts with a forward slash.
	LimitPathQuery NullableString         `json:"limit_path_query,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	Name           *string                `json:"name,omitempty"`
	// The time at which the scheduled reset period has elapsed and the token limits were automatically reset to zero.
	ScheduledResetAt     NullableTime   `json:"scheduled_reset_at,omitempty"`
	ScheduledResetPeriod NullableString `json:"scheduled_reset_period,omitempty"`
	Token                *string        `json:"token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RepositoryTokenRequestPatch RepositoryTokenRequestPatch

// NewRepositoryTokenRequestPatch instantiates a new RepositoryTokenRequestPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryTokenRequestPatch() *RepositoryTokenRequestPatch {
	this := RepositoryTokenRequestPatch{}
	var limitBandwidthUnit string = "Byte"
	this.LimitBandwidthUnit = *NewNullableString(&limitBandwidthUnit)
	var scheduledResetPeriod string = "Never Reset"
	this.ScheduledResetPeriod = *NewNullableString(&scheduledResetPeriod)
	return &this
}

// NewRepositoryTokenRequestPatchWithDefaults instantiates a new RepositoryTokenRequestPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryTokenRequestPatchWithDefaults() *RepositoryTokenRequestPatch {
	this := RepositoryTokenRequestPatch{}
	var limitBandwidthUnit string = "Byte"
	this.LimitBandwidthUnit = *NewNullableString(&limitBandwidthUnit)
	var scheduledResetPeriod string = "Never Reset"
	this.ScheduledResetPeriod = *NewNullableString(&scheduledResetPeriod)
	return &this
}

// GetEulaRequired returns the EulaRequired field value if set, zero value otherwise.
func (o *RepositoryTokenRequestPatch) GetEulaRequired() bool {
	if o == nil || IsNil(o.EulaRequired) {
		var ret bool
		return ret
	}
	return *o.EulaRequired
}

// GetEulaRequiredOk returns a tuple with the EulaRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryTokenRequestPatch) GetEulaRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.EulaRequired) {
		return nil, false
	}
	return o.EulaRequired, true
}

// HasEulaRequired returns a boolean if a field has been set.
func (o *RepositoryTokenRequestPatch) HasEulaRequired() bool {
	if o != nil && !IsNil(o.EulaRequired) {
		return true
	}

	return false
}

// SetEulaRequired gets a reference to the given bool and assigns it to the EulaRequired field.
func (o *RepositoryTokenRequestPatch) SetEulaRequired(v bool) {
	o.EulaRequired = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *RepositoryTokenRequestPatch) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryTokenRequestPatch) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *RepositoryTokenRequestPatch) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *RepositoryTokenRequestPatch) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLimitBandwidth returns the LimitBandwidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequestPatch) GetLimitBandwidth() int64 {
	if o == nil || IsNil(o.LimitBandwidth.Get()) {
		var ret int64
		return ret
	}
	return *o.LimitBandwidth.Get()
}

// GetLimitBandwidthOk returns a tuple with the LimitBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequestPatch) GetLimitBandwidthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitBandwidth.Get(), o.LimitBandwidth.IsSet()
}

// HasLimitBandwidth returns a boolean if a field has been set.
func (o *RepositoryTokenRequestPatch) HasLimitBandwidth() bool {
	if o != nil && o.LimitBandwidth.IsSet() {
		return true
	}

	return false
}

// SetLimitBandwidth gets a reference to the given NullableInt64 and assigns it to the LimitBandwidth field.
func (o *RepositoryTokenRequestPatch) SetLimitBandwidth(v int64) {
	o.LimitBandwidth.Set(&v)
}

// SetLimitBandwidthNil sets the value for LimitBandwidth to be an explicit nil
func (o *RepositoryTokenRequestPatch) SetLimitBandwidthNil() {
	o.LimitBandwidth.Set(nil)
}

// UnsetLimitBandwidth ensures that no value is present for LimitBandwidth, not even an explicit nil
func (o *RepositoryTokenRequestPatch) UnsetLimitBandwidth() {
	o.LimitBandwidth.Unset()
}

// GetLimitBandwidthUnit returns the LimitBandwidthUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequestPatch) GetLimitBandwidthUnit() string {
	if o == nil || IsNil(o.LimitBandwidthUnit.Get()) {
		var ret string
		return ret
	}
	return *o.LimitBandwidthUnit.Get()
}

// GetLimitBandwidthUnitOk returns a tuple with the LimitBandwidthUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequestPatch) GetLimitBandwidthUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitBandwidthUnit.Get(), o.LimitBandwidthUnit.IsSet()
}

// HasLimitBandwidthUnit returns a boolean if a field has been set.
func (o *RepositoryTokenRequestPatch) HasLimitBandwidthUnit() bool {
	if o != nil && o.LimitBandwidthUnit.IsSet() {
		return true
	}

	return false
}

// SetLimitBandwidthUnit gets a reference to the given NullableString and assigns it to the LimitBandwidthUnit field.
func (o *RepositoryTokenRequestPatch) SetLimitBandwidthUnit(v string) {
	o.LimitBandwidthUnit.Set(&v)
}

// SetLimitBandwidthUnitNil sets the value for LimitBandwidthUnit to be an explicit nil
func (o *RepositoryTokenRequestPatch) SetLimitBandwidthUnitNil() {
	o.LimitBandwidthUnit.Set(nil)
}

// UnsetLimitBandwidthUnit ensures that no value is present for LimitBandwidthUnit, not even an explicit nil
func (o *RepositoryTokenRequestPatch) UnsetLimitBandwidthUnit() {
	o.LimitBandwidthUnit.Unset()
}

// GetLimitDateRangeFrom returns the LimitDateRangeFrom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequestPatch) GetLimitDateRangeFrom() time.Time {
	if o == nil || IsNil(o.LimitDateRangeFrom.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LimitDateRangeFrom.Get()
}

// GetLimitDateRangeFromOk returns a tuple with the LimitDateRangeFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequestPatch) GetLimitDateRangeFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitDateRangeFrom.Get(), o.LimitDateRangeFrom.IsSet()
}

// HasLimitDateRangeFrom returns a boolean if a field has been set.
func (o *RepositoryTokenRequestPatch) HasLimitDateRangeFrom() bool {
	if o != nil && o.LimitDateRangeFrom.IsSet() {
		return true
	}

	return false
}

// SetLimitDateRangeFrom gets a reference to the given NullableTime and assigns it to the LimitDateRangeFrom field.
func (o *RepositoryTokenRequestPatch) SetLimitDateRangeFrom(v time.Time) {
	o.LimitDateRangeFrom.Set(&v)
}

// SetLimitDateRangeFromNil sets the value for LimitDateRangeFrom to be an explicit nil
func (o *RepositoryTokenRequestPatch) SetLimitDateRangeFromNil() {
	o.LimitDateRangeFrom.Set(nil)
}

// UnsetLimitDateRangeFrom ensures that no value is present for LimitDateRangeFrom, not even an explicit nil
func (o *RepositoryTokenRequestPatch) UnsetLimitDateRangeFrom() {
	o.LimitDateRangeFrom.Unset()
}

// GetLimitDateRangeTo returns the LimitDateRangeTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequestPatch) GetLimitDateRangeTo() time.Time {
	if o == nil || IsNil(o.LimitDateRangeTo.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LimitDateRangeTo.Get()
}

// GetLimitDateRangeToOk returns a tuple with the LimitDateRangeTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequestPatch) GetLimitDateRangeToOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitDateRangeTo.Get(), o.LimitDateRangeTo.IsSet()
}

// HasLimitDateRangeTo returns a boolean if a field has been set.
func (o *RepositoryTokenRequestPatch) HasLimitDateRangeTo() bool {
	if o != nil && o.LimitDateRangeTo.IsSet() {
		return true
	}

	return false
}

// SetLimitDateRangeTo gets a reference to the given NullableTime and assigns it to the LimitDateRangeTo field.
func (o *RepositoryTokenRequestPatch) SetLimitDateRangeTo(v time.Time) {
	o.LimitDateRangeTo.Set(&v)
}

// SetLimitDateRangeToNil sets the value for LimitDateRangeTo to be an explicit nil
func (o *RepositoryTokenRequestPatch) SetLimitDateRangeToNil() {
	o.LimitDateRangeTo.Set(nil)
}

// UnsetLimitDateRangeTo ensures that no value is present for LimitDateRangeTo, not even an explicit nil
func (o *RepositoryTokenRequestPatch) UnsetLimitDateRangeTo() {
	o.LimitDateRangeTo.Unset()
}

// GetLimitNumClients returns the LimitNumClients field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequestPatch) GetLimitNumClients() int64 {
	if o == nil || IsNil(o.LimitNumClients.Get()) {
		var ret int64
		return ret
	}
	return *o.LimitNumClients.Get()
}

// GetLimitNumClientsOk returns a tuple with the LimitNumClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequestPatch) GetLimitNumClientsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitNumClients.Get(), o.LimitNumClients.IsSet()
}

// HasLimitNumClients returns a boolean if a field has been set.
func (o *RepositoryTokenRequestPatch) HasLimitNumClients() bool {
	if o != nil && o.LimitNumClients.IsSet() {
		return true
	}

	return false
}

// SetLimitNumClients gets a reference to the given NullableInt64 and assigns it to the LimitNumClients field.
func (o *RepositoryTokenRequestPatch) SetLimitNumClients(v int64) {
	o.LimitNumClients.Set(&v)
}

// SetLimitNumClientsNil sets the value for LimitNumClients to be an explicit nil
func (o *RepositoryTokenRequestPatch) SetLimitNumClientsNil() {
	o.LimitNumClients.Set(nil)
}

// UnsetLimitNumClients ensures that no value is present for LimitNumClients, not even an explicit nil
func (o *RepositoryTokenRequestPatch) UnsetLimitNumClients() {
	o.LimitNumClients.Unset()
}

// GetLimitNumDownloads returns the LimitNumDownloads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequestPatch) GetLimitNumDownloads() int64 {
	if o == nil || IsNil(o.LimitNumDownloads.Get()) {
		var ret int64
		return ret
	}
	return *o.LimitNumDownloads.Get()
}

// GetLimitNumDownloadsOk returns a tuple with the LimitNumDownloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequestPatch) GetLimitNumDownloadsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitNumDownloads.Get(), o.LimitNumDownloads.IsSet()
}

// HasLimitNumDownloads returns a boolean if a field has been set.
func (o *RepositoryTokenRequestPatch) HasLimitNumDownloads() bool {
	if o != nil && o.LimitNumDownloads.IsSet() {
		return true
	}

	return false
}

// SetLimitNumDownloads gets a reference to the given NullableInt64 and assigns it to the LimitNumDownloads field.
func (o *RepositoryTokenRequestPatch) SetLimitNumDownloads(v int64) {
	o.LimitNumDownloads.Set(&v)
}

// SetLimitNumDownloadsNil sets the value for LimitNumDownloads to be an explicit nil
func (o *RepositoryTokenRequestPatch) SetLimitNumDownloadsNil() {
	o.LimitNumDownloads.Set(nil)
}

// UnsetLimitNumDownloads ensures that no value is present for LimitNumDownloads, not even an explicit nil
func (o *RepositoryTokenRequestPatch) UnsetLimitNumDownloads() {
	o.LimitNumDownloads.Unset()
}

// GetLimitPackageQuery returns the LimitPackageQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequestPatch) GetLimitPackageQuery() string {
	if o == nil || IsNil(o.LimitPackageQuery.Get()) {
		var ret string
		return ret
	}
	return *o.LimitPackageQuery.Get()
}

// GetLimitPackageQueryOk returns a tuple with the LimitPackageQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequestPatch) GetLimitPackageQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitPackageQuery.Get(), o.LimitPackageQuery.IsSet()
}

// HasLimitPackageQuery returns a boolean if a field has been set.
func (o *RepositoryTokenRequestPatch) HasLimitPackageQuery() bool {
	if o != nil && o.LimitPackageQuery.IsSet() {
		return true
	}

	return false
}

// SetLimitPackageQuery gets a reference to the given NullableString and assigns it to the LimitPackageQuery field.
func (o *RepositoryTokenRequestPatch) SetLimitPackageQuery(v string) {
	o.LimitPackageQuery.Set(&v)
}

// SetLimitPackageQueryNil sets the value for LimitPackageQuery to be an explicit nil
func (o *RepositoryTokenRequestPatch) SetLimitPackageQueryNil() {
	o.LimitPackageQuery.Set(nil)
}

// UnsetLimitPackageQuery ensures that no value is present for LimitPackageQuery, not even an explicit nil
func (o *RepositoryTokenRequestPatch) UnsetLimitPackageQuery() {
	o.LimitPackageQuery.Unset()
}

// GetLimitPathQuery returns the LimitPathQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequestPatch) GetLimitPathQuery() string {
	if o == nil || IsNil(o.LimitPathQuery.Get()) {
		var ret string
		return ret
	}
	return *o.LimitPathQuery.Get()
}

// GetLimitPathQueryOk returns a tuple with the LimitPathQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequestPatch) GetLimitPathQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitPathQuery.Get(), o.LimitPathQuery.IsSet()
}

// HasLimitPathQuery returns a boolean if a field has been set.
func (o *RepositoryTokenRequestPatch) HasLimitPathQuery() bool {
	if o != nil && o.LimitPathQuery.IsSet() {
		return true
	}

	return false
}

// SetLimitPathQuery gets a reference to the given NullableString and assigns it to the LimitPathQuery field.
func (o *RepositoryTokenRequestPatch) SetLimitPathQuery(v string) {
	o.LimitPathQuery.Set(&v)
}

// SetLimitPathQueryNil sets the value for LimitPathQuery to be an explicit nil
func (o *RepositoryTokenRequestPatch) SetLimitPathQueryNil() {
	o.LimitPathQuery.Set(nil)
}

// UnsetLimitPathQuery ensures that no value is present for LimitPathQuery, not even an explicit nil
func (o *RepositoryTokenRequestPatch) UnsetLimitPathQuery() {
	o.LimitPathQuery.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequestPatch) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequestPatch) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RepositoryTokenRequestPatch) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RepositoryTokenRequestPatch) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RepositoryTokenRequestPatch) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryTokenRequestPatch) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RepositoryTokenRequestPatch) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RepositoryTokenRequestPatch) SetName(v string) {
	o.Name = &v
}

// GetScheduledResetAt returns the ScheduledResetAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequestPatch) GetScheduledResetAt() time.Time {
	if o == nil || IsNil(o.ScheduledResetAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledResetAt.Get()
}

// GetScheduledResetAtOk returns a tuple with the ScheduledResetAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequestPatch) GetScheduledResetAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledResetAt.Get(), o.ScheduledResetAt.IsSet()
}

// HasScheduledResetAt returns a boolean if a field has been set.
func (o *RepositoryTokenRequestPatch) HasScheduledResetAt() bool {
	if o != nil && o.ScheduledResetAt.IsSet() {
		return true
	}

	return false
}

// SetScheduledResetAt gets a reference to the given NullableTime and assigns it to the ScheduledResetAt field.
func (o *RepositoryTokenRequestPatch) SetScheduledResetAt(v time.Time) {
	o.ScheduledResetAt.Set(&v)
}

// SetScheduledResetAtNil sets the value for ScheduledResetAt to be an explicit nil
func (o *RepositoryTokenRequestPatch) SetScheduledResetAtNil() {
	o.ScheduledResetAt.Set(nil)
}

// UnsetScheduledResetAt ensures that no value is present for ScheduledResetAt, not even an explicit nil
func (o *RepositoryTokenRequestPatch) UnsetScheduledResetAt() {
	o.ScheduledResetAt.Unset()
}

// GetScheduledResetPeriod returns the ScheduledResetPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequestPatch) GetScheduledResetPeriod() string {
	if o == nil || IsNil(o.ScheduledResetPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.ScheduledResetPeriod.Get()
}

// GetScheduledResetPeriodOk returns a tuple with the ScheduledResetPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequestPatch) GetScheduledResetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledResetPeriod.Get(), o.ScheduledResetPeriod.IsSet()
}

// HasScheduledResetPeriod returns a boolean if a field has been set.
func (o *RepositoryTokenRequestPatch) HasScheduledResetPeriod() bool {
	if o != nil && o.ScheduledResetPeriod.IsSet() {
		return true
	}

	return false
}

// SetScheduledResetPeriod gets a reference to the given NullableString and assigns it to the ScheduledResetPeriod field.
func (o *RepositoryTokenRequestPatch) SetScheduledResetPeriod(v string) {
	o.ScheduledResetPeriod.Set(&v)
}

// SetScheduledResetPeriodNil sets the value for ScheduledResetPeriod to be an explicit nil
func (o *RepositoryTokenRequestPatch) SetScheduledResetPeriodNil() {
	o.ScheduledResetPeriod.Set(nil)
}

// UnsetScheduledResetPeriod ensures that no value is present for ScheduledResetPeriod, not even an explicit nil
func (o *RepositoryTokenRequestPatch) UnsetScheduledResetPeriod() {
	o.ScheduledResetPeriod.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RepositoryTokenRequestPatch) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryTokenRequestPatch) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RepositoryTokenRequestPatch) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RepositoryTokenRequestPatch) SetToken(v string) {
	o.Token = &v
}

func (o RepositoryTokenRequestPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryTokenRequestPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EulaRequired) {
		toSerialize["eula_required"] = o.EulaRequired
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if o.LimitBandwidth.IsSet() {
		toSerialize["limit_bandwidth"] = o.LimitBandwidth.Get()
	}
	if o.LimitBandwidthUnit.IsSet() {
		toSerialize["limit_bandwidth_unit"] = o.LimitBandwidthUnit.Get()
	}
	if o.LimitDateRangeFrom.IsSet() {
		toSerialize["limit_date_range_from"] = o.LimitDateRangeFrom.Get()
	}
	if o.LimitDateRangeTo.IsSet() {
		toSerialize["limit_date_range_to"] = o.LimitDateRangeTo.Get()
	}
	if o.LimitNumClients.IsSet() {
		toSerialize["limit_num_clients"] = o.LimitNumClients.Get()
	}
	if o.LimitNumDownloads.IsSet() {
		toSerialize["limit_num_downloads"] = o.LimitNumDownloads.Get()
	}
	if o.LimitPackageQuery.IsSet() {
		toSerialize["limit_package_query"] = o.LimitPackageQuery.Get()
	}
	if o.LimitPathQuery.IsSet() {
		toSerialize["limit_path_query"] = o.LimitPathQuery.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.ScheduledResetAt.IsSet() {
		toSerialize["scheduled_reset_at"] = o.ScheduledResetAt.Get()
	}
	if o.ScheduledResetPeriod.IsSet() {
		toSerialize["scheduled_reset_period"] = o.ScheduledResetPeriod.Get()
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RepositoryTokenRequestPatch) UnmarshalJSON(data []byte) (err error) {
	varRepositoryTokenRequestPatch := _RepositoryTokenRequestPatch{}

	err = json.Unmarshal(data, &varRepositoryTokenRequestPatch)

	if err != nil {
		return err
	}

	*o = RepositoryTokenRequestPatch(varRepositoryTokenRequestPatch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "eula_required")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "limit_bandwidth")
		delete(additionalProperties, "limit_bandwidth_unit")
		delete(additionalProperties, "limit_date_range_from")
		delete(additionalProperties, "limit_date_range_to")
		delete(additionalProperties, "limit_num_clients")
		delete(additionalProperties, "limit_num_downloads")
		delete(additionalProperties, "limit_package_query")
		delete(additionalProperties, "limit_path_query")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "name")
		delete(additionalProperties, "scheduled_reset_at")
		delete(additionalProperties, "scheduled_reset_period")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRepositoryTokenRequestPatch struct {
	value *RepositoryTokenRequestPatch
	isSet bool
}

func (v NullableRepositoryTokenRequestPatch) Get() *RepositoryTokenRequestPatch {
	return v.value
}

func (v *NullableRepositoryTokenRequestPatch) Set(val *RepositoryTokenRequestPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryTokenRequestPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryTokenRequestPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryTokenRequestPatch(val *RepositoryTokenRequestPatch) *NullableRepositoryTokenRequestPatch {
	return &NullableRepositoryTokenRequestPatch{value: val, isSet: true}
}

func (v NullableRepositoryTokenRequestPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryTokenRequestPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
