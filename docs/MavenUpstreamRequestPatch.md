# MavenUpstreamRequestPatch

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
**IsActive** | Pointer to **bool** | Whether or not this upstream is active and ready for requests. | [optional] 
**Mode** | Pointer to **string** | The mode that this upstream should operate in. Upstream sources can be used to proxy resolved packages, as well as operate in a proxy/cache or cache only mode. | [optional] [default to "Proxy Only"]
**Name** | Pointer to **string** | A descriptive name for this upstream source. A shortened version of this name will be used for tagging cached packages retrieved from this upstream. | [optional] 
**Priority** | Pointer to **int64** | Upstream sources are selected for resolving requests by sequential order (1..n), followed by creation date. | [optional] 
**UpstreamUrl** | Pointer to **string** | The URL for this upstream source. This must be a fully qualified URL including any path elements required to reach the root of the repository.  | [optional] 
**VerifySsl** | Pointer to **bool** | If enabled, SSL certificates are verified when requests are made to this upstream. It&#39;s recommended to leave this enabled for all public sources to help mitigate Man-In-The-Middle (MITM) attacks. Please note this only applies to HTTPS upstreams. | [optional] 

## Methods

### NewMavenUpstreamRequestPatch

`func NewMavenUpstreamRequestPatch() *MavenUpstreamRequestPatch`

NewMavenUpstreamRequestPatch instantiates a new MavenUpstreamRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenUpstreamRequestPatchWithDefaults

`func NewMavenUpstreamRequestPatchWithDefaults() *MavenUpstreamRequestPatch`

NewMavenUpstreamRequestPatchWithDefaults instantiates a new MavenUpstreamRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *MavenUpstreamRequestPatch) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *MavenUpstreamRequestPatch) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *MavenUpstreamRequestPatch) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *MavenUpstreamRequestPatch) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAuthSecret

`func (o *MavenUpstreamRequestPatch) GetAuthSecret() string`

GetAuthSecret returns the AuthSecret field if non-nil, zero value otherwise.

### GetAuthSecretOk

`func (o *MavenUpstreamRequestPatch) GetAuthSecretOk() (*string, bool)`

GetAuthSecretOk returns a tuple with the AuthSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecret

`func (o *MavenUpstreamRequestPatch) SetAuthSecret(v string)`

SetAuthSecret sets AuthSecret field to given value.

### HasAuthSecret

`func (o *MavenUpstreamRequestPatch) HasAuthSecret() bool`

HasAuthSecret returns a boolean if a field has been set.

### SetAuthSecretNil

`func (o *MavenUpstreamRequestPatch) SetAuthSecretNil(b bool)`

 SetAuthSecretNil sets the value for AuthSecret to be an explicit nil

### UnsetAuthSecret
`func (o *MavenUpstreamRequestPatch) UnsetAuthSecret()`

UnsetAuthSecret ensures that no value is present for AuthSecret, not even an explicit nil
### GetAuthUsername

