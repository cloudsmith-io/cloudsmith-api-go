# ProviderSettingsRequestPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claims** | Pointer to **map[string]interface{}** | The set of claims that any received tokens from the provider must contain to authenticate as the configured service account. | [optional] 
**Enabled** | Pointer to **bool** | Whether the provider settings should be used for incoming OIDC requests. | [optional] 
**Name** | Pointer to **string** | The name of the provider settings are being configured for | [optional] 
**ProviderUrl** | Pointer to **string** | The URL from the provider that serves as the base for the OpenID configuration. For example, if the OpenID configuration is available at https://token.actions.githubusercontent.com/.well-known/openid-configuration, the provider URL would be https://token.actions.githubusercontent.com/ | [optional] 
**ServiceAccounts** | Pointer to **[]string** | The service accounts associated with these provider settings | [optional] 

## Methods

### NewProviderSettingsRequestPatch

`func NewProviderSettingsRequestPatch() *ProviderSettingsRequestPatch`

NewProviderSettingsRequestPatch instantiates a new ProviderSettingsRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSettingsRequestPatchWithDefaults

`func NewProviderSettingsRequestPatchWithDefaults() *ProviderSettingsRequestPatch`

NewProviderSettingsRequestPatchWithDefaults instantiates a new ProviderSettingsRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaims

`func (o *ProviderSettingsRequestPatch) GetClaims() map[string]interface{}`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *ProviderSettingsRequestPatch) GetClaimsOk() (*map[string]interface{}, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *ProviderSettingsRequestPatch) SetClaims(v map[string]interface{})`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *ProviderSettingsRequestPatch) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetEnabled

`func (o *ProviderSettingsRequestPatch) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProviderSettingsRequestPatch) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProviderSettingsRequestPatch) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ProviderSettingsRequestPatch) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *ProviderSettingsRequestPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderSettingsRequestPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderSettingsRequestPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderSettingsRequestPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProviderUrl

`func (o *ProviderSettingsRequestPatch) GetProviderUrl() string`

GetProviderUrl returns the ProviderUrl field if non-nil, zero value otherwise.

### GetProviderUrlOk

`func (o *ProviderSettingsRequestPatch) GetProviderUrlOk() (*string, bool)`

GetProviderUrlOk returns a tuple with the ProviderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrl

`func (o *ProviderSettingsRequestPatch) SetProviderUrl(v string)`

SetProviderUrl sets ProviderUrl field to given value.

### HasProviderUrl

`func (o *ProviderSettingsRequestPatch) HasProviderUrl() bool`

HasProviderUrl returns a boolean if a field has been set.

### GetServiceAccounts

`func (o *ProviderSettingsRequestPatch) GetServiceAccounts() []string`

GetServiceAccounts returns the ServiceAccounts field if non-nil, zero value otherwise.

### GetServiceAccountsOk

`func (o *ProviderSettingsRequestPatch) GetServiceAccountsOk() (*[]string, bool)`

GetServiceAccountsOk returns a tuple with the ServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccounts

`func (o *ProviderSettingsRequestPatch) SetServiceAccounts(v []string)`

SetServiceAccounts sets ServiceAccounts field to given value.

### HasServiceAccounts

`func (o *ProviderSettingsRequestPatch) HasServiceAccounts() bool`

HasServiceAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


