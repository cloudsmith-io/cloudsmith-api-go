# RepositoryWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | **[]string** |  | 
**IsActive** | Pointer to **bool** | If enabled, the webhook will trigger on subscribed events and send payloads to the configured target URL. | [optional] 
**PackageQuery** | Pointer to **NullableString** | The package-based search query for webhooks to fire. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. If a package does not match, the webhook will not fire. | [optional] 
**RequestBodyFormat** | Pointer to **int64** | The format of the payloads for webhook requests. | [optional] 
**RequestBodyTemplateFormat** | Pointer to **int64** | The format of the payloads for webhook requests. | [optional] 
**RequestContentType** | Pointer to **NullableString** | The value that will be sent for the &#39;Content Type&#39; header.  | [optional] 
**SecretHeader** | Pointer to **NullableString** | The header to send the predefined secret in. This must be unique from existing headers or it won&#39;t be sent. You can use this as a form of authentication on the endpoint side. | [optional] 
**SecretValue** | Pointer to **NullableString** | The value for the predefined secret (note: this is treated as a passphrase and is encrypted when we store it). You can use this as a form of authentication on the endpoint side. | [optional] 
**SignatureKey** | Pointer to **string** | The value for the signature key - This is used to generate an HMAC-based hex digest of the request body, which we send as the X-Cloudsmith-Signature header so that you can ensure that the request wasn&#39;t modified by a malicious party (note: this is treated as a passphrase and is encrypted when we store it). | [optional] 
**TargetUrl** | **string** | The destination URL that webhook payloads will be POST&#39;ed to. | 
**Templates** | [**[]WebhookTemplate**](WebhookTemplate.md) |  | 
**VerifySsl** | Pointer to **bool** | If enabled, SSL certificates is verified when webhooks are sent. It&#39;s recommended to leave this enabled as not verifying the integrity of SSL certificates leaves you susceptible to Man-in-the-Middle (MITM) attacks. | [optional] 

## Methods

### NewRepositoryWebhookRequest

`func NewRepositoryWebhookRequest(events []string, targetUrl string, templates []WebhookTemplate, ) *RepositoryWebhookRequest`

NewRepositoryWebhookRequest instantiates a new RepositoryWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryWebhookRequestWithDefaults

`func NewRepositoryWebhookRequestWithDefaults() *RepositoryWebhookRequest`

NewRepositoryWebhookRequestWithDefaults instantiates a new RepositoryWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *RepositoryWebhookRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RepositoryWebhookRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RepositoryWebhookRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.


### SetEventsNil

