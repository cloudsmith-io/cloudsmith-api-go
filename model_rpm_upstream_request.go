/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.372.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the RpmUpstreamRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmUpstreamRequest{}

// RpmUpstreamRequest struct for RpmUpstreamRequest
type RpmUpstreamRequest struct {
	// The authentication mode to use when accessing this upstream.
	AuthMode *string `json:"auth_mode,omitempty"`
	// Secret to provide with requests to upstream.
	AuthSecret NullableString `json:"auth_secret,omitempty"`
	// Username to provide with requests to upstream.
	AuthUsername NullableString `json:"auth_username,omitempty"`
	// The distribution version that packages found on this upstream will be associated with.
	DistroVersion string `json:"distro_version"`
	// The key for extra header #1 to send to upstream.
	ExtraHeader1 NullableString `json:"extra_header_1,omitempty"`
	// The key for extra header #2 to send to upstream.
	ExtraHeader2 NullableString `json:"extra_header_2,omitempty"`
	// The value for extra header #1 to send to upstream. This is stored as plaintext, and is NOT encrypted.
	ExtraValue1 NullableString `json:"extra_value_1,omitempty"`
	// The value for extra header #2 to send to upstream. This is stored as plaintext, and is NOT encrypted.
	ExtraValue2 NullableString `json:"extra_value_2,omitempty"`
	// When checked, source packages will be available from this upstream.
	IncludeSources *bool `json:"include_sources,omitempty"`
	// Whether or not this upstream is active and ready for requests.
	IsActive *bool `json:"is_active,omitempty"`
	// The mode that this upstream should operate in. Upstream sources can be used to proxy resolved packages, as well as operate in a proxy/cache or cache only mode.
	Mode *string `json:"mode,omitempty"`
	// A descriptive name for this upstream source. A shortened version of this name will be used for tagging cached packages retrieved from this upstream.
	Name string `json:"name"`
	// Upstream sources are selected for resolving requests by sequential order (1..n), followed by creation date.
	Priority *int64 `json:"priority,omitempty"`
	// The URL for this upstream source. This must be a fully qualified URL including any path elements required to reach the root of the repository.
	UpstreamUrl string `json:"upstream_url"`
	// If enabled, SSL certificates are verified when requests are made to this upstream. It's recommended to leave this enabled for all public sources to help mitigate Man-In-The-Middle (MITM) attacks. Please note this only applies to HTTPS upstreams.
	VerifySsl *bool `json:"verify_ssl,omitempty"`
}

// NewRpmUpstreamRequest instantiates a new RpmUpstreamRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmUpstreamRequest(distroVersion string, name string, upstreamUrl string) *RpmUpstreamRequest {
	this := RpmUpstreamRequest{}
	var authMode string = "None"
	this.AuthMode = &authMode
	this.DistroVersion = distroVersion
	var mode string = "Proxy Only"
	this.Mode = &mode
	this.Name = name
	this.UpstreamUrl = upstreamUrl
	return &this
}

// NewRpmUpstreamRequestWithDefaults instantiates a new RpmUpstreamRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmUpstreamRequestWithDefaults() *RpmUpstreamRequest {
	this := RpmUpstreamRequest{}
	var authMode string = "None"
	this.AuthMode = &authMode
	var mode string = "Proxy Only"
	this.Mode = &mode
	return &this
}

// GetAuthMode returns the AuthMode field value if set, zero value otherwise.
func (o *RpmUpstreamRequest) GetAuthMode() string {
	if o == nil || IsNil(o.AuthMode) {
		var ret string
		return ret
	}
	return *o.AuthMode
}

// GetAuthModeOk returns a tuple with the AuthMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpstreamRequest) GetAuthModeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthMode) {
		return nil, false
	}
	return o.AuthMode, true
}

// HasAuthMode returns a boolean if a field has been set.
func (o *RpmUpstreamRequest) HasAuthMode() bool {
	if o != nil && !IsNil(o.AuthMode) {
		return true
	}

	return false
}

