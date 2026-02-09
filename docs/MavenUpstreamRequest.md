# MavenUpstreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | Pointer to **string** | The authentication mode to use when accessing this upstream.  | [optional] [default to "None"]
**AuthSecret** | Pointer to **NullableString** | Secret to provide with requests to upstream. | [optional] 
**AuthUsername** | Pointer to **NullableString** | Username to provide with requests to upstream. | [optional] 
**ExtraHeader1** | Pointer to **NullableString** | The key for extra header #1 to send to upstream. | [optional] 
**ExtraHeader2** | Pointer to **NullableString** | The key for extra header #2 to send to upstream. | [optional] 
**ExtraValue1** | Pointer to **NullableString** | The value for extra header #1 to send to upstream. This is stored as plaintext, and is NOT encrypted. | [optional] 
**ExtraValue2** | Pointer to **NullableString** | The value for extra header #2 to send to upstream. This is stored as plaintext, and is NOT encrypted. | [optional] 
**GpgKeyInline** | Pointer to **NullableString** | A public GPG key to associate with packages found on this upstream. When using the Cloudsmith setup script, this GPG key will be automatically imported on your deployment machines to allow upstream packages to validate and install. | [optional] 
**GpgKeyUrl** | Pointer to **NullableString** | When provided, Cloudsmith will fetch, validate, and associate a public GPG key found at the provided URL. When using the Cloudsmith setup script, this GPG key will be automatically imported on your deployment machines to allow upstream packages to validate and install. | [optional] 
**GpgVerification** | Pointer to **string** | The GPG signature verification mode for this upstream. | [optional] [default to "Allow All"]
**IsActive** | Pointer to **bool** | Whether or not this upstream is active and ready for requests. | [optional] 
**Mode** | Pointer to **string** | The mode that this upstream should operate in. Upstream sources can be used to proxy resolved packages, as well as operate in a proxy/cache or cache only mode. | [optional] [default to "Proxy Only"]
**Name** | **string** | A descriptive name for this upstream source. A shortened version of this name will be used for tagging cached packages retrieved from this upstream. | 
**Priority** | Pointer to **int64** | Upstream sources are selected for resolving requests by sequential order (1..n), followed by creation date. | [optional] 
**TrustLevel** | Pointer to **string** | Trust level allows for control of the visibility of upstream artifacts to native package managers. Where supported by formats, the default level (untrusted) is recommended for all upstreams, and helps to safeguard against common dependency confusion attack vectors. | [optional] [default to "Trusted"]
**UpstreamUrl** | **string** | The URL for this upstream source. This must be a fully qualified URL including any path elements required to reach the root of the repository.  | 
**VerifySsl** | Pointer to **bool** | If enabled, SSL certificates are verified when requests are made to this upstream. It&#39;s recommended to leave this enabled for all public sources to help mitigate Man-In-The-Middle (MITM) attacks. Please note this only applies to HTTPS upstreams. | [optional] 

## Methods

### NewMavenUpstreamRequest

`func NewMavenUpstreamRequest(name string, upstreamUrl string, ) *MavenUpstreamRequest`

NewMavenUpstreamRequest instantiates a new MavenUpstreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenUpstreamRequestWithDefaults

`func NewMavenUpstreamRequestWithDefaults() *MavenUpstreamRequest`

NewMavenUpstreamRequestWithDefaults instantiates a new MavenUpstreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *MavenUpstreamRequest) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *MavenUpstreamRequest) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *MavenUpstreamRequest) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *MavenUpstreamRequest) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAuthSecret

`func (o *MavenUpstreamRequest) GetAuthSecret() string`

GetAuthSecret returns the AuthSecret field if non-nil, zero value otherwise.

### GetAuthSecretOk

`func (o *MavenUpstreamRequest) GetAuthSecretOk() (*string, bool)`

GetAuthSecretOk returns a tuple with the AuthSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecret

`func (o *MavenUpstreamRequest) SetAuthSecret(v string)`

SetAuthSecret sets AuthSecret field to given value.

### HasAuthSecret

`func (o *MavenUpstreamRequest) HasAuthSecret() bool`

HasAuthSecret returns a boolean if a field has been set.

### SetAuthSecretNil

`func (o *MavenUpstreamRequest) SetAuthSecretNil(b bool)`

 SetAuthSecretNil sets the value for AuthSecret to be an explicit nil

### UnsetAuthSecret
`func (o *MavenUpstreamRequest) UnsetAuthSecret()`

UnsetAuthSecret ensures that no value is present for AuthSecret, not even an explicit nil
### GetAuthUsername

