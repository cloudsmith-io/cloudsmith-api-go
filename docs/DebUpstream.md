# DebUpstream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | Pointer to **string** | The authentication mode to use when accessing this upstream.  | [optional] [default to "None"]
**AuthSecret** | Pointer to **NullableString** | Secret to provide with requests to upstream. | [optional] 
**AuthUsername** | Pointer to **NullableString** | Username to provide with requests to upstream. | [optional] 
**Component** | Pointer to **string** | The component to fetch from the upstream | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime the upstream source was created. | [optional] [readonly] 
**DisableReason** | Pointer to **string** |  | [optional] [readonly] [default to "N/A"]
**DistroVersions** | **[]string** | The distribution version that packages found on this upstream could be associated with. | 
**ExtraHeader1** | Pointer to **NullableString** | The key for extra header #1 to send to upstream. | [optional] 
**ExtraHeader2** | Pointer to **NullableString** | The key for extra header #2 to send to upstream. | [optional] 
**ExtraValue1** | Pointer to **NullableString** | The value for extra header #1 to send to upstream. This is stored as plaintext, and is NOT encrypted. | [optional] 
**ExtraValue2** | Pointer to **NullableString** | The value for extra header #2 to send to upstream. This is stored as plaintext, and is NOT encrypted. | [optional] 
**GpgKeyInline** | Pointer to **NullableString** | A public GPG key to associate with packages found on this upstream. When using the Cloudsmith setup script, this GPG key will be automatically imported on your deployment machines to allow upstream packages to validate and install. | [optional] 
**GpgKeyUrl** | Pointer to **NullableString** | When provided, Cloudsmith will fetch, validate, and associate a public GPG key found at the provided URL. When using the Cloudsmith setup script, this GPG key will be automatically imported on your deployment machines to allow upstream packages to validate and install. | [optional] 
**GpgVerification** | Pointer to **string** | The GPG signature verification mode for this upstream. | [optional] [default to "Allow All"]
**IncludeSources** | Pointer to **bool** | When true, source packages will be available from this upstream. | [optional] 
**IsActive** | Pointer to **bool** | Whether or not this upstream is active and ready for requests. | [optional] 
**Mode** | Pointer to **string** | The mode that this upstream should operate in. Upstream sources can be used to proxy resolved packages, as well as operate in a proxy/cache or cache only mode. | [optional] [default to "Proxy Only"]
**Name** | **string** | A descriptive name for this upstream source. A shortened version of this name will be used for tagging cached packages retrieved from this upstream. | 
**PendingValidation** | Pointer to **bool** | When true, this upstream source is pending validation. | [optional] [readonly] 
**Priority** | Pointer to **int64** | Upstream sources are selected for resolving requests by sequential order (1..n), followed by creation date. | [optional] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpstreamDistribution** | Pointer to **NullableString** | The distribution to fetch from the upstream | [optional] 
**UpstreamUrl** | **string** | The URL for this upstream source. This must be a fully qualified URL including any path elements required to reach the root of the repository.  | 
**VerificationStatus** | Pointer to **string** | The signature verification status for this upstream. | [optional] [readonly] [default to "Unknown"]
**VerifySsl** | Pointer to **bool** | If enabled, SSL certificates are verified when requests are made to this upstream. It&#39;s recommended to leave this enabled for all public sources to help mitigate Man-In-The-Middle (MITM) attacks. Please note this only applies to HTTPS upstreams. | [optional] 

## Methods

### NewDebUpstream

`func NewDebUpstream(distroVersions []string, name string, upstreamUrl string, ) *DebUpstream`

NewDebUpstream instantiates a new DebUpstream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebUpstreamWithDefaults

`func NewDebUpstreamWithDefaults() *DebUpstream`

NewDebUpstreamWithDefaults instantiates a new DebUpstream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *DebUpstream) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *DebUpstream) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *DebUpstream) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *DebUpstream) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAuthSecret

`func (o *DebUpstream) GetAuthSecret() string`

GetAuthSecret returns the AuthSecret field if non-nil, zero value otherwise.

### GetAuthSecretOk

`func (o *DebUpstream) GetAuthSecretOk() (*string, bool)`

GetAuthSecretOk returns a tuple with the AuthSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecret

`func (o *DebUpstream) SetAuthSecret(v string)`

SetAuthSecret sets AuthSecret field to given value.

