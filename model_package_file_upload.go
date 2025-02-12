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
)

// checks if the PackageFileUpload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageFileUpload{}

// PackageFileUpload struct for PackageFileUpload
type PackageFileUpload struct {
	// The identifier for the file to use when creating packages
	Identifier *string `json:"identifier,omitempty"`
	// The dictionary of fields that must be sent with POST uploads
	UploadFields map[string]interface{} `json:"upload_fields,omitempty"`
	// The dictionary of headers that must be sent with uploads
	UploadHeaders map[string]interface{} `json:"upload_headers,omitempty"`
	// The querystring to use for the next-step POST or PUT upload
	UploadQuerystring NullableString `json:"upload_querystring,omitempty"`
	// The URL to use for the next-step POST or PUT upload
	UploadUrl            *string `json:"upload_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PackageFileUpload PackageFileUpload

// NewPackageFileUpload instantiates a new PackageFileUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageFileUpload() *PackageFileUpload {
	this := PackageFileUpload{}
	return &this
}

// NewPackageFileUploadWithDefaults instantiates a new PackageFileUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageFileUploadWithDefaults() *PackageFileUpload {
	this := PackageFileUpload{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *PackageFileUpload) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageFileUpload) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *PackageFileUpload) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *PackageFileUpload) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetUploadFields returns the UploadFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageFileUpload) GetUploadFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.UploadFields
}

// GetUploadFieldsOk returns a tuple with the UploadFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageFileUpload) GetUploadFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UploadFields) {
		return map[string]interface{}{}, false
	}
	return o.UploadFields, true
}

// HasUploadFields returns a boolean if a field has been set.
func (o *PackageFileUpload) HasUploadFields() bool {
	if o != nil && !IsNil(o.UploadFields) {
		return true
	}

	return false
}

// SetUploadFields gets a reference to the given map[string]interface{} and assigns it to the UploadFields field.
func (o *PackageFileUpload) SetUploadFields(v map[string]interface{}) {
	o.UploadFields = v
}

// GetUploadHeaders returns the UploadHeaders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageFileUpload) GetUploadHeaders() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.UploadHeaders
}

// GetUploadHeadersOk returns a tuple with the UploadHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageFileUpload) GetUploadHeadersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UploadHeaders) {
		return map[string]interface{}{}, false
	}
	return o.UploadHeaders, true
}

// HasUploadHeaders returns a boolean if a field has been set.
func (o *PackageFileUpload) HasUploadHeaders() bool {
	if o != nil && !IsNil(o.UploadHeaders) {
		return true
	}

	return false
}

// SetUploadHeaders gets a reference to the given map[string]interface{} and assigns it to the UploadHeaders field.
func (o *PackageFileUpload) SetUploadHeaders(v map[string]interface{}) {
	o.UploadHeaders = v
}

// GetUploadQuerystring returns the UploadQuerystring field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageFileUpload) GetUploadQuerystring() string {
	if o == nil || IsNil(o.UploadQuerystring.Get()) {
		var ret string
		return ret
	}
	return *o.UploadQuerystring.Get()
}

// GetUploadQuerystringOk returns a tuple with the UploadQuerystring field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageFileUpload) GetUploadQuerystringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UploadQuerystring.Get(), o.UploadQuerystring.IsSet()
}

// HasUploadQuerystring returns a boolean if a field has been set.
func (o *PackageFileUpload) HasUploadQuerystring() bool {
	if o != nil && o.UploadQuerystring.IsSet() {
		return true
	}

	return false
}

// SetUploadQuerystring gets a reference to the given NullableString and assigns it to the UploadQuerystring field.
func (o *PackageFileUpload) SetUploadQuerystring(v string) {
	o.UploadQuerystring.Set(&v)
}

// SetUploadQuerystringNil sets the value for UploadQuerystring to be an explicit nil
func (o *PackageFileUpload) SetUploadQuerystringNil() {
	o.UploadQuerystring.Set(nil)
}

// UnsetUploadQuerystring ensures that no value is present for UploadQuerystring, not even an explicit nil
func (o *PackageFileUpload) UnsetUploadQuerystring() {
	o.UploadQuerystring.Unset()
}

// GetUploadUrl returns the UploadUrl field value if set, zero value otherwise.
func (o *PackageFileUpload) GetUploadUrl() string {
	if o == nil || IsNil(o.UploadUrl) {
		var ret string
		return ret
	}
	return *o.UploadUrl
}

// GetUploadUrlOk returns a tuple with the UploadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageFileUpload) GetUploadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.UploadUrl) {
		return nil, false
	}
	return o.UploadUrl, true
}

// HasUploadUrl returns a boolean if a field has been set.
func (o *PackageFileUpload) HasUploadUrl() bool {
	if o != nil && !IsNil(o.UploadUrl) {
		return true
	}

	return false
}

// SetUploadUrl gets a reference to the given string and assigns it to the UploadUrl field.
func (o *PackageFileUpload) SetUploadUrl(v string) {
	o.UploadUrl = &v
}

func (o PackageFileUpload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageFileUpload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if o.UploadFields != nil {
		toSerialize["upload_fields"] = o.UploadFields
	}
	if o.UploadHeaders != nil {
		toSerialize["upload_headers"] = o.UploadHeaders
	}
	if o.UploadQuerystring.IsSet() {
		toSerialize["upload_querystring"] = o.UploadQuerystring.Get()
	}
	if !IsNil(o.UploadUrl) {
		toSerialize["upload_url"] = o.UploadUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PackageFileUpload) UnmarshalJSON(data []byte) (err error) {
	varPackageFileUpload := _PackageFileUpload{}

	err = json.Unmarshal(data, &varPackageFileUpload)

	if err != nil {
		return err
	}

	*o = PackageFileUpload(varPackageFileUpload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "upload_fields")
		delete(additionalProperties, "upload_headers")
		delete(additionalProperties, "upload_querystring")
		delete(additionalProperties, "upload_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePackageFileUpload struct {
	value *PackageFileUpload
	isSet bool
}

func (v NullablePackageFileUpload) Get() *PackageFileUpload {
	return v.value
}

func (v *NullablePackageFileUpload) Set(val *PackageFileUpload) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageFileUpload) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageFileUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageFileUpload(val *PackageFileUpload) *NullablePackageFileUpload {
	return &NullablePackageFileUpload{value: val, isSet: true}
}

func (v NullablePackageFileUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageFileUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
