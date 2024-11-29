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
	"time"
)

// checks if the RepositoryGpgKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryGpgKey{}

// RepositoryGpgKey struct for RepositoryGpgKey
type RepositoryGpgKey struct {
	// If selected this is the active key for this repository.
	Active    *bool      `json:"active,omitempty"`
	Comment   string     `json:"comment"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// If selected this is the default key for this repository.
	Default *bool `json:"default,omitempty"`
	// The long identifier used by GPG for this key.
	Fingerprint      *string `json:"fingerprint,omitempty"`
	FingerprintShort *string `json:"fingerprint_short,omitempty"`
	// The public key given to repository users.
	PublicKey            *string `json:"public_key,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RepositoryGpgKey RepositoryGpgKey

// NewRepositoryGpgKey instantiates a new RepositoryGpgKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGpgKey(comment string) *RepositoryGpgKey {
	this := RepositoryGpgKey{}
	this.Comment = comment
	return &this
}

// NewRepositoryGpgKeyWithDefaults instantiates a new RepositoryGpgKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGpgKeyWithDefaults() *RepositoryGpgKey {
	this := RepositoryGpgKey{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *RepositoryGpgKey) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGpgKey) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *RepositoryGpgKey) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *RepositoryGpgKey) SetActive(v bool) {
	o.Active = &v
}

// GetComment returns the Comment field value
func (o *RepositoryGpgKey) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *RepositoryGpgKey) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *RepositoryGpgKey) SetComment(v string) {
	o.Comment = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RepositoryGpgKey) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGpgKey) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RepositoryGpgKey) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RepositoryGpgKey) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *RepositoryGpgKey) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGpgKey) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *RepositoryGpgKey) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *RepositoryGpgKey) SetDefault(v bool) {
	o.Default = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *RepositoryGpgKey) GetFingerprint() string {
	if o == nil || IsNil(o.Fingerprint) {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGpgKey) GetFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *RepositoryGpgKey) HasFingerprint() bool {
	if o != nil && !IsNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *RepositoryGpgKey) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetFingerprintShort returns the FingerprintShort field value if set, zero value otherwise.
func (o *RepositoryGpgKey) GetFingerprintShort() string {
	if o == nil || IsNil(o.FingerprintShort) {
		var ret string
		return ret
	}
	return *o.FingerprintShort
}

// GetFingerprintShortOk returns a tuple with the FingerprintShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGpgKey) GetFingerprintShortOk() (*string, bool) {
	if o == nil || IsNil(o.FingerprintShort) {
		return nil, false
	}
	return o.FingerprintShort, true
}

// HasFingerprintShort returns a boolean if a field has been set.
func (o *RepositoryGpgKey) HasFingerprintShort() bool {
	if o != nil && !IsNil(o.FingerprintShort) {
		return true
	}

	return false
}

// SetFingerprintShort gets a reference to the given string and assigns it to the FingerprintShort field.
func (o *RepositoryGpgKey) SetFingerprintShort(v string) {
	o.FingerprintShort = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *RepositoryGpgKey) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGpgKey) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *RepositoryGpgKey) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *RepositoryGpgKey) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o RepositoryGpgKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryGpgKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["comment"] = o.Comment
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RepositoryGpgKey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"comment",
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

	varRepositoryGpgKey := _RepositoryGpgKey{}

	err = json.Unmarshal(data, &varRepositoryGpgKey)

	if err != nil {
		return err
	}

	*o = RepositoryGpgKey(varRepositoryGpgKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "default")
		delete(additionalProperties, "fingerprint")
		delete(additionalProperties, "fingerprint_short")
		delete(additionalProperties, "public_key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRepositoryGpgKey struct {
	value *RepositoryGpgKey
	isSet bool
}

func (v NullableRepositoryGpgKey) Get() *RepositoryGpgKey {
	return v.value
}

func (v *NullableRepositoryGpgKey) Set(val *RepositoryGpgKey) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGpgKey) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGpgKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGpgKey(val *RepositoryGpgKey) *NullableRepositoryGpgKey {
	return &NullableRepositoryGpgKey{value: val, isSet: true}
}

func (v NullableRepositoryGpgKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGpgKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
