/*
Cloudsmith API

The API to the Cloudsmith Service

API version: 1.42.3
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// PackagesTag struct for PackagesTag
type PackagesTag struct {
	// None
	Action *string `json:"action,omitempty"`
	// If true, created tags will be immutable. An immutable flag is a tag that cannot be removed from a package.
	IsImmutable *bool `json:"is_immutable,omitempty"`
	// A list of tags to apply the action to. Not required for clears.
	Tags []string `json:"tags,omitempty"`
}

// NewPackagesTag instantiates a new PackagesTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackagesTag() *PackagesTag {
	this := PackagesTag{}
	return &this
}

// NewPackagesTagWithDefaults instantiates a new PackagesTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackagesTagWithDefaults() *PackagesTag {
	this := PackagesTag{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PackagesTag) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesTag) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PackagesTag) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PackagesTag) SetAction(v string) {
	o.Action = &v
}

// GetIsImmutable returns the IsImmutable field value if set, zero value otherwise.
func (o *PackagesTag) GetIsImmutable() bool {
	if o == nil || o.IsImmutable == nil {
		var ret bool
		return ret
	}
	return *o.IsImmutable
}

// GetIsImmutableOk returns a tuple with the IsImmutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesTag) GetIsImmutableOk() (*bool, bool) {
	if o == nil || o.IsImmutable == nil {
		return nil, false
	}
	return o.IsImmutable, true
}

// HasIsImmutable returns a boolean if a field has been set.
func (o *PackagesTag) HasIsImmutable() bool {
	if o != nil && o.IsImmutable != nil {
		return true
	}

	return false
}

// SetIsImmutable gets a reference to the given bool and assigns it to the IsImmutable field.
func (o *PackagesTag) SetIsImmutable(v bool) {
	o.IsImmutable = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PackagesTag) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesTag) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PackagesTag) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *PackagesTag) SetTags(v []string) {
	o.Tags = v
}

func (o PackagesTag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.IsImmutable != nil {
		toSerialize["is_immutable"] = o.IsImmutable
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullablePackagesTag struct {
	value *PackagesTag
	isSet bool
}

func (v NullablePackagesTag) Get() *PackagesTag {
	return v.value
}

func (v *NullablePackagesTag) Set(val *PackagesTag) {
	v.value = val
	v.isSet = true
}

func (v NullablePackagesTag) IsSet() bool {
	return v.isSet
}

func (v *NullablePackagesTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackagesTag(val *PackagesTag) *NullablePackagesTag {
	return &NullablePackagesTag{value: val, isSet: true}
}

func (v NullablePackagesTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackagesTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
