# PythonUpstream

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
**TrustLevel** | Pointer to **string** | Trust level allows for control of the visibility of upstream artifacts to native package managers. Where supported by formats, the default level (untrusted) is recommended for all upstreams, and helps to safeguard against common dependency confusion attack vectors. | [optional] [default to "Trusted"]
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpstreamUrl** | **string** | The URL for this upstream source. This must be a fully qualified URL including any path elements required to reach the root of the repository.  | 
**VerifySsl** | Pointer to **bool** | If enabled, SSL certificates are verified when requests are made to this upstream. It&#39;s recommended to leave this enabled for all public sources to help mitigate Man-In-The-Middle (MITM) attacks. Please note this only applies to HTTPS upstreams. | [optional] 

## Methods

### NewPythonUpstream

`func NewPythonUpstream(name string, upstreamUrl string, ) *PythonUpstream`

NewPythonUpstream instantiates a new PythonUpstream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPythonUpstreamWithDefaults

`func NewPythonUpstreamWithDefaults() *PythonUpstream`

NewPythonUpstreamWithDefaults instantiates a new PythonUpstream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *PythonUpstream) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *PythonUpstream) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *PythonUpstream) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *PythonUpstream) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAuthSecret

`func (o *PythonUpstream) GetAuthSecret() string`

GetAuthSecret returns the AuthSecret field if non-nil, zero value otherwise.

### GetAuthSecretOk

`func (o *PythonUpstream) GetAuthSecretOk() (*string, bool)`

GetAuthSecretOk returns a tuple with the AuthSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecret

`func (o *PythonUpstream) SetAuthSecret(v string)`

SetAuthSecret sets AuthSecret field to given value.

### HasAuthSecret

`func (o *PythonUpstream) HasAuthSecret() bool`

HasAuthSecret returns a boolean if a field has been set.

### SetAuthSecretNil

`func (o *PythonUpstream) SetAuthSecretNil(b bool)`

 SetAuthSecretNil sets the value for AuthSecret to be an explicit nil

### UnsetAuthSecret
`func (o *PythonUpstream) UnsetAuthSecret()`

UnsetAuthSecret ensures that no value is present for AuthSecret, not even an explicit nil
### GetAuthUsername

`func (o *PythonUpstream) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *PythonUpstream) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *PythonUpstream) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *PythonUpstream) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### SetAuthUsernameNil

`func (o *PythonUpstream) SetAuthUsernameNil(b bool)`

 SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil

### UnsetAuthUsername
`func (o *PythonUpstream) UnsetAuthUsername()`

UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
### GetAvailable

`func (o *PythonUpstream) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *PythonUpstream) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *PythonUpstream) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *PythonUpstream) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCanReindex

`func (o *PythonUpstream) GetCanReindex() bool`

GetCanReindex returns the CanReindex field if non-nil, zero value otherwise.

### GetCanReindexOk

`func (o *PythonUpstream) GetCanReindexOk() (*bool, bool)`

GetCanReindexOk returns a tuple with the CanReindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReindex

`func (o *PythonUpstream) SetCanReindex(v bool)`

SetCanReindex sets CanReindex field to given value.

### HasCanReindex

`func (o *PythonUpstream) HasCanReindex() bool`

HasCanReindex returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PythonUpstream) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PythonUpstream) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PythonUpstream) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PythonUpstream) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisableReason

`func (o *PythonUpstream) GetDisableReason() string`

GetDisableReason returns the DisableReason field if non-nil, zero value otherwise.

### GetDisableReasonOk

`func (o *PythonUpstream) GetDisableReasonOk() (*string, bool)`

GetDisableReasonOk returns a tuple with the DisableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReason

`func (o *PythonUpstream) SetDisableReason(v string)`

SetDisableReason sets DisableReason field to given value.

### HasDisableReason

`func (o *PythonUpstream) HasDisableReason() bool`

HasDisableReason returns a boolean if a field has been set.

### GetDisableReasonText

`func (o *PythonUpstream) GetDisableReasonText() string`

GetDisableReasonText returns the DisableReasonText field if non-nil, zero value otherwise.

### GetDisableReasonTextOk

`func (o *PythonUpstream) GetDisableReasonTextOk() (*string, bool)`

GetDisableReasonTextOk returns a tuple with the DisableReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReasonText

`func (o *PythonUpstream) SetDisableReasonText(v string)`

SetDisableReasonText sets DisableReasonText field to given value.

### HasDisableReasonText

`func (o *PythonUpstream) HasDisableReasonText() bool`

HasDisableReasonText returns a boolean if a field has been set.

### GetExtraHeader1

`func (o *PythonUpstream) GetExtraHeader1() string`

GetExtraHeader1 returns the ExtraHeader1 field if non-nil, zero value otherwise.

