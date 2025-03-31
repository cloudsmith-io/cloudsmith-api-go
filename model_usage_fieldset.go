/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.654.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"fmt"
)

// checks if the UsageFieldset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageFieldset{}

// UsageFieldset struct for UsageFieldset
type UsageFieldset struct {
	Display              UsageLimits    `json:"display"`
	Raw                  UsageLimitsRaw `json:"raw"`
	AdditionalProperties map[string]interface{}
}

type _UsageFieldset UsageFieldset

// NewUsageFieldset instantiates a new UsageFieldset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageFieldset(display UsageLimits, raw UsageLimitsRaw) *UsageFieldset {
	this := UsageFieldset{}
	this.Display = display
	this.Raw = raw
	return &this
}

// NewUsageFieldsetWithDefaults instantiates a new UsageFieldset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageFieldsetWithDefaults() *UsageFieldset {
	this := UsageFieldset{}
	return &this
}

// GetDisplay returns the Display field value
func (o *UsageFieldset) GetDisplay() UsageLimits {
	if o == nil {
		var ret UsageLimits
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *UsageFieldset) GetDisplayOk() (*UsageLimits, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *UsageFieldset) SetDisplay(v UsageLimits) {
	o.Display = v
}

// GetRaw returns the Raw field value
func (o *UsageFieldset) GetRaw() UsageLimitsRaw {
	if o == nil {
		var ret UsageLimitsRaw
		return ret
	}

	return o.Raw
}

// GetRawOk returns a tuple with the Raw field value
// and a boolean to check if the value has been set.
func (o *UsageFieldset) GetRawOk() (*UsageLimitsRaw, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Raw, true
}

// SetRaw sets field value
func (o *UsageFieldset) SetRaw(v UsageLimitsRaw) {
	o.Raw = v
}

func (o UsageFieldset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageFieldset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["display"] = o.Display
	toSerialize["raw"] = o.Raw

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UsageFieldset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"display",
		"raw",
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

	varUsageFieldset := _UsageFieldset{}

	err = json.Unmarshal(data, &varUsageFieldset)

	if err != nil {
		return err
	}

	*o = UsageFieldset(varUsageFieldset)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "display")
		delete(additionalProperties, "raw")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsageFieldset struct {
	value *UsageFieldset
	isSet bool
}

func (v NullableUsageFieldset) Get() *UsageFieldset {
	return v.value
}

func (v *NullableUsageFieldset) Set(val *UsageFieldset) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageFieldset) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageFieldset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageFieldset(val *UsageFieldset) *NullableUsageFieldset {
	return &NullableUsageFieldset{value: val, isSet: true}
}

func (v NullableUsageFieldset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageFieldset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
