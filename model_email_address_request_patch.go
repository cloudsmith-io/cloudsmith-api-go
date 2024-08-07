/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.478.2
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the EmailAddressRequestPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailAddressRequestPatch{}

// EmailAddressRequestPatch struct for EmailAddressRequestPatch
type EmailAddressRequestPatch struct {
	Email    *string `json:"email,omitempty"`
	Primary  *bool   `json:"primary,omitempty"`
	Verified *bool   `json:"verified,omitempty"`
}

// NewEmailAddressRequestPatch instantiates a new EmailAddressRequestPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailAddressRequestPatch() *EmailAddressRequestPatch {
	this := EmailAddressRequestPatch{}
	return &this
}

// NewEmailAddressRequestPatchWithDefaults instantiates a new EmailAddressRequestPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailAddressRequestPatchWithDefaults() *EmailAddressRequestPatch {
	this := EmailAddressRequestPatch{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EmailAddressRequestPatch) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAddressRequestPatch) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EmailAddressRequestPatch) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EmailAddressRequestPatch) SetEmail(v string) {
	o.Email = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *EmailAddressRequestPatch) GetPrimary() bool {
	if o == nil || IsNil(o.Primary) {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAddressRequestPatch) GetPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *EmailAddressRequestPatch) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *EmailAddressRequestPatch) SetPrimary(v bool) {
	o.Primary = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *EmailAddressRequestPatch) GetVerified() bool {
	if o == nil || IsNil(o.Verified) {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAddressRequestPatch) GetVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Verified) {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *EmailAddressRequestPatch) HasVerified() bool {
	if o != nil && !IsNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *EmailAddressRequestPatch) SetVerified(v bool) {
	o.Verified = &v
}

func (o EmailAddressRequestPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailAddressRequestPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if !IsNil(o.Verified) {
		toSerialize["verified"] = o.Verified
	}
	return toSerialize, nil
}

type NullableEmailAddressRequestPatch struct {
	value *EmailAddressRequestPatch
	isSet bool
}

func (v NullableEmailAddressRequestPatch) Get() *EmailAddressRequestPatch {
	return v.value
}

func (v *NullableEmailAddressRequestPatch) Set(val *EmailAddressRequestPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailAddressRequestPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailAddressRequestPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailAddressRequestPatch(val *EmailAddressRequestPatch) *NullableEmailAddressRequestPatch {
	return &NullableEmailAddressRequestPatch{value: val, isSet: true}
}

func (v NullableEmailAddressRequestPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailAddressRequestPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
