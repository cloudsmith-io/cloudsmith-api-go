/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.236.5
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// StorageAllocatedLimitRaw struct for StorageAllocatedLimitRaw
type StorageAllocatedLimitRaw struct {
	Configured     *int64   `json:"configured,omitempty"`
	Peak           *int64   `json:"peak,omitempty"`
	PercentageUsed *float64 `json:"percentage_used,omitempty"`
	PlanLimit      *int64   `json:"plan_limit,omitempty"`
	Used           *int64   `json:"used,omitempty"`
}

// NewStorageAllocatedLimitRaw instantiates a new StorageAllocatedLimitRaw object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageAllocatedLimitRaw() *StorageAllocatedLimitRaw {
	this := StorageAllocatedLimitRaw{}
	return &this
}

// NewStorageAllocatedLimitRawWithDefaults instantiates a new StorageAllocatedLimitRaw object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageAllocatedLimitRawWithDefaults() *StorageAllocatedLimitRaw {
	this := StorageAllocatedLimitRaw{}
	return &this
}

// GetConfigured returns the Configured field value if set, zero value otherwise.
func (o *StorageAllocatedLimitRaw) GetConfigured() int64 {
	if o == nil || isNil(o.Configured) {
		var ret int64
		return ret
	}
	return *o.Configured
}

// GetConfiguredOk returns a tuple with the Configured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAllocatedLimitRaw) GetConfiguredOk() (*int64, bool) {
	if o == nil || isNil(o.Configured) {
		return nil, false
	}
	return o.Configured, true
}

// HasConfigured returns a boolean if a field has been set.
func (o *StorageAllocatedLimitRaw) HasConfigured() bool {
	if o != nil && !isNil(o.Configured) {
		return true
	}

	return false
}

// SetConfigured gets a reference to the given int64 and assigns it to the Configured field.
func (o *StorageAllocatedLimitRaw) SetConfigured(v int64) {
	o.Configured = &v
}

// GetPeak returns the Peak field value if set, zero value otherwise.
func (o *StorageAllocatedLimitRaw) GetPeak() int64 {
	if o == nil || isNil(o.Peak) {
		var ret int64
		return ret
	}
	return *o.Peak
}

// GetPeakOk returns a tuple with the Peak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAllocatedLimitRaw) GetPeakOk() (*int64, bool) {
	if o == nil || isNil(o.Peak) {
		return nil, false
	}
	return o.Peak, true
}

// HasPeak returns a boolean if a field has been set.
func (o *StorageAllocatedLimitRaw) HasPeak() bool {
	if o != nil && !isNil(o.Peak) {
		return true
	}

	return false
}

// SetPeak gets a reference to the given int64 and assigns it to the Peak field.
func (o *StorageAllocatedLimitRaw) SetPeak(v int64) {
	o.Peak = &v
}

// GetPercentageUsed returns the PercentageUsed field value if set, zero value otherwise.
func (o *StorageAllocatedLimitRaw) GetPercentageUsed() float64 {
	if o == nil || isNil(o.PercentageUsed) {
		var ret float64
		return ret
	}
	return *o.PercentageUsed
}

// GetPercentageUsedOk returns a tuple with the PercentageUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAllocatedLimitRaw) GetPercentageUsedOk() (*float64, bool) {
	if o == nil || isNil(o.PercentageUsed) {
		return nil, false
	}
	return o.PercentageUsed, true
}

// HasPercentageUsed returns a boolean if a field has been set.
func (o *StorageAllocatedLimitRaw) HasPercentageUsed() bool {
	if o != nil && !isNil(o.PercentageUsed) {
		return true
	}

	return false
}

// SetPercentageUsed gets a reference to the given float64 and assigns it to the PercentageUsed field.
func (o *StorageAllocatedLimitRaw) SetPercentageUsed(v float64) {
	o.PercentageUsed = &v
}

// GetPlanLimit returns the PlanLimit field value if set, zero value otherwise.
func (o *StorageAllocatedLimitRaw) GetPlanLimit() int64 {
	if o == nil || isNil(o.PlanLimit) {
		var ret int64
		return ret
	}
	return *o.PlanLimit
}

// GetPlanLimitOk returns a tuple with the PlanLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAllocatedLimitRaw) GetPlanLimitOk() (*int64, bool) {
	if o == nil || isNil(o.PlanLimit) {
		return nil, false
	}
	return o.PlanLimit, true
}

// HasPlanLimit returns a boolean if a field has been set.
func (o *StorageAllocatedLimitRaw) HasPlanLimit() bool {
	if o != nil && !isNil(o.PlanLimit) {
		return true
	}

	return false
}

// SetPlanLimit gets a reference to the given int64 and assigns it to the PlanLimit field.
func (o *StorageAllocatedLimitRaw) SetPlanLimit(v int64) {
	o.PlanLimit = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *StorageAllocatedLimitRaw) GetUsed() int64 {
	if o == nil || isNil(o.Used) {
		var ret int64
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAllocatedLimitRaw) GetUsedOk() (*int64, bool) {
	if o == nil || isNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *StorageAllocatedLimitRaw) HasUsed() bool {
	if o != nil && !isNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int64 and assigns it to the Used field.
func (o *StorageAllocatedLimitRaw) SetUsed(v int64) {
	o.Used = &v
}

func (o StorageAllocatedLimitRaw) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Configured) {
		toSerialize["configured"] = o.Configured
	}
	if !isNil(o.Peak) {
		toSerialize["peak"] = o.Peak
	}
	if !isNil(o.PercentageUsed) {
		toSerialize["percentage_used"] = o.PercentageUsed
	}
	if !isNil(o.PlanLimit) {
		toSerialize["plan_limit"] = o.PlanLimit
	}
	if !isNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	return json.Marshal(toSerialize)
}

type NullableStorageAllocatedLimitRaw struct {
	value *StorageAllocatedLimitRaw
	isSet bool
}

func (v NullableStorageAllocatedLimitRaw) Get() *StorageAllocatedLimitRaw {
	return v.value
}

func (v *NullableStorageAllocatedLimitRaw) Set(val *StorageAllocatedLimitRaw) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageAllocatedLimitRaw) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageAllocatedLimitRaw) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageAllocatedLimitRaw(val *StorageAllocatedLimitRaw) *NullableStorageAllocatedLimitRaw {
	return &NullableStorageAllocatedLimitRaw{value: val, isSet: true}
}

func (v NullableStorageAllocatedLimitRaw) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageAllocatedLimitRaw) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}