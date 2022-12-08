# RepositoryWebhookResponse

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

### NewRepositoryWebhookResponse

`func NewRepositoryWebhookResponse(events []string, targetUrl string, templates []WebhookTemplate, ) *RepositoryWebhookResponse`

NewRepositoryWebhookResponse instantiates a new RepositoryWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryWebhookResponseWithDefaults

`func NewRepositoryWebhookResponseWithDefaults() *RepositoryWebhookResponse`

NewRepositoryWebhookResponseWithDefaults instantiates a new RepositoryWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *RepositoryWebhookResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryWebhookResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryWebhookResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepositoryWebhookResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RepositoryWebhookResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RepositoryWebhookResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RepositoryWebhookResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RepositoryWebhookResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUrl

`func (o *RepositoryWebhookResponse) GetCreatedByUrl() string`

GetCreatedByUrl returns the CreatedByUrl field if non-nil, zero value otherwise.

### GetCreatedByUrlOk

`func (o *RepositoryWebhookResponse) GetCreatedByUrlOk() (*string, bool)`

GetCreatedByUrlOk returns a tuple with the CreatedByUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUrl

`func (o *RepositoryWebhookResponse) SetCreatedByUrl(v string)`

SetCreatedByUrl sets CreatedByUrl field to given value.

### HasCreatedByUrl

`func (o *RepositoryWebhookResponse) HasCreatedByUrl() bool`

HasCreatedByUrl returns a boolean if a field has been set.

### GetDisableReason

`func (o *RepositoryWebhookResponse) GetDisableReason() int64`

GetDisableReason returns the DisableReason field if non-nil, zero value otherwise.

### GetDisableReasonOk

`func (o *RepositoryWebhookResponse) GetDisableReasonOk() (*int64, bool)`

GetDisableReasonOk returns a tuple with the DisableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReason

`func (o *RepositoryWebhookResponse) SetDisableReason(v int64)`

SetDisableReason sets DisableReason field to given value.

### HasDisableReason

`func (o *RepositoryWebhookResponse) HasDisableReason() bool`

HasDisableReason returns a boolean if a field has been set.

### GetDisableReasonStr

`func (o *RepositoryWebhookResponse) GetDisableReasonStr() string`

GetDisableReasonStr returns the DisableReasonStr field if non-nil, zero value otherwise.

### GetDisableReasonStrOk

`func (o *RepositoryWebhookResponse) GetDisableReasonStrOk() (*string, bool)`

GetDisableReasonStrOk returns a tuple with the DisableReasonStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReasonStr

`func (o *RepositoryWebhookResponse) SetDisableReasonStr(v string)`

SetDisableReasonStr sets DisableReasonStr field to given value.

### HasDisableReasonStr

`func (o *RepositoryWebhookResponse) HasDisableReasonStr() bool`

HasDisableReasonStr returns a boolean if a field has been set.

### GetEvents

`func (o *RepositoryWebhookResponse) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RepositoryWebhookResponse) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RepositoryWebhookResponse) SetEvents(v []string)`

SetEvents sets Events field to given value.


### SetEventsNil

`func (o *RepositoryWebhookResponse) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *RepositoryWebhookResponse) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetIdentifier

`func (o *RepositoryWebhookResponse) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *RepositoryWebhookResponse) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *RepositoryWebhookResponse) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *RepositoryWebhookResponse) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIsActive

