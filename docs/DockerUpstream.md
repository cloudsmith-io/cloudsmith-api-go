# DockerUpstream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | Pointer to **string** | The authentication mode to use when accessing this upstream.  | [optional] [default to "None"]
**AuthSecret** | Pointer to **NullableString** | Secret to provide with requests to upstream. | [optional] 
**AuthUsername** | Pointer to **NullableString** | Username to provide with requests to upstream. | [optional] 
**Available** | Pointer to **bool** | Whether the upstream is available for use. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The datetime the upstream source was created. | [optional] [readonly] 
**DisableReason** | Pointer to **string** |  | [optional] [readonly] [default to "N/A"]
**DisableReasonText** | Pointer to **string** | Human-readable explanation of why this upstream is disabled | [optional] [readonly] 
**ExtraHeader1** | Pointer to **NullableString** | The key for extra header #1 to send to upstream. | [optional] 
**ExtraHeader2** | Pointer to **NullableString** | The key for extra header #2 to send to upstream. | [optional] 
**ExtraValue1** | Pointer to **NullableString** | The value for extra header #1 to send to upstream. This is stored as plaintext, and is NOT encrypted. | [optional] 
**ExtraValue2** | Pointer to **NullableString** | The value for extra header #2 to send to upstream. This is stored as plaintext, and is NOT encrypted. | [optional] 
**HasFailedSignatureVerification** | Pointer to **bool** | Whether the upstream has failed signature verification. | [optional] [readonly] 
**IsActive** | Pointer to **bool** | Whether or not this upstream is active and ready for requests. | [optional] 
**Mode** | Pointer to **string** | The mode that this upstream should operate in. Upstream sources can be used to proxy resolved packages, as well as operate in a proxy/cache or cache only mode. | [optional] [default to "Proxy Only"]
**Name** | **string** | A descriptive name for this upstream source. A shortened version of this name will be used for tagging cached packages retrieved from this upstream. | 
**PendingValidation** | Pointer to **bool** | When true, this upstream source is pending validation. | [optional] [readonly] 
**Priority** | Pointer to **int64** | Upstream sources are selected for resolving requests by sequential order (1..n), followed by creation date. | [optional] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpstreamUrl** | **string** | The URL for this upstream source. This must be a fully qualified URL including any path elements required to reach the root of the repository.  | 
**VerifySsl** | Pointer to **bool** | If enabled, SSL certificates are verified when requests are made to this upstream. It&#39;s recommended to leave this enabled for all public sources to help mitigate Man-In-The-Middle (MITM) attacks. Please note this only applies to HTTPS upstreams. | [optional] 

## Methods

### NewDockerUpstream

`func NewDockerUpstream(name string, upstreamUrl string, ) *DockerUpstream`

NewDockerUpstream instantiates a new DockerUpstream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerUpstreamWithDefaults

`func NewDockerUpstreamWithDefaults() *DockerUpstream`

NewDockerUpstreamWithDefaults instantiates a new DockerUpstream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *DockerUpstream) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *DockerUpstream) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *DockerUpstream) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *DockerUpstream) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAuthSecret

`func (o *DockerUpstream) GetAuthSecret() string`

GetAuthSecret returns the AuthSecret field if non-nil, zero value otherwise.

### GetAuthSecretOk

`func (o *DockerUpstream) GetAuthSecretOk() (*string, bool)`

GetAuthSecretOk returns a tuple with the AuthSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecret

`func (o *DockerUpstream) SetAuthSecret(v string)`

SetAuthSecret sets AuthSecret field to given value.

### HasAuthSecret

`func (o *DockerUpstream) HasAuthSecret() bool`

HasAuthSecret returns a boolean if a field has been set.

### SetAuthSecretNil

`func (o *DockerUpstream) SetAuthSecretNil(b bool)`

 SetAuthSecretNil sets the value for AuthSecret to be an explicit nil

### UnsetAuthSecret
`func (o *DockerUpstream) UnsetAuthSecret()`

UnsetAuthSecret ensures that no value is present for AuthSecret, not even an explicit nil
### GetAuthUsername

