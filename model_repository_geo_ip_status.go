/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.478.2
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the RepositoryGeoIpStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryGeoIpStatus{}

// RepositoryGeoIpStatus struct for RepositoryGeoIpStatus
type RepositoryGeoIpStatus struct {
	// If checked, any access to the website or downloads for this repository is allowed/denied according to the configured Geo/IP restriction rules.
	GeoipEnabled *bool `json:"geoip_enabled,omitempty"`
}

// NewRepositoryGeoIpStatus instantiates a new RepositoryGeoIpStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGeoIpStatus() *RepositoryGeoIpStatus {
	this := RepositoryGeoIpStatus{}
	return &this
}

// NewRepositoryGeoIpStatusWithDefaults instantiates a new RepositoryGeoIpStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGeoIpStatusWithDefaults() *RepositoryGeoIpStatus {
	this := RepositoryGeoIpStatus{}
	return &this
}

// GetGeoipEnabled returns the GeoipEnabled field value if set, zero value otherwise.
func (o *RepositoryGeoIpStatus) GetGeoipEnabled() bool {
	if o == nil || IsNil(o.GeoipEnabled) {
		var ret bool
		return ret
	}
	return *o.GeoipEnabled
}

// GetGeoipEnabledOk returns a tuple with the GeoipEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIpStatus) GetGeoipEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.GeoipEnabled) {
		return nil, false
	}
	return o.GeoipEnabled, true
}

// HasGeoipEnabled returns a boolean if a field has been set.
func (o *RepositoryGeoIpStatus) HasGeoipEnabled() bool {
	if o != nil && !IsNil(o.GeoipEnabled) {
		return true
	}

	return false
}

// SetGeoipEnabled gets a reference to the given bool and assigns it to the GeoipEnabled field.
func (o *RepositoryGeoIpStatus) SetGeoipEnabled(v bool) {
	o.GeoipEnabled = &v
}

func (o RepositoryGeoIpStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryGeoIpStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GeoipEnabled) {
		toSerialize["geoip_enabled"] = o.GeoipEnabled
	}
	return toSerialize, nil
}

type NullableRepositoryGeoIpStatus struct {
	value *RepositoryGeoIpStatus
	isSet bool
}

func (v NullableRepositoryGeoIpStatus) Get() *RepositoryGeoIpStatus {
	return v.value
}

func (v *NullableRepositoryGeoIpStatus) Set(val *RepositoryGeoIpStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGeoIpStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGeoIpStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGeoIpStatus(val *RepositoryGeoIpStatus) *NullableRepositoryGeoIpStatus {
	return &NullableRepositoryGeoIpStatus{value: val, isSet: true}
}

func (v NullableRepositoryGeoIpStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGeoIpStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}