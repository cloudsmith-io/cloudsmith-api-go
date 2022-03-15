# WebhooksPartialUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to **[]string** | None | [optional] 
**IsActive** | Pointer to **bool** | If enabled, the webhook will trigger on events and send payloads to the configured target URL. | [optional] 
**PackageQuery** | Pointer to **string** | The package-based search query for webhooks to fire. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. If a package does not match, the webhook will not fire. | [optional] 
**RequestBodyFormat** | Pointer to **string** | The format of the payloads for webhook requests. | [optional] 
**RequestBodyTemplateFormat** | Pointer to **string** | The format of the payloads for webhook requests. | [optional] 
**RequestContentType** | Pointer to **string** | The value that will be sent for the &#39;Content Type&#39; header.  | [optional] 
**SecretHeader** | Pointer to **string** | The header to send the predefined secret in. This must be unique from existing headers or it won&#39;t be sent. You can use this as a form of authentication on the endpoint side. | [optional] 
**SecretValue** | Pointer to **string** | The value for the predefined secret (note: this is treated as a passphrase and is encrypted when we store it). You can use this as a form of authentication on the endpoint side. | [optional] 
**SignatureKey** | Pointer to **string** | The value for the signature key - This is used to generate an HMAC-based hex digest of the request body, which we send as the X-Cloudsmith-Signature header so that you can ensure that the request wasn&#39;t modified by a malicious party (note: this is treated as a passphrase and is encrypted when we store it). | [optional] 
**TargetUrl** | Pointer to **string** | The destination URL that webhook payloads will be POST&#39;ed to. | [optional] 
**Templates** | Pointer to [**[]WebhooksOwnerRepoTemplates**](WebhooksOwnerRepoTemplates.md) | None | [optional] 
**VerifySsl** | Pointer to **bool** | If enabled, SSL certificates is verified when webhooks are sent. It&#39;s recommended to leave this enabled as not verifying the integrity of SSL certificates leaves you susceptible to Man-in-the-Middle (MITM) attacks. | [optional] 

## Methods

### NewWebhooksPartialUpdate

`func NewWebhooksPartialUpdate() *WebhooksPartialUpdate`

NewWebhooksPartialUpdate instantiates a new WebhooksPartialUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksPartialUpdateWithDefaults

`func NewWebhooksPartialUpdateWithDefaults() *WebhooksPartialUpdate`

NewWebhooksPartialUpdateWithDefaults instantiates a new WebhooksPartialUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhooksPartialUpdate) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhooksPartialUpdate) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhooksPartialUpdate) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WebhooksPartialUpdate) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetIsActive

`func (o *WebhooksPartialUpdate) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *WebhooksPartialUpdate) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *WebhooksPartialUpdate) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *WebhooksPartialUpdate) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetPackageQuery

`func (o *WebhooksPartialUpdate) GetPackageQuery() string`

GetPackageQuery returns the PackageQuery field if non-nil, zero value otherwise.

### GetPackageQueryOk

`func (o *WebhooksPartialUpdate) GetPackageQueryOk() (*string, bool)`

GetPackageQueryOk returns a tuple with the PackageQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageQuery

`func (o *WebhooksPartialUpdate) SetPackageQuery(v string)`

SetPackageQuery sets PackageQuery field to given value.

### HasPackageQuery

`func (o *WebhooksPartialUpdate) HasPackageQuery() bool`

HasPackageQuery returns a boolean if a field has been set.

### GetRequestBodyFormat

`func (o *WebhooksPartialUpdate) GetRequestBodyFormat() string`

GetRequestBodyFormat returns the RequestBodyFormat field if non-nil, zero value otherwise.

### GetRequestBodyFormatOk

`func (o *WebhooksPartialUpdate) GetRequestBodyFormatOk() (*string, bool)`

GetRequestBodyFormatOk returns a tuple with the RequestBodyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyFormat

`func (o *WebhooksPartialUpdate) SetRequestBodyFormat(v string)`

SetRequestBodyFormat sets RequestBodyFormat field to given value.

### HasRequestBodyFormat

`func (o *WebhooksPartialUpdate) HasRequestBodyFormat() bool`

HasRequestBodyFormat returns a boolean if a field has been set.

### GetRequestBodyTemplateFormat

`func (o *WebhooksPartialUpdate) GetRequestBodyTemplateFormat() string`

GetRequestBodyTemplateFormat returns the RequestBodyTemplateFormat field if non-nil, zero value otherwise.

### GetRequestBodyTemplateFormatOk

`func (o *WebhooksPartialUpdate) GetRequestBodyTemplateFormatOk() (*string, bool)`

GetRequestBodyTemplateFormatOk returns a tuple with the RequestBodyTemplateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyTemplateFormat

`func (o *WebhooksPartialUpdate) SetRequestBodyTemplateFormat(v string)`

SetRequestBodyTemplateFormat sets RequestBodyTemplateFormat field to given value.

### HasRequestBodyTemplateFormat

`func (o *WebhooksPartialUpdate) HasRequestBodyTemplateFormat() bool`

HasRequestBodyTemplateFormat returns a boolean if a field has been set.

### GetRequestContentType

`func (o *WebhooksPartialUpdate) GetRequestContentType() string`

GetRequestContentType returns the RequestContentType field if non-nil, zero value otherwise.

### GetRequestContentTypeOk

`func (o *WebhooksPartialUpdate) GetRequestContentTypeOk() (*string, bool)`

GetRequestContentTypeOk returns a tuple with the RequestContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContentType

`func (o *WebhooksPartialUpdate) SetRequestContentType(v string)`

SetRequestContentType sets RequestContentType field to given value.

### HasRequestContentType

`func (o *WebhooksPartialUpdate) HasRequestContentType() bool`

HasRequestContentType returns a boolean if a field has been set.

### GetSecretHeader

`func (o *WebhooksPartialUpdate) GetSecretHeader() string`

GetSecretHeader returns the SecretHeader field if non-nil, zero value otherwise.

### GetSecretHeaderOk

`func (o *WebhooksPartialUpdate) GetSecretHeaderOk() (*string, bool)`

GetSecretHeaderOk returns a tuple with the SecretHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretHeader

`func (o *WebhooksPartialUpdate) SetSecretHeader(v string)`

SetSecretHeader sets SecretHeader field to given value.

### HasSecretHeader

`func (o *WebhooksPartialUpdate) HasSecretHeader() bool`

HasSecretHeader returns a boolean if a field has been set.

### GetSecretValue

`func (o *WebhooksPartialUpdate) GetSecretValue() string`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *WebhooksPartialUpdate) GetSecretValueOk() (*string, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *WebhooksPartialUpdate) SetSecretValue(v string)`

SetSecretValue sets SecretValue field to given value.

### HasSecretValue

`func (o *WebhooksPartialUpdate) HasSecretValue() bool`

HasSecretValue returns a boolean if a field has been set.

### GetSignatureKey

`func (o *WebhooksPartialUpdate) GetSignatureKey() string`

GetSignatureKey returns the SignatureKey field if non-nil, zero value otherwise.

### GetSignatureKeyOk

`func (o *WebhooksPartialUpdate) GetSignatureKeyOk() (*string, bool)`

GetSignatureKeyOk returns a tuple with the SignatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureKey

`func (o *WebhooksPartialUpdate) SetSignatureKey(v string)`

SetSignatureKey sets SignatureKey field to given value.

### HasSignatureKey

`func (o *WebhooksPartialUpdate) HasSignatureKey() bool`

HasSignatureKey returns a boolean if a field has been set.

### GetTargetUrl

`func (o *WebhooksPartialUpdate) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *WebhooksPartialUpdate) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *WebhooksPartialUpdate) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *WebhooksPartialUpdate) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetTemplates

`func (o *WebhooksPartialUpdate) GetTemplates() []WebhooksOwnerRepoTemplates`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *WebhooksPartialUpdate) GetTemplatesOk() (*[]WebhooksOwnerRepoTemplates, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *WebhooksPartialUpdate) SetTemplates(v []WebhooksOwnerRepoTemplates)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *WebhooksPartialUpdate) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetVerifySsl

`func (o *WebhooksPartialUpdate) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *WebhooksPartialUpdate) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *WebhooksPartialUpdate) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *WebhooksPartialUpdate) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