`func (o *DockerUpstream) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *DockerUpstream) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *DockerUpstream) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *DockerUpstream) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### SetAuthUsernameNil

`func (o *DockerUpstream) SetAuthUsernameNil(b bool)`

 SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil

### UnsetAuthUsername
`func (o *DockerUpstream) UnsetAuthUsername()`

UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
### GetAvailable

`func (o *DockerUpstream) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *DockerUpstream) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *DockerUpstream) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *DockerUpstream) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DockerUpstream) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DockerUpstream) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DockerUpstream) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DockerUpstream) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisableReason

`func (o *DockerUpstream) GetDisableReason() string`

GetDisableReason returns the DisableReason field if non-nil, zero value otherwise.

### GetDisableReasonOk

`func (o *DockerUpstream) GetDisableReasonOk() (*string, bool)`

GetDisableReasonOk returns a tuple with the DisableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReason

`func (o *DockerUpstream) SetDisableReason(v string)`

SetDisableReason sets DisableReason field to given value.

### HasDisableReason

`func (o *DockerUpstream) HasDisableReason() bool`

HasDisableReason returns a boolean if a field has been set.

### GetDisableReasonText

`func (o *DockerUpstream) GetDisableReasonText() string`

GetDisableReasonText returns the DisableReasonText field if non-nil, zero value otherwise.

### GetDisableReasonTextOk

`func (o *DockerUpstream) GetDisableReasonTextOk() (*string, bool)`

GetDisableReasonTextOk returns a tuple with the DisableReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReasonText

`func (o *DockerUpstream) SetDisableReasonText(v string)`

SetDisableReasonText sets DisableReasonText field to given value.

### HasDisableReasonText

`func (o *DockerUpstream) HasDisableReasonText() bool`

HasDisableReasonText returns a boolean if a field has been set.

### GetExtraHeader1

`func (o *DockerUpstream) GetExtraHeader1() string`

GetExtraHeader1 returns the ExtraHeader1 field if non-nil, zero value otherwise.

### GetExtraHeader1Ok

`func (o *DockerUpstream) GetExtraHeader1Ok() (*string, bool)`

GetExtraHeader1Ok returns a tuple with the ExtraHeader1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader1

`func (o *DockerUpstream) SetExtraHeader1(v string)`

SetExtraHeader1 sets ExtraHeader1 field to given value.

### HasExtraHeader1

`func (o *DockerUpstream) HasExtraHeader1() bool`

HasExtraHeader1 returns a boolean if a field has been set.

### SetExtraHeader1Nil

`func (o *DockerUpstream) SetExtraHeader1Nil(b bool)`

 SetExtraHeader1Nil sets the value for ExtraHeader1 to be an explicit nil

### UnsetExtraHeader1
`func (o *DockerUpstream) UnsetExtraHeader1()`

UnsetExtraHeader1 ensures that no value is present for ExtraHeader1, not even an explicit nil
### GetExtraHeader2

`func (o *DockerUpstream) GetExtraHeader2() string`

GetExtraHeader2 returns the ExtraHeader2 field if non-nil, zero value otherwise.

### GetExtraHeader2Ok

`func (o *DockerUpstream) GetExtraHeader2Ok() (*string, bool)`

GetExtraHeader2Ok returns a tuple with the ExtraHeader2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader2

`func (o *DockerUpstream) SetExtraHeader2(v string)`

SetExtraHeader2 sets ExtraHeader2 field to given value.

### HasExtraHeader2

`func (o *DockerUpstream) HasExtraHeader2() bool`

HasExtraHeader2 returns a boolean if a field has been set.

### SetExtraHeader2Nil

`func (o *DockerUpstream) SetExtraHeader2Nil(b bool)`

 SetExtraHeader2Nil sets the value for ExtraHeader2 to be an explicit nil

### UnsetExtraHeader2
`func (o *DockerUpstream) UnsetExtraHeader2()`

UnsetExtraHeader2 ensures that no value is present for ExtraHeader2, not even an explicit nil
### GetExtraValue1

`func (o *DockerUpstream) GetExtraValue1() string`

GetExtraValue1 returns the ExtraValue1 field if non-nil, zero value otherwise.

### GetExtraValue1Ok

`func (o *DockerUpstream) GetExtraValue1Ok() (*string, bool)`

GetExtraValue1Ok returns a tuple with the ExtraValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue1

`func (o *DockerUpstream) SetExtraValue1(v string)`

SetExtraValue1 sets ExtraValue1 field to given value.

### HasExtraValue1

`func (o *DockerUpstream) HasExtraValue1() bool`

