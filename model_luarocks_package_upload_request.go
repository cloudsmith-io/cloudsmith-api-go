/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.566.9
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the LuarocksPackageUploadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LuarocksPackageUploadRequest{}

// LuarocksPackageUploadRequest struct for LuarocksPackageUploadRequest
type LuarocksPackageUploadRequest struct {
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags NullableString `json:"tags,omitempty"`
}

type _LuarocksPackageUploadRequest LuarocksPackageUploadRequest

// NewLuarocksPackageUploadRequest instantiates a new LuarocksPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLuarocksPackageUploadRequest(packageFile string) *LuarocksPackageUploadRequest {
	this := LuarocksPackageUploadRequest{}
	this.PackageFile = packageFile
	return &this
}

// NewLuarocksPackageUploadRequestWithDefaults instantiates a new LuarocksPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLuarocksPackageUploadRequestWithDefaults() *LuarocksPackageUploadRequest {
	this := LuarocksPackageUploadRequest{}
	return &this
}

// GetPackageFile returns the PackageFile field value
func (o *LuarocksPackageUploadRequest) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *LuarocksPackageUploadRequest) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *LuarocksPackageUploadRequest) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *LuarocksPackageUploadRequest) GetRepublish() bool {
	if o == nil || IsNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LuarocksPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || IsNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *LuarocksPackageUploadRequest) HasRepublish() bool {
	if o != nil && !IsNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *LuarocksPackageUploadRequest) SetRepublish(v bool) {
	o.Republish = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LuarocksPackageUploadRequest) GetTags() string {
	if o == nil || IsNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LuarocksPackageUploadRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *LuarocksPackageUploadRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *LuarocksPackageUploadRequest) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *LuarocksPackageUploadRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *LuarocksPackageUploadRequest) UnsetTags() {
	o.Tags.Unset()
}

func (o LuarocksPackageUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LuarocksPackageUploadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["package_file"] = o.PackageFile
	if !IsNil(o.Republish) {
		toSerialize["republish"] = o.Republish
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	return toSerialize, nil
}

func (o *LuarocksPackageUploadRequest) UnmarshalJSON(data []byte) (err error) {
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

	varLuarocksPackageUploadRequest := _LuarocksPackageUploadRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLuarocksPackageUploadRequest)

	if err != nil {
		return err
	}

	*o = LuarocksPackageUploadRequest(varLuarocksPackageUploadRequest)

	return err
}

type NullableLuarocksPackageUploadRequest struct {
	value *LuarocksPackageUploadRequest
	isSet bool
}

func (v NullableLuarocksPackageUploadRequest) Get() *LuarocksPackageUploadRequest {
	return v.value
}

func (v *NullableLuarocksPackageUploadRequest) Set(val *LuarocksPackageUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLuarocksPackageUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLuarocksPackageUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLuarocksPackageUploadRequest(val *LuarocksPackageUploadRequest) *NullableLuarocksPackageUploadRequest {
	return &NullableLuarocksPackageUploadRequest{value: val, isSet: true}
}

func (v NullableLuarocksPackageUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLuarocksPackageUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
