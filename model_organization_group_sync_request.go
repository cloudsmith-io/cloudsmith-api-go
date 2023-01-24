/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.202.5
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// OrganizationGroupSyncRequest struct for OrganizationGroupSyncRequest
type OrganizationGroupSyncRequest struct {
	IdpKey       string `json:"idp_key"`
	IdpValue     string `json:"idp_value"`
	Organization string `json:"organization"`
	//  User role within the team.   A `manager` is capable of adding/removing others to/from the team, and  can set the role of other users and other settings pertaining to the  team.   A 'member' is a normal user that inherits the settings and privileges  assigned to the team.
	Role *string `json:"role,omitempty"`
	Team string  `json:"team"`
}

// NewOrganizationGroupSyncRequest instantiates a new OrganizationGroupSyncRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationGroupSyncRequest(idpKey string, idpValue string, organization string, team string) *OrganizationGroupSyncRequest {
	this := OrganizationGroupSyncRequest{}
	this.IdpKey = idpKey
	this.IdpValue = idpValue
	this.Organization = organization
	var role string = "Member"
	this.Role = &role
	this.Team = team
	return &this
}

// NewOrganizationGroupSyncRequestWithDefaults instantiates a new OrganizationGroupSyncRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationGroupSyncRequestWithDefaults() *OrganizationGroupSyncRequest {
	this := OrganizationGroupSyncRequest{}
	var role string = "Member"
	this.Role = &role
	return &this
}

// GetIdpKey returns the IdpKey field value
func (o *OrganizationGroupSyncRequest) GetIdpKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdpKey
}

// GetIdpKeyOk returns a tuple with the IdpKey field value
// and a boolean to check if the value has been set.
func (o *OrganizationGroupSyncRequest) GetIdpKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpKey, true
}

// SetIdpKey sets field value
func (o *OrganizationGroupSyncRequest) SetIdpKey(v string) {
	o.IdpKey = v
}

// GetIdpValue returns the IdpValue field value
func (o *OrganizationGroupSyncRequest) GetIdpValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdpValue
}

// GetIdpValueOk returns a tuple with the IdpValue field value
// and a boolean to check if the value has been set.
func (o *OrganizationGroupSyncRequest) GetIdpValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpValue, true
}

// SetIdpValue sets field value
func (o *OrganizationGroupSyncRequest) SetIdpValue(v string) {
	o.IdpValue = v
}

// GetOrganization returns the Organization field value
func (o *OrganizationGroupSyncRequest) GetOrganization() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *OrganizationGroupSyncRequest) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *OrganizationGroupSyncRequest) SetOrganization(v string) {
	o.Organization = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OrganizationGroupSyncRequest) GetRole() string {
	if o == nil || isNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationGroupSyncRequest) GetRoleOk() (*string, bool) {
	if o == nil || isNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OrganizationGroupSyncRequest) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OrganizationGroupSyncRequest) SetRole(v string) {
	o.Role = &v
}

// GetTeam returns the Team field value
func (o *OrganizationGroupSyncRequest) GetTeam() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Team
}

// GetTeamOk returns a tuple with the Team field value
// and a boolean to check if the value has been set.
func (o *OrganizationGroupSyncRequest) GetTeamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Team, true
}

// SetTeam sets field value
func (o *OrganizationGroupSyncRequest) SetTeam(v string) {
	o.Team = v
}

func (o OrganizationGroupSyncRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["idp_key"] = o.IdpKey
	}
	if true {
		toSerialize["idp_value"] = o.IdpValue
	}
	if true {
		toSerialize["organization"] = o.Organization
	}
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if true {
		toSerialize["team"] = o.Team
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationGroupSyncRequest struct {
	value *OrganizationGroupSyncRequest
	isSet bool
}

func (v NullableOrganizationGroupSyncRequest) Get() *OrganizationGroupSyncRequest {
	return v.value
}

func (v *NullableOrganizationGroupSyncRequest) Set(val *OrganizationGroupSyncRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationGroupSyncRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationGroupSyncRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationGroupSyncRequest(val *OrganizationGroupSyncRequest) *NullableOrganizationGroupSyncRequest {
	return &NullableOrganizationGroupSyncRequest{value: val, isSet: true}
}

func (v NullableOrganizationGroupSyncRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationGroupSyncRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
