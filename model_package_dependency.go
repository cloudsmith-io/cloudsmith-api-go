/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.237.2
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// PackageDependency struct for PackageDependency
type PackageDependency struct {
	DepType *string `json:"dep_type,omitempty"`
	// The name of the package dependency.
	Name     string         `json:"name"`
	Operator NullableString `json:"operator,omitempty"`
	// The raw dependency version (if any).
	Version NullableString `json:"version,omitempty"`
}

// NewPackageDependency instantiates a new PackageDependency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageDependency(name string) *PackageDependency {
	this := PackageDependency{}
	this.Name = name
	return &this
}

// NewPackageDependencyWithDefaults instantiates a new PackageDependency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDependencyWithDefaults() *PackageDependency {
	this := PackageDependency{}
	return &this
}

// GetDepType returns the DepType field value if set, zero value otherwise.
func (o *PackageDependency) GetDepType() string {
	if o == nil || isNil(o.DepType) {
		var ret string
		return ret
	}
	return *o.DepType
}

// GetDepTypeOk returns a tuple with the DepType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDependency) GetDepTypeOk() (*string, bool) {
	if o == nil || isNil(o.DepType) {
		return nil, false
	}
	return o.DepType, true
}

// HasDepType returns a boolean if a field has been set.
func (o *PackageDependency) HasDepType() bool {
	if o != nil && !isNil(o.DepType) {
		return true
	}

	return false
}

// SetDepType gets a reference to the given string and assigns it to the DepType field.
func (o *PackageDependency) SetDepType(v string) {
	o.DepType = &v
}

// GetName returns the Name field value
func (o *PackageDependency) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PackageDependency) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PackageDependency) SetName(v string) {
	o.Name = v
}

// GetOperator returns the Operator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageDependency) GetOperator() string {
	if o == nil || isNil(o.Operator.Get()) {
		var ret string
		return ret
	}
	return *o.Operator.Get()
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageDependency) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operator.Get(), o.Operator.IsSet()
}

// HasOperator returns a boolean if a field has been set.
func (o *PackageDependency) HasOperator() bool {
	if o != nil && o.Operator.IsSet() {
		return true
	}

	return false
}

// SetOperator gets a reference to the given NullableString and assigns it to the Operator field.
func (o *PackageDependency) SetOperator(v string) {
	o.Operator.Set(&v)
}

// SetOperatorNil sets the value for Operator to be an explicit nil
func (o *PackageDependency) SetOperatorNil() {
	o.Operator.Set(nil)
}

// UnsetOperator ensures that no value is present for Operator, not even an explicit nil
func (o *PackageDependency) UnsetOperator() {
	o.Operator.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageDependency) GetVersion() string {
	if o == nil || isNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageDependency) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *PackageDependency) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *PackageDependency) SetVersion(v string) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil
func (o *PackageDependency) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *PackageDependency) UnsetVersion() {
	o.Version.Unset()
}

func (o PackageDependency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DepType) {
		toSerialize["dep_type"] = o.DepType
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Operator.IsSet() {
		toSerialize["operator"] = o.Operator.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePackageDependency struct {
	value *PackageDependency
	isSet bool
}

func (v NullablePackageDependency) Get() *PackageDependency {
	return v.value
}

func (v *NullablePackageDependency) Set(val *PackageDependency) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageDependency) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageDependency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageDependency(val *PackageDependency) *NullablePackageDependency {
	return &NullablePackageDependency{value: val, isSet: true}
}

func (v NullablePackageDependency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageDependency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
