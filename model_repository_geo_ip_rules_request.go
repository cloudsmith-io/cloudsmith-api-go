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

// RepositoryGeoIpRulesRequest struct for RepositoryGeoIpRulesRequest
type RepositoryGeoIpRulesRequest struct {
	Cidr        RepositoryGeoIpCidr        `json:"cidr"`
	CountryCode RepositoryGeoIpCountryCode `json:"country_code"`
}

// NewRepositoryGeoIpRulesRequest instantiates a new RepositoryGeoIpRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGeoIpRulesRequest(cidr RepositoryGeoIpCidr, countryCode RepositoryGeoIpCountryCode) *RepositoryGeoIpRulesRequest {
	this := RepositoryGeoIpRulesRequest{}
	this.Cidr = cidr
	this.CountryCode = countryCode
	return &this
}

// NewRepositoryGeoIpRulesRequestWithDefaults instantiates a new RepositoryGeoIpRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGeoIpRulesRequestWithDefaults() *RepositoryGeoIpRulesRequest {
	this := RepositoryGeoIpRulesRequest{}
	return &this
}

// GetCidr returns the Cidr field value
func (o *RepositoryGeoIpRulesRequest) GetCidr() RepositoryGeoIpCidr {
	if o == nil {
		var ret RepositoryGeoIpCidr
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIpRulesRequest) GetCidrOk() (*RepositoryGeoIpCidr, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *RepositoryGeoIpRulesRequest) SetCidr(v RepositoryGeoIpCidr) {
	o.Cidr = v
}

// GetCountryCode returns the CountryCode field value
func (o *RepositoryGeoIpRulesRequest) GetCountryCode() RepositoryGeoIpCountryCode {
	if o == nil {
		var ret RepositoryGeoIpCountryCode
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIpRulesRequest) GetCountryCodeOk() (*RepositoryGeoIpCountryCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *RepositoryGeoIpRulesRequest) SetCountryCode(v RepositoryGeoIpCountryCode) {
	o.CountryCode = v
}

func (o RepositoryGeoIpRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cidr"] = o.Cidr
	}
	if true {
		toSerialize["country_code"] = o.CountryCode
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryGeoIpRulesRequest struct {
	value *RepositoryGeoIpRulesRequest
	isSet bool
}

func (v NullableRepositoryGeoIpRulesRequest) Get() *RepositoryGeoIpRulesRequest {
	return v.value
}

func (v *NullableRepositoryGeoIpRulesRequest) Set(val *RepositoryGeoIpRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGeoIpRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGeoIpRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGeoIpRulesRequest(val *RepositoryGeoIpRulesRequest) *NullableRepositoryGeoIpRulesRequest {
	return &NullableRepositoryGeoIpRulesRequest{value: val, isSet: true}
}

func (v NullableRepositoryGeoIpRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGeoIpRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
