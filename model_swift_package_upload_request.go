/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.568.8
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"fmt"
)

// checks if the SwiftPackageUploadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwiftPackageUploadRequest{}

// SwiftPackageUploadRequest struct for SwiftPackageUploadRequest
type SwiftPackageUploadRequest struct {
	// The name of the author of the package.
	AuthorName *string `json:"author_name,omitempty"`
	// The organization of the author.
	AuthorOrg *string `json:"author_org,omitempty"`
	// The license URL of this package.
	LicenseUrl NullableString `json:"license_url,omitempty"`
	// The name of this package.
	Name string `json:"name"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// The URL of the readme for the package.
	ReadmeUrl *string `json:"readme_url,omitempty"`
	// The URL of the SCM repository for the package.
	RepositoryUrl *string `json:"repository_url,omitempty"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A scope provides a namespace for related packages within the package registry.
	Scope string `json:"scope"`
	// A comma-separated values list of tags to add to the package.
	Tags NullableString `json:"tags,omitempty"`
	// The raw version for this package.
	Version              string `json:"version"`
	AdditionalProperties map[string]interface{}
}

type _SwiftPackageUploadRequest SwiftPackageUploadRequest

// NewSwiftPackageUploadRequest instantiates a new SwiftPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwiftPackageUploadRequest(name string, packageFile string, scope string, version string) *SwiftPackageUploadRequest {
	this := SwiftPackageUploadRequest{}
	this.Name = name
	this.PackageFile = packageFile
	this.Scope = scope
	this.Version = version
	return &this
}

// NewSwiftPackageUploadRequestWithDefaults instantiates a new SwiftPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwiftPackageUploadRequestWithDefaults() *SwiftPackageUploadRequest {
	this := SwiftPackageUploadRequest{}
	return &this
}

// GetAuthorName returns the AuthorName field value if set, zero value otherwise.
func (o *SwiftPackageUploadRequest) GetAuthorName() string {
	if o == nil || IsNil(o.AuthorName) {
		var ret string
		return ret
	}
	return *o.AuthorName
}

// GetAuthorNameOk returns a tuple with the AuthorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwiftPackageUploadRequest) GetAuthorNameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorName) {
		return nil, false
	}
	return o.AuthorName, true
}

// HasAuthorName returns a boolean if a field has been set.
func (o *SwiftPackageUploadRequest) HasAuthorName() bool {
	if o != nil && !IsNil(o.AuthorName) {
		return true
	}

	return false
}

// SetAuthorName gets a reference to the given string and assigns it to the AuthorName field.
func (o *SwiftPackageUploadRequest) SetAuthorName(v string) {
	o.AuthorName = &v
}

// GetAuthorOrg returns the AuthorOrg field value if set, zero value otherwise.
func (o *SwiftPackageUploadRequest) GetAuthorOrg() string {
	if o == nil || IsNil(o.AuthorOrg) {
		var ret string
		return ret
	}
	return *o.AuthorOrg
}

// GetAuthorOrgOk returns a tuple with the AuthorOrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwiftPackageUploadRequest) GetAuthorOrgOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorOrg) {
		return nil, false
	}
	return o.AuthorOrg, true
}

// HasAuthorOrg returns a boolean if a field has been set.
func (o *SwiftPackageUploadRequest) HasAuthorOrg() bool {
	if o != nil && !IsNil(o.AuthorOrg) {
		return true
	}

	return false
}

// SetAuthorOrg gets a reference to the given string and assigns it to the AuthorOrg field.
func (o *SwiftPackageUploadRequest) SetAuthorOrg(v string) {
	o.AuthorOrg = &v
}

// GetLicenseUrl returns the LicenseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SwiftPackageUploadRequest) GetLicenseUrl() string {
	if o == nil || IsNil(o.LicenseUrl.Get()) {
		var ret string
		return ret
	}
	return *o.LicenseUrl.Get()
}

// GetLicenseUrlOk returns a tuple with the LicenseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SwiftPackageUploadRequest) GetLicenseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LicenseUrl.Get(), o.LicenseUrl.IsSet()
}

// HasLicenseUrl returns a boolean if a field has been set.
func (o *SwiftPackageUploadRequest) HasLicenseUrl() bool {
	if o != nil && o.LicenseUrl.IsSet() {
		return true
	}

	return false
}

// SetLicenseUrl gets a reference to the given NullableString and assigns it to the LicenseUrl field.
func (o *SwiftPackageUploadRequest) SetLicenseUrl(v string) {
	o.LicenseUrl.Set(&v)
}

// SetLicenseUrlNil sets the value for LicenseUrl to be an explicit nil
func (o *SwiftPackageUploadRequest) SetLicenseUrlNil() {
	o.LicenseUrl.Set(nil)
}

// UnsetLicenseUrl ensures that no value is present for LicenseUrl, not even an explicit nil
func (o *SwiftPackageUploadRequest) UnsetLicenseUrl() {
	o.LicenseUrl.Unset()
}

// GetName returns the Name field value
func (o *SwiftPackageUploadRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SwiftPackageUploadRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SwiftPackageUploadRequest) SetName(v string) {
	o.Name = v
}

