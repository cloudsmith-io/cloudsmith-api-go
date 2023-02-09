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

// ErrorDetail struct for ErrorDetail
type ErrorDetail struct {
	// An extended message for the response.
	Detail string `json:"detail"`
}

// NewErrorDetail instantiates a new ErrorDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorDetail(detail string) *ErrorDetail {
	this := ErrorDetail{}
	this.Detail = detail
	return &this
}

// NewErrorDetailWithDefaults instantiates a new ErrorDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorDetailWithDefaults() *ErrorDetail {
	this := ErrorDetail{}
	return &this
}

// GetDetail returns the Detail field value
func (o *ErrorDetail) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *ErrorDetail) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *ErrorDetail) SetDetail(v string) {
	o.Detail = v
}

func (o ErrorDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["detail"] = o.Detail
	}
	return json.Marshal(toSerialize)
}

type NullableErrorDetail struct {
	value *ErrorDetail
	isSet bool
}

func (v NullableErrorDetail) Get() *ErrorDetail {
	return v.value
}

func (v *NullableErrorDetail) Set(val *ErrorDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorDetail(val *ErrorDetail) *NullableErrorDetail {
	return &NullableErrorDetail{value: val, isSet: true}
}

func (v NullableErrorDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
