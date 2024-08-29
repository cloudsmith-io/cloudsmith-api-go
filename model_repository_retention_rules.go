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
)

// checks if the RepositoryRetentionRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryRetentionRules{}

// RepositoryRetentionRules struct for RepositoryRetentionRules
type RepositoryRetentionRules struct {
	// The maximum X number of packages to retain.
	RetentionCountLimit *int64 `json:"retention_count_limit,omitempty"`
	// The X number of days of packages to retain.
	RetentionDaysLimit *int64 `json:"retention_days_limit,omitempty"`
	// If checked, the retention lifecycle rules will be activated for the repository. Any packages that don't match will be deleted automatically, and the rest are retained.
	RetentionEnabled *bool `json:"retention_enabled,omitempty"`
	// If checked, retention will apply to packages by package formats rather than across all package formats.For example, when retaining by a limit of 1 and you upload PythonPkg 1.0 and RubyPkg 1.0, no packages are deleted because they are different formats.
	RetentionGroupByFormat *bool `json:"retention_group_by_format,omitempty"`
	// If checked, retention will apply to groups of packages by name rather than all packages.<br>For example, when retaining by a limit of 1 and you upload PkgA 1.0, PkgB 1.0 and PkgB 1.1; only PkgB 1.0 is deleted because there are two (2) PkgBs and one (1) PkgA.
	RetentionGroupByName *bool `json:"retention_group_by_name,omitempty"`
	// If checked, retention will apply to packages by package type (e.g. by binary, by source, etc.), rather than across all package types for one or more formats. <br>For example, when retaining by a limit of 1 and you upload DebPackage 1.0 and DebSourcePackage 1.0, no packages are deleted because they are different package types, binary and source respectively.
	RetentionGroupByPackageType *bool `json:"retention_group_by_package_type,omitempty"`
	// The maximum X total size (in bytes) of packages to retain.
	RetentionSizeLimit *int64 `json:"retention_size_limit,omitempty"`
}

// NewRepositoryRetentionRules instantiates a new RepositoryRetentionRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryRetentionRules() *RepositoryRetentionRules {
	this := RepositoryRetentionRules{}
	return &this
}

// NewRepositoryRetentionRulesWithDefaults instantiates a new RepositoryRetentionRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryRetentionRulesWithDefaults() *RepositoryRetentionRules {
	this := RepositoryRetentionRules{}
	return &this
}

// GetRetentionCountLimit returns the RetentionCountLimit field value if set, zero value otherwise.
func (o *RepositoryRetentionRules) GetRetentionCountLimit() int64 {
	if o == nil || IsNil(o.RetentionCountLimit) {
		var ret int64
		return ret
	}
	return *o.RetentionCountLimit
}

// GetRetentionCountLimitOk returns a tuple with the RetentionCountLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRetentionRules) GetRetentionCountLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.RetentionCountLimit) {
		return nil, false
	}
	return o.RetentionCountLimit, true
}

// HasRetentionCountLimit returns a boolean if a field has been set.
func (o *RepositoryRetentionRules) HasRetentionCountLimit() bool {
	if o != nil && !IsNil(o.RetentionCountLimit) {
		return true
	}

	return false
}

// SetRetentionCountLimit gets a reference to the given int64 and assigns it to the RetentionCountLimit field.
func (o *RepositoryRetentionRules) SetRetentionCountLimit(v int64) {
	o.RetentionCountLimit = &v
}

// GetRetentionDaysLimit returns the RetentionDaysLimit field value if set, zero value otherwise.
func (o *RepositoryRetentionRules) GetRetentionDaysLimit() int64 {
	if o == nil || IsNil(o.RetentionDaysLimit) {
		var ret int64
		return ret
	}
	return *o.RetentionDaysLimit
}

// GetRetentionDaysLimitOk returns a tuple with the RetentionDaysLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRetentionRules) GetRetentionDaysLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.RetentionDaysLimit) {
		return nil, false
	}
	return o.RetentionDaysLimit, true
}

