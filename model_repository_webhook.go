/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.498.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"time"
)

// checks if the RepositoryWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryWebhook{}

// RepositoryWebhook struct for RepositoryWebhook
type RepositoryWebhook struct {
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	CreatedBy        *string    `json:"created_by,omitempty"`
	CreatedByUrl     *string    `json:"created_by_url,omitempty"`
	DisableReason    *int64     `json:"disable_reason,omitempty"`
	DisableReasonStr *string    `json:"disable_reason_str,omitempty"`
	Events           []string   `json:"events"`
	// Deprecated (23-05-15): Please use 'slug_perm' instead. Previously: A monotonically increasing number that identified a webhook request within a repository.
	Identifier NullableInt64 `json:"identifier,omitempty"`
	// If enabled, the webhook will trigger on subscribed events and send payloads to the configured target URL.
	IsActive              *bool   `json:"is_active,omitempty"`
	IsLastResponseBad     *bool   `json:"is_last_response_bad,omitempty"`
	LastResponseStatus    *int64  `json:"last_response_status,omitempty"`
	LastResponseStatusStr *string `json:"last_response_status_str,omitempty"`
	NumSent               *int64  `json:"num_sent,omitempty"`
	// The package-based search query for webhooks to fire. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. If a package does not match, the webhook will not fire.
	PackageQuery NullableString `json:"package_query,omitempty"`
	// The format of the payloads for webhook requests.
	RequestBodyFormat    *int64  `json:"request_body_format,omitempty"`
	RequestBodyFormatStr *string `json:"request_body_format_str,omitempty"`
	// The format of the payloads for webhook requests.
	RequestBodyTemplateFormat    *int64  `json:"request_body_template_format,omitempty"`
	RequestBodyTemplateFormatStr *string `json:"request_body_template_format_str,omitempty"`
	// The value that will be sent for the 'Content Type' header.
	RequestContentType NullableString `json:"request_content_type,omitempty"`
	// The header to send the predefined secret in. This must be unique from existing headers or it won't be sent. You can use this as a form of authentication on the endpoint side.
	SecretHeader NullableString `json:"secret_header,omitempty"`
	SelfUrl      *string        `json:"self_url,omitempty"`
	SlugPerm     *string        `json:"slug_perm,omitempty"`
	// The destination URL that webhook payloads will be POST'ed to.
	TargetUrl    string            `json:"target_url"`
	Templates    []WebhookTemplate `json:"templates"`
	UpdatedAt    *time.Time        `json:"updated_at,omitempty"`
	UpdatedBy    *string           `json:"updated_by,omitempty"`
	UpdatedByUrl *string           `json:"updated_by_url,omitempty"`
	// If enabled, SSL certificates is verified when webhooks are sent. It's recommended to leave this enabled as not verifying the integrity of SSL certificates leaves you susceptible to Man-in-the-Middle (MITM) attacks.
	VerifySsl *bool `json:"verify_ssl,omitempty"`
}

// NewRepositoryWebhook instantiates a new RepositoryWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryWebhook(events []string, targetUrl string, templates []WebhookTemplate) *RepositoryWebhook {
	this := RepositoryWebhook{}
	this.Events = events
	this.TargetUrl = targetUrl
	this.Templates = templates
	return &this
}

// NewRepositoryWebhookWithDefaults instantiates a new RepositoryWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryWebhookWithDefaults() *RepositoryWebhook {
	this := RepositoryWebhook{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RepositoryWebhook) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *RepositoryWebhook) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedByUrl returns the CreatedByUrl field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetCreatedByUrl() string {
	if o == nil || IsNil(o.CreatedByUrl) {
		var ret string
		return ret
	}
	return *o.CreatedByUrl
}

// GetCreatedByUrlOk returns a tuple with the CreatedByUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetCreatedByUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedByUrl) {
		return nil, false
	}
	return o.CreatedByUrl, true
}

// HasCreatedByUrl returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasCreatedByUrl() bool {
	if o != nil && !IsNil(o.CreatedByUrl) {
		return true
	}

	return false
}

// SetCreatedByUrl gets a reference to the given string and assigns it to the CreatedByUrl field.
func (o *RepositoryWebhook) SetCreatedByUrl(v string) {
	o.CreatedByUrl = &v
}

