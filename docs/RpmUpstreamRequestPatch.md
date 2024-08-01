# RpmUpstreamRequestPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | Pointer to **string** | The authentication mode to use when accessing this upstream.  | [optional] [default to "None"]
**AuthSecret** | Pointer to **NullableString** | Secret to provide with requests to upstream. | [optional] 
**AuthUsername** | Pointer to **NullableString** | Username to provide with requests to upstream. | [optional] 
**DistroVersion** | Pointer to **string** | The distribution version that packages found on this upstream will be associated with. | [optional] 
**ExtraHeader1** | Pointer to **NullableString** | The key for extra header #1 to send to upstream. | [optional] 
**ExtraHeader2** | Pointer to **NullableString** | The key for extra header #2 to send to upstream. | [optional] 
**ExtraValue1** | Pointer to **NullableString** | The value for extra header #1 to send to upstream. This is stored as plaintext, and is NOT encrypted. | [optional] 
**ExtraValue2** | Pointer to **NullableString** | The value for extra header #2 to send to upstream. This is stored as plaintext, and is NOT encrypted. | [optional] 
**GpgKeyInline** | Pointer to **NullableString** | A public GPG key to associate with packages found on this upstream. When using the Cloudsmith setup script, this GPG key will be automatically imported on your deployment machines to allow upstream packages to validate and install. | [optional] 
**GpgKeyUrl** | Pointer to **NullableString** | When provided, Cloudsmith will fetch, validate, and associate a public GPG key found at the provided URL. When using the Cloudsmith setup script, this GPG key will be automatically imported on your deployment machines to allow upstream packages to validate and install. | [optional] 
**GpgVerification** | Pointer to **string** | The GPG signature verification mode for this upstream. | [optional] [default to "Allow All"]
**IncludeSources** | Pointer to **bool** | When checked, source packages will be available from this upstream. | [optional] 
**IsActive** | Pointer to **bool** | Whether or not this upstream is active and ready for requests. | [optional] 
**Mode** | Pointer to **string** | The mode that this upstream should operate in. Upstream sources can be used to proxy resolved packages, as well as operate in a proxy/cache or cache only mode. | [optional] [default to "Proxy Only"]
**Name** | Pointer to **string** | A descriptive name for this upstream source. A shortened version of this name will be used for tagging cached packages retrieved from this upstream. | [optional] 
**Priority** | Pointer to **int64** | Upstream sources are selected for resolving requests by sequential order (1..n), followed by creation date. | [optional] 
**UpstreamUrl** | Pointer to **string** | The URL for this upstream source. This must be a fully qualified URL including any path elements required to reach the root of the repository.  | [optional] 
**VerifySsl** | Pointer to **bool** | If enabled, SSL certificates are verified when requests are made to this upstream. It&#39;s recommended to leave this enabled for all public sources to help mitigate Man-In-The-Middle (MITM) attacks. Please note this only applies to HTTPS upstreams. | [optional] 

## Methods

### NewRpmUpstreamRequestPatch

`func NewRpmUpstreamRequestPatch() *RpmUpstreamRequestPatch`

NewRpmUpstreamRequestPatch instantiates a new RpmUpstreamRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmUpstreamRequestPatchWithDefaults

`func NewRpmUpstreamRequestPatchWithDefaults() *RpmUpstreamRequestPatch`

NewRpmUpstreamRequestPatchWithDefaults instantiates a new RpmUpstreamRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *RpmUpstreamRequestPatch) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *RpmUpstreamRequestPatch) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *RpmUpstreamRequestPatch) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *RpmUpstreamRequestPatch) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAuthSecret

`func (o *RpmUpstreamRequestPatch) GetAuthSecret() string`

GetAuthSecret returns the AuthSecret field if non-nil, zero value otherwise.

### GetAuthSecretOk

`func (o *RpmUpstreamRequestPatch) GetAuthSecretOk() (*string, bool)`

GetAuthSecretOk returns a tuple with the AuthSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecret

`func (o *RpmUpstreamRequestPatch) SetAuthSecret(v string)`

