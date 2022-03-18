# EntitlementsRefresh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EulaRequired** | Pointer to **bool** | If checked, a EULA acceptance is required for this token. | [optional] 
**IsActive** | Pointer to **bool** | If enabled, the token will allow downloads based on configured restrictions (if any). | [optional] 
**LimitBandwidth** | Pointer to **int64** | The maximum download bandwidth allowed for the token. Values are expressed as the selected unit of bandwidth. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.  | [optional] 
**LimitBandwidthUnit** | Pointer to **string** | None | [optional] 
**LimitDateRangeFrom** | Pointer to **string** | The starting date/time the token is allowed to be used from. | [optional] 
**LimitDateRangeTo** | Pointer to **string** | The ending date/time the token is allowed to be used until. | [optional] 
**LimitNumClients** | Pointer to **int64** | The maximum number of unique clients allowed for the token. Please note that since clients are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point. | [optional] 
**LimitNumDownloads** | Pointer to **int64** | The maximum number of downloads allowed for the token. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point. | [optional] 
**LimitPackageQuery** | Pointer to **string** | The package-based search query to apply to restrict downloads to. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. This will still allow access to non-package files, such as metadata. | [optional] 
**LimitPathQuery** | Pointer to **string** | The path-based search query to apply to restrict downloads to. This supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. The path evaluated does not include the domain name, the namespace, the entitlement code used, the package format, etc. and it always starts with a forward slash. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | None | [optional] 
**ScheduledResetAt** | Pointer to **string** | The time at which the scheduled reset period has elapsed and the token limits were automatically reset to zero. | [optional] 
**ScheduledResetPeriod** | Pointer to **string** | None | [optional] 
**Token** | Pointer to **string** | None | [optional] 

## Methods

### NewEntitlementsRefresh

`func NewEntitlementsRefresh() *EntitlementsRefresh`

NewEntitlementsRefresh instantiates a new EntitlementsRefresh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementsRefreshWithDefaults

`func NewEntitlementsRefreshWithDefaults() *EntitlementsRefresh`

NewEntitlementsRefreshWithDefaults instantiates a new EntitlementsRefresh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEulaRequired

`func (o *EntitlementsRefresh) GetEulaRequired() bool`

GetEulaRequired returns the EulaRequired field if non-nil, zero value otherwise.

### GetEulaRequiredOk

`func (o *EntitlementsRefresh) GetEulaRequiredOk() (*bool, bool)`

GetEulaRequiredOk returns a tuple with the EulaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaRequired

`func (o *EntitlementsRefresh) SetEulaRequired(v bool)`

SetEulaRequired sets EulaRequired field to given value.

### HasEulaRequired

`func (o *EntitlementsRefresh) HasEulaRequired() bool`

HasEulaRequired returns a boolean if a field has been set.

### GetIsActive