// GetDisableReason returns the DisableReason field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetDisableReason() int64 {
	if o == nil || IsNil(o.DisableReason) {
		var ret int64
		return ret
	}
	return *o.DisableReason
}

// GetDisableReasonOk returns a tuple with the DisableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetDisableReasonOk() (*int64, bool) {
	if o == nil || IsNil(o.DisableReason) {
		return nil, false
	}
	return o.DisableReason, true
}

// HasDisableReason returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasDisableReason() bool {
	if o != nil && !IsNil(o.DisableReason) {
		return true
	}

	return false
}

// SetDisableReason gets a reference to the given int64 and assigns it to the DisableReason field.
func (o *RepositoryWebhook) SetDisableReason(v int64) {
	o.DisableReason = &v
}

// GetDisableReasonStr returns the DisableReasonStr field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetDisableReasonStr() string {
	if o == nil || IsNil(o.DisableReasonStr) {
		var ret string
		return ret
	}
	return *o.DisableReasonStr
}

// GetDisableReasonStrOk returns a tuple with the DisableReasonStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetDisableReasonStrOk() (*string, bool) {
	if o == nil || IsNil(o.DisableReasonStr) {
		return nil, false
	}
	return o.DisableReasonStr, true
}

// HasDisableReasonStr returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasDisableReasonStr() bool {
	if o != nil && !IsNil(o.DisableReasonStr) {
		return true
	}

	return false
}

// SetDisableReasonStr gets a reference to the given string and assigns it to the DisableReasonStr field.
func (o *RepositoryWebhook) SetDisableReasonStr(v string) {
	o.DisableReasonStr = &v
}

// GetEvents returns the Events field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *RepositoryWebhook) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhook) GetEventsOk() ([]string, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *RepositoryWebhook) SetEvents(v []string) {
	o.Events = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryWebhook) GetIdentifier() int64 {
	if o == nil || IsNil(o.Identifier.Get()) {
		var ret int64
		return ret
	}
	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhook) GetIdentifierOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// HasIdentifier returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasIdentifier() bool {
	if o != nil && o.Identifier.IsSet() {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given NullableInt64 and assigns it to the Identifier field.
func (o *RepositoryWebhook) SetIdentifier(v int64) {
	o.Identifier.Set(&v)
}

// SetIdentifierNil sets the value for Identifier to be an explicit nil
func (o *RepositoryWebhook) SetIdentifierNil() {
	o.Identifier.Set(nil)
}

// UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
func (o *RepositoryWebhook) UnsetIdentifier() {
	o.Identifier.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *RepositoryWebhook) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsLastResponseBad returns the IsLastResponseBad field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetIsLastResponseBad() bool {
	if o == nil || IsNil(o.IsLastResponseBad) {
		var ret bool
		return ret
	}
	return *o.IsLastResponseBad
}

// GetIsLastResponseBadOk returns a tuple with the IsLastResponseBad field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetIsLastResponseBadOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLastResponseBad) {
		return nil, false
	}
	return o.IsLastResponseBad, true
}

// HasIsLastResponseBad returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasIsLastResponseBad() bool {
	if o != nil && !IsNil(o.IsLastResponseBad) {
		return true
	}

	return false
}

// SetIsLastResponseBad gets a reference to the given bool and assigns it to the IsLastResponseBad field.
func (o *RepositoryWebhook) SetIsLastResponseBad(v bool) {
	o.IsLastResponseBad = &v
}

// GetLastResponseStatus returns the LastResponseStatus field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetLastResponseStatus() int64 {
	if o == nil || IsNil(o.LastResponseStatus) {
		var ret int64
		return ret
	}
	return *o.LastResponseStatus
}

// GetLastResponseStatusOk returns a tuple with the LastResponseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetLastResponseStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.LastResponseStatus) {
		return nil, false
	}
	return o.LastResponseStatus, true
}

// HasLastResponseStatus returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasLastResponseStatus() bool {
	if o != nil && !IsNil(o.LastResponseStatus) {
		return true
	}

	return false
}