SetAuthSecret sets AuthSecret field to given value.

### HasAuthSecret

`func (o *RpmUpstreamRequestPatch) HasAuthSecret() bool`

HasAuthSecret returns a boolean if a field has been set.

### SetAuthSecretNil

`func (o *RpmUpstreamRequestPatch) SetAuthSecretNil(b bool)`

 SetAuthSecretNil sets the value for AuthSecret to be an explicit nil

### UnsetAuthSecret
`func (o *RpmUpstreamRequestPatch) UnsetAuthSecret()`

UnsetAuthSecret ensures that no value is present for AuthSecret, not even an explicit nil
### GetAuthUsername

`func (o *RpmUpstreamRequestPatch) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *RpmUpstreamRequestPatch) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *RpmUpstreamRequestPatch) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *RpmUpstreamRequestPatch) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### SetAuthUsernameNil

`func (o *RpmUpstreamRequestPatch) SetAuthUsernameNil(b bool)`

 SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil

### UnsetAuthUsername
`func (o *RpmUpstreamRequestPatch) UnsetAuthUsername()`

UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
### GetDistroVersion

`func (o *RpmUpstreamRequestPatch) GetDistroVersion() string`

GetDistroVersion returns the DistroVersion field if non-nil, zero value otherwise.

### GetDistroVersionOk

`func (o *RpmUpstreamRequestPatch) GetDistroVersionOk() (*string, bool)`

GetDistroVersionOk returns a tuple with the DistroVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroVersion

`func (o *RpmUpstreamRequestPatch) SetDistroVersion(v string)`

SetDistroVersion sets DistroVersion field to given value.

### HasDistroVersion

`func (o *RpmUpstreamRequestPatch) HasDistroVersion() bool`

HasDistroVersion returns a boolean if a field has been set.

### GetExtraHeader1

`func (o *RpmUpstreamRequestPatch) GetExtraHeader1() string`

GetExtraHeader1 returns the ExtraHeader1 field if non-nil, zero value otherwise.

### GetExtraHeader1Ok

`func (o *RpmUpstreamRequestPatch) GetExtraHeader1Ok() (*string, bool)`

GetExtraHeader1Ok returns a tuple with the ExtraHeader1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader1

`func (o *RpmUpstreamRequestPatch) SetExtraHeader1(v string)`

SetExtraHeader1 sets ExtraHeader1 field to given value.

### HasExtraHeader1

`func (o *RpmUpstreamRequestPatch) HasExtraHeader1() bool`

HasExtraHeader1 returns a boolean if a field has been set.

### SetExtraHeader1Nil

`func (o *RpmUpstreamRequestPatch) SetExtraHeader1Nil(b bool)`

 SetExtraHeader1Nil sets the value for ExtraHeader1 to be an explicit nil

### UnsetExtraHeader1
`func (o *RpmUpstreamRequestPatch) UnsetExtraHeader1()`

UnsetExtraHeader1 ensures that no value is present for ExtraHeader1, not even an explicit nil
### GetExtraHeader2

`func (o *RpmUpstreamRequestPatch) GetExtraHeader2() string`

GetExtraHeader2 returns the ExtraHeader2 field if non-nil, zero value otherwise.

### GetExtraHeader2Ok

`func (o *RpmUpstreamRequestPatch) GetExtraHeader2Ok() (*string, bool)`

GetExtraHeader2Ok returns a tuple with the ExtraHeader2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader2

`func (o *RpmUpstreamRequestPatch) SetExtraHeader2(v string)`

SetExtraHeader2 sets ExtraHeader2 field to given value.

### HasExtraHeader2

`func (o *RpmUpstreamRequestPatch) HasExtraHeader2() bool`

HasExtraHeader2 returns a boolean if a field has been set.

### SetExtraHeader2Nil

`func (o *RpmUpstreamRequestPatch) SetExtraHeader2Nil(b bool)`

 SetExtraHeader2Nil sets the value for ExtraHeader2 to be an explicit nil

### UnsetExtraHeader2
`func (o *RpmUpstreamRequestPatch) UnsetExtraHeader2()`

UnsetExtraHeader2 ensures that no value is present for ExtraHeader2, not even an explicit nil
### GetExtraValue1

`func (o *RpmUpstreamRequestPatch) GetExtraValue1() string`

GetExtraValue1 returns the ExtraValue1 field if non-nil, zero value otherwise.

### GetExtraValue1Ok

`func (o *RpmUpstreamRequestPatch) GetExtraValue1Ok() (*string, bool)`

GetExtraValue1Ok returns a tuple with the ExtraValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue1

`func (o *RpmUpstreamRequestPatch) SetExtraValue1(v string)`

SetExtraValue1 sets ExtraValue1 field to given value.

### HasExtraValue1

`func (o *RpmUpstreamRequestPatch) HasExtraValue1() bool`

HasExtraValue1 returns a boolean if a field has been set.

### SetExtraValue1Nil

`func (o *RpmUpstreamRequestPatch) SetExtraValue1Nil(b bool)`

 SetExtraValue1Nil sets the value for ExtraValue1 to be an explicit nil

### UnsetExtraValue1
`func (o *RpmUpstreamRequestPatch) UnsetExtraValue1()`

UnsetExtraValue1 ensures that no value is present for ExtraValue1, not even an explicit nil
### GetExtraValue2

`func (o *RpmUpstreamRequestPatch) GetExtraValue2() string`

GetExtraValue2 returns the ExtraValue2 field if non-nil, zero value otherwise.

### GetExtraValue2Ok

`func (o *RpmUpstreamRequestPatch) GetExtraValue2Ok() (*string, bool)`

GetExtraValue2Ok returns a tuple with the ExtraValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue2

`func (o *RpmUpstreamRequestPatch) SetExtraValue2(v string)`

SetExtraValue2 sets ExtraValue2 field to given value.

### HasExtraValue2

`func (o *RpmUpstreamRequestPatch) HasExtraValue2() bool`

HasExtraValue2 returns a boolean if a field has been set.

### SetExtraValue2Nil

`func (o *RpmUpstreamRequestPatch) SetExtraValue2Nil(b bool)`

 SetExtraValue2Nil sets the value for ExtraValue2 to be an explicit nil

### UnsetExtraValue2
`func (o *RpmUpstreamRequestPatch) UnsetExtraValue2()`

UnsetExtraValue2 ensures that no value is present for ExtraValue2, not even an explicit nil
### GetGpgKeyInline

`func (o *RpmUpstreamRequestPatch) GetGpgKeyInline() string`

GetGpgKeyInline returns the GpgKeyInline field if non-nil, zero value otherwise.

### GetGpgKeyInlineOk

`func (o *RpmUpstreamRequestPatch) GetGpgKeyInlineOk() (*string, bool)`

GetGpgKeyInlineOk returns a tuple with the GpgKeyInline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgKeyInline

`func (o *RpmUpstreamRequestPatch) SetGpgKeyInline(v string)`

SetGpgKeyInline sets GpgKeyInline field to given value.

### HasGpgKeyInline

`func (o *RpmUpstreamRequestPatch) HasGpgKeyInline() bool`

HasGpgKeyInline returns a boolean if a field has been set.

### SetGpgKeyInlineNil

`func (o *RpmUpstreamRequestPatch) SetGpgKeyInlineNil(b bool)`

 SetGpgKeyInlineNil sets the value for GpgKeyInline to be an explicit nil

### UnsetGpgKeyInline
`func (o *RpmUpstreamRequestPatch) UnsetGpgKeyInline()`

UnsetGpgKeyInline ensures that no value is present for GpgKeyInline, not even an explicit nil
### GetGpgKeyUrl

`func (o *RpmUpstreamRequestPatch) GetGpgKeyUrl() string`

GetGpgKeyUrl returns the GpgKeyUrl field if non-nil, zero value otherwise.

### GetGpgKeyUrlOk

`func (o *RpmUpstreamRequestPatch) GetGpgKeyUrlOk() (*string, bool)`

GetGpgKeyUrlOk returns a tuple with the GpgKeyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgKeyUrl

`func (o *RpmUpstreamRequestPatch) SetGpgKeyUrl(v string)`

SetGpgKeyUrl sets GpgKeyUrl field to given value.

### HasGpgKeyUrl

`func (o *RpmUpstreamRequestPatch) HasGpgKeyUrl() bool`

HasGpgKeyUrl returns a boolean if a field has been set.

### SetGpgKeyUrlNil

`func (o *RpmUpstreamRequestPatch) SetGpgKeyUrlNil(b bool)`

 SetGpgKeyUrlNil sets the value for GpgKeyUrl to be an explicit nil

### UnsetGpgKeyUrl
`func (o *RpmUpstreamRequestPatch) UnsetGpgKeyUrl()`

UnsetGpgKeyUrl ensures that no value is present for GpgKeyUrl, not even an explicit nil
### GetGpgVerification

`func (o *RpmUpstreamRequestPatch) GetGpgVerification() string`

GetGpgVerification returns the GpgVerification field if non-nil, zero value otherwise.

### GetGpgVerificationOk

`func (o *RpmUpstreamRequestPatch) GetGpgVerificationOk() (*string, bool)`

GetGpgVerificationOk returns a tuple with the GpgVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgVerification

`func (o *RpmUpstreamRequestPatch) SetGpgVerification(v string)`

SetGpgVerification sets GpgVerification field to given value.

### HasGpgVerification

`func (o *RpmUpstreamRequestPatch) HasGpgVerification() bool`

HasGpgVerification returns a boolean if a field has been set.

### GetIncludeSources

`func (o *RpmUpstreamRequestPatch) GetIncludeSources() bool`

GetIncludeSources returns the IncludeSources field if non-nil, zero value otherwise.

### GetIncludeSourcesOk

`func (o *RpmUpstreamRequestPatch) GetIncludeSourcesOk() (*bool, bool)`

GetIncludeSourcesOk returns a tuple with the IncludeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSources

`func (o *RpmUpstreamRequestPatch) SetIncludeSources(v bool)`

SetIncludeSources sets IncludeSources field to given value.

### HasIncludeSources

`func (o *RpmUpstreamRequestPatch) HasIncludeSources() bool`

HasIncludeSources returns a boolean if a field has been set.

### GetIsActive

`func (o *RpmUpstreamRequestPatch) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RpmUpstreamRequestPatch) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RpmUpstreamRequestPatch) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RpmUpstreamRequestPatch) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetMode

