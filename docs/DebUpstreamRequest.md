# DebUpstreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | Pointer to **string** | The authentication mode to use when accessing this upstream.  | [optional] [default to "None"]
**AuthSecret** | Pointer to **NullableString** | Secret to provide with requests to upstream. | [optional] 
**AuthUsername** | Pointer to **NullableString** | Username to provide with requests to upstream. | [optional] 
**Component** | Pointer to **string** | The component to fetch from the upstream | [optional] 
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
**Priority** | Pointer to **int64** | Upstream sources are selected for resolving requests by sequential order (1..n), followed by creation date. | [optional] 
**UpstreamDistribution** | Pointer to **NullableString** | The distribution to fetch from the upstream | [optional] 
**UpstreamUrl** | **string** | The URL for this upstream source. This must be a fully qualified URL including any path elements required to reach the root of the repository.  | 
**VerifySsl** | Pointer to **bool** | If enabled, SSL certificates are verified when requests are made to this upstream. It&#39;s recommended to leave this enabled for all public sources to help mitigate Man-In-The-Middle (MITM) attacks. Please note this only applies to HTTPS upstreams. | [optional] 

## Methods

### NewDebUpstreamRequest

`func NewDebUpstreamRequest(distroVersions []string, name string, upstreamUrl string, ) *DebUpstreamRequest`

NewDebUpstreamRequest instantiates a new DebUpstreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebUpstreamRequestWithDefaults

`func NewDebUpstreamRequestWithDefaults() *DebUpstreamRequest`

NewDebUpstreamRequestWithDefaults instantiates a new DebUpstreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *DebUpstreamRequest) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *DebUpstreamRequest) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *DebUpstreamRequest) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *DebUpstreamRequest) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAuthSecret

`func (o *DebUpstreamRequest) GetAuthSecret() string`

GetAuthSecret returns the AuthSecret field if non-nil, zero value otherwise.

### GetAuthSecretOk

`func (o *DebUpstreamRequest) GetAuthSecretOk() (*string, bool)`

GetAuthSecretOk returns a tuple with the AuthSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecret

`func (o *DebUpstreamRequest) SetAuthSecret(v string)`

SetAuthSecret sets AuthSecret field to given value.

### HasAuthSecret

`func (o *DebUpstreamRequest) HasAuthSecret() bool`

HasAuthSecret returns a boolean if a field has been set.

### SetAuthSecretNil

`func (o *DebUpstreamRequest) SetAuthSecretNil(b bool)`

 SetAuthSecretNil sets the value for AuthSecret to be an explicit nil

### UnsetAuthSecret
`func (o *DebUpstreamRequest) UnsetAuthSecret()`

UnsetAuthSecret ensures that no value is present for AuthSecret, not even an explicit nil
### GetAuthUsername