// SetLastResponseStatus gets a reference to the given int64 and assigns it to the LastResponseStatus field.
func (o *RepositoryWebhook) SetLastResponseStatus(v int64) {
	o.LastResponseStatus = &v
}

// GetLastResponseStatusStr returns the LastResponseStatusStr field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetLastResponseStatusStr() string {
	if o == nil || IsNil(o.LastResponseStatusStr) {
		var ret string
		return ret
	}
	return *o.LastResponseStatusStr
}

// GetLastResponseStatusStrOk returns a tuple with the LastResponseStatusStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetLastResponseStatusStrOk() (*string, bool) {
	if o == nil || IsNil(o.LastResponseStatusStr) {
		return nil, false
	}
	return o.LastResponseStatusStr, true
}

// HasLastResponseStatusStr returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasLastResponseStatusStr() bool {
	if o != nil && !IsNil(o.LastResponseStatusStr) {
		return true
	}

	return false
}

// SetLastResponseStatusStr gets a reference to the given string and assigns it to the LastResponseStatusStr field.
func (o *RepositoryWebhook) SetLastResponseStatusStr(v string) {
	o.LastResponseStatusStr = &v
}

// GetNumSent returns the NumSent field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetNumSent() int64 {
	if o == nil || IsNil(o.NumSent) {
		var ret int64
		return ret
	}
	return *o.NumSent
}

// GetNumSentOk returns a tuple with the NumSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetNumSentOk() (*int64, bool) {
	if o == nil || IsNil(o.NumSent) {
		return nil, false
	}
	return o.NumSent, true
}

// HasNumSent returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasNumSent() bool {
	if o != nil && !IsNil(o.NumSent) {
		return true
	}

	return false
}

// SetNumSent gets a reference to the given int64 and assigns it to the NumSent field.
func (o *RepositoryWebhook) SetNumSent(v int64) {
	o.NumSent = &v
}

// GetPackageQuery returns the PackageQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryWebhook) GetPackageQuery() string {
	if o == nil || IsNil(o.PackageQuery.Get()) {
		var ret string
		return ret
	}
	return *o.PackageQuery.Get()
}

// GetPackageQueryOk returns a tuple with the PackageQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhook) GetPackageQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageQuery.Get(), o.PackageQuery.IsSet()
}

// HasPackageQuery returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasPackageQuery() bool {
	if o != nil && o.PackageQuery.IsSet() {
		return true
	}

	return false
}

// SetPackageQuery gets a reference to the given NullableString and assigns it to the PackageQuery field.
func (o *RepositoryWebhook) SetPackageQuery(v string) {
	o.PackageQuery.Set(&v)
}

// SetPackageQueryNil sets the value for PackageQuery to be an explicit nil
func (o *RepositoryWebhook) SetPackageQueryNil() {
	o.PackageQuery.Set(nil)
}

// UnsetPackageQuery ensures that no value is present for PackageQuery, not even an explicit nil
func (o *RepositoryWebhook) UnsetPackageQuery() {
	o.PackageQuery.Unset()
}

// GetRequestBodyFormat returns the RequestBodyFormat field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetRequestBodyFormat() int64 {
	if o == nil || IsNil(o.RequestBodyFormat) {
		var ret int64
		return ret
	}
	return *o.RequestBodyFormat
}

// GetRequestBodyFormatOk returns a tuple with the RequestBodyFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetRequestBodyFormatOk() (*int64, bool) {
	if o == nil || IsNil(o.RequestBodyFormat) {
		return nil, false
	}
	return o.RequestBodyFormat, true
}

// HasRequestBodyFormat returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasRequestBodyFormat() bool {
	if o != nil && !IsNil(o.RequestBodyFormat) {
		return true
	}

	return false
}

// SetRequestBodyFormat gets a reference to the given int64 and assigns it to the RequestBodyFormat field.
func (o *RepositoryWebhook) SetRequestBodyFormat(v int64) {
	o.RequestBodyFormat = &v
}

// GetRequestBodyFormatStr returns the RequestBodyFormatStr field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetRequestBodyFormatStr() string {
	if o == nil || IsNil(o.RequestBodyFormatStr) {
		var ret string
		return ret
	}
	return *o.RequestBodyFormatStr
}

