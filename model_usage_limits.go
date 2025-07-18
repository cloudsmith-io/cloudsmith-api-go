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
	"fmt"
)

// checks if the UsageLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageLimits{}

// UsageLimits struct for UsageLimits
type UsageLimits struct {
	Bandwidth            AllocatedLimit        `json:"bandwidth"`
	Storage              StorageAllocatedLimit `json:"storage"`
	AdditionalProperties map[string]interface{}
}

type _UsageLimits UsageLimits

// NewUsageLimits instantiates a new UsageLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageLimits(bandwidth AllocatedLimit, storage StorageAllocatedLimit) *UsageLimits {
	this := UsageLimits{}
	this.Bandwidth = bandwidth
	this.Storage = storage
	return &this
}

// NewUsageLimitsWithDefaults instantiates a new UsageLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageLimitsWithDefaults() *UsageLimits {
	this := UsageLimits{}
	return &this
}

// GetBandwidth returns the Bandwidth field value
func (o *UsageLimits) GetBandwidth() AllocatedLimit {
	if o == nil {
		var ret AllocatedLimit
		return ret
	}

	return o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value
// and a boolean to check if the value has been set.
func (o *UsageLimits) GetBandwidthOk() (*AllocatedLimit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bandwidth, true
}

// SetBandwidth sets field value
func (o *UsageLimits) SetBandwidth(v AllocatedLimit) {
	o.Bandwidth = v
}

// GetStorage returns the Storage field value
func (o *UsageLimits) GetStorage() StorageAllocatedLimit {
	if o == nil {
		var ret StorageAllocatedLimit
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *UsageLimits) GetStorageOk() (*StorageAllocatedLimit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *UsageLimits) SetStorage(v StorageAllocatedLimit) {
	o.Storage = v
}

func (o UsageLimits) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bandwidth"] = o.Bandwidth
	toSerialize["storage"] = o.Storage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UsageLimits) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bandwidth",
		"storage",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUsageLimits := _UsageLimits{}

	err = json.Unmarshal(data, &varUsageLimits)

	if err != nil {
		return err
	}

	*o = UsageLimits(varUsageLimits)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bandwidth")
		delete(additionalProperties, "storage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsageLimits struct {
	value *UsageLimits
	isSet bool
}

func (v NullableUsageLimits) Get() *UsageLimits {
	return v.value
}

func (v *NullableUsageLimits) Set(val *UsageLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageLimits(val *UsageLimits) *NullableUsageLimits {
	return &NullableUsageLimits{value: val, isSet: true}
}

func (v NullableUsageLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
