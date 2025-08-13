# ProviderSettingsWriteRequestPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claims** | Pointer to **map[string]interface{}** | The set of claims that any received tokens from the provider must contain to authenticate as the configured service account. | [optional] 
**DynamicMappings** | Pointer to [**[]DynamicMapping**](DynamicMapping.md) | The dynamic mappings of &#x60;mapping_claim&#x60; values to service accounts. Cannot be provided if &#x60;service_accounts&#x60; is also set.  Note: This field and the dynamic mappings feature are still in early access. Breaking changes are possible as we receive feedback on this feature. | [optional] 
**Enabled** | Pointer to **bool** | Whether the provider settings should be used for incoming OIDC requests. | [optional] 
**MappingClaim** | Pointer to **NullableString** | The OIDC claim to use for mapping to service accounts in dynamic_mappings. Cannot be provided if &#x60;service_accounts&#x60; is also set.  Note: This field and the dynamic mappings feature are still in early access. Breaking changes are possible as we receive feedback on this feature. | [optional] 
**Name** | Pointer to **string** | The name of the provider settings are being configured for | [optional] 
**ProviderUrl** | Pointer to **string** | The URL from the provider that serves as the base for the OpenID configuration. For example, if the OpenID configuration is available at https://token.actions.githubusercontent.com/.well-known/openid-configuration, the provider URL would be https://token.actions.githubusercontent.com/ | [optional] 
**ServiceAccounts** | Pointer to **[]string** | The service accounts associated with these provider settings. Cannot be provided if &#x60;mapping_claim&#x60; or &#x60;dynamic_mappings&#x60; are specified. | [optional] 

## Methods

### NewProviderSettingsWriteRequestPatch

`func NewProviderSettingsWriteRequestPatch() *ProviderSettingsWriteRequestPatch`

NewProviderSettingsWriteRequestPatch instantiates a new ProviderSettingsWriteRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSettingsWriteRequestPatchWithDefaults

`func NewProviderSettingsWriteRequestPatchWithDefaults() *ProviderSettingsWriteRequestPatch`

NewProviderSettingsWriteRequestPatchWithDefaults instantiates a new ProviderSettingsWriteRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaims

`func (o *ProviderSettingsWriteRequestPatch) GetClaims() map[string]interface{}`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *ProviderSettingsWriteRequestPatch) GetClaimsOk() (*map[string]interface{}, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *ProviderSettingsWriteRequestPatch) SetClaims(v map[string]interface{})`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *ProviderSettingsWriteRequestPatch) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetDynamicMappings

`func (o *ProviderSettingsWriteRequestPatch) GetDynamicMappings() []DynamicMapping`

GetDynamicMappings returns the DynamicMappings field if non-nil, zero value otherwise.

### GetDynamicMappingsOk

`func (o *ProviderSettingsWriteRequestPatch) GetDynamicMappingsOk() (*[]DynamicMapping, bool)`

GetDynamicMappingsOk returns a tuple with the DynamicMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicMappings

`func (o *ProviderSettingsWriteRequestPatch) SetDynamicMappings(v []DynamicMapping)`

SetDynamicMappings sets DynamicMappings field to given value.

### HasDynamicMappings

`func (o *ProviderSettingsWriteRequestPatch) HasDynamicMappings() bool`

HasDynamicMappings returns a boolean if a field has been set.

### GetEnabled

`func (o *ProviderSettingsWriteRequestPatch) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProviderSettingsWriteRequestPatch) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProviderSettingsWriteRequestPatch) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ProviderSettingsWriteRequestPatch) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMappingClaim

`func (o *ProviderSettingsWriteRequestPatch) GetMappingClaim() string`

GetMappingClaim returns the MappingClaim field if non-nil, zero value otherwise.

### GetMappingClaimOk

`func (o *ProviderSettingsWriteRequestPatch) GetMappingClaimOk() (*string, bool)`

GetMappingClaimOk returns a tuple with the MappingClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingClaim

`func (o *ProviderSettingsWriteRequestPatch) SetMappingClaim(v string)`

SetMappingClaim sets MappingClaim field to given value.

### HasMappingClaim

`func (o *ProviderSettingsWriteRequestPatch) HasMappingClaim() bool`

HasMappingClaim returns a boolean if a field has been set.

### SetMappingClaimNil

`func (o *ProviderSettingsWriteRequestPatch) SetMappingClaimNil(b bool)`

 SetMappingClaimNil sets the value for MappingClaim to be an explicit nil

### UnsetMappingClaim
`func (o *ProviderSettingsWriteRequestPatch) UnsetMappingClaim()`

UnsetMappingClaim ensures that no value is present for MappingClaim, not even an explicit nil
### GetName

`func (o *ProviderSettingsWriteRequestPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderSettingsWriteRequestPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderSettingsWriteRequestPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderSettingsWriteRequestPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProviderUrl

`func (o *ProviderSettingsWriteRequestPatch) GetProviderUrl() string`

GetProviderUrl returns the ProviderUrl field if non-nil, zero value otherwise.

### GetProviderUrlOk

`func (o *ProviderSettingsWriteRequestPatch) GetProviderUrlOk() (*string, bool)`

GetProviderUrlOk returns a tuple with the ProviderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrl

`func (o *ProviderSettingsWriteRequestPatch) SetProviderUrl(v string)`

SetProviderUrl sets ProviderUrl field to given value.

### HasProviderUrl

`func (o *ProviderSettingsWriteRequestPatch) HasProviderUrl() bool`

HasProviderUrl returns a boolean if a field has been set.

### GetServiceAccounts

`func (o *ProviderSettingsWriteRequestPatch) GetServiceAccounts() []string`

GetServiceAccounts returns the ServiceAccounts field if non-nil, zero value otherwise.

### GetServiceAccountsOk

`func (o *ProviderSettingsWriteRequestPatch) GetServiceAccountsOk() (*[]string, bool)`

GetServiceAccountsOk returns a tuple with the ServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccounts

`func (o *ProviderSettingsWriteRequestPatch) SetServiceAccounts(v []string)`

SetServiceAccounts sets ServiceAccounts field to given value.

### HasServiceAccounts

`func (o *ProviderSettingsWriteRequestPatch) HasServiceAccounts() bool`

HasServiceAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


