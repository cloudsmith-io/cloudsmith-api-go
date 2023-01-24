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

// RepositoryGeoIPList struct for RepositoryGeoIPList
type RepositoryGeoIPList struct {
	// List all CIDR geographic rules within the repository
	Cidr *string `json:"cidr,omitempty"`
	// List all GeoIP geographic rules within the repository
	CountryCode *string `json:"country_code,omitempty"`
}

// NewRepositoryGeoIPList instantiates a new RepositoryGeoIPList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGeoIPList() *RepositoryGeoIPList {
	this := RepositoryGeoIPList{}
	return &this
}

// NewRepositoryGeoIPListWithDefaults instantiates a new RepositoryGeoIPList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGeoIPListWithDefaults() *RepositoryGeoIPList {
	this := RepositoryGeoIPList{}
	return &this
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *RepositoryGeoIPList) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIPList) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *RepositoryGeoIPList) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *RepositoryGeoIPList) SetCidr(v string) {
	o.Cidr = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *RepositoryGeoIPList) GetCountryCode() string {
	if o == nil || isNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIPList) GetCountryCodeOk() (*string, bool) {
	if o == nil || isNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *RepositoryGeoIPList) HasCountryCode() bool {
	if o != nil && !isNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *RepositoryGeoIPList) SetCountryCode(v string) {
	o.CountryCode = &v
}

func (o RepositoryGeoIPList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !isNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryGeoIPList struct {
	value *RepositoryGeoIPList
	isSet bool
}

func (v NullableRepositoryGeoIPList) Get() *RepositoryGeoIPList {
	return v.value
}

func (v *NullableRepositoryGeoIPList) Set(val *RepositoryGeoIPList) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGeoIPList) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGeoIPList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGeoIPList(val *RepositoryGeoIPList) *NullableRepositoryGeoIPList {
	return &NullableRepositoryGeoIPList{value: val, isSet: true}
}

func (v NullableRepositoryGeoIPList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGeoIPList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}