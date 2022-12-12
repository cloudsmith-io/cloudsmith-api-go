/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.182.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// StatusBasic struct for StatusBasic
type StatusBasic struct {
	// The message describing the state of the API.
	Detail *string `json:"detail,omitempty"`
	// The current version for the Cloudsmith service.
	Version *string `json:"version,omitempty"`
}

// NewStatusBasic instantiates a new StatusBasic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusBasic() *StatusBasic {
	this := StatusBasic{}
	return &this
}

// NewStatusBasicWithDefaults instantiates a new StatusBasic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusBasicWithDefaults() *StatusBasic {
	this := StatusBasic{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *StatusBasic) GetDetail() string {
	if o == nil || isNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusBasic) GetDetailOk() (*string, bool) {
	if o == nil || isNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *StatusBasic) HasDetail() bool {
	if o != nil && !isNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *StatusBasic) SetDetail(v string) {
	o.Detail = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *StatusBasic) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusBasic) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *StatusBasic) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *StatusBasic) SetVersion(v string) {
	o.Version = &v
}

func (o StatusBasic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableStatusBasic struct {
	value *StatusBasic
	isSet bool
}

func (v NullableStatusBasic) Get() *StatusBasic {
	return v.value
}

func (v *NullableStatusBasic) Set(val *StatusBasic) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusBasic) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusBasic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusBasic(val *StatusBasic) *NullableStatusBasic {
	return &NullableStatusBasic{value: val, isSet: true}
}

func (v NullableStatusBasic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusBasic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}