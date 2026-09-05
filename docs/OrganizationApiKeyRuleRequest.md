# OrganizationApiKeyRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnforceRefresh** | Pointer to **bool** | When enabled, API&#39;s keys that violate the organization&#39;s policy will be replaced automatically. | [optional] 
**IsEnabled** | Pointer to **bool** | Whether this policy is currently active and enforced. | [optional] 
**MaxAgeHours** | Pointer to **NullableInt64** | The maximum permitted age of an API key for use in the organization. API keys older than this will no longer have access until they are refreshed. To disable the policy unset this value. | [optional] 
**RefreshImmediately** | Pointer to **bool** | If set to True, API keys that violate this rule will be replaced immediately after the request is made. There is no way to undo this. | [optional] 
**RuleType** | **string** | Specify which account types this rule applies to: all accounts (deprecated — prefer separate service/user rules), service accounts only, or user accounts only. | 

## Methods

### NewOrganizationApiKeyRuleRequest

`func NewOrganizationApiKeyRuleRequest(ruleType string, ) *OrganizationApiKeyRuleRequest`

NewOrganizationApiKeyRuleRequest instantiates a new OrganizationApiKeyRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationApiKeyRuleRequestWithDefaults

`func NewOrganizationApiKeyRuleRequestWithDefaults() *OrganizationApiKeyRuleRequest`

NewOrganizationApiKeyRuleRequestWithDefaults instantiates a new OrganizationApiKeyRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnforceRefresh

`func (o *OrganizationApiKeyRuleRequest) GetEnforceRefresh() bool`

GetEnforceRefresh returns the EnforceRefresh field if non-nil, zero value otherwise.

### GetEnforceRefreshOk

`func (o *OrganizationApiKeyRuleRequest) GetEnforceRefreshOk() (*bool, bool)`

GetEnforceRefreshOk returns a tuple with the EnforceRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceRefresh

`func (o *OrganizationApiKeyRuleRequest) SetEnforceRefresh(v bool)`

SetEnforceRefresh sets EnforceRefresh field to given value.

### HasEnforceRefresh

`func (o *OrganizationApiKeyRuleRequest) HasEnforceRefresh() bool`

HasEnforceRefresh returns a boolean if a field has been set.

### GetIsEnabled

`func (o *OrganizationApiKeyRuleRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *OrganizationApiKeyRuleRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *OrganizationApiKeyRuleRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *OrganizationApiKeyRuleRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMaxAgeHours

`func (o *OrganizationApiKeyRuleRequest) GetMaxAgeHours() int64`

GetMaxAgeHours returns the MaxAgeHours field if non-nil, zero value otherwise.

### GetMaxAgeHoursOk

`func (o *OrganizationApiKeyRuleRequest) GetMaxAgeHoursOk() (*int64, bool)`

GetMaxAgeHoursOk returns a tuple with the MaxAgeHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAgeHours

`func (o *OrganizationApiKeyRuleRequest) SetMaxAgeHours(v int64)`

SetMaxAgeHours sets MaxAgeHours field to given value.

### HasMaxAgeHours

`func (o *OrganizationApiKeyRuleRequest) HasMaxAgeHours() bool`

HasMaxAgeHours returns a boolean if a field has been set.

### SetMaxAgeHoursNil

`func (o *OrganizationApiKeyRuleRequest) SetMaxAgeHoursNil(b bool)`

 SetMaxAgeHoursNil sets the value for MaxAgeHours to be an explicit nil

### UnsetMaxAgeHours
`func (o *OrganizationApiKeyRuleRequest) UnsetMaxAgeHours()`

UnsetMaxAgeHours ensures that no value is present for MaxAgeHours, not even an explicit nil
### GetRefreshImmediately

`func (o *OrganizationApiKeyRuleRequest) GetRefreshImmediately() bool`

GetRefreshImmediately returns the RefreshImmediately field if non-nil, zero value otherwise.

### GetRefreshImmediatelyOk

`func (o *OrganizationApiKeyRuleRequest) GetRefreshImmediatelyOk() (*bool, bool)`

GetRefreshImmediatelyOk returns a tuple with the RefreshImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshImmediately

`func (o *OrganizationApiKeyRuleRequest) SetRefreshImmediately(v bool)`

SetRefreshImmediately sets RefreshImmediately field to given value.

### HasRefreshImmediately

`func (o *OrganizationApiKeyRuleRequest) HasRefreshImmediately() bool`

HasRefreshImmediately returns a boolean if a field has been set.

### GetRuleType

`func (o *OrganizationApiKeyRuleRequest) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *OrganizationApiKeyRuleRequest) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *OrganizationApiKeyRuleRequest) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


