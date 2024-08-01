/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.477.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the ComposerPackageUploadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComposerPackageUploadRequest{}

// ComposerPackageUploadRequest struct for ComposerPackageUploadRequest
type ComposerPackageUploadRequest struct {
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags NullableString `json:"tags,omitempty"`
}

// NewComposerPackageUploadRequest instantiates a new ComposerPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComposerPackageUploadRequest(packageFile string) *ComposerPackageUploadRequest {
	this := ComposerPackageUploadRequest{}
	this.PackageFile = packageFile
	return &this
}

// NewComposerPackageUploadRequestWithDefaults instantiates a new ComposerPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComposerPackageUploadRequestWithDefaults() *ComposerPackageUploadRequest {
	this := ComposerPackageUploadRequest{}
	return &this
}

// GetPackageFile returns the PackageFile field value
func (o *ComposerPackageUploadRequest) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *ComposerPackageUploadRequest) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *ComposerPackageUploadRequest) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *ComposerPackageUploadRequest) GetRepublish() bool {
	if o == nil || IsNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposerPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || IsNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *ComposerPackageUploadRequest) HasRepublish() bool {
	if o != nil && !IsNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *ComposerPackageUploadRequest) SetRepublish(v bool) {
	o.Republish = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComposerPackageUploadRequest) GetTags() string {
	if o == nil || IsNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComposerPackageUploadRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *ComposerPackageUploadRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *ComposerPackageUploadRequest) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *ComposerPackageUploadRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *ComposerPackageUploadRequest) UnsetTags() {
	o.Tags.Unset()
}

func (o ComposerPackageUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComposerPackageUploadRequest) ToMap() (map[string]interface{}, error) {
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

type NullableComposerPackageUploadRequest struct {
	value *ComposerPackageUploadRequest
	isSet bool
}

func (v NullableComposerPackageUploadRequest) Get() *ComposerPackageUploadRequest {
	return v.value
}

func (v *NullableComposerPackageUploadRequest) Set(val *ComposerPackageUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableComposerPackageUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableComposerPackageUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComposerPackageUploadRequest(val *ComposerPackageUploadRequest) *NullableComposerPackageUploadRequest {
	return &NullableComposerPackageUploadRequest{value: val, isSet: true}
}

func (v NullableComposerPackageUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComposerPackageUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