`func (o *RepositoryWebhookResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RepositoryWebhookResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RepositoryWebhookResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RepositoryWebhookResponse) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsLastResponseBad

`func (o *RepositoryWebhookResponse) GetIsLastResponseBad() bool`

GetIsLastResponseBad returns the IsLastResponseBad field if non-nil, zero value otherwise.

### GetIsLastResponseBadOk

`func (o *RepositoryWebhookResponse) GetIsLastResponseBadOk() (*bool, bool)`

GetIsLastResponseBadOk returns a tuple with the IsLastResponseBad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastResponseBad

`func (o *RepositoryWebhookResponse) SetIsLastResponseBad(v bool)`

SetIsLastResponseBad sets IsLastResponseBad field to given value.

### HasIsLastResponseBad

`func (o *RepositoryWebhookResponse) HasIsLastResponseBad() bool`

HasIsLastResponseBad returns a boolean if a field has been set.

### GetLastResponseStatus

`func (o *RepositoryWebhookResponse) GetLastResponseStatus() int64`

GetLastResponseStatus returns the LastResponseStatus field if non-nil, zero value otherwise.

### GetLastResponseStatusOk

`func (o *RepositoryWebhookResponse) GetLastResponseStatusOk() (*int64, bool)`

GetLastResponseStatusOk returns a tuple with the LastResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResponseStatus

`func (o *RepositoryWebhookResponse) SetLastResponseStatus(v int64)`

SetLastResponseStatus sets LastResponseStatus field to given value.

### HasLastResponseStatus

`func (o *RepositoryWebhookResponse) HasLastResponseStatus() bool`

HasLastResponseStatus returns a boolean if a field has been set.

### GetLastResponseStatusStr

`func (o *RepositoryWebhookResponse) GetLastResponseStatusStr() string`

GetLastResponseStatusStr returns the LastResponseStatusStr field if non-nil, zero value otherwise.

### GetLastResponseStatusStrOk

`func (o *RepositoryWebhookResponse) GetLastResponseStatusStrOk() (*string, bool)`

GetLastResponseStatusStrOk returns a tuple with the LastResponseStatusStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResponseStatusStr

`func (o *RepositoryWebhookResponse) SetLastResponseStatusStr(v string)`

SetLastResponseStatusStr sets LastResponseStatusStr field to given value.

### HasLastResponseStatusStr

`func (o *RepositoryWebhookResponse) HasLastResponseStatusStr() bool`

HasLastResponseStatusStr returns a boolean if a field has been set.

### GetNumSent

`func (o *RepositoryWebhookResponse) GetNumSent() int64`

GetNumSent returns the NumSent field if non-nil, zero value otherwise.

### GetNumSentOk

`func (o *RepositoryWebhookResponse) GetNumSentOk() (*int64, bool)`

GetNumSentOk returns a tuple with the NumSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSent

`func (o *RepositoryWebhookResponse) SetNumSent(v int64)`

SetNumSent sets NumSent field to given value.

### HasNumSent

`func (o *RepositoryWebhookResponse) HasNumSent() bool`

HasNumSent returns a boolean if a field has been set.

### GetPackageQuery

`func (o *RepositoryWebhookResponse) GetPackageQuery() string`

GetPackageQuery returns the PackageQuery field if non-nil, zero value otherwise.

### GetPackageQueryOk

`func (o *RepositoryWebhookResponse) GetPackageQueryOk() (*string, bool)`

GetPackageQueryOk returns a tuple with the PackageQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageQuery

`func (o *RepositoryWebhookResponse) SetPackageQuery(v string)`

SetPackageQuery sets PackageQuery field to given value.

### HasPackageQuery

`func (o *RepositoryWebhookResponse) HasPackageQuery() bool`

HasPackageQuery returns a boolean if a field has been set.

### SetPackageQueryNil

`func (o *RepositoryWebhookResponse) SetPackageQueryNil(b bool)`

 SetPackageQueryNil sets the value for PackageQuery to be an explicit nil

### UnsetPackageQuery
`func (o *RepositoryWebhookResponse) UnsetPackageQuery()`

UnsetPackageQuery ensures that no value is present for PackageQuery, not even an explicit nil
### GetRequestBodyFormat

`func (o *RepositoryWebhookResponse) GetRequestBodyFormat() int64`

GetRequestBodyFormat returns the RequestBodyFormat field if non-nil, zero value otherwise.

### GetRequestBodyFormatOk

`func (o *RepositoryWebhookResponse) GetRequestBodyFormatOk() (*int64, bool)`

GetRequestBodyFormatOk returns a tuple with the RequestBodyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyFormat

`func (o *RepositoryWebhookResponse) SetRequestBodyFormat(v int64)`

SetRequestBodyFormat sets RequestBodyFormat field to given value.

### HasRequestBodyFormat

`func (o *RepositoryWebhookResponse) HasRequestBodyFormat() bool`

HasRequestBodyFormat returns a boolean if a field has been set.

### GetRequestBodyFormatStr

`func (o *RepositoryWebhookResponse) GetRequestBodyFormatStr() string`

GetRequestBodyFormatStr returns the RequestBodyFormatStr field if non-nil, zero value otherwise.

### GetRequestBodyFormatStrOk

`func (o *RepositoryWebhookResponse) GetRequestBodyFormatStrOk() (*string, bool)`

GetRequestBodyFormatStrOk returns a tuple with the RequestBodyFormatStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyFormatStr

`func (o *RepositoryWebhookResponse) SetRequestBodyFormatStr(v string)`

SetRequestBodyFormatStr sets RequestBodyFormatStr field to given value.

### HasRequestBodyFormatStr

`func (o *RepositoryWebhookResponse) HasRequestBodyFormatStr() bool`

HasRequestBodyFormatStr returns a boolean if a field has been set.

### GetRequestBodyTemplateFormat

`func (o *RepositoryWebhookResponse) GetRequestBodyTemplateFormat() int64`

GetRequestBodyTemplateFormat returns the RequestBodyTemplateFormat field if non-nil, zero value otherwise.

### GetRequestBodyTemplateFormatOk

`func (o *RepositoryWebhookResponse) GetRequestBodyTemplateFormatOk() (*int64, bool)`

GetRequestBodyTemplateFormatOk returns a tuple with the RequestBodyTemplateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyTemplateFormat

`func (o *RepositoryWebhookResponse) SetRequestBodyTemplateFormat(v int64)`

SetRequestBodyTemplateFormat sets RequestBodyTemplateFormat field to given value.

### HasRequestBodyTemplateFormat

`func (o *RepositoryWebhookResponse) HasRequestBodyTemplateFormat() bool`

HasRequestBodyTemplateFormat returns a boolean if a field has been set.

### GetRequestBodyTemplateFormatStr

`func (o *RepositoryWebhookResponse) GetRequestBodyTemplateFormatStr() string`

GetRequestBodyTemplateFormatStr returns the RequestBodyTemplateFormatStr field if non-nil, zero value otherwise.

### GetRequestBodyTemplateFormatStrOk

`func (o *RepositoryWebhookResponse) GetRequestBodyTemplateFormatStrOk() (*string, bool)`

GetRequestBodyTemplateFormatStrOk returns a tuple with the RequestBodyTemplateFormatStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyTemplateFormatStr

`func (o *RepositoryWebhookResponse) SetRequestBodyTemplateFormatStr(v string)`

SetRequestBodyTemplateFormatStr sets RequestBodyTemplateFormatStr field to given value.

### HasRequestBodyTemplateFormatStr

`func (o *RepositoryWebhookResponse) HasRequestBodyTemplateFormatStr() bool`

HasRequestBodyTemplateFormatStr returns a boolean if a field has been set.

### GetRequestContentType

`func (o *RepositoryWebhookResponse) GetRequestContentType() string`

GetRequestContentType returns the RequestContentType field if non-nil, zero value otherwise.

### GetRequestContentTypeOk

`func (o *RepositoryWebhookResponse) GetRequestContentTypeOk() (*string, bool)`

GetRequestContentTypeOk returns a tuple with the RequestContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContentType

`func (o *RepositoryWebhookResponse) SetRequestContentType(v string)`

SetRequestContentType sets RequestContentType field to given value.

### HasRequestContentType

`func (o *RepositoryWebhookResponse) HasRequestContentType() bool`

HasRequestContentType returns a boolean if a field has been set.

### SetRequestContentTypeNil

`func (o *RepositoryWebhookResponse) SetRequestContentTypeNil(b bool)`

 SetRequestContentTypeNil sets the value for RequestContentType to be an explicit nil

### UnsetRequestContentType
`func (o *RepositoryWebhookResponse) UnsetRequestContentType()`

UnsetRequestContentType ensures that no value is present for RequestContentType, not even an explicit nil
### GetSecretHeader

`func (o *RepositoryWebhookResponse) GetSecretHeader() string`

GetSecretHeader returns the SecretHeader field if non-nil, zero value otherwise.

### GetSecretHeaderOk

`func (o *RepositoryWebhookResponse) GetSecretHeaderOk() (*string, bool)`

GetSecretHeaderOk returns a tuple with the SecretHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretHeader

`func (o *RepositoryWebhookResponse) SetSecretHeader(v string)`

SetSecretHeader sets SecretHeader field to given value.

### HasSecretHeader

`func (o *RepositoryWebhookResponse) HasSecretHeader() bool`

HasSecretHeader returns a boolean if a field has been set.

### SetSecretHeaderNil

`func (o *RepositoryWebhookResponse) SetSecretHeaderNil(b bool)`

 SetSecretHeaderNil sets the value for SecretHeader to be an explicit nil

### UnsetSecretHeader
`func (o *RepositoryWebhookResponse) UnsetSecretHeader()`

UnsetSecretHeader ensures that no value is present for SecretHeader, not even an explicit nil
### GetSelfUrl

`func (o *RepositoryWebhookResponse) GetSelfUrl() string`

GetSelfUrl returns the SelfUrl field if non-nil, zero value otherwise.

### GetSelfUrlOk

`func (o *RepositoryWebhookResponse) GetSelfUrlOk() (*string, bool)`

GetSelfUrlOk returns a tuple with the SelfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfUrl

`func (o *RepositoryWebhookResponse) SetSelfUrl(v string)`

SetSelfUrl sets SelfUrl field to given value.

### HasSelfUrl

`func (o *RepositoryWebhookResponse) HasSelfUrl() bool`

HasSelfUrl returns a boolean if a field has been set.

### GetSlugPerm

`func (o *RepositoryWebhookResponse) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *RepositoryWebhookResponse) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *RepositoryWebhookResponse) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *RepositoryWebhookResponse) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetTargetUrl

`func (o *RepositoryWebhookResponse) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *RepositoryWebhookResponse) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *RepositoryWebhookResponse) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### GetTemplates

`func (o *RepositoryWebhookResponse) GetTemplates() []WebhookTemplate`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *RepositoryWebhookResponse) GetTemplatesOk() (*[]WebhookTemplate, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *RepositoryWebhookResponse) SetTemplates(v []WebhookTemplate)`

SetTemplates sets Templates field to given value.


### SetTemplatesNil

`func (o *RepositoryWebhookResponse) SetTemplatesNil(b bool)`

 SetTemplatesNil sets the value for Templates to be an explicit nil

### UnsetTemplates
`func (o *RepositoryWebhookResponse) UnsetTemplates()`

UnsetTemplates ensures that no value is present for Templates, not even an explicit nil
### GetUpdatedAt

`func (o *RepositoryWebhookResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RepositoryWebhookResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RepositoryWebhookResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RepositoryWebhookResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RepositoryWebhookResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RepositoryWebhookResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RepositoryWebhookResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RepositoryWebhookResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedByUrl

`func (o *RepositoryWebhookResponse) GetUpdatedByUrl() string`

GetUpdatedByUrl returns the UpdatedByUrl field if non-nil, zero value otherwise.

### GetUpdatedByUrlOk

`func (o *RepositoryWebhookResponse) GetUpdatedByUrlOk() (*string, bool)`

GetUpdatedByUrlOk returns a tuple with the UpdatedByUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUrl

`func (o *RepositoryWebhookResponse) SetUpdatedByUrl(v string)`

SetUpdatedByUrl sets UpdatedByUrl field to given value.

### HasUpdatedByUrl

`func (o *RepositoryWebhookResponse) HasUpdatedByUrl() bool`

HasUpdatedByUrl returns a boolean if a field has been set.

### GetVerifySsl

`func (o *RepositoryWebhookResponse) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *RepositoryWebhookResponse) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *RepositoryWebhookResponse) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *RepositoryWebhookResponse) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


