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
)

// checks if the UserBrief type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserBrief{}

// UserBrief struct for UserBrief
type UserBrief struct {
	// If true then you're logged in as a user.
	Authenticated *bool `json:"authenticated,omitempty"`
	// Your email address that we use to contact you. This is only visible to you.
	Email NullableString `json:"email,omitempty"`
	// The full name of the user (if any).
	Name NullableString `json:"name,omitempty"`
	// The URL for the full profile of the user.
	ProfileUrl           NullableString `json:"profile_url,omitempty"`
	SelfUrl              *string        `json:"self_url,omitempty"`
	Slug                 NullableString `json:"slug,omitempty"`
	SlugPerm             NullableString `json:"slug_perm,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserBrief UserBrief

// NewUserBrief instantiates a new UserBrief object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserBrief() *UserBrief {
	this := UserBrief{}
	return &this
}

// NewUserBriefWithDefaults instantiates a new UserBrief object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserBriefWithDefaults() *UserBrief {
	this := UserBrief{}
	return &this
}

// GetAuthenticated returns the Authenticated field value if set, zero value otherwise.
func (o *UserBrief) GetAuthenticated() bool {
	if o == nil || IsNil(o.Authenticated) {
		var ret bool
		return ret
	}
	return *o.Authenticated
}

// GetAuthenticatedOk returns a tuple with the Authenticated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserBrief) GetAuthenticatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Authenticated) {
		return nil, false
	}
	return o.Authenticated, true
}

// HasAuthenticated returns a boolean if a field has been set.
func (o *UserBrief) HasAuthenticated() bool {
	if o != nil && !IsNil(o.Authenticated) {
		return true
	}

	return false
}

// SetAuthenticated gets a reference to the given bool and assigns it to the Authenticated field.
func (o *UserBrief) SetAuthenticated(v bool) {
	o.Authenticated = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserBrief) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserBrief) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UserBrief) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UserBrief) SetEmail(v string) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil
func (o *UserBrief) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UserBrief) UnsetEmail() {
	o.Email.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserBrief) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserBrief) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UserBrief) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UserBrief) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *UserBrief) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UserBrief) UnsetName() {
	o.Name.Unset()
}

// GetProfileUrl returns the ProfileUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserBrief) GetProfileUrl() string {
	if o == nil || IsNil(o.ProfileUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileUrl.Get()
}

// GetProfileUrlOk returns a tuple with the ProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserBrief) GetProfileUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileUrl.Get(), o.ProfileUrl.IsSet()
}

// HasProfileUrl returns a boolean if a field has been set.
func (o *UserBrief) HasProfileUrl() bool {
	if o != nil && o.ProfileUrl.IsSet() {
		return true
	}

	return false
}

// SetProfileUrl gets a reference to the given NullableString and assigns it to the ProfileUrl field.
func (o *UserBrief) SetProfileUrl(v string) {
	o.ProfileUrl.Set(&v)
}

// SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil
func (o *UserBrief) SetProfileUrlNil() {
	o.ProfileUrl.Set(nil)
}

// UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
func (o *UserBrief) UnsetProfileUrl() {
	o.ProfileUrl.Unset()
}

// GetSelfUrl returns the SelfUrl field value if set, zero value otherwise.
func (o *UserBrief) GetSelfUrl() string {
	if o == nil || IsNil(o.SelfUrl) {
		var ret string
		return ret
	}
	return *o.SelfUrl
}

// GetSelfUrlOk returns a tuple with the SelfUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserBrief) GetSelfUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SelfUrl) {
		return nil, false
	}
	return o.SelfUrl, true
}

// HasSelfUrl returns a boolean if a field has been set.
func (o *UserBrief) HasSelfUrl() bool {
	if o != nil && !IsNil(o.SelfUrl) {
		return true
	}

	return false
}

// SetSelfUrl gets a reference to the given string and assigns it to the SelfUrl field.
func (o *UserBrief) SetSelfUrl(v string) {
	o.SelfUrl = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserBrief) GetSlug() string {
	if o == nil || IsNil(o.Slug.Get()) {
		var ret string
		return ret
	}
	return *o.Slug.Get()
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserBrief) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Slug.Get(), o.Slug.IsSet()
}

// HasSlug returns a boolean if a field has been set.
func (o *UserBrief) HasSlug() bool {
	if o != nil && o.Slug.IsSet() {
		return true
	}

	return false
}

// SetSlug gets a reference to the given NullableString and assigns it to the Slug field.
func (o *UserBrief) SetSlug(v string) {
	o.Slug.Set(&v)
}

// SetSlugNil sets the value for Slug to be an explicit nil
func (o *UserBrief) SetSlugNil() {
	o.Slug.Set(nil)
}

// UnsetSlug ensures that no value is present for Slug, not even an explicit nil
func (o *UserBrief) UnsetSlug() {
	o.Slug.Unset()
}

// GetSlugPerm returns the SlugPerm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserBrief) GetSlugPerm() string {
	if o == nil || IsNil(o.SlugPerm.Get()) {
		var ret string
		return ret
	}
	return *o.SlugPerm.Get()
}

// GetSlugPermOk returns a tuple with the SlugPerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserBrief) GetSlugPermOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlugPerm.Get(), o.SlugPerm.IsSet()
}

// HasSlugPerm returns a boolean if a field has been set.
func (o *UserBrief) HasSlugPerm() bool {
	if o != nil && o.SlugPerm.IsSet() {
		return true
	}

	return false
}

// SetSlugPerm gets a reference to the given NullableString and assigns it to the SlugPerm field.
func (o *UserBrief) SetSlugPerm(v string) {
	o.SlugPerm.Set(&v)
}

// SetSlugPermNil sets the value for SlugPerm to be an explicit nil
func (o *UserBrief) SetSlugPermNil() {
	o.SlugPerm.Set(nil)
}

// UnsetSlugPerm ensures that no value is present for SlugPerm, not even an explicit nil
func (o *UserBrief) UnsetSlugPerm() {
	o.SlugPerm.Unset()
}

func (o UserBrief) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserBrief) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authenticated) {
		toSerialize["authenticated"] = o.Authenticated
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ProfileUrl.IsSet() {
		toSerialize["profile_url"] = o.ProfileUrl.Get()
	}
	if !IsNil(o.SelfUrl) {
		toSerialize["self_url"] = o.SelfUrl
	}
	if o.Slug.IsSet() {
		toSerialize["slug"] = o.Slug.Get()
	}
	if o.SlugPerm.IsSet() {
		toSerialize["slug_perm"] = o.SlugPerm.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserBrief) UnmarshalJSON(data []byte) (err error) {
	varUserBrief := _UserBrief{}

	err = json.Unmarshal(data, &varUserBrief)

	if err != nil {
		return err
	}

	*o = UserBrief(varUserBrief)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticated")
		delete(additionalProperties, "email")
		delete(additionalProperties, "name")
		delete(additionalProperties, "profile_url")
		delete(additionalProperties, "self_url")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "slug_perm")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserBrief struct {
	value *UserBrief
	isSet bool
}

func (v NullableUserBrief) Get() *UserBrief {
	return v.value
}

func (v *NullableUserBrief) Set(val *UserBrief) {
	v.value = val
	v.isSet = true
}

func (v NullableUserBrief) IsSet() bool {
	return v.isSet
}

func (v *NullableUserBrief) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserBrief(val *UserBrief) *NullableUserBrief {
	return &NullableUserBrief{value: val, isSet: true}
}

func (v NullableUserBrief) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserBrief) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
