# OrganizationApiKeyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**EnforceRefresh** | Pointer to **bool** | When enabled, API&#39;s keys that violate the organization&#39;s policy will be replaced automatically. | [optional] 
**IsEnabled** | Pointer to **bool** | Whether this policy is currently active and enforced. | [optional] 
**LastAppliedAt** | Pointer to **time.Time** | The last time this policy was evaluated and applied by the expiry task. | [optional] [readonly] 
**MaxAgeHours** | Pointer to **NullableInt64** | The maximum permitted age of an API key for use in the organization. API keys older than this will no longer have access until they are refreshed. To disable the policy unset this value. | [optional] 
**RuleType** | **string** | Specify which account types this rule applies to: all accounts (deprecated — prefer separate service/user rules), service accounts only, or user accounts only. | 
**Slug** | Pointer to **string** | Human-readable identifier for the policy type. Automatically generated based on policy_type. | [optional] [readonly] 
**SlugPerm** | Pointer to **string** | An auto-generated id that uniquely identifies the API key policy. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewOrganizationApiKeyRule

`func NewOrganizationApiKeyRule(ruleType string, ) *OrganizationApiKeyRule`

NewOrganizationApiKeyRule instantiates a new OrganizationApiKeyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationApiKeyRuleWithDefaults

`func NewOrganizationApiKeyRuleWithDefaults() *OrganizationApiKeyRule`

NewOrganizationApiKeyRuleWithDefaults instantiates a new OrganizationApiKeyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *OrganizationApiKeyRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationApiKeyRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationApiKeyRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganizationApiKeyRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnforceRefresh

`func (o *OrganizationApiKeyRule) GetEnforceRefresh() bool`

GetEnforceRefresh returns the EnforceRefresh field if non-nil, zero value otherwise.

### GetEnforceRefreshOk

`func (o *OrganizationApiKeyRule) GetEnforceRefreshOk() (*bool, bool)`

GetEnforceRefreshOk returns a tuple with the EnforceRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceRefresh

`func (o *OrganizationApiKeyRule) SetEnforceRefresh(v bool)`

SetEnforceRefresh sets EnforceRefresh field to given value.

### HasEnforceRefresh

`func (o *OrganizationApiKeyRule) HasEnforceRefresh() bool`

HasEnforceRefresh returns a boolean if a field has been set.

### GetIsEnabled

`func (o *OrganizationApiKeyRule) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *OrganizationApiKeyRule) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *OrganizationApiKeyRule) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *OrganizationApiKeyRule) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetLastAppliedAt

`func (o *OrganizationApiKeyRule) GetLastAppliedAt() time.Time`

GetLastAppliedAt returns the LastAppliedAt field if non-nil, zero value otherwise.

### GetLastAppliedAtOk

`func (o *OrganizationApiKeyRule) GetLastAppliedAtOk() (*time.Time, bool)`

GetLastAppliedAtOk returns a tuple with the LastAppliedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAppliedAt

`func (o *OrganizationApiKeyRule) SetLastAppliedAt(v time.Time)`

SetLastAppliedAt sets LastAppliedAt field to given value.

### HasLastAppliedAt

`func (o *OrganizationApiKeyRule) HasLastAppliedAt() bool`

HasLastAppliedAt returns a boolean if a field has been set.

### GetMaxAgeHours

`func (o *OrganizationApiKeyRule) GetMaxAgeHours() int64`

GetMaxAgeHours returns the MaxAgeHours field if non-nil, zero value otherwise.

### GetMaxAgeHoursOk

`func (o *OrganizationApiKeyRule) GetMaxAgeHoursOk() (*int64, bool)`

GetMaxAgeHoursOk returns a tuple with the MaxAgeHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAgeHours

`func (o *OrganizationApiKeyRule) SetMaxAgeHours(v int64)`

SetMaxAgeHours sets MaxAgeHours field to given value.

### HasMaxAgeHours

`func (o *OrganizationApiKeyRule) HasMaxAgeHours() bool`

HasMaxAgeHours returns a boolean if a field has been set.

### SetMaxAgeHoursNil

`func (o *OrganizationApiKeyRule) SetMaxAgeHoursNil(b bool)`

 SetMaxAgeHoursNil sets the value for MaxAgeHours to be an explicit nil

### UnsetMaxAgeHours
`func (o *OrganizationApiKeyRule) UnsetMaxAgeHours()`

UnsetMaxAgeHours ensures that no value is present for MaxAgeHours, not even an explicit nil
### GetRuleType

`func (o *OrganizationApiKeyRule) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *OrganizationApiKeyRule) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *OrganizationApiKeyRule) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetSlug

`func (o *OrganizationApiKeyRule) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OrganizationApiKeyRule) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OrganizationApiKeyRule) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *OrganizationApiKeyRule) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSlugPerm

`func (o *OrganizationApiKeyRule) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *OrganizationApiKeyRule) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *OrganizationApiKeyRule) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *OrganizationApiKeyRule) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrganizationApiKeyRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationApiKeyRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationApiKeyRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationApiKeyRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


