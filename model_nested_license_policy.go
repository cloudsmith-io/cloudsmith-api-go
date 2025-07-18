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

// checks if the NestedLicensePolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedLicensePolicy{}

// NestedLicensePolicy struct for NestedLicensePolicy
type NestedLicensePolicy struct {
	AllowUnknownLicenses  *bool          `json:"allow_unknown_licenses,omitempty"`
	CreatedAt             *time.Time     `json:"created_at,omitempty"`
	Description           NullableString `json:"description,omitempty"`
	Name                  *string        `json:"name,omitempty"`
	OnViolationQuarantine *bool          `json:"on_violation_quarantine,omitempty"`
	PackageQueryString    NullableString `json:"package_query_string,omitempty"`
	SlugPerm              *string        `json:"slug_perm,omitempty" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	SpdxIdentifiers       []string       `json:"spdx_identifiers"`
	UpdatedAt             *time.Time     `json:"updated_at,omitempty"`
	Url                   *string        `json:"url,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _NestedLicensePolicy NestedLicensePolicy

// NewNestedLicensePolicy instantiates a new NestedLicensePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedLicensePolicy(spdxIdentifiers []string) *NestedLicensePolicy {
	this := NestedLicensePolicy{}
	this.SpdxIdentifiers = spdxIdentifiers
	return &this
}

// NewNestedLicensePolicyWithDefaults instantiates a new NestedLicensePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedLicensePolicyWithDefaults() *NestedLicensePolicy {
	this := NestedLicensePolicy{}
	return &this
}

// GetAllowUnknownLicenses returns the AllowUnknownLicenses field value if set, zero value otherwise.
func (o *NestedLicensePolicy) GetAllowUnknownLicenses() bool {
	if o == nil || IsNil(o.AllowUnknownLicenses) {
		var ret bool
		return ret
	}
	return *o.AllowUnknownLicenses
}

// GetAllowUnknownLicensesOk returns a tuple with the AllowUnknownLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedLicensePolicy) GetAllowUnknownLicensesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowUnknownLicenses) {
		return nil, false
	}
	return o.AllowUnknownLicenses, true
}

// HasAllowUnknownLicenses returns a boolean if a field has been set.
func (o *NestedLicensePolicy) HasAllowUnknownLicenses() bool {
	if o != nil && !IsNil(o.AllowUnknownLicenses) {
		return true
	}

	return false
}

// SetAllowUnknownLicenses gets a reference to the given bool and assigns it to the AllowUnknownLicenses field.
func (o *NestedLicensePolicy) SetAllowUnknownLicenses(v bool) {
	o.AllowUnknownLicenses = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NestedLicensePolicy) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedLicensePolicy) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NestedLicensePolicy) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NestedLicensePolicy) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedLicensePolicy) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedLicensePolicy) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *NestedLicensePolicy) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *NestedLicensePolicy) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *NestedLicensePolicy) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *NestedLicensePolicy) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NestedLicensePolicy) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedLicensePolicy) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NestedLicensePolicy) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NestedLicensePolicy) SetName(v string) {
	o.Name = &v
}

// GetOnViolationQuarantine returns the OnViolationQuarantine field value if set, zero value otherwise.
func (o *NestedLicensePolicy) GetOnViolationQuarantine() bool {
	if o == nil || IsNil(o.OnViolationQuarantine) {
		var ret bool
		return ret
	}
	return *o.OnViolationQuarantine
}

// GetOnViolationQuarantineOk returns a tuple with the OnViolationQuarantine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedLicensePolicy) GetOnViolationQuarantineOk() (*bool, bool) {
	if o == nil || IsNil(o.OnViolationQuarantine) {
		return nil, false
	}
	return o.OnViolationQuarantine, true
}

// HasOnViolationQuarantine returns a boolean if a field has been set.
func (o *NestedLicensePolicy) HasOnViolationQuarantine() bool {
	if o != nil && !IsNil(o.OnViolationQuarantine) {
		return true
	}

	return false
}

// SetOnViolationQuarantine gets a reference to the given bool and assigns it to the OnViolationQuarantine field.
func (o *NestedLicensePolicy) SetOnViolationQuarantine(v bool) {
	o.OnViolationQuarantine = &v
}

// GetPackageQueryString returns the PackageQueryString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedLicensePolicy) GetPackageQueryString() string {
	if o == nil || IsNil(o.PackageQueryString.Get()) {
		var ret string
		return ret
	}
	return *o.PackageQueryString.Get()
}

