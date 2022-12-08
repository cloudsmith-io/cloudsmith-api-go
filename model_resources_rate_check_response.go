/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.181.6
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// ResourcesRateCheckResponse struct for ResourcesRateCheckResponse
type ResourcesRateCheckResponse struct {
	// Rate limit values per resource
	Resources *map[string]RateCheck `json:"resources,omitempty"`
}

// NewResourcesRateCheckResponse instantiates a new ResourcesRateCheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcesRateCheckResponse() *ResourcesRateCheckResponse {
	this := ResourcesRateCheckResponse{}
	return &this
}

// NewResourcesRateCheckResponseWithDefaults instantiates a new ResourcesRateCheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcesRateCheckResponseWithDefaults() *ResourcesRateCheckResponse {
	this := ResourcesRateCheckResponse{}
	return &this
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ResourcesRateCheckResponse) GetResources() map[string]RateCheck {
	if o == nil || isNil(o.Resources) {
		var ret map[string]RateCheck
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcesRateCheckResponse) GetResourcesOk() (*map[string]RateCheck, bool) {
	if o == nil || isNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ResourcesRateCheckResponse) HasResources() bool {
	if o != nil && !isNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given map[string]RateCheck and assigns it to the Resources field.
func (o *ResourcesRateCheckResponse) SetResources(v map[string]RateCheck) {
	o.Resources = &v
}

func (o ResourcesRateCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableResourcesRateCheckResponse struct {
	value *ResourcesRateCheckResponse
	isSet bool
}

func (v NullableResourcesRateCheckResponse) Get() *ResourcesRateCheckResponse {
	return v.value
}

func (v *NullableResourcesRateCheckResponse) Set(val *ResourcesRateCheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcesRateCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcesRateCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcesRateCheckResponse(val *ResourcesRateCheckResponse) *NullableResourcesRateCheckResponse {
	return &NullableResourcesRateCheckResponse{value: val, isSet: true}
}

func (v NullableResourcesRateCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcesRateCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
