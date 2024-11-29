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

// checks if the PackageDependencies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageDependencies{}

// PackageDependencies struct for PackageDependencies
type PackageDependencies struct {
	Dependencies         []PackageDependency `json:"dependencies"`
	AdditionalProperties map[string]interface{}
}

type _PackageDependencies PackageDependencies

// NewPackageDependencies instantiates a new PackageDependencies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageDependencies(dependencies []PackageDependency) *PackageDependencies {
	this := PackageDependencies{}
	this.Dependencies = dependencies
	return &this
}

// NewPackageDependenciesWithDefaults instantiates a new PackageDependencies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDependenciesWithDefaults() *PackageDependencies {
	this := PackageDependencies{}
	return &this
}

// GetDependencies returns the Dependencies field value
func (o *PackageDependencies) GetDependencies() []PackageDependency {
	if o == nil {
		var ret []PackageDependency
		return ret
	}

	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value
// and a boolean to check if the value has been set.
func (o *PackageDependencies) GetDependenciesOk() ([]PackageDependency, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dependencies, true
}

// SetDependencies sets field value
func (o *PackageDependencies) SetDependencies(v []PackageDependency) {
	o.Dependencies = v
}

func (o PackageDependencies) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageDependencies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dependencies"] = o.Dependencies

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PackageDependencies) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dependencies",
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

	varPackageDependencies := _PackageDependencies{}

	err = json.Unmarshal(data, &varPackageDependencies)

	if err != nil {
		return err
	}

	*o = PackageDependencies(varPackageDependencies)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dependencies")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePackageDependencies struct {
	value *PackageDependencies
	isSet bool
}

func (v NullablePackageDependencies) Get() *PackageDependencies {
	return v.value
}

func (v *NullablePackageDependencies) Set(val *PackageDependencies) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageDependencies) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageDependencies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageDependencies(val *PackageDependencies) *NullablePackageDependencies {
	return &NullablePackageDependencies{value: val, isSet: true}
}

func (v NullablePackageDependencies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageDependencies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
