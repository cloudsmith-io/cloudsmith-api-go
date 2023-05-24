/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.273.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"time"
)

// RepositoryRsaKey struct for RepositoryRsaKey
type RepositoryRsaKey struct {
	// If selected this is the active key for this repository.
	Active    *bool      `json:"active,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// If selected this is the default key for this repository.
	Default *bool `json:"default,omitempty"`
	// The long identifier used by RSA for this key.
	Fingerprint      *string `json:"fingerprint,omitempty"`
	FingerprintShort *string `json:"fingerprint_short,omitempty"`
	// The public key given to repository users.
	PublicKey *string `json:"public_key,omitempty"`
}

// NewRepositoryRsaKey instantiates a new RepositoryRsaKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryRsaKey() *RepositoryRsaKey {
	this := RepositoryRsaKey{}
	return &this
}

// NewRepositoryRsaKeyWithDefaults instantiates a new RepositoryRsaKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryRsaKeyWithDefaults() *RepositoryRsaKey {
	this := RepositoryRsaKey{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *RepositoryRsaKey) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRsaKey) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *RepositoryRsaKey) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *RepositoryRsaKey) SetActive(v bool) {
	o.Active = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RepositoryRsaKey) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRsaKey) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RepositoryRsaKey) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RepositoryRsaKey) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *RepositoryRsaKey) GetDefault() bool {
	if o == nil || isNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRsaKey) GetDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *RepositoryRsaKey) HasDefault() bool {
	if o != nil && !isNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *RepositoryRsaKey) SetDefault(v bool) {
	o.Default = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *RepositoryRsaKey) GetFingerprint() string {
	if o == nil || isNil(o.Fingerprint) {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRsaKey) GetFingerprintOk() (*string, bool) {
	if o == nil || isNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *RepositoryRsaKey) HasFingerprint() bool {
	if o != nil && !isNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *RepositoryRsaKey) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetFingerprintShort returns the FingerprintShort field value if set, zero value otherwise.
func (o *RepositoryRsaKey) GetFingerprintShort() string {
	if o == nil || isNil(o.FingerprintShort) {
		var ret string
		return ret
	}
	return *o.FingerprintShort
}

// GetFingerprintShortOk returns a tuple with the FingerprintShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRsaKey) GetFingerprintShortOk() (*string, bool) {
	if o == nil || isNil(o.FingerprintShort) {
		return nil, false
	}
	return o.FingerprintShort, true
}

// HasFingerprintShort returns a boolean if a field has been set.
func (o *RepositoryRsaKey) HasFingerprintShort() bool {
	if o != nil && !isNil(o.FingerprintShort) {
		return true
	}

	return false
}

// SetFingerprintShort gets a reference to the given string and assigns it to the FingerprintShort field.
func (o *RepositoryRsaKey) SetFingerprintShort(v string) {
	o.FingerprintShort = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *RepositoryRsaKey) GetPublicKey() string {
	if o == nil || isNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryRsaKey) GetPublicKeyOk() (*string, bool) {
	if o == nil || isNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *RepositoryRsaKey) HasPublicKey() bool {
	if o != nil && !isNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *RepositoryRsaKey) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o RepositoryRsaKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !isNil(o.Fingerprint) {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if !isNil(o.FingerprintShort) {
		toSerialize["fingerprint_short"] = o.FingerprintShort
	}
	if !isNil(o.PublicKey) {
		toSerialize["public_key"] = o.PublicKey
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryRsaKey struct {
	value *RepositoryRsaKey
	isSet bool
}

func (v NullableRepositoryRsaKey) Get() *RepositoryRsaKey {
	return v.value
}

func (v *NullableRepositoryRsaKey) Set(val *RepositoryRsaKey) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryRsaKey) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryRsaKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryRsaKey(val *RepositoryRsaKey) *NullableRepositoryRsaKey {
	return &NullableRepositoryRsaKey{value: val, isSet: true}
}

func (v NullableRepositoryRsaKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryRsaKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
