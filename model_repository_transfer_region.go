/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.568.8
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the RepositoryTransferRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryTransferRegion{}

// RepositoryTransferRegion struct for RepositoryTransferRegion
type RepositoryTransferRegion struct {
	// The Cloudsmith region in which package files are stored.
	StorageRegion        *string `json:"storage_region,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RepositoryTransferRegion RepositoryTransferRegion

// NewRepositoryTransferRegion instantiates a new RepositoryTransferRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryTransferRegion() *RepositoryTransferRegion {
	this := RepositoryTransferRegion{}
	var storageRegion string = "default"
	this.StorageRegion = &storageRegion
	return &this
}

// NewRepositoryTransferRegionWithDefaults instantiates a new RepositoryTransferRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryTransferRegionWithDefaults() *RepositoryTransferRegion {
	this := RepositoryTransferRegion{}
	var storageRegion string = "default"
	this.StorageRegion = &storageRegion
	return &this
}

// GetStorageRegion returns the StorageRegion field value if set, zero value otherwise.
func (o *RepositoryTransferRegion) GetStorageRegion() string {
	if o == nil || IsNil(o.StorageRegion) {
		var ret string
		return ret
	}
	return *o.StorageRegion
}

// GetStorageRegionOk returns a tuple with the StorageRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryTransferRegion) GetStorageRegionOk() (*string, bool) {
	if o == nil || IsNil(o.StorageRegion) {
		return nil, false
	}
	return o.StorageRegion, true
}

// HasStorageRegion returns a boolean if a field has been set.
func (o *RepositoryTransferRegion) HasStorageRegion() bool {
	if o != nil && !IsNil(o.StorageRegion) {
		return true
	}

	return false
}

// SetStorageRegion gets a reference to the given string and assigns it to the StorageRegion field.
func (o *RepositoryTransferRegion) SetStorageRegion(v string) {
	o.StorageRegion = &v
}

func (o RepositoryTransferRegion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryTransferRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StorageRegion) {
		toSerialize["storage_region"] = o.StorageRegion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RepositoryTransferRegion) UnmarshalJSON(data []byte) (err error) {
	varRepositoryTransferRegion := _RepositoryTransferRegion{}

	err = json.Unmarshal(data, &varRepositoryTransferRegion)

	if err != nil {
		return err
	}

	*o = RepositoryTransferRegion(varRepositoryTransferRegion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "storage_region")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRepositoryTransferRegion struct {
	value *RepositoryTransferRegion
	isSet bool
}

func (v NullableRepositoryTransferRegion) Get() *RepositoryTransferRegion {
	return v.value
}

func (v *NullableRepositoryTransferRegion) Set(val *RepositoryTransferRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryTransferRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryTransferRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryTransferRegion(val *RepositoryTransferRegion) *NullableRepositoryTransferRegion {
	return &NullableRepositoryTransferRegion{value: val, isSet: true}
}

func (v NullableRepositoryTransferRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryTransferRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
