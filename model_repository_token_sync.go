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
)

// checks if the RepositoryTokenSync type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryTokenSync{}

// RepositoryTokenSync struct for RepositoryTokenSync
type RepositoryTokenSync struct {
	// The entitlements that have been synchronised.
	Tokens               []RepositoryToken `json:"tokens,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RepositoryTokenSync RepositoryTokenSync

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
	if o == nil || IsNil(o.Tokens) {
		var ret []RepositoryToken
		return ret
	}
	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryTokenSync) GetTokensOk() ([]RepositoryToken, bool) {
	if o == nil || IsNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *RepositoryTokenSync) HasTokens() bool {
	if o != nil && !IsNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given []RepositoryToken and assigns it to the Tokens field.
func (o *RepositoryTokenSync) SetTokens(v []RepositoryToken) {
	o.Tokens = v
}

func (o RepositoryTokenSync) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryTokenSync) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RepositoryTokenSync) UnmarshalJSON(data []byte) (err error) {
	varRepositoryTokenSync := _RepositoryTokenSync{}

	err = json.Unmarshal(data, &varRepositoryTokenSync)

	if err != nil {
		return err
	}

	*o = RepositoryTokenSync(varRepositoryTokenSync)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tokens")
		o.AdditionalProperties = additionalProperties
	}

	return err
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