`func (o *MavenUpstreamRequestPatch) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *MavenUpstreamRequestPatch) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *MavenUpstreamRequestPatch) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *MavenUpstreamRequestPatch) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### SetAuthUsernameNil

`func (o *MavenUpstreamRequestPatch) SetAuthUsernameNil(b bool)`

 SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil

### UnsetAuthUsername
`func (o *MavenUpstreamRequestPatch) UnsetAuthUsername()`

UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
### GetExtraHeader1

`func (o *MavenUpstreamRequestPatch) GetExtraHeader1() string`

GetExtraHeader1 returns the ExtraHeader1 field if non-nil, zero value otherwise.

### GetExtraHeader1Ok

`func (o *MavenUpstreamRequestPatch) GetExtraHeader1Ok() (*string, bool)`

GetExtraHeader1Ok returns a tuple with the ExtraHeader1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader1

`func (o *MavenUpstreamRequestPatch) SetExtraHeader1(v string)`

SetExtraHeader1 sets ExtraHeader1 field to given value.

### HasExtraHeader1

`func (o *MavenUpstreamRequestPatch) HasExtraHeader1() bool`

HasExtraHeader1 returns a boolean if a field has been set.

### SetExtraHeader1Nil

`func (o *MavenUpstreamRequestPatch) SetExtraHeader1Nil(b bool)`

 SetExtraHeader1Nil sets the value for ExtraHeader1 to be an explicit nil

### UnsetExtraHeader1
`func (o *MavenUpstreamRequestPatch) UnsetExtraHeader1()`

UnsetExtraHeader1 ensures that no value is present for ExtraHeader1, not even an explicit nil
### GetExtraHeader2

`func (o *MavenUpstreamRequestPatch) GetExtraHeader2() string`

GetExtraHeader2 returns the ExtraHeader2 field if non-nil, zero value otherwise.

### GetExtraHeader2Ok

`func (o *MavenUpstreamRequestPatch) GetExtraHeader2Ok() (*string, bool)`

GetExtraHeader2Ok returns a tuple with the ExtraHeader2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader2

`func (o *MavenUpstreamRequestPatch) SetExtraHeader2(v string)`

SetExtraHeader2 sets ExtraHeader2 field to given value.

### HasExtraHeader2

`func (o *MavenUpstreamRequestPatch) HasExtraHeader2() bool`

HasExtraHeader2 returns a boolean if a field has been set.

### SetExtraHeader2Nil

`func (o *MavenUpstreamRequestPatch) SetExtraHeader2Nil(b bool)`

 SetExtraHeader2Nil sets the value for ExtraHeader2 to be an explicit nil

### UnsetExtraHeader2
`func (o *MavenUpstreamRequestPatch) UnsetExtraHeader2()`

UnsetExtraHeader2 ensures that no value is present for ExtraHeader2, not even an explicit nil
### GetExtraValue1

`func (o *MavenUpstreamRequestPatch) GetExtraValue1() string`

GetExtraValue1 returns the ExtraValue1 field if non-nil, zero value otherwise.

### GetExtraValue1Ok

`func (o *MavenUpstreamRequestPatch) GetExtraValue1Ok() (*string, bool)`

GetExtraValue1Ok returns a tuple with the ExtraValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue1

`func (o *MavenUpstreamRequestPatch) SetExtraValue1(v string)`

SetExtraValue1 sets ExtraValue1 field to given value.

### HasExtraValue1

`func (o *MavenUpstreamRequestPatch) HasExtraValue1() bool`

HasExtraValue1 returns a boolean if a field has been set.

### SetExtraValue1Nil

`func (o *MavenUpstreamRequestPatch) SetExtraValue1Nil(b bool)`

 SetExtraValue1Nil sets the value for ExtraValue1 to be an explicit nil

### UnsetExtraValue1
`func (o *MavenUpstreamRequestPatch) UnsetExtraValue1()`

UnsetExtraValue1 ensures that no value is present for ExtraValue1, not even an explicit nil
### GetExtraValue2

`func (o *MavenUpstreamRequestPatch) GetExtraValue2() string`

GetExtraValue2 returns the ExtraValue2 field if non-nil, zero value otherwise.

### GetExtraValue2Ok

`func (o *MavenUpstreamRequestPatch) GetExtraValue2Ok() (*string, bool)`

GetExtraValue2Ok returns a tuple with the ExtraValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue2

`func (o *MavenUpstreamRequestPatch) SetExtraValue2(v string)`

SetExtraValue2 sets ExtraValue2 field to given value.

### HasExtraValue2

`func (o *MavenUpstreamRequestPatch) HasExtraValue2() bool`

HasExtraValue2 returns a boolean if a field has been set.

### SetExtraValue2Nil

`func (o *MavenUpstreamRequestPatch) SetExtraValue2Nil(b bool)`

 SetExtraValue2Nil sets the value for ExtraValue2 to be an explicit nil

### UnsetExtraValue2
`func (o *MavenUpstreamRequestPatch) UnsetExtraValue2()`

UnsetExtraValue2 ensures that no value is present for ExtraValue2, not even an explicit nil
### GetIsActive

`func (o *MavenUpstreamRequestPatch) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *MavenUpstreamRequestPatch) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *MavenUpstreamRequestPatch) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *MavenUpstreamRequestPatch) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetMode

`func (o *MavenUpstreamRequestPatch) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MavenUpstreamRequestPatch) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MavenUpstreamRequestPatch) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *MavenUpstreamRequestPatch) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *MavenUpstreamRequestPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MavenUpstreamRequestPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MavenUpstreamRequestPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MavenUpstreamRequestPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *MavenUpstreamRequestPatch) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MavenUpstreamRequestPatch) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MavenUpstreamRequestPatch) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MavenUpstreamRequestPatch) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetUpstreamUrl

`func (o *MavenUpstreamRequestPatch) GetUpstreamUrl() string`

GetUpstreamUrl returns the UpstreamUrl field if non-nil, zero value otherwise.

### GetUpstreamUrlOk

`func (o *MavenUpstreamRequestPatch) GetUpstreamUrlOk() (*string, bool)`

GetUpstreamUrlOk returns a tuple with the UpstreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamUrl

`func (o *MavenUpstreamRequestPatch) SetUpstreamUrl(v string)`

SetUpstreamUrl sets UpstreamUrl field to given value.

### HasUpstreamUrl

`func (o *MavenUpstreamRequestPatch) HasUpstreamUrl() bool`

HasUpstreamUrl returns a boolean if a field has been set.

### GetVerifySsl

`func (o *MavenUpstreamRequestPatch) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *MavenUpstreamRequestPatch) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *MavenUpstreamRequestPatch) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *MavenUpstreamRequestPatch) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


