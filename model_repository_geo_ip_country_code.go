/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.247.7
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// RepositoryGeoIPCountryCode struct for RepositoryGeoIPCountryCode
type RepositoryGeoIPCountryCode struct {
	// The allowed country codes for this repository
	Allow []string `json:"allow"`
	// The denied country codes for this repository
	Deny []string `json:"deny"`
}

// NewRepositoryGeoIPCountryCode instantiates a new RepositoryGeoIPCountryCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGeoIPCountryCode(allow []string, deny []string) *RepositoryGeoIPCountryCode {
	this := RepositoryGeoIPCountryCode{}
	this.Allow = allow
	this.Deny = deny
	return &this
}

// NewRepositoryGeoIPCountryCodeWithDefaults instantiates a new RepositoryGeoIPCountryCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGeoIPCountryCodeWithDefaults() *RepositoryGeoIPCountryCode {
	this := RepositoryGeoIPCountryCode{}
	return &this
}

// GetAllow returns the Allow field value
func (o *RepositoryGeoIPCountryCode) GetAllow() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIPCountryCode) GetAllowOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Allow, true
}

// SetAllow sets field value
func (o *RepositoryGeoIPCountryCode) SetAllow(v []string) {
	o.Allow = v
}

// GetDeny returns the Deny field value
func (o *RepositoryGeoIPCountryCode) GetDeny() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Deny
}

// GetDenyOk returns a tuple with the Deny field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIPCountryCode) GetDenyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deny, true
}

// SetDeny sets field value
func (o *RepositoryGeoIPCountryCode) SetDeny(v []string) {
	o.Deny = v
}

func (o RepositoryGeoIPCountryCode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allow"] = o.Allow
	}
	if true {
		toSerialize["deny"] = o.Deny
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryGeoIPCountryCode struct {
	value *RepositoryGeoIPCountryCode
	isSet bool
}

func (v NullableRepositoryGeoIPCountryCode) Get() *RepositoryGeoIPCountryCode {
	return v.value
}

func (v *NullableRepositoryGeoIPCountryCode) Set(val *RepositoryGeoIPCountryCode) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGeoIPCountryCode) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGeoIPCountryCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGeoIPCountryCode(val *RepositoryGeoIPCountryCode) *NullableRepositoryGeoIPCountryCode {
	return &NullableRepositoryGeoIPCountryCode{value: val, isSet: true}
}

func (v NullableRepositoryGeoIPCountryCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGeoIPCountryCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}