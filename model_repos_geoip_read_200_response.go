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

// ReposGeoipRead200Response struct for ReposGeoipRead200Response
type ReposGeoipRead200Response struct {
	CountryCode *ReposGeoipRead200ResponseCountryCode `json:"country_code,omitempty"`
	Cidr        *ReposGeoipRead200ResponseCountryCode `json:"cidr,omitempty"`
}

// NewReposGeoipRead200Response instantiates a new ReposGeoipRead200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReposGeoipRead200Response() *ReposGeoipRead200Response {
	this := ReposGeoipRead200Response{}
	return &this
}

// NewReposGeoipRead200ResponseWithDefaults instantiates a new ReposGeoipRead200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReposGeoipRead200ResponseWithDefaults() *ReposGeoipRead200Response {
	this := ReposGeoipRead200Response{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *ReposGeoipRead200Response) GetCountryCode() ReposGeoipRead200ResponseCountryCode {
	if o == nil || isNil(o.CountryCode) {
		var ret ReposGeoipRead200ResponseCountryCode
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposGeoipRead200Response) GetCountryCodeOk() (*ReposGeoipRead200ResponseCountryCode, bool) {
	if o == nil || isNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *ReposGeoipRead200Response) HasCountryCode() bool {
	if o != nil && !isNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given ReposGeoipRead200ResponseCountryCode and assigns it to the CountryCode field.
func (o *ReposGeoipRead200Response) SetCountryCode(v ReposGeoipRead200ResponseCountryCode) {
	o.CountryCode = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *ReposGeoipRead200Response) GetCidr() ReposGeoipRead200ResponseCountryCode {
	if o == nil || isNil(o.Cidr) {
		var ret ReposGeoipRead200ResponseCountryCode
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposGeoipRead200Response) GetCidrOk() (*ReposGeoipRead200ResponseCountryCode, bool) {
	if o == nil || isNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *ReposGeoipRead200Response) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given ReposGeoipRead200ResponseCountryCode and assigns it to the Cidr field.
func (o *ReposGeoipRead200Response) SetCidr(v ReposGeoipRead200ResponseCountryCode) {
	o.Cidr = &v
}

func (o ReposGeoipRead200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	return json.Marshal(toSerialize)
}

type NullableReposGeoipRead200Response struct {
	value *ReposGeoipRead200Response
	isSet bool
}

func (v NullableReposGeoipRead200Response) Get() *ReposGeoipRead200Response {
	return v.value
}

func (v *NullableReposGeoipRead200Response) Set(val *ReposGeoipRead200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReposGeoipRead200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReposGeoipRead200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReposGeoipRead200Response(val *ReposGeoipRead200Response) *NullableReposGeoipRead200Response {
	return &NullableReposGeoipRead200Response{value: val, isSet: true}
}

func (v NullableReposGeoipRead200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReposGeoipRead200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
