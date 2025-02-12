/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.617.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the RepositoryWebhookRequestPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryWebhookRequestPatch{}

// RepositoryWebhookRequestPatch struct for RepositoryWebhookRequestPatch
type RepositoryWebhookRequestPatch struct {
	Events []string `json:"events,omitempty"`
	// If enabled, the webhook will trigger on subscribed events and send payloads to the configured target URL.
	IsActive *bool `json:"is_active,omitempty"`
	// The package-based search query for webhooks to fire. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. If a package does not match, the webhook will not fire.
	PackageQuery NullableString `json:"package_query,omitempty"`
	// The format of the payloads for webhook requests. Valid options are: (0) JSON, (1) JSON array, (2) form encoded JSON and (3) Handlebars template.
	RequestBodyFormat *int64 `json:"request_body_format,omitempty"`
	// The format of the payloads for webhook requests. Valid options are: (0) Generic/user defined, (1) JSON and (2) XML.
	RequestBodyTemplateFormat *int64 `json:"request_body_template_format,omitempty"`
	// The value that will be sent for the 'Content Type' header.
	RequestContentType NullableString `json:"request_content_type,omitempty"`
	// The header to send the predefined secret in. This must be unique from existing headers or it won't be sent. You can use this as a form of authentication on the endpoint side.
	SecretHeader NullableString `json:"secret_header,omitempty" validate:"regexp=^[-\\\\w]+$"`
	// The value for the predefined secret (note: this is treated as a passphrase and is encrypted when we store it). You can use this as a form of authentication on the endpoint side.
	SecretValue NullableString `json:"secret_value,omitempty" validate:"regexp=^[^\\\\n\\\\r]+$"`
	// The value for the signature key - This is used to generate an HMAC-based hex digest of the request body, which we send as the X-Cloudsmith-Signature header so that you can ensure that the request wasn't modified by a malicious party (note: this is treated as a passphrase and is encrypted when we store it).
	SignatureKey *string `json:"signature_key,omitempty"`
	// The destination URL that webhook payloads will be POST'ed to.
	TargetUrl *string           `json:"target_url,omitempty"`
	Templates []WebhookTemplate `json:"templates,omitempty"`
	// If enabled, SSL certificates is verified when webhooks are sent. It's recommended to leave this enabled as not verifying the integrity of SSL certificates leaves you susceptible to Man-in-the-Middle (MITM) attacks.
	VerifySsl            *bool `json:"verify_ssl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RepositoryWebhookRequestPatch RepositoryWebhookRequestPatch

// NewRepositoryWebhookRequestPatch instantiates a new RepositoryWebhookRequestPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryWebhookRequestPatch() *RepositoryWebhookRequestPatch {
	this := RepositoryWebhookRequestPatch{}
	return &this
}

// NewRepositoryWebhookRequestPatchWithDefaults instantiates a new RepositoryWebhookRequestPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryWebhookRequestPatchWithDefaults() *RepositoryWebhookRequestPatch {
	this := RepositoryWebhookRequestPatch{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryWebhookRequestPatch) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhookRequestPatch) GetEventsOk() ([]string, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *RepositoryWebhookRequestPatch) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *RepositoryWebhookRequestPatch) SetEvents(v []string) {
	o.Events = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *RepositoryWebhookRequestPatch) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhookRequestPatch) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *RepositoryWebhookRequestPatch) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *RepositoryWebhookRequestPatch) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetPackageQuery returns the PackageQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryWebhookRequestPatch) GetPackageQuery() string {
	if o == nil || IsNil(o.PackageQuery.Get()) {
		var ret string
		return ret
	}
	return *o.PackageQuery.Get()
}

// GetPackageQueryOk returns a tuple with the PackageQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhookRequestPatch) GetPackageQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageQuery.Get(), o.PackageQuery.IsSet()
}

// HasPackageQuery returns a boolean if a field has been set.
func (o *RepositoryWebhookRequestPatch) HasPackageQuery() bool {
	if o != nil && o.PackageQuery.IsSet() {
		return true
	}

	return false
}

// SetPackageQuery gets a reference to the given NullableString and assigns it to the PackageQuery field.
func (o *RepositoryWebhookRequestPatch) SetPackageQuery(v string) {
	o.PackageQuery.Set(&v)
}

// SetPackageQueryNil sets the value for PackageQuery to be an explicit nil
func (o *RepositoryWebhookRequestPatch) SetPackageQueryNil() {
	o.PackageQuery.Set(nil)
}

// UnsetPackageQuery ensures that no value is present for PackageQuery, not even an explicit nil
func (o *RepositoryWebhookRequestPatch) UnsetPackageQuery() {
	o.PackageQuery.Unset()
}

// GetRequestBodyFormat returns the RequestBodyFormat field value if set, zero value otherwise.
func (o *RepositoryWebhookRequestPatch) GetRequestBodyFormat() int64 {
	if o == nil || IsNil(o.RequestBodyFormat) {
		var ret int64
		return ret
	}
	return *o.RequestBodyFormat
}

// GetRequestBodyFormatOk returns a tuple with the RequestBodyFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhookRequestPatch) GetRequestBodyFormatOk() (*int64, bool) {
	if o == nil || IsNil(o.RequestBodyFormat) {
		return nil, false
	}
	return o.RequestBodyFormat, true
}

// HasRequestBodyFormat returns a boolean if a field has been set.
func (o *RepositoryWebhookRequestPatch) HasRequestBodyFormat() bool {
	if o != nil && !IsNil(o.RequestBodyFormat) {
		return true
	}

	return false
}

// SetRequestBodyFormat gets a reference to the given int64 and assigns it to the RequestBodyFormat field.
func (o *RepositoryWebhookRequestPatch) SetRequestBodyFormat(v int64) {
	o.RequestBodyFormat = &v
}

// GetRequestBodyTemplateFormat returns the RequestBodyTemplateFormat field value if set, zero value otherwise.
func (o *RepositoryWebhookRequestPatch) GetRequestBodyTemplateFormat() int64 {
	if o == nil || IsNil(o.RequestBodyTemplateFormat) {
		var ret int64
		return ret
	}
	return *o.RequestBodyTemplateFormat
}

// GetRequestBodyTemplateFormatOk returns a tuple with the RequestBodyTemplateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhookRequestPatch) GetRequestBodyTemplateFormatOk() (*int64, bool) {
	if o == nil || IsNil(o.RequestBodyTemplateFormat) {
		return nil, false
	}
	return o.RequestBodyTemplateFormat, true
}

// HasRequestBodyTemplateFormat returns a boolean if a field has been set.
func (o *RepositoryWebhookRequestPatch) HasRequestBodyTemplateFormat() bool {
	if o != nil && !IsNil(o.RequestBodyTemplateFormat) {
		return true
	}

	return false
}

// SetRequestBodyTemplateFormat gets a reference to the given int64 and assigns it to the RequestBodyTemplateFormat field.
func (o *RepositoryWebhookRequestPatch) SetRequestBodyTemplateFormat(v int64) {
	o.RequestBodyTemplateFormat = &v
}

// GetRequestContentType returns the RequestContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryWebhookRequestPatch) GetRequestContentType() string {
	if o == nil || IsNil(o.RequestContentType.Get()) {
		var ret string
		return ret
	}
	return *o.RequestContentType.Get()
}

// GetRequestContentTypeOk returns a tuple with the RequestContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhookRequestPatch) GetRequestContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestContentType.Get(), o.RequestContentType.IsSet()
}

// HasRequestContentType returns a boolean if a field has been set.
func (o *RepositoryWebhookRequestPatch) HasRequestContentType() bool {
	if o != nil && o.RequestContentType.IsSet() {
		return true
	}

	return false
}

// SetRequestContentType gets a reference to the given NullableString and assigns it to the RequestContentType field.
func (o *RepositoryWebhookRequestPatch) SetRequestContentType(v string) {
	o.RequestContentType.Set(&v)
}

// SetRequestContentTypeNil sets the value for RequestContentType to be an explicit nil
func (o *RepositoryWebhookRequestPatch) SetRequestContentTypeNil() {
	o.RequestContentType.Set(nil)
}

// UnsetRequestContentType ensures that no value is present for RequestContentType, not even an explicit nil
func (o *RepositoryWebhookRequestPatch) UnsetRequestContentType() {
	o.RequestContentType.Unset()
}

// GetSecretHeader returns the SecretHeader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryWebhookRequestPatch) GetSecretHeader() string {
	if o == nil || IsNil(o.SecretHeader.Get()) {
		var ret string
		return ret
	}
	return *o.SecretHeader.Get()
}

// GetSecretHeaderOk returns a tuple with the SecretHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhookRequestPatch) GetSecretHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretHeader.Get(), o.SecretHeader.IsSet()
}

// HasSecretHeader returns a boolean if a field has been set.
func (o *RepositoryWebhookRequestPatch) HasSecretHeader() bool {
	if o != nil && o.SecretHeader.IsSet() {
		return true
	}

	return false
}

// SetSecretHeader gets a reference to the given NullableString and assigns it to the SecretHeader field.
func (o *RepositoryWebhookRequestPatch) SetSecretHeader(v string) {
	o.SecretHeader.Set(&v)
}

// SetSecretHeaderNil sets the value for SecretHeader to be an explicit nil
func (o *RepositoryWebhookRequestPatch) SetSecretHeaderNil() {
	o.SecretHeader.Set(nil)
}

// UnsetSecretHeader ensures that no value is present for SecretHeader, not even an explicit nil
func (o *RepositoryWebhookRequestPatch) UnsetSecretHeader() {
	o.SecretHeader.Unset()
}

// GetSecretValue returns the SecretValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryWebhookRequestPatch) GetSecretValue() string {
	if o == nil || IsNil(o.SecretValue.Get()) {
		var ret string
		return ret
	}
	return *o.SecretValue.Get()
}

// GetSecretValueOk returns a tuple with the SecretValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhookRequestPatch) GetSecretValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretValue.Get(), o.SecretValue.IsSet()
}

// HasSecretValue returns a boolean if a field has been set.
func (o *RepositoryWebhookRequestPatch) HasSecretValue() bool {
	if o != nil && o.SecretValue.IsSet() {
		return true
	}

	return false
}

// SetSecretValue gets a reference to the given NullableString and assigns it to the SecretValue field.
func (o *RepositoryWebhookRequestPatch) SetSecretValue(v string) {
	o.SecretValue.Set(&v)
}

// SetSecretValueNil sets the value for SecretValue to be an explicit nil
func (o *RepositoryWebhookRequestPatch) SetSecretValueNil() {
	o.SecretValue.Set(nil)
}

// UnsetSecretValue ensures that no value is present for SecretValue, not even an explicit nil
func (o *RepositoryWebhookRequestPatch) UnsetSecretValue() {
	o.SecretValue.Unset()
}

// GetSignatureKey returns the SignatureKey field value if set, zero value otherwise.
func (o *RepositoryWebhookRequestPatch) GetSignatureKey() string {
	if o == nil || IsNil(o.SignatureKey) {
		var ret string
		return ret
	}
	return *o.SignatureKey
}

// GetSignatureKeyOk returns a tuple with the SignatureKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhookRequestPatch) GetSignatureKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureKey) {
		return nil, false
	}
	return o.SignatureKey, true
}

// HasSignatureKey returns a boolean if a field has been set.
func (o *RepositoryWebhookRequestPatch) HasSignatureKey() bool {
	if o != nil && !IsNil(o.SignatureKey) {
		return true
	}

	return false
}

// SetSignatureKey gets a reference to the given string and assigns it to the SignatureKey field.
func (o *RepositoryWebhookRequestPatch) SetSignatureKey(v string) {
	o.SignatureKey = &v
}

// GetTargetUrl returns the TargetUrl field value if set, zero value otherwise.
func (o *RepositoryWebhookRequestPatch) GetTargetUrl() string {
	if o == nil || IsNil(o.TargetUrl) {
		var ret string
		return ret
	}
	return *o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhookRequestPatch) GetTargetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUrl) {
		return nil, false
	}
	return o.TargetUrl, true
}

// HasTargetUrl returns a boolean if a field has been set.
func (o *RepositoryWebhookRequestPatch) HasTargetUrl() bool {
	if o != nil && !IsNil(o.TargetUrl) {
		return true
	}

	return false
}

// SetTargetUrl gets a reference to the given string and assigns it to the TargetUrl field.
func (o *RepositoryWebhookRequestPatch) SetTargetUrl(v string) {
	o.TargetUrl = &v
}

// GetTemplates returns the Templates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryWebhookRequestPatch) GetTemplates() []WebhookTemplate {
	if o == nil {
		var ret []WebhookTemplate
		return ret
	}
	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhookRequestPatch) GetTemplatesOk() ([]WebhookTemplate, bool) {
	if o == nil || IsNil(o.Templates) {
		return nil, false
	}
	return o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *RepositoryWebhookRequestPatch) HasTemplates() bool {
	if o != nil && !IsNil(o.Templates) {
		return true
	}

	return false
}

// SetTemplates gets a reference to the given []WebhookTemplate and assigns it to the Templates field.
func (o *RepositoryWebhookRequestPatch) SetTemplates(v []WebhookTemplate) {
	o.Templates = v
}

// GetVerifySsl returns the VerifySsl field value if set, zero value otherwise.
func (o *RepositoryWebhookRequestPatch) GetVerifySsl() bool {
	if o == nil || IsNil(o.VerifySsl) {
		var ret bool
		return ret
	}
	return *o.VerifySsl
}

// GetVerifySslOk returns a tuple with the VerifySsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhookRequestPatch) GetVerifySslOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifySsl) {
		return nil, false
	}
	return o.VerifySsl, true
}

// HasVerifySsl returns a boolean if a field has been set.
func (o *RepositoryWebhookRequestPatch) HasVerifySsl() bool {
	if o != nil && !IsNil(o.VerifySsl) {
		return true
	}

	return false
}

// SetVerifySsl gets a reference to the given bool and assigns it to the VerifySsl field.
func (o *RepositoryWebhookRequestPatch) SetVerifySsl(v bool) {
	o.VerifySsl = &v
}

func (o RepositoryWebhookRequestPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryWebhookRequestPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if o.PackageQuery.IsSet() {
		toSerialize["package_query"] = o.PackageQuery.Get()
	}
	if !IsNil(o.RequestBodyFormat) {
		toSerialize["request_body_format"] = o.RequestBodyFormat
	}
	if !IsNil(o.RequestBodyTemplateFormat) {
		toSerialize["request_body_template_format"] = o.RequestBodyTemplateFormat
	}
	if o.RequestContentType.IsSet() {
		toSerialize["request_content_type"] = o.RequestContentType.Get()
	}
	if o.SecretHeader.IsSet() {
		toSerialize["secret_header"] = o.SecretHeader.Get()
	}
	if o.SecretValue.IsSet() {
		toSerialize["secret_value"] = o.SecretValue.Get()
	}
	if !IsNil(o.SignatureKey) {
		toSerialize["signature_key"] = o.SignatureKey
	}
	if !IsNil(o.TargetUrl) {
		toSerialize["target_url"] = o.TargetUrl
	}
	if o.Templates != nil {
		toSerialize["templates"] = o.Templates
	}
	if !IsNil(o.VerifySsl) {
		toSerialize["verify_ssl"] = o.VerifySsl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RepositoryWebhookRequestPatch) UnmarshalJSON(data []byte) (err error) {
	varRepositoryWebhookRequestPatch := _RepositoryWebhookRequestPatch{}

	err = json.Unmarshal(data, &varRepositoryWebhookRequestPatch)

	if err != nil {
		return err
	}

	*o = RepositoryWebhookRequestPatch(varRepositoryWebhookRequestPatch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "events")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "package_query")
		delete(additionalProperties, "request_body_format")
		delete(additionalProperties, "request_body_template_format")
		delete(additionalProperties, "request_content_type")
		delete(additionalProperties, "secret_header")
		delete(additionalProperties, "secret_value")
		delete(additionalProperties, "signature_key")
		delete(additionalProperties, "target_url")
		delete(additionalProperties, "templates")
		delete(additionalProperties, "verify_ssl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRepositoryWebhookRequestPatch struct {
	value *RepositoryWebhookRequestPatch
	isSet bool
}

func (v NullableRepositoryWebhookRequestPatch) Get() *RepositoryWebhookRequestPatch {
	return v.value
}

func (v *NullableRepositoryWebhookRequestPatch) Set(val *RepositoryWebhookRequestPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryWebhookRequestPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryWebhookRequestPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryWebhookRequestPatch(val *RepositoryWebhookRequestPatch) *NullableRepositoryWebhookRequestPatch {
	return &NullableRepositoryWebhookRequestPatch{value: val, isSet: true}
}

func (v NullableRepositoryWebhookRequestPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryWebhookRequestPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
