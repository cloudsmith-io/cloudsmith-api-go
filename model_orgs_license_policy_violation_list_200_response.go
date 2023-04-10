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

// OrgsLicensePolicyViolationList200Response struct for OrgsLicensePolicyViolationList200Response
type OrgsLicensePolicyViolationList200Response struct {
	Next     NullableString                     `json:"next,omitempty"`
	Previous NullableString                     `json:"previous,omitempty"`
	Results  []PackageLicensePolicyViolationLog `json:"results"`
}

// NewOrgsLicensePolicyViolationList200Response instantiates a new OrgsLicensePolicyViolationList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgsLicensePolicyViolationList200Response(results []PackageLicensePolicyViolationLog) *OrgsLicensePolicyViolationList200Response {
	this := OrgsLicensePolicyViolationList200Response{}
	this.Results = results
	return &this
}

// NewOrgsLicensePolicyViolationList200ResponseWithDefaults instantiates a new OrgsLicensePolicyViolationList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgsLicensePolicyViolationList200ResponseWithDefaults() *OrgsLicensePolicyViolationList200Response {
	this := OrgsLicensePolicyViolationList200Response{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgsLicensePolicyViolationList200Response) GetNext() string {
	if o == nil || isNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgsLicensePolicyViolationList200Response) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *OrgsLicensePolicyViolationList200Response) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *OrgsLicensePolicyViolationList200Response) SetNext(v string) {
	o.Next.Set(&v)
}

// SetNextNil sets the value for Next to be an explicit nil
func (o *OrgsLicensePolicyViolationList200Response) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *OrgsLicensePolicyViolationList200Response) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgsLicensePolicyViolationList200Response) GetPrevious() string {
	if o == nil || isNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgsLicensePolicyViolationList200Response) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *OrgsLicensePolicyViolationList200Response) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *OrgsLicensePolicyViolationList200Response) SetPrevious(v string) {
	o.Previous.Set(&v)
}

// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *OrgsLicensePolicyViolationList200Response) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *OrgsLicensePolicyViolationList200Response) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value
func (o *OrgsLicensePolicyViolationList200Response) GetResults() []PackageLicensePolicyViolationLog {
	if o == nil {
		var ret []PackageLicensePolicyViolationLog
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *OrgsLicensePolicyViolationList200Response) GetResultsOk() ([]PackageLicensePolicyViolationLog, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *OrgsLicensePolicyViolationList200Response) SetResults(v []PackageLicensePolicyViolationLog) {
	o.Results = v
}

func (o OrgsLicensePolicyViolationList200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableOrgsLicensePolicyViolationList200Response struct {
	value *OrgsLicensePolicyViolationList200Response
	isSet bool
}

func (v NullableOrgsLicensePolicyViolationList200Response) Get() *OrgsLicensePolicyViolationList200Response {
	return v.value
}

func (v *NullableOrgsLicensePolicyViolationList200Response) Set(val *OrgsLicensePolicyViolationList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgsLicensePolicyViolationList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgsLicensePolicyViolationList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgsLicensePolicyViolationList200Response(val *OrgsLicensePolicyViolationList200Response) *NullableOrgsLicensePolicyViolationList200Response {
	return &NullableOrgsLicensePolicyViolationList200Response{value: val, isSet: true}
}

func (v NullableOrgsLicensePolicyViolationList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgsLicensePolicyViolationList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
