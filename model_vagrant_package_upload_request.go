/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.736.13
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"fmt"
)

// checks if the VagrantPackageUploadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VagrantPackageUploadRequest{}

// VagrantPackageUploadRequest struct for VagrantPackageUploadRequest
type VagrantPackageUploadRequest struct {
	// The name of this package.
	Name string `json:"name"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// The virtual machine provider for the box.
	Provider string `json:"provider"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags NullableString `json:"tags,omitempty"`
	// The raw version for this package.
	Version              string `json:"version"`
	AdditionalProperties map[string]interface{}
}

type _VagrantPackageUploadRequest VagrantPackageUploadRequest

// NewVagrantPackageUploadRequest instantiates a new VagrantPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVagrantPackageUploadRequest(name string, packageFile string, provider string, version string) *VagrantPackageUploadRequest {
	this := VagrantPackageUploadRequest{}
	this.Name = name
	this.PackageFile = packageFile
	this.Provider = provider
	this.Version = version
	return &this
}

// NewVagrantPackageUploadRequestWithDefaults instantiates a new VagrantPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVagrantPackageUploadRequestWithDefaults() *VagrantPackageUploadRequest {
	this := VagrantPackageUploadRequest{}
	return &this
}

// GetName returns the Name field value
func (o *VagrantPackageUploadRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VagrantPackageUploadRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VagrantPackageUploadRequest) SetName(v string) {
	o.Name = v
}

// GetPackageFile returns the PackageFile field value
func (o *VagrantPackageUploadRequest) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *VagrantPackageUploadRequest) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *VagrantPackageUploadRequest) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetProvider returns the Provider field value
func (o *VagrantPackageUploadRequest) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *VagrantPackageUploadRequest) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *VagrantPackageUploadRequest) SetProvider(v string) {
	o.Provider = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *VagrantPackageUploadRequest) GetRepublish() bool {
	if o == nil || IsNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VagrantPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || IsNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *VagrantPackageUploadRequest) HasRepublish() bool {
	if o != nil && !IsNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *VagrantPackageUploadRequest) SetRepublish(v bool) {
	o.Republish = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VagrantPackageUploadRequest) GetTags() string {
	if o == nil || IsNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VagrantPackageUploadRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *VagrantPackageUploadRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *VagrantPackageUploadRequest) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *VagrantPackageUploadRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *VagrantPackageUploadRequest) UnsetTags() {
	o.Tags.Unset()
}

// GetVersion returns the Version field value
func (o *VagrantPackageUploadRequest) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *VagrantPackageUploadRequest) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *VagrantPackageUploadRequest) SetVersion(v string) {
	o.Version = v
}

func (o VagrantPackageUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VagrantPackageUploadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["package_file"] = o.PackageFile
	toSerialize["provider"] = o.Provider
	if !IsNil(o.Republish) {
		toSerialize["republish"] = o.Republish
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VagrantPackageUploadRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"package_file",
		"provider",
		"version",
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

	varVagrantPackageUploadRequest := _VagrantPackageUploadRequest{}

	err = json.Unmarshal(data, &varVagrantPackageUploadRequest)

	if err != nil {
		return err
	}

	*o = VagrantPackageUploadRequest(varVagrantPackageUploadRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "package_file")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "republish")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVagrantPackageUploadRequest struct {
	value *VagrantPackageUploadRequest
	isSet bool
}

func (v NullableVagrantPackageUploadRequest) Get() *VagrantPackageUploadRequest {
	return v.value
}

func (v *NullableVagrantPackageUploadRequest) Set(val *VagrantPackageUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVagrantPackageUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVagrantPackageUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVagrantPackageUploadRequest(val *VagrantPackageUploadRequest) *NullableVagrantPackageUploadRequest {
	return &NullableVagrantPackageUploadRequest{value: val, isSet: true}
}

func (v NullableVagrantPackageUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVagrantPackageUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
