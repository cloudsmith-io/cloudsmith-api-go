/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.236.5
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// RepositoryRsaKeyCreate struct for RepositoryRsaKeyCreate
type RepositoryRsaKeyCreate struct {
	// The RSA passphrase used for signing.
	RsaPassphrase *string `json:"rsa_passphrase,omitempty"`
	// The RSA private key.
	RsaPrivateKey string `json:"rsa_private_key"`
}

// NewRepositoryRsaKeyCreate instantiates a new RepositoryRsaKeyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryRsaKeyCreate(rsaPrivateKey string) *RepositoryRsaKeyCreate {
	this := RepositoryRsaKeyCreate{}
	this.RsaPrivateKey = rsaPrivateKey
	return &this
}

// NewRepositoryRsaKeyCreateWithDefaults instantiates a new RepositoryRsaKeyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryRsaKeyCreateWithDefaults() *RepositoryRsaKeyCreate {
	this := RepositoryRsaKeyCreate{}
	return &this
}

// GetRsaPassphrase returns the RsaPassphrase field value if set, zero value otherwise.
func (o *RepositoryRsaKeyCreate) GetRsaPassphrase() string {
	if o == nil || isNil(o.RsaPassphrase) {
		var ret string
		return ret
	}
	return *o.RsaPassphrase
}

// GetRsaPassphraseOk returns a tuple with the RsaPassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRsaKeyCreate) GetRsaPassphraseOk() (*string, bool) {
	if o == nil || isNil(o.RsaPassphrase) {
		return nil, false
	}
	return o.RsaPassphrase, true
}

// HasRsaPassphrase returns a boolean if a field has been set.
func (o *RepositoryRsaKeyCreate) HasRsaPassphrase() bool {
	if o != nil && !isNil(o.RsaPassphrase) {
		return true
	}

	return false
}

// SetRsaPassphrase gets a reference to the given string and assigns it to the RsaPassphrase field.
func (o *RepositoryRsaKeyCreate) SetRsaPassphrase(v string) {
	o.RsaPassphrase = &v
}

// GetRsaPrivateKey returns the RsaPrivateKey field value
func (o *RepositoryRsaKeyCreate) GetRsaPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RsaPrivateKey
}

// GetRsaPrivateKeyOk returns a tuple with the RsaPrivateKey field value
// and a boolean to check if the value has been set.
func (o *RepositoryRsaKeyCreate) GetRsaPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RsaPrivateKey, true
}

// SetRsaPrivateKey sets field value
func (o *RepositoryRsaKeyCreate) SetRsaPrivateKey(v string) {
	o.RsaPrivateKey = v
}

func (o RepositoryRsaKeyCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RsaPassphrase) {
		toSerialize["rsa_passphrase"] = o.RsaPassphrase
	}
	if true {
		toSerialize["rsa_private_key"] = o.RsaPrivateKey
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryRsaKeyCreate struct {
	value *RepositoryRsaKeyCreate
	isSet bool
}

func (v NullableRepositoryRsaKeyCreate) Get() *RepositoryRsaKeyCreate {
	return v.value
}

func (v *NullableRepositoryRsaKeyCreate) Set(val *RepositoryRsaKeyCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryRsaKeyCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryRsaKeyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryRsaKeyCreate(val *RepositoryRsaKeyCreate) *NullableRepositoryRsaKeyCreate {
	return &NullableRepositoryRsaKeyCreate{value: val, isSet: true}
}

func (v NullableRepositoryRsaKeyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryRsaKeyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
