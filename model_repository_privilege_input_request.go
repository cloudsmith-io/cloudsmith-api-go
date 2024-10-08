/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.533.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the RepositoryPrivilegeInputRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryPrivilegeInputRequest{}

// RepositoryPrivilegeInputRequest struct for RepositoryPrivilegeInputRequest
type RepositoryPrivilegeInputRequest struct {
	// List of objects with explicit privileges to the repository.
	Privileges []RepositoryPrivilegeDict `json:"privileges"`
}

// NewRepositoryPrivilegeInputRequest instantiates a new RepositoryPrivilegeInputRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryPrivilegeInputRequest(privileges []RepositoryPrivilegeDict) *RepositoryPrivilegeInputRequest {
	this := RepositoryPrivilegeInputRequest{}
	this.Privileges = privileges
	return &this
}

// NewRepositoryPrivilegeInputRequestWithDefaults instantiates a new RepositoryPrivilegeInputRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryPrivilegeInputRequestWithDefaults() *RepositoryPrivilegeInputRequest {
	this := RepositoryPrivilegeInputRequest{}
	return &this
}

// GetPrivileges returns the Privileges field value
func (o *RepositoryPrivilegeInputRequest) GetPrivileges() []RepositoryPrivilegeDict {
	if o == nil {
		var ret []RepositoryPrivilegeDict
		return ret
	}

	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value
// and a boolean to check if the value has been set.
func (o *RepositoryPrivilegeInputRequest) GetPrivilegesOk() ([]RepositoryPrivilegeDict, bool) {
	if o == nil {
		return nil, false
	}
	return o.Privileges, true
}

// SetPrivileges sets field value
func (o *RepositoryPrivilegeInputRequest) SetPrivileges(v []RepositoryPrivilegeDict) {
	o.Privileges = v
}

func (o RepositoryPrivilegeInputRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryPrivilegeInputRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["privileges"] = o.Privileges
	return toSerialize, nil
}

type NullableRepositoryPrivilegeInputRequest struct {
	value *RepositoryPrivilegeInputRequest
	isSet bool
}

func (v NullableRepositoryPrivilegeInputRequest) Get() *RepositoryPrivilegeInputRequest {
	return v.value
}

func (v *NullableRepositoryPrivilegeInputRequest) Set(val *RepositoryPrivilegeInputRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryPrivilegeInputRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryPrivilegeInputRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryPrivilegeInputRequest(val *RepositoryPrivilegeInputRequest) *NullableRepositoryPrivilegeInputRequest {
	return &NullableRepositoryPrivilegeInputRequest{value: val, isSet: true}
}

func (v NullableRepositoryPrivilegeInputRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryPrivilegeInputRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