`func (o *MavenUpstreamRequest) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *MavenUpstreamRequest) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *MavenUpstreamRequest) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *MavenUpstreamRequest) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### SetAuthUsernameNil

`func (o *MavenUpstreamRequest) SetAuthUsernameNil(b bool)`

 SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil

### UnsetAuthUsername
`func (o *MavenUpstreamRequest) UnsetAuthUsername()`

UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
### GetExtraHeader1

`func (o *MavenUpstreamRequest) GetExtraHeader1() string`

GetExtraHeader1 returns the ExtraHeader1 field if non-nil, zero value otherwise.

### GetExtraHeader1Ok

`func (o *MavenUpstreamRequest) GetExtraHeader1Ok() (*string, bool)`

GetExtraHeader1Ok returns a tuple with the ExtraHeader1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader1

`func (o *MavenUpstreamRequest) SetExtraHeader1(v string)`

SetExtraHeader1 sets ExtraHeader1 field to given value.

### HasExtraHeader1

`func (o *MavenUpstreamRequest) HasExtraHeader1() bool`

HasExtraHeader1 returns a boolean if a field has been set.

### SetExtraHeader1Nil

`func (o *MavenUpstreamRequest) SetExtraHeader1Nil(b bool)`

 SetExtraHeader1Nil sets the value for ExtraHeader1 to be an explicit nil

### UnsetExtraHeader1
`func (o *MavenUpstreamRequest) UnsetExtraHeader1()`

UnsetExtraHeader1 ensures that no value is present for ExtraHeader1, not even an explicit nil
### GetExtraHeader2

`func (o *MavenUpstreamRequest) GetExtraHeader2() string`

GetExtraHeader2 returns the ExtraHeader2 field if non-nil, zero value otherwise.

### GetExtraHeader2Ok

`func (o *MavenUpstreamRequest) GetExtraHeader2Ok() (*string, bool)`

GetExtraHeader2Ok returns a tuple with the ExtraHeader2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader2

`func (o *MavenUpstreamRequest) SetExtraHeader2(v string)`

SetExtraHeader2 sets ExtraHeader2 field to given value.

### HasExtraHeader2

`func (o *MavenUpstreamRequest) HasExtraHeader2() bool`

HasExtraHeader2 returns a boolean if a field has been set.

### SetExtraHeader2Nil

`func (o *MavenUpstreamRequest) SetExtraHeader2Nil(b bool)`

 SetExtraHeader2Nil sets the value for ExtraHeader2 to be an explicit nil

### UnsetExtraHeader2
`func (o *MavenUpstreamRequest) UnsetExtraHeader2()`

UnsetExtraHeader2 ensures that no value is present for ExtraHeader2, not even an explicit nil
### GetExtraValue1

`func (o *MavenUpstreamRequest) GetExtraValue1() string`

GetExtraValue1 returns the ExtraValue1 field if non-nil, zero value otherwise.

### GetExtraValue1Ok

`func (o *MavenUpstreamRequest) GetExtraValue1Ok() (*string, bool)`

GetExtraValue1Ok returns a tuple with the ExtraValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue1

`func (o *MavenUpstreamRequest) SetExtraValue1(v string)`

SetExtraValue1 sets ExtraValue1 field to given value.

### HasExtraValue1

`func (o *MavenUpstreamRequest) HasExtraValue1() bool`

HasExtraValue1 returns a boolean if a field has been set.

### SetExtraValue1Nil

`func (o *MavenUpstreamRequest) SetExtraValue1Nil(b bool)`

 SetExtraValue1Nil sets the value for ExtraValue1 to be an explicit nil

### UnsetExtraValue1
`func (o *MavenUpstreamRequest) UnsetExtraValue1()`

UnsetExtraValue1 ensures that no value is present for ExtraValue1, not even an explicit nil
### GetExtraValue2

`func (o *MavenUpstreamRequest) GetExtraValue2() string`

GetExtraValue2 returns the ExtraValue2 field if non-nil, zero value otherwise.

### GetExtraValue2Ok

`func (o *MavenUpstreamRequest) GetExtraValue2Ok() (*string, bool)`

GetExtraValue2Ok returns a tuple with the ExtraValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue2

`func (o *MavenUpstreamRequest) SetExtraValue2(v string)`

SetExtraValue2 sets ExtraValue2 field to given value.

### HasExtraValue2

`func (o *MavenUpstreamRequest) HasExtraValue2() bool`

HasExtraValue2 returns a boolean if a field has been set.

### SetExtraValue2Nil

`func (o *MavenUpstreamRequest) SetExtraValue2Nil(b bool)`

 SetExtraValue2Nil sets the value for ExtraValue2 to be an explicit nil

### UnsetExtraValue2
`func (o *MavenUpstreamRequest) UnsetExtraValue2()`

UnsetExtraValue2 ensures that no value is present for ExtraValue2, not even an explicit nil
### GetGpgKeyInline

`func (o *MavenUpstreamRequest) GetGpgKeyInline() string`

GetGpgKeyInline returns the GpgKeyInline field if non-nil, zero value otherwise.

### GetGpgKeyInlineOk

`func (o *MavenUpstreamRequest) GetGpgKeyInlineOk() (*string, bool)`

GetGpgKeyInlineOk returns a tuple with the GpgKeyInline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgKeyInline

`func (o *MavenUpstreamRequest) SetGpgKeyInline(v string)`

SetGpgKeyInline sets GpgKeyInline field to given value.

### HasGpgKeyInline

`func (o *MavenUpstreamRequest) HasGpgKeyInline() bool`

HasGpgKeyInline returns a boolean if a field has been set.

### SetGpgKeyInlineNil

`func (o *MavenUpstreamRequest) SetGpgKeyInlineNil(b bool)`

 SetGpgKeyInlineNil sets the value for GpgKeyInline to be an explicit nil

### UnsetGpgKeyInline
`func (o *MavenUpstreamRequest) UnsetGpgKeyInline()`

UnsetGpgKeyInline ensures that no value is present for GpgKeyInline, not even an explicit nil
### GetGpgKeyUrl

`func (o *MavenUpstreamRequest) GetGpgKeyUrl() string`

GetGpgKeyUrl returns the GpgKeyUrl field if non-nil, zero value otherwise.

### GetGpgKeyUrlOk

`func (o *MavenUpstreamRequest) GetGpgKeyUrlOk() (*string, bool)`

GetGpgKeyUrlOk returns a tuple with the GpgKeyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgKeyUrl

`func (o *MavenUpstreamRequest) SetGpgKeyUrl(v string)`

SetGpgKeyUrl sets GpgKeyUrl field to given value.

### HasGpgKeyUrl

`func (o *MavenUpstreamRequest) HasGpgKeyUrl() bool`

HasGpgKeyUrl returns a boolean if a field has been set.

### SetGpgKeyUrlNil

`func (o *MavenUpstreamRequest) SetGpgKeyUrlNil(b bool)`

 SetGpgKeyUrlNil sets the value for GpgKeyUrl to be an explicit nil

### UnsetGpgKeyUrl
`func (o *MavenUpstreamRequest) UnsetGpgKeyUrl()`

UnsetGpgKeyUrl ensures that no value is present for GpgKeyUrl, not even an explicit nil
### GetGpgVerification

`func (o *MavenUpstreamRequest) GetGpgVerification() string`

GetGpgVerification returns the GpgVerification field if non-nil, zero value otherwise.

### GetGpgVerificationOk

`func (o *MavenUpstreamRequest) GetGpgVerificationOk() (*string, bool)`

GetGpgVerificationOk returns a tuple with the GpgVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgVerification

`func (o *MavenUpstreamRequest) SetGpgVerification(v string)`

SetGpgVerification sets GpgVerification field to given value.

### HasGpgVerification

`func (o *MavenUpstreamRequest) HasGpgVerification() bool`

HasGpgVerification returns a boolean if a field has been set.

### GetIsActive

`func (o *MavenUpstreamRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *MavenUpstreamRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *MavenUpstreamRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *MavenUpstreamRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetMode

`func (o *MavenUpstreamRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MavenUpstreamRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MavenUpstreamRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *MavenUpstreamRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *MavenUpstreamRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MavenUpstreamRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MavenUpstreamRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *MavenUpstreamRequest) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MavenUpstreamRequest) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MavenUpstreamRequest) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MavenUpstreamRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTrustLevel

`func (o *MavenUpstreamRequest) GetTrustLevel() string`

GetTrustLevel returns the TrustLevel field if non-nil, zero value otherwise.

### GetTrustLevelOk

`func (o *MavenUpstreamRequest) GetTrustLevelOk() (*string, bool)`

GetTrustLevelOk returns a tuple with the TrustLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustLevel

`func (o *MavenUpstreamRequest) SetTrustLevel(v string)`

SetTrustLevel sets TrustLevel field to given value.

### HasTrustLevel

`func (o *MavenUpstreamRequest) HasTrustLevel() bool`

HasTrustLevel returns a boolean if a field has been set.

### GetUpstreamUrl

`func (o *MavenUpstreamRequest) GetUpstreamUrl() string`

GetUpstreamUrl returns the UpstreamUrl field if non-nil, zero value otherwise.

### GetUpstreamUrlOk

`func (o *MavenUpstreamRequest) GetUpstreamUrlOk() (*string, bool)`

GetUpstreamUrlOk returns a tuple with the UpstreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamUrl

`func (o *MavenUpstreamRequest) SetUpstreamUrl(v string)`

SetUpstreamUrl sets UpstreamUrl field to given value.


### GetVerifySsl

`func (o *MavenUpstreamRequest) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *MavenUpstreamRequest) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *MavenUpstreamRequest) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *MavenUpstreamRequest) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