`func (o *EntitlementsRefresh) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *EntitlementsRefresh) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *EntitlementsRefresh) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *EntitlementsRefresh) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLimitBandwidth

`func (o *EntitlementsRefresh) GetLimitBandwidth() int64`

GetLimitBandwidth returns the LimitBandwidth field if non-nil, zero value otherwise.

### GetLimitBandwidthOk

`func (o *EntitlementsRefresh) GetLimitBandwidthOk() (*int64, bool)`

GetLimitBandwidthOk returns a tuple with the LimitBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBandwidth

`func (o *EntitlementsRefresh) SetLimitBandwidth(v int64)`

SetLimitBandwidth sets LimitBandwidth field to given value.

### HasLimitBandwidth

`func (o *EntitlementsRefresh) HasLimitBandwidth() bool`

HasLimitBandwidth returns a boolean if a field has been set.

### GetLimitBandwidthUnit

`func (o *EntitlementsRefresh) GetLimitBandwidthUnit() string`

GetLimitBandwidthUnit returns the LimitBandwidthUnit field if non-nil, zero value otherwise.

### GetLimitBandwidthUnitOk

`func (o *EntitlementsRefresh) GetLimitBandwidthUnitOk() (*string, bool)`

GetLimitBandwidthUnitOk returns a tuple with the LimitBandwidthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBandwidthUnit

`func (o *EntitlementsRefresh) SetLimitBandwidthUnit(v string)`

SetLimitBandwidthUnit sets LimitBandwidthUnit field to given value.

### HasLimitBandwidthUnit

`func (o *EntitlementsRefresh) HasLimitBandwidthUnit() bool`

HasLimitBandwidthUnit returns a boolean if a field has been set.

### GetLimitDateRangeFrom

`func (o *EntitlementsRefresh) GetLimitDateRangeFrom() string`

GetLimitDateRangeFrom returns the LimitDateRangeFrom field if non-nil, zero value otherwise.

### GetLimitDateRangeFromOk

`func (o *EntitlementsRefresh) GetLimitDateRangeFromOk() (*string, bool)`

GetLimitDateRangeFromOk returns a tuple with the LimitDateRangeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitDateRangeFrom

`func (o *EntitlementsRefresh) SetLimitDateRangeFrom(v string)`

SetLimitDateRangeFrom sets LimitDateRangeFrom field to given value.

### HasLimitDateRangeFrom

`func (o *EntitlementsRefresh) HasLimitDateRangeFrom() bool`

HasLimitDateRangeFrom returns a boolean if a field has been set.

### GetLimitDateRangeTo

`func (o *EntitlementsRefresh) GetLimitDateRangeTo() string`

GetLimitDateRangeTo returns the LimitDateRangeTo field if non-nil, zero value otherwise.

### GetLimitDateRangeToOk

`func (o *EntitlementsRefresh) GetLimitDateRangeToOk() (*string, bool)`

GetLimitDateRangeToOk returns a tuple with the LimitDateRangeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitDateRangeTo

`func (o *EntitlementsRefresh) SetLimitDateRangeTo(v string)`

SetLimitDateRangeTo sets LimitDateRangeTo field to given value.

### HasLimitDateRangeTo

`func (o *EntitlementsRefresh) HasLimitDateRangeTo() bool`

HasLimitDateRangeTo returns a boolean if a field has been set.

### GetLimitNumClients

`func (o *EntitlementsRefresh) GetLimitNumClients() int64`

GetLimitNumClients returns the LimitNumClients field if non-nil, zero value otherwise.

### GetLimitNumClientsOk

`func (o *EntitlementsRefresh) GetLimitNumClientsOk() (*int64, bool)`

GetLimitNumClientsOk returns a tuple with the LimitNumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNumClients

`func (o *EntitlementsRefresh) SetLimitNumClients(v int64)`

SetLimitNumClients sets LimitNumClients field to given value.

### HasLimitNumClients

`func (o *EntitlementsRefresh) HasLimitNumClients() bool`

HasLimitNumClients returns a boolean if a field has been set.

### GetLimitNumDownloads

`func (o *EntitlementsRefresh) GetLimitNumDownloads() int64`

GetLimitNumDownloads returns the LimitNumDownloads field if non-nil, zero value otherwise.

### GetLimitNumDownloadsOk

`func (o *EntitlementsRefresh) GetLimitNumDownloadsOk() (*int64, bool)`

GetLimitNumDownloadsOk returns a tuple with the LimitNumDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNumDownloads

`func (o *EntitlementsRefresh) SetLimitNumDownloads(v int64)`

SetLimitNumDownloads sets LimitNumDownloads field to given value.

### HasLimitNumDownloads

`func (o *EntitlementsRefresh) HasLimitNumDownloads() bool`

HasLimitNumDownloads returns a boolean if a field has been set.

### GetLimitPackageQuery

`func (o *EntitlementsRefresh) GetLimitPackageQuery() string`

GetLimitPackageQuery returns the LimitPackageQuery field if non-nil, zero value otherwise.

### GetLimitPackageQueryOk

`func (o *EntitlementsRefresh) GetLimitPackageQueryOk() (*string, bool)`

GetLimitPackageQueryOk returns a tuple with the LimitPackageQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPackageQuery

`func (o *EntitlementsRefresh) SetLimitPackageQuery(v string)`

SetLimitPackageQuery sets LimitPackageQuery field to given value.

### HasLimitPackageQuery

`func (o *EntitlementsRefresh) HasLimitPackageQuery() bool`

HasLimitPackageQuery returns a boolean if a field has been set.

### GetLimitPathQuery

`func (o *EntitlementsRefresh) GetLimitPathQuery() string`

GetLimitPathQuery returns the LimitPathQuery field if non-nil, zero value otherwise.

### GetLimitPathQueryOk

`func (o *EntitlementsRefresh) GetLimitPathQueryOk() (*string, bool)`

GetLimitPathQueryOk returns a tuple with the LimitPathQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPathQuery

`func (o *EntitlementsRefresh) SetLimitPathQuery(v string)`

SetLimitPathQuery sets LimitPathQuery field to given value.

### HasLimitPathQuery

`func (o *EntitlementsRefresh) HasLimitPathQuery() bool`

HasLimitPathQuery returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementsRefresh) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementsRefresh) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementsRefresh) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementsRefresh) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetScheduledResetAt

`func (o *EntitlementsRefresh) GetScheduledResetAt() string`

GetScheduledResetAt returns the ScheduledResetAt field if non-nil, zero value otherwise.

### GetScheduledResetAtOk

`func (o *EntitlementsRefresh) GetScheduledResetAtOk() (*string, bool)`

GetScheduledResetAtOk returns a tuple with the ScheduledResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledResetAt

`func (o *EntitlementsRefresh) SetScheduledResetAt(v string)`

SetScheduledResetAt sets ScheduledResetAt field to given value.

### HasScheduledResetAt

`func (o *EntitlementsRefresh) HasScheduledResetAt() bool`

HasScheduledResetAt returns a boolean if a field has been set.

### GetScheduledResetPeriod

`func (o *EntitlementsRefresh) GetScheduledResetPeriod() string`

GetScheduledResetPeriod returns the ScheduledResetPeriod field if non-nil, zero value otherwise.

### GetScheduledResetPeriodOk

`func (o *EntitlementsRefresh) GetScheduledResetPeriodOk() (*string, bool)`

GetScheduledResetPeriodOk returns a tuple with the ScheduledResetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledResetPeriod

`func (o *EntitlementsRefresh) SetScheduledResetPeriod(v string)`

SetScheduledResetPeriod sets ScheduledResetPeriod field to given value.

### HasScheduledResetPeriod

`func (o *EntitlementsRefresh) HasScheduledResetPeriod() bool`

HasScheduledResetPeriod returns a boolean if a field has been set.

### GetToken

`func (o *EntitlementsRefresh) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EntitlementsRefresh) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EntitlementsRefresh) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EntitlementsRefresh) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


