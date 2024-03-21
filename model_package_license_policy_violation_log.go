/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.372.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"time"
)

// checks if the PackageLicensePolicyViolationLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageLicensePolicyViolationLog{}

// PackageLicensePolicyViolationLog struct for PackageLicensePolicyViolationLog
type PackageLicensePolicyViolationLog struct {
	EventAt *time.Time           `json:"event_at,omitempty"`
	Package PackageVulnerability `json:"package"`
	Policy  NestedLicensePolicy  `json:"policy"`
	Reasons []string             `json:"reasons"`
}

// NewPackageLicensePolicyViolationLog instantiates a new PackageLicensePolicyViolationLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageLicensePolicyViolationLog(package_ PackageVulnerability, policy NestedLicensePolicy, reasons []string) *PackageLicensePolicyViolationLog {
	this := PackageLicensePolicyViolationLog{}
	this.Package = package_
	this.Policy = policy
	this.Reasons = reasons
	return &this
}

// NewPackageLicensePolicyViolationLogWithDefaults instantiates a new PackageLicensePolicyViolationLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageLicensePolicyViolationLogWithDefaults() *PackageLicensePolicyViolationLog {
	this := PackageLicensePolicyViolationLog{}
	return &this
}

// GetEventAt returns the EventAt field value if set, zero value otherwise.
func (o *PackageLicensePolicyViolationLog) GetEventAt() time.Time {
	if o == nil || IsNil(o.EventAt) {
		var ret time.Time
		return ret
	}
	return *o.EventAt
}

// GetEventAtOk returns a tuple with the EventAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLicensePolicyViolationLog) GetEventAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EventAt) {
		return nil, false
	}
	return o.EventAt, true
}

// HasEventAt returns a boolean if a field has been set.
func (o *PackageLicensePolicyViolationLog) HasEventAt() bool {
	if o != nil && !IsNil(o.EventAt) {
		return true
	}

	return false
}

// SetEventAt gets a reference to the given time.Time and assigns it to the EventAt field.
func (o *PackageLicensePolicyViolationLog) SetEventAt(v time.Time) {
	o.EventAt = &v
}

// GetPackage returns the Package field value
func (o *PackageLicensePolicyViolationLog) GetPackage() PackageVulnerability {
	if o == nil {
		var ret PackageVulnerability
		return ret
	}

	return o.Package
}

// GetPackageOk returns a tuple with the Package field value
// and a boolean to check if the value has been set.
func (o *PackageLicensePolicyViolationLog) GetPackageOk() (*PackageVulnerability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Package, true
}

// SetPackage sets field value
func (o *PackageLicensePolicyViolationLog) SetPackage(v PackageVulnerability) {
	o.Package = v
}

// GetPolicy returns the Policy field value
func (o *PackageLicensePolicyViolationLog) GetPolicy() NestedLicensePolicy {
	if o == nil {
		var ret NestedLicensePolicy
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *PackageLicensePolicyViolationLog) GetPolicyOk() (*NestedLicensePolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *PackageLicensePolicyViolationLog) SetPolicy(v NestedLicensePolicy) {
	o.Policy = v
}

// GetReasons returns the Reasons field value
func (o *PackageLicensePolicyViolationLog) GetReasons() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value
// and a boolean to check if the value has been set.
func (o *PackageLicensePolicyViolationLog) GetReasonsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reasons, true
}

// SetReasons sets field value
func (o *PackageLicensePolicyViolationLog) SetReasons(v []string) {
	o.Reasons = v
}

func (o PackageLicensePolicyViolationLog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageLicensePolicyViolationLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventAt) {
		toSerialize["event_at"] = o.EventAt
	}
	toSerialize["package"] = o.Package
	toSerialize["policy"] = o.Policy
	toSerialize["reasons"] = o.Reasons
	return toSerialize, nil
}

type NullablePackageLicensePolicyViolationLog struct {
	value *PackageLicensePolicyViolationLog
	isSet bool
}

func (v NullablePackageLicensePolicyViolationLog) Get() *PackageLicensePolicyViolationLog {
	return v.value
}

func (v *NullablePackageLicensePolicyViolationLog) Set(val *PackageLicensePolicyViolationLog) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageLicensePolicyViolationLog) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageLicensePolicyViolationLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageLicensePolicyViolationLog(val *PackageLicensePolicyViolationLog) *NullablePackageLicensePolicyViolationLog {
	return &NullablePackageLicensePolicyViolationLog{value: val, isSet: true}
}

func (v NullablePackageLicensePolicyViolationLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageLicensePolicyViolationLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
