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
	"time"
)

// checks if the PackageDenyPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageDenyPolicy{}

// PackageDenyPolicy struct for PackageDenyPolicy
type PackageDenyPolicy struct {
	Action      *string        `json:"action,omitempty"`
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	Description NullableString `json:"description,omitempty"`
	// Whether this rule is enabled or disabled.
	Enabled *bool          `json:"enabled,omitempty"`
	Name    NullableString `json:"name,omitempty"`
	// Packages that match this query will trigger this deny rule.
	PackageQueryString   string     `json:"package_query_string"`
	SlugPerm             *string    `json:"slug_perm,omitempty" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Status               *string    `json:"status,omitempty"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PackageDenyPolicy PackageDenyPolicy

// NewPackageDenyPolicy instantiates a new PackageDenyPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageDenyPolicy(packageQueryString string) *PackageDenyPolicy {
	this := PackageDenyPolicy{}
	this.PackageQueryString = packageQueryString
	return &this
}

// NewPackageDenyPolicyWithDefaults instantiates a new PackageDenyPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDenyPolicyWithDefaults() *PackageDenyPolicy {
	this := PackageDenyPolicy{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PackageDenyPolicy) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDenyPolicy) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PackageDenyPolicy) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PackageDenyPolicy) SetAction(v string) {
	o.Action = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PackageDenyPolicy) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDenyPolicy) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PackageDenyPolicy) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PackageDenyPolicy) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageDenyPolicy) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageDenyPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PackageDenyPolicy) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PackageDenyPolicy) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PackageDenyPolicy) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PackageDenyPolicy) UnsetDescription() {
	o.Description.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PackageDenyPolicy) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDenyPolicy) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PackageDenyPolicy) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PackageDenyPolicy) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageDenyPolicy) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageDenyPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PackageDenyPolicy) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PackageDenyPolicy) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PackageDenyPolicy) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PackageDenyPolicy) UnsetName() {
	o.Name.Unset()
}

// GetPackageQueryString returns the PackageQueryString field value
func (o *PackageDenyPolicy) GetPackageQueryString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageQueryString
}

// GetPackageQueryStringOk returns a tuple with the PackageQueryString field value
// and a boolean to check if the value has been set.
func (o *PackageDenyPolicy) GetPackageQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageQueryString, true
}

// SetPackageQueryString sets field value
func (o *PackageDenyPolicy) SetPackageQueryString(v string) {
	o.PackageQueryString = v
}

// GetSlugPerm returns the SlugPerm field value if set, zero value otherwise.
func (o *PackageDenyPolicy) GetSlugPerm() string {
	if o == nil || IsNil(o.SlugPerm) {
		var ret string
		return ret
	}
	return *o.SlugPerm
}

// GetSlugPermOk returns a tuple with the SlugPerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDenyPolicy) GetSlugPermOk() (*string, bool) {
	if o == nil || IsNil(o.SlugPerm) {
		return nil, false
	}
	return o.SlugPerm, true
}

// HasSlugPerm returns a boolean if a field has been set.
func (o *PackageDenyPolicy) HasSlugPerm() bool {
	if o != nil && !IsNil(o.SlugPerm) {
		return true
	}

	return false
}

// SetSlugPerm gets a reference to the given string and assigns it to the SlugPerm field.
func (o *PackageDenyPolicy) SetSlugPerm(v string) {
	o.SlugPerm = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PackageDenyPolicy) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDenyPolicy) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PackageDenyPolicy) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PackageDenyPolicy) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PackageDenyPolicy) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDenyPolicy) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PackageDenyPolicy) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PackageDenyPolicy) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PackageDenyPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageDenyPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["package_query_string"] = o.PackageQueryString
	if !IsNil(o.SlugPerm) {
		toSerialize["slug_perm"] = o.SlugPerm
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PackageDenyPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"package_query_string",
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

	varPackageDenyPolicy := _PackageDenyPolicy{}

	err = json.Unmarshal(data, &varPackageDenyPolicy)

	if err != nil {
		return err
	}

	*o = PackageDenyPolicy(varPackageDenyPolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "description")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "name")
		delete(additionalProperties, "package_query_string")
		delete(additionalProperties, "slug_perm")
		delete(additionalProperties, "status")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePackageDenyPolicy struct {
	value *PackageDenyPolicy
	isSet bool
}

func (v NullablePackageDenyPolicy) Get() *PackageDenyPolicy {
	return v.value
}

func (v *NullablePackageDenyPolicy) Set(val *PackageDenyPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageDenyPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageDenyPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageDenyPolicy(val *PackageDenyPolicy) *NullablePackageDenyPolicy {
	return &NullablePackageDenyPolicy{value: val, isSet: true}
}

func (v NullablePackageDenyPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageDenyPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
