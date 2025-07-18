/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.736.13
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"fmt"
)

// checks if the PackagesGroupsList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackagesGroupsList200Response{}

// PackagesGroupsList200Response struct for PackagesGroupsList200Response
type PackagesGroupsList200Response struct {
	Results              []PackageGroup `json:"results"`
	AdditionalProperties map[string]interface{}
}

type _PackagesGroupsList200Response PackagesGroupsList200Response

// NewPackagesGroupsList200Response instantiates a new PackagesGroupsList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackagesGroupsList200Response(results []PackageGroup) *PackagesGroupsList200Response {
	this := PackagesGroupsList200Response{}
	this.Results = results
	return &this
}

// NewPackagesGroupsList200ResponseWithDefaults instantiates a new PackagesGroupsList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackagesGroupsList200ResponseWithDefaults() *PackagesGroupsList200Response {
	this := PackagesGroupsList200Response{}
	return &this
}

// GetResults returns the Results field value
func (o *PackagesGroupsList200Response) GetResults() []PackageGroup {
	if o == nil {
		var ret []PackageGroup
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PackagesGroupsList200Response) GetResultsOk() ([]PackageGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PackagesGroupsList200Response) SetResults(v []PackageGroup) {
	o.Results = v
}

func (o PackagesGroupsList200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackagesGroupsList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PackagesGroupsList200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"results",
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

	varPackagesGroupsList200Response := _PackagesGroupsList200Response{}

	err = json.Unmarshal(data, &varPackagesGroupsList200Response)

	if err != nil {
		return err
	}

	*o = PackagesGroupsList200Response(varPackagesGroupsList200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePackagesGroupsList200Response struct {
	value *PackagesGroupsList200Response
	isSet bool
}

func (v NullablePackagesGroupsList200Response) Get() *PackagesGroupsList200Response {
	return v.value
}

func (v *NullablePackagesGroupsList200Response) Set(val *PackagesGroupsList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePackagesGroupsList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePackagesGroupsList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackagesGroupsList200Response(val *PackagesGroupsList200Response) *NullablePackagesGroupsList200Response {
	return &NullablePackagesGroupsList200Response{value: val, isSet: true}
}

func (v NullablePackagesGroupsList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackagesGroupsList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
