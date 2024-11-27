/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.566.9
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the OrganizationTeamInvite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationTeamInvite{}

// OrganizationTeamInvite struct for OrganizationTeamInvite
type OrganizationTeamInvite struct {
	// The role to be assigned to the invited user in the team.
	Role *string `json:"role,omitempty"`
	// The team identifier (slug).
	Team string `json:"team" validate:"regexp=^[-a-zA-Z0-9_]+$"`
}

type _OrganizationTeamInvite OrganizationTeamInvite

// NewOrganizationTeamInvite instantiates a new OrganizationTeamInvite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationTeamInvite(team string) *OrganizationTeamInvite {
	this := OrganizationTeamInvite{}
	var role string = "Member"
	this.Role = &role
	this.Team = team
	return &this
}

// NewOrganizationTeamInviteWithDefaults instantiates a new OrganizationTeamInvite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationTeamInviteWithDefaults() *OrganizationTeamInvite {
	this := OrganizationTeamInvite{}
	var role string = "Member"
	this.Role = &role
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OrganizationTeamInvite) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeamInvite) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OrganizationTeamInvite) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OrganizationTeamInvite) SetRole(v string) {
	o.Role = &v
}

// GetTeam returns the Team field value
func (o *OrganizationTeamInvite) GetTeam() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Team
}

// GetTeamOk returns a tuple with the Team field value
// and a boolean to check if the value has been set.
func (o *OrganizationTeamInvite) GetTeamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Team, true
}

// SetTeam sets field value
func (o *OrganizationTeamInvite) SetTeam(v string) {
	o.Team = v
}

func (o OrganizationTeamInvite) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationTeamInvite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	toSerialize["team"] = o.Team
	return toSerialize, nil
}

func (o *OrganizationTeamInvite) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"team",
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

	varOrganizationTeamInvite := _OrganizationTeamInvite{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrganizationTeamInvite)

	if err != nil {
		return err
	}

	*o = OrganizationTeamInvite(varOrganizationTeamInvite)

	return err
}

type NullableOrganizationTeamInvite struct {
	value *OrganizationTeamInvite
	isSet bool
}

func (v NullableOrganizationTeamInvite) Get() *OrganizationTeamInvite {
	return v.value
}

func (v *NullableOrganizationTeamInvite) Set(val *OrganizationTeamInvite) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationTeamInvite) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationTeamInvite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationTeamInvite(val *OrganizationTeamInvite) *NullableOrganizationTeamInvite {
	return &NullableOrganizationTeamInvite{value: val, isSet: true}
}

func (v NullableOrganizationTeamInvite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationTeamInvite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