### HasAuthSecret

`func (o *DebUpstream) HasAuthSecret() bool`

HasAuthSecret returns a boolean if a field has been set.

### SetAuthSecretNil

`func (o *DebUpstream) SetAuthSecretNil(b bool)`

 SetAuthSecretNil sets the value for AuthSecret to be an explicit nil

### UnsetAuthSecret
`func (o *DebUpstream) UnsetAuthSecret()`

UnsetAuthSecret ensures that no value is present for AuthSecret, not even an explicit nil
### GetAuthUsername

`func (o *DebUpstream) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *DebUpstream) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *DebUpstream) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *DebUpstream) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### SetAuthUsernameNil

`func (o *DebUpstream) SetAuthUsernameNil(b bool)`

 SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil

### UnsetAuthUsername
`func (o *DebUpstream) UnsetAuthUsername()`

UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
### GetComponent

`func (o *DebUpstream) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *DebUpstream) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *DebUpstream) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *DebUpstream) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DebUpstream) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DebUpstream) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DebUpstream) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DebUpstream) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisableReason

`func (o *DebUpstream) GetDisableReason() string`

GetDisableReason returns the DisableReason field if non-nil, zero value otherwise.

### GetDisableReasonOk

`func (o *DebUpstream) GetDisableReasonOk() (*string, bool)`

GetDisableReasonOk returns a tuple with the DisableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReason

`func (o *DebUpstream) SetDisableReason(v string)`

SetDisableReason sets DisableReason field to given value.

### HasDisableReason

`func (o *DebUpstream) HasDisableReason() bool`

HasDisableReason returns a boolean if a field has been set.

### GetDistroVersions

`func (o *DebUpstream) GetDistroVersions() []string`

GetDistroVersions returns the DistroVersions field if non-nil, zero value otherwise.

### GetDistroVersionsOk

`func (o *DebUpstream) GetDistroVersionsOk() (*[]string, bool)`

GetDistroVersionsOk returns a tuple with the DistroVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroVersions

`func (o *DebUpstream) SetDistroVersions(v []string)`

SetDistroVersions sets DistroVersions field to given value.


### GetExtraHeader1

`func (o *DebUpstream) GetExtraHeader1() string`

GetExtraHeader1 returns the ExtraHeader1 field if non-nil, zero value otherwise.

### GetExtraHeader1Ok

`func (o *DebUpstream) GetExtraHeader1Ok() (*string, bool)`

GetExtraHeader1Ok returns a tuple with the ExtraHeader1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader1

`func (o *DebUpstream) SetExtraHeader1(v string)`

SetExtraHeader1 sets ExtraHeader1 field to given value.

### HasExtraHeader1

`func (o *DebUpstream) HasExtraHeader1() bool`

HasExtraHeader1 returns a boolean if a field has been set.

### SetExtraHeader1Nil

`func (o *DebUpstream) SetExtraHeader1Nil(b bool)`

 SetExtraHeader1Nil sets the value for ExtraHeader1 to be an explicit nil

### UnsetExtraHeader1
`func (o *DebUpstream) UnsetExtraHeader1()`

UnsetExtraHeader1 ensures that no value is present for ExtraHeader1, not even an explicit nil
### GetExtraHeader2

`func (o *DebUpstream) GetExtraHeader2() string`

GetExtraHeader2 returns the ExtraHeader2 field if non-nil, zero value otherwise.

### GetExtraHeader2Ok

`func (o *DebUpstream) GetExtraHeader2Ok() (*string, bool)`

GetExtraHeader2Ok returns a tuple with the ExtraHeader2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader2

`func (o *DebUpstream) SetExtraHeader2(v string)`

SetExtraHeader2 sets ExtraHeader2 field to given value.

### HasExtraHeader2

`func (o *DebUpstream) HasExtraHeader2() bool`

HasExtraHeader2 returns a boolean if a field has been set.

### SetExtraHeader2Nil

`func (o *DebUpstream) SetExtraHeader2Nil(b bool)`

 SetExtraHeader2Nil sets the value for ExtraHeader2 to be an explicit nil

### UnsetExtraHeader2
`func (o *DebUpstream) UnsetExtraHeader2()`

UnsetExtraHeader2 ensures that no value is present for ExtraHeader2, not even an explicit nil
### GetExtraValue1

