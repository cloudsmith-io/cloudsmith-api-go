/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.568.8
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"fmt"
)

// checks if the PackageMoveRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageMoveRequest{}

// PackageMoveRequest struct for PackageMoveRequest
type PackageMoveRequest struct {
	Destination          string `json:"destination"`
	AdditionalProperties map[string]interface{}
}

type _PackageMoveRequest PackageMoveRequest

// NewPackageMoveRequest instantiates a new PackageMoveRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageMoveRequest(destination string) *PackageMoveRequest {
	this := PackageMoveRequest{}
	this.Destination = destination
	return &this
}

// NewPackageMoveRequestWithDefaults instantiates a new PackageMoveRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageMoveRequestWithDefaults() *PackageMoveRequest {
	this := PackageMoveRequest{}
	return &this
}

// GetDestination returns the Destination field value
func (o *PackageMoveRequest) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *PackageMoveRequest) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *PackageMoveRequest) SetDestination(v string) {
	o.Destination = v
}

func (o PackageMoveRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageMoveRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination"] = o.Destination

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PackageMoveRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination",
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

	varPackageMoveRequest := _PackageMoveRequest{}

	err = json.Unmarshal(data, &varPackageMoveRequest)

	if err != nil {
		return err
	}

	*o = PackageMoveRequest(varPackageMoveRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "destination")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePackageMoveRequest struct {
	value *PackageMoveRequest
	isSet bool
}

func (v NullablePackageMoveRequest) Get() *PackageMoveRequest {
	return v.value
}

func (v *NullablePackageMoveRequest) Set(val *PackageMoveRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageMoveRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageMoveRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageMoveRequest(val *PackageMoveRequest) *NullablePackageMoveRequest {
	return &NullablePackageMoveRequest{value: val, isSet: true}
}

func (v NullablePackageMoveRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageMoveRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
