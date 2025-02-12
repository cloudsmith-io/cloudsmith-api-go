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
)

// checks if the RepositoryGeoIpRulesRequestPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryGeoIpRulesRequestPatch{}

// RepositoryGeoIpRulesRequestPatch struct for RepositoryGeoIpRulesRequestPatch
type RepositoryGeoIpRulesRequestPatch struct {
	Cidr                 *RepositoryGeoIpCidr        `json:"cidr,omitempty"`
	CountryCode          *RepositoryGeoIpCountryCode `json:"country_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RepositoryGeoIpRulesRequestPatch RepositoryGeoIpRulesRequestPatch

// NewRepositoryGeoIpRulesRequestPatch instantiates a new RepositoryGeoIpRulesRequestPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGeoIpRulesRequestPatch() *RepositoryGeoIpRulesRequestPatch {
	this := RepositoryGeoIpRulesRequestPatch{}
	return &this
}

// NewRepositoryGeoIpRulesRequestPatchWithDefaults instantiates a new RepositoryGeoIpRulesRequestPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGeoIpRulesRequestPatchWithDefaults() *RepositoryGeoIpRulesRequestPatch {
	this := RepositoryGeoIpRulesRequestPatch{}
	return &this
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *RepositoryGeoIpRulesRequestPatch) GetCidr() RepositoryGeoIpCidr {
	if o == nil || IsNil(o.Cidr) {
		var ret RepositoryGeoIpCidr
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIpRulesRequestPatch) GetCidrOk() (*RepositoryGeoIpCidr, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *RepositoryGeoIpRulesRequestPatch) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given RepositoryGeoIpCidr and assigns it to the Cidr field.
func (o *RepositoryGeoIpRulesRequestPatch) SetCidr(v RepositoryGeoIpCidr) {
	o.Cidr = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *RepositoryGeoIpRulesRequestPatch) GetCountryCode() RepositoryGeoIpCountryCode {
	if o == nil || IsNil(o.CountryCode) {
		var ret RepositoryGeoIpCountryCode
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIpRulesRequestPatch) GetCountryCodeOk() (*RepositoryGeoIpCountryCode, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *RepositoryGeoIpRulesRequestPatch) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given RepositoryGeoIpCountryCode and assigns it to the CountryCode field.
func (o *RepositoryGeoIpRulesRequestPatch) SetCountryCode(v RepositoryGeoIpCountryCode) {
	o.CountryCode = &v
}

func (o RepositoryGeoIpRulesRequestPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryGeoIpRulesRequestPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RepositoryGeoIpRulesRequestPatch) UnmarshalJSON(data []byte) (err error) {
	varRepositoryGeoIpRulesRequestPatch := _RepositoryGeoIpRulesRequestPatch{}

	err = json.Unmarshal(data, &varRepositoryGeoIpRulesRequestPatch)

	if err != nil {
		return err
	}

	*o = RepositoryGeoIpRulesRequestPatch(varRepositoryGeoIpRulesRequestPatch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cidr")
		delete(additionalProperties, "country_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRepositoryGeoIpRulesRequestPatch struct {
	value *RepositoryGeoIpRulesRequestPatch
	isSet bool
}

func (v NullableRepositoryGeoIpRulesRequestPatch) Get() *RepositoryGeoIpRulesRequestPatch {
	return v.value
}

func (v *NullableRepositoryGeoIpRulesRequestPatch) Set(val *RepositoryGeoIpRulesRequestPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGeoIpRulesRequestPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGeoIpRulesRequestPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGeoIpRulesRequestPatch(val *RepositoryGeoIpRulesRequestPatch) *NullableRepositoryGeoIpRulesRequestPatch {
	return &NullableRepositoryGeoIpRulesRequestPatch{value: val, isSet: true}
}

func (v NullableRepositoryGeoIpRulesRequestPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGeoIpRulesRequestPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