// GetPackageQueryStringOk returns a tuple with the PackageQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedLicensePolicy) GetPackageQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageQueryString.Get(), o.PackageQueryString.IsSet()
}

// HasPackageQueryString returns a boolean if a field has been set.
func (o *NestedLicensePolicy) HasPackageQueryString() bool {
	if o != nil && o.PackageQueryString.IsSet() {
		return true
	}

	return false
}

// SetPackageQueryString gets a reference to the given NullableString and assigns it to the PackageQueryString field.
func (o *NestedLicensePolicy) SetPackageQueryString(v string) {
	o.PackageQueryString.Set(&v)
}

// SetPackageQueryStringNil sets the value for PackageQueryString to be an explicit nil
func (o *NestedLicensePolicy) SetPackageQueryStringNil() {
	o.PackageQueryString.Set(nil)
}

// UnsetPackageQueryString ensures that no value is present for PackageQueryString, not even an explicit nil
func (o *NestedLicensePolicy) UnsetPackageQueryString() {
	o.PackageQueryString.Unset()
}

// GetSlugPerm returns the SlugPerm field value if set, zero value otherwise.
func (o *NestedLicensePolicy) GetSlugPerm() string {
	if o == nil || IsNil(o.SlugPerm) {
		var ret string
		return ret
	}
	return *o.SlugPerm
}

// GetSlugPermOk returns a tuple with the SlugPerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedLicensePolicy) GetSlugPermOk() (*string, bool) {
	if o == nil || IsNil(o.SlugPerm) {
		return nil, false
	}
	return o.SlugPerm, true
}

// HasSlugPerm returns a boolean if a field has been set.
func (o *NestedLicensePolicy) HasSlugPerm() bool {
	if o != nil && !IsNil(o.SlugPerm) {
		return true
	}

	return false
}

// SetSlugPerm gets a reference to the given string and assigns it to the SlugPerm field.
func (o *NestedLicensePolicy) SetSlugPerm(v string) {
	o.SlugPerm = &v
}

// GetSpdxIdentifiers returns the SpdxIdentifiers field value
func (o *NestedLicensePolicy) GetSpdxIdentifiers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SpdxIdentifiers
}

// GetSpdxIdentifiersOk returns a tuple with the SpdxIdentifiers field value
// and a boolean to check if the value has been set.
func (o *NestedLicensePolicy) GetSpdxIdentifiersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpdxIdentifiers, true
}

// SetSpdxIdentifiers sets field value
func (o *NestedLicensePolicy) SetSpdxIdentifiers(v []string) {
	o.SpdxIdentifiers = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NestedLicensePolicy) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedLicensePolicy) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NestedLicensePolicy) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NestedLicensePolicy) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NestedLicensePolicy) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedLicensePolicy) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NestedLicensePolicy) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NestedLicensePolicy) SetUrl(v string) {
	o.Url = &v
}

func (o NestedLicensePolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedLicensePolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowUnknownLicenses) {
		toSerialize["allow_unknown_licenses"] = o.AllowUnknownLicenses
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OnViolationQuarantine) {
		toSerialize["on_violation_quarantine"] = o.OnViolationQuarantine
	}
	if o.PackageQueryString.IsSet() {
		toSerialize["package_query_string"] = o.PackageQueryString.Get()
	}
	if !IsNil(o.SlugPerm) {
		toSerialize["slug_perm"] = o.SlugPerm
	}
	toSerialize["spdx_identifiers"] = o.SpdxIdentifiers
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedLicensePolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"spdx_identifiers",
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

	varNestedLicensePolicy := _NestedLicensePolicy{}

	err = json.Unmarshal(data, &varNestedLicensePolicy)

	if err != nil {
		return err
	}

	*o = NestedLicensePolicy(varNestedLicensePolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allow_unknown_licenses")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "on_violation_quarantine")
		delete(additionalProperties, "package_query_string")
		delete(additionalProperties, "slug_perm")
		delete(additionalProperties, "spdx_identifiers")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedLicensePolicy struct {
	value *NestedLicensePolicy
	isSet bool
}

func (v NullableNestedLicensePolicy) Get() *NestedLicensePolicy {
	return v.value
}

func (v *NullableNestedLicensePolicy) Set(val *NestedLicensePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedLicensePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedLicensePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedLicensePolicy(val *NestedLicensePolicy) *NullableNestedLicensePolicy {
	return &NullableNestedLicensePolicy{value: val, isSet: true}
}

func (v NullableNestedLicensePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedLicensePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
