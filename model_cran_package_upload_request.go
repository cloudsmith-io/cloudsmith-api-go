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

// checks if the CranPackageUploadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CranPackageUploadRequest{}

// CranPackageUploadRequest struct for CranPackageUploadRequest
type CranPackageUploadRequest struct {
	// Binary package uploads for macOS should specify the architecture they were built for.
	Architecture *string `json:"architecture,omitempty"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// Binary package uploads should specify the version of R they were built for.
	RVersion NullableString `json:"r_version,omitempty"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags                 NullableString `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CranPackageUploadRequest CranPackageUploadRequest

// NewCranPackageUploadRequest instantiates a new CranPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCranPackageUploadRequest(packageFile string) *CranPackageUploadRequest {
	this := CranPackageUploadRequest{}
	this.PackageFile = packageFile
	return &this
}

// NewCranPackageUploadRequestWithDefaults instantiates a new CranPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCranPackageUploadRequestWithDefaults() *CranPackageUploadRequest {
	this := CranPackageUploadRequest{}
	return &this
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *CranPackageUploadRequest) GetArchitecture() string {
	if o == nil || IsNil(o.Architecture) {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CranPackageUploadRequest) GetArchitectureOk() (*string, bool) {
	if o == nil || IsNil(o.Architecture) {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *CranPackageUploadRequest) HasArchitecture() bool {
	if o != nil && !IsNil(o.Architecture) {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *CranPackageUploadRequest) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetPackageFile returns the PackageFile field value
func (o *CranPackageUploadRequest) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *CranPackageUploadRequest) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *CranPackageUploadRequest) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRVersion returns the RVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CranPackageUploadRequest) GetRVersion() string {
	if o == nil || IsNil(o.RVersion.Get()) {
		var ret string
		return ret
	}
	return *o.RVersion.Get()
}

// GetRVersionOk returns a tuple with the RVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CranPackageUploadRequest) GetRVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RVersion.Get(), o.RVersion.IsSet()
}

// HasRVersion returns a boolean if a field has been set.
func (o *CranPackageUploadRequest) HasRVersion() bool {
	if o != nil && o.RVersion.IsSet() {
		return true
	}

	return false
}

// SetRVersion gets a reference to the given NullableString and assigns it to the RVersion field.
func (o *CranPackageUploadRequest) SetRVersion(v string) {
	o.RVersion.Set(&v)
}

// SetRVersionNil sets the value for RVersion to be an explicit nil
func (o *CranPackageUploadRequest) SetRVersionNil() {
	o.RVersion.Set(nil)
}

// UnsetRVersion ensures that no value is present for RVersion, not even an explicit nil
func (o *CranPackageUploadRequest) UnsetRVersion() {
	o.RVersion.Unset()
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *CranPackageUploadRequest) GetRepublish() bool {
	if o == nil || IsNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CranPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || IsNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *CranPackageUploadRequest) HasRepublish() bool {
	if o != nil && !IsNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *CranPackageUploadRequest) SetRepublish(v bool) {
	o.Republish = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CranPackageUploadRequest) GetTags() string {
	if o == nil || IsNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CranPackageUploadRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *CranPackageUploadRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *CranPackageUploadRequest) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *CranPackageUploadRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *CranPackageUploadRequest) UnsetTags() {
	o.Tags.Unset()
}

func (o CranPackageUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CranPackageUploadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Architecture) {
		toSerialize["architecture"] = o.Architecture
	}
	toSerialize["package_file"] = o.PackageFile
	if o.RVersion.IsSet() {
		toSerialize["r_version"] = o.RVersion.Get()
	}
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

func (o *CranPackageUploadRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCranPackageUploadRequest := _CranPackageUploadRequest{}

	err = json.Unmarshal(data, &varCranPackageUploadRequest)

	if err != nil {
		return err
	}

	*o = CranPackageUploadRequest(varCranPackageUploadRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "architecture")
		delete(additionalProperties, "package_file")
		delete(additionalProperties, "r_version")
		delete(additionalProperties, "republish")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCranPackageUploadRequest struct {
	value *CranPackageUploadRequest
	isSet bool
}

func (v NullableCranPackageUploadRequest) Get() *CranPackageUploadRequest {
	return v.value
}

func (v *NullableCranPackageUploadRequest) Set(val *CranPackageUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCranPackageUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCranPackageUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCranPackageUploadRequest(val *CranPackageUploadRequest) *NullableCranPackageUploadRequest {
	return &NullableCranPackageUploadRequest{value: val, isSet: true}
}

func (v NullableCranPackageUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCranPackageUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
