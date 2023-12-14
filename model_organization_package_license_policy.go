/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.327.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"time"
)

// checks if the OrganizationPackageLicensePolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationPackageLicensePolicy{}

// OrganizationPackageLicensePolicy struct for OrganizationPackageLicensePolicy
type OrganizationPackageLicensePolicy struct {
	AllowUnknownLicenses  *bool          `json:"allow_unknown_licenses,omitempty"`
	CreatedAt             *time.Time     `json:"created_at,omitempty"`
	Description           NullableString `json:"description,omitempty"`
	Name                  string         `json:"name"`
	OnViolationQuarantine *bool          `json:"on_violation_quarantine,omitempty"`
	PackageQueryString    NullableString `json:"package_query_string,omitempty"`
	SlugPerm              *string        `json:"slug_perm,omitempty"`
	SpdxIdentifiers       []string       `json:"spdx_identifiers"`
	UpdatedAt             *time.Time     `json:"updated_at,omitempty"`
}

// NewOrganizationPackageLicensePolicy instantiates a new OrganizationPackageLicensePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationPackageLicensePolicy(name string, spdxIdentifiers []string) *OrganizationPackageLicensePolicy {
	this := OrganizationPackageLicensePolicy{}
	this.Name = name
	this.SpdxIdentifiers = spdxIdentifiers
	return &this
}

// NewOrganizationPackageLicensePolicyWithDefaults instantiates a new OrganizationPackageLicensePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationPackageLicensePolicyWithDefaults() *OrganizationPackageLicensePolicy {
	this := OrganizationPackageLicensePolicy{}
	return &this
}

// GetAllowUnknownLicenses returns the AllowUnknownLicenses field value if set, zero value otherwise.
func (o *OrganizationPackageLicensePolicy) GetAllowUnknownLicenses() bool {
	if o == nil || IsNil(o.AllowUnknownLicenses) {
		var ret bool
		return ret
	}
	return *o.AllowUnknownLicenses
}

// GetAllowUnknownLicensesOk returns a tuple with the AllowUnknownLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationPackageLicensePolicy) GetAllowUnknownLicensesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowUnknownLicenses) {
		return nil, false
	}
	return o.AllowUnknownLicenses, true
}

// HasAllowUnknownLicenses returns a boolean if a field has been set.
func (o *OrganizationPackageLicensePolicy) HasAllowUnknownLicenses() bool {
	if o != nil && !IsNil(o.AllowUnknownLicenses) {
		return true
	}

	return false
}

// SetAllowUnknownLicenses gets a reference to the given bool and assigns it to the AllowUnknownLicenses field.
func (o *OrganizationPackageLicensePolicy) SetAllowUnknownLicenses(v bool) {
	o.AllowUnknownLicenses = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrganizationPackageLicensePolicy) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationPackageLicensePolicy) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrganizationPackageLicensePolicy) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OrganizationPackageLicensePolicy) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationPackageLicensePolicy) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationPackageLicensePolicy) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationPackageLicensePolicy) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *OrganizationPackageLicensePolicy) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *OrganizationPackageLicensePolicy) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *OrganizationPackageLicensePolicy) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value
func (o *OrganizationPackageLicensePolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationPackageLicensePolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationPackageLicensePolicy) SetName(v string) {
	o.Name = v
}

// GetOnViolationQuarantine returns the OnViolationQuarantine field value if set, zero value otherwise.
func (o *OrganizationPackageLicensePolicy) GetOnViolationQuarantine() bool {
	if o == nil || IsNil(o.OnViolationQuarantine) {
		var ret bool
		return ret
	}
	return *o.OnViolationQuarantine
}

// GetOnViolationQuarantineOk returns a tuple with the OnViolationQuarantine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationPackageLicensePolicy) GetOnViolationQuarantineOk() (*bool, bool) {
	if o == nil || IsNil(o.OnViolationQuarantine) {
		return nil, false
	}
	return o.OnViolationQuarantine, true
}

// HasOnViolationQuarantine returns a boolean if a field has been set.
func (o *OrganizationPackageLicensePolicy) HasOnViolationQuarantine() bool {
	if o != nil && !IsNil(o.OnViolationQuarantine) {
		return true
	}

	return false
}

