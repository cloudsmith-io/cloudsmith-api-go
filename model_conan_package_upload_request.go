/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.617.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"fmt"
)

// checks if the ConanPackageUploadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConanPackageUploadRequest{}

// ConanPackageUploadRequest struct for ConanPackageUploadRequest
type ConanPackageUploadRequest struct {
	// Conan channel.
	ConanChannel NullableString `json:"conan_channel,omitempty"`
	// Conan prefix (User).
	ConanPrefix NullableString `json:"conan_prefix,omitempty"`
	// The info file is an python file containing the package metadata.
	InfoFile string `json:"info_file"`
	// The info file is an python file containing the package metadata.
	ManifestFile string `json:"manifest_file"`
	// The conan file is an python file containing the package metadata.
	MetadataFile string `json:"metadata_file"`
	// The name of this package.
	Name NullableString `json:"name,omitempty"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags NullableString `json:"tags,omitempty"`
	// The raw version for this package.
	Version              NullableString `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConanPackageUploadRequest ConanPackageUploadRequest

// NewConanPackageUploadRequest instantiates a new ConanPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConanPackageUploadRequest(infoFile string, manifestFile string, metadataFile string, packageFile string) *ConanPackageUploadRequest {
	this := ConanPackageUploadRequest{}
	this.InfoFile = infoFile
	this.ManifestFile = manifestFile
	this.MetadataFile = metadataFile
	this.PackageFile = packageFile
	return &this
}

// NewConanPackageUploadRequestWithDefaults instantiates a new ConanPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConanPackageUploadRequestWithDefaults() *ConanPackageUploadRequest {
	this := ConanPackageUploadRequest{}
	return &this
}

// GetConanChannel returns the ConanChannel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConanPackageUploadRequest) GetConanChannel() string {
	if o == nil || IsNil(o.ConanChannel.Get()) {
		var ret string
		return ret
	}
	return *o.ConanChannel.Get()
}

// GetConanChannelOk returns a tuple with the ConanChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConanPackageUploadRequest) GetConanChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConanChannel.Get(), o.ConanChannel.IsSet()
}

// HasConanChannel returns a boolean if a field has been set.
func (o *ConanPackageUploadRequest) HasConanChannel() bool {
	if o != nil && o.ConanChannel.IsSet() {
		return true
	}

	return false
}

// SetConanChannel gets a reference to the given NullableString and assigns it to the ConanChannel field.
func (o *ConanPackageUploadRequest) SetConanChannel(v string) {
	o.ConanChannel.Set(&v)
}

// SetConanChannelNil sets the value for ConanChannel to be an explicit nil
func (o *ConanPackageUploadRequest) SetConanChannelNil() {
	o.ConanChannel.Set(nil)
}

// UnsetConanChannel ensures that no value is present for ConanChannel, not even an explicit nil
func (o *ConanPackageUploadRequest) UnsetConanChannel() {
	o.ConanChannel.Unset()
}

// GetConanPrefix returns the ConanPrefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConanPackageUploadRequest) GetConanPrefix() string {
	if o == nil || IsNil(o.ConanPrefix.Get()) {
		var ret string
		return ret
	}
	return *o.ConanPrefix.Get()
}

// GetConanPrefixOk returns a tuple with the ConanPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConanPackageUploadRequest) GetConanPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConanPrefix.Get(), o.ConanPrefix.IsSet()
}

// HasConanPrefix returns a boolean if a field has been set.
func (o *ConanPackageUploadRequest) HasConanPrefix() bool {
	if o != nil && o.ConanPrefix.IsSet() {
		return true
	}

	return false
}

// SetConanPrefix gets a reference to the given NullableString and assigns it to the ConanPrefix field.
func (o *ConanPackageUploadRequest) SetConanPrefix(v string) {
	o.ConanPrefix.Set(&v)
}

// SetConanPrefixNil sets the value for ConanPrefix to be an explicit nil
func (o *ConanPackageUploadRequest) SetConanPrefixNil() {
	o.ConanPrefix.Set(nil)
}

// UnsetConanPrefix ensures that no value is present for ConanPrefix, not even an explicit nil
func (o *ConanPackageUploadRequest) UnsetConanPrefix() {
	o.ConanPrefix.Unset()
}

// GetInfoFile returns the InfoFile field value
func (o *ConanPackageUploadRequest) GetInfoFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InfoFile
}