// GetPackageFile returns the PackageFile field value
func (o *SwiftPackageUploadRequest) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *SwiftPackageUploadRequest) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *SwiftPackageUploadRequest) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetReadmeUrl returns the ReadmeUrl field value if set, zero value otherwise.
func (o *SwiftPackageUploadRequest) GetReadmeUrl() string {
	if o == nil || IsNil(o.ReadmeUrl) {
		var ret string
		return ret
	}
	return *o.ReadmeUrl
}

// GetReadmeUrlOk returns a tuple with the ReadmeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwiftPackageUploadRequest) GetReadmeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReadmeUrl) {
		return nil, false
	}
	return o.ReadmeUrl, true
}

// HasReadmeUrl returns a boolean if a field has been set.
func (o *SwiftPackageUploadRequest) HasReadmeUrl() bool {
	if o != nil && !IsNil(o.ReadmeUrl) {
		return true
	}

	return false
}

// SetReadmeUrl gets a reference to the given string and assigns it to the ReadmeUrl field.
func (o *SwiftPackageUploadRequest) SetReadmeUrl(v string) {
	o.ReadmeUrl = &v
}

// GetRepositoryUrl returns the RepositoryUrl field value if set, zero value otherwise.
func (o *SwiftPackageUploadRequest) GetRepositoryUrl() string {
	if o == nil || IsNil(o.RepositoryUrl) {
		var ret string
		return ret
	}
	return *o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwiftPackageUploadRequest) GetRepositoryUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryUrl) {
		return nil, false
	}
	return o.RepositoryUrl, true
}

// HasRepositoryUrl returns a boolean if a field has been set.
func (o *SwiftPackageUploadRequest) HasRepositoryUrl() bool {
	if o != nil && !IsNil(o.RepositoryUrl) {
		return true
	}

	return false
}

// SetRepositoryUrl gets a reference to the given string and assigns it to the RepositoryUrl field.
func (o *SwiftPackageUploadRequest) SetRepositoryUrl(v string) {
	o.RepositoryUrl = &v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *SwiftPackageUploadRequest) GetRepublish() bool {
	if o == nil || IsNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwiftPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || IsNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *SwiftPackageUploadRequest) HasRepublish() bool {
	if o != nil && !IsNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *SwiftPackageUploadRequest) SetRepublish(v bool) {
	o.Republish = &v
}

// GetScope returns the Scope field value
func (o *SwiftPackageUploadRequest) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *SwiftPackageUploadRequest) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *SwiftPackageUploadRequest) SetScope(v string) {
	o.Scope = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SwiftPackageUploadRequest) GetTags() string {
	if o == nil || IsNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SwiftPackageUploadRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *SwiftPackageUploadRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *SwiftPackageUploadRequest) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *SwiftPackageUploadRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *SwiftPackageUploadRequest) UnsetTags() {
	o.Tags.Unset()
}

// GetVersion returns the Version field value
func (o *SwiftPackageUploadRequest) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SwiftPackageUploadRequest) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *SwiftPackageUploadRequest) SetVersion(v string) {
	o.Version = v
}

func (o SwiftPackageUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwiftPackageUploadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorName) {
		toSerialize["author_name"] = o.AuthorName
	}
	if !IsNil(o.AuthorOrg) {
		toSerialize["author_org"] = o.AuthorOrg
	}
	if o.LicenseUrl.IsSet() {
		toSerialize["license_url"] = o.LicenseUrl.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["package_file"] = o.PackageFile
	if !IsNil(o.ReadmeUrl) {
		toSerialize["readme_url"] = o.ReadmeUrl
	}
	if !IsNil(o.RepositoryUrl) {
		toSerialize["repository_url"] = o.RepositoryUrl
	}
	if !IsNil(o.Republish) {
		toSerialize["republish"] = o.Republish
	}
	toSerialize["scope"] = o.Scope
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SwiftPackageUploadRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"package_file",
		"scope",
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

	varSwiftPackageUploadRequest := _SwiftPackageUploadRequest{}

	err = json.Unmarshal(data, &varSwiftPackageUploadRequest)

	if err != nil {
		return err
	}

	*o = SwiftPackageUploadRequest(varSwiftPackageUploadRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "author_name")
		delete(additionalProperties, "author_org")
		delete(additionalProperties, "license_url")
		delete(additionalProperties, "name")
		delete(additionalProperties, "package_file")
		delete(additionalProperties, "readme_url")
		delete(additionalProperties, "repository_url")
		delete(additionalProperties, "republish")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSwiftPackageUploadRequest struct {
	value *SwiftPackageUploadRequest
	isSet bool
}

func (v NullableSwiftPackageUploadRequest) Get() *SwiftPackageUploadRequest {
	return v.value
}

func (v *NullableSwiftPackageUploadRequest) Set(val *SwiftPackageUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSwiftPackageUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSwiftPackageUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwiftPackageUploadRequest(val *SwiftPackageUploadRequest) *NullableSwiftPackageUploadRequest {
	return &NullableSwiftPackageUploadRequest{value: val, isSet: true}
}

func (v NullableSwiftPackageUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwiftPackageUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
