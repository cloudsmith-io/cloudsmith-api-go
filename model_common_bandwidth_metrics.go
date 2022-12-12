/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.182.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// CommonBandwidthMetrics struct for CommonBandwidthMetrics
type CommonBandwidthMetrics struct {
	Average CommonBandwidthMetricsValue `json:"average"`
	Highest CommonBandwidthMetricsValue `json:"highest"`
	Lowest  CommonBandwidthMetricsValue `json:"lowest"`
	Total   CommonBandwidthMetricsValue `json:"total"`
}

// NewCommonBandwidthMetrics instantiates a new CommonBandwidthMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonBandwidthMetrics(average CommonBandwidthMetricsValue, highest CommonBandwidthMetricsValue, lowest CommonBandwidthMetricsValue, total CommonBandwidthMetricsValue) *CommonBandwidthMetrics {
	this := CommonBandwidthMetrics{}
	this.Average = average
	this.Highest = highest
	this.Lowest = lowest
	this.Total = total
	return &this
}

// NewCommonBandwidthMetricsWithDefaults instantiates a new CommonBandwidthMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonBandwidthMetricsWithDefaults() *CommonBandwidthMetrics {
	this := CommonBandwidthMetrics{}
	return &this
}

// GetAverage returns the Average field value
func (o *CommonBandwidthMetrics) GetAverage() CommonBandwidthMetricsValue {
	if o == nil {
		var ret CommonBandwidthMetricsValue
		return ret
	}

	return o.Average
}

// GetAverageOk returns a tuple with the Average field value
// and a boolean to check if the value has been set.
func (o *CommonBandwidthMetrics) GetAverageOk() (*CommonBandwidthMetricsValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Average, true
}

// SetAverage sets field value
func (o *CommonBandwidthMetrics) SetAverage(v CommonBandwidthMetricsValue) {
	o.Average = v
}

// GetHighest returns the Highest field value
func (o *CommonBandwidthMetrics) GetHighest() CommonBandwidthMetricsValue {
	if o == nil {
		var ret CommonBandwidthMetricsValue
		return ret
	}

	return o.Highest
}

// GetHighestOk returns a tuple with the Highest field value
// and a boolean to check if the value has been set.
func (o *CommonBandwidthMetrics) GetHighestOk() (*CommonBandwidthMetricsValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Highest, true
}

// SetHighest sets field value
func (o *CommonBandwidthMetrics) SetHighest(v CommonBandwidthMetricsValue) {
	o.Highest = v
}

// GetLowest returns the Lowest field value
func (o *CommonBandwidthMetrics) GetLowest() CommonBandwidthMetricsValue {
	if o == nil {
		var ret CommonBandwidthMetricsValue
		return ret
	}

	return o.Lowest
}

// GetLowestOk returns a tuple with the Lowest field value
// and a boolean to check if the value has been set.
func (o *CommonBandwidthMetrics) GetLowestOk() (*CommonBandwidthMetricsValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lowest, true
}

// SetLowest sets field value
func (o *CommonBandwidthMetrics) SetLowest(v CommonBandwidthMetricsValue) {
	o.Lowest = v
}

// GetTotal returns the Total field value
func (o *CommonBandwidthMetrics) GetTotal() CommonBandwidthMetricsValue {
	if o == nil {
		var ret CommonBandwidthMetricsValue
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CommonBandwidthMetrics) GetTotalOk() (*CommonBandwidthMetricsValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CommonBandwidthMetrics) SetTotal(v CommonBandwidthMetricsValue) {
	o.Total = v
}

func (o CommonBandwidthMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["average"] = o.Average
	}
	if true {
		toSerialize["highest"] = o.Highest
	}
	if true {
		toSerialize["lowest"] = o.Lowest
	}
	if true {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableCommonBandwidthMetrics struct {
	value *CommonBandwidthMetrics
	isSet bool
}

func (v NullableCommonBandwidthMetrics) Get() *CommonBandwidthMetrics {
	return v.value
}

func (v *NullableCommonBandwidthMetrics) Set(val *CommonBandwidthMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonBandwidthMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonBandwidthMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonBandwidthMetrics(val *CommonBandwidthMetrics) *NullableCommonBandwidthMetrics {
	return &NullableCommonBandwidthMetrics{value: val, isSet: true}
}

func (v NullableCommonBandwidthMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonBandwidthMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
