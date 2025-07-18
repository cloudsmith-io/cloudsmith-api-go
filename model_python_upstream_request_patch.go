/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.736.13
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the PythonUpstreamRequestPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PythonUpstreamRequestPatch{}

// PythonUpstreamRequestPatch struct for PythonUpstreamRequestPatch
type PythonUpstreamRequestPatch struct {
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
	Name *string `json:"name,omitempty" validate:"regexp=^\\\\w[\\\\w \\\\-'\\\\.\\/()]+$"`
	// Upstream sources are selected for resolving requests by sequential order (1..n), followed by creation date.
	Priority *int64 `json:"priority,omitempty"`
	// The URL for this upstream source. This must be a fully qualified URL including any path elements required to reach the root of the repository.
	UpstreamUrl *string `json:"upstream_url,omitempty"`
	// If enabled, SSL certificates are verified when requests are made to this upstream. It's recommended to leave this enabled for all public sources to help mitigate Man-In-The-Middle (MITM) attacks. Please note this only applies to HTTPS upstreams.
	VerifySsl            *bool `json:"verify_ssl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PythonUpstreamRequestPatch PythonUpstreamRequestPatch

// NewPythonUpstreamRequestPatch instantiates a new PythonUpstreamRequestPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPythonUpstreamRequestPatch() *PythonUpstreamRequestPatch {
	this := PythonUpstreamRequestPatch{}
	var authMode string = "None"
	this.AuthMode = &authMode
	var mode string = "Proxy Only"
	this.Mode = &mode
	return &this
}

// NewPythonUpstreamRequestPatchWithDefaults instantiates a new PythonUpstreamRequestPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPythonUpstreamRequestPatchWithDefaults() *PythonUpstreamRequestPatch {
	this := PythonUpstreamRequestPatch{}
	var authMode string = "None"
	this.AuthMode = &authMode
	var mode string = "Proxy Only"
	this.Mode = &mode
	return &this
}

// GetAuthMode returns the AuthMode field value if set, zero value otherwise.
func (o *PythonUpstreamRequestPatch) GetAuthMode() string {
	if o == nil || IsNil(o.AuthMode) {
		var ret string
		return ret
	}
	return *o.AuthMode
}

// GetAuthModeOk returns a tuple with the AuthMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonUpstreamRequestPatch) GetAuthModeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthMode) {
		return nil, false
	}
	return o.AuthMode, true
}

// HasAuthMode returns a boolean if a field has been set.
func (o *PythonUpstreamRequestPatch) HasAuthMode() bool {
	if o != nil && !IsNil(o.AuthMode) {
		return true
	}

	return false
}

// SetAuthMode gets a reference to the given string and assigns it to the AuthMode field.
func (o *PythonUpstreamRequestPatch) SetAuthMode(v string) {
	o.AuthMode = &v
}

// GetAuthSecret returns the AuthSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PythonUpstreamRequestPatch) GetAuthSecret() string {
	if o == nil || IsNil(o.AuthSecret.Get()) {
		var ret string
		return ret
	}
	return *o.AuthSecret.Get()
}

// GetAuthSecretOk returns a tuple with the AuthSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PythonUpstreamRequestPatch) GetAuthSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthSecret.Get(), o.AuthSecret.IsSet()
}

// HasAuthSecret returns a boolean if a field has been set.
func (o *PythonUpstreamRequestPatch) HasAuthSecret() bool {
	if o != nil && o.AuthSecret.IsSet() {
		return true
	}

	return false
}

// SetAuthSecret gets a reference to the given NullableString and assigns it to the AuthSecret field.
func (o *PythonUpstreamRequestPatch) SetAuthSecret(v string) {
	o.AuthSecret.Set(&v)
}

// SetAuthSecretNil sets the value for AuthSecret to be an explicit nil
func (o *PythonUpstreamRequestPatch) SetAuthSecretNil() {
	o.AuthSecret.Set(nil)
}

// UnsetAuthSecret ensures that no value is present for AuthSecret, not even an explicit nil
func (o *PythonUpstreamRequestPatch) UnsetAuthSecret() {
	o.AuthSecret.Unset()
}

// GetAuthUsername returns the AuthUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PythonUpstreamRequestPatch) GetAuthUsername() string {
	if o == nil || IsNil(o.AuthUsername.Get()) {
		var ret string
		return ret
	}
	return *o.AuthUsername.Get()
}

// GetAuthUsernameOk returns a tuple with the AuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PythonUpstreamRequestPatch) GetAuthUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthUsername.Get(), o.AuthUsername.IsSet()
}

// HasAuthUsername returns a boolean if a field has been set.
func (o *PythonUpstreamRequestPatch) HasAuthUsername() bool {
	if o != nil && o.AuthUsername.IsSet() {
		return true
	}

	return false
}

// SetAuthUsername gets a reference to the given NullableString and assigns it to the AuthUsername field.
func (o *PythonUpstreamRequestPatch) SetAuthUsername(v string) {
	o.AuthUsername.Set(&v)
}

// SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil
func (o *PythonUpstreamRequestPatch) SetAuthUsernameNil() {
	o.AuthUsername.Set(nil)
}

// UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
func (o *PythonUpstreamRequestPatch) UnsetAuthUsername() {
	o.AuthUsername.Unset()
}

// GetExtraHeader1 returns the ExtraHeader1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PythonUpstreamRequestPatch) GetExtraHeader1() string {
	if o == nil || IsNil(o.ExtraHeader1.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraHeader1.Get()
}

// GetExtraHeader1Ok returns a tuple with the ExtraHeader1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PythonUpstreamRequestPatch) GetExtraHeader1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraHeader1.Get(), o.ExtraHeader1.IsSet()
}

// HasExtraHeader1 returns a boolean if a field has been set.
func (o *PythonUpstreamRequestPatch) HasExtraHeader1() bool {
	if o != nil && o.ExtraHeader1.IsSet() {
		return true
	}

	return false
}

// SetExtraHeader1 gets a reference to the given NullableString and assigns it to the ExtraHeader1 field.
func (o *PythonUpstreamRequestPatch) SetExtraHeader1(v string) {
	o.ExtraHeader1.Set(&v)
}

// SetExtraHeader1Nil sets the value for ExtraHeader1 to be an explicit nil
func (o *PythonUpstreamRequestPatch) SetExtraHeader1Nil() {
	o.ExtraHeader1.Set(nil)
}

// UnsetExtraHeader1 ensures that no value is present for ExtraHeader1, not even an explicit nil
func (o *PythonUpstreamRequestPatch) UnsetExtraHeader1() {
	o.ExtraHeader1.Unset()
}

// GetExtraHeader2 returns the ExtraHeader2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PythonUpstreamRequestPatch) GetExtraHeader2() string {
	if o == nil || IsNil(o.ExtraHeader2.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraHeader2.Get()
}

// GetExtraHeader2Ok returns a tuple with the ExtraHeader2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PythonUpstreamRequestPatch) GetExtraHeader2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraHeader2.Get(), o.ExtraHeader2.IsSet()
}

// HasExtraHeader2 returns a boolean if a field has been set.
func (o *PythonUpstreamRequestPatch) HasExtraHeader2() bool {
	if o != nil && o.ExtraHeader2.IsSet() {
		return true
	}

	return false
}

// SetExtraHeader2 gets a reference to the given NullableString and assigns it to the ExtraHeader2 field.
func (o *PythonUpstreamRequestPatch) SetExtraHeader2(v string) {
	o.ExtraHeader2.Set(&v)
}

// SetExtraHeader2Nil sets the value for ExtraHeader2 to be an explicit nil
func (o *PythonUpstreamRequestPatch) SetExtraHeader2Nil() {
	o.ExtraHeader2.Set(nil)
}

// UnsetExtraHeader2 ensures that no value is present for ExtraHeader2, not even an explicit nil
func (o *PythonUpstreamRequestPatch) UnsetExtraHeader2() {
	o.ExtraHeader2.Unset()
}

// GetExtraValue1 returns the ExtraValue1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PythonUpstreamRequestPatch) GetExtraValue1() string {
	if o == nil || IsNil(o.ExtraValue1.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraValue1.Get()
}

// GetExtraValue1Ok returns a tuple with the ExtraValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PythonUpstreamRequestPatch) GetExtraValue1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraValue1.Get(), o.ExtraValue1.IsSet()
}

// HasExtraValue1 returns a boolean if a field has been set.
func (o *PythonUpstreamRequestPatch) HasExtraValue1() bool {
	if o != nil && o.ExtraValue1.IsSet() {
		return true
	}

	return false
}

// SetExtraValue1 gets a reference to the given NullableString and assigns it to the ExtraValue1 field.
func (o *PythonUpstreamRequestPatch) SetExtraValue1(v string) {
	o.ExtraValue1.Set(&v)
}

// SetExtraValue1Nil sets the value for ExtraValue1 to be an explicit nil
func (o *PythonUpstreamRequestPatch) SetExtraValue1Nil() {
	o.ExtraValue1.Set(nil)
}

// UnsetExtraValue1 ensures that no value is present for ExtraValue1, not even an explicit nil
func (o *PythonUpstreamRequestPatch) UnsetExtraValue1() {
	o.ExtraValue1.Unset()
}

// GetExtraValue2 returns the ExtraValue2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PythonUpstreamRequestPatch) GetExtraValue2() string {
	if o == nil || IsNil(o.ExtraValue2.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraValue2.Get()
}

// GetExtraValue2Ok returns a tuple with the ExtraValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PythonUpstreamRequestPatch) GetExtraValue2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraValue2.Get(), o.ExtraValue2.IsSet()
}

// HasExtraValue2 returns a boolean if a field has been set.
func (o *PythonUpstreamRequestPatch) HasExtraValue2() bool {
	if o != nil && o.ExtraValue2.IsSet() {
		return true
	}

	return false
}

// SetExtraValue2 gets a reference to the given NullableString and assigns it to the ExtraValue2 field.
func (o *PythonUpstreamRequestPatch) SetExtraValue2(v string) {
	o.ExtraValue2.Set(&v)
}

// SetExtraValue2Nil sets the value for ExtraValue2 to be an explicit nil
func (o *PythonUpstreamRequestPatch) SetExtraValue2Nil() {
	o.ExtraValue2.Set(nil)
}

// UnsetExtraValue2 ensures that no value is present for ExtraValue2, not even an explicit nil
func (o *PythonUpstreamRequestPatch) UnsetExtraValue2() {
	o.ExtraValue2.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *PythonUpstreamRequestPatch) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonUpstreamRequestPatch) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *PythonUpstreamRequestPatch) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *PythonUpstreamRequestPatch) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *PythonUpstreamRequestPatch) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonUpstreamRequestPatch) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *PythonUpstreamRequestPatch) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *PythonUpstreamRequestPatch) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PythonUpstreamRequestPatch) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonUpstreamRequestPatch) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PythonUpstreamRequestPatch) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PythonUpstreamRequestPatch) SetName(v string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *PythonUpstreamRequestPatch) GetPriority() int64 {
	if o == nil || IsNil(o.Priority) {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonUpstreamRequestPatch) GetPriorityOk() (*int64, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *PythonUpstreamRequestPatch) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *PythonUpstreamRequestPatch) SetPriority(v int64) {
	o.Priority = &v
}

// GetUpstreamUrl returns the UpstreamUrl field value if set, zero value otherwise.
func (o *PythonUpstreamRequestPatch) GetUpstreamUrl() string {
	if o == nil || IsNil(o.UpstreamUrl) {
		var ret string
		return ret
	}
	return *o.UpstreamUrl
}

// GetUpstreamUrlOk returns a tuple with the UpstreamUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonUpstreamRequestPatch) GetUpstreamUrlOk() (*string, bool) {
	if o == nil || IsNil(o.UpstreamUrl) {
		return nil, false
	}
	return o.UpstreamUrl, true
}

// HasUpstreamUrl returns a boolean if a field has been set.
func (o *PythonUpstreamRequestPatch) HasUpstreamUrl() bool {
	if o != nil && !IsNil(o.UpstreamUrl) {
		return true
	}

	return false
}

// SetUpstreamUrl gets a reference to the given string and assigns it to the UpstreamUrl field.
func (o *PythonUpstreamRequestPatch) SetUpstreamUrl(v string) {
	o.UpstreamUrl = &v
}

// GetVerifySsl returns the VerifySsl field value if set, zero value otherwise.
func (o *PythonUpstreamRequestPatch) GetVerifySsl() bool {
	if o == nil || IsNil(o.VerifySsl) {
		var ret bool
		return ret
	}
	return *o.VerifySsl
}

// GetVerifySslOk returns a tuple with the VerifySsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PythonUpstreamRequestPatch) GetVerifySslOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifySsl) {
		return nil, false
	}
	return o.VerifySsl, true
}

// HasVerifySsl returns a boolean if a field has been set.
func (o *PythonUpstreamRequestPatch) HasVerifySsl() bool {
	if o != nil && !IsNil(o.VerifySsl) {
		return true
	}

	return false
}

// SetVerifySsl gets a reference to the given bool and assigns it to the VerifySsl field.
func (o *PythonUpstreamRequestPatch) SetVerifySsl(v bool) {
	o.VerifySsl = &v
}

func (o PythonUpstreamRequestPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PythonUpstreamRequestPatch) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.UpstreamUrl) {
		toSerialize["upstream_url"] = o.UpstreamUrl
	}
	if !IsNil(o.VerifySsl) {
		toSerialize["verify_ssl"] = o.VerifySsl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PythonUpstreamRequestPatch) UnmarshalJSON(data []byte) (err error) {
	varPythonUpstreamRequestPatch := _PythonUpstreamRequestPatch{}

	err = json.Unmarshal(data, &varPythonUpstreamRequestPatch)

	if err != nil {
		return err
	}

	*o = PythonUpstreamRequestPatch(varPythonUpstreamRequestPatch)

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

type NullablePythonUpstreamRequestPatch struct {
	value *PythonUpstreamRequestPatch
	isSet bool
}

func (v NullablePythonUpstreamRequestPatch) Get() *PythonUpstreamRequestPatch {
	return v.value
}

func (v *NullablePythonUpstreamRequestPatch) Set(val *PythonUpstreamRequestPatch) {
	v.value = val
	v.isSet = true
}

func (v NullablePythonUpstreamRequestPatch) IsSet() bool {
	return v.isSet
}

func (v *NullablePythonUpstreamRequestPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePythonUpstreamRequestPatch(val *PythonUpstreamRequestPatch) *NullablePythonUpstreamRequestPatch {
	return &NullablePythonUpstreamRequestPatch{value: val, isSet: true}
}

func (v NullablePythonUpstreamRequestPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePythonUpstreamRequestPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
