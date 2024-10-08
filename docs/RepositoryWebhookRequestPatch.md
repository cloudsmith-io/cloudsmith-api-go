# RepositoryWebhookRequestPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to **[]string** |  | [optional] 
**IsActive** | Pointer to **bool** | If enabled, the webhook will trigger on subscribed events and send payloads to the configured target URL. | [optional] 
**PackageQuery** | Pointer to **NullableString** | The package-based search query for webhooks to fire. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. If a package does not match, the webhook will not fire. | [optional] 
**RequestBodyFormat** | Pointer to **int64** | The format of the payloads for webhook requests. Valid options are: (0) JSON, (1) JSON array, (2) form encoded JSON and (3) Handlebars template. | [optional] 
**RequestBodyTemplateFormat** | Pointer to **int64** | The format of the payloads for webhook requests. Valid options are: (0) Generic/user defined, (1) JSON and (2) XML. | [optional] 
**RequestContentType** | Pointer to **NullableString** | The value that will be sent for the &#39;Content Type&#39; header.  | [optional] 
**SecretHeader** | Pointer to **NullableString** | The header to send the predefined secret in. This must be unique from existing headers or it won&#39;t be sent. You can use this as a form of authentication on the endpoint side. | [optional] 
**SecretValue** | Pointer to **NullableString** | The value for the predefined secret (note: this is treated as a passphrase and is encrypted when we store it). You can use this as a form of authentication on the endpoint side. | [optional] 
**SignatureKey** | Pointer to **string** | The value for the signature key - This is used to generate an HMAC-based hex digest of the request body, which we send as the X-Cloudsmith-Signature header so that you can ensure that the request wasn&#39;t modified by a malicious party (note: this is treated as a passphrase and is encrypted when we store it). | [optional] 
**TargetUrl** | Pointer to **string** | The destination URL that webhook payloads will be POST&#39;ed to. | [optional] 
**Templates** | Pointer to [**[]WebhookTemplate**](WebhookTemplate.md) |  | [optional] 
**VerifySsl** | Pointer to **bool** | If enabled, SSL certificates is verified when webhooks are sent. It&#39;s recommended to leave this enabled as not verifying the integrity of SSL certificates leaves you susceptible to Man-in-the-Middle (MITM) attacks. | [optional] 

## Methods

### NewRepositoryWebhookRequestPatch

`func NewRepositoryWebhookRequestPatch() *RepositoryWebhookRequestPatch`

NewRepositoryWebhookRequestPatch instantiates a new RepositoryWebhookRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryWebhookRequestPatchWithDefaults

`func NewRepositoryWebhookRequestPatchWithDefaults() *RepositoryWebhookRequestPatch`

NewRepositoryWebhookRequestPatchWithDefaults instantiates a new RepositoryWebhookRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *RepositoryWebhookRequestPatch) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RepositoryWebhookRequestPatch) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RepositoryWebhookRequestPatch) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *RepositoryWebhookRequestPatch) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *RepositoryWebhookRequestPatch) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *RepositoryWebhookRequestPatch) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetIsActive

`func (o *RepositoryWebhookRequestPatch) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RepositoryWebhookRequestPatch) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RepositoryWebhookRequestPatch) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RepositoryWebhookRequestPatch) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetPackageQuery

`func (o *RepositoryWebhookRequestPatch) GetPackageQuery() string`

GetPackageQuery returns the PackageQuery field if non-nil, zero value otherwise.

### GetPackageQueryOk

`func (o *RepositoryWebhookRequestPatch) GetPackageQueryOk() (*string, bool)`

GetPackageQueryOk returns a tuple with the PackageQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageQuery

`func (o *RepositoryWebhookRequestPatch) SetPackageQuery(v string)`

SetPackageQuery sets PackageQuery field to given value.

### HasPackageQuery

`func (o *RepositoryWebhookRequestPatch) HasPackageQuery() bool`

HasPackageQuery returns a boolean if a field has been set.

### SetPackageQueryNil

`func (o *RepositoryWebhookRequestPatch) SetPackageQueryNil(b bool)`

 SetPackageQueryNil sets the value for PackageQuery to be an explicit nil

### UnsetPackageQuery
`func (o *RepositoryWebhookRequestPatch) UnsetPackageQuery()`

UnsetPackageQuery ensures that no value is present for PackageQuery, not even an explicit nil
### GetRequestBodyFormat

`func (o *RepositoryWebhookRequestPatch) GetRequestBodyFormat() int64`

GetRequestBodyFormat returns the RequestBodyFormat field if non-nil, zero value otherwise.

### GetRequestBodyFormatOk

`func (o *RepositoryWebhookRequestPatch) GetRequestBodyFormatOk() (*int64, bool)`

GetRequestBodyFormatOk returns a tuple with the RequestBodyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyFormat

`func (o *RepositoryWebhookRequestPatch) SetRequestBodyFormat(v int64)`

SetRequestBodyFormat sets RequestBodyFormat field to given value.

### HasRequestBodyFormat

`func (o *RepositoryWebhookRequestPatch) HasRequestBodyFormat() bool`

HasRequestBodyFormat returns a boolean if a field has been set.

### GetRequestBodyTemplateFormat

`func (o *RepositoryWebhookRequestPatch) GetRequestBodyTemplateFormat() int64`

GetRequestBodyTemplateFormat returns the RequestBodyTemplateFormat field if non-nil, zero value otherwise.

### GetRequestBodyTemplateFormatOk

`func (o *RepositoryWebhookRequestPatch) GetRequestBodyTemplateFormatOk() (*int64, bool)`

GetRequestBodyTemplateFormatOk returns a tuple with the RequestBodyTemplateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyTemplateFormat

`func (o *RepositoryWebhookRequestPatch) SetRequestBodyTemplateFormat(v int64)`

SetRequestBodyTemplateFormat sets RequestBodyTemplateFormat field to given value.

### HasRequestBodyTemplateFormat

`func (o *RepositoryWebhookRequestPatch) HasRequestBodyTemplateFormat() bool`

HasRequestBodyTemplateFormat returns a boolean if a field has been set.

### GetRequestContentType

`func (o *RepositoryWebhookRequestPatch) GetRequestContentType() string`

GetRequestContentType returns the RequestContentType field if non-nil, zero value otherwise.

### GetRequestContentTypeOk

`func (o *RepositoryWebhookRequestPatch) GetRequestContentTypeOk() (*string, bool)`

GetRequestContentTypeOk returns a tuple with the RequestContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContentType

`func (o *RepositoryWebhookRequestPatch) SetRequestContentType(v string)`

SetRequestContentType sets RequestContentType field to given value.

### HasRequestContentType

`func (o *RepositoryWebhookRequestPatch) HasRequestContentType() bool`

HasRequestContentType returns a boolean if a field has been set.

### SetRequestContentTypeNil

`func (o *RepositoryWebhookRequestPatch) SetRequestContentTypeNil(b bool)`

 SetRequestContentTypeNil sets the value for RequestContentType to be an explicit nil

### UnsetRequestContentType
`func (o *RepositoryWebhookRequestPatch) UnsetRequestContentType()`

UnsetRequestContentType ensures that no value is present for RequestContentType, not even an explicit nil
### GetSecretHeader

`func (o *RepositoryWebhookRequestPatch) GetSecretHeader() string`

GetSecretHeader returns the SecretHeader field if non-nil, zero value otherwise.

### GetSecretHeaderOk

`func (o *RepositoryWebhookRequestPatch) GetSecretHeaderOk() (*string, bool)`

GetSecretHeaderOk returns a tuple with the SecretHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretHeader

`func (o *RepositoryWebhookRequestPatch) SetSecretHeader(v string)`

SetSecretHeader sets SecretHeader field to given value.

### HasSecretHeader

`func (o *RepositoryWebhookRequestPatch) HasSecretHeader() bool`

HasSecretHeader returns a boolean if a field has been set.

### SetSecretHeaderNil

`func (o *RepositoryWebhookRequestPatch) SetSecretHeaderNil(b bool)`

 SetSecretHeaderNil sets the value for SecretHeader to be an explicit nil

### UnsetSecretHeader
`func (o *RepositoryWebhookRequestPatch) UnsetSecretHeader()`

UnsetSecretHeader ensures that no value is present for SecretHeader, not even an explicit nil
### GetSecretValue

`func (o *RepositoryWebhookRequestPatch) GetSecretValue() string`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *RepositoryWebhookRequestPatch) GetSecretValueOk() (*string, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *RepositoryWebhookRequestPatch) SetSecretValue(v string)`

SetSecretValue sets SecretValue field to given value.

### HasSecretValue

`func (o *RepositoryWebhookRequestPatch) HasSecretValue() bool`

HasSecretValue returns a boolean if a field has been set.

### SetSecretValueNil

`func (o *RepositoryWebhookRequestPatch) SetSecretValueNil(b bool)`

 SetSecretValueNil sets the value for SecretValue to be an explicit nil

### UnsetSecretValue
`func (o *RepositoryWebhookRequestPatch) UnsetSecretValue()`

UnsetSecretValue ensures that no value is present for SecretValue, not even an explicit nil
### GetSignatureKey

`func (o *RepositoryWebhookRequestPatch) GetSignatureKey() string`

GetSignatureKey returns the SignatureKey field if non-nil, zero value otherwise.

### GetSignatureKeyOk

`func (o *RepositoryWebhookRequestPatch) GetSignatureKeyOk() (*string, bool)`

GetSignatureKeyOk returns a tuple with the SignatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureKey

`func (o *RepositoryWebhookRequestPatch) SetSignatureKey(v string)`

SetSignatureKey sets SignatureKey field to given value.

### HasSignatureKey

`func (o *RepositoryWebhookRequestPatch) HasSignatureKey() bool`

HasSignatureKey returns a boolean if a field has been set.

### GetTargetUrl

`func (o *RepositoryWebhookRequestPatch) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *RepositoryWebhookRequestPatch) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *RepositoryWebhookRequestPatch) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *RepositoryWebhookRequestPatch) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetTemplates

`func (o *RepositoryWebhookRequestPatch) GetTemplates() []WebhookTemplate`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *RepositoryWebhookRequestPatch) GetTemplatesOk() (*[]WebhookTemplate, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *RepositoryWebhookRequestPatch) SetTemplates(v []WebhookTemplate)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *RepositoryWebhookRequestPatch) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### SetTemplatesNil

`func (o *RepositoryWebhookRequestPatch) SetTemplatesNil(b bool)`

 SetTemplatesNil sets the value for Templates to be an explicit nil

### UnsetTemplates
`func (o *RepositoryWebhookRequestPatch) UnsetTemplates()`

UnsetTemplates ensures that no value is present for Templates, not even an explicit nil
### GetVerifySsl

`func (o *RepositoryWebhookRequestPatch) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *RepositoryWebhookRequestPatch) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *RepositoryWebhookRequestPatch) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *RepositoryWebhookRequestPatch) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