// GetRequestBodyFormatStrOk returns a tuple with the RequestBodyFormatStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetRequestBodyFormatStrOk() (*string, bool) {
	if o == nil || IsNil(o.RequestBodyFormatStr) {
		return nil, false
	}
	return o.RequestBodyFormatStr, true
}

// HasRequestBodyFormatStr returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasRequestBodyFormatStr() bool {
	if o != nil && !IsNil(o.RequestBodyFormatStr) {
		return true
	}

	return false
}

// SetRequestBodyFormatStr gets a reference to the given string and assigns it to the RequestBodyFormatStr field.
func (o *RepositoryWebhook) SetRequestBodyFormatStr(v string) {
	o.RequestBodyFormatStr = &v
}

// GetRequestBodyTemplateFormat returns the RequestBodyTemplateFormat field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetRequestBodyTemplateFormat() int64 {
	if o == nil || IsNil(o.RequestBodyTemplateFormat) {
		var ret int64
		return ret
	}
	return *o.RequestBodyTemplateFormat
}

// GetRequestBodyTemplateFormatOk returns a tuple with the RequestBodyTemplateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetRequestBodyTemplateFormatOk() (*int64, bool) {
	if o == nil || IsNil(o.RequestBodyTemplateFormat) {
		return nil, false
	}
	return o.RequestBodyTemplateFormat, true
}

// HasRequestBodyTemplateFormat returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasRequestBodyTemplateFormat() bool {
	if o != nil && !IsNil(o.RequestBodyTemplateFormat) {
		return true
	}

	return false
}

// SetRequestBodyTemplateFormat gets a reference to the given int64 and assigns it to the RequestBodyTemplateFormat field.
func (o *RepositoryWebhook) SetRequestBodyTemplateFormat(v int64) {
	o.RequestBodyTemplateFormat = &v
}

// GetRequestBodyTemplateFormatStr returns the RequestBodyTemplateFormatStr field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetRequestBodyTemplateFormatStr() string {
	if o == nil || IsNil(o.RequestBodyTemplateFormatStr) {
		var ret string
		return ret
	}
	return *o.RequestBodyTemplateFormatStr
}

// GetRequestBodyTemplateFormatStrOk returns a tuple with the RequestBodyTemplateFormatStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetRequestBodyTemplateFormatStrOk() (*string, bool) {
	if o == nil || IsNil(o.RequestBodyTemplateFormatStr) {
		return nil, false
	}
	return o.RequestBodyTemplateFormatStr, true
}

// HasRequestBodyTemplateFormatStr returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasRequestBodyTemplateFormatStr() bool {
	if o != nil && !IsNil(o.RequestBodyTemplateFormatStr) {
		return true
	}

	return false
}

// SetRequestBodyTemplateFormatStr gets a reference to the given string and assigns it to the RequestBodyTemplateFormatStr field.
func (o *RepositoryWebhook) SetRequestBodyTemplateFormatStr(v string) {
	o.RequestBodyTemplateFormatStr = &v
}

// GetRequestContentType returns the RequestContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryWebhook) GetRequestContentType() string {
	if o == nil || IsNil(o.RequestContentType.Get()) {
		var ret string
		return ret
	}
	return *o.RequestContentType.Get()
}

// GetRequestContentTypeOk returns a tuple with the RequestContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhook) GetRequestContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestContentType.Get(), o.RequestContentType.IsSet()
}

// HasRequestContentType returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasRequestContentType() bool {
	if o != nil && o.RequestContentType.IsSet() {
		return true
	}

	return false
}

// SetRequestContentType gets a reference to the given NullableString and assigns it to the RequestContentType field.
func (o *RepositoryWebhook) SetRequestContentType(v string) {
	o.RequestContentType.Set(&v)
}

// SetRequestContentTypeNil sets the value for RequestContentType to be an explicit nil
func (o *RepositoryWebhook) SetRequestContentTypeNil() {
	o.RequestContentType.Set(nil)
}

// UnsetRequestContentType ensures that no value is present for RequestContentType, not even an explicit nil
func (o *RepositoryWebhook) UnsetRequestContentType() {
	o.RequestContentType.Unset()
}

