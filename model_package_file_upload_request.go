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

// checks if the PackageFileUploadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageFileUploadRequest{}

// PackageFileUploadRequest struct for PackageFileUploadRequest
type PackageFileUploadRequest struct {
	// Filename for the package file upload.
	Filename string `json:"filename"`
	// MD5 checksum for a POST-based package file upload.
	Md5Checksum *string `json:"md5_checksum,omitempty"`
	// The method to use for package file upload.
	Method *string `json:"method,omitempty"`
	// SHA256 checksum for a PUT-based package file upload.
	Sha256Checksum       *string `json:"sha256_checksum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PackageFileUploadRequest PackageFileUploadRequest

// NewPackageFileUploadRequest instantiates a new PackageFileUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageFileUploadRequest(filename string) *PackageFileUploadRequest {
	this := PackageFileUploadRequest{}
	this.Filename = filename
	var method string = "post"
	this.Method = &method
	return &this
}

// NewPackageFileUploadRequestWithDefaults instantiates a new PackageFileUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageFileUploadRequestWithDefaults() *PackageFileUploadRequest {
	this := PackageFileUploadRequest{}
	var method string = "post"
	this.Method = &method
	return &this
}

// GetFilename returns the Filename field value
func (o *PackageFileUploadRequest) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *PackageFileUploadRequest) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *PackageFileUploadRequest) SetFilename(v string) {
	o.Filename = v
}

// GetMd5Checksum returns the Md5Checksum field value if set, zero value otherwise.
func (o *PackageFileUploadRequest) GetMd5Checksum() string {
	if o == nil || IsNil(o.Md5Checksum) {
		var ret string
		return ret
	}
	return *o.Md5Checksum
}

// GetMd5ChecksumOk returns a tuple with the Md5Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageFileUploadRequest) GetMd5ChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.Md5Checksum) {
		return nil, false
	}
	return o.Md5Checksum, true
}

// HasMd5Checksum returns a boolean if a field has been set.
func (o *PackageFileUploadRequest) HasMd5Checksum() bool {
	if o != nil && !IsNil(o.Md5Checksum) {
		return true
	}

	return false
}

// SetMd5Checksum gets a reference to the given string and assigns it to the Md5Checksum field.
func (o *PackageFileUploadRequest) SetMd5Checksum(v string) {
	o.Md5Checksum = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PackageFileUploadRequest) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageFileUploadRequest) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PackageFileUploadRequest) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *PackageFileUploadRequest) SetMethod(v string) {
	o.Method = &v
}

// GetSha256Checksum returns the Sha256Checksum field value if set, zero value otherwise.
func (o *PackageFileUploadRequest) GetSha256Checksum() string {
	if o == nil || IsNil(o.Sha256Checksum) {
		var ret string
		return ret
	}
	return *o.Sha256Checksum
}

// GetSha256ChecksumOk returns a tuple with the Sha256Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageFileUploadRequest) GetSha256ChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.Sha256Checksum) {
		return nil, false
	}
	return o.Sha256Checksum, true
}

// HasSha256Checksum returns a boolean if a field has been set.
func (o *PackageFileUploadRequest) HasSha256Checksum() bool {
	if o != nil && !IsNil(o.Sha256Checksum) {
		return true
	}

	return false
}

// SetSha256Checksum gets a reference to the given string and assigns it to the Sha256Checksum field.
func (o *PackageFileUploadRequest) SetSha256Checksum(v string) {
	o.Sha256Checksum = &v
}

func (o PackageFileUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageFileUploadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filename"] = o.Filename
	if !IsNil(o.Md5Checksum) {
		toSerialize["md5_checksum"] = o.Md5Checksum
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Sha256Checksum) {
		toSerialize["sha256_checksum"] = o.Sha256Checksum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PackageFileUploadRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filename",
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

	varPackageFileUploadRequest := _PackageFileUploadRequest{}

	err = json.Unmarshal(data, &varPackageFileUploadRequest)

	if err != nil {
		return err
	}

	*o = PackageFileUploadRequest(varPackageFileUploadRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filename")
		delete(additionalProperties, "md5_checksum")
		delete(additionalProperties, "method")
		delete(additionalProperties, "sha256_checksum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePackageFileUploadRequest struct {
	value *PackageFileUploadRequest
	isSet bool
}

func (v NullablePackageFileUploadRequest) Get() *PackageFileUploadRequest {
	return v.value
}

func (v *NullablePackageFileUploadRequest) Set(val *PackageFileUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageFileUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageFileUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageFileUploadRequest(val *PackageFileUploadRequest) *NullablePackageFileUploadRequest {
	return &NullablePackageFileUploadRequest{value: val, isSet: true}
}

func (v NullablePackageFileUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageFileUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
