/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.202.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// PackageQuarantineRequest struct for PackageQuarantineRequest
type PackageQuarantineRequest struct {
	// If true, the package be restored from quarantine.
	Restore *bool `json:"restore,omitempty"`
}

// NewPackageQuarantineRequest instantiates a new PackageQuarantineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageQuarantineRequest() *PackageQuarantineRequest {
	this := PackageQuarantineRequest{}
	return &this
}

// NewPackageQuarantineRequestWithDefaults instantiates a new PackageQuarantineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageQuarantineRequestWithDefaults() *PackageQuarantineRequest {
	this := PackageQuarantineRequest{}
	return &this
}

// GetRestore returns the Restore field value if set, zero value otherwise.
func (o *PackageQuarantineRequest) GetRestore() bool {
	if o == nil || isNil(o.Restore) {
		var ret bool
		return ret
	}
	return *o.Restore
}

// GetRestoreOk returns a tuple with the Restore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageQuarantineRequest) GetRestoreOk() (*bool, bool) {
	if o == nil || isNil(o.Restore) {
		return nil, false
	}
	return o.Restore, true
}

// HasRestore returns a boolean if a field has been set.
func (o *PackageQuarantineRequest) HasRestore() bool {
	if o != nil && !isNil(o.Restore) {
		return true
	}

	return false
}

// SetRestore gets a reference to the given bool and assigns it to the Restore field.
func (o *PackageQuarantineRequest) SetRestore(v bool) {
	o.Restore = &v
}

func (o PackageQuarantineRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Restore) {
		toSerialize["restore"] = o.Restore
	}
	return json.Marshal(toSerialize)
}

type NullablePackageQuarantineRequest struct {
	value *PackageQuarantineRequest
	isSet bool
}

func (v NullablePackageQuarantineRequest) Get() *PackageQuarantineRequest {
	return v.value
}

func (v *NullablePackageQuarantineRequest) Set(val *PackageQuarantineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageQuarantineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageQuarantineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageQuarantineRequest(val *PackageQuarantineRequest) *NullablePackageQuarantineRequest {
	return &NullablePackageQuarantineRequest{value: val, isSet: true}
}

func (v NullablePackageQuarantineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageQuarantineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
