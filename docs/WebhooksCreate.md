# WebhooksCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | **[]string** | None | 
**IsActive** | Pointer to **bool** | If enabled, the webhook will trigger on events and send payloads to the configured target URL. | [optional] 
**PackageQuery** | Pointer to **string** | The package-based search query for webhooks to fire. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. If a package does not match, the webhook will not fire. | [optional] 
**RequestBodyFormat** | Pointer to **string** | The format of the payloads for webhook requests. | [optional] 
**RequestBodyTemplateFormat** | Pointer to **string** | The format of the payloads for webhook requests. | [optional] 
**RequestContentType** | Pointer to **string** | The value that will be sent for the &#39;Content Type&#39; header.  | [optional] 
**SecretHeader** | Pointer to **string** | The header to send the predefined secret in. This must be unique from existing headers or it won&#39;t be sent. You can use this as a form of authentication on the endpoint side. | [optional] 
**SecretValue** | Pointer to **string** | The value for the predefined secret (note: this is treated as a passphrase and is encrypted when we store it). You can use this as a form of authentication on the endpoint side. | [optional] 
**SignatureKey** | Pointer to **string** | The value for the signature key - This is used to generate an HMAC-based hex digest of the request body, which we send as the X-Cloudsmith-Signature header so that you can ensure that the request wasn&#39;t modified by a malicious party (note: this is treated as a passphrase and is encrypted when we store it). | [optional] 
**TargetUrl** | **string** | The destination URL that webhook payloads will be POST&#39;ed to. | 
**Templates** | [**[]WebhooksOwnerRepoTemplates**](WebhooksOwnerRepoTemplates.md) | None | 
**VerifySsl** | Pointer to **bool** | If enabled, SSL certificates is verified when webhooks are sent. It&#39;s recommended to leave this enabled as not verifying the integrity of SSL certificates leaves you susceptible to Man-in-the-Middle (MITM) attacks. | [optional] 

## Methods

### NewWebhooksCreate

`func NewWebhooksCreate(events []string, targetUrl string, templates []WebhooksOwnerRepoTemplates, ) *WebhooksCreate`

NewWebhooksCreate instantiates a new WebhooksCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksCreateWithDefaults

`func NewWebhooksCreateWithDefaults() *WebhooksCreate`

NewWebhooksCreateWithDefaults instantiates a new WebhooksCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *WebhooksCreate) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhooksCreate) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhooksCreate) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetIsActive

`func (o *WebhooksCreate) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *WebhooksCreate) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *WebhooksCreate) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *WebhooksCreate) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetPackageQuery

`func (o *WebhooksCreate) GetPackageQuery() string`

GetPackageQuery returns the PackageQuery field if non-nil, zero value otherwise.

### GetPackageQueryOk

`func (o *WebhooksCreate) GetPackageQueryOk() (*string, bool)`

GetPackageQueryOk returns a tuple with the PackageQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageQuery

`func (o *WebhooksCreate) SetPackageQuery(v string)`

SetPackageQuery sets PackageQuery field to given value.

### HasPackageQuery

`func (o *WebhooksCreate) HasPackageQuery() bool`

HasPackageQuery returns a boolean if a field has been set.

### GetRequestBodyFormat

`func (o *WebhooksCreate) GetRequestBodyFormat() string`

GetRequestBodyFormat returns the RequestBodyFormat field if non-nil, zero value otherwise.

### GetRequestBodyFormatOk

`func (o *WebhooksCreate) GetRequestBodyFormatOk() (*string, bool)`

GetRequestBodyFormatOk returns a tuple with the RequestBodyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyFormat

`func (o *WebhooksCreate) SetRequestBodyFormat(v string)`

SetRequestBodyFormat sets RequestBodyFormat field to given value.

### HasRequestBodyFormat

`func (o *WebhooksCreate) HasRequestBodyFormat() bool`

HasRequestBodyFormat returns a boolean if a field has been set.

### GetRequestBodyTemplateFormat

`func (o *WebhooksCreate) GetRequestBodyTemplateFormat() string`

GetRequestBodyTemplateFormat returns the RequestBodyTemplateFormat field if non-nil, zero value otherwise.

### GetRequestBodyTemplateFormatOk

`func (o *WebhooksCreate) GetRequestBodyTemplateFormatOk() (*string, bool)`

GetRequestBodyTemplateFormatOk returns a tuple with the RequestBodyTemplateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyTemplateFormat

`func (o *WebhooksCreate) SetRequestBodyTemplateFormat(v string)`

SetRequestBodyTemplateFormat sets RequestBodyTemplateFormat field to given value.

### HasRequestBodyTemplateFormat

`func (o *WebhooksCreate) HasRequestBodyTemplateFormat() bool`

HasRequestBodyTemplateFormat returns a boolean if a field has been set.

### GetRequestContentType

`func (o *WebhooksCreate) GetRequestContentType() string`

GetRequestContentType returns the RequestContentType field if non-nil, zero value otherwise.

### GetRequestContentTypeOk

`func (o *WebhooksCreate) GetRequestContentTypeOk() (*string, bool)`

GetRequestContentTypeOk returns a tuple with the RequestContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContentType

`func (o *WebhooksCreate) SetRequestContentType(v string)`

SetRequestContentType sets RequestContentType field to given value.

### HasRequestContentType

`func (o *WebhooksCreate) HasRequestContentType() bool`

HasRequestContentType returns a boolean if a field has been set.

### GetSecretHeader

`func (o *WebhooksCreate) GetSecretHeader() string`

GetSecretHeader returns the SecretHeader field if non-nil, zero value otherwise.

### GetSecretHeaderOk

`func (o *WebhooksCreate) GetSecretHeaderOk() (*string, bool)`

GetSecretHeaderOk returns a tuple with the SecretHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretHeader

`func (o *WebhooksCreate) SetSecretHeader(v string)`

SetSecretHeader sets SecretHeader field to given value.

### HasSecretHeader

`func (o *WebhooksCreate) HasSecretHeader() bool`

HasSecretHeader returns a boolean if a field has been set.

### GetSecretValue

`func (o *WebhooksCreate) GetSecretValue() string`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *WebhooksCreate) GetSecretValueOk() (*string, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *WebhooksCreate) SetSecretValue(v string)`

SetSecretValue sets SecretValue field to given value.

### HasSecretValue

`func (o *WebhooksCreate) HasSecretValue() bool`

HasSecretValue returns a boolean if a field has been set.

### GetSignatureKey

`func (o *WebhooksCreate) GetSignatureKey() string`

GetSignatureKey returns the SignatureKey field if non-nil, zero value otherwise.

### GetSignatureKeyOk

`func (o *WebhooksCreate) GetSignatureKeyOk() (*string, bool)`

GetSignatureKeyOk returns a tuple with the SignatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureKey

`func (o *WebhooksCreate) SetSignatureKey(v string)`

SetSignatureKey sets SignatureKey field to given value.

### HasSignatureKey

`func (o *WebhooksCreate) HasSignatureKey() bool`

HasSignatureKey returns a boolean if a field has been set.

### GetTargetUrl

`func (o *WebhooksCreate) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *WebhooksCreate) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *WebhooksCreate) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### GetTemplates

`func (o *WebhooksCreate) GetTemplates() []WebhooksOwnerRepoTemplates`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *WebhooksCreate) GetTemplatesOk() (*[]WebhooksOwnerRepoTemplates, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *WebhooksCreate) SetTemplates(v []WebhooksOwnerRepoTemplates)`

SetTemplates sets Templates field to given value.


### GetVerifySsl

`func (o *WebhooksCreate) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *WebhooksCreate) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *WebhooksCreate) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *WebhooksCreate) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


