# OrganizationApiKeyRuleRequestPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | Whether this rule is currently active and enforced. | [optional] 
**MaxAgeHours** | Pointer to **NullableInt64** | The maximum permitted age of an API key for use in the organization. API keys older than this will no longer have access until they are refreshed. To disable the rule unset this value. | [optional] 
**RuleType** | Pointer to **string** | Specify which account types this rule applies to: all accounts (deprecated — prefer separate service/user rules), service accounts only, or user accounts only. | [optional] 

## Methods

### NewOrganizationApiKeyRuleRequestPatch

`func NewOrganizationApiKeyRuleRequestPatch() *OrganizationApiKeyRuleRequestPatch`

NewOrganizationApiKeyRuleRequestPatch instantiates a new OrganizationApiKeyRuleRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationApiKeyRuleRequestPatchWithDefaults

`func NewOrganizationApiKeyRuleRequestPatchWithDefaults() *OrganizationApiKeyRuleRequestPatch`

NewOrganizationApiKeyRuleRequestPatchWithDefaults instantiates a new OrganizationApiKeyRuleRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *OrganizationApiKeyRuleRequestPatch) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *OrganizationApiKeyRuleRequestPatch) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *OrganizationApiKeyRuleRequestPatch) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *OrganizationApiKeyRuleRequestPatch) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMaxAgeHours

`func (o *OrganizationApiKeyRuleRequestPatch) GetMaxAgeHours() int64`

GetMaxAgeHours returns the MaxAgeHours field if non-nil, zero value otherwise.

### GetMaxAgeHoursOk

`func (o *OrganizationApiKeyRuleRequestPatch) GetMaxAgeHoursOk() (*int64, bool)`

GetMaxAgeHoursOk returns a tuple with the MaxAgeHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAgeHours

`func (o *OrganizationApiKeyRuleRequestPatch) SetMaxAgeHours(v int64)`

SetMaxAgeHours sets MaxAgeHours field to given value.

### HasMaxAgeHours

`func (o *OrganizationApiKeyRuleRequestPatch) HasMaxAgeHours() bool`

HasMaxAgeHours returns a boolean if a field has been set.

### SetMaxAgeHoursNil

`func (o *OrganizationApiKeyRuleRequestPatch) SetMaxAgeHoursNil(b bool)`

 SetMaxAgeHoursNil sets the value for MaxAgeHours to be an explicit nil

### UnsetMaxAgeHours
`func (o *OrganizationApiKeyRuleRequestPatch) UnsetMaxAgeHours()`

UnsetMaxAgeHours ensures that no value is present for MaxAgeHours, not even an explicit nil
### GetRuleType

`func (o *OrganizationApiKeyRuleRequestPatch) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *OrganizationApiKeyRuleRequestPatch) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *OrganizationApiKeyRuleRequestPatch) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *OrganizationApiKeyRuleRequestPatch) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