`func (o *DebUpstreamRequest) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *DebUpstreamRequest) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *DebUpstreamRequest) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *DebUpstreamRequest) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### SetAuthUsernameNil

`func (o *DebUpstreamRequest) SetAuthUsernameNil(b bool)`

 SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil

### UnsetAuthUsername
`func (o *DebUpstreamRequest) UnsetAuthUsername()`

UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
### GetComponent

`func (o *DebUpstreamRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *DebUpstreamRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *DebUpstreamRequest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *DebUpstreamRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetDistroVersions

`func (o *DebUpstreamRequest) GetDistroVersions() []string`

GetDistroVersions returns the DistroVersions field if non-nil, zero value otherwise.

### GetDistroVersionsOk

`func (o *DebUpstreamRequest) GetDistroVersionsOk() (*[]string, bool)`

GetDistroVersionsOk returns a tuple with the DistroVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroVersions

`func (o *DebUpstreamRequest) SetDistroVersions(v []string)`

SetDistroVersions sets DistroVersions field to given value.


### GetExtraHeader1

`func (o *DebUpstreamRequest) GetExtraHeader1() string`

GetExtraHeader1 returns the ExtraHeader1 field if non-nil, zero value otherwise.

### GetExtraHeader1Ok

`func (o *DebUpstreamRequest) GetExtraHeader1Ok() (*string, bool)`

GetExtraHeader1Ok returns a tuple with the ExtraHeader1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader1

`func (o *DebUpstreamRequest) SetExtraHeader1(v string)`

SetExtraHeader1 sets ExtraHeader1 field to given value.

### HasExtraHeader1

`func (o *DebUpstreamRequest) HasExtraHeader1() bool`

HasExtraHeader1 returns a boolean if a field has been set.

### SetExtraHeader1Nil

`func (o *DebUpstreamRequest) SetExtraHeader1Nil(b bool)`

 SetExtraHeader1Nil sets the value for ExtraHeader1 to be an explicit nil

### UnsetExtraHeader1
`func (o *DebUpstreamRequest) UnsetExtraHeader1()`

UnsetExtraHeader1 ensures that no value is present for ExtraHeader1, not even an explicit nil
### GetExtraHeader2

`func (o *DebUpstreamRequest) GetExtraHeader2() string`

GetExtraHeader2 returns the ExtraHeader2 field if non-nil, zero value otherwise.

### GetExtraHeader2Ok

`func (o *DebUpstreamRequest) GetExtraHeader2Ok() (*string, bool)`

GetExtraHeader2Ok returns a tuple with the ExtraHeader2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader2

`func (o *DebUpstreamRequest) SetExtraHeader2(v string)`

SetExtraHeader2 sets ExtraHeader2 field to given value.

### HasExtraHeader2

`func (o *DebUpstreamRequest) HasExtraHeader2() bool`

HasExtraHeader2 returns a boolean if a field has been set.

### SetExtraHeader2Nil

`func (o *DebUpstreamRequest) SetExtraHeader2Nil(b bool)`

 SetExtraHeader2Nil sets the value for ExtraHeader2 to be an explicit nil

### UnsetExtraHeader2
`func (o *DebUpstreamRequest) UnsetExtraHeader2()`

UnsetExtraHeader2 ensures that no value is present for ExtraHeader2, not even an explicit nil
### GetExtraValue1

`func (o *DebUpstreamRequest) GetExtraValue1() string`

GetExtraValue1 returns the ExtraValue1 field if non-nil, zero value otherwise.

### GetExtraValue1Ok

`func (o *DebUpstreamRequest) GetExtraValue1Ok() (*string, bool)`

GetExtraValue1Ok returns a tuple with the ExtraValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue1

`func (o *DebUpstreamRequest) SetExtraValue1(v string)`

SetExtraValue1 sets ExtraValue1 field to given value.

### HasExtraValue1

`func (o *DebUpstreamRequest) HasExtraValue1() bool`

HasExtraValue1 returns a boolean if a field has been set.

### SetExtraValue1Nil

`func (o *DebUpstreamRequest) SetExtraValue1Nil(b bool)`

 SetExtraValue1Nil sets the value for ExtraValue1 to be an explicit nil

### UnsetExtraValue1
`func (o *DebUpstreamRequest) UnsetExtraValue1()`

UnsetExtraValue1 ensures that no value is present for ExtraValue1, not even an explicit nil
### GetExtraValue2

`func (o *DebUpstreamRequest) GetExtraValue2() string`

GetExtraValue2 returns the ExtraValue2 field if non-nil, zero value otherwise.

### GetExtraValue2Ok

`func (o *DebUpstreamRequest) GetExtraValue2Ok() (*string, bool)`

GetExtraValue2Ok returns a tuple with the ExtraValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue2

`func (o *DebUpstreamRequest) SetExtraValue2(v string)`

SetExtraValue2 sets ExtraValue2 field to given value.

### HasExtraValue2

`func (o *DebUpstreamRequest) HasExtraValue2() bool`

HasExtraValue2 returns a boolean if a field has been set.

### SetExtraValue2Nil

`func (o *DebUpstreamRequest) SetExtraValue2Nil(b bool)`

 SetExtraValue2Nil sets the value for ExtraValue2 to be an explicit nil

### UnsetExtraValue2
`func (o *DebUpstreamRequest) UnsetExtraValue2()`

UnsetExtraValue2 ensures that no value is present for ExtraValue2, not even an explicit nil
### GetGpgKeyInline

`func (o *DebUpstreamRequest) GetGpgKeyInline() string`

GetGpgKeyInline returns the GpgKeyInline field if non-nil, zero value otherwise.

### GetGpgKeyInlineOk

`func (o *DebUpstreamRequest) GetGpgKeyInlineOk() (*string, bool)`

GetGpgKeyInlineOk returns a tuple with the GpgKeyInline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgKeyInline

`func (o *DebUpstreamRequest) SetGpgKeyInline(v string)`

SetGpgKeyInline sets GpgKeyInline field to given value.

### HasGpgKeyInline

`func (o *DebUpstreamRequest) HasGpgKeyInline() bool`

HasGpgKeyInline returns a boolean if a field has been set.

### SetGpgKeyInlineNil

`func (o *DebUpstreamRequest) SetGpgKeyInlineNil(b bool)`

 SetGpgKeyInlineNil sets the value for GpgKeyInline to be an explicit nil

### UnsetGpgKeyInline
`func (o *DebUpstreamRequest) UnsetGpgKeyInline()`

UnsetGpgKeyInline ensures that no value is present for GpgKeyInline, not even an explicit nil
### GetGpgKeyUrl

`func (o *DebUpstreamRequest) GetGpgKeyUrl() string`

GetGpgKeyUrl returns the GpgKeyUrl field if non-nil, zero value otherwise.

### GetGpgKeyUrlOk

`func (o *DebUpstreamRequest) GetGpgKeyUrlOk() (*string, bool)`

GetGpgKeyUrlOk returns a tuple with the GpgKeyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgKeyUrl

`func (o *DebUpstreamRequest) SetGpgKeyUrl(v string)`

SetGpgKeyUrl sets GpgKeyUrl field to given value.

### HasGpgKeyUrl

`func (o *DebUpstreamRequest) HasGpgKeyUrl() bool`

HasGpgKeyUrl returns a boolean if a field has been set.

### SetGpgKeyUrlNil

`func (o *DebUpstreamRequest) SetGpgKeyUrlNil(b bool)`

 SetGpgKeyUrlNil sets the value for GpgKeyUrl to be an explicit nil

### UnsetGpgKeyUrl
`func (o *DebUpstreamRequest) UnsetGpgKeyUrl()`

UnsetGpgKeyUrl ensures that no value is present for GpgKeyUrl, not even an explicit nil
### GetGpgVerification

`func (o *DebUpstreamRequest) GetGpgVerification() string`

GetGpgVerification returns the GpgVerification field if non-nil, zero value otherwise.

### GetGpgVerificationOk

`func (o *DebUpstreamRequest) GetGpgVerificationOk() (*string, bool)`

GetGpgVerificationOk returns a tuple with the GpgVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgVerification

`func (o *DebUpstreamRequest) SetGpgVerification(v string)`

SetGpgVerification sets GpgVerification field to given value.

### HasGpgVerification

`func (o *DebUpstreamRequest) HasGpgVerification() bool`

HasGpgVerification returns a boolean if a field has been set.

### GetIncludeSources

`func (o *DebUpstreamRequest) GetIncludeSources() bool`

GetIncludeSources returns the IncludeSources field if non-nil, zero value otherwise.

### GetIncludeSourcesOk

`func (o *DebUpstreamRequest) GetIncludeSourcesOk() (*bool, bool)`

GetIncludeSourcesOk returns a tuple with the IncludeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSources

`func (o *DebUpstreamRequest) SetIncludeSources(v bool)`

SetIncludeSources sets IncludeSources field to given value.

### HasIncludeSources

`func (o *DebUpstreamRequest) HasIncludeSources() bool`

HasIncludeSources returns a boolean if a field has been set.

### GetIsActive

`func (o *DebUpstreamRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DebUpstreamRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DebUpstreamRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *DebUpstreamRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetMode

`func (o *DebUpstreamRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DebUpstreamRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DebUpstreamRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DebUpstreamRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *DebUpstreamRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DebUpstreamRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DebUpstreamRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *DebUpstreamRequest) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DebUpstreamRequest) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DebUpstreamRequest) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DebUpstreamRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetUpstreamDistribution

`func (o *DebUpstreamRequest) GetUpstreamDistribution() string`

GetUpstreamDistribution returns the UpstreamDistribution field if non-nil, zero value otherwise.

### GetUpstreamDistributionOk

`func (o *DebUpstreamRequest) GetUpstreamDistributionOk() (*string, bool)`

GetUpstreamDistributionOk returns a tuple with the UpstreamDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamDistribution

`func (o *DebUpstreamRequest) SetUpstreamDistribution(v string)`

SetUpstreamDistribution sets UpstreamDistribution field to given value.

### HasUpstreamDistribution

`func (o *DebUpstreamRequest) HasUpstreamDistribution() bool`

HasUpstreamDistribution returns a boolean if a field has been set.

### SetUpstreamDistributionNil

`func (o *DebUpstreamRequest) SetUpstreamDistributionNil(b bool)`

 SetUpstreamDistributionNil sets the value for UpstreamDistribution to be an explicit nil

### UnsetUpstreamDistribution
`func (o *DebUpstreamRequest) UnsetUpstreamDistribution()`

UnsetUpstreamDistribution ensures that no value is present for UpstreamDistribution, not even an explicit nil
### GetUpstreamUrl

`func (o *DebUpstreamRequest) GetUpstreamUrl() string`

GetUpstreamUrl returns the UpstreamUrl field if non-nil, zero value otherwise.

### GetUpstreamUrlOk

`func (o *DebUpstreamRequest) GetUpstreamUrlOk() (*string, bool)`

GetUpstreamUrlOk returns a tuple with the UpstreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamUrl

`func (o *DebUpstreamRequest) SetUpstreamUrl(v string)`

SetUpstreamUrl sets UpstreamUrl field to given value.


### GetVerifySsl

`func (o *DebUpstreamRequest) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *DebUpstreamRequest) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *DebUpstreamRequest) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *DebUpstreamRequest) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


