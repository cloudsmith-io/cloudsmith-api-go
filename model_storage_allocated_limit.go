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
)

// checks if the StorageAllocatedLimit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageAllocatedLimit{}

// StorageAllocatedLimit struct for StorageAllocatedLimit
type StorageAllocatedLimit struct {
	Configured           *string `json:"configured,omitempty"`
	Peak                 *string `json:"peak,omitempty"`
	PercentageUsed       *string `json:"percentage_used,omitempty"`
	PlanLimit            *string `json:"plan_limit,omitempty"`
	Used                 *string `json:"used,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageAllocatedLimit StorageAllocatedLimit

// NewStorageAllocatedLimit instantiates a new StorageAllocatedLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageAllocatedLimit() *StorageAllocatedLimit {
	this := StorageAllocatedLimit{}
	return &this
}

// NewStorageAllocatedLimitWithDefaults instantiates a new StorageAllocatedLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageAllocatedLimitWithDefaults() *StorageAllocatedLimit {
	this := StorageAllocatedLimit{}
	return &this
}

// GetConfigured returns the Configured field value if set, zero value otherwise.
func (o *StorageAllocatedLimit) GetConfigured() string {
	if o == nil || IsNil(o.Configured) {
		var ret string
		return ret
	}
	return *o.Configured
}

// GetConfiguredOk returns a tuple with the Configured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAllocatedLimit) GetConfiguredOk() (*string, bool) {
	if o == nil || IsNil(o.Configured) {
		return nil, false
	}
	return o.Configured, true
}

// HasConfigured returns a boolean if a field has been set.
func (o *StorageAllocatedLimit) HasConfigured() bool {
	if o != nil && !IsNil(o.Configured) {
		return true
	}

	return false
}

// SetConfigured gets a reference to the given string and assigns it to the Configured field.
func (o *StorageAllocatedLimit) SetConfigured(v string) {
	o.Configured = &v
}

// GetPeak returns the Peak field value if set, zero value otherwise.
func (o *StorageAllocatedLimit) GetPeak() string {
	if o == nil || IsNil(o.Peak) {
		var ret string
		return ret
	}
	return *o.Peak
}

// GetPeakOk returns a tuple with the Peak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAllocatedLimit) GetPeakOk() (*string, bool) {
	if o == nil || IsNil(o.Peak) {
		return nil, false
	}
	return o.Peak, true
}

// HasPeak returns a boolean if a field has been set.
func (o *StorageAllocatedLimit) HasPeak() bool {
	if o != nil && !IsNil(o.Peak) {
		return true
	}

	return false
}

// SetPeak gets a reference to the given string and assigns it to the Peak field.
func (o *StorageAllocatedLimit) SetPeak(v string) {
	o.Peak = &v
}

// GetPercentageUsed returns the PercentageUsed field value if set, zero value otherwise.
func (o *StorageAllocatedLimit) GetPercentageUsed() string {
	if o == nil || IsNil(o.PercentageUsed) {
		var ret string
		return ret
	}
	return *o.PercentageUsed
}

// GetPercentageUsedOk returns a tuple with the PercentageUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAllocatedLimit) GetPercentageUsedOk() (*string, bool) {
	if o == nil || IsNil(o.PercentageUsed) {
		return nil, false
	}
	return o.PercentageUsed, true
}

// HasPercentageUsed returns a boolean if a field has been set.
func (o *StorageAllocatedLimit) HasPercentageUsed() bool {
	if o != nil && !IsNil(o.PercentageUsed) {
		return true
	}

	return false
}

// SetPercentageUsed gets a reference to the given string and assigns it to the PercentageUsed field.
func (o *StorageAllocatedLimit) SetPercentageUsed(v string) {
	o.PercentageUsed = &v
}

// GetPlanLimit returns the PlanLimit field value if set, zero value otherwise.
func (o *StorageAllocatedLimit) GetPlanLimit() string {
	if o == nil || IsNil(o.PlanLimit) {
		var ret string
		return ret
	}
	return *o.PlanLimit
}

// GetPlanLimitOk returns a tuple with the PlanLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAllocatedLimit) GetPlanLimitOk() (*string, bool) {
	if o == nil || IsNil(o.PlanLimit) {
		return nil, false
	}
	return o.PlanLimit, true
}

// HasPlanLimit returns a boolean if a field has been set.
func (o *StorageAllocatedLimit) HasPlanLimit() bool {
	if o != nil && !IsNil(o.PlanLimit) {
		return true
	}

	return false
}

// SetPlanLimit gets a reference to the given string and assigns it to the PlanLimit field.
func (o *StorageAllocatedLimit) SetPlanLimit(v string) {
	o.PlanLimit = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *StorageAllocatedLimit) GetUsed() string {
	if o == nil || IsNil(o.Used) {
		var ret string
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAllocatedLimit) GetUsedOk() (*string, bool) {
	if o == nil || IsNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *StorageAllocatedLimit) HasUsed() bool {
	if o != nil && !IsNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given string and assigns it to the Used field.
func (o *StorageAllocatedLimit) SetUsed(v string) {
	o.Used = &v
}

func (o StorageAllocatedLimit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageAllocatedLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configured) {
		toSerialize["configured"] = o.Configured
	}
	if !IsNil(o.Peak) {
		toSerialize["peak"] = o.Peak
	}
	if !IsNil(o.PercentageUsed) {
		toSerialize["percentage_used"] = o.PercentageUsed
	}
	if !IsNil(o.PlanLimit) {
		toSerialize["plan_limit"] = o.PlanLimit
	}
	if !IsNil(o.Used) {
		toSerialize["used"] = o.Used
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageAllocatedLimit) UnmarshalJSON(data []byte) (err error) {
	varStorageAllocatedLimit := _StorageAllocatedLimit{}

	err = json.Unmarshal(data, &varStorageAllocatedLimit)

	if err != nil {
		return err
	}

	*o = StorageAllocatedLimit(varStorageAllocatedLimit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "configured")
		delete(additionalProperties, "peak")
		delete(additionalProperties, "percentage_used")
		delete(additionalProperties, "plan_limit")
		delete(additionalProperties, "used")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageAllocatedLimit struct {
	value *StorageAllocatedLimit
	isSet bool
}

func (v NullableStorageAllocatedLimit) Get() *StorageAllocatedLimit {
	return v.value
}

func (v *NullableStorageAllocatedLimit) Set(val *StorageAllocatedLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageAllocatedLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageAllocatedLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageAllocatedLimit(val *StorageAllocatedLimit) *NullableStorageAllocatedLimit {
	return &NullableStorageAllocatedLimit{value: val, isSet: true}
}

func (v NullableStorageAllocatedLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageAllocatedLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
