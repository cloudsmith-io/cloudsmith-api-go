/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.617.2
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"fmt"
)

// checks if the RepositoryTokenSyncRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryTokenSyncRequest{}

// RepositoryTokenSyncRequest struct for RepositoryTokenSyncRequest
type RepositoryTokenSyncRequest struct {
	// The source repository slug (in the same owner namespace).
	Source               string `json:"source"`
	AdditionalProperties map[string]interface{}
}

type _RepositoryTokenSyncRequest RepositoryTokenSyncRequest

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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryTokenSyncRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RepositoryTokenSyncRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
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

	varRepositoryTokenSyncRequest := _RepositoryTokenSyncRequest{}

	err = json.Unmarshal(data, &varRepositoryTokenSyncRequest)

	if err != nil {
		return err
	}

	*o = RepositoryTokenSyncRequest(varRepositoryTokenSyncRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	}

	return err
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
