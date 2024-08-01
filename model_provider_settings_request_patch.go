/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.477.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the ProviderSettingsRequestPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderSettingsRequestPatch{}

// ProviderSettingsRequestPatch struct for ProviderSettingsRequestPatch
type ProviderSettingsRequestPatch struct {
	// The set of claims that any received tokens from the provider must contain to authenticate as the configured service account.
	Claims map[string]interface{} `json:"claims,omitempty"`
	// Whether the provider settings should be used for incoming OIDC requests.
	Enabled *bool `json:"enabled,omitempty"`
	// The name of the provider settings are being configured for
	Name *string `json:"name,omitempty"`
	// The URL from the provider that serves as the base for the OpenID configuration. For example, if the OpenID configuration is available at https://token.actions.githubusercontent.com/.well-known/openid-configuration, the provider URL would be https://token.actions.githubusercontent.com/
	ProviderUrl *string `json:"provider_url,omitempty"`
	// The service accounts associated with these provider settings
	ServiceAccounts []string `json:"service_accounts,omitempty"`
}

// NewProviderSettingsRequestPatch instantiates a new ProviderSettingsRequestPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderSettingsRequestPatch() *ProviderSettingsRequestPatch {
	this := ProviderSettingsRequestPatch{}
	return &this
}

// NewProviderSettingsRequestPatchWithDefaults instantiates a new ProviderSettingsRequestPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderSettingsRequestPatchWithDefaults() *ProviderSettingsRequestPatch {
	this := ProviderSettingsRequestPatch{}
	return &this
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *ProviderSettingsRequestPatch) GetClaims() map[string]interface{} {
	if o == nil || IsNil(o.Claims) {
		var ret map[string]interface{}
		return ret
	}
	return o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSettingsRequestPatch) GetClaimsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Claims) {
		return map[string]interface{}{}, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *ProviderSettingsRequestPatch) HasClaims() bool {
	if o != nil && !IsNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given map[string]interface{} and assigns it to the Claims field.
func (o *ProviderSettingsRequestPatch) SetClaims(v map[string]interface{}) {
	o.Claims = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ProviderSettingsRequestPatch) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSettingsRequestPatch) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ProviderSettingsRequestPatch) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ProviderSettingsRequestPatch) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProviderSettingsRequestPatch) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSettingsRequestPatch) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProviderSettingsRequestPatch) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProviderSettingsRequestPatch) SetName(v string) {
	o.Name = &v
}

// GetProviderUrl returns the ProviderUrl field value if set, zero value otherwise.
func (o *ProviderSettingsRequestPatch) GetProviderUrl() string {
	if o == nil || IsNil(o.ProviderUrl) {
		var ret string
		return ret
	}
	return *o.ProviderUrl
}

// GetProviderUrlOk returns a tuple with the ProviderUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSettingsRequestPatch) GetProviderUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderUrl) {
		return nil, false
	}
	return o.ProviderUrl, true
}

// HasProviderUrl returns a boolean if a field has been set.
func (o *ProviderSettingsRequestPatch) HasProviderUrl() bool {
	if o != nil && !IsNil(o.ProviderUrl) {
		return true
	}

	return false
}

// SetProviderUrl gets a reference to the given string and assigns it to the ProviderUrl field.
func (o *ProviderSettingsRequestPatch) SetProviderUrl(v string) {
	o.ProviderUrl = &v
}

// GetServiceAccounts returns the ServiceAccounts field value if set, zero value otherwise.
func (o *ProviderSettingsRequestPatch) GetServiceAccounts() []string {
	if o == nil || IsNil(o.ServiceAccounts) {
		var ret []string
		return ret
	}
	return o.ServiceAccounts
}

// GetServiceAccountsOk returns a tuple with the ServiceAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSettingsRequestPatch) GetServiceAccountsOk() ([]string, bool) {
	if o == nil || IsNil(o.ServiceAccounts) {
		return nil, false
	}
	return o.ServiceAccounts, true
}

// HasServiceAccounts returns a boolean if a field has been set.
func (o *ProviderSettingsRequestPatch) HasServiceAccounts() bool {
	if o != nil && !IsNil(o.ServiceAccounts) {
		return true
	}

	return false
}

// SetServiceAccounts gets a reference to the given []string and assigns it to the ServiceAccounts field.
func (o *ProviderSettingsRequestPatch) SetServiceAccounts(v []string) {
	o.ServiceAccounts = v
}

func (o ProviderSettingsRequestPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderSettingsRequestPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Claims) {
		toSerialize["claims"] = o.Claims
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProviderUrl) {
		toSerialize["provider_url"] = o.ProviderUrl
	}
	if !IsNil(o.ServiceAccounts) {
		toSerialize["service_accounts"] = o.ServiceAccounts
	}
	return toSerialize, nil
}

type NullableProviderSettingsRequestPatch struct {
	value *ProviderSettingsRequestPatch
	isSet bool
}

func (v NullableProviderSettingsRequestPatch) Get() *ProviderSettingsRequestPatch {
	return v.value
}

func (v *NullableProviderSettingsRequestPatch) Set(val *ProviderSettingsRequestPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderSettingsRequestPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderSettingsRequestPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderSettingsRequestPatch(val *ProviderSettingsRequestPatch) *NullableProviderSettingsRequestPatch {
	return &NullableProviderSettingsRequestPatch{value: val, isSet: true}
}

func (v NullableProviderSettingsRequestPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderSettingsRequestPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
