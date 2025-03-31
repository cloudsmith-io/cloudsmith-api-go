/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.654.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the DistributionVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DistributionVersion{}

// DistributionVersion A list of the versions for this distribution
type DistributionVersion struct {
	// The textual name for this version.
	Name *string `json:"name,omitempty"`
	// The slug identifier for this version
	Slug                 *string `json:"slug,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DistributionVersion DistributionVersion

// NewDistributionVersion instantiates a new DistributionVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributionVersion() *DistributionVersion {
	this := DistributionVersion{}
	return &this
}

// NewDistributionVersionWithDefaults instantiates a new DistributionVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionVersionWithDefaults() *DistributionVersion {
	this := DistributionVersion{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DistributionVersion) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionVersion) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DistributionVersion) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DistributionVersion) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *DistributionVersion) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionVersion) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *DistributionVersion) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *DistributionVersion) SetSlug(v string) {
	o.Slug = &v
}

func (o DistributionVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DistributionVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DistributionVersion) UnmarshalJSON(data []byte) (err error) {
	varDistributionVersion := _DistributionVersion{}

	err = json.Unmarshal(data, &varDistributionVersion)

	if err != nil {
		return err
	}

	*o = DistributionVersion(varDistributionVersion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDistributionVersion struct {
	value *DistributionVersion
	isSet bool
}

func (v NullableDistributionVersion) Get() *DistributionVersion {
	return v.value
}

func (v *NullableDistributionVersion) Set(val *DistributionVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionVersion(val *DistributionVersion) *NullableDistributionVersion {
	return &NullableDistributionVersion{value: val, isSet: true}
}

func (v NullableDistributionVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
