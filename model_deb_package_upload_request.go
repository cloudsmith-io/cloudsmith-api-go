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

// DebPackageUploadRequest struct for DebPackageUploadRequest
type DebPackageUploadRequest struct {
	// The changes archive containing the changes made to the source and debian packaging files
	ChangesFile NullableString `json:"changes_file,omitempty"`
	// The distribution to store the package for.
	Distribution string `json:"distribution"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// The sources archive containing the source code for the binary
	SourcesFile NullableString `json:"sources_file,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags NullableString `json:"tags,omitempty"`
}

// NewDebPackageUploadRequest instantiates a new DebPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebPackageUploadRequest(distribution string, packageFile string) *DebPackageUploadRequest {
	this := DebPackageUploadRequest{}
	this.Distribution = distribution
	this.PackageFile = packageFile
	return &this
}

// NewDebPackageUploadRequestWithDefaults instantiates a new DebPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebPackageUploadRequestWithDefaults() *DebPackageUploadRequest {
	this := DebPackageUploadRequest{}
	return &this
}

// GetChangesFile returns the ChangesFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebPackageUploadRequest) GetChangesFile() string {
	if o == nil || isNil(o.ChangesFile.Get()) {
		var ret string
		return ret
	}
	return *o.ChangesFile.Get()
}

// GetChangesFileOk returns a tuple with the ChangesFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebPackageUploadRequest) GetChangesFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChangesFile.Get(), o.ChangesFile.IsSet()
}

// HasChangesFile returns a boolean if a field has been set.
func (o *DebPackageUploadRequest) HasChangesFile() bool {
	if o != nil && o.ChangesFile.IsSet() {
		return true
	}

	return false
}

// SetChangesFile gets a reference to the given NullableString and assigns it to the ChangesFile field.
func (o *DebPackageUploadRequest) SetChangesFile(v string) {
	o.ChangesFile.Set(&v)
}

// SetChangesFileNil sets the value for ChangesFile to be an explicit nil
func (o *DebPackageUploadRequest) SetChangesFileNil() {
	o.ChangesFile.Set(nil)
}

// UnsetChangesFile ensures that no value is present for ChangesFile, not even an explicit nil
func (o *DebPackageUploadRequest) UnsetChangesFile() {
	o.ChangesFile.Unset()
}

// GetDistribution returns the Distribution field value
func (o *DebPackageUploadRequest) GetDistribution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value
// and a boolean to check if the value has been set.
func (o *DebPackageUploadRequest) GetDistributionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Distribution, true
}

// SetDistribution sets field value
func (o *DebPackageUploadRequest) SetDistribution(v string) {
	o.Distribution = v
}

// GetPackageFile returns the PackageFile field value
func (o *DebPackageUploadRequest) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *DebPackageUploadRequest) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *DebPackageUploadRequest) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *DebPackageUploadRequest) GetRepublish() bool {
	if o == nil || isNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || isNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *DebPackageUploadRequest) HasRepublish() bool {
	if o != nil && !isNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *DebPackageUploadRequest) SetRepublish(v bool) {
	o.Republish = &v
}

// GetSourcesFile returns the SourcesFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebPackageUploadRequest) GetSourcesFile() string {
	if o == nil || isNil(o.SourcesFile.Get()) {
		var ret string
		return ret
	}
	return *o.SourcesFile.Get()
}

// GetSourcesFileOk returns a tuple with the SourcesFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebPackageUploadRequest) GetSourcesFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourcesFile.Get(), o.SourcesFile.IsSet()
}

// HasSourcesFile returns a boolean if a field has been set.
func (o *DebPackageUploadRequest) HasSourcesFile() bool {
	if o != nil && o.SourcesFile.IsSet() {
		return true
	}

	return false
}

// SetSourcesFile gets a reference to the given NullableString and assigns it to the SourcesFile field.
func (o *DebPackageUploadRequest) SetSourcesFile(v string) {
	o.SourcesFile.Set(&v)
}

// SetSourcesFileNil sets the value for SourcesFile to be an explicit nil
func (o *DebPackageUploadRequest) SetSourcesFileNil() {
	o.SourcesFile.Set(nil)
}

// UnsetSourcesFile ensures that no value is present for SourcesFile, not even an explicit nil
func (o *DebPackageUploadRequest) UnsetSourcesFile() {
	o.SourcesFile.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebPackageUploadRequest) GetTags() string {
	if o == nil || isNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebPackageUploadRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *DebPackageUploadRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *DebPackageUploadRequest) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *DebPackageUploadRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *DebPackageUploadRequest) UnsetTags() {
	o.Tags.Unset()
}

func (o DebPackageUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangesFile.IsSet() {
		toSerialize["changes_file"] = o.ChangesFile.Get()
	}
	if true {
		toSerialize["distribution"] = o.Distribution
	}
	if true {
		toSerialize["package_file"] = o.PackageFile
	}
	if !isNil(o.Republish) {
		toSerialize["republish"] = o.Republish
	}
	if o.SourcesFile.IsSet() {
		toSerialize["sources_file"] = o.SourcesFile.Get()
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDebPackageUploadRequest struct {
	value *DebPackageUploadRequest
	isSet bool
}

func (v NullableDebPackageUploadRequest) Get() *DebPackageUploadRequest {
	return v.value
}

func (v *NullableDebPackageUploadRequest) Set(val *DebPackageUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDebPackageUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDebPackageUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebPackageUploadRequest(val *DebPackageUploadRequest) *NullableDebPackageUploadRequest {
	return &NullableDebPackageUploadRequest{value: val, isSet: true}
}

func (v NullableDebPackageUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebPackageUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
