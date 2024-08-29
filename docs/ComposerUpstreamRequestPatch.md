# ComposerUpstreamRequestPatch

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

### NewComposerUpstreamRequestPatch

`func NewComposerUpstreamRequestPatch() *ComposerUpstreamRequestPatch`

NewComposerUpstreamRequestPatch instantiates a new ComposerUpstreamRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComposerUpstreamRequestPatchWithDefaults

`func NewComposerUpstreamRequestPatchWithDefaults() *ComposerUpstreamRequestPatch`

NewComposerUpstreamRequestPatchWithDefaults instantiates a new ComposerUpstreamRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *ComposerUpstreamRequestPatch) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *ComposerUpstreamRequestPatch) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *ComposerUpstreamRequestPatch) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *ComposerUpstreamRequestPatch) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAuthSecret

`func (o *ComposerUpstreamRequestPatch) GetAuthSecret() string`

GetAuthSecret returns the AuthSecret field if non-nil, zero value otherwise.

### GetAuthSecretOk

`func (o *ComposerUpstreamRequestPatch) GetAuthSecretOk() (*string, bool)`

GetAuthSecretOk returns a tuple with the AuthSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecret

`func (o *ComposerUpstreamRequestPatch) SetAuthSecret(v string)`

SetAuthSecret sets AuthSecret field to given value.

### HasAuthSecret

`func (o *ComposerUpstreamRequestPatch) HasAuthSecret() bool`

HasAuthSecret returns a boolean if a field has been set.

### SetAuthSecretNil

`func (o *ComposerUpstreamRequestPatch) SetAuthSecretNil(b bool)`

 SetAuthSecretNil sets the value for AuthSecret to be an explicit nil

### UnsetAuthSecret
`func (o *ComposerUpstreamRequestPatch) UnsetAuthSecret()`

UnsetAuthSecret ensures that no value is present for AuthSecret, not even an explicit nil
### GetAuthUsername

