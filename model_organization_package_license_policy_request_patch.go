/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.401.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the OrganizationPackageLicensePolicyRequestPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationPackageLicensePolicyRequestPatch{}

// OrganizationPackageLicensePolicyRequestPatch struct for OrganizationPackageLicensePolicyRequestPatch
type OrganizationPackageLicensePolicyRequestPatch struct {
	AllowUnknownLicenses  *bool          `json:"allow_unknown_licenses,omitempty"`
	Description           NullableString `json:"description,omitempty"`
	Name                  *string        `json:"name,omitempty"`
	OnViolationQuarantine *bool          `json:"on_violation_quarantine,omitempty"`
	PackageQueryString    NullableString `json:"package_query_string,omitempty"`
	SpdxIdentifiers       []string       `json:"spdx_identifiers,omitempty"`
}

// NewOrganizationPackageLicensePolicyRequestPatch instantiates a new OrganizationPackageLicensePolicyRequestPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationPackageLicensePolicyRequestPatch() *OrganizationPackageLicensePolicyRequestPatch {
	this := OrganizationPackageLicensePolicyRequestPatch{}
	return &this
}

// NewOrganizationPackageLicensePolicyRequestPatchWithDefaults instantiates a new OrganizationPackageLicensePolicyRequestPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationPackageLicensePolicyRequestPatchWithDefaults() *OrganizationPackageLicensePolicyRequestPatch {
	this := OrganizationPackageLicensePolicyRequestPatch{}
	return &this
}

// GetAllowUnknownLicenses returns the AllowUnknownLicenses field value if set, zero value otherwise.
func (o *OrganizationPackageLicensePolicyRequestPatch) GetAllowUnknownLicenses() bool {
	if o == nil || IsNil(o.AllowUnknownLicenses) {
		var ret bool
		return ret
	}
	return *o.AllowUnknownLicenses
}

// GetAllowUnknownLicensesOk returns a tuple with the AllowUnknownLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationPackageLicensePolicyRequestPatch) GetAllowUnknownLicensesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowUnknownLicenses) {
		return nil, false
	}
	return o.AllowUnknownLicenses, true
}

// HasAllowUnknownLicenses returns a boolean if a field has been set.
func (o *OrganizationPackageLicensePolicyRequestPatch) HasAllowUnknownLicenses() bool {
	if o != nil && !IsNil(o.AllowUnknownLicenses) {
		return true
	}

	return false
}

// SetAllowUnknownLicenses gets a reference to the given bool and assigns it to the AllowUnknownLicenses field.
func (o *OrganizationPackageLicensePolicyRequestPatch) SetAllowUnknownLicenses(v bool) {
	o.AllowUnknownLicenses = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationPackageLicensePolicyRequestPatch) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationPackageLicensePolicyRequestPatch) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationPackageLicensePolicyRequestPatch) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *OrganizationPackageLicensePolicyRequestPatch) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *OrganizationPackageLicensePolicyRequestPatch) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *OrganizationPackageLicensePolicyRequestPatch) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationPackageLicensePolicyRequestPatch) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationPackageLicensePolicyRequestPatch) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationPackageLicensePolicyRequestPatch) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationPackageLicensePolicyRequestPatch) SetName(v string) {
	o.Name = &v
}

// GetOnViolationQuarantine returns the OnViolationQuarantine field value if set, zero value otherwise.
func (o *OrganizationPackageLicensePolicyRequestPatch) GetOnViolationQuarantine() bool {
	if o == nil || IsNil(o.OnViolationQuarantine) {
		var ret bool
		return ret
	}
	return *o.OnViolationQuarantine
}

// GetOnViolationQuarantineOk returns a tuple with the OnViolationQuarantine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationPackageLicensePolicyRequestPatch) GetOnViolationQuarantineOk() (*bool, bool) {
	if o == nil || IsNil(o.OnViolationQuarantine) {
		return nil, false
	}
	return o.OnViolationQuarantine, true
}

// HasOnViolationQuarantine returns a boolean if a field has been set.
func (o *OrganizationPackageLicensePolicyRequestPatch) HasOnViolationQuarantine() bool {
	if o != nil && !IsNil(o.OnViolationQuarantine) {
		return true
	}

	return false
}

