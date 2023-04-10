/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.237.2
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// EntitlementUsageMetrics struct for EntitlementUsageMetrics
type EntitlementUsageMetrics struct {
	Tokens CommonMetrics `json:"tokens"`
}

// NewEntitlementUsageMetrics instantiates a new EntitlementUsageMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementUsageMetrics(tokens CommonMetrics) *EntitlementUsageMetrics {
	this := EntitlementUsageMetrics{}
	this.Tokens = tokens
	return &this
}

// NewEntitlementUsageMetricsWithDefaults instantiates a new EntitlementUsageMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementUsageMetricsWithDefaults() *EntitlementUsageMetrics {
	this := EntitlementUsageMetrics{}
	return &this
}

// GetTokens returns the Tokens field value
func (o *EntitlementUsageMetrics) GetTokens() CommonMetrics {
	if o == nil {
		var ret CommonMetrics
		return ret
	}

	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value
// and a boolean to check if the value has been set.
func (o *EntitlementUsageMetrics) GetTokensOk() (*CommonMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tokens, true
}

// SetTokens sets field value
func (o *EntitlementUsageMetrics) SetTokens(v CommonMetrics) {
	o.Tokens = v
}

func (o EntitlementUsageMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tokens"] = o.Tokens
	}
	return json.Marshal(toSerialize)
}

type NullableEntitlementUsageMetrics struct {
	value *EntitlementUsageMetrics
	isSet bool
}

func (v NullableEntitlementUsageMetrics) Get() *EntitlementUsageMetrics {
	return v.value
}

func (v *NullableEntitlementUsageMetrics) Set(val *EntitlementUsageMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementUsageMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementUsageMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementUsageMetrics(val *EntitlementUsageMetrics) *NullableEntitlementUsageMetrics {
	return &NullableEntitlementUsageMetrics{value: val, isSet: true}
}

func (v NullableEntitlementUsageMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementUsageMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
