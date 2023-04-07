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

// RepositoryPrivilegeInputRequestPatch struct for RepositoryPrivilegeInputRequestPatch
type RepositoryPrivilegeInputRequestPatch struct {
	// List of objects with explicit privileges to the repository.
	Privileges []RepositoryPrivilegeDict `json:"privileges,omitempty"`
}

// NewRepositoryPrivilegeInputRequestPatch instantiates a new RepositoryPrivilegeInputRequestPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryPrivilegeInputRequestPatch() *RepositoryPrivilegeInputRequestPatch {
	this := RepositoryPrivilegeInputRequestPatch{}
	return &this
}

// NewRepositoryPrivilegeInputRequestPatchWithDefaults instantiates a new RepositoryPrivilegeInputRequestPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryPrivilegeInputRequestPatchWithDefaults() *RepositoryPrivilegeInputRequestPatch {
	this := RepositoryPrivilegeInputRequestPatch{}
	return &this
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *RepositoryPrivilegeInputRequestPatch) GetPrivileges() []RepositoryPrivilegeDict {
	if o == nil || isNil(o.Privileges) {
		var ret []RepositoryPrivilegeDict
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryPrivilegeInputRequestPatch) GetPrivilegesOk() ([]RepositoryPrivilegeDict, bool) {
	if o == nil || isNil(o.Privileges) {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *RepositoryPrivilegeInputRequestPatch) HasPrivileges() bool {
	if o != nil && !isNil(o.Privileges) {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []RepositoryPrivilegeDict and assigns it to the Privileges field.
func (o *RepositoryPrivilegeInputRequestPatch) SetPrivileges(v []RepositoryPrivilegeDict) {
	o.Privileges = v
}

func (o RepositoryPrivilegeInputRequestPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Privileges) {
		toSerialize["privileges"] = o.Privileges
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryPrivilegeInputRequestPatch struct {
	value *RepositoryPrivilegeInputRequestPatch
	isSet bool
}

func (v NullableRepositoryPrivilegeInputRequestPatch) Get() *RepositoryPrivilegeInputRequestPatch {
	return v.value
}

func (v *NullableRepositoryPrivilegeInputRequestPatch) Set(val *RepositoryPrivilegeInputRequestPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryPrivilegeInputRequestPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryPrivilegeInputRequestPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryPrivilegeInputRequestPatch(val *RepositoryPrivilegeInputRequestPatch) *NullableRepositoryPrivilegeInputRequestPatch {
	return &NullableRepositoryPrivilegeInputRequestPatch{value: val, isSet: true}
}

func (v NullableRepositoryPrivilegeInputRequestPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryPrivilegeInputRequestPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