// SetAuthMode gets a reference to the given string and assigns it to the AuthMode field.
func (o *RpmUpstreamRequest) SetAuthMode(v string) {
	o.AuthMode = &v
}

// GetAuthSecret returns the AuthSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmUpstreamRequest) GetAuthSecret() string {
	if o == nil || IsNil(o.AuthSecret.Get()) {
		var ret string
		return ret
	}
	return *o.AuthSecret.Get()
}

// GetAuthSecretOk returns a tuple with the AuthSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmUpstreamRequest) GetAuthSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthSecret.Get(), o.AuthSecret.IsSet()
}

// HasAuthSecret returns a boolean if a field has been set.
func (o *RpmUpstreamRequest) HasAuthSecret() bool {
	if o != nil && o.AuthSecret.IsSet() {
		return true
	}

	return false
}

// SetAuthSecret gets a reference to the given NullableString and assigns it to the AuthSecret field.
func (o *RpmUpstreamRequest) SetAuthSecret(v string) {
	o.AuthSecret.Set(&v)
}

// SetAuthSecretNil sets the value for AuthSecret to be an explicit nil
func (o *RpmUpstreamRequest) SetAuthSecretNil() {
	o.AuthSecret.Set(nil)
}

// UnsetAuthSecret ensures that no value is present for AuthSecret, not even an explicit nil
func (o *RpmUpstreamRequest) UnsetAuthSecret() {
	o.AuthSecret.Unset()
}

// GetAuthUsername returns the AuthUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmUpstreamRequest) GetAuthUsername() string {
	if o == nil || IsNil(o.AuthUsername.Get()) {
		var ret string
		return ret
	}
	return *o.AuthUsername.Get()
}

// GetAuthUsernameOk returns a tuple with the AuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmUpstreamRequest) GetAuthUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthUsername.Get(), o.AuthUsername.IsSet()
}

// HasAuthUsername returns a boolean if a field has been set.
func (o *RpmUpstreamRequest) HasAuthUsername() bool {
	if o != nil && o.AuthUsername.IsSet() {
		return true
	}

	return false
}

// SetAuthUsername gets a reference to the given NullableString and assigns it to the AuthUsername field.
func (o *RpmUpstreamRequest) SetAuthUsername(v string) {
	o.AuthUsername.Set(&v)
}

// SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil
func (o *RpmUpstreamRequest) SetAuthUsernameNil() {
	o.AuthUsername.Set(nil)
}

// UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
func (o *RpmUpstreamRequest) UnsetAuthUsername() {
	o.AuthUsername.Unset()
}

// GetDistroVersion returns the DistroVersion field value
func (o *RpmUpstreamRequest) GetDistroVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DistroVersion
}

// GetDistroVersionOk returns a tuple with the DistroVersion field value
// and a boolean to check if the value has been set.
func (o *RpmUpstreamRequest) GetDistroVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DistroVersion, true
}

// SetDistroVersion sets field value
func (o *RpmUpstreamRequest) SetDistroVersion(v string) {
	o.DistroVersion = v
}

// GetExtraHeader1 returns the ExtraHeader1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmUpstreamRequest) GetExtraHeader1() string {
	if o == nil || IsNil(o.ExtraHeader1.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraHeader1.Get()
}

// GetExtraHeader1Ok returns a tuple with the ExtraHeader1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmUpstreamRequest) GetExtraHeader1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraHeader1.Get(), o.ExtraHeader1.IsSet()
}

// HasExtraHeader1 returns a boolean if a field has been set.
func (o *RpmUpstreamRequest) HasExtraHeader1() bool {
	if o != nil && o.ExtraHeader1.IsSet() {
		return true
	}

	return false
}

// SetExtraHeader1 gets a reference to the given NullableString and assigns it to the ExtraHeader1 field.
func (o *RpmUpstreamRequest) SetExtraHeader1(v string) {
	o.ExtraHeader1.Set(&v)
}

// SetExtraHeader1Nil sets the value for ExtraHeader1 to be an explicit nil
func (o *RpmUpstreamRequest) SetExtraHeader1Nil() {
	o.ExtraHeader1.Set(nil)
}

