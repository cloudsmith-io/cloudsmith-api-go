/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.568.8
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the Eula type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Eula{}

// Eula struct for Eula
type Eula struct {
	// A unique identifier that you can use for your own EULA tracking purposes. This might be a date, or a semantic version, etc. The only requirement is that it is unique across multiple EULAs.
	Identifier NullableString `json:"identifier,omitempty" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	// A sequential identifier that increments by one for each new commit in a repository.
	Number               NullableInt64 `json:"number,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Eula Eula

// NewEula instantiates a new Eula object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEula() *Eula {
	this := Eula{}
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
	if o == nil || IsNil(o.Identifier.Get()) {
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

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Eula) GetNumber() int64 {
	if o == nil || IsNil(o.Number.Get()) {
		var ret int64
		return ret
	}
	return *o.Number.Get()
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Eula) GetNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Number.Get(), o.Number.IsSet()
}

// HasNumber returns a boolean if a field has been set.
func (o *Eula) HasNumber() bool {
	if o != nil && o.Number.IsSet() {
		return true
	}

	return false
}

// SetNumber gets a reference to the given NullableInt64 and assigns it to the Number field.
func (o *Eula) SetNumber(v int64) {
	o.Number.Set(&v)
}

// SetNumberNil sets the value for Number to be an explicit nil
func (o *Eula) SetNumberNil() {
	o.Number.Set(nil)
}

// UnsetNumber ensures that no value is present for Number, not even an explicit nil
func (o *Eula) UnsetNumber() {
	o.Number.Unset()
}

func (o Eula) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Eula) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}
	if o.Number.IsSet() {
		toSerialize["number"] = o.Number.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Eula) UnmarshalJSON(data []byte) (err error) {
	varEula := _Eula{}

	err = json.Unmarshal(data, &varEula)

	if err != nil {
		return err
	}

	*o = Eula(varEula)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "number")
		o.AdditionalProperties = additionalProperties
	}

	return err
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
