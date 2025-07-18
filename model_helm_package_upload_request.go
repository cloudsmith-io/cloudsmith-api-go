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

// checks if the HelmPackageUploadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmPackageUploadRequest{}

// HelmPackageUploadRequest struct for HelmPackageUploadRequest
type HelmPackageUploadRequest struct {
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// The provenance file containing the signature for the chart. If one is not provided, it will be generated automatically.
	ProvenanceFile NullableString `json:"provenance_file,omitempty"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags                 NullableString `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HelmPackageUploadRequest HelmPackageUploadRequest

// NewHelmPackageUploadRequest instantiates a new HelmPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmPackageUploadRequest(packageFile string) *HelmPackageUploadRequest {
	this := HelmPackageUploadRequest{}
	this.PackageFile = packageFile
	return &this
}

// NewHelmPackageUploadRequestWithDefaults instantiates a new HelmPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmPackageUploadRequestWithDefaults() *HelmPackageUploadRequest {
	this := HelmPackageUploadRequest{}
	return &this
}

// GetPackageFile returns the PackageFile field value
func (o *HelmPackageUploadRequest) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *HelmPackageUploadRequest) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *HelmPackageUploadRequest) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetProvenanceFile returns the ProvenanceFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HelmPackageUploadRequest) GetProvenanceFile() string {
	if o == nil || IsNil(o.ProvenanceFile.Get()) {
		var ret string
		return ret
	}
	return *o.ProvenanceFile.Get()
}

// GetProvenanceFileOk returns a tuple with the ProvenanceFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HelmPackageUploadRequest) GetProvenanceFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProvenanceFile.Get(), o.ProvenanceFile.IsSet()
}

// HasProvenanceFile returns a boolean if a field has been set.
func (o *HelmPackageUploadRequest) HasProvenanceFile() bool {
	if o != nil && o.ProvenanceFile.IsSet() {
		return true
	}

	return false
}

// SetProvenanceFile gets a reference to the given NullableString and assigns it to the ProvenanceFile field.
func (o *HelmPackageUploadRequest) SetProvenanceFile(v string) {
	o.ProvenanceFile.Set(&v)
}

// SetProvenanceFileNil sets the value for ProvenanceFile to be an explicit nil
func (o *HelmPackageUploadRequest) SetProvenanceFileNil() {
	o.ProvenanceFile.Set(nil)
}

// UnsetProvenanceFile ensures that no value is present for ProvenanceFile, not even an explicit nil
func (o *HelmPackageUploadRequest) UnsetProvenanceFile() {
	o.ProvenanceFile.Unset()
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *HelmPackageUploadRequest) GetRepublish() bool {
	if o == nil || IsNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || IsNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *HelmPackageUploadRequest) HasRepublish() bool {
	if o != nil && !IsNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *HelmPackageUploadRequest) SetRepublish(v bool) {
	o.Republish = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HelmPackageUploadRequest) GetTags() string {
	if o == nil || IsNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HelmPackageUploadRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *HelmPackageUploadRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *HelmPackageUploadRequest) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *HelmPackageUploadRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *HelmPackageUploadRequest) UnsetTags() {
	o.Tags.Unset()
}

func (o HelmPackageUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmPackageUploadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["package_file"] = o.PackageFile
	if o.ProvenanceFile.IsSet() {
		toSerialize["provenance_file"] = o.ProvenanceFile.Get()
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

func (o *HelmPackageUploadRequest) UnmarshalJSON(data []byte) (err error) {
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

	varHelmPackageUploadRequest := _HelmPackageUploadRequest{}

	err = json.Unmarshal(data, &varHelmPackageUploadRequest)

	if err != nil {
		return err
	}

	*o = HelmPackageUploadRequest(varHelmPackageUploadRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "package_file")
		delete(additionalProperties, "provenance_file")
		delete(additionalProperties, "republish")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHelmPackageUploadRequest struct {
	value *HelmPackageUploadRequest
	isSet bool
}

func (v NullableHelmPackageUploadRequest) Get() *HelmPackageUploadRequest {
	return v.value
}

func (v *NullableHelmPackageUploadRequest) Set(val *HelmPackageUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmPackageUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmPackageUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmPackageUploadRequest(val *HelmPackageUploadRequest) *NullableHelmPackageUploadRequest {
	return &NullableHelmPackageUploadRequest{value: val, isSet: true}
}

func (v NullableHelmPackageUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmPackageUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