// SetOnViolationQuarantine gets a reference to the given bool and assigns it to the OnViolationQuarantine field.
func (o *OrganizationPackageLicensePolicyRequestPatch) SetOnViolationQuarantine(v bool) {
	o.OnViolationQuarantine = &v
}

// GetPackageQueryString returns the PackageQueryString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationPackageLicensePolicyRequestPatch) GetPackageQueryString() string {
	if o == nil || IsNil(o.PackageQueryString.Get()) {
		var ret string
		return ret
	}
	return *o.PackageQueryString.Get()
}

// GetPackageQueryStringOk returns a tuple with the PackageQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationPackageLicensePolicyRequestPatch) GetPackageQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageQueryString.Get(), o.PackageQueryString.IsSet()
}

// HasPackageQueryString returns a boolean if a field has been set.
func (o *OrganizationPackageLicensePolicyRequestPatch) HasPackageQueryString() bool {
	if o != nil && o.PackageQueryString.IsSet() {
		return true
	}

	return false
}

// SetPackageQueryString gets a reference to the given NullableString and assigns it to the PackageQueryString field.
func (o *OrganizationPackageLicensePolicyRequestPatch) SetPackageQueryString(v string) {
	o.PackageQueryString.Set(&v)
}

// SetPackageQueryStringNil sets the value for PackageQueryString to be an explicit nil
func (o *OrganizationPackageLicensePolicyRequestPatch) SetPackageQueryStringNil() {
	o.PackageQueryString.Set(nil)
}

// UnsetPackageQueryString ensures that no value is present for PackageQueryString, not even an explicit nil
func (o *OrganizationPackageLicensePolicyRequestPatch) UnsetPackageQueryString() {
	o.PackageQueryString.Unset()
}

// GetSpdxIdentifiers returns the SpdxIdentifiers field value if set, zero value otherwise.
func (o *OrganizationPackageLicensePolicyRequestPatch) GetSpdxIdentifiers() []string {
	if o == nil || IsNil(o.SpdxIdentifiers) {
		var ret []string
		return ret
	}
	return o.SpdxIdentifiers
}

// GetSpdxIdentifiersOk returns a tuple with the SpdxIdentifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationPackageLicensePolicyRequestPatch) GetSpdxIdentifiersOk() ([]string, bool) {
	if o == nil || IsNil(o.SpdxIdentifiers) {
		return nil, false
	}
	return o.SpdxIdentifiers, true
}

// HasSpdxIdentifiers returns a boolean if a field has been set.
func (o *OrganizationPackageLicensePolicyRequestPatch) HasSpdxIdentifiers() bool {
	if o != nil && !IsNil(o.SpdxIdentifiers) {
		return true
	}

	return false
}

// SetSpdxIdentifiers gets a reference to the given []string and assigns it to the SpdxIdentifiers field.
func (o *OrganizationPackageLicensePolicyRequestPatch) SetSpdxIdentifiers(v []string) {
	o.SpdxIdentifiers = v
}

func (o OrganizationPackageLicensePolicyRequestPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationPackageLicensePolicyRequestPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowUnknownLicenses) {
		toSerialize["allow_unknown_licenses"] = o.AllowUnknownLicenses
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
	if !IsNil(o.SpdxIdentifiers) {
		toSerialize["spdx_identifiers"] = o.SpdxIdentifiers
	}
	return toSerialize, nil
}

type NullableOrganizationPackageLicensePolicyRequestPatch struct {
	value *OrganizationPackageLicensePolicyRequestPatch
	isSet bool
}

func (v NullableOrganizationPackageLicensePolicyRequestPatch) Get() *OrganizationPackageLicensePolicyRequestPatch {
	return v.value
}

func (v *NullableOrganizationPackageLicensePolicyRequestPatch) Set(val *OrganizationPackageLicensePolicyRequestPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationPackageLicensePolicyRequestPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationPackageLicensePolicyRequestPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationPackageLicensePolicyRequestPatch(val *OrganizationPackageLicensePolicyRequestPatch) *NullableOrganizationPackageLicensePolicyRequestPatch {
	return &NullableOrganizationPackageLicensePolicyRequestPatch{value: val, isSet: true}
}

func (v NullableOrganizationPackageLicensePolicyRequestPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationPackageLicensePolicyRequestPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
