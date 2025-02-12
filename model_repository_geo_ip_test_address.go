/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.616.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"fmt"
)

// checks if the RepositoryGeoIpTestAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryGeoIpTestAddress{}

// RepositoryGeoIpTestAddress struct for RepositoryGeoIpTestAddress
type RepositoryGeoIpTestAddress struct {
	// The IP addresses to test against this repository
	Addresses            []string `json:"addresses"`
	AdditionalProperties map[string]interface{}
}

type _RepositoryGeoIpTestAddress RepositoryGeoIpTestAddress

// NewRepositoryGeoIpTestAddress instantiates a new RepositoryGeoIpTestAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGeoIpTestAddress(addresses []string) *RepositoryGeoIpTestAddress {
	this := RepositoryGeoIpTestAddress{}
	this.Addresses = addresses
	return &this
}

// NewRepositoryGeoIpTestAddressWithDefaults instantiates a new RepositoryGeoIpTestAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGeoIpTestAddressWithDefaults() *RepositoryGeoIpTestAddress {
	this := RepositoryGeoIpTestAddress{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *RepositoryGeoIpTestAddress) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIpTestAddress) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *RepositoryGeoIpTestAddress) SetAddresses(v []string) {
	o.Addresses = v
}

func (o RepositoryGeoIpTestAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryGeoIpTestAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addresses"] = o.Addresses

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RepositoryGeoIpTestAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addresses",
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

	varRepositoryGeoIpTestAddress := _RepositoryGeoIpTestAddress{}

	err = json.Unmarshal(data, &varRepositoryGeoIpTestAddress)

	if err != nil {
		return err
	}

	*o = RepositoryGeoIpTestAddress(varRepositoryGeoIpTestAddress)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "addresses")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRepositoryGeoIpTestAddress struct {
	value *RepositoryGeoIpTestAddress
	isSet bool
}

func (v NullableRepositoryGeoIpTestAddress) Get() *RepositoryGeoIpTestAddress {
	return v.value
}

func (v *NullableRepositoryGeoIpTestAddress) Set(val *RepositoryGeoIpTestAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGeoIpTestAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGeoIpTestAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGeoIpTestAddress(val *RepositoryGeoIpTestAddress) *NullableRepositoryGeoIpTestAddress {
	return &NullableRepositoryGeoIpTestAddress{value: val, isSet: true}
}

func (v NullableRepositoryGeoIpTestAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGeoIpTestAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
