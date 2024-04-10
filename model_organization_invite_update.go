/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.392.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the OrganizationInviteUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationInviteUpdate{}

// OrganizationInviteUpdate struct for OrganizationInviteUpdate
type OrganizationInviteUpdate struct {
	// The role to be assigned to the invited user.
	Role *string `json:"role,omitempty"`
}

// NewOrganizationInviteUpdate instantiates a new OrganizationInviteUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInviteUpdate() *OrganizationInviteUpdate {
	this := OrganizationInviteUpdate{}
	var role string = "Member"
	this.Role = &role
	return &this
}

// NewOrganizationInviteUpdateWithDefaults instantiates a new OrganizationInviteUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInviteUpdateWithDefaults() *OrganizationInviteUpdate {
	this := OrganizationInviteUpdate{}
	var role string = "Member"
	this.Role = &role
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OrganizationInviteUpdate) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInviteUpdate) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OrganizationInviteUpdate) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OrganizationInviteUpdate) SetRole(v string) {
	o.Role = &v
}

func (o OrganizationInviteUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationInviteUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableOrganizationInviteUpdate struct {
	value *OrganizationInviteUpdate
	isSet bool
}

func (v NullableOrganizationInviteUpdate) Get() *OrganizationInviteUpdate {
	return v.value
}

func (v *NullableOrganizationInviteUpdate) Set(val *OrganizationInviteUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationInviteUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationInviteUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationInviteUpdate(val *OrganizationInviteUpdate) *NullableOrganizationInviteUpdate {
	return &NullableOrganizationInviteUpdate{value: val, isSet: true}
}

func (v NullableOrganizationInviteUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationInviteUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
