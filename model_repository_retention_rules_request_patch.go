/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.617.2
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the RepositoryRetentionRulesRequestPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryRetentionRulesRequestPatch{}

// RepositoryRetentionRulesRequestPatch struct for RepositoryRetentionRulesRequestPatch
type RepositoryRetentionRulesRequestPatch struct {
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
	RetentionSizeLimit   *int64 `json:"retention_size_limit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RepositoryRetentionRulesRequestPatch RepositoryRetentionRulesRequestPatch

// NewRepositoryRetentionRulesRequestPatch instantiates a new RepositoryRetentionRulesRequestPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryRetentionRulesRequestPatch() *RepositoryRetentionRulesRequestPatch {
	this := RepositoryRetentionRulesRequestPatch{}
	return &this
}

// NewRepositoryRetentionRulesRequestPatchWithDefaults instantiates a new RepositoryRetentionRulesRequestPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryRetentionRulesRequestPatchWithDefaults() *RepositoryRetentionRulesRequestPatch {
	this := RepositoryRetentionRulesRequestPatch{}
	return &this
}

// GetRetentionCountLimit returns the RetentionCountLimit field value if set, zero value otherwise.
func (o *RepositoryRetentionRulesRequestPatch) GetRetentionCountLimit() int64 {
	if o == nil || IsNil(o.RetentionCountLimit) {
		var ret int64
		return ret
	}
	return *o.RetentionCountLimit
}

// GetRetentionCountLimitOk returns a tuple with the RetentionCountLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRetentionRulesRequestPatch) GetRetentionCountLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.RetentionCountLimit) {
		return nil, false
	}
	return o.RetentionCountLimit, true
}

// HasRetentionCountLimit returns a boolean if a field has been set.
func (o *RepositoryRetentionRulesRequestPatch) HasRetentionCountLimit() bool {
	if o != nil && !IsNil(o.RetentionCountLimit) {
		return true
	}

	return false
}

// SetRetentionCountLimit gets a reference to the given int64 and assigns it to the RetentionCountLimit field.
func (o *RepositoryRetentionRulesRequestPatch) SetRetentionCountLimit(v int64) {
	o.RetentionCountLimit = &v
}

// GetRetentionDaysLimit returns the RetentionDaysLimit field value if set, zero value otherwise.
func (o *RepositoryRetentionRulesRequestPatch) GetRetentionDaysLimit() int64 {
	if o == nil || IsNil(o.RetentionDaysLimit) {
		var ret int64
		return ret
	}
	return *o.RetentionDaysLimit
}

// GetRetentionDaysLimitOk returns a tuple with the RetentionDaysLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRetentionRulesRequestPatch) GetRetentionDaysLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.RetentionDaysLimit) {
		return nil, false
	}
	return o.RetentionDaysLimit, true
}

// HasRetentionDaysLimit returns a boolean if a field has been set.
func (o *RepositoryRetentionRulesRequestPatch) HasRetentionDaysLimit() bool {
	if o != nil && !IsNil(o.RetentionDaysLimit) {
		return true
	}

	return false
}

// SetRetentionDaysLimit gets a reference to the given int64 and assigns it to the RetentionDaysLimit field.
func (o *RepositoryRetentionRulesRequestPatch) SetRetentionDaysLimit(v int64) {
	o.RetentionDaysLimit = &v
}

// GetRetentionEnabled returns the RetentionEnabled field value if set, zero value otherwise.
func (o *RepositoryRetentionRulesRequestPatch) GetRetentionEnabled() bool {
	if o == nil || IsNil(o.RetentionEnabled) {
		var ret bool
		return ret
	}
	return *o.RetentionEnabled
}

// GetRetentionEnabledOk returns a tuple with the RetentionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRetentionRulesRequestPatch) GetRetentionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RetentionEnabled) {
		return nil, false
	}
	return o.RetentionEnabled, true
}

// HasRetentionEnabled returns a boolean if a field has been set.
func (o *RepositoryRetentionRulesRequestPatch) HasRetentionEnabled() bool {
	if o != nil && !IsNil(o.RetentionEnabled) {
		return true
	}

	return false
}

// SetRetentionEnabled gets a reference to the given bool and assigns it to the RetentionEnabled field.
func (o *RepositoryRetentionRulesRequestPatch) SetRetentionEnabled(v bool) {
	o.RetentionEnabled = &v
}

// GetRetentionGroupByFormat returns the RetentionGroupByFormat field value if set, zero value otherwise.
func (o *RepositoryRetentionRulesRequestPatch) GetRetentionGroupByFormat() bool {
	if o == nil || IsNil(o.RetentionGroupByFormat) {
		var ret bool
		return ret
	}
	return *o.RetentionGroupByFormat
}

// GetRetentionGroupByFormatOk returns a tuple with the RetentionGroupByFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRetentionRulesRequestPatch) GetRetentionGroupByFormatOk() (*bool, bool) {
	if o == nil || IsNil(o.RetentionGroupByFormat) {
		return nil, false
	}
	return o.RetentionGroupByFormat, true
}

// HasRetentionGroupByFormat returns a boolean if a field has been set.
func (o *RepositoryRetentionRulesRequestPatch) HasRetentionGroupByFormat() bool {
	if o != nil && !IsNil(o.RetentionGroupByFormat) {
		return true
	}

	return false
}

// SetRetentionGroupByFormat gets a reference to the given bool and assigns it to the RetentionGroupByFormat field.
func (o *RepositoryRetentionRulesRequestPatch) SetRetentionGroupByFormat(v bool) {
	o.RetentionGroupByFormat = &v
}

// GetRetentionGroupByName returns the RetentionGroupByName field value if set, zero value otherwise.
func (o *RepositoryRetentionRulesRequestPatch) GetRetentionGroupByName() bool {
	if o == nil || IsNil(o.RetentionGroupByName) {
		var ret bool
		return ret
	}
	return *o.RetentionGroupByName
}

// GetRetentionGroupByNameOk returns a tuple with the RetentionGroupByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRetentionRulesRequestPatch) GetRetentionGroupByNameOk() (*bool, bool) {
	if o == nil || IsNil(o.RetentionGroupByName) {
		return nil, false
	}
	return o.RetentionGroupByName, true
}

// HasRetentionGroupByName returns a boolean if a field has been set.
func (o *RepositoryRetentionRulesRequestPatch) HasRetentionGroupByName() bool {
	if o != nil && !IsNil(o.RetentionGroupByName) {
		return true
	}

	return false
}

// SetRetentionGroupByName gets a reference to the given bool and assigns it to the RetentionGroupByName field.
func (o *RepositoryRetentionRulesRequestPatch) SetRetentionGroupByName(v bool) {
	o.RetentionGroupByName = &v
}

// GetRetentionGroupByPackageType returns the RetentionGroupByPackageType field value if set, zero value otherwise.
func (o *RepositoryRetentionRulesRequestPatch) GetRetentionGroupByPackageType() bool {
	if o == nil || IsNil(o.RetentionGroupByPackageType) {
		var ret bool
		return ret
	}
	return *o.RetentionGroupByPackageType
}

// GetRetentionGroupByPackageTypeOk returns a tuple with the RetentionGroupByPackageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRetentionRulesRequestPatch) GetRetentionGroupByPackageTypeOk() (*bool, bool) {
	if o == nil || IsNil(o.RetentionGroupByPackageType) {
		return nil, false
	}
	return o.RetentionGroupByPackageType, true
}

// HasRetentionGroupByPackageType returns a boolean if a field has been set.
func (o *RepositoryRetentionRulesRequestPatch) HasRetentionGroupByPackageType() bool {
	if o != nil && !IsNil(o.RetentionGroupByPackageType) {
		return true
	}

	return false
}

// SetRetentionGroupByPackageType gets a reference to the given bool and assigns it to the RetentionGroupByPackageType field.
func (o *RepositoryRetentionRulesRequestPatch) SetRetentionGroupByPackageType(v bool) {
	o.RetentionGroupByPackageType = &v
}

// GetRetentionSizeLimit returns the RetentionSizeLimit field value if set, zero value otherwise.
func (o *RepositoryRetentionRulesRequestPatch) GetRetentionSizeLimit() int64 {
	if o == nil || IsNil(o.RetentionSizeLimit) {
		var ret int64
		return ret
	}
	return *o.RetentionSizeLimit
}

// GetRetentionSizeLimitOk returns a tuple with the RetentionSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRetentionRulesRequestPatch) GetRetentionSizeLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.RetentionSizeLimit) {
		return nil, false
	}
	return o.RetentionSizeLimit, true
}

// HasRetentionSizeLimit returns a boolean if a field has been set.
func (o *RepositoryRetentionRulesRequestPatch) HasRetentionSizeLimit() bool {
	if o != nil && !IsNil(o.RetentionSizeLimit) {
		return true
	}

	return false
}

// SetRetentionSizeLimit gets a reference to the given int64 and assigns it to the RetentionSizeLimit field.
func (o *RepositoryRetentionRulesRequestPatch) SetRetentionSizeLimit(v int64) {
	o.RetentionSizeLimit = &v
}

func (o RepositoryRetentionRulesRequestPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryRetentionRulesRequestPatch) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RepositoryRetentionRulesRequestPatch) UnmarshalJSON(data []byte) (err error) {
	varRepositoryRetentionRulesRequestPatch := _RepositoryRetentionRulesRequestPatch{}

	err = json.Unmarshal(data, &varRepositoryRetentionRulesRequestPatch)

	if err != nil {
		return err
	}

	*o = RepositoryRetentionRulesRequestPatch(varRepositoryRetentionRulesRequestPatch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "retention_count_limit")
		delete(additionalProperties, "retention_days_limit")
		delete(additionalProperties, "retention_enabled")
		delete(additionalProperties, "retention_group_by_format")
		delete(additionalProperties, "retention_group_by_name")
		delete(additionalProperties, "retention_group_by_package_type")
		delete(additionalProperties, "retention_size_limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRepositoryRetentionRulesRequestPatch struct {
	value *RepositoryRetentionRulesRequestPatch
	isSet bool
}

func (v NullableRepositoryRetentionRulesRequestPatch) Get() *RepositoryRetentionRulesRequestPatch {
	return v.value
}

func (v *NullableRepositoryRetentionRulesRequestPatch) Set(val *RepositoryRetentionRulesRequestPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryRetentionRulesRequestPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryRetentionRulesRequestPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryRetentionRulesRequestPatch(val *RepositoryRetentionRulesRequestPatch) *NullableRepositoryRetentionRulesRequestPatch {
	return &NullableRepositoryRetentionRulesRequestPatch{value: val, isSet: true}
}

func (v NullableRepositoryRetentionRulesRequestPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryRetentionRulesRequestPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
