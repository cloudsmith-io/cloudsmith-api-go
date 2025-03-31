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
	"time"
)

// checks if the RepositoryEcdsaKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryEcdsaKey{}

// RepositoryEcdsaKey struct for RepositoryEcdsaKey
type RepositoryEcdsaKey struct {
	// If selected this is the active key for this repository.
	Active    *bool      `json:"active,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// If selected this is the default key for this repository.
	Default *bool `json:"default,omitempty"`
	// The long identifier used by ECDSA for this key.
	Fingerprint      *string `json:"fingerprint,omitempty"`
	FingerprintShort *string `json:"fingerprint_short,omitempty"`
	// The public key given to repository users.
	PublicKey *string `json:"public_key,omitempty"`
	// The SSH fingerprint used by ECDSA for this key.
	SshFingerprint       NullableString `json:"ssh_fingerprint,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RepositoryEcdsaKey RepositoryEcdsaKey

// NewRepositoryEcdsaKey instantiates a new RepositoryEcdsaKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryEcdsaKey() *RepositoryEcdsaKey {
	this := RepositoryEcdsaKey{}
	return &this
}

// NewRepositoryEcdsaKeyWithDefaults instantiates a new RepositoryEcdsaKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryEcdsaKeyWithDefaults() *RepositoryEcdsaKey {
	this := RepositoryEcdsaKey{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *RepositoryEcdsaKey) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryEcdsaKey) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *RepositoryEcdsaKey) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *RepositoryEcdsaKey) SetActive(v bool) {
	o.Active = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RepositoryEcdsaKey) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryEcdsaKey) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RepositoryEcdsaKey) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RepositoryEcdsaKey) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *RepositoryEcdsaKey) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryEcdsaKey) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *RepositoryEcdsaKey) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *RepositoryEcdsaKey) SetDefault(v bool) {
	o.Default = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *RepositoryEcdsaKey) GetFingerprint() string {
	if o == nil || IsNil(o.Fingerprint) {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryEcdsaKey) GetFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *RepositoryEcdsaKey) HasFingerprint() bool {
	if o != nil && !IsNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *RepositoryEcdsaKey) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetFingerprintShort returns the FingerprintShort field value if set, zero value otherwise.
func (o *RepositoryEcdsaKey) GetFingerprintShort() string {
	if o == nil || IsNil(o.FingerprintShort) {
		var ret string
		return ret
	}
	return *o.FingerprintShort
}

// GetFingerprintShortOk returns a tuple with the FingerprintShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryEcdsaKey) GetFingerprintShortOk() (*string, bool) {
	if o == nil || IsNil(o.FingerprintShort) {
		return nil, false
	}
	return o.FingerprintShort, true
}

// HasFingerprintShort returns a boolean if a field has been set.
func (o *RepositoryEcdsaKey) HasFingerprintShort() bool {
	if o != nil && !IsNil(o.FingerprintShort) {
		return true
	}

	return false
}

// SetFingerprintShort gets a reference to the given string and assigns it to the FingerprintShort field.
func (o *RepositoryEcdsaKey) SetFingerprintShort(v string) {
	o.FingerprintShort = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *RepositoryEcdsaKey) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryEcdsaKey) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *RepositoryEcdsaKey) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *RepositoryEcdsaKey) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetSshFingerprint returns the SshFingerprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryEcdsaKey) GetSshFingerprint() string {
	if o == nil || IsNil(o.SshFingerprint.Get()) {
		var ret string
		return ret
	}
	return *o.SshFingerprint.Get()
}

// GetSshFingerprintOk returns a tuple with the SshFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryEcdsaKey) GetSshFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SshFingerprint.Get(), o.SshFingerprint.IsSet()
}

// HasSshFingerprint returns a boolean if a field has been set.
func (o *RepositoryEcdsaKey) HasSshFingerprint() bool {
	if o != nil && o.SshFingerprint.IsSet() {
		return true
	}

	return false
}

// SetSshFingerprint gets a reference to the given NullableString and assigns it to the SshFingerprint field.
func (o *RepositoryEcdsaKey) SetSshFingerprint(v string) {
	o.SshFingerprint.Set(&v)
}

// SetSshFingerprintNil sets the value for SshFingerprint to be an explicit nil
func (o *RepositoryEcdsaKey) SetSshFingerprintNil() {
	o.SshFingerprint.Set(nil)
}

// UnsetSshFingerprint ensures that no value is present for SshFingerprint, not even an explicit nil
func (o *RepositoryEcdsaKey) UnsetSshFingerprint() {
	o.SshFingerprint.Unset()
}

func (o RepositoryEcdsaKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryEcdsaKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Fingerprint) {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if !IsNil(o.FingerprintShort) {
		toSerialize["fingerprint_short"] = o.FingerprintShort
	}
	if !IsNil(o.PublicKey) {
		toSerialize["public_key"] = o.PublicKey
	}
	if o.SshFingerprint.IsSet() {
		toSerialize["ssh_fingerprint"] = o.SshFingerprint.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RepositoryEcdsaKey) UnmarshalJSON(data []byte) (err error) {
	varRepositoryEcdsaKey := _RepositoryEcdsaKey{}

	err = json.Unmarshal(data, &varRepositoryEcdsaKey)

	if err != nil {
		return err
	}

	*o = RepositoryEcdsaKey(varRepositoryEcdsaKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "default")
		delete(additionalProperties, "fingerprint")
		delete(additionalProperties, "fingerprint_short")
		delete(additionalProperties, "public_key")
		delete(additionalProperties, "ssh_fingerprint")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRepositoryEcdsaKey struct {
	value *RepositoryEcdsaKey
	isSet bool
}

func (v NullableRepositoryEcdsaKey) Get() *RepositoryEcdsaKey {
	return v.value
}

func (v *NullableRepositoryEcdsaKey) Set(val *RepositoryEcdsaKey) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryEcdsaKey) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryEcdsaKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryEcdsaKey(val *RepositoryEcdsaKey) *NullableRepositoryEcdsaKey {
	return &NullableRepositoryEcdsaKey{value: val, isSet: true}
}

func (v NullableRepositoryEcdsaKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryEcdsaKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
