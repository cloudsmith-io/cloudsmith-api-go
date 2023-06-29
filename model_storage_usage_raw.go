/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.290.2
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// StorageUsageRaw struct for StorageUsageRaw
type StorageUsageRaw struct {
	Limit      *int64   `json:"limit,omitempty"`
	Peak       *int64   `json:"peak,omitempty"`
	Percentage *float64 `json:"percentage,omitempty"`
	Used       *int64   `json:"used,omitempty"`
}

// NewStorageUsageRaw instantiates a new StorageUsageRaw object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageUsageRaw() *StorageUsageRaw {
	this := StorageUsageRaw{}
	return &this
}

// NewStorageUsageRawWithDefaults instantiates a new StorageUsageRaw object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageUsageRawWithDefaults() *StorageUsageRaw {
	this := StorageUsageRaw{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *StorageUsageRaw) GetLimit() int64 {
	if o == nil || isNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUsageRaw) GetLimitOk() (*int64, bool) {
	if o == nil || isNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *StorageUsageRaw) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *StorageUsageRaw) SetLimit(v int64) {
	o.Limit = &v
}

// GetPeak returns the Peak field value if set, zero value otherwise.
func (o *StorageUsageRaw) GetPeak() int64 {
	if o == nil || isNil(o.Peak) {
		var ret int64
		return ret
	}
	return *o.Peak
}

// GetPeakOk returns a tuple with the Peak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUsageRaw) GetPeakOk() (*int64, bool) {
	if o == nil || isNil(o.Peak) {
		return nil, false
	}
	return o.Peak, true
}

// HasPeak returns a boolean if a field has been set.
func (o *StorageUsageRaw) HasPeak() bool {
	if o != nil && !isNil(o.Peak) {
		return true
	}

	return false
}

// SetPeak gets a reference to the given int64 and assigns it to the Peak field.
func (o *StorageUsageRaw) SetPeak(v int64) {
	o.Peak = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *StorageUsageRaw) GetPercentage() float64 {
	if o == nil || isNil(o.Percentage) {
		var ret float64
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUsageRaw) GetPercentageOk() (*float64, bool) {
	if o == nil || isNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *StorageUsageRaw) HasPercentage() bool {
	if o != nil && !isNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float64 and assigns it to the Percentage field.
func (o *StorageUsageRaw) SetPercentage(v float64) {
	o.Percentage = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *StorageUsageRaw) GetUsed() int64 {
	if o == nil || isNil(o.Used) {
		var ret int64
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUsageRaw) GetUsedOk() (*int64, bool) {
	if o == nil || isNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *StorageUsageRaw) HasUsed() bool {
	if o != nil && !isNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int64 and assigns it to the Used field.
func (o *StorageUsageRaw) SetUsed(v int64) {
	o.Used = &v
}

func (o StorageUsageRaw) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !isNil(o.Peak) {
		toSerialize["peak"] = o.Peak
	}
	if !isNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	if !isNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	return json.Marshal(toSerialize)
}

type NullableStorageUsageRaw struct {
	value *StorageUsageRaw
	isSet bool
}

func (v NullableStorageUsageRaw) Get() *StorageUsageRaw {
	return v.value
}

func (v *NullableStorageUsageRaw) Set(val *StorageUsageRaw) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageUsageRaw) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageUsageRaw) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageUsageRaw(val *StorageUsageRaw) *NullableStorageUsageRaw {
	return &NullableStorageUsageRaw{value: val, isSet: true}
}

func (v NullableStorageUsageRaw) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageUsageRaw) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
