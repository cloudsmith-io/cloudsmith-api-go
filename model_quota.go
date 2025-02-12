/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.616.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"fmt"
)

// checks if the Quota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Quota{}

// Quota struct for Quota
type Quota struct {
	Usage                UsageFieldset `json:"usage"`
	AdditionalProperties map[string]interface{}
}

type _Quota Quota

// NewQuota instantiates a new Quota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuota(usage UsageFieldset) *Quota {
	this := Quota{}
	this.Usage = usage
	return &this
}

// NewQuotaWithDefaults instantiates a new Quota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaWithDefaults() *Quota {
	this := Quota{}
	return &this
}

// GetUsage returns the Usage field value
func (o *Quota) GetUsage() UsageFieldset {
	if o == nil {
		var ret UsageFieldset
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *Quota) GetUsageOk() (*UsageFieldset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *Quota) SetUsage(v UsageFieldset) {
	o.Usage = v
}

func (o Quota) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Quota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["usage"] = o.Usage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Quota) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"usage",
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

	varQuota := _Quota{}

	err = json.Unmarshal(data, &varQuota)

	if err != nil {
		return err
	}

	*o = Quota(varQuota)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "usage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuota struct {
	value *Quota
	isSet bool
}

func (v NullableQuota) Get() *Quota {
	return v.value
}

func (v *NullableQuota) Set(val *Quota) {
	v.value = val
	v.isSet = true
}

func (v NullableQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuota(val *Quota) *NullableQuota {
	return &NullableQuota{value: val, isSet: true}
}

func (v NullableQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
