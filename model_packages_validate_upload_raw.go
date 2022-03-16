/*
Cloudsmith API

The API to the Cloudsmith Service

API version: 1.42.3
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// PackagesValidateUploadRaw struct for PackagesValidateUploadRaw
type PackagesValidateUploadRaw struct {
	// A custom content/media (also known as MIME) type to be sent when downloading this file. By default Cloudsmith will attempt to detect the type, but if you need to override it, you can specify it here.
	ContentType *string `json:"content_type,omitempty"`
	// A textual description of this package.
	Description *string `json:"description,omitempty"`
	// The name of this package.
	Name *string `json:"name,omitempty"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A one-liner synopsis of this package.
	Summary *string `json:"summary,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags *string `json:"tags,omitempty"`
	// The raw version for this package.
	Version *string `json:"version,omitempty"`
}

// NewPackagesValidateUploadRaw instantiates a new PackagesValidateUploadRaw object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackagesValidateUploadRaw(packageFile string) *PackagesValidateUploadRaw {
	this := PackagesValidateUploadRaw{}
	this.PackageFile = packageFile
	return &this
}

// NewPackagesValidateUploadRawWithDefaults instantiates a new PackagesValidateUploadRaw object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackagesValidateUploadRawWithDefaults() *PackagesValidateUploadRaw {
	this := PackagesValidateUploadRaw{}
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *PackagesValidateUploadRaw) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadRaw) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *PackagesValidateUploadRaw) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *PackagesValidateUploadRaw) SetContentType(v string) {
	o.ContentType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PackagesValidateUploadRaw) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadRaw) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PackagesValidateUploadRaw) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PackagesValidateUploadRaw) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PackagesValidateUploadRaw) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadRaw) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PackagesValidateUploadRaw) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PackagesValidateUploadRaw) SetName(v string) {
	o.Name = &v
}

// GetPackageFile returns the PackageFile field value
func (o *PackagesValidateUploadRaw) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadRaw) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *PackagesValidateUploadRaw) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *PackagesValidateUploadRaw) GetRepublish() bool {
	if o == nil || o.Republish == nil {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadRaw) GetRepublishOk() (*bool, bool) {
	if o == nil || o.Republish == nil {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *PackagesValidateUploadRaw) HasRepublish() bool {
	if o != nil && o.Republish != nil {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *PackagesValidateUploadRaw) SetRepublish(v bool) {
	o.Republish = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *PackagesValidateUploadRaw) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadRaw) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *PackagesValidateUploadRaw) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *PackagesValidateUploadRaw) SetSummary(v string) {
	o.Summary = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PackagesValidateUploadRaw) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadRaw) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PackagesValidateUploadRaw) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *PackagesValidateUploadRaw) SetTags(v string) {
	o.Tags = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PackagesValidateUploadRaw) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadRaw) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PackagesValidateUploadRaw) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PackagesValidateUploadRaw) SetVersion(v string) {
	o.Version = &v
}

func (o PackagesValidateUploadRaw) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentType != nil {
		toSerialize["content_type"] = o.ContentType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["package_file"] = o.PackageFile
	}
	if o.Republish != nil {
		toSerialize["republish"] = o.Republish
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullablePackagesValidateUploadRaw struct {
	value *PackagesValidateUploadRaw
	isSet bool
}

func (v NullablePackagesValidateUploadRaw) Get() *PackagesValidateUploadRaw {
	return v.value
}

func (v *NullablePackagesValidateUploadRaw) Set(val *PackagesValidateUploadRaw) {
	v.value = val
	v.isSet = true
}

func (v NullablePackagesValidateUploadRaw) IsSet() bool {
	return v.isSet
}

func (v *NullablePackagesValidateUploadRaw) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackagesValidateUploadRaw(val *PackagesValidateUploadRaw) *NullablePackagesValidateUploadRaw {
	return &NullablePackagesValidateUploadRaw{value: val, isSet: true}
}

func (v NullablePackagesValidateUploadRaw) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackagesValidateUploadRaw) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
