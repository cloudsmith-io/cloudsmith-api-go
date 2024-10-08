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

// checks if the OrganizationTeamMembership type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationTeamMembership{}

// OrganizationTeamMembership The team members
type OrganizationTeamMembership struct {
	Role string `json:"role"`
	User string `json:"user"`
}

// NewOrganizationTeamMembership instantiates a new OrganizationTeamMembership object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationTeamMembership(role string, user string) *OrganizationTeamMembership {
	this := OrganizationTeamMembership{}
	this.Role = role
	this.User = user
	return &this
}

// NewOrganizationTeamMembershipWithDefaults instantiates a new OrganizationTeamMembership object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationTeamMembershipWithDefaults() *OrganizationTeamMembership {
	this := OrganizationTeamMembership{}
	return &this
}

// GetRole returns the Role field value
func (o *OrganizationTeamMembership) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *OrganizationTeamMembership) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *OrganizationTeamMembership) SetRole(v string) {
	o.Role = v
}

// GetUser returns the User field value
func (o *OrganizationTeamMembership) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *OrganizationTeamMembership) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *OrganizationTeamMembership) SetUser(v string) {
	o.User = v
}

func (o OrganizationTeamMembership) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationTeamMembership) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role
	toSerialize["user"] = o.User
	return toSerialize, nil
}

type NullableOrganizationTeamMembership struct {
	value *OrganizationTeamMembership
	isSet bool
}

func (v NullableOrganizationTeamMembership) Get() *OrganizationTeamMembership {
	return v.value
}

func (v *NullableOrganizationTeamMembership) Set(val *OrganizationTeamMembership) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationTeamMembership) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationTeamMembership) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationTeamMembership(val *OrganizationTeamMembership) *NullableOrganizationTeamMembership {
	return &NullableOrganizationTeamMembership{value: val, isSet: true}
}

func (v NullableOrganizationTeamMembership) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationTeamMembership) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