// HasRetentionDaysLimit returns a boolean if a field has been set.
func (o *RepositoryRetentionRules) HasRetentionDaysLimit() bool {
	if o != nil && !IsNil(o.RetentionDaysLimit) {
		return true
	}

	return false
}

// SetRetentionDaysLimit gets a reference to the given int64 and assigns it to the RetentionDaysLimit field.
func (o *RepositoryRetentionRules) SetRetentionDaysLimit(v int64) {
	o.RetentionDaysLimit = &v
}

// GetRetentionEnabled returns the RetentionEnabled field value if set, zero value otherwise.
func (o *RepositoryRetentionRules) GetRetentionEnabled() bool {
	if o == nil || IsNil(o.RetentionEnabled) {
		var ret bool
		return ret
	}
	return *o.RetentionEnabled
}

// GetRetentionEnabledOk returns a tuple with the RetentionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRetentionRules) GetRetentionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RetentionEnabled) {
		return nil, false
	}
	return o.RetentionEnabled, true
}

// HasRetentionEnabled returns a boolean if a field has been set.
func (o *RepositoryRetentionRules) HasRetentionEnabled() bool {
	if o != nil && !IsNil(o.RetentionEnabled) {
		return true
	}

	return false
}

// SetRetentionEnabled gets a reference to the given bool and assigns it to the RetentionEnabled field.
func (o *RepositoryRetentionRules) SetRetentionEnabled(v bool) {
	o.RetentionEnabled = &v
}

// GetRetentionGroupByFormat returns the RetentionGroupByFormat field value if set, zero value otherwise.
func (o *RepositoryRetentionRules) GetRetentionGroupByFormat() bool {
	if o == nil || IsNil(o.RetentionGroupByFormat) {
		var ret bool
		return ret
	}
	return *o.RetentionGroupByFormat
}

// GetRetentionGroupByFormatOk returns a tuple with the RetentionGroupByFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRetentionRules) GetRetentionGroupByFormatOk() (*bool, bool) {
	if o == nil || IsNil(o.RetentionGroupByFormat) {
		return nil, false
	}
	return o.RetentionGroupByFormat, true
}

// HasRetentionGroupByFormat returns a boolean if a field has been set.
func (o *RepositoryRetentionRules) HasRetentionGroupByFormat() bool {
	if o != nil && !IsNil(o.RetentionGroupByFormat) {
		return true
	}

	return false
}

// SetRetentionGroupByFormat gets a reference to the given bool and assigns it to the RetentionGroupByFormat field.
func (o *RepositoryRetentionRules) SetRetentionGroupByFormat(v bool) {
	o.RetentionGroupByFormat = &v
}

// GetRetentionGroupByName returns the RetentionGroupByName field value if set, zero value otherwise.
func (o *RepositoryRetentionRules) GetRetentionGroupByName() bool {
	if o == nil || IsNil(o.RetentionGroupByName) {
		var ret bool
		return ret
	}
	return *o.RetentionGroupByName
}

// GetRetentionGroupByNameOk returns a tuple with the RetentionGroupByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRetentionRules) GetRetentionGroupByNameOk() (*bool, bool) {
	if o == nil || IsNil(o.RetentionGroupByName) {
		return nil, false
	}
	return o.RetentionGroupByName, true
}

// HasRetentionGroupByName returns a boolean if a field has been set.
func (o *RepositoryRetentionRules) HasRetentionGroupByName() bool {
	if o != nil && !IsNil(o.RetentionGroupByName) {
		return true
	}

	return false
}

// SetRetentionGroupByName gets a reference to the given bool and assigns it to the RetentionGroupByName field.
func (o *RepositoryRetentionRules) SetRetentionGroupByName(v bool) {
	o.RetentionGroupByName = &v
}

