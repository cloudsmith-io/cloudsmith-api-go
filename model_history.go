/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.617.2
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the History type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &History{}

// History struct for History
type History struct {
	Days                 *int64             `json:"days,omitempty"`
	Display              HistoryFieldset    `json:"display"`
	End                  time.Time          `json:"end"`
	Plan                 string             `json:"plan"`
	Raw                  HistoryFieldsetRaw `json:"raw"`
	Start                time.Time          `json:"start"`
	AdditionalProperties map[string]interface{}
}

type _History History

// NewHistory instantiates a new History object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistory(display HistoryFieldset, end time.Time, plan string, raw HistoryFieldsetRaw, start time.Time) *History {
	this := History{}
	this.Display = display
	this.End = end
	this.Plan = plan
	this.Raw = raw
	this.Start = start
	return &this
}

// NewHistoryWithDefaults instantiates a new History object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryWithDefaults() *History {
	this := History{}
	return &this
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *History) GetDays() int64 {
	if o == nil || IsNil(o.Days) {
		var ret int64
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *History) GetDaysOk() (*int64, bool) {
	if o == nil || IsNil(o.Days) {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *History) HasDays() bool {
	if o != nil && !IsNil(o.Days) {
		return true
	}

	return false
}

// SetDays gets a reference to the given int64 and assigns it to the Days field.
func (o *History) SetDays(v int64) {
	o.Days = &v
}

// GetDisplay returns the Display field value
func (o *History) GetDisplay() HistoryFieldset {
	if o == nil {
		var ret HistoryFieldset
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *History) GetDisplayOk() (*HistoryFieldset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *History) SetDisplay(v HistoryFieldset) {
	o.Display = v
}

// GetEnd returns the End field value
func (o *History) GetEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *History) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *History) SetEnd(v time.Time) {
	o.End = v
}

// GetPlan returns the Plan field value
func (o *History) GetPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *History) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *History) SetPlan(v string) {
	o.Plan = v
}

// GetRaw returns the Raw field value
func (o *History) GetRaw() HistoryFieldsetRaw {
	if o == nil {
		var ret HistoryFieldsetRaw
		return ret
	}

	return o.Raw
}

// GetRawOk returns a tuple with the Raw field value
// and a boolean to check if the value has been set.
func (o *History) GetRawOk() (*HistoryFieldsetRaw, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Raw, true
}

// SetRaw sets field value
func (o *History) SetRaw(v HistoryFieldsetRaw) {
	o.Raw = v
}

// GetStart returns the Start field value
func (o *History) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *History) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *History) SetStart(v time.Time) {
	o.Start = v
}

func (o History) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o History) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Days) {
		toSerialize["days"] = o.Days
	}
	toSerialize["display"] = o.Display
	toSerialize["end"] = o.End
	toSerialize["plan"] = o.Plan
	toSerialize["raw"] = o.Raw
	toSerialize["start"] = o.Start

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *History) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"display",
		"end",
		"plan",
		"raw",
		"start",
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

	varHistory := _History{}

	err = json.Unmarshal(data, &varHistory)

	if err != nil {
		return err
	}

	*o = History(varHistory)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "days")
		delete(additionalProperties, "display")
		delete(additionalProperties, "end")
		delete(additionalProperties, "plan")
		delete(additionalProperties, "raw")
		delete(additionalProperties, "start")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHistory struct {
	value *History
	isSet bool
}

func (v NullableHistory) Get() *History {
	return v.value
}

func (v *NullableHistory) Set(val *History) {
	v.value = val
	v.isSet = true
}

func (v NullableHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistory(val *History) *NullableHistory {
	return &NullableHistory{value: val, isSet: true}
}

func (v NullableHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