// UnsetExtraHeader1 ensures that no value is present for ExtraHeader1, not even an explicit nil
func (o *RpmUpstreamRequest) UnsetExtraHeader1() {
	o.ExtraHeader1.Unset()
}

// GetExtraHeader2 returns the ExtraHeader2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmUpstreamRequest) GetExtraHeader2() string {
	if o == nil || IsNil(o.ExtraHeader2.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraHeader2.Get()
}

// GetExtraHeader2Ok returns a tuple with the ExtraHeader2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmUpstreamRequest) GetExtraHeader2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraHeader2.Get(), o.ExtraHeader2.IsSet()
}

// HasExtraHeader2 returns a boolean if a field has been set.
func (o *RpmUpstreamRequest) HasExtraHeader2() bool {
	if o != nil && o.ExtraHeader2.IsSet() {
		return true
	}

	return false
}

// SetExtraHeader2 gets a reference to the given NullableString and assigns it to the ExtraHeader2 field.
func (o *RpmUpstreamRequest) SetExtraHeader2(v string) {
	o.ExtraHeader2.Set(&v)
}

// SetExtraHeader2Nil sets the value for ExtraHeader2 to be an explicit nil
func (o *RpmUpstreamRequest) SetExtraHeader2Nil() {
	o.ExtraHeader2.Set(nil)
}

// UnsetExtraHeader2 ensures that no value is present for ExtraHeader2, not even an explicit nil
func (o *RpmUpstreamRequest) UnsetExtraHeader2() {
	o.ExtraHeader2.Unset()
}

// GetExtraValue1 returns the ExtraValue1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmUpstreamRequest) GetExtraValue1() string {
	if o == nil || IsNil(o.ExtraValue1.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraValue1.Get()
}

// GetExtraValue1Ok returns a tuple with the ExtraValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmUpstreamRequest) GetExtraValue1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraValue1.Get(), o.ExtraValue1.IsSet()
}

// HasExtraValue1 returns a boolean if a field has been set.
func (o *RpmUpstreamRequest) HasExtraValue1() bool {
	if o != nil && o.ExtraValue1.IsSet() {
		return true
	}

	return false
}

// SetExtraValue1 gets a reference to the given NullableString and assigns it to the ExtraValue1 field.
func (o *RpmUpstreamRequest) SetExtraValue1(v string) {
	o.ExtraValue1.Set(&v)
}

// SetExtraValue1Nil sets the value for ExtraValue1 to be an explicit nil
func (o *RpmUpstreamRequest) SetExtraValue1Nil() {
	o.ExtraValue1.Set(nil)
}

// UnsetExtraValue1 ensures that no value is present for ExtraValue1, not even an explicit nil
func (o *RpmUpstreamRequest) UnsetExtraValue1() {
	o.ExtraValue1.Unset()
}

// GetExtraValue2 returns the ExtraValue2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmUpstreamRequest) GetExtraValue2() string {
	if o == nil || IsNil(o.ExtraValue2.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraValue2.Get()
}

// GetExtraValue2Ok returns a tuple with the ExtraValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmUpstreamRequest) GetExtraValue2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraValue2.Get(), o.ExtraValue2.IsSet()
}

// HasExtraValue2 returns a boolean if a field has been set.
func (o *RpmUpstreamRequest) HasExtraValue2() bool {
	if o != nil && o.ExtraValue2.IsSet() {
		return true
	}

	return false
}

// SetExtraValue2 gets a reference to the given NullableString and assigns it to the ExtraValue2 field.
func (o *RpmUpstreamRequest) SetExtraValue2(v string) {
	o.ExtraValue2.Set(&v)
}

// SetExtraValue2Nil sets the value for ExtraValue2 to be an explicit nil
func (o *RpmUpstreamRequest) SetExtraValue2Nil() {
	o.ExtraValue2.Set(nil)
}

