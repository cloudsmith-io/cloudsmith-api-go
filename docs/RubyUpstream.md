# RubyUpstream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | Pointer to **string** | The authentication mode to use when accessing this upstream.  | [optional] [default to "None"]
**AuthSecret** | Pointer to **NullableString** | Secret to provide with requests to upstream. | [optional] 
**AuthUsername** | Pointer to **NullableString** | Username to provide with requests to upstream. | [optional] 
**Available** | Pointer to **bool** | Whether the upstream is available for use. | [optional] [readonly] 
**CanReindex** | Pointer to **bool** | Whether the upstream can be reindexed. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The datetime the upstream source was created. | [optional] [readonly] 
**DisableReason** | Pointer to **string** |  | [optional] [readonly] [default to "N/A"]
**DisableReasonText** | Pointer to **string** | Human-readable explanation of why this upstream is disabled | [optional] [readonly] 
**ExtraHeader1** | Pointer to **NullableString** | The key for extra header #1 to send to upstream. | [optional] 
**ExtraHeader2** | Pointer to **NullableString** | The key for extra header #2 to send to upstream. | [optional] 
**ExtraValue1** | Pointer to **NullableString** | The value for extra header #1 to send to upstream. This is stored as plaintext, and is NOT encrypted. | [optional] 
**ExtraValue2** | Pointer to **NullableString** | The value for extra header #2 to send to upstream. This is stored as plaintext, and is NOT encrypted. | [optional] 
**HasFailedSignatureVerification** | Pointer to **bool** | Whether the upstream has failed signature verification. | [optional] [readonly] 
**IndexPackageCount** | Pointer to **string** | The number of packages available in this upstream source | [optional] [readonly] 
**IndexStatus** | Pointer to **string** | The current indexing status of this upstream source | [optional] [readonly] 
**IsActive** | Pointer to **bool** | Whether or not this upstream is active and ready for requests. | [optional] 
**LastIndexed** | Pointer to **string** | The last time this upstream source was indexed | [optional] [readonly] 
**Mode** | Pointer to **string** | The mode that this upstream should operate in. Upstream sources can be used to proxy resolved packages, as well as operate in a proxy/cache or cache only mode. | [optional] [default to "Proxy Only"]
**Name** | **string** | A descriptive name for this upstream source. A shortened version of this name will be used for tagging cached packages retrieved from this upstream. | 
**PendingValidation** | Pointer to **bool** | When true, this upstream source is pending validation. | [optional] [readonly] 
**Priority** | Pointer to **int64** | Upstream sources are selected for resolving requests by sequential order (1..n), followed by creation date. | [optional] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpstreamUrl** | **string** | The URL for this upstream source. This must be a fully qualified URL including any path elements required to reach the root of the repository.  | 
**VerifySsl** | Pointer to **bool** | If enabled, SSL certificates are verified when requests are made to this upstream. It&#39;s recommended to leave this enabled for all public sources to help mitigate Man-In-The-Middle (MITM) attacks. Please note this only applies to HTTPS upstreams. | [optional] 

## Methods

### NewRubyUpstream

`func NewRubyUpstream(name string, upstreamUrl string, ) *RubyUpstream`

NewRubyUpstream instantiates a new RubyUpstream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRubyUpstreamWithDefaults

`func NewRubyUpstreamWithDefaults() *RubyUpstream`

NewRubyUpstreamWithDefaults instantiates a new RubyUpstream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *RubyUpstream) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *RubyUpstream) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *RubyUpstream) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *RubyUpstream) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAuthSecret

`func (o *RubyUpstream) GetAuthSecret() string`

GetAuthSecret returns the AuthSecret field if non-nil, zero value otherwise.

### GetAuthSecretOk

`func (o *RubyUpstream) GetAuthSecretOk() (*string, bool)`

GetAuthSecretOk returns a tuple with the AuthSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecret

`func (o *RubyUpstream) SetAuthSecret(v string)`

SetAuthSecret sets AuthSecret field to given value.

