/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.417.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the NugetPackageUploadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NugetPackageUploadRequest{}

// NugetPackageUploadRequest struct for NugetPackageUploadRequest
type NugetPackageUploadRequest struct {
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// Uploads a symbols file as a separate package
	SymbolsFile NullableString `json:"symbols_file,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags NullableString `json:"tags,omitempty"`
}

// NewNugetPackageUploadRequest instantiates a new NugetPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNugetPackageUploadRequest(packageFile string) *NugetPackageUploadRequest {
	this := NugetPackageUploadRequest{}
	this.PackageFile = packageFile
	return &this
}

// NewNugetPackageUploadRequestWithDefaults instantiates a new NugetPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNugetPackageUploadRequestWithDefaults() *NugetPackageUploadRequest {
	this := NugetPackageUploadRequest{}
	return &this
}

// GetPackageFile returns the PackageFile field value
func (o *NugetPackageUploadRequest) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *NugetPackageUploadRequest) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *NugetPackageUploadRequest) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *NugetPackageUploadRequest) GetRepublish() bool {
	if o == nil || IsNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NugetPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || IsNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *NugetPackageUploadRequest) HasRepublish() bool {
	if o != nil && !IsNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *NugetPackageUploadRequest) SetRepublish(v bool) {
	o.Republish = &v
}

// GetSymbolsFile returns the SymbolsFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NugetPackageUploadRequest) GetSymbolsFile() string {
	if o == nil || IsNil(o.SymbolsFile.Get()) {
		var ret string
		return ret
	}
	return *o.SymbolsFile.Get()
}

// GetSymbolsFileOk returns a tuple with the SymbolsFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NugetPackageUploadRequest) GetSymbolsFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SymbolsFile.Get(), o.SymbolsFile.IsSet()
}

// HasSymbolsFile returns a boolean if a field has been set.
func (o *NugetPackageUploadRequest) HasSymbolsFile() bool {
	if o != nil && o.SymbolsFile.IsSet() {
		return true
	}

	return false
}

// SetSymbolsFile gets a reference to the given NullableString and assigns it to the SymbolsFile field.
func (o *NugetPackageUploadRequest) SetSymbolsFile(v string) {
	o.SymbolsFile.Set(&v)
}

// SetSymbolsFileNil sets the value for SymbolsFile to be an explicit nil
func (o *NugetPackageUploadRequest) SetSymbolsFileNil() {
	o.SymbolsFile.Set(nil)
}

// UnsetSymbolsFile ensures that no value is present for SymbolsFile, not even an explicit nil
func (o *NugetPackageUploadRequest) UnsetSymbolsFile() {
	o.SymbolsFile.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NugetPackageUploadRequest) GetTags() string {
	if o == nil || IsNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NugetPackageUploadRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *NugetPackageUploadRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *NugetPackageUploadRequest) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *NugetPackageUploadRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *NugetPackageUploadRequest) UnsetTags() {
	o.Tags.Unset()
}

func (o NugetPackageUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NugetPackageUploadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["package_file"] = o.PackageFile
	if !IsNil(o.Republish) {
		toSerialize["republish"] = o.Republish
	}
	if o.SymbolsFile.IsSet() {
		toSerialize["symbols_file"] = o.SymbolsFile.Get()
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	return toSerialize, nil
}

type NullableNugetPackageUploadRequest struct {
	value *NugetPackageUploadRequest
	isSet bool
}

func (v NullableNugetPackageUploadRequest) Get() *NugetPackageUploadRequest {
	return v.value
}

func (v *NullableNugetPackageUploadRequest) Set(val *NugetPackageUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNugetPackageUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNugetPackageUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNugetPackageUploadRequest(val *NugetPackageUploadRequest) *NullableNugetPackageUploadRequest {
	return &NullableNugetPackageUploadRequest{value: val, isSet: true}
}

func (v NullableNugetPackageUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNugetPackageUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