// UnsetExtraValue2 ensures that no value is present for ExtraValue2, not even an explicit nil
func (o *RpmUpstreamRequest) UnsetExtraValue2() {
	o.ExtraValue2.Unset()
}

// GetIncludeSources returns the IncludeSources field value if set, zero value otherwise.
func (o *RpmUpstreamRequest) GetIncludeSources() bool {
	if o == nil || IsNil(o.IncludeSources) {
		var ret bool
		return ret
	}
	return *o.IncludeSources
}

// GetIncludeSourcesOk returns a tuple with the IncludeSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpstreamRequest) GetIncludeSourcesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeSources) {
		return nil, false
	}
	return o.IncludeSources, true
}

// HasIncludeSources returns a boolean if a field has been set.
func (o *RpmUpstreamRequest) HasIncludeSources() bool {
	if o != nil && !IsNil(o.IncludeSources) {
		return true
	}

	return false
}

// SetIncludeSources gets a reference to the given bool and assigns it to the IncludeSources field.
func (o *RpmUpstreamRequest) SetIncludeSources(v bool) {
	o.IncludeSources = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *RpmUpstreamRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpstreamRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *RpmUpstreamRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *RpmUpstreamRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *RpmUpstreamRequest) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpstreamRequest) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *RpmUpstreamRequest) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *RpmUpstreamRequest) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value
func (o *RpmUpstreamRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RpmUpstreamRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RpmUpstreamRequest) SetName(v string) {
	o.Name = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *RpmUpstreamRequest) GetPriority() int64 {
	if o == nil || IsNil(o.Priority) {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpstreamRequest) GetPriorityOk() (*int64, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *RpmUpstreamRequest) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *RpmUpstreamRequest) SetPriority(v int64) {
	o.Priority = &v
}

// GetUpstreamUrl returns the UpstreamUrl field value
func (o *RpmUpstreamRequest) GetUpstreamUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpstreamUrl
}

// GetUpstreamUrlOk returns a tuple with the UpstreamUrl field value
// and a boolean to check if the value has been set.
func (o *RpmUpstreamRequest) GetUpstreamUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpstreamUrl, true
}

// SetUpstreamUrl sets field value
func (o *RpmUpstreamRequest) SetUpstreamUrl(v string) {
	o.UpstreamUrl = v
}

// GetVerifySsl returns the VerifySsl field value if set, zero value otherwise.
func (o *RpmUpstreamRequest) GetVerifySsl() bool {
	if o == nil || IsNil(o.VerifySsl) {
		var ret bool
		return ret
	}
	return *o.VerifySsl
}

// GetVerifySslOk returns a tuple with the VerifySsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpstreamRequest) GetVerifySslOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifySsl) {
		return nil, false
	}
	return o.VerifySsl, true
}

// HasVerifySsl returns a boolean if a field has been set.
func (o *RpmUpstreamRequest) HasVerifySsl() bool {
	if o != nil && !IsNil(o.VerifySsl) {
		return true
	}

	return false
}

// SetVerifySsl gets a reference to the given bool and assigns it to the VerifySsl field.
func (o *RpmUpstreamRequest) SetVerifySsl(v bool) {
	o.VerifySsl = &v
}

func (o RpmUpstreamRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmUpstreamRequest) ToMap() (map[string]interface{}, error) {
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
	toSerialize["distro_version"] = o.DistroVersion
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
	if !IsNil(o.IncludeSources) {
		toSerialize["include_sources"] = o.IncludeSources
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
	return toSerialize, nil
}

type NullableRpmUpstreamRequest struct {
	value *RpmUpstreamRequest
	isSet bool
}

func (v NullableRpmUpstreamRequest) Get() *RpmUpstreamRequest {
	return v.value
}

func (v *NullableRpmUpstreamRequest) Set(val *RpmUpstreamRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmUpstreamRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmUpstreamRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmUpstreamRequest(val *RpmUpstreamRequest) *NullableRpmUpstreamRequest {
	return &NullableRpmUpstreamRequest{value: val, isSet: true}
}

func (v NullableRpmUpstreamRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmUpstreamRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