### HasAuthSecret

`func (o *RubyUpstream) HasAuthSecret() bool`

HasAuthSecret returns a boolean if a field has been set.

### SetAuthSecretNil

`func (o *RubyUpstream) SetAuthSecretNil(b bool)`

 SetAuthSecretNil sets the value for AuthSecret to be an explicit nil

### UnsetAuthSecret
`func (o *RubyUpstream) UnsetAuthSecret()`

UnsetAuthSecret ensures that no value is present for AuthSecret, not even an explicit nil
### GetAuthUsername

`func (o *RubyUpstream) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *RubyUpstream) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *RubyUpstream) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *RubyUpstream) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### SetAuthUsernameNil

`func (o *RubyUpstream) SetAuthUsernameNil(b bool)`

 SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil

### UnsetAuthUsername
`func (o *RubyUpstream) UnsetAuthUsername()`

UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
### GetAvailable

`func (o *RubyUpstream) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *RubyUpstream) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *RubyUpstream) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *RubyUpstream) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCanReindex

`func (o *RubyUpstream) GetCanReindex() bool`

GetCanReindex returns the CanReindex field if non-nil, zero value otherwise.

### GetCanReindexOk

`func (o *RubyUpstream) GetCanReindexOk() (*bool, bool)`

GetCanReindexOk returns a tuple with the CanReindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReindex

`func (o *RubyUpstream) SetCanReindex(v bool)`

SetCanReindex sets CanReindex field to given value.

### HasCanReindex

`func (o *RubyUpstream) HasCanReindex() bool`

HasCanReindex returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RubyUpstream) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RubyUpstream) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RubyUpstream) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RubyUpstream) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisableReason

`func (o *RubyUpstream) GetDisableReason() string`

GetDisableReason returns the DisableReason field if non-nil, zero value otherwise.

### GetDisableReasonOk

`func (o *RubyUpstream) GetDisableReasonOk() (*string, bool)`

GetDisableReasonOk returns a tuple with the DisableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReason

`func (o *RubyUpstream) SetDisableReason(v string)`

SetDisableReason sets DisableReason field to given value.

### HasDisableReason

`func (o *RubyUpstream) HasDisableReason() bool`

HasDisableReason returns a boolean if a field has been set.

### GetDisableReasonText

`func (o *RubyUpstream) GetDisableReasonText() string`

GetDisableReasonText returns the DisableReasonText field if non-nil, zero value otherwise.

### GetDisableReasonTextOk

`func (o *RubyUpstream) GetDisableReasonTextOk() (*string, bool)`

GetDisableReasonTextOk returns a tuple with the DisableReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReasonText

`func (o *RubyUpstream) SetDisableReasonText(v string)`

SetDisableReasonText sets DisableReasonText field to given value.

### HasDisableReasonText

`func (o *RubyUpstream) HasDisableReasonText() bool`

HasDisableReasonText returns a boolean if a field has been set.

### GetExtraHeader1

`func (o *RubyUpstream) GetExtraHeader1() string`

GetExtraHeader1 returns the ExtraHeader1 field if non-nil, zero value otherwise.

### GetExtraHeader1Ok

`func (o *RubyUpstream) GetExtraHeader1Ok() (*string, bool)`

GetExtraHeader1Ok returns a tuple with the ExtraHeader1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader1

`func (o *RubyUpstream) SetExtraHeader1(v string)`

SetExtraHeader1 sets ExtraHeader1 field to given value.

### HasExtraHeader1

`func (o *RubyUpstream) HasExtraHeader1() bool`

HasExtraHeader1 returns a boolean if a field has been set.

### SetExtraHeader1Nil

`func (o *RubyUpstream) SetExtraHeader1Nil(b bool)`

 SetExtraHeader1Nil sets the value for ExtraHeader1 to be an explicit nil

### UnsetExtraHeader1
`func (o *RubyUpstream) UnsetExtraHeader1()`

UnsetExtraHeader1 ensures that no value is present for ExtraHeader1, not even an explicit nil
### GetExtraHeader2

