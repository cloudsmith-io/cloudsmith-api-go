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

// checks if the OrganizationGroupSyncStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationGroupSyncStatus{}

// OrganizationGroupSyncStatus struct for OrganizationGroupSyncStatus
type OrganizationGroupSyncStatus struct {
	SamlGroupSyncStatus *bool `json:"saml_group_sync_status,omitempty"`
}

// NewOrganizationGroupSyncStatus instantiates a new OrganizationGroupSyncStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationGroupSyncStatus() *OrganizationGroupSyncStatus {
	this := OrganizationGroupSyncStatus{}
	return &this
}

// NewOrganizationGroupSyncStatusWithDefaults instantiates a new OrganizationGroupSyncStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationGroupSyncStatusWithDefaults() *OrganizationGroupSyncStatus {
	this := OrganizationGroupSyncStatus{}
	return &this
}

// GetSamlGroupSyncStatus returns the SamlGroupSyncStatus field value if set, zero value otherwise.
func (o *OrganizationGroupSyncStatus) GetSamlGroupSyncStatus() bool {
	if o == nil || IsNil(o.SamlGroupSyncStatus) {
		var ret bool
		return ret
	}
	return *o.SamlGroupSyncStatus
}

// GetSamlGroupSyncStatusOk returns a tuple with the SamlGroupSyncStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationGroupSyncStatus) GetSamlGroupSyncStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.SamlGroupSyncStatus) {
		return nil, false
	}
	return o.SamlGroupSyncStatus, true
}

// HasSamlGroupSyncStatus returns a boolean if a field has been set.
func (o *OrganizationGroupSyncStatus) HasSamlGroupSyncStatus() bool {
	if o != nil && !IsNil(o.SamlGroupSyncStatus) {
		return true
	}

	return false
}

// SetSamlGroupSyncStatus gets a reference to the given bool and assigns it to the SamlGroupSyncStatus field.
func (o *OrganizationGroupSyncStatus) SetSamlGroupSyncStatus(v bool) {
	o.SamlGroupSyncStatus = &v
}

func (o OrganizationGroupSyncStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationGroupSyncStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SamlGroupSyncStatus) {
		toSerialize["saml_group_sync_status"] = o.SamlGroupSyncStatus
	}
	return toSerialize, nil
}

type NullableOrganizationGroupSyncStatus struct {
	value *OrganizationGroupSyncStatus
	isSet bool
}

func (v NullableOrganizationGroupSyncStatus) Get() *OrganizationGroupSyncStatus {
	return v.value
}

func (v *NullableOrganizationGroupSyncStatus) Set(val *OrganizationGroupSyncStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationGroupSyncStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationGroupSyncStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationGroupSyncStatus(val *OrganizationGroupSyncStatus) *NullableOrganizationGroupSyncStatus {
	return &NullableOrganizationGroupSyncStatus{value: val, isSet: true}
}

func (v NullableOrganizationGroupSyncStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationGroupSyncStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