### GetExtraHeader1Ok

`func (o *PythonUpstream) GetExtraHeader1Ok() (*string, bool)`

GetExtraHeader1Ok returns a tuple with the ExtraHeader1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader1

`func (o *PythonUpstream) SetExtraHeader1(v string)`

SetExtraHeader1 sets ExtraHeader1 field to given value.

### HasExtraHeader1

`func (o *PythonUpstream) HasExtraHeader1() bool`

HasExtraHeader1 returns a boolean if a field has been set.

### SetExtraHeader1Nil

`func (o *PythonUpstream) SetExtraHeader1Nil(b bool)`

 SetExtraHeader1Nil sets the value for ExtraHeader1 to be an explicit nil

### UnsetExtraHeader1
`func (o *PythonUpstream) UnsetExtraHeader1()`

UnsetExtraHeader1 ensures that no value is present for ExtraHeader1, not even an explicit nil
### GetExtraHeader2

`func (o *PythonUpstream) GetExtraHeader2() string`

GetExtraHeader2 returns the ExtraHeader2 field if non-nil, zero value otherwise.

### GetExtraHeader2Ok

`func (o *PythonUpstream) GetExtraHeader2Ok() (*string, bool)`

GetExtraHeader2Ok returns a tuple with the ExtraHeader2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader2

`func (o *PythonUpstream) SetExtraHeader2(v string)`

SetExtraHeader2 sets ExtraHeader2 field to given value.

### HasExtraHeader2

`func (o *PythonUpstream) HasExtraHeader2() bool`

HasExtraHeader2 returns a boolean if a field has been set.

### SetExtraHeader2Nil

`func (o *PythonUpstream) SetExtraHeader2Nil(b bool)`

 SetExtraHeader2Nil sets the value for ExtraHeader2 to be an explicit nil

### UnsetExtraHeader2
`func (o *PythonUpstream) UnsetExtraHeader2()`

UnsetExtraHeader2 ensures that no value is present for ExtraHeader2, not even an explicit nil
### GetExtraValue1

`func (o *PythonUpstream) GetExtraValue1() string`

GetExtraValue1 returns the ExtraValue1 field if non-nil, zero value otherwise.

### GetExtraValue1Ok

`func (o *PythonUpstream) GetExtraValue1Ok() (*string, bool)`

GetExtraValue1Ok returns a tuple with the ExtraValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue1

`func (o *PythonUpstream) SetExtraValue1(v string)`

SetExtraValue1 sets ExtraValue1 field to given value.

### HasExtraValue1

`func (o *PythonUpstream) HasExtraValue1() bool`

HasExtraValue1 returns a boolean if a field has been set.

### SetExtraValue1Nil

`func (o *PythonUpstream) SetExtraValue1Nil(b bool)`

 SetExtraValue1Nil sets the value for ExtraValue1 to be an explicit nil

### UnsetExtraValue1
`func (o *PythonUpstream) UnsetExtraValue1()`

UnsetExtraValue1 ensures that no value is present for ExtraValue1, not even an explicit nil
### GetExtraValue2

`func (o *PythonUpstream) GetExtraValue2() string`

GetExtraValue2 returns the ExtraValue2 field if non-nil, zero value otherwise.

### GetExtraValue2Ok

`func (o *PythonUpstream) GetExtraValue2Ok() (*string, bool)`

GetExtraValue2Ok returns a tuple with the ExtraValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue2

`func (o *PythonUpstream) SetExtraValue2(v string)`

SetExtraValue2 sets ExtraValue2 field to given value.

### HasExtraValue2

`func (o *PythonUpstream) HasExtraValue2() bool`

HasExtraValue2 returns a boolean if a field has been set.

### SetExtraValue2Nil

`func (o *PythonUpstream) SetExtraValue2Nil(b bool)`

 SetExtraValue2Nil sets the value for ExtraValue2 to be an explicit nil

### UnsetExtraValue2
`func (o *PythonUpstream) UnsetExtraValue2()`

UnsetExtraValue2 ensures that no value is present for ExtraValue2, not even an explicit nil
### GetHasFailedSignatureVerification

`func (o *PythonUpstream) GetHasFailedSignatureVerification() bool`

GetHasFailedSignatureVerification returns the HasFailedSignatureVerification field if non-nil, zero value otherwise.

### GetHasFailedSignatureVerificationOk

`func (o *PythonUpstream) GetHasFailedSignatureVerificationOk() (*bool, bool)`

GetHasFailedSignatureVerificationOk returns a tuple with the HasFailedSignatureVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFailedSignatureVerification

`func (o *PythonUpstream) SetHasFailedSignatureVerification(v bool)`

SetHasFailedSignatureVerification sets HasFailedSignatureVerification field to given value.

