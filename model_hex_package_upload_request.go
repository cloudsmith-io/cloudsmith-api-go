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
	"fmt"
)

// checks if the HexPackageUploadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HexPackageUploadRequest{}

// HexPackageUploadRequest struct for HexPackageUploadRequest
type HexPackageUploadRequest struct {
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags                 NullableString `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HexPackageUploadRequest HexPackageUploadRequest

// NewHexPackageUploadRequest instantiates a new HexPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHexPackageUploadRequest(packageFile string) *HexPackageUploadRequest {
	this := HexPackageUploadRequest{}
	this.PackageFile = packageFile
	return &this
}

// NewHexPackageUploadRequestWithDefaults instantiates a new HexPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHexPackageUploadRequestWithDefaults() *HexPackageUploadRequest {
	this := HexPackageUploadRequest{}
	return &this
}

// GetPackageFile returns the PackageFile field value
func (o *HexPackageUploadRequest) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *HexPackageUploadRequest) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *HexPackageUploadRequest) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *HexPackageUploadRequest) GetRepublish() bool {
	if o == nil || IsNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HexPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || IsNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *HexPackageUploadRequest) HasRepublish() bool {
	if o != nil && !IsNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *HexPackageUploadRequest) SetRepublish(v bool) {
	o.Republish = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HexPackageUploadRequest) GetTags() string {
	if o == nil || IsNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HexPackageUploadRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *HexPackageUploadRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *HexPackageUploadRequest) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *HexPackageUploadRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *HexPackageUploadRequest) UnsetTags() {
	o.Tags.Unset()
}

func (o HexPackageUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HexPackageUploadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["package_file"] = o.PackageFile
	if !IsNil(o.Republish) {
		toSerialize["republish"] = o.Republish
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HexPackageUploadRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"package_file",
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

	varHexPackageUploadRequest := _HexPackageUploadRequest{}

	err = json.Unmarshal(data, &varHexPackageUploadRequest)

	if err != nil {
		return err
	}

	*o = HexPackageUploadRequest(varHexPackageUploadRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "package_file")
		delete(additionalProperties, "republish")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHexPackageUploadRequest struct {
	value *HexPackageUploadRequest
	isSet bool
}

func (v NullableHexPackageUploadRequest) Get() *HexPackageUploadRequest {
	return v.value
}

func (v *NullableHexPackageUploadRequest) Set(val *HexPackageUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHexPackageUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHexPackageUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHexPackageUploadRequest(val *HexPackageUploadRequest) *NullableHexPackageUploadRequest {
	return &NullableHexPackageUploadRequest{value: val, isSet: true}
}

func (v NullableHexPackageUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHexPackageUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
