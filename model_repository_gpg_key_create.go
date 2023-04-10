/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.237.2
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// RepositoryGpgKeyCreate struct for RepositoryGpgKeyCreate
type RepositoryGpgKeyCreate struct {
	// The GPG passphrase used for signing.
	GpgPassphrase *string `json:"gpg_passphrase,omitempty"`
	// The GPG private key.
	GpgPrivateKey string `json:"gpg_private_key"`
}

// NewRepositoryGpgKeyCreate instantiates a new RepositoryGpgKeyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGpgKeyCreate(gpgPrivateKey string) *RepositoryGpgKeyCreate {
	this := RepositoryGpgKeyCreate{}
	this.GpgPrivateKey = gpgPrivateKey
	return &this
}

// NewRepositoryGpgKeyCreateWithDefaults instantiates a new RepositoryGpgKeyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGpgKeyCreateWithDefaults() *RepositoryGpgKeyCreate {
	this := RepositoryGpgKeyCreate{}
	return &this
}

// GetGpgPassphrase returns the GpgPassphrase field value if set, zero value otherwise.
func (o *RepositoryGpgKeyCreate) GetGpgPassphrase() string {
	if o == nil || isNil(o.GpgPassphrase) {
		var ret string
		return ret
	}
	return *o.GpgPassphrase
}

// GetGpgPassphraseOk returns a tuple with the GpgPassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGpgKeyCreate) GetGpgPassphraseOk() (*string, bool) {
	if o == nil || isNil(o.GpgPassphrase) {
		return nil, false
	}
	return o.GpgPassphrase, true
}

// HasGpgPassphrase returns a boolean if a field has been set.
func (o *RepositoryGpgKeyCreate) HasGpgPassphrase() bool {
	if o != nil && !isNil(o.GpgPassphrase) {
		return true
	}

	return false
}

// SetGpgPassphrase gets a reference to the given string and assigns it to the GpgPassphrase field.
func (o *RepositoryGpgKeyCreate) SetGpgPassphrase(v string) {
	o.GpgPassphrase = &v
}

// GetGpgPrivateKey returns the GpgPrivateKey field value
func (o *RepositoryGpgKeyCreate) GetGpgPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GpgPrivateKey
}

// GetGpgPrivateKeyOk returns a tuple with the GpgPrivateKey field value
// and a boolean to check if the value has been set.
func (o *RepositoryGpgKeyCreate) GetGpgPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GpgPrivateKey, true
}

// SetGpgPrivateKey sets field value
func (o *RepositoryGpgKeyCreate) SetGpgPrivateKey(v string) {
	o.GpgPrivateKey = v
}

func (o RepositoryGpgKeyCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GpgPassphrase) {
		toSerialize["gpg_passphrase"] = o.GpgPassphrase
	}
	if true {
		toSerialize["gpg_private_key"] = o.GpgPrivateKey
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryGpgKeyCreate struct {
	value *RepositoryGpgKeyCreate
	isSet bool
}

func (v NullableRepositoryGpgKeyCreate) Get() *RepositoryGpgKeyCreate {
	return v.value
}

func (v *NullableRepositoryGpgKeyCreate) Set(val *RepositoryGpgKeyCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGpgKeyCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGpgKeyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGpgKeyCreate(val *RepositoryGpgKeyCreate) *NullableRepositoryGpgKeyCreate {
	return &NullableRepositoryGpgKeyCreate{value: val, isSet: true}
}

func (v NullableRepositoryGpgKeyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGpgKeyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
