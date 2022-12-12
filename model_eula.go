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

// Eula struct for Eula
type Eula struct {
	// A unique identifier that you can use for your own EULA tracking purposes. This might be a date, or a semantic version, etc. The only requirement is that it is unique across multiple EULAs.
	Identifier NullableString `json:"identifier,omitempty"`
	// A sequential identifier that increments by one for each new commit in a repository.
	Number int64 `json:"number"`
}

// NewEula instantiates a new Eula object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEula(number int64) *Eula {
	this := Eula{}
	this.Number = number
	return &this
}

// NewEulaWithDefaults instantiates a new Eula object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEulaWithDefaults() *Eula {
	this := Eula{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Eula) GetIdentifier() string {
	if o == nil || isNil(o.Identifier.Get()) {
		var ret string
		return ret
	}
	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Eula) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Eula) HasIdentifier() bool {
	if o != nil && o.Identifier.IsSet() {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given NullableString and assigns it to the Identifier field.
func (o *Eula) SetIdentifier(v string) {
	o.Identifier.Set(&v)
}

// SetIdentifierNil sets the value for Identifier to be an explicit nil
func (o *Eula) SetIdentifierNil() {
	o.Identifier.Set(nil)
}

// UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
func (o *Eula) UnsetIdentifier() {
	o.Identifier.Unset()
}

// GetNumber returns the Number field value
func (o *Eula) GetNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *Eula) GetNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *Eula) SetNumber(v int64) {
	o.Number = v
}

func (o Eula) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}
	if true {
		toSerialize["number"] = o.Number
	}
	return json.Marshal(toSerialize)
}

type NullableEula struct {
	value *Eula
	isSet bool
}

func (v NullableEula) Get() *Eula {
	return v.value
}

func (v *NullableEula) Set(val *Eula) {
	v.value = val
	v.isSet = true
}

func (v NullableEula) IsSet() bool {
	return v.isSet
}

func (v *NullableEula) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEula(val *Eula) *NullableEula {
	return &NullableEula{value: val, isSet: true}
}

func (v NullableEula) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEula) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