// SetOnViolationQuarantine gets a reference to the given bool and assigns it to the OnViolationQuarantine field.
func (o *OrganizationPackageLicensePolicy) SetOnViolationQuarantine(v bool) {
	o.OnViolationQuarantine = &v
}

// GetPackageQueryString returns the PackageQueryString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationPackageLicensePolicy) GetPackageQueryString() string {
	if o == nil || IsNil(o.PackageQueryString.Get()) {
		var ret string
		return ret
	}
	return *o.PackageQueryString.Get()
}

// GetPackageQueryStringOk returns a tuple with the PackageQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationPackageLicensePolicy) GetPackageQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageQueryString.Get(), o.PackageQueryString.IsSet()
}

// HasPackageQueryString returns a boolean if a field has been set.
func (o *OrganizationPackageLicensePolicy) HasPackageQueryString() bool {
	if o != nil && o.PackageQueryString.IsSet() {
		return true
	}

	return false
}

// SetPackageQueryString gets a reference to the given NullableString and assigns it to the PackageQueryString field.
func (o *OrganizationPackageLicensePolicy) SetPackageQueryString(v string) {
	o.PackageQueryString.Set(&v)
}

// SetPackageQueryStringNil sets the value for PackageQueryString to be an explicit nil
func (o *OrganizationPackageLicensePolicy) SetPackageQueryStringNil() {
	o.PackageQueryString.Set(nil)
}

// UnsetPackageQueryString ensures that no value is present for PackageQueryString, not even an explicit nil
func (o *OrganizationPackageLicensePolicy) UnsetPackageQueryString() {
	o.PackageQueryString.Unset()
}

// GetSlugPerm returns the SlugPerm field value if set, zero value otherwise.
func (o *OrganizationPackageLicensePolicy) GetSlugPerm() string {
	if o == nil || IsNil(o.SlugPerm) {
		var ret string
		return ret
	}
	return *o.SlugPerm
}

// GetSlugPermOk returns a tuple with the SlugPerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationPackageLicensePolicy) GetSlugPermOk() (*string, bool) {
	if o == nil || IsNil(o.SlugPerm) {
		return nil, false
	}
	return o.SlugPerm, true
}

// HasSlugPerm returns a boolean if a field has been set.
func (o *OrganizationPackageLicensePolicy) HasSlugPerm() bool {
	if o != nil && !IsNil(o.SlugPerm) {
		return true
	}

	return false
}

// SetSlugPerm gets a reference to the given string and assigns it to the SlugPerm field.
func (o *OrganizationPackageLicensePolicy) SetSlugPerm(v string) {
	o.SlugPerm = &v
}

// GetSpdxIdentifiers returns the SpdxIdentifiers field value
func (o *OrganizationPackageLicensePolicy) GetSpdxIdentifiers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SpdxIdentifiers
}

// GetSpdxIdentifiersOk returns a tuple with the SpdxIdentifiers field value
// and a boolean to check if the value has been set.
func (o *OrganizationPackageLicensePolicy) GetSpdxIdentifiersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpdxIdentifiers, true
}

// SetSpdxIdentifiers sets field value
func (o *OrganizationPackageLicensePolicy) SetSpdxIdentifiers(v []string) {
	o.SpdxIdentifiers = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrganizationPackageLicensePolicy) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationPackageLicensePolicy) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrganizationPackageLicensePolicy) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OrganizationPackageLicensePolicy) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o OrganizationPackageLicensePolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationPackageLicensePolicy) ToMap() (map[string]interface{}, error) {
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
	toSerialize["name"] = o.Name
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
	return toSerialize, nil
}

type NullableOrganizationPackageLicensePolicy struct {
	value *OrganizationPackageLicensePolicy
	isSet bool
}

func (v NullableOrganizationPackageLicensePolicy) Get() *OrganizationPackageLicensePolicy {
	return v.value
}

func (v *NullableOrganizationPackageLicensePolicy) Set(val *OrganizationPackageLicensePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationPackageLicensePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationPackageLicensePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationPackageLicensePolicy(val *OrganizationPackageLicensePolicy) *NullableOrganizationPackageLicensePolicy {
	return &NullableOrganizationPackageLicensePolicy{value: val, isSet: true}
}

func (v NullableOrganizationPackageLicensePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationPackageLicensePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
