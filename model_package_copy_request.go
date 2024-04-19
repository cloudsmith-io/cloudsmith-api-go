/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.401.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the PackageCopyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageCopyRequest{}

// PackageCopyRequest struct for PackageCopyRequest
type PackageCopyRequest struct {
	Destination string `json:"destination"`
	// If true, the package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
}

// NewPackageCopyRequest instantiates a new PackageCopyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageCopyRequest(destination string) *PackageCopyRequest {
	this := PackageCopyRequest{}
	this.Destination = destination
	return &this
}

// NewPackageCopyRequestWithDefaults instantiates a new PackageCopyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageCopyRequestWithDefaults() *PackageCopyRequest {
	this := PackageCopyRequest{}
	return &this
}

// GetDestination returns the Destination field value
func (o *PackageCopyRequest) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *PackageCopyRequest) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *PackageCopyRequest) SetDestination(v string) {
	o.Destination = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *PackageCopyRequest) GetRepublish() bool {
	if o == nil || IsNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageCopyRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || IsNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *PackageCopyRequest) HasRepublish() bool {
	if o != nil && !IsNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *PackageCopyRequest) SetRepublish(v bool) {
	o.Republish = &v
}

func (o PackageCopyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageCopyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination"] = o.Destination
	if !IsNil(o.Republish) {
		toSerialize["republish"] = o.Republish
	}
	return toSerialize, nil
}

type NullablePackageCopyRequest struct {
	value *PackageCopyRequest
	isSet bool
}

func (v NullablePackageCopyRequest) Get() *PackageCopyRequest {
	return v.value
}

func (v *NullablePackageCopyRequest) Set(val *PackageCopyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageCopyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageCopyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageCopyRequest(val *PackageCopyRequest) *NullablePackageCopyRequest {
	return &NullablePackageCopyRequest{value: val, isSet: true}
}

func (v NullablePackageCopyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageCopyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