### HasHasFailedSignatureVerification

`func (o *PythonUpstream) HasHasFailedSignatureVerification() bool`

HasHasFailedSignatureVerification returns a boolean if a field has been set.

### GetIndexPackageCount

`func (o *PythonUpstream) GetIndexPackageCount() string`

GetIndexPackageCount returns the IndexPackageCount field if non-nil, zero value otherwise.

### GetIndexPackageCountOk

`func (o *PythonUpstream) GetIndexPackageCountOk() (*string, bool)`

GetIndexPackageCountOk returns a tuple with the IndexPackageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPackageCount

`func (o *PythonUpstream) SetIndexPackageCount(v string)`

SetIndexPackageCount sets IndexPackageCount field to given value.

### HasIndexPackageCount

`func (o *PythonUpstream) HasIndexPackageCount() bool`

HasIndexPackageCount returns a boolean if a field has been set.

### GetIndexStatus

`func (o *PythonUpstream) GetIndexStatus() string`

GetIndexStatus returns the IndexStatus field if non-nil, zero value otherwise.

### GetIndexStatusOk

`func (o *PythonUpstream) GetIndexStatusOk() (*string, bool)`

GetIndexStatusOk returns a tuple with the IndexStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexStatus

`func (o *PythonUpstream) SetIndexStatus(v string)`

SetIndexStatus sets IndexStatus field to given value.

### HasIndexStatus

`func (o *PythonUpstream) HasIndexStatus() bool`

HasIndexStatus returns a boolean if a field has been set.

### GetIsActive

`func (o *PythonUpstream) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PythonUpstream) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PythonUpstream) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PythonUpstream) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastIndexed

`func (o *PythonUpstream) GetLastIndexed() string`

GetLastIndexed returns the LastIndexed field if non-nil, zero value otherwise.

### GetLastIndexedOk

`func (o *PythonUpstream) GetLastIndexedOk() (*string, bool)`

GetLastIndexedOk returns a tuple with the LastIndexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndexed

`func (o *PythonUpstream) SetLastIndexed(v string)`

SetLastIndexed sets LastIndexed field to given value.

### HasLastIndexed

`func (o *PythonUpstream) HasLastIndexed() bool`

HasLastIndexed returns a boolean if a field has been set.

### GetMode

`func (o *PythonUpstream) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PythonUpstream) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PythonUpstream) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PythonUpstream) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *PythonUpstream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PythonUpstream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PythonUpstream) SetName(v string)`

SetName sets Name field to given value.


### GetPendingValidation

`func (o *PythonUpstream) GetPendingValidation() bool`

GetPendingValidation returns the PendingValidation field if non-nil, zero value otherwise.

### GetPendingValidationOk

`func (o *PythonUpstream) GetPendingValidationOk() (*bool, bool)`

GetPendingValidationOk returns a tuple with the PendingValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingValidation

`func (o *PythonUpstream) SetPendingValidation(v bool)`

SetPendingValidation sets PendingValidation field to given value.

### HasPendingValidation

`func (o *PythonUpstream) HasPendingValidation() bool`

HasPendingValidation returns a boolean if a field has been set.

### GetPriority

`func (o *PythonUpstream) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PythonUpstream) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PythonUpstream) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PythonUpstream) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSlugPerm

`func (o *PythonUpstream) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *PythonUpstream) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *PythonUpstream) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *PythonUpstream) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetTrustLevel

`func (o *PythonUpstream) GetTrustLevel() string`

GetTrustLevel returns the TrustLevel field if non-nil, zero value otherwise.

### GetTrustLevelOk

`func (o *PythonUpstream) GetTrustLevelOk() (*string, bool)`

GetTrustLevelOk returns a tuple with the TrustLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustLevel

`func (o *PythonUpstream) SetTrustLevel(v string)`

SetTrustLevel sets TrustLevel field to given value.

### HasTrustLevel

`func (o *PythonUpstream) HasTrustLevel() bool`

HasTrustLevel returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PythonUpstream) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PythonUpstream) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PythonUpstream) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PythonUpstream) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpstreamUrl

`func (o *PythonUpstream) GetUpstreamUrl() string`

GetUpstreamUrl returns the UpstreamUrl field if non-nil, zero value otherwise.

### GetUpstreamUrlOk

`func (o *PythonUpstream) GetUpstreamUrlOk() (*string, bool)`

GetUpstreamUrlOk returns a tuple with the UpstreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamUrl

`func (o *PythonUpstream) SetUpstreamUrl(v string)`

SetUpstreamUrl sets UpstreamUrl field to given value.


### GetVerifySsl

`func (o *PythonUpstream) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *PythonUpstream) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *PythonUpstream) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *PythonUpstream) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


