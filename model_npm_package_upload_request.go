/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.654.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"fmt"
)

// checks if the NpmPackageUploadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NpmPackageUploadRequest{}

// NpmPackageUploadRequest struct for NpmPackageUploadRequest
type NpmPackageUploadRequest struct {
	// The default npm dist-tag for this package/version - This will replace any other package/version if they are using the same tag.
	NpmDistTag *string `json:"npm_dist_tag,omitempty"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags                 NullableString `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NpmPackageUploadRequest NpmPackageUploadRequest

// NewNpmPackageUploadRequest instantiates a new NpmPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNpmPackageUploadRequest(packageFile string) *NpmPackageUploadRequest {
	this := NpmPackageUploadRequest{}
	var npmDistTag string = "latest"
	this.NpmDistTag = &npmDistTag
	this.PackageFile = packageFile
	return &this
}

// NewNpmPackageUploadRequestWithDefaults instantiates a new NpmPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNpmPackageUploadRequestWithDefaults() *NpmPackageUploadRequest {
	this := NpmPackageUploadRequest{}
	var npmDistTag string = "latest"
	this.NpmDistTag = &npmDistTag
	return &this
}

// GetNpmDistTag returns the NpmDistTag field value if set, zero value otherwise.
func (o *NpmPackageUploadRequest) GetNpmDistTag() string {
	if o == nil || IsNil(o.NpmDistTag) {
		var ret string
		return ret
	}
	return *o.NpmDistTag
}

// GetNpmDistTagOk returns a tuple with the NpmDistTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NpmPackageUploadRequest) GetNpmDistTagOk() (*string, bool) {
	if o == nil || IsNil(o.NpmDistTag) {
		return nil, false
	}
	return o.NpmDistTag, true
}

// HasNpmDistTag returns a boolean if a field has been set.
func (o *NpmPackageUploadRequest) HasNpmDistTag() bool {
	if o != nil && !IsNil(o.NpmDistTag) {
		return true
	}

	return false
}

// SetNpmDistTag gets a reference to the given string and assigns it to the NpmDistTag field.
func (o *NpmPackageUploadRequest) SetNpmDistTag(v string) {
	o.NpmDistTag = &v
}

// GetPackageFile returns the PackageFile field value
func (o *NpmPackageUploadRequest) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *NpmPackageUploadRequest) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *NpmPackageUploadRequest) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *NpmPackageUploadRequest) GetRepublish() bool {
	if o == nil || IsNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NpmPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || IsNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *NpmPackageUploadRequest) HasRepublish() bool {
	if o != nil && !IsNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *NpmPackageUploadRequest) SetRepublish(v bool) {
	o.Republish = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NpmPackageUploadRequest) GetTags() string {
	if o == nil || IsNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NpmPackageUploadRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *NpmPackageUploadRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *NpmPackageUploadRequest) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *NpmPackageUploadRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *NpmPackageUploadRequest) UnsetTags() {
	o.Tags.Unset()
}

func (o NpmPackageUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NpmPackageUploadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NpmDistTag) {
		toSerialize["npm_dist_tag"] = o.NpmDistTag
	}
	toSerialize["package_file"] = o.PackageFile
	if !IsNil(o.Republish) {
		toSerialize["republish"] = o.Republish
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NpmPackageUploadRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varNpmPackageUploadRequest := _NpmPackageUploadRequest{}

	err = json.Unmarshal(data, &varNpmPackageUploadRequest)

	if err != nil {
		return err
	}

	*o = NpmPackageUploadRequest(varNpmPackageUploadRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "npm_dist_tag")
		delete(additionalProperties, "package_file")
		delete(additionalProperties, "republish")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNpmPackageUploadRequest struct {
	value *NpmPackageUploadRequest
	isSet bool
}

func (v NullableNpmPackageUploadRequest) Get() *NpmPackageUploadRequest {
	return v.value
}

func (v *NullableNpmPackageUploadRequest) Set(val *NpmPackageUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNpmPackageUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNpmPackageUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNpmPackageUploadRequest(val *NpmPackageUploadRequest) *NullableNpmPackageUploadRequest {
	return &NullableNpmPackageUploadRequest{value: val, isSet: true}
}

func (v NullableNpmPackageUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNpmPackageUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