`func (o *RubyUpstream) GetExtraHeader2() string`

GetExtraHeader2 returns the ExtraHeader2 field if non-nil, zero value otherwise.

### GetExtraHeader2Ok

`func (o *RubyUpstream) GetExtraHeader2Ok() (*string, bool)`

GetExtraHeader2Ok returns a tuple with the ExtraHeader2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader2

`func (o *RubyUpstream) SetExtraHeader2(v string)`

SetExtraHeader2 sets ExtraHeader2 field to given value.

### HasExtraHeader2

`func (o *RubyUpstream) HasExtraHeader2() bool`

HasExtraHeader2 returns a boolean if a field has been set.

### SetExtraHeader2Nil

`func (o *RubyUpstream) SetExtraHeader2Nil(b bool)`

 SetExtraHeader2Nil sets the value for ExtraHeader2 to be an explicit nil

### UnsetExtraHeader2
`func (o *RubyUpstream) UnsetExtraHeader2()`

UnsetExtraHeader2 ensures that no value is present for ExtraHeader2, not even an explicit nil
### GetExtraValue1

`func (o *RubyUpstream) GetExtraValue1() string`

GetExtraValue1 returns the ExtraValue1 field if non-nil, zero value otherwise.

### GetExtraValue1Ok

`func (o *RubyUpstream) GetExtraValue1Ok() (*string, bool)`

GetExtraValue1Ok returns a tuple with the ExtraValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue1

`func (o *RubyUpstream) SetExtraValue1(v string)`

SetExtraValue1 sets ExtraValue1 field to given value.

### HasExtraValue1

`func (o *RubyUpstream) HasExtraValue1() bool`

HasExtraValue1 returns a boolean if a field has been set.

### SetExtraValue1Nil

`func (o *RubyUpstream) SetExtraValue1Nil(b bool)`

 SetExtraValue1Nil sets the value for ExtraValue1 to be an explicit nil

### UnsetExtraValue1
`func (o *RubyUpstream) UnsetExtraValue1()`

UnsetExtraValue1 ensures that no value is present for ExtraValue1, not even an explicit nil
### GetExtraValue2

`func (o *RubyUpstream) GetExtraValue2() string`

GetExtraValue2 returns the ExtraValue2 field if non-nil, zero value otherwise.

### GetExtraValue2Ok

`func (o *RubyUpstream) GetExtraValue2Ok() (*string, bool)`

GetExtraValue2Ok returns a tuple with the ExtraValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue2

`func (o *RubyUpstream) SetExtraValue2(v string)`

SetExtraValue2 sets ExtraValue2 field to given value.

### HasExtraValue2

`func (o *RubyUpstream) HasExtraValue2() bool`

HasExtraValue2 returns a boolean if a field has been set.

### SetExtraValue2Nil

`func (o *RubyUpstream) SetExtraValue2Nil(b bool)`

 SetExtraValue2Nil sets the value for ExtraValue2 to be an explicit nil

### UnsetExtraValue2
`func (o *RubyUpstream) UnsetExtraValue2()`

UnsetExtraValue2 ensures that no value is present for ExtraValue2, not even an explicit nil
### GetHasFailedSignatureVerification

`func (o *RubyUpstream) GetHasFailedSignatureVerification() bool`

GetHasFailedSignatureVerification returns the HasFailedSignatureVerification field if non-nil, zero value otherwise.

### GetHasFailedSignatureVerificationOk

`func (o *RubyUpstream) GetHasFailedSignatureVerificationOk() (*bool, bool)`

GetHasFailedSignatureVerificationOk returns a tuple with the HasFailedSignatureVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFailedSignatureVerification

`func (o *RubyUpstream) SetHasFailedSignatureVerification(v bool)`

SetHasFailedSignatureVerification sets HasFailedSignatureVerification field to given value.

### HasHasFailedSignatureVerification

`func (o *RubyUpstream) HasHasFailedSignatureVerification() bool`

