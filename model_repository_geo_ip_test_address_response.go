/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.206.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// RepositoryGeoIPTestAddressResponse struct for RepositoryGeoIPTestAddressResponse
type RepositoryGeoIPTestAddressResponse struct {
	// The IP address test results ordered by allowed
	Addresses []RepositoryGeoIPTestAddressResponseDict `json:"addresses"`
}

// NewRepositoryGeoIPTestAddressResponse instantiates a new RepositoryGeoIPTestAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGeoIPTestAddressResponse(addresses []RepositoryGeoIPTestAddressResponseDict) *RepositoryGeoIPTestAddressResponse {
	this := RepositoryGeoIPTestAddressResponse{}
	this.Addresses = addresses
	return &this
}

// NewRepositoryGeoIPTestAddressResponseWithDefaults instantiates a new RepositoryGeoIPTestAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGeoIPTestAddressResponseWithDefaults() *RepositoryGeoIPTestAddressResponse {
	this := RepositoryGeoIPTestAddressResponse{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *RepositoryGeoIPTestAddressResponse) GetAddresses() []RepositoryGeoIPTestAddressResponseDict {
	if o == nil {
		var ret []RepositoryGeoIPTestAddressResponseDict
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIPTestAddressResponse) GetAddressesOk() ([]RepositoryGeoIPTestAddressResponseDict, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *RepositoryGeoIPTestAddressResponse) SetAddresses(v []RepositoryGeoIPTestAddressResponseDict) {
	o.Addresses = v
}

func (o RepositoryGeoIPTestAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryGeoIPTestAddressResponse struct {
	value *RepositoryGeoIPTestAddressResponse
	isSet bool
}

func (v NullableRepositoryGeoIPTestAddressResponse) Get() *RepositoryGeoIPTestAddressResponse {
	return v.value
}

func (v *NullableRepositoryGeoIPTestAddressResponse) Set(val *RepositoryGeoIPTestAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGeoIPTestAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGeoIPTestAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGeoIPTestAddressResponse(val *RepositoryGeoIPTestAddressResponse) *NullableRepositoryGeoIPTestAddressResponse {
	return &NullableRepositoryGeoIPTestAddressResponse{value: val, isSet: true}
}

func (v NullableRepositoryGeoIPTestAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGeoIPTestAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