`func (o *DebUpstream) GetExtraValue1() string`

GetExtraValue1 returns the ExtraValue1 field if non-nil, zero value otherwise.

### GetExtraValue1Ok

`func (o *DebUpstream) GetExtraValue1Ok() (*string, bool)`

GetExtraValue1Ok returns a tuple with the ExtraValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue1

`func (o *DebUpstream) SetExtraValue1(v string)`

SetExtraValue1 sets ExtraValue1 field to given value.

### HasExtraValue1

`func (o *DebUpstream) HasExtraValue1() bool`

HasExtraValue1 returns a boolean if a field has been set.

### SetExtraValue1Nil

`func (o *DebUpstream) SetExtraValue1Nil(b bool)`

 SetExtraValue1Nil sets the value for ExtraValue1 to be an explicit nil

### UnsetExtraValue1
`func (o *DebUpstream) UnsetExtraValue1()`

UnsetExtraValue1 ensures that no value is present for ExtraValue1, not even an explicit nil
### GetExtraValue2

`func (o *DebUpstream) GetExtraValue2() string`

GetExtraValue2 returns the ExtraValue2 field if non-nil, zero value otherwise.

### GetExtraValue2Ok

`func (o *DebUpstream) GetExtraValue2Ok() (*string, bool)`

GetExtraValue2Ok returns a tuple with the ExtraValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue2

`func (o *DebUpstream) SetExtraValue2(v string)`

SetExtraValue2 sets ExtraValue2 field to given value.

### HasExtraValue2

`func (o *DebUpstream) HasExtraValue2() bool`

HasExtraValue2 returns a boolean if a field has been set.

### SetExtraValue2Nil

`func (o *DebUpstream) SetExtraValue2Nil(b bool)`

 SetExtraValue2Nil sets the value for ExtraValue2 to be an explicit nil

### UnsetExtraValue2
`func (o *DebUpstream) UnsetExtraValue2()`

UnsetExtraValue2 ensures that no value is present for ExtraValue2, not even an explicit nil
### GetGpgKeyInline

`func (o *DebUpstream) GetGpgKeyInline() string`

GetGpgKeyInline returns the GpgKeyInline field if non-nil, zero value otherwise.

### GetGpgKeyInlineOk

`func (o *DebUpstream) GetGpgKeyInlineOk() (*string, bool)`

GetGpgKeyInlineOk returns a tuple with the GpgKeyInline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgKeyInline

`func (o *DebUpstream) SetGpgKeyInline(v string)`

SetGpgKeyInline sets GpgKeyInline field to given value.

### HasGpgKeyInline

`func (o *DebUpstream) HasGpgKeyInline() bool`

HasGpgKeyInline returns a boolean if a field has been set.

### SetGpgKeyInlineNil

`func (o *DebUpstream) SetGpgKeyInlineNil(b bool)`

 SetGpgKeyInlineNil sets the value for GpgKeyInline to be an explicit nil

### UnsetGpgKeyInline
`func (o *DebUpstream) UnsetGpgKeyInline()`

UnsetGpgKeyInline ensures that no value is present for GpgKeyInline, not even an explicit nil
### GetGpgKeyUrl

`func (o *DebUpstream) GetGpgKeyUrl() string`

GetGpgKeyUrl returns the GpgKeyUrl field if non-nil, zero value otherwise.

### GetGpgKeyUrlOk

`func (o *DebUpstream) GetGpgKeyUrlOk() (*string, bool)`

GetGpgKeyUrlOk returns a tuple with the GpgKeyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgKeyUrl

`func (o *DebUpstream) SetGpgKeyUrl(v string)`

SetGpgKeyUrl sets GpgKeyUrl field to given value.

### HasGpgKeyUrl

`func (o *DebUpstream) HasGpgKeyUrl() bool`

HasGpgKeyUrl returns a boolean if a field has been set.

### SetGpgKeyUrlNil

`func (o *DebUpstream) SetGpgKeyUrlNil(b bool)`

 SetGpgKeyUrlNil sets the value for GpgKeyUrl to be an explicit nil

### UnsetGpgKeyUrl
`func (o *DebUpstream) UnsetGpgKeyUrl()`

UnsetGpgKeyUrl ensures that no value is present for GpgKeyUrl, not even an explicit nil
### GetGpgVerification

`func (o *DebUpstream) GetGpgVerification() string`