HasHasFailedSignatureVerification returns a boolean if a field has been set.

### GetIndexPackageCount

`func (o *RubyUpstream) GetIndexPackageCount() string`

GetIndexPackageCount returns the IndexPackageCount field if non-nil, zero value otherwise.

### GetIndexPackageCountOk

`func (o *RubyUpstream) GetIndexPackageCountOk() (*string, bool)`

GetIndexPackageCountOk returns a tuple with the IndexPackageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPackageCount

`func (o *RubyUpstream) SetIndexPackageCount(v string)`

SetIndexPackageCount sets IndexPackageCount field to given value.

### HasIndexPackageCount

`func (o *RubyUpstream) HasIndexPackageCount() bool`

HasIndexPackageCount returns a boolean if a field has been set.

### GetIndexStatus

`func (o *RubyUpstream) GetIndexStatus() string`

GetIndexStatus returns the IndexStatus field if non-nil, zero value otherwise.

### GetIndexStatusOk

`func (o *RubyUpstream) GetIndexStatusOk() (*string, bool)`

GetIndexStatusOk returns a tuple with the IndexStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexStatus

`func (o *RubyUpstream) SetIndexStatus(v string)`

SetIndexStatus sets IndexStatus field to given value.

### HasIndexStatus

`func (o *RubyUpstream) HasIndexStatus() bool`

HasIndexStatus returns a boolean if a field has been set.

### GetIsActive

`func (o *RubyUpstream) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RubyUpstream) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RubyUpstream) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RubyUpstream) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastIndexed

`func (o *RubyUpstream) GetLastIndexed() string`

GetLastIndexed returns the LastIndexed field if non-nil, zero value otherwise.

### GetLastIndexedOk

`func (o *RubyUpstream) GetLastIndexedOk() (*string, bool)`

GetLastIndexedOk returns a tuple with the LastIndexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndexed

`func (o *RubyUpstream) SetLastIndexed(v string)`

SetLastIndexed sets LastIndexed field to given value.

### HasLastIndexed

`func (o *RubyUpstream) HasLastIndexed() bool`

HasLastIndexed returns a boolean if a field has been set.

### GetMode

`func (o *RubyUpstream) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RubyUpstream) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RubyUpstream) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RubyUpstream) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *RubyUpstream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RubyUpstream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RubyUpstream) SetName(v string)`

SetName sets Name field to given value.


### GetPendingValidation

`func (o *RubyUpstream) GetPendingValidation() bool`

GetPendingValidation returns the PendingValidation field if non-nil, zero value otherwise.

### GetPendingValidationOk

`func (o *RubyUpstream) GetPendingValidationOk() (*bool, bool)`

GetPendingValidationOk returns a tuple with the PendingValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingValidation

`func (o *RubyUpstream) SetPendingValidation(v bool)`

SetPendingValidation sets PendingValidation field to given value.

### HasPendingValidation

`func (o *RubyUpstream) HasPendingValidation() bool`

HasPendingValidation returns a boolean if a field has been set.

### GetPriority

`func (o *RubyUpstream) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RubyUpstream) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RubyUpstream) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RubyUpstream) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSlugPerm

`func (o *RubyUpstream) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *RubyUpstream) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *RubyUpstream) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *RubyUpstream) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RubyUpstream) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RubyUpstream) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RubyUpstream) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RubyUpstream) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpstreamUrl

`func (o *RubyUpstream) GetUpstreamUrl() string`

GetUpstreamUrl returns the UpstreamUrl field if non-nil, zero value otherwise.

### GetUpstreamUrlOk

`func (o *RubyUpstream) GetUpstreamUrlOk() (*string, bool)`

GetUpstreamUrlOk returns a tuple with the UpstreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamUrl

`func (o *RubyUpstream) SetUpstreamUrl(v string)`

SetUpstreamUrl sets UpstreamUrl field to given value.


### GetVerifySsl

`func (o *RubyUpstream) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *RubyUpstream) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *RubyUpstream) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *RubyUpstream) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


