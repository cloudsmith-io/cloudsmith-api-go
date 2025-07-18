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

// checks if the FormatSupport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormatSupport{}

// FormatSupport A set of what the package format supports
type FormatSupport struct {
	// If true the package format supports dependencies
	Dependencies bool `json:"dependencies"`
	// If true the package format supports distributions
	Distributions bool `json:"distributions"`
	// If true the package format supports file lists
	FileLists bool `json:"file_lists"`
	// If true the package format supports metadata
	Metadata  bool                  `json:"metadata"`
	Upstreams FormatSupportUpstream `json:"upstreams"`
	// If true the package format supports versioning
	Versioning           bool `json:"versioning"`
	AdditionalProperties map[string]interface{}
}

type _FormatSupport FormatSupport

// NewFormatSupport instantiates a new FormatSupport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormatSupport(dependencies bool, distributions bool, fileLists bool, metadata bool, upstreams FormatSupportUpstream, versioning bool) *FormatSupport {
	this := FormatSupport{}
	this.Dependencies = dependencies
	this.Distributions = distributions
	this.FileLists = fileLists
	this.Metadata = metadata
	this.Upstreams = upstreams
	this.Versioning = versioning
	return &this
}

// NewFormatSupportWithDefaults instantiates a new FormatSupport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormatSupportWithDefaults() *FormatSupport {
	this := FormatSupport{}
	return &this
}

// GetDependencies returns the Dependencies field value
func (o *FormatSupport) GetDependencies() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value
// and a boolean to check if the value has been set.
func (o *FormatSupport) GetDependenciesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dependencies, true
}

// SetDependencies sets field value
func (o *FormatSupport) SetDependencies(v bool) {
	o.Dependencies = v
}

// GetDistributions returns the Distributions field value
func (o *FormatSupport) GetDistributions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Distributions
}

// GetDistributionsOk returns a tuple with the Distributions field value
// and a boolean to check if the value has been set.
func (o *FormatSupport) GetDistributionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Distributions, true
}

// SetDistributions sets field value
func (o *FormatSupport) SetDistributions(v bool) {
	o.Distributions = v
}

// GetFileLists returns the FileLists field value
func (o *FormatSupport) GetFileLists() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FileLists
}

// GetFileListsOk returns a tuple with the FileLists field value
// and a boolean to check if the value has been set.
func (o *FormatSupport) GetFileListsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileLists, true
}

// SetFileLists sets field value
func (o *FormatSupport) SetFileLists(v bool) {
	o.FileLists = v
}

// GetMetadata returns the Metadata field value
func (o *FormatSupport) GetMetadata() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *FormatSupport) GetMetadataOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *FormatSupport) SetMetadata(v bool) {
	o.Metadata = v
}

// GetUpstreams returns the Upstreams field value
func (o *FormatSupport) GetUpstreams() FormatSupportUpstream {
	if o == nil {
		var ret FormatSupportUpstream
		return ret
	}

	return o.Upstreams
}

// GetUpstreamsOk returns a tuple with the Upstreams field value
// and a boolean to check if the value has been set.
func (o *FormatSupport) GetUpstreamsOk() (*FormatSupportUpstream, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Upstreams, true
}

// SetUpstreams sets field value
func (o *FormatSupport) SetUpstreams(v FormatSupportUpstream) {
	o.Upstreams = v
}

// GetVersioning returns the Versioning field value
func (o *FormatSupport) GetVersioning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Versioning
}

// GetVersioningOk returns a tuple with the Versioning field value
// and a boolean to check if the value has been set.
func (o *FormatSupport) GetVersioningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Versioning, true
}

// SetVersioning sets field value
func (o *FormatSupport) SetVersioning(v bool) {
	o.Versioning = v
}

func (o FormatSupport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormatSupport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dependencies"] = o.Dependencies
	toSerialize["distributions"] = o.Distributions
	toSerialize["file_lists"] = o.FileLists
	toSerialize["metadata"] = o.Metadata
	toSerialize["upstreams"] = o.Upstreams
	toSerialize["versioning"] = o.Versioning

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FormatSupport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dependencies",
		"distributions",
		"file_lists",
		"metadata",
		"upstreams",
		"versioning",
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

	varFormatSupport := _FormatSupport{}

	err = json.Unmarshal(data, &varFormatSupport)

	if err != nil {
		return err
	}

	*o = FormatSupport(varFormatSupport)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dependencies")
		delete(additionalProperties, "distributions")
		delete(additionalProperties, "file_lists")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "upstreams")
		delete(additionalProperties, "versioning")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormatSupport struct {
	value *FormatSupport
	isSet bool
}

func (v NullableFormatSupport) Get() *FormatSupport {
	return v.value
}

func (v *NullableFormatSupport) Set(val *FormatSupport) {
	v.value = val
	v.isSet = true
}

func (v NullableFormatSupport) IsSet() bool {
	return v.isSet
}

func (v *NullableFormatSupport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormatSupport(val *FormatSupport) *NullableFormatSupport {
	return &NullableFormatSupport{value: val, isSet: true}
}

func (v NullableFormatSupport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormatSupport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
