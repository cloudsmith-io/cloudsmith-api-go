# RepositoryWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to **string** |  | [optional] [readonly] 
**CreatedByUrl** | Pointer to **string** |  | [optional] [readonly] 
**DisableReason** | Pointer to **int64** |  | [optional] [readonly] 
**DisableReasonStr** | Pointer to **string** |  | [optional] [readonly] 
**Events** | **[]string** |  | 
**Identifier** | Pointer to **int64** |  | [optional] [readonly] 
**IsActive** | Pointer to **bool** | If enabled, the webhook will trigger on subscribed events and send payloads to the configured target URL. | [optional] 
**IsLastResponseBad** | Pointer to **bool** |  | [optional] [readonly] 
**LastResponseStatus** | Pointer to **int64** |  | [optional] [readonly] 
**LastResponseStatusStr** | Pointer to **string** |  | [optional] [readonly] 
**NumSent** | Pointer to **int64** |  | [optional] [readonly] 
**PackageQuery** | Pointer to **NullableString** | The package-based search query for webhooks to fire. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. If a package does not match, the webhook will not fire. | [optional] 
**RequestBodyFormat** | Pointer to **int64** | The format of the payloads for webhook requests. | [optional] 
**RequestBodyFormatStr** | Pointer to **string** |  | [optional] [readonly] 
**RequestBodyTemplateFormat** | Pointer to **int64** | The format of the payloads for webhook requests. | [optional] 
**RequestBodyTemplateFormatStr** | Pointer to **string** |  | [optional] [readonly] 
**RequestContentType** | Pointer to **NullableString** | The value that will be sent for the &#39;Content Type&#39; header.  | [optional] 
**SecretHeader** | Pointer to **NullableString** | The header to send the predefined secret in. This must be unique from existing headers or it won&#39;t be sent. You can use this as a form of authentication on the endpoint side. | [optional] 
**SelfUrl** | Pointer to **string** |  | [optional] [readonly] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**TargetUrl** | **string** | The destination URL that webhook payloads will be POST&#39;ed to. | 
**Templates** | [**[]WebhookTemplate**](WebhookTemplate.md) |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedBy** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedByUrl** | Pointer to **string** |  | [optional] [readonly] 
**VerifySsl** | Pointer to **bool** | If enabled, SSL certificates is verified when webhooks are sent. It&#39;s recommended to leave this enabled as not verifying the integrity of SSL certificates leaves you susceptible to Man-in-the-Middle (MITM) attacks. | [optional] 

## Methods

### NewRepositoryWebhook

`func NewRepositoryWebhook(events []string, targetUrl string, templates []WebhookTemplate, ) *RepositoryWebhook`

NewRepositoryWebhook instantiates a new RepositoryWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryWebhookWithDefaults

`func NewRepositoryWebhookWithDefaults() *RepositoryWebhook`

NewRepositoryWebhookWithDefaults instantiates a new RepositoryWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *RepositoryWebhook) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryWebhook) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryWebhook) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepositoryWebhook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RepositoryWebhook) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RepositoryWebhook) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RepositoryWebhook) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RepositoryWebhook) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUrl

`func (o *RepositoryWebhook) GetCreatedByUrl() string`

GetCreatedByUrl returns the CreatedByUrl field if non-nil, zero value otherwise.

### GetCreatedByUrlOk

`func (o *RepositoryWebhook) GetCreatedByUrlOk() (*string, bool)`

GetCreatedByUrlOk returns a tuple with the CreatedByUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUrl

`func (o *RepositoryWebhook) SetCreatedByUrl(v string)`

SetCreatedByUrl sets CreatedByUrl field to given value.

### HasCreatedByUrl

`func (o *RepositoryWebhook) HasCreatedByUrl() bool`

HasCreatedByUrl returns a boolean if a field has been set.

### GetDisableReason

`func (o *RepositoryWebhook) GetDisableReason() int64`

GetDisableReason returns the DisableReason field if non-nil, zero value otherwise.

### GetDisableReasonOk

`func (o *RepositoryWebhook) GetDisableReasonOk() (*int64, bool)`

GetDisableReasonOk returns a tuple with the DisableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReason

`func (o *RepositoryWebhook) SetDisableReason(v int64)`

SetDisableReason sets DisableReason field to given value.

### HasDisableReason

`func (o *RepositoryWebhook) HasDisableReason() bool`

HasDisableReason returns a boolean if a field has been set.

### GetDisableReasonStr