`func (o *RpmUpstreamRequestPatch) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RpmUpstreamRequestPatch) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RpmUpstreamRequestPatch) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RpmUpstreamRequestPatch) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *RpmUpstreamRequestPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RpmUpstreamRequestPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RpmUpstreamRequestPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RpmUpstreamRequestPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *RpmUpstreamRequestPatch) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RpmUpstreamRequestPatch) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RpmUpstreamRequestPatch) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RpmUpstreamRequestPatch) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetUpstreamUrl

`func (o *RpmUpstreamRequestPatch) GetUpstreamUrl() string`

GetUpstreamUrl returns the UpstreamUrl field if non-nil, zero value otherwise.

### GetUpstreamUrlOk

`func (o *RpmUpstreamRequestPatch) GetUpstreamUrlOk() (*string, bool)`

GetUpstreamUrlOk returns a tuple with the UpstreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamUrl

`func (o *RpmUpstreamRequestPatch) SetUpstreamUrl(v string)`

SetUpstreamUrl sets UpstreamUrl field to given value.

### HasUpstreamUrl

`func (o *RpmUpstreamRequestPatch) HasUpstreamUrl() bool`

HasUpstreamUrl returns a boolean if a field has been set.

### GetVerifySsl

`func (o *RpmUpstreamRequestPatch) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *RpmUpstreamRequestPatch) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *RpmUpstreamRequestPatch) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *RpmUpstreamRequestPatch) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


