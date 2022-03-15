/*
Cloudsmith API

The API to the Cloudsmith Service

API version: 1.40.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// FormatsDistributions struct for FormatsDistributions
type FormatsDistributions struct {
	//
	Name *string `json:"name,omitempty"`
	//
	SelfUrl *string `json:"self_url,omitempty"`
	// The slug identifier for this distribution
	Slug *string `json:"slug,omitempty"`
	//
	Variants *string `json:"variants,omitempty"`
}

// NewFormatsDistributions instantiates a new FormatsDistributions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormatsDistributions() *FormatsDistributions {
	this := FormatsDistributions{}
	return &this
}

// NewFormatsDistributionsWithDefaults instantiates a new FormatsDistributions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormatsDistributionsWithDefaults() *FormatsDistributions {
	this := FormatsDistributions{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FormatsDistributions) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatsDistributions) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FormatsDistributions) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FormatsDistributions) SetName(v string) {
	o.Name = &v
}

// GetSelfUrl returns the SelfUrl field value if set, zero value otherwise.
func (o *FormatsDistributions) GetSelfUrl() string {
	if o == nil || o.SelfUrl == nil {
		var ret string
		return ret
	}
	return *o.SelfUrl
}

// GetSelfUrlOk returns a tuple with the SelfUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatsDistributions) GetSelfUrlOk() (*string, bool) {
	if o == nil || o.SelfUrl == nil {
		return nil, false
	}
	return o.SelfUrl, true
}

// HasSelfUrl returns a boolean if a field has been set.
func (o *FormatsDistributions) HasSelfUrl() bool {
	if o != nil && o.SelfUrl != nil {
		return true
	}

	return false
}

// SetSelfUrl gets a reference to the given string and assigns it to the SelfUrl field.
func (o *FormatsDistributions) SetSelfUrl(v string) {
	o.SelfUrl = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *FormatsDistributions) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatsDistributions) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *FormatsDistributions) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *FormatsDistributions) SetSlug(v string) {
	o.Slug = &v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *FormatsDistributions) GetVariants() string {
	if o == nil || o.Variants == nil {
		var ret string
		return ret
	}
	return *o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatsDistributions) GetVariantsOk() (*string, bool) {
	if o == nil || o.Variants == nil {
		return nil, false
	}
	return o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *FormatsDistributions) HasVariants() bool {
	if o != nil && o.Variants != nil {
		return true
	}

	return false
}

// SetVariants gets a reference to the given string and assigns it to the Variants field.
func (o *FormatsDistributions) SetVariants(v string) {
	o.Variants = &v
}

func (o FormatsDistributions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SelfUrl != nil {
		toSerialize["self_url"] = o.SelfUrl
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Variants != nil {
		toSerialize["variants"] = o.Variants
	}
	return json.Marshal(toSerialize)
}

type NullableFormatsDistributions struct {
	value *FormatsDistributions
	isSet bool
}

func (v NullableFormatsDistributions) Get() *FormatsDistributions {
	return v.value
}

func (v *NullableFormatsDistributions) Set(val *FormatsDistributions) {
	v.value = val
	v.isSet = true
}

func (v NullableFormatsDistributions) IsSet() bool {
	return v.isSet
}

func (v *NullableFormatsDistributions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormatsDistributions(val *FormatsDistributions) *NullableFormatsDistributions {
	return &NullableFormatsDistributions{value: val, isSet: true}
}

func (v NullableFormatsDistributions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormatsDistributions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
