/*
Cloudsmith API

The API to the Cloudsmith Service

API version: 1.40.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// EntitlementsRefresh struct for EntitlementsRefresh
type EntitlementsRefresh struct {
	// If checked, a EULA acceptance is required for this token.
	EulaRequired *bool `json:"eula_required,omitempty"`
	// If enabled, the token will allow downloads based on configured restrictions (if any).
	IsActive *bool `json:"is_active,omitempty"`
	// The maximum download bandwidth allowed for the token. Values are expressed as the selected unit of bandwidth. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.
	LimitBandwidth *int64 `json:"limit_bandwidth,omitempty"`
	// None
	LimitBandwidthUnit *string `json:"limit_bandwidth_unit,omitempty"`
	// The starting date/time the token is allowed to be used from.
	LimitDateRangeFrom *string `json:"limit_date_range_from,omitempty"`
	// The ending date/time the token is allowed to be used until.
	LimitDateRangeTo *string `json:"limit_date_range_to,omitempty"`
	// The maximum number of unique clients allowed for the token. Please note that since clients are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.
	LimitNumClients *int64 `json:"limit_num_clients,omitempty"`
	// The maximum number of downloads allowed for the token. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.
	LimitNumDownloads *int64 `json:"limit_num_downloads,omitempty"`
	// The package-based search query to apply to restrict downloads to. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. This will still allow access to non-package files, such as metadata.
	LimitPackageQuery *string `json:"limit_package_query,omitempty"`
	// The path-based search query to apply to restrict downloads to. This supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. The path evaluated does not include the domain name, the namespace, the entitlement code used, the package format, etc. and it always starts with a forward slash.
	LimitPathQuery *string `json:"limit_path_query,omitempty"`
	// None
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The time at which the scheduled reset period has elapsed and the token limits were automatically reset to zero.
	ScheduledResetAt *string `json:"scheduled_reset_at,omitempty"`
	// None
	ScheduledResetPeriod *string `json:"scheduled_reset_period,omitempty"`
	// None
	Token *string `json:"token,omitempty"`
}

// NewEntitlementsRefresh instantiates a new EntitlementsRefresh object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementsRefresh() *EntitlementsRefresh {
	this := EntitlementsRefresh{}
	return &this
}

// NewEntitlementsRefreshWithDefaults instantiates a new EntitlementsRefresh object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementsRefreshWithDefaults() *EntitlementsRefresh {
	this := EntitlementsRefresh{}
	return &this
}

// GetEulaRequired returns the EulaRequired field value if set, zero value otherwise.
func (o *EntitlementsRefresh) GetEulaRequired() bool {
	if o == nil || o.EulaRequired == nil {
		var ret bool
		return ret
	}
	return *o.EulaRequired
}

// GetEulaRequiredOk returns a tuple with the EulaRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsRefresh) GetEulaRequiredOk() (*bool, bool) {
	if o == nil || o.EulaRequired == nil {
		return nil, false
	}
	return o.EulaRequired, true
}

// HasEulaRequired returns a boolean if a field has been set.
func (o *EntitlementsRefresh) HasEulaRequired() bool {
	if o != nil && o.EulaRequired != nil {
		return true
	}

	return false
}

// SetEulaRequired gets a reference to the given bool and assigns it to the EulaRequired field.
func (o *EntitlementsRefresh) SetEulaRequired(v bool) {
	o.EulaRequired = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *EntitlementsRefresh) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsRefresh) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *EntitlementsRefresh) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *EntitlementsRefresh) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLimitBandwidth returns the LimitBandwidth field value if set, zero value otherwise.
func (o *EntitlementsRefresh) GetLimitBandwidth() int64 {
	if o == nil || o.LimitBandwidth == nil {
		var ret int64
		return ret
	}
	return *o.LimitBandwidth
}

// GetLimitBandwidthOk returns a tuple with the LimitBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsRefresh) GetLimitBandwidthOk() (*int64, bool) {
	if o == nil || o.LimitBandwidth == nil {
		return nil, false
	}
	return o.LimitBandwidth, true
}

// HasLimitBandwidth returns a boolean if a field has been set.
func (o *EntitlementsRefresh) HasLimitBandwidth() bool {
	if o != nil && o.LimitBandwidth != nil {
		return true
	}

	return false
}

// SetLimitBandwidth gets a reference to the given int64 and assigns it to the LimitBandwidth field.
func (o *EntitlementsRefresh) SetLimitBandwidth(v int64) {
	o.LimitBandwidth = &v
}

// GetLimitBandwidthUnit returns the LimitBandwidthUnit field value if set, zero value otherwise.
func (o *EntitlementsRefresh) GetLimitBandwidthUnit() string {
	if o == nil || o.LimitBandwidthUnit == nil {
		var ret string
		return ret
	}
	return *o.LimitBandwidthUnit
}

// GetLimitBandwidthUnitOk returns a tuple with the LimitBandwidthUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsRefresh) GetLimitBandwidthUnitOk() (*string, bool) {
	if o == nil || o.LimitBandwidthUnit == nil {
		return nil, false
	}
	return o.LimitBandwidthUnit, true
}

// HasLimitBandwidthUnit returns a boolean if a field has been set.
func (o *EntitlementsRefresh) HasLimitBandwidthUnit() bool {
	if o != nil && o.LimitBandwidthUnit != nil {
		return true
	}

	return false
}

// SetLimitBandwidthUnit gets a reference to the given string and assigns it to the LimitBandwidthUnit field.
func (o *EntitlementsRefresh) SetLimitBandwidthUnit(v string) {
	o.LimitBandwidthUnit = &v
}

// GetLimitDateRangeFrom returns the LimitDateRangeFrom field value if set, zero value otherwise.
func (o *EntitlementsRefresh) GetLimitDateRangeFrom() string {
	if o == nil || o.LimitDateRangeFrom == nil {
		var ret string
		return ret
	}
	return *o.LimitDateRangeFrom
}

// GetLimitDateRangeFromOk returns a tuple with the LimitDateRangeFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsRefresh) GetLimitDateRangeFromOk() (*string, bool) {
	if o == nil || o.LimitDateRangeFrom == nil {
		return nil, false
	}
	return o.LimitDateRangeFrom, true
}

// HasLimitDateRangeFrom returns a boolean if a field has been set.
func (o *EntitlementsRefresh) HasLimitDateRangeFrom() bool {
	if o != nil && o.LimitDateRangeFrom != nil {
		return true
	}

	return false
}

// SetLimitDateRangeFrom gets a reference to the given string and assigns it to the LimitDateRangeFrom field.
func (o *EntitlementsRefresh) SetLimitDateRangeFrom(v string) {
	o.LimitDateRangeFrom = &v
}

// GetLimitDateRangeTo returns the LimitDateRangeTo field value if set, zero value otherwise.
func (o *EntitlementsRefresh) GetLimitDateRangeTo() string {
	if o == nil || o.LimitDateRangeTo == nil {
		var ret string
		return ret
	}
	return *o.LimitDateRangeTo
}

// GetLimitDateRangeToOk returns a tuple with the LimitDateRangeTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsRefresh) GetLimitDateRangeToOk() (*string, bool) {
	if o == nil || o.LimitDateRangeTo == nil {
		return nil, false
	}
	return o.LimitDateRangeTo, true
}

// HasLimitDateRangeTo returns a boolean if a field has been set.
func (o *EntitlementsRefresh) HasLimitDateRangeTo() bool {
	if o != nil && o.LimitDateRangeTo != nil {
		return true
	}

	return false
}

// SetLimitDateRangeTo gets a reference to the given string and assigns it to the LimitDateRangeTo field.
func (o *EntitlementsRefresh) SetLimitDateRangeTo(v string) {
	o.LimitDateRangeTo = &v
}

// GetLimitNumClients returns the LimitNumClients field value if set, zero value otherwise.
func (o *EntitlementsRefresh) GetLimitNumClients() int64 {
	if o == nil || o.LimitNumClients == nil {
		var ret int64
		return ret
	}
	return *o.LimitNumClients
}

// GetLimitNumClientsOk returns a tuple with the LimitNumClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsRefresh) GetLimitNumClientsOk() (*int64, bool) {
	if o == nil || o.LimitNumClients == nil {
		return nil, false
	}
	return o.LimitNumClients, true
}

// HasLimitNumClients returns a boolean if a field has been set.
func (o *EntitlementsRefresh) HasLimitNumClients() bool {
	if o != nil && o.LimitNumClients != nil {
		return true
	}

	return false
}

// SetLimitNumClients gets a reference to the given int64 and assigns it to the LimitNumClients field.
func (o *EntitlementsRefresh) SetLimitNumClients(v int64) {
	o.LimitNumClients = &v
}

// GetLimitNumDownloads returns the LimitNumDownloads field value if set, zero value otherwise.
func (o *EntitlementsRefresh) GetLimitNumDownloads() int64 {
	if o == nil || o.LimitNumDownloads == nil {
		var ret int64
		return ret
	}
	return *o.LimitNumDownloads
}

// GetLimitNumDownloadsOk returns a tuple with the LimitNumDownloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsRefresh) GetLimitNumDownloadsOk() (*int64, bool) {
	if o == nil || o.LimitNumDownloads == nil {
		return nil, false
	}
	return o.LimitNumDownloads, true
}

// HasLimitNumDownloads returns a boolean if a field has been set.
func (o *EntitlementsRefresh) HasLimitNumDownloads() bool {
	if o != nil && o.LimitNumDownloads != nil {
		return true
	}

	return false
}

// SetLimitNumDownloads gets a reference to the given int64 and assigns it to the LimitNumDownloads field.
func (o *EntitlementsRefresh) SetLimitNumDownloads(v int64) {
	o.LimitNumDownloads = &v
}

// GetLimitPackageQuery returns the LimitPackageQuery field value if set, zero value otherwise.
func (o *EntitlementsRefresh) GetLimitPackageQuery() string {
	if o == nil || o.LimitPackageQuery == nil {
		var ret string
		return ret
	}
	return *o.LimitPackageQuery
}

// GetLimitPackageQueryOk returns a tuple with the LimitPackageQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsRefresh) GetLimitPackageQueryOk() (*string, bool) {
	if o == nil || o.LimitPackageQuery == nil {
		return nil, false
	}
	return o.LimitPackageQuery, true
}

// HasLimitPackageQuery returns a boolean if a field has been set.
func (o *EntitlementsRefresh) HasLimitPackageQuery() bool {
	if o != nil && o.LimitPackageQuery != nil {
		return true
	}

	return false
}

// SetLimitPackageQuery gets a reference to the given string and assigns it to the LimitPackageQuery field.
func (o *EntitlementsRefresh) SetLimitPackageQuery(v string) {
	o.LimitPackageQuery = &v
}

// GetLimitPathQuery returns the LimitPathQuery field value if set, zero value otherwise.
func (o *EntitlementsRefresh) GetLimitPathQuery() string {
	if o == nil || o.LimitPathQuery == nil {
		var ret string
		return ret
	}
	return *o.LimitPathQuery
}

// GetLimitPathQueryOk returns a tuple with the LimitPathQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsRefresh) GetLimitPathQueryOk() (*string, bool) {
	if o == nil || o.LimitPathQuery == nil {
		return nil, false
	}
	return o.LimitPathQuery, true
}

// HasLimitPathQuery returns a boolean if a field has been set.
func (o *EntitlementsRefresh) HasLimitPathQuery() bool {
	if o != nil && o.LimitPathQuery != nil {
		return true
	}

	return false
}

// SetLimitPathQuery gets a reference to the given string and assigns it to the LimitPathQuery field.
func (o *EntitlementsRefresh) SetLimitPathQuery(v string) {
	o.LimitPathQuery = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *EntitlementsRefresh) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsRefresh) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *EntitlementsRefresh) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *EntitlementsRefresh) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetScheduledResetAt returns the ScheduledResetAt field value if set, zero value otherwise.
func (o *EntitlementsRefresh) GetScheduledResetAt() string {
	if o == nil || o.ScheduledResetAt == nil {
		var ret string
		return ret
	}
	return *o.ScheduledResetAt
}

// GetScheduledResetAtOk returns a tuple with the ScheduledResetAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsRefresh) GetScheduledResetAtOk() (*string, bool) {
	if o == nil || o.ScheduledResetAt == nil {
		return nil, false
	}
	return o.ScheduledResetAt, true
}

// HasScheduledResetAt returns a boolean if a field has been set.
func (o *EntitlementsRefresh) HasScheduledResetAt() bool {
	if o != nil && o.ScheduledResetAt != nil {
		return true
	}

	return false
}

// SetScheduledResetAt gets a reference to the given string and assigns it to the ScheduledResetAt field.
func (o *EntitlementsRefresh) SetScheduledResetAt(v string) {
	o.ScheduledResetAt = &v
}

// GetScheduledResetPeriod returns the ScheduledResetPeriod field value if set, zero value otherwise.
func (o *EntitlementsRefresh) GetScheduledResetPeriod() string {
	if o == nil || o.ScheduledResetPeriod == nil {
		var ret string
		return ret
	}
	return *o.ScheduledResetPeriod
}

// GetScheduledResetPeriodOk returns a tuple with the ScheduledResetPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsRefresh) GetScheduledResetPeriodOk() (*string, bool) {
	if o == nil || o.ScheduledResetPeriod == nil {
		return nil, false
	}
	return o.ScheduledResetPeriod, true
}

// HasScheduledResetPeriod returns a boolean if a field has been set.
func (o *EntitlementsRefresh) HasScheduledResetPeriod() bool {
	if o != nil && o.ScheduledResetPeriod != nil {
		return true
	}

	return false
}

// SetScheduledResetPeriod gets a reference to the given string and assigns it to the ScheduledResetPeriod field.
func (o *EntitlementsRefresh) SetScheduledResetPeriod(v string) {
	o.ScheduledResetPeriod = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *EntitlementsRefresh) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsRefresh) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *EntitlementsRefresh) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *EntitlementsRefresh) SetToken(v string) {
	o.Token = &v
}

func (o EntitlementsRefresh) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EulaRequired != nil {
		toSerialize["eula_required"] = o.EulaRequired
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.LimitBandwidth != nil {
		toSerialize["limit_bandwidth"] = o.LimitBandwidth
	}
	if o.LimitBandwidthUnit != nil {
		toSerialize["limit_bandwidth_unit"] = o.LimitBandwidthUnit
	}
	if o.LimitDateRangeFrom != nil {
		toSerialize["limit_date_range_from"] = o.LimitDateRangeFrom
	}
	if o.LimitDateRangeTo != nil {
		toSerialize["limit_date_range_to"] = o.LimitDateRangeTo
	}
	if o.LimitNumClients != nil {
		toSerialize["limit_num_clients"] = o.LimitNumClients
	}
	if o.LimitNumDownloads != nil {
		toSerialize["limit_num_downloads"] = o.LimitNumDownloads
	}
	if o.LimitPackageQuery != nil {
		toSerialize["limit_package_query"] = o.LimitPackageQuery
	}
	if o.LimitPathQuery != nil {
		toSerialize["limit_path_query"] = o.LimitPathQuery
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.ScheduledResetAt != nil {
		toSerialize["scheduled_reset_at"] = o.ScheduledResetAt
	}
	if o.ScheduledResetPeriod != nil {
		toSerialize["scheduled_reset_period"] = o.ScheduledResetPeriod
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableEntitlementsRefresh struct {
	value *EntitlementsRefresh
	isSet bool
}

func (v NullableEntitlementsRefresh) Get() *EntitlementsRefresh {
	return v.value
}

func (v *NullableEntitlementsRefresh) Set(val *EntitlementsRefresh) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementsRefresh) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementsRefresh) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementsRefresh(val *EntitlementsRefresh) *NullableEntitlementsRefresh {
	return &NullableEntitlementsRefresh{value: val, isSet: true}
}

func (v NullableEntitlementsRefresh) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementsRefresh) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