// GetSecretHeader returns the SecretHeader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryWebhook) GetSecretHeader() string {
	if o == nil || IsNil(o.SecretHeader.Get()) {
		var ret string
		return ret
	}
	return *o.SecretHeader.Get()
}

// GetSecretHeaderOk returns a tuple with the SecretHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhook) GetSecretHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretHeader.Get(), o.SecretHeader.IsSet()
}

// HasSecretHeader returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasSecretHeader() bool {
	if o != nil && o.SecretHeader.IsSet() {
		return true
	}

	return false
}

// SetSecretHeader gets a reference to the given NullableString and assigns it to the SecretHeader field.
func (o *RepositoryWebhook) SetSecretHeader(v string) {
	o.SecretHeader.Set(&v)
}

// SetSecretHeaderNil sets the value for SecretHeader to be an explicit nil
func (o *RepositoryWebhook) SetSecretHeaderNil() {
	o.SecretHeader.Set(nil)
}

// UnsetSecretHeader ensures that no value is present for SecretHeader, not even an explicit nil
func (o *RepositoryWebhook) UnsetSecretHeader() {
	o.SecretHeader.Unset()
}

// GetSelfUrl returns the SelfUrl field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetSelfUrl() string {
	if o == nil || IsNil(o.SelfUrl) {
		var ret string
		return ret
	}
	return *o.SelfUrl
}

// GetSelfUrlOk returns a tuple with the SelfUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetSelfUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SelfUrl) {
		return nil, false
	}
	return o.SelfUrl, true
}

// HasSelfUrl returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasSelfUrl() bool {
	if o != nil && !IsNil(o.SelfUrl) {
		return true
	}

	return false
}

// SetSelfUrl gets a reference to the given string and assigns it to the SelfUrl field.
func (o *RepositoryWebhook) SetSelfUrl(v string) {
	o.SelfUrl = &v
}

// GetSlugPerm returns the SlugPerm field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetSlugPerm() string {
	if o == nil || IsNil(o.SlugPerm) {
		var ret string
		return ret
	}
	return *o.SlugPerm
}

// GetSlugPermOk returns a tuple with the SlugPerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetSlugPermOk() (*string, bool) {
	if o == nil || IsNil(o.SlugPerm) {
		return nil, false
	}
	return o.SlugPerm, true
}

// HasSlugPerm returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasSlugPerm() bool {
	if o != nil && !IsNil(o.SlugPerm) {
		return true
	}

	return false
}

// SetSlugPerm gets a reference to the given string and assigns it to the SlugPerm field.
func (o *RepositoryWebhook) SetSlugPerm(v string) {
	o.SlugPerm = &v
}

// GetTargetUrl returns the TargetUrl field value
func (o *RepositoryWebhook) GetTargetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetTargetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetUrl, true
}

// SetTargetUrl sets field value
func (o *RepositoryWebhook) SetTargetUrl(v string) {
	o.TargetUrl = v
}

// GetTemplates returns the Templates field value
// If the value is explicit nil, the zero value for []WebhookTemplate will be returned
func (o *RepositoryWebhook) GetTemplates() []WebhookTemplate {
	if o == nil {
		var ret []WebhookTemplate
		return ret
	}

	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhook) GetTemplatesOk() ([]WebhookTemplate, bool) {
	if o == nil || IsNil(o.Templates) {
		return nil, false
	}
	return o.Templates, true
}

// SetTemplates sets field value
func (o *RepositoryWebhook) SetTemplates(v []WebhookTemplate) {
	o.Templates = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RepositoryWebhook) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *RepositoryWebhook) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetUpdatedByUrl returns the UpdatedByUrl field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetUpdatedByUrl() string {
	if o == nil || IsNil(o.UpdatedByUrl) {
		var ret string
		return ret
	}
	return *o.UpdatedByUrl
}

// GetUpdatedByUrlOk returns a tuple with the UpdatedByUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetUpdatedByUrlOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedByUrl) {
		return nil, false
	}
	return o.UpdatedByUrl, true
}

