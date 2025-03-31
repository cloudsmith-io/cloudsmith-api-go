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
	"fmt"
)

// checks if the OrganizationTeam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationTeam{}

// OrganizationTeam struct for OrganizationTeam
type OrganizationTeam struct {
	Description          *string `json:"description,omitempty"`
	Name                 string  `json:"name"`
	Slug                 *string `json:"slug,omitempty" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	SlugPerm             *string `json:"slug_perm,omitempty" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Visibility           *string `json:"visibility,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationTeam OrganizationTeam

// NewOrganizationTeam instantiates a new OrganizationTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationTeam(name string) *OrganizationTeam {
	this := OrganizationTeam{}
	this.Name = name
	var visibility string = "Visible"
	this.Visibility = &visibility
	return &this
}

// NewOrganizationTeamWithDefaults instantiates a new OrganizationTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationTeamWithDefaults() *OrganizationTeam {
	this := OrganizationTeam{}
	var visibility string = "Visible"
	this.Visibility = &visibility
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationTeam) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeam) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationTeam) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationTeam) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *OrganizationTeam) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationTeam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationTeam) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *OrganizationTeam) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeam) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *OrganizationTeam) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *OrganizationTeam) SetSlug(v string) {
	o.Slug = &v
}

// GetSlugPerm returns the SlugPerm field value if set, zero value otherwise.
func (o *OrganizationTeam) GetSlugPerm() string {
	if o == nil || IsNil(o.SlugPerm) {
		var ret string
		return ret
	}
	return *o.SlugPerm
}

// GetSlugPermOk returns a tuple with the SlugPerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeam) GetSlugPermOk() (*string, bool) {
	if o == nil || IsNil(o.SlugPerm) {
		return nil, false
	}
	return o.SlugPerm, true
}

// HasSlugPerm returns a boolean if a field has been set.
func (o *OrganizationTeam) HasSlugPerm() bool {
	if o != nil && !IsNil(o.SlugPerm) {
		return true
	}

	return false
}

// SetSlugPerm gets a reference to the given string and assigns it to the SlugPerm field.
func (o *OrganizationTeam) SetSlugPerm(v string) {
	o.SlugPerm = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *OrganizationTeam) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeam) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *OrganizationTeam) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *OrganizationTeam) SetVisibility(v string) {
	o.Visibility = &v
}

func (o OrganizationTeam) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationTeam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.SlugPerm) {
		toSerialize["slug_perm"] = o.SlugPerm
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationTeam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varOrganizationTeam := _OrganizationTeam{}

	err = json.Unmarshal(data, &varOrganizationTeam)

	if err != nil {
		return err
	}

	*o = OrganizationTeam(varOrganizationTeam)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "slug_perm")
		delete(additionalProperties, "visibility")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationTeam struct {
	value *OrganizationTeam
	isSet bool
}

func (v NullableOrganizationTeam) Get() *OrganizationTeam {
	return v.value
}

func (v *NullableOrganizationTeam) Set(val *OrganizationTeam) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationTeam(val *OrganizationTeam) *NullableOrganizationTeam {
	return &NullableOrganizationTeam{value: val, isSet: true}
}

func (v NullableOrganizationTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