`func (o *ComposerUpstreamRequestPatch) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *ComposerUpstreamRequestPatch) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *ComposerUpstreamRequestPatch) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *ComposerUpstreamRequestPatch) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### SetAuthUsernameNil

`func (o *ComposerUpstreamRequestPatch) SetAuthUsernameNil(b bool)`

 SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil

### UnsetAuthUsername
`func (o *ComposerUpstreamRequestPatch) UnsetAuthUsername()`

UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
### GetExtraHeader1

`func (o *ComposerUpstreamRequestPatch) GetExtraHeader1() string`

GetExtraHeader1 returns the ExtraHeader1 field if non-nil, zero value otherwise.

### GetExtraHeader1Ok

`func (o *ComposerUpstreamRequestPatch) GetExtraHeader1Ok() (*string, bool)`

GetExtraHeader1Ok returns a tuple with the ExtraHeader1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader1

`func (o *ComposerUpstreamRequestPatch) SetExtraHeader1(v string)`

SetExtraHeader1 sets ExtraHeader1 field to given value.

### HasExtraHeader1

`func (o *ComposerUpstreamRequestPatch) HasExtraHeader1() bool`

HasExtraHeader1 returns a boolean if a field has been set.

### SetExtraHeader1Nil

`func (o *ComposerUpstreamRequestPatch) SetExtraHeader1Nil(b bool)`

 SetExtraHeader1Nil sets the value for ExtraHeader1 to be an explicit nil

### UnsetExtraHeader1
`func (o *ComposerUpstreamRequestPatch) UnsetExtraHeader1()`

UnsetExtraHeader1 ensures that no value is present for ExtraHeader1, not even an explicit nil
### GetExtraHeader2

`func (o *ComposerUpstreamRequestPatch) GetExtraHeader2() string`

GetExtraHeader2 returns the ExtraHeader2 field if non-nil, zero value otherwise.

### GetExtraHeader2Ok

`func (o *ComposerUpstreamRequestPatch) GetExtraHeader2Ok() (*string, bool)`

GetExtraHeader2Ok returns a tuple with the ExtraHeader2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeader2

`func (o *ComposerUpstreamRequestPatch) SetExtraHeader2(v string)`

SetExtraHeader2 sets ExtraHeader2 field to given value.

### HasExtraHeader2

`func (o *ComposerUpstreamRequestPatch) HasExtraHeader2() bool`

HasExtraHeader2 returns a boolean if a field has been set.

### SetExtraHeader2Nil

`func (o *ComposerUpstreamRequestPatch) SetExtraHeader2Nil(b bool)`

 SetExtraHeader2Nil sets the value for ExtraHeader2 to be an explicit nil

### UnsetExtraHeader2
`func (o *ComposerUpstreamRequestPatch) UnsetExtraHeader2()`

UnsetExtraHeader2 ensures that no value is present for ExtraHeader2, not even an explicit nil
### GetExtraValue1

`func (o *ComposerUpstreamRequestPatch) GetExtraValue1() string`

GetExtraValue1 returns the ExtraValue1 field if non-nil, zero value otherwise.

### GetExtraValue1Ok

`func (o *ComposerUpstreamRequestPatch) GetExtraValue1Ok() (*string, bool)`

GetExtraValue1Ok returns a tuple with the ExtraValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue1

`func (o *ComposerUpstreamRequestPatch) SetExtraValue1(v string)`

SetExtraValue1 sets ExtraValue1 field to given value.

### HasExtraValue1

`func (o *ComposerUpstreamRequestPatch) HasExtraValue1() bool`

HasExtraValue1 returns a boolean if a field has been set.

### SetExtraValue1Nil

`func (o *ComposerUpstreamRequestPatch) SetExtraValue1Nil(b bool)`

 SetExtraValue1Nil sets the value for ExtraValue1 to be an explicit nil

### UnsetExtraValue1
`func (o *ComposerUpstreamRequestPatch) UnsetExtraValue1()`

UnsetExtraValue1 ensures that no value is present for ExtraValue1, not even an explicit nil
### GetExtraValue2

`func (o *ComposerUpstreamRequestPatch) GetExtraValue2() string`

GetExtraValue2 returns the ExtraValue2 field if non-nil, zero value otherwise.

### GetExtraValue2Ok

`func (o *ComposerUpstreamRequestPatch) GetExtraValue2Ok() (*string, bool)`

GetExtraValue2Ok returns a tuple with the ExtraValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraValue2

`func (o *ComposerUpstreamRequestPatch) SetExtraValue2(v string)`

SetExtraValue2 sets ExtraValue2 field to given value.

### HasExtraValue2

`func (o *ComposerUpstreamRequestPatch) HasExtraValue2() bool`

HasExtraValue2 returns a boolean if a field has been set.

### SetExtraValue2Nil

`func (o *ComposerUpstreamRequestPatch) SetExtraValue2Nil(b bool)`

 SetExtraValue2Nil sets the value for ExtraValue2 to be an explicit nil

### UnsetExtraValue2
`func (o *ComposerUpstreamRequestPatch) UnsetExtraValue2()`

UnsetExtraValue2 ensures that no value is present for ExtraValue2, not even an explicit nil
### GetIsActive

`func (o *ComposerUpstreamRequestPatch) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ComposerUpstreamRequestPatch) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ComposerUpstreamRequestPatch) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ComposerUpstreamRequestPatch) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetMode

`func (o *ComposerUpstreamRequestPatch) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ComposerUpstreamRequestPatch) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ComposerUpstreamRequestPatch) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ComposerUpstreamRequestPatch) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *ComposerUpstreamRequestPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComposerUpstreamRequestPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComposerUpstreamRequestPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComposerUpstreamRequestPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *ComposerUpstreamRequestPatch) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ComposerUpstreamRequestPatch) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ComposerUpstreamRequestPatch) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ComposerUpstreamRequestPatch) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetUpstreamUrl

`func (o *ComposerUpstreamRequestPatch) GetUpstreamUrl() string`

GetUpstreamUrl returns the UpstreamUrl field if non-nil, zero value otherwise.

### GetUpstreamUrlOk

`func (o *ComposerUpstreamRequestPatch) GetUpstreamUrlOk() (*string, bool)`

GetUpstreamUrlOk returns a tuple with the UpstreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamUrl

`func (o *ComposerUpstreamRequestPatch) SetUpstreamUrl(v string)`

SetUpstreamUrl sets UpstreamUrl field to given value.

### HasUpstreamUrl

`func (o *ComposerUpstreamRequestPatch) HasUpstreamUrl() bool`

HasUpstreamUrl returns a boolean if a field has been set.

### GetVerifySsl

`func (o *ComposerUpstreamRequestPatch) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *ComposerUpstreamRequestPatch) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *ComposerUpstreamRequestPatch) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *ComposerUpstreamRequestPatch) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


