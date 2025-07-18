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

// checks if the RateCheck type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateCheck{}

// RateCheck struct for RateCheck
type RateCheck struct {
	// The time in seconds that you are suggested to wait until the next request in order to avoid consuming too much within the rate limit window.
	Interval *float64 `json:"interval,omitempty"`
	// The maximum number of requests that you are permitted to send per hour
	Limit *int64 `json:"limit,omitempty"`
	// The number of requests that are remaining in the current rate limit window
	Remaining *int64 `json:"remaining,omitempty"`
	// The UTC epoch timestamp at which the current rate limit window will reset
	Reset *int64 `json:"reset,omitempty"`
	// The ISO 8601 datetime at which the current rate limit window will reset
	ResetIso8601 *string `json:"reset_iso_8601,omitempty"`
	// If true, throttling is currently being enforced.
	Throttled            *bool `json:"throttled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RateCheck RateCheck

// NewRateCheck instantiates a new RateCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateCheck() *RateCheck {
	this := RateCheck{}
	return &this
}

// NewRateCheckWithDefaults instantiates a new RateCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateCheckWithDefaults() *RateCheck {
	this := RateCheck{}
	return &this
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *RateCheck) GetInterval() float64 {
	if o == nil || IsNil(o.Interval) {
		var ret float64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateCheck) GetIntervalOk() (*float64, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *RateCheck) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given float64 and assigns it to the Interval field.
func (o *RateCheck) SetInterval(v float64) {
	o.Interval = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *RateCheck) GetLimit() int64 {
	if o == nil || IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateCheck) GetLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *RateCheck) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *RateCheck) SetLimit(v int64) {
	o.Limit = &v
}

// GetRemaining returns the Remaining field value if set, zero value otherwise.
func (o *RateCheck) GetRemaining() int64 {
	if o == nil || IsNil(o.Remaining) {
		var ret int64
		return ret
	}
	return *o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateCheck) GetRemainingOk() (*int64, bool) {
	if o == nil || IsNil(o.Remaining) {
		return nil, false
	}
	return o.Remaining, true
}

// HasRemaining returns a boolean if a field has been set.
func (o *RateCheck) HasRemaining() bool {
	if o != nil && !IsNil(o.Remaining) {
		return true
	}

	return false
}

// SetRemaining gets a reference to the given int64 and assigns it to the Remaining field.
func (o *RateCheck) SetRemaining(v int64) {
	o.Remaining = &v
}

// GetReset returns the Reset field value if set, zero value otherwise.
func (o *RateCheck) GetReset() int64 {
	if o == nil || IsNil(o.Reset) {
		var ret int64
		return ret
	}
	return *o.Reset
}

// GetResetOk returns a tuple with the Reset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateCheck) GetResetOk() (*int64, bool) {
	if o == nil || IsNil(o.Reset) {
		return nil, false
	}
	return o.Reset, true
}

// HasReset returns a boolean if a field has been set.
func (o *RateCheck) HasReset() bool {
	if o != nil && !IsNil(o.Reset) {
		return true
	}

	return false
}

// SetReset gets a reference to the given int64 and assigns it to the Reset field.
func (o *RateCheck) SetReset(v int64) {
	o.Reset = &v
}

// GetResetIso8601 returns the ResetIso8601 field value if set, zero value otherwise.
func (o *RateCheck) GetResetIso8601() string {
	if o == nil || IsNil(o.ResetIso8601) {
		var ret string
		return ret
	}
	return *o.ResetIso8601
}

// GetResetIso8601Ok returns a tuple with the ResetIso8601 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateCheck) GetResetIso8601Ok() (*string, bool) {
	if o == nil || IsNil(o.ResetIso8601) {
		return nil, false
	}
	return o.ResetIso8601, true
}

// HasResetIso8601 returns a boolean if a field has been set.
func (o *RateCheck) HasResetIso8601() bool {
	if o != nil && !IsNil(o.ResetIso8601) {
		return true
	}

	return false
}

// SetResetIso8601 gets a reference to the given string and assigns it to the ResetIso8601 field.
func (o *RateCheck) SetResetIso8601(v string) {
	o.ResetIso8601 = &v
}

// GetThrottled returns the Throttled field value if set, zero value otherwise.
func (o *RateCheck) GetThrottled() bool {
	if o == nil || IsNil(o.Throttled) {
		var ret bool
		return ret
	}
	return *o.Throttled
}

// GetThrottledOk returns a tuple with the Throttled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateCheck) GetThrottledOk() (*bool, bool) {
	if o == nil || IsNil(o.Throttled) {
		return nil, false
	}
	return o.Throttled, true
}

// HasThrottled returns a boolean if a field has been set.
func (o *RateCheck) HasThrottled() bool {
	if o != nil && !IsNil(o.Throttled) {
		return true
	}

	return false
}

// SetThrottled gets a reference to the given bool and assigns it to the Throttled field.
func (o *RateCheck) SetThrottled(v bool) {
	o.Throttled = &v
}

func (o RateCheck) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateCheck) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Remaining) {
		toSerialize["remaining"] = o.Remaining
	}
	if !IsNil(o.Reset) {
		toSerialize["reset"] = o.Reset
	}
	if !IsNil(o.ResetIso8601) {
		toSerialize["reset_iso_8601"] = o.ResetIso8601
	}
	if !IsNil(o.Throttled) {
		toSerialize["throttled"] = o.Throttled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RateCheck) UnmarshalJSON(data []byte) (err error) {
	varRateCheck := _RateCheck{}

	err = json.Unmarshal(data, &varRateCheck)

	if err != nil {
		return err
	}

	*o = RateCheck(varRateCheck)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "interval")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "remaining")
		delete(additionalProperties, "reset")
		delete(additionalProperties, "reset_iso_8601")
		delete(additionalProperties, "throttled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRateCheck struct {
	value *RateCheck
	isSet bool
}

func (v NullableRateCheck) Get() *RateCheck {
	return v.value
}

func (v *NullableRateCheck) Set(val *RateCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableRateCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableRateCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateCheck(val *RateCheck) *NullableRateCheck {
	return &NullableRateCheck{value: val, isSet: true}
}

func (v NullableRateCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
