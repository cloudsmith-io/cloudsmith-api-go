/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.236.5
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// RepositoryPrivilegeInput struct for RepositoryPrivilegeInput
type RepositoryPrivilegeInput struct {
	// List of objects with explicit privileges to the repository.
	Privileges []RepositoryPrivilegeDict `json:"privileges"`
}

// NewRepositoryPrivilegeInput instantiates a new RepositoryPrivilegeInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryPrivilegeInput(privileges []RepositoryPrivilegeDict) *RepositoryPrivilegeInput {
	this := RepositoryPrivilegeInput{}
	this.Privileges = privileges
	return &this
}

// NewRepositoryPrivilegeInputWithDefaults instantiates a new RepositoryPrivilegeInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryPrivilegeInputWithDefaults() *RepositoryPrivilegeInput {
	this := RepositoryPrivilegeInput{}
	return &this
}

// GetPrivileges returns the Privileges field value
func (o *RepositoryPrivilegeInput) GetPrivileges() []RepositoryPrivilegeDict {
	if o == nil {
		var ret []RepositoryPrivilegeDict
		return ret
	}

	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value
// and a boolean to check if the value has been set.
func (o *RepositoryPrivilegeInput) GetPrivilegesOk() ([]RepositoryPrivilegeDict, bool) {
	if o == nil {
		return nil, false
	}
	return o.Privileges, true
}

// SetPrivileges sets field value
func (o *RepositoryPrivilegeInput) SetPrivileges(v []RepositoryPrivilegeDict) {
	o.Privileges = v
}

func (o RepositoryPrivilegeInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["privileges"] = o.Privileges
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryPrivilegeInput struct {
	value *RepositoryPrivilegeInput
	isSet bool
}

func (v NullableRepositoryPrivilegeInput) Get() *RepositoryPrivilegeInput {
	return v.value
}

func (v *NullableRepositoryPrivilegeInput) Set(val *RepositoryPrivilegeInput) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryPrivilegeInput) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryPrivilegeInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryPrivilegeInput(val *RepositoryPrivilegeInput) *NullableRepositoryPrivilegeInput {
	return &NullableRepositoryPrivilegeInput{value: val, isSet: true}
}

func (v NullableRepositoryPrivilegeInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryPrivilegeInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
