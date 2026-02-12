# CargoUpstream

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

### NewCargoUpstream

`func NewCargoUpstream(name string, upstreamUrl string, ) *CargoUpstream`

NewCargoUpstream instantiates a new CargoUpstream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCargoUpstreamWithDefaults

`func NewCargoUpstreamWithDefaults() *CargoUpstream`

NewCargoUpstreamWithDefaults instantiates a new CargoUpstream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *CargoUpstream) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *CargoUpstream) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *CargoUpstream) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *CargoUpstream) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAuthSecret

`func (o *CargoUpstream) GetAuthSecret() string`

GetAuthSecret returns the AuthSecret field if non-nil, zero value otherwise.

### GetAuthSecretOk

`func (o *CargoUpstream) GetAuthSecretOk() (*string, bool)`

GetAuthSecretOk returns a tuple with the AuthSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecret

`func (o *CargoUpstream) SetAuthSecret(v string)`

SetAuthSecret sets AuthSecret field to given value.

### HasAuthSecret

`func (o *CargoUpstream) HasAuthSecret() bool`

HasAuthSecret returns a boolean if a field has been set.

### SetAuthSecretNil

`func (o *CargoUpstream) SetAuthSecretNil(b bool)`

 SetAuthSecretNil sets the value for AuthSecret to be an explicit nil

### UnsetAuthSecret
`func (o *CargoUpstream) UnsetAuthSecret()`

UnsetAuthSecret ensures that no value is present for AuthSecret, not even an explicit nil
### GetAuthUsername

`func (o *CargoUpstream) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *CargoUpstream) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *CargoUpstream) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *CargoUpstream) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### SetAuthUsernameNil

`func (o *CargoUpstream) SetAuthUsernameNil(b bool)`

 SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil

### UnsetAuthUsername
`func (o *CargoUpstream) UnsetAuthUsername()`

UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
### GetAvailable

`func (o *CargoUpstream) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CargoUpstream) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CargoUpstream) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CargoUpstream) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCanReindex

`func (o *CargoUpstream) GetCanReindex() bool`

GetCanReindex returns the CanReindex field if non-nil, zero value otherwise.

### GetCanReindexOk

`func (o *CargoUpstream) GetCanReindexOk() (*bool, bool)`

GetCanReindexOk returns a tuple with the CanReindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReindex

`func (o *CargoUpstream) SetCanReindex(v bool)`

SetCanReindex sets CanReindex field to given value.

### HasCanReindex

`func (o *CargoUpstream) HasCanReindex() bool`

HasCanReindex returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CargoUpstream) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CargoUpstream) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CargoUpstream) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CargoUpstream) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisableReason

`func (o *CargoUpstream) GetDisableReason() string`

GetDisableReason returns the DisableReason field if non-nil, zero value otherwise.

### GetDisableReasonOk

`func (o *CargoUpstream) GetDisableReasonOk() (*string, bool)`

GetDisableReasonOk returns a tuple with the DisableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReason

`func (o *CargoUpstream) SetDisableReason(v string)`

SetDisableReason sets DisableReason field to given value.

### HasDisableReason

`func (o *CargoUpstream) HasDisableReason() bool`

HasDisableReason returns a boolean if a field has been set.

### GetDisableReasonText

`func (o *CargoUpstream) GetDisableReasonText() string`

GetDisableReasonText returns the DisableReasonText field if non-nil, zero value otherwise.

### GetDisableReasonTextOk

`func (o *CargoUpstream) GetDisableReasonTextOk() (*string, bool)`

GetDisableReasonTextOk returns a tuple with the DisableReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReasonText

`func (o *CargoUpstream) SetDisableReasonText(v string)`

SetDisableReasonText sets DisableReasonText field to given value.

### HasDisableReasonText

`func (o *CargoUpstream) HasDisableReasonText() bool`

HasDisableReasonText returns a boolean if a field has been set.

### GetExtraHeader1

`func (o *CargoUpstream) GetExtraHeader1() string`

GetExtraHeader1 returns the ExtraHeader1 field if non-nil, zero value otherwise.

### GetExtraHeader1Ok

`func (o *CargoUpstream) GetExtraHeader1Ok() (*string, bool)`

GetExtraHeader1Ok returns a tuple with the ExtraHeader1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader1