// HasUpdatedByUrl returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasUpdatedByUrl() bool {
	if o != nil && !IsNil(o.UpdatedByUrl) {
		return true
	}

	return false
}

// SetUpdatedByUrl gets a reference to the given string and assigns it to the UpdatedByUrl field.
func (o *RepositoryWebhook) SetUpdatedByUrl(v string) {
	o.UpdatedByUrl = &v
}

// GetVerifySsl returns the VerifySsl field value if set, zero value otherwise.
func (o *RepositoryWebhook) GetVerifySsl() bool {
	if o == nil || IsNil(o.VerifySsl) {
		var ret bool
		return ret
	}
	return *o.VerifySsl
}

// GetVerifySslOk returns a tuple with the VerifySsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhook) GetVerifySslOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifySsl) {
		return nil, false
	}
	return o.VerifySsl, true
}

// HasVerifySsl returns a boolean if a field has been set.
func (o *RepositoryWebhook) HasVerifySsl() bool {
	if o != nil && !IsNil(o.VerifySsl) {
		return true
	}

	return false
}

// SetVerifySsl gets a reference to the given bool and assigns it to the VerifySsl field.
func (o *RepositoryWebhook) SetVerifySsl(v bool) {
	o.VerifySsl = &v
}

func (o RepositoryWebhook) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.CreatedByUrl) {
		toSerialize["created_by_url"] = o.CreatedByUrl
	}
	if !IsNil(o.DisableReason) {
		toSerialize["disable_reason"] = o.DisableReason
	}
	if !IsNil(o.DisableReasonStr) {
		toSerialize["disable_reason_str"] = o.DisableReasonStr
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.IsLastResponseBad) {
		toSerialize["is_last_response_bad"] = o.IsLastResponseBad
	}
	if !IsNil(o.LastResponseStatus) {
		toSerialize["last_response_status"] = o.LastResponseStatus
	}
	if !IsNil(o.LastResponseStatusStr) {
		toSerialize["last_response_status_str"] = o.LastResponseStatusStr
	}
	if !IsNil(o.NumSent) {
		toSerialize["num_sent"] = o.NumSent
	}
	if o.PackageQuery.IsSet() {
		toSerialize["package_query"] = o.PackageQuery.Get()
	}
	if !IsNil(o.RequestBodyFormat) {
		toSerialize["request_body_format"] = o.RequestBodyFormat
	}
	if !IsNil(o.RequestBodyFormatStr) {
		toSerialize["request_body_format_str"] = o.RequestBodyFormatStr
	}
	if !IsNil(o.RequestBodyTemplateFormat) {
		toSerialize["request_body_template_format"] = o.RequestBodyTemplateFormat
	}
	if !IsNil(o.RequestBodyTemplateFormatStr) {
		toSerialize["request_body_template_format_str"] = o.RequestBodyTemplateFormatStr
	}
	if o.RequestContentType.IsSet() {
		toSerialize["request_content_type"] = o.RequestContentType.Get()
	}
	if o.SecretHeader.IsSet() {
		toSerialize["secret_header"] = o.SecretHeader.Get()
	}
	if !IsNil(o.SelfUrl) {
		toSerialize["self_url"] = o.SelfUrl
	}
	if !IsNil(o.SlugPerm) {
		toSerialize["slug_perm"] = o.SlugPerm
	}
	toSerialize["target_url"] = o.TargetUrl
	if o.Templates != nil {
		toSerialize["templates"] = o.Templates
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	if !IsNil(o.UpdatedByUrl) {
		toSerialize["updated_by_url"] = o.UpdatedByUrl
	}
	if !IsNil(o.VerifySsl) {
		toSerialize["verify_ssl"] = o.VerifySsl
	}
	return toSerialize, nil
}

type NullableRepositoryWebhook struct {
	value *RepositoryWebhook
	isSet bool
}

func (v NullableRepositoryWebhook) Get() *RepositoryWebhook {
	return v.value
}

func (v *NullableRepositoryWebhook) Set(val *RepositoryWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryWebhook(val *RepositoryWebhook) *NullableRepositoryWebhook {
	return &NullableRepositoryWebhook{value: val, isSet: true}
}

func (v NullableRepositoryWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
