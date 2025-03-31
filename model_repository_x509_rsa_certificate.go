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

// checks if the RepositoryX509RsaCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryX509RsaCertificate{}

// RepositoryX509RsaCertificate struct for RepositoryX509RsaCertificate
type RepositoryX509RsaCertificate struct {
	// If selected this is the active key for this repository.
	Active *bool `json:"active,omitempty"`
	// The issued certificate.
	Certificate NullableString `json:"certificate,omitempty"`
	// Base64 encoded CA certificate chain.
	CertificateChain                 NullableString `json:"certificate_chain,omitempty"`
	CertificateChainFingerprint      *string        `json:"certificate_chain_fingerprint,omitempty"`
	CertificateChainFingerprintShort *string        `json:"certificate_chain_fingerprint_short,omitempty"`
	// The SHA-256 long identifier used
	CertificateFingerprint      NullableString `json:"certificate_fingerprint,omitempty"`
	CertificateFingerprintShort *string        `json:"certificate_fingerprint_short,omitempty"`
	CreatedAt                   *time.Time     `json:"created_at,omitempty"`
	// If selected this is the default key for this repository.
	Default              *bool   `json:"default,omitempty"`
	IssuingStatus        *string `json:"issuing_status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RepositoryX509RsaCertificate RepositoryX509RsaCertificate

// NewRepositoryX509RsaCertificate instantiates a new RepositoryX509RsaCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryX509RsaCertificate() *RepositoryX509RsaCertificate {
	this := RepositoryX509RsaCertificate{}
	var issuingStatus string = "Certificate is pending to be issued"
	this.IssuingStatus = &issuingStatus
	return &this
}

// NewRepositoryX509RsaCertificateWithDefaults instantiates a new RepositoryX509RsaCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryX509RsaCertificateWithDefaults() *RepositoryX509RsaCertificate {
	this := RepositoryX509RsaCertificate{}
	var issuingStatus string = "Certificate is pending to be issued"
	this.IssuingStatus = &issuingStatus
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *RepositoryX509RsaCertificate) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryX509RsaCertificate) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *RepositoryX509RsaCertificate) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *RepositoryX509RsaCertificate) SetActive(v bool) {
	o.Active = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryX509RsaCertificate) GetCertificate() string {
	if o == nil || IsNil(o.Certificate.Get()) {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryX509RsaCertificate) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *RepositoryX509RsaCertificate) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *RepositoryX509RsaCertificate) SetCertificate(v string) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *RepositoryX509RsaCertificate) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *RepositoryX509RsaCertificate) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetCertificateChain returns the CertificateChain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryX509RsaCertificate) GetCertificateChain() string {
	if o == nil || IsNil(o.CertificateChain.Get()) {
		var ret string
		return ret
	}
	return *o.CertificateChain.Get()
}

// GetCertificateChainOk returns a tuple with the CertificateChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryX509RsaCertificate) GetCertificateChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateChain.Get(), o.CertificateChain.IsSet()
}

// HasCertificateChain returns a boolean if a field has been set.
func (o *RepositoryX509RsaCertificate) HasCertificateChain() bool {
	if o != nil && o.CertificateChain.IsSet() {
		return true
	}

	return false
}

// SetCertificateChain gets a reference to the given NullableString and assigns it to the CertificateChain field.
func (o *RepositoryX509RsaCertificate) SetCertificateChain(v string) {
	o.CertificateChain.Set(&v)
}

// SetCertificateChainNil sets the value for CertificateChain to be an explicit nil
func (o *RepositoryX509RsaCertificate) SetCertificateChainNil() {
	o.CertificateChain.Set(nil)
}

// UnsetCertificateChain ensures that no value is present for CertificateChain, not even an explicit nil
func (o *RepositoryX509RsaCertificate) UnsetCertificateChain() {
	o.CertificateChain.Unset()
}

// GetCertificateChainFingerprint returns the CertificateChainFingerprint field value if set, zero value otherwise.
func (o *RepositoryX509RsaCertificate) GetCertificateChainFingerprint() string {
	if o == nil || IsNil(o.CertificateChainFingerprint) {
		var ret string
		return ret
	}
	return *o.CertificateChainFingerprint
}

// GetCertificateChainFingerprintOk returns a tuple with the CertificateChainFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryX509RsaCertificate) GetCertificateChainFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateChainFingerprint) {
		return nil, false
	}
	return o.CertificateChainFingerprint, true
}

// HasCertificateChainFingerprint returns a boolean if a field has been set.
func (o *RepositoryX509RsaCertificate) HasCertificateChainFingerprint() bool {
	if o != nil && !IsNil(o.CertificateChainFingerprint) {
		return true
	}

	return false
}

// SetCertificateChainFingerprint gets a reference to the given string and assigns it to the CertificateChainFingerprint field.
func (o *RepositoryX509RsaCertificate) SetCertificateChainFingerprint(v string) {
	o.CertificateChainFingerprint = &v
}

// GetCertificateChainFingerprintShort returns the CertificateChainFingerprintShort field value if set, zero value otherwise.
func (o *RepositoryX509RsaCertificate) GetCertificateChainFingerprintShort() string {
	if o == nil || IsNil(o.CertificateChainFingerprintShort) {
		var ret string
		return ret
	}
	return *o.CertificateChainFingerprintShort
}

// GetCertificateChainFingerprintShortOk returns a tuple with the CertificateChainFingerprintShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryX509RsaCertificate) GetCertificateChainFingerprintShortOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateChainFingerprintShort) {
		return nil, false
	}
	return o.CertificateChainFingerprintShort, true
}

// HasCertificateChainFingerprintShort returns a boolean if a field has been set.
func (o *RepositoryX509RsaCertificate) HasCertificateChainFingerprintShort() bool {
	if o != nil && !IsNil(o.CertificateChainFingerprintShort) {
		return true
	}

	return false
}

// SetCertificateChainFingerprintShort gets a reference to the given string and assigns it to the CertificateChainFingerprintShort field.
func (o *RepositoryX509RsaCertificate) SetCertificateChainFingerprintShort(v string) {
	o.CertificateChainFingerprintShort = &v
}

// GetCertificateFingerprint returns the CertificateFingerprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryX509RsaCertificate) GetCertificateFingerprint() string {
	if o == nil || IsNil(o.CertificateFingerprint.Get()) {
		var ret string
		return ret
	}
	return *o.CertificateFingerprint.Get()
}

// GetCertificateFingerprintOk returns a tuple with the CertificateFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryX509RsaCertificate) GetCertificateFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateFingerprint.Get(), o.CertificateFingerprint.IsSet()
}

// HasCertificateFingerprint returns a boolean if a field has been set.
func (o *RepositoryX509RsaCertificate) HasCertificateFingerprint() bool {
	if o != nil && o.CertificateFingerprint.IsSet() {
		return true
	}

	return false
}

// SetCertificateFingerprint gets a reference to the given NullableString and assigns it to the CertificateFingerprint field.
func (o *RepositoryX509RsaCertificate) SetCertificateFingerprint(v string) {
	o.CertificateFingerprint.Set(&v)
}

// SetCertificateFingerprintNil sets the value for CertificateFingerprint to be an explicit nil
func (o *RepositoryX509RsaCertificate) SetCertificateFingerprintNil() {
	o.CertificateFingerprint.Set(nil)
}

// UnsetCertificateFingerprint ensures that no value is present for CertificateFingerprint, not even an explicit nil
func (o *RepositoryX509RsaCertificate) UnsetCertificateFingerprint() {
	o.CertificateFingerprint.Unset()
}

// GetCertificateFingerprintShort returns the CertificateFingerprintShort field value if set, zero value otherwise.
func (o *RepositoryX509RsaCertificate) GetCertificateFingerprintShort() string {
	if o == nil || IsNil(o.CertificateFingerprintShort) {
		var ret string
		return ret
	}
	return *o.CertificateFingerprintShort
}

// GetCertificateFingerprintShortOk returns a tuple with the CertificateFingerprintShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryX509RsaCertificate) GetCertificateFingerprintShortOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateFingerprintShort) {
		return nil, false
	}
	return o.CertificateFingerprintShort, true
}

// HasCertificateFingerprintShort returns a boolean if a field has been set.
func (o *RepositoryX509RsaCertificate) HasCertificateFingerprintShort() bool {
	if o != nil && !IsNil(o.CertificateFingerprintShort) {
		return true
	}

	return false
}

// SetCertificateFingerprintShort gets a reference to the given string and assigns it to the CertificateFingerprintShort field.
func (o *RepositoryX509RsaCertificate) SetCertificateFingerprintShort(v string) {
	o.CertificateFingerprintShort = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RepositoryX509RsaCertificate) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryX509RsaCertificate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RepositoryX509RsaCertificate) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RepositoryX509RsaCertificate) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *RepositoryX509RsaCertificate) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryX509RsaCertificate) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *RepositoryX509RsaCertificate) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *RepositoryX509RsaCertificate) SetDefault(v bool) {
	o.Default = &v
}

// GetIssuingStatus returns the IssuingStatus field value if set, zero value otherwise.
func (o *RepositoryX509RsaCertificate) GetIssuingStatus() string {
	if o == nil || IsNil(o.IssuingStatus) {
		var ret string
		return ret
	}
	return *o.IssuingStatus
}

// GetIssuingStatusOk returns a tuple with the IssuingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryX509RsaCertificate) GetIssuingStatusOk() (*string, bool) {
	if o == nil || IsNil(o.IssuingStatus) {
		return nil, false
	}
	return o.IssuingStatus, true
}

// HasIssuingStatus returns a boolean if a field has been set.
func (o *RepositoryX509RsaCertificate) HasIssuingStatus() bool {
	if o != nil && !IsNil(o.IssuingStatus) {
		return true
	}

	return false
}

// SetIssuingStatus gets a reference to the given string and assigns it to the IssuingStatus field.
func (o *RepositoryX509RsaCertificate) SetIssuingStatus(v string) {
	o.IssuingStatus = &v
}

func (o RepositoryX509RsaCertificate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryX509RsaCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.CertificateChain.IsSet() {
		toSerialize["certificate_chain"] = o.CertificateChain.Get()
	}
	if !IsNil(o.CertificateChainFingerprint) {
		toSerialize["certificate_chain_fingerprint"] = o.CertificateChainFingerprint
	}
	if !IsNil(o.CertificateChainFingerprintShort) {
		toSerialize["certificate_chain_fingerprint_short"] = o.CertificateChainFingerprintShort
	}
	if o.CertificateFingerprint.IsSet() {
		toSerialize["certificate_fingerprint"] = o.CertificateFingerprint.Get()
	}
	if !IsNil(o.CertificateFingerprintShort) {
		toSerialize["certificate_fingerprint_short"] = o.CertificateFingerprintShort
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.IssuingStatus) {
		toSerialize["issuing_status"] = o.IssuingStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RepositoryX509RsaCertificate) UnmarshalJSON(data []byte) (err error) {
	varRepositoryX509RsaCertificate := _RepositoryX509RsaCertificate{}

	err = json.Unmarshal(data, &varRepositoryX509RsaCertificate)

	if err != nil {
		return err
	}

	*o = RepositoryX509RsaCertificate(varRepositoryX509RsaCertificate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "certificate_chain")
		delete(additionalProperties, "certificate_chain_fingerprint")
		delete(additionalProperties, "certificate_chain_fingerprint_short")
		delete(additionalProperties, "certificate_fingerprint")
		delete(additionalProperties, "certificate_fingerprint_short")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "default")
		delete(additionalProperties, "issuing_status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRepositoryX509RsaCertificate struct {
	value *RepositoryX509RsaCertificate
	isSet bool
}

func (v NullableRepositoryX509RsaCertificate) Get() *RepositoryX509RsaCertificate {
	return v.value
}

func (v *NullableRepositoryX509RsaCertificate) Set(val *RepositoryX509RsaCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryX509RsaCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryX509RsaCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryX509RsaCertificate(val *RepositoryX509RsaCertificate) *NullableRepositoryX509RsaCertificate {
	return &NullableRepositoryX509RsaCertificate{value: val, isSet: true}
}

func (v NullableRepositoryX509RsaCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryX509RsaCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
