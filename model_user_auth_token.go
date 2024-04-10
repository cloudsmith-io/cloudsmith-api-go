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

// checks if the UserAuthToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAuthToken{}

// UserAuthToken struct for UserAuthToken
type UserAuthToken struct {
	// API token for the authenticated user
	Token *string `json:"token,omitempty"`
}

// NewUserAuthToken instantiates a new UserAuthToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAuthToken() *UserAuthToken {
	this := UserAuthToken{}
	return &this
}

// NewUserAuthTokenWithDefaults instantiates a new UserAuthToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAuthTokenWithDefaults() *UserAuthToken {
	this := UserAuthToken{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UserAuthToken) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAuthToken) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UserAuthToken) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UserAuthToken) SetToken(v string) {
	o.Token = &v
}

func (o UserAuthToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAuthToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableUserAuthToken struct {
	value *UserAuthToken
	isSet bool
}

func (v NullableUserAuthToken) Get() *UserAuthToken {
	return v.value
}

func (v *NullableUserAuthToken) Set(val *UserAuthToken) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAuthToken) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAuthToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAuthToken(val *UserAuthToken) *NullableUserAuthToken {
	return &NullableUserAuthToken{value: val, isSet: true}
}

func (v NullableUserAuthToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAuthToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
