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

// RepositoryTokenSyncRequest struct for RepositoryTokenSyncRequest
type RepositoryTokenSyncRequest struct {
	// The source repository slug (in the same owner namespace).
	Source string `json:"source"`
}

// NewRepositoryTokenSyncRequest instantiates a new RepositoryTokenSyncRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryTokenSyncRequest(source string) *RepositoryTokenSyncRequest {
	this := RepositoryTokenSyncRequest{}
	this.Source = source
	return &this
}

// NewRepositoryTokenSyncRequestWithDefaults instantiates a new RepositoryTokenSyncRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryTokenSyncRequestWithDefaults() *RepositoryTokenSyncRequest {
	this := RepositoryTokenSyncRequest{}
	return &this
}

// GetSource returns the Source field value
func (o *RepositoryTokenSyncRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *RepositoryTokenSyncRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *RepositoryTokenSyncRequest) SetSource(v string) {
	o.Source = v
}

func (o RepositoryTokenSyncRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryTokenSyncRequest struct {
	value *RepositoryTokenSyncRequest
	isSet bool
}

func (v NullableRepositoryTokenSyncRequest) Get() *RepositoryTokenSyncRequest {
	return v.value
}

func (v *NullableRepositoryTokenSyncRequest) Set(val *RepositoryTokenSyncRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryTokenSyncRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryTokenSyncRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryTokenSyncRequest(val *RepositoryTokenSyncRequest) *NullableRepositoryTokenSyncRequest {
	return &NullableRepositoryTokenSyncRequest{value: val, isSet: true}
}

func (v NullableRepositoryTokenSyncRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryTokenSyncRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