// GetRetentionGroupByPackageType returns the RetentionGroupByPackageType field value if set, zero value otherwise.
func (o *RepositoryRetentionRules) GetRetentionGroupByPackageType() bool {
	if o == nil || IsNil(o.RetentionGroupByPackageType) {
		var ret bool
		return ret
	}
	return *o.RetentionGroupByPackageType
}

// GetRetentionGroupByPackageTypeOk returns a tuple with the RetentionGroupByPackageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRetentionRules) GetRetentionGroupByPackageTypeOk() (*bool, bool) {
	if o == nil || IsNil(o.RetentionGroupByPackageType) {
		return nil, false
	}
	return o.RetentionGroupByPackageType, true
}

// HasRetentionGroupByPackageType returns a boolean if a field has been set.
func (o *RepositoryRetentionRules) HasRetentionGroupByPackageType() bool {
	if o != nil && !IsNil(o.RetentionGroupByPackageType) {
		return true
	}

	return false
}

// SetRetentionGroupByPackageType gets a reference to the given bool and assigns it to the RetentionGroupByPackageType field.
func (o *RepositoryRetentionRules) SetRetentionGroupByPackageType(v bool) {
	o.RetentionGroupByPackageType = &v
}

// GetRetentionSizeLimit returns the RetentionSizeLimit field value if set, zero value otherwise.
func (o *RepositoryRetentionRules) GetRetentionSizeLimit() int64 {
	if o == nil || IsNil(o.RetentionSizeLimit) {
		var ret int64
		return ret
	}
	return *o.RetentionSizeLimit
}

// GetRetentionSizeLimitOk returns a tuple with the RetentionSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRetentionRules) GetRetentionSizeLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.RetentionSizeLimit) {
		return nil, false
	}
	return o.RetentionSizeLimit, true
}

// HasRetentionSizeLimit returns a boolean if a field has been set.
func (o *RepositoryRetentionRules) HasRetentionSizeLimit() bool {
	if o != nil && !IsNil(o.RetentionSizeLimit) {
		return true
	}

	return false
}

// SetRetentionSizeLimit gets a reference to the given int64 and assigns it to the RetentionSizeLimit field.
func (o *RepositoryRetentionRules) SetRetentionSizeLimit(v int64) {
	o.RetentionSizeLimit = &v
}

func (o RepositoryRetentionRules) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryRetentionRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RetentionCountLimit) {
		toSerialize["retention_count_limit"] = o.RetentionCountLimit
	}
	if !IsNil(o.RetentionDaysLimit) {
		toSerialize["retention_days_limit"] = o.RetentionDaysLimit
	}
	if !IsNil(o.RetentionEnabled) {
		toSerialize["retention_enabled"] = o.RetentionEnabled
	}
	if !IsNil(o.RetentionGroupByFormat) {
		toSerialize["retention_group_by_format"] = o.RetentionGroupByFormat
	}
	if !IsNil(o.RetentionGroupByName) {
		toSerialize["retention_group_by_name"] = o.RetentionGroupByName
	}
	if !IsNil(o.RetentionGroupByPackageType) {
		toSerialize["retention_group_by_package_type"] = o.RetentionGroupByPackageType
	}
	if !IsNil(o.RetentionSizeLimit) {
		toSerialize["retention_size_limit"] = o.RetentionSizeLimit
	}
	return toSerialize, nil
}

type NullableRepositoryRetentionRules struct {
	value *RepositoryRetentionRules
	isSet bool
}

func (v NullableRepositoryRetentionRules) Get() *RepositoryRetentionRules {
	return v.value
}

func (v *NullableRepositoryRetentionRules) Set(val *RepositoryRetentionRules) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryRetentionRules) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryRetentionRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryRetentionRules(val *RepositoryRetentionRules) *NullableRepositoryRetentionRules {
	return &NullableRepositoryRetentionRules{value: val, isSet: true}
}

func (v NullableRepositoryRetentionRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryRetentionRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