GetGpgVerification returns the GpgVerification field if non-nil, zero value otherwise.

### GetGpgVerificationOk

`func (o *DebUpstream) GetGpgVerificationOk() (*string, bool)`

GetGpgVerificationOk returns a tuple with the GpgVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgVerification

`func (o *DebUpstream) SetGpgVerification(v string)`

SetGpgVerification sets GpgVerification field to given value.

### HasGpgVerification

`func (o *DebUpstream) HasGpgVerification() bool`

HasGpgVerification returns a boolean if a field has been set.

### GetIncludeSources

`func (o *DebUpstream) GetIncludeSources() bool`

GetIncludeSources returns the IncludeSources field if non-nil, zero value otherwise.

### GetIncludeSourcesOk

`func (o *DebUpstream) GetIncludeSourcesOk() (*bool, bool)`

GetIncludeSourcesOk returns a tuple with the IncludeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSources

`func (o *DebUpstream) SetIncludeSources(v bool)`

SetIncludeSources sets IncludeSources field to given value.

### HasIncludeSources

`func (o *DebUpstream) HasIncludeSources() bool`

HasIncludeSources returns a boolean if a field has been set.

### GetIsActive

`func (o *DebUpstream) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DebUpstream) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DebUpstream) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *DebUpstream) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetMode

`func (o *DebUpstream) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DebUpstream) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DebUpstream) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DebUpstream) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *DebUpstream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DebUpstream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DebUpstream) SetName(v string)`

SetName sets Name field to given value.


### GetPendingValidation

`func (o *DebUpstream) GetPendingValidation() bool`

GetPendingValidation returns the PendingValidation field if non-nil, zero value otherwise.

### GetPendingValidationOk

`func (o *DebUpstream) GetPendingValidationOk() (*bool, bool)`

GetPendingValidationOk returns a tuple with the PendingValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingValidation

`func (o *DebUpstream) SetPendingValidation(v bool)`

SetPendingValidation sets PendingValidation field to given value.

### HasPendingValidation

`func (o *DebUpstream) HasPendingValidation() bool`

HasPendingValidation returns a boolean if a field has been set.

### GetPriority

`func (o *DebUpstream) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DebUpstream) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DebUpstream) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DebUpstream) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSlugPerm

`func (o *DebUpstream) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *DebUpstream) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *DebUpstream) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *DebUpstream) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DebUpstream) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DebUpstream) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DebUpstream) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DebUpstream) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpstreamDistribution

`func (o *DebUpstream) GetUpstreamDistribution() string`

GetUpstreamDistribution returns the UpstreamDistribution field if non-nil, zero value otherwise.

### GetUpstreamDistributionOk

`func (o *DebUpstream) GetUpstreamDistributionOk() (*string, bool)`

GetUpstreamDistributionOk returns a tuple with the UpstreamDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamDistribution

`func (o *DebUpstream) SetUpstreamDistribution(v string)`

SetUpstreamDistribution sets UpstreamDistribution field to given value.

### HasUpstreamDistribution

`func (o *DebUpstream) HasUpstreamDistribution() bool`

HasUpstreamDistribution returns a boolean if a field has been set.

### SetUpstreamDistributionNil

`func (o *DebUpstream) SetUpstreamDistributionNil(b bool)`

 SetUpstreamDistributionNil sets the value for UpstreamDistribution to be an explicit nil

### UnsetUpstreamDistribution
`func (o *DebUpstream) UnsetUpstreamDistribution()`

UnsetUpstreamDistribution ensures that no value is present for UpstreamDistribution, not even an explicit nil
### GetUpstreamUrl

`func (o *DebUpstream) GetUpstreamUrl() string`

GetUpstreamUrl returns the UpstreamUrl field if non-nil, zero value otherwise.

### GetUpstreamUrlOk

`func (o *DebUpstream) GetUpstreamUrlOk() (*string, bool)`

GetUpstreamUrlOk returns a tuple with the UpstreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamUrl

`func (o *DebUpstream) SetUpstreamUrl(v string)`

SetUpstreamUrl sets UpstreamUrl field to given value.


### GetVerificationStatus

`func (o *DebUpstream) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *DebUpstream) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *DebUpstream) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *DebUpstream) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### GetVerifySsl

`func (o *DebUpstream) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *DebUpstream) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *DebUpstream) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *DebUpstream) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


