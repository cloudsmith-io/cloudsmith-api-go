/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.206.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// DartPackageUploadRequest struct for DartPackageUploadRequest
type DartPackageUploadRequest struct {
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags NullableString `json:"tags,omitempty"`
}

// NewDartPackageUploadRequest instantiates a new DartPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDartPackageUploadRequest(packageFile string) *DartPackageUploadRequest {
	this := DartPackageUploadRequest{}
	this.PackageFile = packageFile
	return &this
}

// NewDartPackageUploadRequestWithDefaults instantiates a new DartPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDartPackageUploadRequestWithDefaults() *DartPackageUploadRequest {
	this := DartPackageUploadRequest{}
	return &this
}

// GetPackageFile returns the PackageFile field value
func (o *DartPackageUploadRequest) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *DartPackageUploadRequest) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *DartPackageUploadRequest) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *DartPackageUploadRequest) GetRepublish() bool {
	if o == nil || isNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DartPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || isNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *DartPackageUploadRequest) HasRepublish() bool {
	if o != nil && !isNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *DartPackageUploadRequest) SetRepublish(v bool) {
	o.Republish = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DartPackageUploadRequest) GetTags() string {
	if o == nil || isNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DartPackageUploadRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *DartPackageUploadRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *DartPackageUploadRequest) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *DartPackageUploadRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *DartPackageUploadRequest) UnsetTags() {
	o.Tags.Unset()
}

func (o DartPackageUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["package_file"] = o.PackageFile
	}
	if !isNil(o.Republish) {
		toSerialize["republish"] = o.Republish
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDartPackageUploadRequest struct {
	value *DartPackageUploadRequest
	isSet bool
}

func (v NullableDartPackageUploadRequest) Get() *DartPackageUploadRequest {
	return v.value
}

func (v *NullableDartPackageUploadRequest) Set(val *DartPackageUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDartPackageUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDartPackageUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDartPackageUploadRequest(val *DartPackageUploadRequest) *NullableDartPackageUploadRequest {
	return &NullableDartPackageUploadRequest{value: val, isSet: true}
}

func (v NullableDartPackageUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDartPackageUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