`func (o *RepositoryWebhookRequest) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *RepositoryWebhookRequest) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetIsActive

`func (o *RepositoryWebhookRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RepositoryWebhookRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RepositoryWebhookRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RepositoryWebhookRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetPackageQuery

`func (o *RepositoryWebhookRequest) GetPackageQuery() string`

GetPackageQuery returns the PackageQuery field if non-nil, zero value otherwise.

### GetPackageQueryOk

`func (o *RepositoryWebhookRequest) GetPackageQueryOk() (*string, bool)`

GetPackageQueryOk returns a tuple with the PackageQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageQuery

`func (o *RepositoryWebhookRequest) SetPackageQuery(v string)`

SetPackageQuery sets PackageQuery field to given value.

### HasPackageQuery

`func (o *RepositoryWebhookRequest) HasPackageQuery() bool`

HasPackageQuery returns a boolean if a field has been set.

### SetPackageQueryNil

`func (o *RepositoryWebhookRequest) SetPackageQueryNil(b bool)`

 SetPackageQueryNil sets the value for PackageQuery to be an explicit nil

### UnsetPackageQuery
`func (o *RepositoryWebhookRequest) UnsetPackageQuery()`

UnsetPackageQuery ensures that no value is present for PackageQuery, not even an explicit nil
### GetRequestBodyFormat

`func (o *RepositoryWebhookRequest) GetRequestBodyFormat() int64`

GetRequestBodyFormat returns the RequestBodyFormat field if non-nil, zero value otherwise.

### GetRequestBodyFormatOk

`func (o *RepositoryWebhookRequest) GetRequestBodyFormatOk() (*int64, bool)`

GetRequestBodyFormatOk returns a tuple with the RequestBodyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyFormat

`func (o *RepositoryWebhookRequest) SetRequestBodyFormat(v int64)`

SetRequestBodyFormat sets RequestBodyFormat field to given value.

### HasRequestBodyFormat

`func (o *RepositoryWebhookRequest) HasRequestBodyFormat() bool`

HasRequestBodyFormat returns a boolean if a field has been set.

### GetRequestBodyTemplateFormat

`func (o *RepositoryWebhookRequest) GetRequestBodyTemplateFormat() int64`

GetRequestBodyTemplateFormat returns the RequestBodyTemplateFormat field if non-nil, zero value otherwise.

### GetRequestBodyTemplateFormatOk

`func (o *RepositoryWebhookRequest) GetRequestBodyTemplateFormatOk() (*int64, bool)`

GetRequestBodyTemplateFormatOk returns a tuple with the RequestBodyTemplateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyTemplateFormat

`func (o *RepositoryWebhookRequest) SetRequestBodyTemplateFormat(v int64)`

SetRequestBodyTemplateFormat sets RequestBodyTemplateFormat field to given value.

### HasRequestBodyTemplateFormat

`func (o *RepositoryWebhookRequest) HasRequestBodyTemplateFormat() bool`

HasRequestBodyTemplateFormat returns a boolean if a field has been set.

### GetRequestContentType

`func (o *RepositoryWebhookRequest) GetRequestContentType() string`

GetRequestContentType returns the RequestContentType field if non-nil, zero value otherwise.

### GetRequestContentTypeOk

`func (o *RepositoryWebhookRequest) GetRequestContentTypeOk() (*string, bool)`

GetRequestContentTypeOk returns a tuple with the RequestContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContentType

`func (o *RepositoryWebhookRequest) SetRequestContentType(v string)`

SetRequestContentType sets RequestContentType field to given value.

### HasRequestContentType

`func (o *RepositoryWebhookRequest) HasRequestContentType() bool`

HasRequestContentType returns a boolean if a field has been set.

### SetRequestContentTypeNil

`func (o *RepositoryWebhookRequest) SetRequestContentTypeNil(b bool)`

 SetRequestContentTypeNil sets the value for RequestContentType to be an explicit nil

### UnsetRequestContentType
`func (o *RepositoryWebhookRequest) UnsetRequestContentType()`

UnsetRequestContentType ensures that no value is present for RequestContentType, not even an explicit nil
### GetSecretHeader

`func (o *RepositoryWebhookRequest) GetSecretHeader() string`

GetSecretHeader returns the SecretHeader field if non-nil, zero value otherwise.

### GetSecretHeaderOk

`func (o *RepositoryWebhookRequest) GetSecretHeaderOk() (*string, bool)`

GetSecretHeaderOk returns a tuple with the SecretHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretHeader

`func (o *RepositoryWebhookRequest) SetSecretHeader(v string)`

SetSecretHeader sets SecretHeader field to given value.

### HasSecretHeader

`func (o *RepositoryWebhookRequest) HasSecretHeader() bool`

HasSecretHeader returns a boolean if a field has been set.

### SetSecretHeaderNil

`func (o *RepositoryWebhookRequest) SetSecretHeaderNil(b bool)`

 SetSecretHeaderNil sets the value for SecretHeader to be an explicit nil

### UnsetSecretHeader
`func (o *RepositoryWebhookRequest) UnsetSecretHeader()`

UnsetSecretHeader ensures that no value is present for SecretHeader, not even an explicit nil
### GetSecretValue

`func (o *RepositoryWebhookRequest) GetSecretValue() string`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *RepositoryWebhookRequest) GetSecretValueOk() (*string, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *RepositoryWebhookRequest) SetSecretValue(v string)`

SetSecretValue sets SecretValue field to given value.

### HasSecretValue

`func (o *RepositoryWebhookRequest) HasSecretValue() bool`

HasSecretValue returns a boolean if a field has been set.

### SetSecretValueNil

`func (o *RepositoryWebhookRequest) SetSecretValueNil(b bool)`

 SetSecretValueNil sets the value for SecretValue to be an explicit nil

### UnsetSecretValue
`func (o *RepositoryWebhookRequest) UnsetSecretValue()`

UnsetSecretValue ensures that no value is present for SecretValue, not even an explicit nil
### GetSignatureKey

`func (o *RepositoryWebhookRequest) GetSignatureKey() string`

GetSignatureKey returns the SignatureKey field if non-nil, zero value otherwise.

### GetSignatureKeyOk

`func (o *RepositoryWebhookRequest) GetSignatureKeyOk() (*string, bool)`

GetSignatureKeyOk returns a tuple with the SignatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureKey

`func (o *RepositoryWebhookRequest) SetSignatureKey(v string)`

SetSignatureKey sets SignatureKey field to given value.

### HasSignatureKey

`func (o *RepositoryWebhookRequest) HasSignatureKey() bool`

HasSignatureKey returns a boolean if a field has been set.

### GetTargetUrl

`func (o *RepositoryWebhookRequest) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *RepositoryWebhookRequest) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *RepositoryWebhookRequest) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### GetTemplates

`func (o *RepositoryWebhookRequest) GetTemplates() []WebhookTemplate`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *RepositoryWebhookRequest) GetTemplatesOk() (*[]WebhookTemplate, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *RepositoryWebhookRequest) SetTemplates(v []WebhookTemplate)`

SetTemplates sets Templates field to given value.


### SetTemplatesNil

`func (o *RepositoryWebhookRequest) SetTemplatesNil(b bool)`

 SetTemplatesNil sets the value for Templates to be an explicit nil

### UnsetTemplates
`func (o *RepositoryWebhookRequest) UnsetTemplates()`

UnsetTemplates ensures that no value is present for Templates, not even an explicit nil
### GetVerifySsl

`func (o *RepositoryWebhookRequest) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *RepositoryWebhookRequest) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *RepositoryWebhookRequest) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *RepositoryWebhookRequest) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


