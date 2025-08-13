# ProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claims** | **map[string]interface{}** | The set of claims that any received tokens from the provider must contain to authenticate as the configured service account. | 
**Enabled** | **bool** | Whether the provider settings should be used for incoming OIDC requests. | 
**MappingClaim** | Pointer to **NullableString** | The OIDC claim to use for mapping to service accounts in dynamic_mappings.  Note: This field and the dynamic mappings feature are still in early access. Breaking changes are possible as we receive feedback on this feature. | [optional] 
**Name** | **string** | The name of the provider settings are being configured for | 
**ProviderUrl** | **string** | The URL from the provider that serves as the base for the OpenID configuration. For example, if the OpenID configuration is available at https://token.actions.githubusercontent.com/.well-known/openid-configuration, the provider URL would be https://token.actions.githubusercontent.com/ | 
**ServiceAccounts** | Pointer to **[]string** | The service accounts associated with these provider settings. | [optional] 
**Slug** | Pointer to **string** | The slug of the provider settings | [optional] [readonly] 
**SlugPerm** | Pointer to **string** | The unique, immutable identifier of the provider settings. | [optional] [readonly] 

## Methods

### NewProviderSettings

`func NewProviderSettings(claims map[string]interface{}, enabled bool, name string, providerUrl string, ) *ProviderSettings`

NewProviderSettings instantiates a new ProviderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSettingsWithDefaults

`func NewProviderSettingsWithDefaults() *ProviderSettings`

NewProviderSettingsWithDefaults instantiates a new ProviderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaims

`func (o *ProviderSettings) GetClaims() map[string]interface{}`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *ProviderSettings) GetClaimsOk() (*map[string]interface{}, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *ProviderSettings) SetClaims(v map[string]interface{})`

SetClaims sets Claims field to given value.


### GetEnabled

`func (o *ProviderSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProviderSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProviderSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMappingClaim

`func (o *ProviderSettings) GetMappingClaim() string`

GetMappingClaim returns the MappingClaim field if non-nil, zero value otherwise.

### GetMappingClaimOk

`func (o *ProviderSettings) GetMappingClaimOk() (*string, bool)`

GetMappingClaimOk returns a tuple with the MappingClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingClaim

`func (o *ProviderSettings) SetMappingClaim(v string)`

SetMappingClaim sets MappingClaim field to given value.

### HasMappingClaim

`func (o *ProviderSettings) HasMappingClaim() bool`

HasMappingClaim returns a boolean if a field has been set.

### SetMappingClaimNil

`func (o *ProviderSettings) SetMappingClaimNil(b bool)`

 SetMappingClaimNil sets the value for MappingClaim to be an explicit nil

### UnsetMappingClaim
`func (o *ProviderSettings) UnsetMappingClaim()`

UnsetMappingClaim ensures that no value is present for MappingClaim, not even an explicit nil
### GetName

`func (o *ProviderSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderSettings) SetName(v string)`

SetName sets Name field to given value.


### GetProviderUrl

`func (o *ProviderSettings) GetProviderUrl() string`

GetProviderUrl returns the ProviderUrl field if non-nil, zero value otherwise.

### GetProviderUrlOk

`func (o *ProviderSettings) GetProviderUrlOk() (*string, bool)`

GetProviderUrlOk returns a tuple with the ProviderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrl

`func (o *ProviderSettings) SetProviderUrl(v string)`

SetProviderUrl sets ProviderUrl field to given value.


### GetServiceAccounts

`func (o *ProviderSettings) GetServiceAccounts() []string`

GetServiceAccounts returns the ServiceAccounts field if non-nil, zero value otherwise.

### GetServiceAccountsOk

`func (o *ProviderSettings) GetServiceAccountsOk() (*[]string, bool)`

GetServiceAccountsOk returns a tuple with the ServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccounts

`func (o *ProviderSettings) SetServiceAccounts(v []string)`

SetServiceAccounts sets ServiceAccounts field to given value.

### HasServiceAccounts

`func (o *ProviderSettings) HasServiceAccounts() bool`

HasServiceAccounts returns a boolean if a field has been set.

### GetSlug

`func (o *ProviderSettings) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProviderSettings) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProviderSettings) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProviderSettings) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSlugPerm

`func (o *ProviderSettings) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *ProviderSettings) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *ProviderSettings) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *ProviderSettings) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