`func (o *RepositoryWebhook) GetDisableReasonStr() string`

GetDisableReasonStr returns the DisableReasonStr field if non-nil, zero value otherwise.

### GetDisableReasonStrOk

`func (o *RepositoryWebhook) GetDisableReasonStrOk() (*string, bool)`

GetDisableReasonStrOk returns a tuple with the DisableReasonStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReasonStr

`func (o *RepositoryWebhook) SetDisableReasonStr(v string)`

SetDisableReasonStr sets DisableReasonStr field to given value.

### HasDisableReasonStr

`func (o *RepositoryWebhook) HasDisableReasonStr() bool`

HasDisableReasonStr returns a boolean if a field has been set.

### GetEvents

`func (o *RepositoryWebhook) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RepositoryWebhook) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RepositoryWebhook) SetEvents(v []string)`

SetEvents sets Events field to given value.


### SetEventsNil

`func (o *RepositoryWebhook) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *RepositoryWebhook) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetIdentifier

`func (o *RepositoryWebhook) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *RepositoryWebhook) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *RepositoryWebhook) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *RepositoryWebhook) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIsActive

`func (o *RepositoryWebhook) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RepositoryWebhook) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RepositoryWebhook) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RepositoryWebhook) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsLastResponseBad

`func (o *RepositoryWebhook) GetIsLastResponseBad() bool`

GetIsLastResponseBad returns the IsLastResponseBad field if non-nil, zero value otherwise.

### GetIsLastResponseBadOk

`func (o *RepositoryWebhook) GetIsLastResponseBadOk() (*bool, bool)`

GetIsLastResponseBadOk returns a tuple with the IsLastResponseBad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastResponseBad

`func (o *RepositoryWebhook) SetIsLastResponseBad(v bool)`

SetIsLastResponseBad sets IsLastResponseBad field to given value.

### HasIsLastResponseBad

`func (o *RepositoryWebhook) HasIsLastResponseBad() bool`

HasIsLastResponseBad returns a boolean if a field has been set.

### GetLastResponseStatus

`func (o *RepositoryWebhook) GetLastResponseStatus() int64`

GetLastResponseStatus returns the LastResponseStatus field if non-nil, zero value otherwise.

### GetLastResponseStatusOk

`func (o *RepositoryWebhook) GetLastResponseStatusOk() (*int64, bool)`

GetLastResponseStatusOk returns a tuple with the LastResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResponseStatus

`func (o *RepositoryWebhook) SetLastResponseStatus(v int64)`

SetLastResponseStatus sets LastResponseStatus field to given value.

### HasLastResponseStatus

`func (o *RepositoryWebhook) HasLastResponseStatus() bool`

HasLastResponseStatus returns a boolean if a field has been set.

### GetLastResponseStatusStr

`func (o *RepositoryWebhook) GetLastResponseStatusStr() string`

GetLastResponseStatusStr returns the LastResponseStatusStr field if non-nil, zero value otherwise.

### GetLastResponseStatusStrOk

`func (o *RepositoryWebhook) GetLastResponseStatusStrOk() (*string, bool)`

GetLastResponseStatusStrOk returns a tuple with the LastResponseStatusStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResponseStatusStr

`func (o *RepositoryWebhook) SetLastResponseStatusStr(v string)`

SetLastResponseStatusStr sets LastResponseStatusStr field to given value.

### HasLastResponseStatusStr

`func (o *RepositoryWebhook) HasLastResponseStatusStr() bool`

HasLastResponseStatusStr returns a boolean if a field has been set.

### GetNumSent

`func (o *RepositoryWebhook) GetNumSent() int64`

GetNumSent returns the NumSent field if non-nil, zero value otherwise.

### GetNumSentOk

`func (o *RepositoryWebhook) GetNumSentOk() (*int64, bool)`

GetNumSentOk returns a tuple with the NumSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSent

`func (o *RepositoryWebhook) SetNumSent(v int64)`

SetNumSent sets NumSent field to given value.

### HasNumSent

`func (o *RepositoryWebhook) HasNumSent() bool`

HasNumSent returns a boolean if a field has been set.

### GetPackageQuery

`func (o *RepositoryWebhook) GetPackageQuery() string`

GetPackageQuery returns the PackageQuery field if non-nil, zero value otherwise.

### GetPackageQueryOk

`func (o *RepositoryWebhook) GetPackageQueryOk() (*string, bool)`

GetPackageQueryOk returns a tuple with the PackageQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageQuery

`func (o *RepositoryWebhook) SetPackageQuery(v string)`

SetPackageQuery sets PackageQuery field to given value.

### HasPackageQuery

`func (o *RepositoryWebhook) HasPackageQuery() bool`

HasPackageQuery returns a boolean if a field has been set.

### SetPackageQueryNil

`func (o *RepositoryWebhook) SetPackageQueryNil(b bool)`

 SetPackageQueryNil sets the value for PackageQuery to be an explicit nil

### UnsetPackageQuery
`func (o *RepositoryWebhook) UnsetPackageQuery()`

UnsetPackageQuery ensures that no value is present for PackageQuery, not even an explicit nil
### GetRequestBodyFormat

`func (o *RepositoryWebhook) GetRequestBodyFormat() int64`

GetRequestBodyFormat returns the RequestBodyFormat field if non-nil, zero value otherwise.

### GetRequestBodyFormatOk

`func (o *RepositoryWebhook) GetRequestBodyFormatOk() (*int64, bool)`

GetRequestBodyFormatOk returns a tuple with the RequestBodyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyFormat

`func (o *RepositoryWebhook) SetRequestBodyFormat(v int64)`

SetRequestBodyFormat sets RequestBodyFormat field to given value.

### HasRequestBodyFormat

`func (o *RepositoryWebhook) HasRequestBodyFormat() bool`

HasRequestBodyFormat returns a boolean if a field has been set.

### GetRequestBodyFormatStr

`func (o *RepositoryWebhook) GetRequestBodyFormatStr() string`

GetRequestBodyFormatStr returns the RequestBodyFormatStr field if non-nil, zero value otherwise.

### GetRequestBodyFormatStrOk

`func (o *RepositoryWebhook) GetRequestBodyFormatStrOk() (*string, bool)`

GetRequestBodyFormatStrOk returns a tuple with the RequestBodyFormatStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyFormatStr

`func (o *RepositoryWebhook) SetRequestBodyFormatStr(v string)`

SetRequestBodyFormatStr sets RequestBodyFormatStr field to given value.

### HasRequestBodyFormatStr

`func (o *RepositoryWebhook) HasRequestBodyFormatStr() bool`

HasRequestBodyFormatStr returns a boolean if a field has been set.

### GetRequestBodyTemplateFormat

`func (o *RepositoryWebhook) GetRequestBodyTemplateFormat() int64`

GetRequestBodyTemplateFormat returns the RequestBodyTemplateFormat field if non-nil, zero value otherwise.

### GetRequestBodyTemplateFormatOk

`func (o *RepositoryWebhook) GetRequestBodyTemplateFormatOk() (*int64, bool)`

GetRequestBodyTemplateFormatOk returns a tuple with the RequestBodyTemplateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyTemplateFormat

`func (o *RepositoryWebhook) SetRequestBodyTemplateFormat(v int64)`

SetRequestBodyTemplateFormat sets RequestBodyTemplateFormat field to given value.

### HasRequestBodyTemplateFormat

`func (o *RepositoryWebhook) HasRequestBodyTemplateFormat() bool`

HasRequestBodyTemplateFormat returns a boolean if a field has been set.

### GetRequestBodyTemplateFormatStr

`func (o *RepositoryWebhook) GetRequestBodyTemplateFormatStr() string`

GetRequestBodyTemplateFormatStr returns the RequestBodyTemplateFormatStr field if non-nil, zero value otherwise.

### GetRequestBodyTemplateFormatStrOk

`func (o *RepositoryWebhook) GetRequestBodyTemplateFormatStrOk() (*string, bool)`

GetRequestBodyTemplateFormatStrOk returns a tuple with the RequestBodyTemplateFormatStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyTemplateFormatStr

`func (o *RepositoryWebhook) SetRequestBodyTemplateFormatStr(v string)`

SetRequestBodyTemplateFormatStr sets RequestBodyTemplateFormatStr field to given value.

### HasRequestBodyTemplateFormatStr

`func (o *RepositoryWebhook) HasRequestBodyTemplateFormatStr() bool`

HasRequestBodyTemplateFormatStr returns a boolean if a field has been set.

### GetRequestContentType

`func (o *RepositoryWebhook) GetRequestContentType() string`

GetRequestContentType returns the RequestContentType field if non-nil, zero value otherwise.

### GetRequestContentTypeOk

`func (o *RepositoryWebhook) GetRequestContentTypeOk() (*string, bool)`

GetRequestContentTypeOk returns a tuple with the RequestContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContentType

`func (o *RepositoryWebhook) SetRequestContentType(v string)`

SetRequestContentType sets RequestContentType field to given value.

### HasRequestContentType

`func (o *RepositoryWebhook) HasRequestContentType() bool`

HasRequestContentType returns a boolean if a field has been set.

### SetRequestContentTypeNil

`func (o *RepositoryWebhook) SetRequestContentTypeNil(b bool)`

 SetRequestContentTypeNil sets the value for RequestContentType to be an explicit nil

### UnsetRequestContentType
`func (o *RepositoryWebhook) UnsetRequestContentType()`

UnsetRequestContentType ensures that no value is present for RequestContentType, not even an explicit nil
### GetSecretHeader

`func (o *RepositoryWebhook) GetSecretHeader() string`

GetSecretHeader returns the SecretHeader field if non-nil, zero value otherwise.

### GetSecretHeaderOk

`func (o *RepositoryWebhook) GetSecretHeaderOk() (*string, bool)`

GetSecretHeaderOk returns a tuple with the SecretHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretHeader

`func (o *RepositoryWebhook) SetSecretHeader(v string)`

SetSecretHeader sets SecretHeader field to given value.

### HasSecretHeader

`func (o *RepositoryWebhook) HasSecretHeader() bool`

HasSecretHeader returns a boolean if a field has been set.

### SetSecretHeaderNil

`func (o *RepositoryWebhook) SetSecretHeaderNil(b bool)`

 SetSecretHeaderNil sets the value for SecretHeader to be an explicit nil

### UnsetSecretHeader
`func (o *RepositoryWebhook) UnsetSecretHeader()`

UnsetSecretHeader ensures that no value is present for SecretHeader, not even an explicit nil
### GetSelfUrl

`func (o *RepositoryWebhook) GetSelfUrl() string`

GetSelfUrl returns the SelfUrl field if non-nil, zero value otherwise.

### GetSelfUrlOk

`func (o *RepositoryWebhook) GetSelfUrlOk() (*string, bool)`

GetSelfUrlOk returns a tuple with the SelfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfUrl

`func (o *RepositoryWebhook) SetSelfUrl(v string)`

SetSelfUrl sets SelfUrl field to given value.

### HasSelfUrl

`func (o *RepositoryWebhook) HasSelfUrl() bool`

HasSelfUrl returns a boolean if a field has been set.

### GetSlugPerm

`func (o *RepositoryWebhook) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *RepositoryWebhook) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *RepositoryWebhook) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *RepositoryWebhook) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetTargetUrl

`func (o *RepositoryWebhook) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *RepositoryWebhook) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *RepositoryWebhook) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### GetTemplates

`func (o *RepositoryWebhook) GetTemplates() []WebhookTemplate`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *RepositoryWebhook) GetTemplatesOk() (*[]WebhookTemplate, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *RepositoryWebhook) SetTemplates(v []WebhookTemplate)`

SetTemplates sets Templates field to given value.


### SetTemplatesNil

`func (o *RepositoryWebhook) SetTemplatesNil(b bool)`

 SetTemplatesNil sets the value for Templates to be an explicit nil

### UnsetTemplates
`func (o *RepositoryWebhook) UnsetTemplates()`

UnsetTemplates ensures that no value is present for Templates, not even an explicit nil
### GetUpdatedAt

`func (o *RepositoryWebhook) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RepositoryWebhook) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RepositoryWebhook) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RepositoryWebhook) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RepositoryWebhook) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RepositoryWebhook) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RepositoryWebhook) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RepositoryWebhook) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedByUrl

`func (o *RepositoryWebhook) GetUpdatedByUrl() string`

GetUpdatedByUrl returns the UpdatedByUrl field if non-nil, zero value otherwise.

### GetUpdatedByUrlOk

`func (o *RepositoryWebhook) GetUpdatedByUrlOk() (*string, bool)`

GetUpdatedByUrlOk returns a tuple with the UpdatedByUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUrl

`func (o *RepositoryWebhook) SetUpdatedByUrl(v string)`

SetUpdatedByUrl sets UpdatedByUrl field to given value.

### HasUpdatedByUrl

`func (o *RepositoryWebhook) HasUpdatedByUrl() bool`

HasUpdatedByUrl returns a boolean if a field has been set.

### GetVerifySsl

`func (o *RepositoryWebhook) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *RepositoryWebhook) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *RepositoryWebhook) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *RepositoryWebhook) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


