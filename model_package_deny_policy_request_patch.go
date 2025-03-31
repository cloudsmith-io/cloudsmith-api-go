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

// checks if the PackageDenyPolicyRequestPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageDenyPolicyRequestPatch{}

// PackageDenyPolicyRequestPatch struct for PackageDenyPolicyRequestPatch
type PackageDenyPolicyRequestPatch struct {
	Description NullableString `json:"description,omitempty"`
	// Whether this rule is enabled or disabled.
	Enabled *bool          `json:"enabled,omitempty"`
	Name    NullableString `json:"name,omitempty"`
	// Packages that match this query will trigger this deny rule.
	PackageQueryString   *string `json:"package_query_string,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PackageDenyPolicyRequestPatch PackageDenyPolicyRequestPatch

// NewPackageDenyPolicyRequestPatch instantiates a new PackageDenyPolicyRequestPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageDenyPolicyRequestPatch() *PackageDenyPolicyRequestPatch {
	this := PackageDenyPolicyRequestPatch{}
	return &this
}

// NewPackageDenyPolicyRequestPatchWithDefaults instantiates a new PackageDenyPolicyRequestPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDenyPolicyRequestPatchWithDefaults() *PackageDenyPolicyRequestPatch {
	this := PackageDenyPolicyRequestPatch{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageDenyPolicyRequestPatch) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageDenyPolicyRequestPatch) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PackageDenyPolicyRequestPatch) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PackageDenyPolicyRequestPatch) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PackageDenyPolicyRequestPatch) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PackageDenyPolicyRequestPatch) UnsetDescription() {
	o.Description.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PackageDenyPolicyRequestPatch) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDenyPolicyRequestPatch) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PackageDenyPolicyRequestPatch) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PackageDenyPolicyRequestPatch) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageDenyPolicyRequestPatch) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageDenyPolicyRequestPatch) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PackageDenyPolicyRequestPatch) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PackageDenyPolicyRequestPatch) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PackageDenyPolicyRequestPatch) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PackageDenyPolicyRequestPatch) UnsetName() {
	o.Name.Unset()
}

// GetPackageQueryString returns the PackageQueryString field value if set, zero value otherwise.
func (o *PackageDenyPolicyRequestPatch) GetPackageQueryString() string {
	if o == nil || IsNil(o.PackageQueryString) {
		var ret string
		return ret
	}
	return *o.PackageQueryString
}

// GetPackageQueryStringOk returns a tuple with the PackageQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDenyPolicyRequestPatch) GetPackageQueryStringOk() (*string, bool) {
	if o == nil || IsNil(o.PackageQueryString) {
		return nil, false
	}
	return o.PackageQueryString, true
}

// HasPackageQueryString returns a boolean if a field has been set.
func (o *PackageDenyPolicyRequestPatch) HasPackageQueryString() bool {
	if o != nil && !IsNil(o.PackageQueryString) {
		return true
	}

	return false
}

// SetPackageQueryString gets a reference to the given string and assigns it to the PackageQueryString field.
func (o *PackageDenyPolicyRequestPatch) SetPackageQueryString(v string) {
	o.PackageQueryString = &v
}

func (o PackageDenyPolicyRequestPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageDenyPolicyRequestPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.PackageQueryString) {
		toSerialize["package_query_string"] = o.PackageQueryString
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PackageDenyPolicyRequestPatch) UnmarshalJSON(data []byte) (err error) {
	varPackageDenyPolicyRequestPatch := _PackageDenyPolicyRequestPatch{}

	err = json.Unmarshal(data, &varPackageDenyPolicyRequestPatch)

	if err != nil {
		return err
	}

	*o = PackageDenyPolicyRequestPatch(varPackageDenyPolicyRequestPatch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "name")
		delete(additionalProperties, "package_query_string")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePackageDenyPolicyRequestPatch struct {
	value *PackageDenyPolicyRequestPatch
	isSet bool
}

func (v NullablePackageDenyPolicyRequestPatch) Get() *PackageDenyPolicyRequestPatch {
	return v.value
}

func (v *NullablePackageDenyPolicyRequestPatch) Set(val *PackageDenyPolicyRequestPatch) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageDenyPolicyRequestPatch) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageDenyPolicyRequestPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageDenyPolicyRequestPatch(val *PackageDenyPolicyRequestPatch) *NullablePackageDenyPolicyRequestPatch {
	return &NullablePackageDenyPolicyRequestPatch{value: val, isSet: true}
}

func (v NullablePackageDenyPolicyRequestPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageDenyPolicyRequestPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
