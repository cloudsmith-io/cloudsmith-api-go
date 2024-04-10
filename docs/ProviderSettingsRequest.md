# ProviderSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claims** | **map[string]interface{}** | The set of claims that any received tokens from the provider must contain to authenticate as the configured service account. | 
**Enabled** | **bool** | Whether the provider settings should be used for incoming OIDC requests. | 
**Name** | **string** | The name of the provider settings are being configured for | 
**ProviderUrl** | **string** | The URL from the provider that serves as the base for the OpenID configuration. For example, if the OpenID configuration is available at https://token.actions.githubusercontent.com/.well-known/openid-configuration, the provider URL would be https://token.actions.githubusercontent.com/ | 
**ServiceAccounts** | **[]string** | The service accounts associated with these provider settings | 

## Methods

### NewProviderSettingsRequest

`func NewProviderSettingsRequest(claims map[string]interface{}, enabled bool, name string, providerUrl string, serviceAccounts []string, ) *ProviderSettingsRequest`

NewProviderSettingsRequest instantiates a new ProviderSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSettingsRequestWithDefaults

`func NewProviderSettingsRequestWithDefaults() *ProviderSettingsRequest`

NewProviderSettingsRequestWithDefaults instantiates a new ProviderSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaims

`func (o *ProviderSettingsRequest) GetClaims() map[string]interface{}`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *ProviderSettingsRequest) GetClaimsOk() (*map[string]interface{}, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *ProviderSettingsRequest) SetClaims(v map[string]interface{})`

SetClaims sets Claims field to given value.


### GetEnabled

`func (o *ProviderSettingsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProviderSettingsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProviderSettingsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetName

`func (o *ProviderSettingsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderSettingsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderSettingsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProviderUrl

`func (o *ProviderSettingsRequest) GetProviderUrl() string`

GetProviderUrl returns the ProviderUrl field if non-nil, zero value otherwise.

### GetProviderUrlOk

`func (o *ProviderSettingsRequest) GetProviderUrlOk() (*string, bool)`

GetProviderUrlOk returns a tuple with the ProviderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrl

`func (o *ProviderSettingsRequest) SetProviderUrl(v string)`

SetProviderUrl sets ProviderUrl field to given value.


### GetServiceAccounts

`func (o *ProviderSettingsRequest) GetServiceAccounts() []string`

GetServiceAccounts returns the ServiceAccounts field if non-nil, zero value otherwise.

### GetServiceAccountsOk

`func (o *ProviderSettingsRequest) GetServiceAccountsOk() (*[]string, bool)`

GetServiceAccountsOk returns a tuple with the ServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccounts

`func (o *ProviderSettingsRequest) SetServiceAccounts(v []string)`

SetServiceAccounts sets ServiceAccounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


