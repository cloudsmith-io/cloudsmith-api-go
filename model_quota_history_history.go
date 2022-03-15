/*
Cloudsmith API

The API to the Cloudsmith Service

API version: 1.40.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// QuotaHistoryHistory struct for QuotaHistoryHistory
type QuotaHistoryHistory struct {
	//
	Days *int64 `json:"days,omitempty"`
	//
	Display map[string]interface{} `json:"display,omitempty"`
	//
	End *string `json:"end,omitempty"`
	//
	Plan *string `json:"plan,omitempty"`
	//
	Raw map[string]interface{} `json:"raw,omitempty"`
	//
	Start *string `json:"start,omitempty"`
}

// NewQuotaHistoryHistory instantiates a new QuotaHistoryHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaHistoryHistory() *QuotaHistoryHistory {
	this := QuotaHistoryHistory{}
	return &this
}

// NewQuotaHistoryHistoryWithDefaults instantiates a new QuotaHistoryHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaHistoryHistoryWithDefaults() *QuotaHistoryHistory {
	this := QuotaHistoryHistory{}
	return &this
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *QuotaHistoryHistory) GetDays() int64 {
	if o == nil || o.Days == nil {
		var ret int64
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaHistoryHistory) GetDaysOk() (*int64, bool) {
	if o == nil || o.Days == nil {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *QuotaHistoryHistory) HasDays() bool {
	if o != nil && o.Days != nil {
		return true
	}

	return false
}

// SetDays gets a reference to the given int64 and assigns it to the Days field.
func (o *QuotaHistoryHistory) SetDays(v int64) {
	o.Days = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *QuotaHistoryHistory) GetDisplay() map[string]interface{} {
	if o == nil || o.Display == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaHistoryHistory) GetDisplayOk() (map[string]interface{}, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *QuotaHistoryHistory) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given map[string]interface{} and assigns it to the Display field.
func (o *QuotaHistoryHistory) SetDisplay(v map[string]interface{}) {
	o.Display = v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *QuotaHistoryHistory) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaHistoryHistory) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *QuotaHistoryHistory) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *QuotaHistoryHistory) SetEnd(v string) {
	o.End = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *QuotaHistoryHistory) GetPlan() string {
	if o == nil || o.Plan == nil {
		var ret string
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaHistoryHistory) GetPlanOk() (*string, bool) {
	if o == nil || o.Plan == nil {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *QuotaHistoryHistory) HasPlan() bool {
	if o != nil && o.Plan != nil {
		return true
	}

	return false
}

// SetPlan gets a reference to the given string and assigns it to the Plan field.
func (o *QuotaHistoryHistory) SetPlan(v string) {
	o.Plan = &v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *QuotaHistoryHistory) GetRaw() map[string]interface{} {
	if o == nil || o.Raw == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaHistoryHistory) GetRawOk() (map[string]interface{}, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *QuotaHistoryHistory) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given map[string]interface{} and assigns it to the Raw field.
func (o *QuotaHistoryHistory) SetRaw(v map[string]interface{}) {
	o.Raw = v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *QuotaHistoryHistory) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaHistoryHistory) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *QuotaHistoryHistory) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *QuotaHistoryHistory) SetStart(v string) {
	o.Start = &v
}

func (o QuotaHistoryHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Days != nil {
		toSerialize["days"] = o.Days
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Plan != nil {
		toSerialize["plan"] = o.Plan
	}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	return json.Marshal(toSerialize)
}

type NullableQuotaHistoryHistory struct {
	value *QuotaHistoryHistory
	isSet bool
}

func (v NullableQuotaHistoryHistory) Get() *QuotaHistoryHistory {
	return v.value
}

func (v *NullableQuotaHistoryHistory) Set(val *QuotaHistoryHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaHistoryHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaHistoryHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaHistoryHistory(val *QuotaHistoryHistory) *NullableQuotaHistoryHistory {
	return &NullableQuotaHistoryHistory{value: val, isSet: true}
}

func (v NullableQuotaHistoryHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaHistoryHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
