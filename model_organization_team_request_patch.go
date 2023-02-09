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

// OrganizationTeamRequestPatch struct for OrganizationTeamRequestPatch
type OrganizationTeamRequestPatch struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	Visibility  *string `json:"visibility,omitempty"`
}

// NewOrganizationTeamRequestPatch instantiates a new OrganizationTeamRequestPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationTeamRequestPatch() *OrganizationTeamRequestPatch {
	this := OrganizationTeamRequestPatch{}
	var visibility string = "Visible"
	this.Visibility = &visibility
	return &this
}

// NewOrganizationTeamRequestPatchWithDefaults instantiates a new OrganizationTeamRequestPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationTeamRequestPatchWithDefaults() *OrganizationTeamRequestPatch {
	this := OrganizationTeamRequestPatch{}
	var visibility string = "Visible"
	this.Visibility = &visibility
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationTeamRequestPatch) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeamRequestPatch) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationTeamRequestPatch) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationTeamRequestPatch) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationTeamRequestPatch) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeamRequestPatch) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationTeamRequestPatch) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationTeamRequestPatch) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *OrganizationTeamRequestPatch) GetSlug() string {
	if o == nil || isNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeamRequestPatch) GetSlugOk() (*string, bool) {
	if o == nil || isNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *OrganizationTeamRequestPatch) HasSlug() bool {
	if o != nil && !isNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *OrganizationTeamRequestPatch) SetSlug(v string) {
	o.Slug = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *OrganizationTeamRequestPatch) GetVisibility() string {
	if o == nil || isNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeamRequestPatch) GetVisibilityOk() (*string, bool) {
	if o == nil || isNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *OrganizationTeamRequestPatch) HasVisibility() bool {
	if o != nil && !isNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *OrganizationTeamRequestPatch) SetVisibility(v string) {
	o.Visibility = &v
}

func (o OrganizationTeamRequestPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !isNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationTeamRequestPatch struct {
	value *OrganizationTeamRequestPatch
	isSet bool
}

func (v NullableOrganizationTeamRequestPatch) Get() *OrganizationTeamRequestPatch {
	return v.value
}

func (v *NullableOrganizationTeamRequestPatch) Set(val *OrganizationTeamRequestPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationTeamRequestPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationTeamRequestPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationTeamRequestPatch(val *OrganizationTeamRequestPatch) *NullableOrganizationTeamRequestPatch {
	return &NullableOrganizationTeamRequestPatch{value: val, isSet: true}
}

func (v NullableOrganizationTeamRequestPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationTeamRequestPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
