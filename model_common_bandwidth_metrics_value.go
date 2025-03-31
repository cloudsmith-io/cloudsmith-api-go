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

// checks if the CommonBandwidthMetricsValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonBandwidthMetricsValue{}

// CommonBandwidthMetricsValue Average bandwidth usage in the specified period, e.g. a day
type CommonBandwidthMetricsValue struct {
	// Bandwidth usage value
	Display string `json:"display"`
	// Unit of measurement e.g. bytes
	Units *string `json:"units,omitempty"`
	// Human readable version of display value
	Value                int64 `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _CommonBandwidthMetricsValue CommonBandwidthMetricsValue

// NewCommonBandwidthMetricsValue instantiates a new CommonBandwidthMetricsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonBandwidthMetricsValue(display string, value int64) *CommonBandwidthMetricsValue {
	this := CommonBandwidthMetricsValue{}
	this.Display = display
	var units string = "bytes"
	this.Units = &units
	this.Value = value
	return &this
}

// NewCommonBandwidthMetricsValueWithDefaults instantiates a new CommonBandwidthMetricsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonBandwidthMetricsValueWithDefaults() *CommonBandwidthMetricsValue {
	this := CommonBandwidthMetricsValue{}
	var units string = "bytes"
	this.Units = &units
	return &this
}

// GetDisplay returns the Display field value
func (o *CommonBandwidthMetricsValue) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *CommonBandwidthMetricsValue) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *CommonBandwidthMetricsValue) SetDisplay(v string) {
	o.Display = v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *CommonBandwidthMetricsValue) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBandwidthMetricsValue) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *CommonBandwidthMetricsValue) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *CommonBandwidthMetricsValue) SetUnits(v string) {
	o.Units = &v
}

// GetValue returns the Value field value
func (o *CommonBandwidthMetricsValue) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CommonBandwidthMetricsValue) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CommonBandwidthMetricsValue) SetValue(v int64) {
	o.Value = v
}

func (o CommonBandwidthMetricsValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonBandwidthMetricsValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["display"] = o.Display
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonBandwidthMetricsValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"display",
		"value",
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

	varCommonBandwidthMetricsValue := _CommonBandwidthMetricsValue{}

	err = json.Unmarshal(data, &varCommonBandwidthMetricsValue)

	if err != nil {
		return err
	}

	*o = CommonBandwidthMetricsValue(varCommonBandwidthMetricsValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "display")
		delete(additionalProperties, "units")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonBandwidthMetricsValue struct {
	value *CommonBandwidthMetricsValue
	isSet bool
}

func (v NullableCommonBandwidthMetricsValue) Get() *CommonBandwidthMetricsValue {
	return v.value
}

func (v *NullableCommonBandwidthMetricsValue) Set(val *CommonBandwidthMetricsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonBandwidthMetricsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonBandwidthMetricsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonBandwidthMetricsValue(val *CommonBandwidthMetricsValue) *NullableCommonBandwidthMetricsValue {
	return &NullableCommonBandwidthMetricsValue{value: val, isSet: true}
}

func (v NullableCommonBandwidthMetricsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonBandwidthMetricsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
