/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.327.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the Architecture type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Architecture{}

// Architecture struct for Architecture
type Architecture struct {
	Description NullableString `json:"description,omitempty"`
	Name        string         `json:"name"`
}

// NewArchitecture instantiates a new Architecture object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchitecture(name string) *Architecture {
	this := Architecture{}
	this.Name = name
	return &this
}

// NewArchitectureWithDefaults instantiates a new Architecture object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchitectureWithDefaults() *Architecture {
	this := Architecture{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Architecture) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Architecture) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Architecture) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Architecture) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Architecture) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Architecture) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value
func (o *Architecture) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Architecture) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Architecture) SetName(v string) {
	o.Name = v
}

func (o Architecture) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Architecture) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableArchitecture struct {
	value *Architecture
	isSet bool
}

func (v NullableArchitecture) Get() *Architecture {
	return v.value
}

func (v *NullableArchitecture) Set(val *Architecture) {
	v.value = val
	v.isSet = true
}

func (v NullableArchitecture) IsSet() bool {
	return v.isSet
}

func (v *NullableArchitecture) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchitecture(val *Architecture) *NullableArchitecture {
	return &NullableArchitecture{value: val, isSet: true}
}

func (v NullableArchitecture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchitecture) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
