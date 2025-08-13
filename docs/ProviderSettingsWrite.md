# ProviderSettingsWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claims** | **map[string]interface{}** | The set of claims that any received tokens from the provider must contain to authenticate as the configured service account. | 
**DynamicMappings** | Pointer to [**[]DynamicMapping**](DynamicMapping.md) | The dynamic mappings of &#x60;mapping_claim&#x60; values to service accounts. Cannot be provided if &#x60;service_accounts&#x60; is also set.  Note: This field and the dynamic mappings feature are still in early access. Breaking changes are possible as we receive feedback on this feature. | [optional] 
**Enabled** | **bool** | Whether the provider settings should be used for incoming OIDC requests. | 
**MappingClaim** | Pointer to **NullableString** | The OIDC claim to use for mapping to service accounts in dynamic_mappings. Cannot be provided if &#x60;service_accounts&#x60; is also set.  Note: This field and the dynamic mappings feature are still in early access. Breaking changes are possible as we receive feedback on this feature. | [optional] 
**Name** | **string** | The name of the provider settings are being configured for | 
**ProviderUrl** | **string** | The URL from the provider that serves as the base for the OpenID configuration. For example, if the OpenID configuration is available at https://token.actions.githubusercontent.com/.well-known/openid-configuration, the provider URL would be https://token.actions.githubusercontent.com/ | 
**ServiceAccounts** | Pointer to **[]string** | The service accounts associated with these provider settings. Cannot be provided if &#x60;mapping_claim&#x60; or &#x60;dynamic_mappings&#x60; are specified. | [optional] 
**Slug** | Pointer to **string** | The slug of the provider settings | [optional] [readonly] 
**SlugPerm** | Pointer to **string** | The unique, immutable identifier of the provider settings. | [optional] [readonly] 

## Methods

### NewProviderSettingsWrite

`func NewProviderSettingsWrite(claims map[string]interface{}, enabled bool, name string, providerUrl string, ) *ProviderSettingsWrite`

NewProviderSettingsWrite instantiates a new ProviderSettingsWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSettingsWriteWithDefaults

`func NewProviderSettingsWriteWithDefaults() *ProviderSettingsWrite`

NewProviderSettingsWriteWithDefaults instantiates a new ProviderSettingsWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaims

`func (o *ProviderSettingsWrite) GetClaims() map[string]interface{}`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *ProviderSettingsWrite) GetClaimsOk() (*map[string]interface{}, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *ProviderSettingsWrite) SetClaims(v map[string]interface{})`

SetClaims sets Claims field to given value.


### GetDynamicMappings

`func (o *ProviderSettingsWrite) GetDynamicMappings() []DynamicMapping`

GetDynamicMappings returns the DynamicMappings field if non-nil, zero value otherwise.

### GetDynamicMappingsOk

`func (o *ProviderSettingsWrite) GetDynamicMappingsOk() (*[]DynamicMapping, bool)`

GetDynamicMappingsOk returns a tuple with the DynamicMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicMappings

`func (o *ProviderSettingsWrite) SetDynamicMappings(v []DynamicMapping)`

SetDynamicMappings sets DynamicMappings field to given value.

### HasDynamicMappings

`func (o *ProviderSettingsWrite) HasDynamicMappings() bool`

HasDynamicMappings returns a boolean if a field has been set.

### GetEnabled

`func (o *ProviderSettingsWrite) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProviderSettingsWrite) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProviderSettingsWrite) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMappingClaim

`func (o *ProviderSettingsWrite) GetMappingClaim() string`

GetMappingClaim returns the MappingClaim field if non-nil, zero value otherwise.

### GetMappingClaimOk

`func (o *ProviderSettingsWrite) GetMappingClaimOk() (*string, bool)`

GetMappingClaimOk returns a tuple with the MappingClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingClaim

`func (o *ProviderSettingsWrite) SetMappingClaim(v string)`

SetMappingClaim sets MappingClaim field to given value.

### HasMappingClaim

`func (o *ProviderSettingsWrite) HasMappingClaim() bool`

HasMappingClaim returns a boolean if a field has been set.

### SetMappingClaimNil

`func (o *ProviderSettingsWrite) SetMappingClaimNil(b bool)`

 SetMappingClaimNil sets the value for MappingClaim to be an explicit nil

### UnsetMappingClaim
`func (o *ProviderSettingsWrite) UnsetMappingClaim()`

UnsetMappingClaim ensures that no value is present for MappingClaim, not even an explicit nil
### GetName

`func (o *ProviderSettingsWrite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderSettingsWrite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderSettingsWrite) SetName(v string)`

SetName sets Name field to given value.


### GetProviderUrl

`func (o *ProviderSettingsWrite) GetProviderUrl() string`

GetProviderUrl returns the ProviderUrl field if non-nil, zero value otherwise.

### GetProviderUrlOk

`func (o *ProviderSettingsWrite) GetProviderUrlOk() (*string, bool)`

GetProviderUrlOk returns a tuple with the ProviderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrl

`func (o *ProviderSettingsWrite) SetProviderUrl(v string)`

SetProviderUrl sets ProviderUrl field to given value.


### GetServiceAccounts

`func (o *ProviderSettingsWrite) GetServiceAccounts() []string`

GetServiceAccounts returns the ServiceAccounts field if non-nil, zero value otherwise.

### GetServiceAccountsOk

`func (o *ProviderSettingsWrite) GetServiceAccountsOk() (*[]string, bool)`

GetServiceAccountsOk returns a tuple with the ServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccounts

`func (o *ProviderSettingsWrite) SetServiceAccounts(v []string)`

SetServiceAccounts sets ServiceAccounts field to given value.

### HasServiceAccounts

`func (o *ProviderSettingsWrite) HasServiceAccounts() bool`

HasServiceAccounts returns a boolean if a field has been set.

### GetSlug

`func (o *ProviderSettingsWrite) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProviderSettingsWrite) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProviderSettingsWrite) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProviderSettingsWrite) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSlugPerm

`func (o *ProviderSettingsWrite) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *ProviderSettingsWrite) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *ProviderSettingsWrite) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *ProviderSettingsWrite) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


