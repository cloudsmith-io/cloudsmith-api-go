/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.616.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"fmt"
)

// checks if the NpmUpstreamRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NpmUpstreamRequest{}

// NpmUpstreamRequest struct for NpmUpstreamRequest
type NpmUpstreamRequest struct {
	// The authentication mode to use when accessing this upstream.
	AuthMode *string `json:"auth_mode,omitempty"`
	// Secret to provide with requests to upstream.
	AuthSecret NullableString `json:"auth_secret,omitempty"`
	// Username to provide with requests to upstream.
	AuthUsername NullableString `json:"auth_username,omitempty"`
	// The key for extra header #1 to send to upstream.
	ExtraHeader1 NullableString `json:"extra_header_1,omitempty" validate:"regexp=^[-\\\\w]+$"`
	// The key for extra header #2 to send to upstream.
	ExtraHeader2 NullableString `json:"extra_header_2,omitempty" validate:"regexp=^[-\\\\w]+$"`
	// The value for extra header #1 to send to upstream. This is stored as plaintext, and is NOT encrypted.
	ExtraValue1 NullableString `json:"extra_value_1,omitempty" validate:"regexp=^[^\\\\n\\\\r]+$"`
	// The value for extra header #2 to send to upstream. This is stored as plaintext, and is NOT encrypted.
	ExtraValue2 NullableString `json:"extra_value_2,omitempty" validate:"regexp=^[^\\\\n\\\\r]+$"`
	// Whether or not this upstream is active and ready for requests.
	IsActive *bool `json:"is_active,omitempty"`
	// The mode that this upstream should operate in. Upstream sources can be used to proxy resolved packages, as well as operate in a proxy/cache or cache only mode.
	Mode *string `json:"mode,omitempty"`
	// A descriptive name for this upstream source. A shortened version of this name will be used for tagging cached packages retrieved from this upstream.
	Name string `json:"name" validate:"regexp=^\\\\w[\\\\w \\\\-'\\\\.\\/()]+$"`
	// Upstream sources are selected for resolving requests by sequential order (1..n), followed by creation date.
	Priority *int64 `json:"priority,omitempty"`
	// The URL for this upstream source. This must be a fully qualified URL including any path elements required to reach the root of the repository.
	UpstreamUrl string `json:"upstream_url"`
	// If enabled, SSL certificates are verified when requests are made to this upstream. It's recommended to leave this enabled for all public sources to help mitigate Man-In-The-Middle (MITM) attacks. Please note this only applies to HTTPS upstreams.
	VerifySsl            *bool `json:"verify_ssl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NpmUpstreamRequest NpmUpstreamRequest

// NewNpmUpstreamRequest instantiates a new NpmUpstreamRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNpmUpstreamRequest(name string, upstreamUrl string) *NpmUpstreamRequest {
	this := NpmUpstreamRequest{}
	var authMode string = "None"
	this.AuthMode = &authMode
	var mode string = "Proxy Only"
	this.Mode = &mode
	this.Name = name
	this.UpstreamUrl = upstreamUrl
	return &this
}

// NewNpmUpstreamRequestWithDefaults instantiates a new NpmUpstreamRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNpmUpstreamRequestWithDefaults() *NpmUpstreamRequest {
	this := NpmUpstreamRequest{}
	var authMode string = "None"
	this.AuthMode = &authMode
	var mode string = "Proxy Only"
	this.Mode = &mode
	return &this
}

// GetAuthMode returns the AuthMode field value if set, zero value otherwise.
func (o *NpmUpstreamRequest) GetAuthMode() string {
	if o == nil || IsNil(o.AuthMode) {
		var ret string
		return ret
	}
	return *o.AuthMode
}

// GetAuthModeOk returns a tuple with the AuthMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NpmUpstreamRequest) GetAuthModeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthMode) {
		return nil, false
	}
	return o.AuthMode, true
}

// HasAuthMode returns a boolean if a field has been set.
func (o *NpmUpstreamRequest) HasAuthMode() bool {
	if o != nil && !IsNil(o.AuthMode) {
		return true
	}

	return false
}

// SetAuthMode gets a reference to the given string and assigns it to the AuthMode field.
func (o *NpmUpstreamRequest) SetAuthMode(v string) {
	o.AuthMode = &v
}

// GetAuthSecret returns the AuthSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NpmUpstreamRequest) GetAuthSecret() string {
	if o == nil || IsNil(o.AuthSecret.Get()) {
		var ret string
		return ret
	}
	return *o.AuthSecret.Get()
}

// GetAuthSecretOk returns a tuple with the AuthSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NpmUpstreamRequest) GetAuthSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthSecret.Get(), o.AuthSecret.IsSet()
}

// HasAuthSecret returns a boolean if a field has been set.
func (o *NpmUpstreamRequest) HasAuthSecret() bool {
	if o != nil && o.AuthSecret.IsSet() {
		return true
	}

	return false
}

// SetAuthSecret gets a reference to the given NullableString and assigns it to the AuthSecret field.
func (o *NpmUpstreamRequest) SetAuthSecret(v string) {
	o.AuthSecret.Set(&v)
}

// SetAuthSecretNil sets the value for AuthSecret to be an explicit nil
func (o *NpmUpstreamRequest) SetAuthSecretNil() {
	o.AuthSecret.Set(nil)
}

// UnsetAuthSecret ensures that no value is present for AuthSecret, not even an explicit nil
func (o *NpmUpstreamRequest) UnsetAuthSecret() {
	o.AuthSecret.Unset()
}

// GetAuthUsername returns the AuthUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NpmUpstreamRequest) GetAuthUsername() string {
	if o == nil || IsNil(o.AuthUsername.Get()) {
		var ret string
		return ret
	}
	return *o.AuthUsername.Get()
}

// GetAuthUsernameOk returns a tuple with the AuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NpmUpstreamRequest) GetAuthUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthUsername.Get(), o.AuthUsername.IsSet()
}

// HasAuthUsername returns a boolean if a field has been set.
func (o *NpmUpstreamRequest) HasAuthUsername() bool {
	if o != nil && o.AuthUsername.IsSet() {
		return true
	}

	return false
}

// SetAuthUsername gets a reference to the given NullableString and assigns it to the AuthUsername field.
func (o *NpmUpstreamRequest) SetAuthUsername(v string) {
	o.AuthUsername.Set(&v)
}

// SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil
func (o *NpmUpstreamRequest) SetAuthUsernameNil() {
	o.AuthUsername.Set(nil)
}

// UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
func (o *NpmUpstreamRequest) UnsetAuthUsername() {
	o.AuthUsername.Unset()
}

// GetExtraHeader1 returns the ExtraHeader1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NpmUpstreamRequest) GetExtraHeader1() string {
	if o == nil || IsNil(o.ExtraHeader1.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraHeader1.Get()
}

// GetExtraHeader1Ok returns a tuple with the ExtraHeader1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NpmUpstreamRequest) GetExtraHeader1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraHeader1.Get(), o.ExtraHeader1.IsSet()
}

// HasExtraHeader1 returns a boolean if a field has been set.
func (o *NpmUpstreamRequest) HasExtraHeader1() bool {
	if o != nil && o.ExtraHeader1.IsSet() {
		return true
	}

	return false
}

// SetExtraHeader1 gets a reference to the given NullableString and assigns it to the ExtraHeader1 field.
func (o *NpmUpstreamRequest) SetExtraHeader1(v string) {
	o.ExtraHeader1.Set(&v)
}

// SetExtraHeader1Nil sets the value for ExtraHeader1 to be an explicit nil
func (o *NpmUpstreamRequest) SetExtraHeader1Nil() {
	o.ExtraHeader1.Set(nil)
}

// UnsetExtraHeader1 ensures that no value is present for ExtraHeader1, not even an explicit nil
func (o *NpmUpstreamRequest) UnsetExtraHeader1() {
	o.ExtraHeader1.Unset()
}

// GetExtraHeader2 returns the ExtraHeader2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NpmUpstreamRequest) GetExtraHeader2() string {
	if o == nil || IsNil(o.ExtraHeader2.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraHeader2.Get()
}

// GetExtraHeader2Ok returns a tuple with the ExtraHeader2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NpmUpstreamRequest) GetExtraHeader2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraHeader2.Get(), o.ExtraHeader2.IsSet()
}

// HasExtraHeader2 returns a boolean if a field has been set.
func (o *NpmUpstreamRequest) HasExtraHeader2() bool {
	if o != nil && o.ExtraHeader2.IsSet() {
		return true
	}

	return false
}

// SetExtraHeader2 gets a reference to the given NullableString and assigns it to the ExtraHeader2 field.
func (o *NpmUpstreamRequest) SetExtraHeader2(v string) {
	o.ExtraHeader2.Set(&v)
}

// SetExtraHeader2Nil sets the value for ExtraHeader2 to be an explicit nil
func (o *NpmUpstreamRequest) SetExtraHeader2Nil() {
	o.ExtraHeader2.Set(nil)
}

// UnsetExtraHeader2 ensures that no value is present for ExtraHeader2, not even an explicit nil
func (o *NpmUpstreamRequest) UnsetExtraHeader2() {
	o.ExtraHeader2.Unset()
}

// GetExtraValue1 returns the ExtraValue1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NpmUpstreamRequest) GetExtraValue1() string {
	if o == nil || IsNil(o.ExtraValue1.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraValue1.Get()
}

// GetExtraValue1Ok returns a tuple with the ExtraValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NpmUpstreamRequest) GetExtraValue1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraValue1.Get(), o.ExtraValue1.IsSet()
}

// HasExtraValue1 returns a boolean if a field has been set.
func (o *NpmUpstreamRequest) HasExtraValue1() bool {
	if o != nil && o.ExtraValue1.IsSet() {
		return true
	}

	return false
}

// SetExtraValue1 gets a reference to the given NullableString and assigns it to the ExtraValue1 field.
func (o *NpmUpstreamRequest) SetExtraValue1(v string) {
	o.ExtraValue1.Set(&v)
}

// SetExtraValue1Nil sets the value for ExtraValue1 to be an explicit nil
func (o *NpmUpstreamRequest) SetExtraValue1Nil() {
	o.ExtraValue1.Set(nil)
}

// UnsetExtraValue1 ensures that no value is present for ExtraValue1, not even an explicit nil
func (o *NpmUpstreamRequest) UnsetExtraValue1() {
	o.ExtraValue1.Unset()
}

// GetExtraValue2 returns the ExtraValue2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NpmUpstreamRequest) GetExtraValue2() string {
	if o == nil || IsNil(o.ExtraValue2.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraValue2.Get()
}

// GetExtraValue2Ok returns a tuple with the ExtraValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NpmUpstreamRequest) GetExtraValue2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraValue2.Get(), o.ExtraValue2.IsSet()
}

// HasExtraValue2 returns a boolean if a field has been set.
func (o *NpmUpstreamRequest) HasExtraValue2() bool {
	if o != nil && o.ExtraValue2.IsSet() {
		return true
	}

	return false
}

// SetExtraValue2 gets a reference to the given NullableString and assigns it to the ExtraValue2 field.
func (o *NpmUpstreamRequest) SetExtraValue2(v string) {
	o.ExtraValue2.Set(&v)
}

// SetExtraValue2Nil sets the value for ExtraValue2 to be an explicit nil
func (o *NpmUpstreamRequest) SetExtraValue2Nil() {
	o.ExtraValue2.Set(nil)
}

// UnsetExtraValue2 ensures that no value is present for ExtraValue2, not even an explicit nil
func (o *NpmUpstreamRequest) UnsetExtraValue2() {
	o.ExtraValue2.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *NpmUpstreamRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NpmUpstreamRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *NpmUpstreamRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *NpmUpstreamRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *NpmUpstreamRequest) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NpmUpstreamRequest) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *NpmUpstreamRequest) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *NpmUpstreamRequest) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value
func (o *NpmUpstreamRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NpmUpstreamRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NpmUpstreamRequest) SetName(v string) {
	o.Name = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *NpmUpstreamRequest) GetPriority() int64 {
	if o == nil || IsNil(o.Priority) {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NpmUpstreamRequest) GetPriorityOk() (*int64, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *NpmUpstreamRequest) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *NpmUpstreamRequest) SetPriority(v int64) {
	o.Priority = &v
}

// GetUpstreamUrl returns the UpstreamUrl field value
func (o *NpmUpstreamRequest) GetUpstreamUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpstreamUrl
}

// GetUpstreamUrlOk returns a tuple with the UpstreamUrl field value
// and a boolean to check if the value has been set.
func (o *NpmUpstreamRequest) GetUpstreamUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpstreamUrl, true
}

// SetUpstreamUrl sets field value
func (o *NpmUpstreamRequest) SetUpstreamUrl(v string) {
	o.UpstreamUrl = v
}

// GetVerifySsl returns the VerifySsl field value if set, zero value otherwise.
func (o *NpmUpstreamRequest) GetVerifySsl() bool {
	if o == nil || IsNil(o.VerifySsl) {
		var ret bool
		return ret
	}
	return *o.VerifySsl
}

// GetVerifySslOk returns a tuple with the VerifySsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NpmUpstreamRequest) GetVerifySslOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifySsl) {
		return nil, false
	}
	return o.VerifySsl, true
}

// HasVerifySsl returns a boolean if a field has been set.
func (o *NpmUpstreamRequest) HasVerifySsl() bool {
	if o != nil && !IsNil(o.VerifySsl) {
		return true
	}

	return false
}

// SetVerifySsl gets a reference to the given bool and assigns it to the VerifySsl field.
func (o *NpmUpstreamRequest) SetVerifySsl(v bool) {
	o.VerifySsl = &v
}

func (o NpmUpstreamRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NpmUpstreamRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthMode) {
		toSerialize["auth_mode"] = o.AuthMode
	}
	if o.AuthSecret.IsSet() {
		toSerialize["auth_secret"] = o.AuthSecret.Get()
	}
	if o.AuthUsername.IsSet() {
		toSerialize["auth_username"] = o.AuthUsername.Get()
	}
	if o.ExtraHeader1.IsSet() {
		toSerialize["extra_header_1"] = o.ExtraHeader1.Get()
	}
	if o.ExtraHeader2.IsSet() {
		toSerialize["extra_header_2"] = o.ExtraHeader2.Get()
	}
	if o.ExtraValue1.IsSet() {
		toSerialize["extra_value_1"] = o.ExtraValue1.Get()
	}
	if o.ExtraValue2.IsSet() {
		toSerialize["extra_value_2"] = o.ExtraValue2.Get()
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	toSerialize["upstream_url"] = o.UpstreamUrl
	if !IsNil(o.VerifySsl) {
		toSerialize["verify_ssl"] = o.VerifySsl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NpmUpstreamRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"upstream_url",
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

	varNpmUpstreamRequest := _NpmUpstreamRequest{}

	err = json.Unmarshal(data, &varNpmUpstreamRequest)

	if err != nil {
		return err
	}

	*o = NpmUpstreamRequest(varNpmUpstreamRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "auth_mode")
		delete(additionalProperties, "auth_secret")
		delete(additionalProperties, "auth_username")
		delete(additionalProperties, "extra_header_1")
		delete(additionalProperties, "extra_header_2")
		delete(additionalProperties, "extra_value_1")
		delete(additionalProperties, "extra_value_2")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "name")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "upstream_url")
		delete(additionalProperties, "verify_ssl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNpmUpstreamRequest struct {
	value *NpmUpstreamRequest
	isSet bool
}

func (v NullableNpmUpstreamRequest) Get() *NpmUpstreamRequest {
	return v.value
}

func (v *NullableNpmUpstreamRequest) Set(val *NpmUpstreamRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNpmUpstreamRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNpmUpstreamRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNpmUpstreamRequest(val *NpmUpstreamRequest) *NullableNpmUpstreamRequest {
	return &NullableNpmUpstreamRequest{value: val, isSet: true}
}

func (v NullableNpmUpstreamRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNpmUpstreamRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
