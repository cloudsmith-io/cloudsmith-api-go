/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.209.19
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// RepositoryTokenSync struct for RepositoryTokenSync
type RepositoryTokenSync struct {
	// The entitlements that have been synchronised.
	Tokens []RepositoryToken `json:"tokens,omitempty"`
}

// NewRepositoryTokenSync instantiates a new RepositoryTokenSync object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryTokenSync() *RepositoryTokenSync {
	this := RepositoryTokenSync{}
	return &this
}

// NewRepositoryTokenSyncWithDefaults instantiates a new RepositoryTokenSync object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryTokenSyncWithDefaults() *RepositoryTokenSync {
	this := RepositoryTokenSync{}
	return &this
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *RepositoryTokenSync) GetTokens() []RepositoryToken {
	if o == nil || isNil(o.Tokens) {
		var ret []RepositoryToken
		return ret
	}
	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryTokenSync) GetTokensOk() ([]RepositoryToken, bool) {
	if o == nil || isNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *RepositoryTokenSync) HasTokens() bool {
	if o != nil && !isNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given []RepositoryToken and assigns it to the Tokens field.
func (o *RepositoryTokenSync) SetTokens(v []RepositoryToken) {
	o.Tokens = v
}

func (o RepositoryTokenSync) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryTokenSync struct {
	value *RepositoryTokenSync
	isSet bool
}

func (v NullableRepositoryTokenSync) Get() *RepositoryTokenSync {
	return v.value
}

func (v *NullableRepositoryTokenSync) Set(val *RepositoryTokenSync) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryTokenSync) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryTokenSync) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryTokenSync(val *RepositoryTokenSync) *NullableRepositoryTokenSync {
	return &NullableRepositoryTokenSync{value: val, isSet: true}
}

func (v NullableRepositoryTokenSync) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryTokenSync) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