`func (o *CargoUpstream) SetExtraHeader1(v string)`

SetExtraHeader1 sets ExtraHeader1 field to given value.

### HasExtraHeader1

`func (o *CargoUpstream) HasExtraHeader1() bool`

HasExtraHeader1 returns a boolean if a field has been set.

### SetExtraHeader1Nil

`func (o *CargoUpstream) SetExtraHeader1Nil(b bool)`

 SetExtraHeader1Nil sets the value for ExtraHeader1 to be an explicit nil

### UnsetExtraHeader1
`func (o *CargoUpstream) UnsetExtraHeader1()`

UnsetExtraHeader1 ensures that no value is present for ExtraHeader1, not even an explicit nil
### GetExtraHeader2

`func (o *CargoUpstream) GetExtraHeader2() string`

GetExtraHeader2 returns the ExtraHeader2 field if non-nil, zero value otherwise.

### GetExtraHeader2Ok

`func (o *CargoUpstream) GetExtraHeader2Ok() (*string, bool)`

GetExtraHeader2Ok returns a tuple with the ExtraHeader2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader2

`func (o *CargoUpstream) SetExtraHeader2(v string)`

SetExtraHeader2 sets ExtraHeader2 field to given value.

### HasExtraHeader2

`func (o *CargoUpstream) HasExtraHeader2() bool`

HasExtraHeader2 returns a boolean if a field has been set.

### SetExtraHeader2Nil

`func (o *CargoUpstream) SetExtraHeader2Nil(b bool)`

 SetExtraHeader2Nil sets the value for ExtraHeader2 to be an explicit nil

### UnsetExtraHeader2
`func (o *CargoUpstream) UnsetExtraHeader2()`

UnsetExtraHeader2 ensures that no value is present for ExtraHeader2, not even an explicit nil
### GetExtraValue1

`func (o *CargoUpstream) GetExtraValue1() string`

GetExtraValue1 returns the ExtraValue1 field if non-nil, zero value otherwise.

### GetExtraValue1Ok

`func (o *CargoUpstream) GetExtraValue1Ok() (*string, bool)`

GetExtraValue1Ok returns a tuple with the ExtraValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue1

`func (o *CargoUpstream) SetExtraValue1(v string)`

SetExtraValue1 sets ExtraValue1 field to given value.

### HasExtraValue1

`func (o *CargoUpstream) HasExtraValue1() bool`

HasExtraValue1 returns a boolean if a field has been set.

### SetExtraValue1Nil

`func (o *CargoUpstream) SetExtraValue1Nil(b bool)`

 SetExtraValue1Nil sets the value for ExtraValue1 to be an explicit nil

### UnsetExtraValue1
`func (o *CargoUpstream) UnsetExtraValue1()`

UnsetExtraValue1 ensures that no value is present for ExtraValue1, not even an explicit nil
### GetExtraValue2

`func (o *CargoUpstream) GetExtraValue2() string`

GetExtraValue2 returns the ExtraValue2 field if non-nil, zero value otherwise.

### GetExtraValue2Ok

`func (o *CargoUpstream) GetExtraValue2Ok() (*string, bool)`

GetExtraValue2Ok returns a tuple with the ExtraValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue2

`func (o *CargoUpstream) SetExtraValue2(v string)`

SetExtraValue2 sets ExtraValue2 field to given value.

### HasExtraValue2

`func (o *CargoUpstream) HasExtraValue2() bool`

HasExtraValue2 returns a boolean if a field has been set.

### SetExtraValue2Nil

`func (o *CargoUpstream) SetExtraValue2Nil(b bool)`

 SetExtraValue2Nil sets the value for ExtraValue2 to be an explicit nil

### UnsetExtraValue2
`func (o *CargoUpstream) UnsetExtraValue2()`

UnsetExtraValue2 ensures that no value is present for ExtraValue2, not even an explicit nil
### GetHasFailedSignatureVerification

`func (o *CargoUpstream) GetHasFailedSignatureVerification() bool`

GetHasFailedSignatureVerification returns the HasFailedSignatureVerification field if non-nil, zero value otherwise.

### GetHasFailedSignatureVerificationOk

`func (o *CargoUpstream) GetHasFailedSignatureVerificationOk() (*bool, bool)`

GetHasFailedSignatureVerificationOk returns a tuple with the HasFailedSignatureVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFailedSignatureVerification

`func (o *CargoUpstream) SetHasFailedSignatureVerification(v bool)`

SetHasFailedSignatureVerification sets HasFailedSignatureVerification field to given value.

### HasHasFailedSignatureVerification

`func (o *CargoUpstream) HasHasFailedSignatureVerification() bool`

HasHasFailedSignatureVerification returns a boolean if a field has been set.

### GetIndexPackageCount

`func (o *CargoUpstream) GetIndexPackageCount() string`

GetIndexPackageCount returns the IndexPackageCount field if non-nil, zero value otherwise.

### GetIndexPackageCountOk

`func (o *CargoUpstream) GetIndexPackageCountOk() (*string, bool)`

GetIndexPackageCountOk returns a tuple with the IndexPackageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPackageCount

`func (o *CargoUpstream) SetIndexPackageCount(v string)`

SetIndexPackageCount sets IndexPackageCount field to given value.

### HasIndexPackageCount

`func (o *CargoUpstream) HasIndexPackageCount() bool`

HasIndexPackageCount returns a boolean if a field has been set.

### GetIndexStatus

`func (o *CargoUpstream) GetIndexStatus() string`

GetIndexStatus returns the IndexStatus field if non-nil, zero value otherwise.

### GetIndexStatusOk

`func (o *CargoUpstream) GetIndexStatusOk() (*string, bool)`

GetIndexStatusOk returns a tuple with the IndexStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexStatus

`func (o *CargoUpstream) SetIndexStatus(v string)`

SetIndexStatus sets IndexStatus field to given value.

### HasIndexStatus

`func (o *CargoUpstream) HasIndexStatus() bool`

HasIndexStatus returns a boolean if a field has been set.

### GetIsActive

`func (o *CargoUpstream) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CargoUpstream) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CargoUpstream) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CargoUpstream) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastIndexed

`func (o *CargoUpstream) GetLastIndexed() string`

GetLastIndexed returns the LastIndexed field if non-nil, zero value otherwise.

### GetLastIndexedOk

`func (o *CargoUpstream) GetLastIndexedOk() (*string, bool)`

GetLastIndexedOk returns a tuple with the LastIndexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndexed

`func (o *CargoUpstream) SetLastIndexed(v string)`

SetLastIndexed sets LastIndexed field to given value.

### HasLastIndexed

`func (o *CargoUpstream) HasLastIndexed() bool`

HasLastIndexed returns a boolean if a field has been set.

### GetMode

`func (o *CargoUpstream) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CargoUpstream) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CargoUpstream) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CargoUpstream) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *CargoUpstream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CargoUpstream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CargoUpstream) SetName(v string)`

SetName sets Name field to given value.


### GetPendingValidation

`func (o *CargoUpstream) GetPendingValidation() bool`

GetPendingValidation returns the PendingValidation field if non-nil, zero value otherwise.

### GetPendingValidationOk

`func (o *CargoUpstream) GetPendingValidationOk() (*bool, bool)`

GetPendingValidationOk returns a tuple with the PendingValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingValidation

`func (o *CargoUpstream) SetPendingValidation(v bool)`

SetPendingValidation sets PendingValidation field to given value.

### HasPendingValidation

`func (o *CargoUpstream) HasPendingValidation() bool`

HasPendingValidation returns a boolean if a field has been set.

### GetPriority

`func (o *CargoUpstream) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CargoUpstream) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CargoUpstream) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CargoUpstream) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSlugPerm

`func (o *CargoUpstream) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *CargoUpstream) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *CargoUpstream) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *CargoUpstream) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CargoUpstream) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CargoUpstream) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CargoUpstream) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CargoUpstream) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpstreamUrl

`func (o *CargoUpstream) GetUpstreamUrl() string`

GetUpstreamUrl returns the UpstreamUrl field if non-nil, zero value otherwise.

### GetUpstreamUrlOk

`func (o *CargoUpstream) GetUpstreamUrlOk() (*string, bool)`

GetUpstreamUrlOk returns a tuple with the UpstreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamUrl

`func (o *CargoUpstream) SetUpstreamUrl(v string)`

SetUpstreamUrl sets UpstreamUrl field to given value.


### GetVerifySsl

`func (o *CargoUpstream) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *CargoUpstream) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *CargoUpstream) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *CargoUpstream) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


