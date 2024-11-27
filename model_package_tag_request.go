/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.566.9
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the PackageTagRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageTagRequest{}

// PackageTagRequest struct for PackageTagRequest
type PackageTagRequest struct {
	Action NullableString `json:"action,omitempty"`
	// If true, created tags will be immutable. An immutable flag is a tag that cannot be removed from a package.
	IsImmutable *bool `json:"is_immutable,omitempty"`
	// A list of tags to apply the action to. Not required for clears.
	Tags []string `json:"tags,omitempty"`
}

// NewPackageTagRequest instantiates a new PackageTagRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageTagRequest() *PackageTagRequest {
	this := PackageTagRequest{}
	var action string = "Add"
	this.Action = *NewNullableString(&action)
	var isImmutable bool = false
	this.IsImmutable = &isImmutable
	return &this
}

// NewPackageTagRequestWithDefaults instantiates a new PackageTagRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageTagRequestWithDefaults() *PackageTagRequest {
	this := PackageTagRequest{}
	var action string = "Add"
	this.Action = *NewNullableString(&action)
	var isImmutable bool = false
	this.IsImmutable = &isImmutable
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageTagRequest) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageTagRequest) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *PackageTagRequest) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *PackageTagRequest) SetAction(v string) {
	o.Action.Set(&v)
}

// SetActionNil sets the value for Action to be an explicit nil
func (o *PackageTagRequest) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *PackageTagRequest) UnsetAction() {
	o.Action.Unset()
}

// GetIsImmutable returns the IsImmutable field value if set, zero value otherwise.
func (o *PackageTagRequest) GetIsImmutable() bool {
	if o == nil || IsNil(o.IsImmutable) {
		var ret bool
		return ret
	}
	return *o.IsImmutable
}

// GetIsImmutableOk returns a tuple with the IsImmutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageTagRequest) GetIsImmutableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsImmutable) {
		return nil, false
	}
	return o.IsImmutable, true
}

// HasIsImmutable returns a boolean if a field has been set.
func (o *PackageTagRequest) HasIsImmutable() bool {
	if o != nil && !IsNil(o.IsImmutable) {
		return true
	}

	return false
}

// SetIsImmutable gets a reference to the given bool and assigns it to the IsImmutable field.
func (o *PackageTagRequest) SetIsImmutable(v bool) {
	o.IsImmutable = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageTagRequest) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageTagRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PackageTagRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *PackageTagRequest) SetTags(v []string) {
	o.Tags = v
}

func (o PackageTagRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageTagRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if !IsNil(o.IsImmutable) {
		toSerialize["is_immutable"] = o.IsImmutable
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullablePackageTagRequest struct {
	value *PackageTagRequest
	isSet bool
}

func (v NullablePackageTagRequest) Get() *PackageTagRequest {
	return v.value
}

func (v *NullablePackageTagRequest) Set(val *PackageTagRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageTagRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageTagRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageTagRequest(val *PackageTagRequest) *NullablePackageTagRequest {
	return &NullablePackageTagRequest{value: val, isSet: true}
}

func (v NullablePackageTagRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageTagRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
