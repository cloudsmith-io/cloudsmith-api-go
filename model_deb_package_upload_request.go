/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.617.2
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"fmt"
)

// checks if the DebPackageUploadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DebPackageUploadRequest{}

// DebPackageUploadRequest struct for DebPackageUploadRequest
type DebPackageUploadRequest struct {
	// The changes archive containing the changes made to the source and debian packaging files
	ChangesFile NullableString `json:"changes_file,omitempty"`
	// The component (channel) for the package (e.g. 'main', 'unstable', etc.)
	Component *string `json:"component,omitempty" validate:"regexp=^[-_.\\\\w]+$"`
	// The distribution to store the package for.
	Distribution string `json:"distribution"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// The sources archive containing the source code for the binary
	SourcesFile NullableString `json:"sources_file,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags                 NullableString `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DebPackageUploadRequest DebPackageUploadRequest

// NewDebPackageUploadRequest instantiates a new DebPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebPackageUploadRequest(distribution string, packageFile string) *DebPackageUploadRequest {
	this := DebPackageUploadRequest{}
	var component string = "main"
	this.Component = &component
	this.Distribution = distribution
	this.PackageFile = packageFile
	return &this
}

// NewDebPackageUploadRequestWithDefaults instantiates a new DebPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebPackageUploadRequestWithDefaults() *DebPackageUploadRequest {
	this := DebPackageUploadRequest{}
	var component string = "main"
	this.Component = &component
	return &this
}

// GetChangesFile returns the ChangesFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebPackageUploadRequest) GetChangesFile() string {
	if o == nil || IsNil(o.ChangesFile.Get()) {
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

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *DebPackageUploadRequest) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebPackageUploadRequest) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *DebPackageUploadRequest) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *DebPackageUploadRequest) SetComponent(v string) {
	o.Component = &v
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
	if o == nil || IsNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || IsNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *DebPackageUploadRequest) HasRepublish() bool {
	if o != nil && !IsNil(o.Republish) {
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
	if o == nil || IsNil(o.SourcesFile.Get()) {
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
	if o == nil || IsNil(o.Tags.Get()) {
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DebPackageUploadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangesFile.IsSet() {
		toSerialize["changes_file"] = o.ChangesFile.Get()
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	toSerialize["distribution"] = o.Distribution
	toSerialize["package_file"] = o.PackageFile
	if !IsNil(o.Republish) {
		toSerialize["republish"] = o.Republish
	}
	if o.SourcesFile.IsSet() {
		toSerialize["sources_file"] = o.SourcesFile.Get()
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DebPackageUploadRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"distribution",
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

	varDebPackageUploadRequest := _DebPackageUploadRequest{}

	err = json.Unmarshal(data, &varDebPackageUploadRequest)

	if err != nil {
		return err
	}

	*o = DebPackageUploadRequest(varDebPackageUploadRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "changes_file")
		delete(additionalProperties, "component")
		delete(additionalProperties, "distribution")
		delete(additionalProperties, "package_file")
		delete(additionalProperties, "republish")
		delete(additionalProperties, "sources_file")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
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
