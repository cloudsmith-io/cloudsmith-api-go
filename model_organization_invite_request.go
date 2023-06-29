/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.290.2
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// OrganizationInviteRequest struct for OrganizationInviteRequest
type OrganizationInviteRequest struct {
	// The email of the user to be invited.
	Email *string `json:"email,omitempty"`
	// The role to be assigned to the invited user.
	Role *string `json:"role,omitempty"`
	// The slug of the user to be invited.
	User *string `json:"user,omitempty"`
}

// NewOrganizationInviteRequest instantiates a new OrganizationInviteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInviteRequest() *OrganizationInviteRequest {
	this := OrganizationInviteRequest{}
	var role string = "Member"
	this.Role = &role
	return &this
}

// NewOrganizationInviteRequestWithDefaults instantiates a new OrganizationInviteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInviteRequestWithDefaults() *OrganizationInviteRequest {
	this := OrganizationInviteRequest{}
	var role string = "Member"
	this.Role = &role
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrganizationInviteRequest) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInviteRequest) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrganizationInviteRequest) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrganizationInviteRequest) SetEmail(v string) {
	o.Email = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OrganizationInviteRequest) GetRole() string {
	if o == nil || isNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInviteRequest) GetRoleOk() (*string, bool) {
	if o == nil || isNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OrganizationInviteRequest) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OrganizationInviteRequest) SetRole(v string) {
	o.Role = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *OrganizationInviteRequest) GetUser() string {
	if o == nil || isNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInviteRequest) GetUserOk() (*string, bool) {
	if o == nil || isNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *OrganizationInviteRequest) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *OrganizationInviteRequest) SetUser(v string) {
	o.User = &v
}

func (o OrganizationInviteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationInviteRequest struct {
	value *OrganizationInviteRequest
	isSet bool
}

func (v NullableOrganizationInviteRequest) Get() *OrganizationInviteRequest {
	return v.value
}

func (v *NullableOrganizationInviteRequest) Set(val *OrganizationInviteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationInviteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationInviteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationInviteRequest(val *OrganizationInviteRequest) *NullableOrganizationInviteRequest {
	return &NullableOrganizationInviteRequest{value: val, isSet: true}
}

func (v NullableOrganizationInviteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationInviteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