HasExtraValue1 returns a boolean if a field has been set.

### SetExtraValue1Nil

`func (o *DockerUpstream) SetExtraValue1Nil(b bool)`

 SetExtraValue1Nil sets the value for ExtraValue1 to be an explicit nil

### UnsetExtraValue1
`func (o *DockerUpstream) UnsetExtraValue1()`

UnsetExtraValue1 ensures that no value is present for ExtraValue1, not even an explicit nil
### GetExtraValue2

`func (o *DockerUpstream) GetExtraValue2() string`

GetExtraValue2 returns the ExtraValue2 field if non-nil, zero value otherwise.

### GetExtraValue2Ok

`func (o *DockerUpstream) GetExtraValue2Ok() (*string, bool)`

GetExtraValue2Ok returns a tuple with the ExtraValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue2

`func (o *DockerUpstream) SetExtraValue2(v string)`

SetExtraValue2 sets ExtraValue2 field to given value.

### HasExtraValue2

`func (o *DockerUpstream) HasExtraValue2() bool`

HasExtraValue2 returns a boolean if a field has been set.

### SetExtraValue2Nil

`func (o *DockerUpstream) SetExtraValue2Nil(b bool)`

 SetExtraValue2Nil sets the value for ExtraValue2 to be an explicit nil

### UnsetExtraValue2
`func (o *DockerUpstream) UnsetExtraValue2()`

UnsetExtraValue2 ensures that no value is present for ExtraValue2, not even an explicit nil
### GetHasFailedSignatureVerification

`func (o *DockerUpstream) GetHasFailedSignatureVerification() bool`

GetHasFailedSignatureVerification returns the HasFailedSignatureVerification field if non-nil, zero value otherwise.

### GetHasFailedSignatureVerificationOk

`func (o *DockerUpstream) GetHasFailedSignatureVerificationOk() (*bool, bool)`

GetHasFailedSignatureVerificationOk returns a tuple with the HasFailedSignatureVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFailedSignatureVerification

`func (o *DockerUpstream) SetHasFailedSignatureVerification(v bool)`

SetHasFailedSignatureVerification sets HasFailedSignatureVerification field to given value.

### HasHasFailedSignatureVerification

`func (o *DockerUpstream) HasHasFailedSignatureVerification() bool`

HasHasFailedSignatureVerification returns a boolean if a field has been set.

### GetIsActive

`func (o *DockerUpstream) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DockerUpstream) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DockerUpstream) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *DockerUpstream) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetMode

`func (o *DockerUpstream) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DockerUpstream) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DockerUpstream) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DockerUpstream) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *DockerUpstream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DockerUpstream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DockerUpstream) SetName(v string)`

SetName sets Name field to given value.


### GetPendingValidation

`func (o *DockerUpstream) GetPendingValidation() bool`

GetPendingValidation returns the PendingValidation field if non-nil, zero value otherwise.

### GetPendingValidationOk

`func (o *DockerUpstream) GetPendingValidationOk() (*bool, bool)`

GetPendingValidationOk returns a tuple with the PendingValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingValidation

`func (o *DockerUpstream) SetPendingValidation(v bool)`

SetPendingValidation sets PendingValidation field to given value.

### HasPendingValidation

`func (o *DockerUpstream) HasPendingValidation() bool`

HasPendingValidation returns a boolean if a field has been set.

### GetPriority

`func (o *DockerUpstream) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DockerUpstream) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DockerUpstream) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DockerUpstream) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSlugPerm

`func (o *DockerUpstream) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *DockerUpstream) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *DockerUpstream) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *DockerUpstream) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DockerUpstream) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DockerUpstream) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DockerUpstream) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DockerUpstream) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpstreamUrl

`func (o *DockerUpstream) GetUpstreamUrl() string`

GetUpstreamUrl returns the UpstreamUrl field if non-nil, zero value otherwise.

### GetUpstreamUrlOk

`func (o *DockerUpstream) GetUpstreamUrlOk() (*string, bool)`

GetUpstreamUrlOk returns a tuple with the UpstreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamUrl

`func (o *DockerUpstream) SetUpstreamUrl(v string)`

SetUpstreamUrl sets UpstreamUrl field to given value.


### GetVerifySsl

`func (o *DockerUpstream) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *DockerUpstream) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *DockerUpstream) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *DockerUpstream) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