// GetInfoFileOk returns a tuple with the InfoFile field value
// and a boolean to check if the value has been set.
func (o *ConanPackageUploadRequest) GetInfoFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InfoFile, true
}

// SetInfoFile sets field value
func (o *ConanPackageUploadRequest) SetInfoFile(v string) {
	o.InfoFile = v
}

// GetManifestFile returns the ManifestFile field value
func (o *ConanPackageUploadRequest) GetManifestFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ManifestFile
}

// GetManifestFileOk returns a tuple with the ManifestFile field value
// and a boolean to check if the value has been set.
func (o *ConanPackageUploadRequest) GetManifestFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManifestFile, true
}

// SetManifestFile sets field value
func (o *ConanPackageUploadRequest) SetManifestFile(v string) {
	o.ManifestFile = v
}

// GetMetadataFile returns the MetadataFile field value
func (o *ConanPackageUploadRequest) GetMetadataFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataFile
}

// GetMetadataFileOk returns a tuple with the MetadataFile field value
// and a boolean to check if the value has been set.
func (o *ConanPackageUploadRequest) GetMetadataFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataFile, true
}

// SetMetadataFile sets field value
func (o *ConanPackageUploadRequest) SetMetadataFile(v string) {
	o.MetadataFile = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConanPackageUploadRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConanPackageUploadRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ConanPackageUploadRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ConanPackageUploadRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ConanPackageUploadRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ConanPackageUploadRequest) UnsetName() {
	o.Name.Unset()
}

// GetPackageFile returns the PackageFile field value
func (o *ConanPackageUploadRequest) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *ConanPackageUploadRequest) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *ConanPackageUploadRequest) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *ConanPackageUploadRequest) GetRepublish() bool {
	if o == nil || IsNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConanPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || IsNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *ConanPackageUploadRequest) HasRepublish() bool {
	if o != nil && !IsNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *ConanPackageUploadRequest) SetRepublish(v bool) {
	o.Republish = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConanPackageUploadRequest) GetTags() string {
	if o == nil || IsNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConanPackageUploadRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *ConanPackageUploadRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *ConanPackageUploadRequest) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *ConanPackageUploadRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *ConanPackageUploadRequest) UnsetTags() {
	o.Tags.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConanPackageUploadRequest) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConanPackageUploadRequest) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *ConanPackageUploadRequest) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *ConanPackageUploadRequest) SetVersion(v string) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil
func (o *ConanPackageUploadRequest) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *ConanPackageUploadRequest) UnsetVersion() {
	o.Version.Unset()
}

func (o ConanPackageUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConanPackageUploadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ConanChannel.IsSet() {
		toSerialize["conan_channel"] = o.ConanChannel.Get()
	}
	if o.ConanPrefix.IsSet() {
		toSerialize["conan_prefix"] = o.ConanPrefix.Get()
	}
	toSerialize["info_file"] = o.InfoFile
	toSerialize["manifest_file"] = o.ManifestFile
	toSerialize["metadata_file"] = o.MetadataFile
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["package_file"] = o.PackageFile
	if !IsNil(o.Republish) {
		toSerialize["republish"] = o.Republish
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConanPackageUploadRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"info_file",
		"manifest_file",
		"metadata_file",
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

	varConanPackageUploadRequest := _ConanPackageUploadRequest{}

	err = json.Unmarshal(data, &varConanPackageUploadRequest)

	if err != nil {
		return err
	}

	*o = ConanPackageUploadRequest(varConanPackageUploadRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "conan_channel")
		delete(additionalProperties, "conan_prefix")
		delete(additionalProperties, "info_file")
		delete(additionalProperties, "manifest_file")
		delete(additionalProperties, "metadata_file")
		delete(additionalProperties, "name")
		delete(additionalProperties, "package_file")
		delete(additionalProperties, "republish")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConanPackageUploadRequest struct {
	value *ConanPackageUploadRequest
	isSet bool
}

func (v NullableConanPackageUploadRequest) Get() *ConanPackageUploadRequest {
	return v.value
}

func (v *NullableConanPackageUploadRequest) Set(val *ConanPackageUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConanPackageUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConanPackageUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConanPackageUploadRequest(val *ConanPackageUploadRequest) *NullableConanPackageUploadRequest {
	return &NullableConanPackageUploadRequest{value: val, isSet: true}
}

func (v NullableConanPackageUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConanPackageUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
