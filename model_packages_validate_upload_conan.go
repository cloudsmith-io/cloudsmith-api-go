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

// PackagesValidateUploadConan struct for PackagesValidateUploadConan
type PackagesValidateUploadConan struct {
	// Conan channel.
	ConanChannel *string `json:"conan_channel,omitempty"`
	// Conan prefix (User).
	ConanPrefix *string `json:"conan_prefix,omitempty"`
	// The info file is an python file containing the package metadata.
	InfoFile string `json:"info_file"`
	// The info file is an python file containing the package metadata.
	ManifestFile string `json:"manifest_file"`
	// The conan file is an python file containing the package metadata.
	MetadataFile string `json:"metadata_file"`
	// The name of this package.
	Name *string `json:"name,omitempty"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags *string `json:"tags,omitempty"`
	// The raw version for this package.
	Version *string `json:"version,omitempty"`
}

// NewPackagesValidateUploadConan instantiates a new PackagesValidateUploadConan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackagesValidateUploadConan(infoFile string, manifestFile string, metadataFile string, packageFile string) *PackagesValidateUploadConan {
	this := PackagesValidateUploadConan{}
	this.InfoFile = infoFile
	this.ManifestFile = manifestFile
	this.MetadataFile = metadataFile
	this.PackageFile = packageFile
	return &this
}

// NewPackagesValidateUploadConanWithDefaults instantiates a new PackagesValidateUploadConan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackagesValidateUploadConanWithDefaults() *PackagesValidateUploadConan {
	this := PackagesValidateUploadConan{}
	return &this
}

// GetConanChannel returns the ConanChannel field value if set, zero value otherwise.
func (o *PackagesValidateUploadConan) GetConanChannel() string {
	if o == nil || o.ConanChannel == nil {
		var ret string
		return ret
	}
	return *o.ConanChannel
}

// GetConanChannelOk returns a tuple with the ConanChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadConan) GetConanChannelOk() (*string, bool) {
	if o == nil || o.ConanChannel == nil {
		return nil, false
	}
	return o.ConanChannel, true
}

// HasConanChannel returns a boolean if a field has been set.
func (o *PackagesValidateUploadConan) HasConanChannel() bool {
	if o != nil && o.ConanChannel != nil {
		return true
	}

	return false
}

// SetConanChannel gets a reference to the given string and assigns it to the ConanChannel field.
func (o *PackagesValidateUploadConan) SetConanChannel(v string) {
	o.ConanChannel = &v
}

// GetConanPrefix returns the ConanPrefix field value if set, zero value otherwise.
func (o *PackagesValidateUploadConan) GetConanPrefix() string {
	if o == nil || o.ConanPrefix == nil {
		var ret string
		return ret
	}
	return *o.ConanPrefix
}

// GetConanPrefixOk returns a tuple with the ConanPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadConan) GetConanPrefixOk() (*string, bool) {
	if o == nil || o.ConanPrefix == nil {
		return nil, false
	}
	return o.ConanPrefix, true
}

// HasConanPrefix returns a boolean if a field has been set.
func (o *PackagesValidateUploadConan) HasConanPrefix() bool {
	if o != nil && o.ConanPrefix != nil {
		return true
	}

	return false
}

// SetConanPrefix gets a reference to the given string and assigns it to the ConanPrefix field.
func (o *PackagesValidateUploadConan) SetConanPrefix(v string) {
	o.ConanPrefix = &v
}

// GetInfoFile returns the InfoFile field value
func (o *PackagesValidateUploadConan) GetInfoFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InfoFile
}

// GetInfoFileOk returns a tuple with the InfoFile field value
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadConan) GetInfoFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InfoFile, true
}

// SetInfoFile sets field value
func (o *PackagesValidateUploadConan) SetInfoFile(v string) {
	o.InfoFile = v
}

// GetManifestFile returns the ManifestFile field value
func (o *PackagesValidateUploadConan) GetManifestFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ManifestFile
}

// GetManifestFileOk returns a tuple with the ManifestFile field value
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadConan) GetManifestFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManifestFile, true
}

// SetManifestFile sets field value
func (o *PackagesValidateUploadConan) SetManifestFile(v string) {
	o.ManifestFile = v
}

// GetMetadataFile returns the MetadataFile field value
func (o *PackagesValidateUploadConan) GetMetadataFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataFile
}

// GetMetadataFileOk returns a tuple with the MetadataFile field value
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadConan) GetMetadataFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataFile, true
}

// SetMetadataFile sets field value
func (o *PackagesValidateUploadConan) SetMetadataFile(v string) {
	o.MetadataFile = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PackagesValidateUploadConan) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadConan) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PackagesValidateUploadConan) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PackagesValidateUploadConan) SetName(v string) {
	o.Name = &v
}

// GetPackageFile returns the PackageFile field value
func (o *PackagesValidateUploadConan) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadConan) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *PackagesValidateUploadConan) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *PackagesValidateUploadConan) GetRepublish() bool {
	if o == nil || o.Republish == nil {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadConan) GetRepublishOk() (*bool, bool) {
	if o == nil || o.Republish == nil {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *PackagesValidateUploadConan) HasRepublish() bool {
	if o != nil && o.Republish != nil {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *PackagesValidateUploadConan) SetRepublish(v bool) {
	o.Republish = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PackagesValidateUploadConan) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadConan) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PackagesValidateUploadConan) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *PackagesValidateUploadConan) SetTags(v string) {
	o.Tags = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PackagesValidateUploadConan) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadConan) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PackagesValidateUploadConan) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PackagesValidateUploadConan) SetVersion(v string) {
	o.Version = &v
}

func (o PackagesValidateUploadConan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConanChannel != nil {
		toSerialize["conan_channel"] = o.ConanChannel
	}
	if o.ConanPrefix != nil {
		toSerialize["conan_prefix"] = o.ConanPrefix
	}
	if true {
		toSerialize["info_file"] = o.InfoFile
	}
	if true {
		toSerialize["manifest_file"] = o.ManifestFile
	}
	if true {
		toSerialize["metadata_file"] = o.MetadataFile
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
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullablePackagesValidateUploadConan struct {
	value *PackagesValidateUploadConan
	isSet bool
}

func (v NullablePackagesValidateUploadConan) Get() *PackagesValidateUploadConan {
	return v.value
}

func (v *NullablePackagesValidateUploadConan) Set(val *PackagesValidateUploadConan) {
	v.value = val
	v.isSet = true
}

func (v NullablePackagesValidateUploadConan) IsSet() bool {
	return v.isSet
}

func (v *NullablePackagesValidateUploadConan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackagesValidateUploadConan(val *PackagesValidateUploadConan) *NullablePackagesValidateUploadConan {
	return &NullablePackagesValidateUploadConan{value: val, isSet: true}
}

func (v NullablePackagesValidateUploadConan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackagesValidateUploadConan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
