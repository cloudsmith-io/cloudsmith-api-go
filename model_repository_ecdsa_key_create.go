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

// checks if the RepositoryEcdsaKeyCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryEcdsaKeyCreate{}

// RepositoryEcdsaKeyCreate struct for RepositoryEcdsaKeyCreate
type RepositoryEcdsaKeyCreate struct {
	// The ECDSA passphrase used for signing.
	EcdsaPassphrase *string `json:"ecdsa_passphrase,omitempty"`
	// The ECDSA private key.
	EcdsaPrivateKey      string `json:"ecdsa_private_key"`
	AdditionalProperties map[string]interface{}
}

type _RepositoryEcdsaKeyCreate RepositoryEcdsaKeyCreate

// NewRepositoryEcdsaKeyCreate instantiates a new RepositoryEcdsaKeyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryEcdsaKeyCreate(ecdsaPrivateKey string) *RepositoryEcdsaKeyCreate {
	this := RepositoryEcdsaKeyCreate{}
	this.EcdsaPrivateKey = ecdsaPrivateKey
	return &this
}

// NewRepositoryEcdsaKeyCreateWithDefaults instantiates a new RepositoryEcdsaKeyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryEcdsaKeyCreateWithDefaults() *RepositoryEcdsaKeyCreate {
	this := RepositoryEcdsaKeyCreate{}
	return &this
}

// GetEcdsaPassphrase returns the EcdsaPassphrase field value if set, zero value otherwise.
func (o *RepositoryEcdsaKeyCreate) GetEcdsaPassphrase() string {
	if o == nil || IsNil(o.EcdsaPassphrase) {
		var ret string
		return ret
	}
	return *o.EcdsaPassphrase
}

// GetEcdsaPassphraseOk returns a tuple with the EcdsaPassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryEcdsaKeyCreate) GetEcdsaPassphraseOk() (*string, bool) {
	if o == nil || IsNil(o.EcdsaPassphrase) {
		return nil, false
	}
	return o.EcdsaPassphrase, true
}

// HasEcdsaPassphrase returns a boolean if a field has been set.
func (o *RepositoryEcdsaKeyCreate) HasEcdsaPassphrase() bool {
	if o != nil && !IsNil(o.EcdsaPassphrase) {
		return true
	}

	return false
}

// SetEcdsaPassphrase gets a reference to the given string and assigns it to the EcdsaPassphrase field.
func (o *RepositoryEcdsaKeyCreate) SetEcdsaPassphrase(v string) {
	o.EcdsaPassphrase = &v
}

// GetEcdsaPrivateKey returns the EcdsaPrivateKey field value
func (o *RepositoryEcdsaKeyCreate) GetEcdsaPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EcdsaPrivateKey
}

// GetEcdsaPrivateKeyOk returns a tuple with the EcdsaPrivateKey field value
// and a boolean to check if the value has been set.
func (o *RepositoryEcdsaKeyCreate) GetEcdsaPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EcdsaPrivateKey, true
}

// SetEcdsaPrivateKey sets field value
func (o *RepositoryEcdsaKeyCreate) SetEcdsaPrivateKey(v string) {
	o.EcdsaPrivateKey = v
}

func (o RepositoryEcdsaKeyCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryEcdsaKeyCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EcdsaPassphrase) {
		toSerialize["ecdsa_passphrase"] = o.EcdsaPassphrase
	}
	toSerialize["ecdsa_private_key"] = o.EcdsaPrivateKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RepositoryEcdsaKeyCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ecdsa_private_key",
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

	varRepositoryEcdsaKeyCreate := _RepositoryEcdsaKeyCreate{}

	err = json.Unmarshal(data, &varRepositoryEcdsaKeyCreate)

	if err != nil {
		return err
	}

	*o = RepositoryEcdsaKeyCreate(varRepositoryEcdsaKeyCreate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ecdsa_passphrase")
		delete(additionalProperties, "ecdsa_private_key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRepositoryEcdsaKeyCreate struct {
	value *RepositoryEcdsaKeyCreate
	isSet bool
}

func (v NullableRepositoryEcdsaKeyCreate) Get() *RepositoryEcdsaKeyCreate {
	return v.value
}

func (v *NullableRepositoryEcdsaKeyCreate) Set(val *RepositoryEcdsaKeyCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryEcdsaKeyCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryEcdsaKeyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryEcdsaKeyCreate(val *RepositoryEcdsaKeyCreate) *NullableRepositoryEcdsaKeyCreate {
	return &NullableRepositoryEcdsaKeyCreate{value: val, isSet: true}
}

func (v NullableRepositoryEcdsaKeyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryEcdsaKeyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
