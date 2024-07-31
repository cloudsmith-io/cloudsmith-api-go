/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.477.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the PackageUsageMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageUsageMetrics{}

// PackageUsageMetrics struct for PackageUsageMetrics
type PackageUsageMetrics struct {
	Packages CommonMetrics `json:"packages"`
}

// NewPackageUsageMetrics instantiates a new PackageUsageMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageUsageMetrics(packages CommonMetrics) *PackageUsageMetrics {
	this := PackageUsageMetrics{}
	this.Packages = packages
	return &this
}

// NewPackageUsageMetricsWithDefaults instantiates a new PackageUsageMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageUsageMetricsWithDefaults() *PackageUsageMetrics {
	this := PackageUsageMetrics{}
	return &this
}

// GetPackages returns the Packages field value
func (o *PackageUsageMetrics) GetPackages() CommonMetrics {
	if o == nil {
		var ret CommonMetrics
		return ret
	}

	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value
// and a boolean to check if the value has been set.
func (o *PackageUsageMetrics) GetPackagesOk() (*CommonMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Packages, true
}

// SetPackages sets field value
func (o *PackageUsageMetrics) SetPackages(v CommonMetrics) {
	o.Packages = v
}

func (o PackageUsageMetrics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageUsageMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["packages"] = o.Packages
	return toSerialize, nil
}

type NullablePackageUsageMetrics struct {
	value *PackageUsageMetrics
	isSet bool
}

func (v NullablePackageUsageMetrics) Get() *PackageUsageMetrics {
	return v.value
}

func (v *NullablePackageUsageMetrics) Set(val *PackageUsageMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageUsageMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageUsageMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageUsageMetrics(val *PackageUsageMetrics) *NullablePackageUsageMetrics {
	return &NullablePackageUsageMetrics{value: val, isSet: true}
}

func (v NullablePackageUsageMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageUsageMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
