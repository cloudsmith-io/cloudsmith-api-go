# EntitlementsPartialUpdate

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
**Name** | Pointer to **string** | None | [optional] 
**ScheduledResetAt** | Pointer to **string** | The time at which the scheduled reset period has elapsed and the token limits were automatically reset to zero. | [optional] 
**ScheduledResetPeriod** | Pointer to **string** | None | [optional] 
**Token** | Pointer to **string** | None | [optional] 

## Methods

### NewEntitlementsPartialUpdate

`func NewEntitlementsPartialUpdate() *EntitlementsPartialUpdate`

NewEntitlementsPartialUpdate instantiates a new EntitlementsPartialUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementsPartialUpdateWithDefaults

`func NewEntitlementsPartialUpdateWithDefaults() *EntitlementsPartialUpdate`

NewEntitlementsPartialUpdateWithDefaults instantiates a new EntitlementsPartialUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEulaRequired

`func (o *EntitlementsPartialUpdate) GetEulaRequired() bool`

GetEulaRequired returns the EulaRequired field if non-nil, zero value otherwise.

### GetEulaRequiredOk

`func (o *EntitlementsPartialUpdate) GetEulaRequiredOk() (*bool, bool)`

GetEulaRequiredOk returns a tuple with the EulaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaRequired

`func (o *EntitlementsPartialUpdate) SetEulaRequired(v bool)`

SetEulaRequired sets EulaRequired field to given value.

### HasEulaRequired

`func (o *EntitlementsPartialUpdate) HasEulaRequired() bool`

HasEulaRequired returns a boolean if a field has been set.

### GetIsActive

`func (o *EntitlementsPartialUpdate) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *EntitlementsPartialUpdate) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *EntitlementsPartialUpdate) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *EntitlementsPartialUpdate) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLimitBandwidth

`func (o *EntitlementsPartialUpdate) GetLimitBandwidth() int64`

GetLimitBandwidth returns the LimitBandwidth field if non-nil, zero value otherwise.

### GetLimitBandwidthOk

`func (o *EntitlementsPartialUpdate) GetLimitBandwidthOk() (*int64, bool)`

GetLimitBandwidthOk returns a tuple with the LimitBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBandwidth

`func (o *EntitlementsPartialUpdate) SetLimitBandwidth(v int64)`

SetLimitBandwidth sets LimitBandwidth field to given value.

### HasLimitBandwidth

`func (o *EntitlementsPartialUpdate) HasLimitBandwidth() bool`

HasLimitBandwidth returns a boolean if a field has been set.

### GetLimitBandwidthUnit

`func (o *EntitlementsPartialUpdate) GetLimitBandwidthUnit() string`

GetLimitBandwidthUnit returns the LimitBandwidthUnit field if non-nil, zero value otherwise.

### GetLimitBandwidthUnitOk

`func (o *EntitlementsPartialUpdate) GetLimitBandwidthUnitOk() (*string, bool)`

GetLimitBandwidthUnitOk returns a tuple with the LimitBandwidthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBandwidthUnit

`func (o *EntitlementsPartialUpdate) SetLimitBandwidthUnit(v string)`

SetLimitBandwidthUnit sets LimitBandwidthUnit field to given value.

### HasLimitBandwidthUnit

`func (o *EntitlementsPartialUpdate) HasLimitBandwidthUnit() bool`

HasLimitBandwidthUnit returns a boolean if a field has been set.

### GetLimitDateRangeFrom

`func (o *EntitlementsPartialUpdate) GetLimitDateRangeFrom() string`

GetLimitDateRangeFrom returns the LimitDateRangeFrom field if non-nil, zero value otherwise.

### GetLimitDateRangeFromOk

`func (o *EntitlementsPartialUpdate) GetLimitDateRangeFromOk() (*string, bool)`

GetLimitDateRangeFromOk returns a tuple with the LimitDateRangeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitDateRangeFrom

`func (o *EntitlementsPartialUpdate) SetLimitDateRangeFrom(v string)`

SetLimitDateRangeFrom sets LimitDateRangeFrom field to given value.

### HasLimitDateRangeFrom

`func (o *EntitlementsPartialUpdate) HasLimitDateRangeFrom() bool`

HasLimitDateRangeFrom returns a boolean if a field has been set.

### GetLimitDateRangeTo

`func (o *EntitlementsPartialUpdate) GetLimitDateRangeTo() string`

GetLimitDateRangeTo returns the LimitDateRangeTo field if non-nil, zero value otherwise.

### GetLimitDateRangeToOk

`func (o *EntitlementsPartialUpdate) GetLimitDateRangeToOk() (*string, bool)`

GetLimitDateRangeToOk returns a tuple with the LimitDateRangeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitDateRangeTo

`func (o *EntitlementsPartialUpdate) SetLimitDateRangeTo(v string)`

SetLimitDateRangeTo sets LimitDateRangeTo field to given value.

### HasLimitDateRangeTo

`func (o *EntitlementsPartialUpdate) HasLimitDateRangeTo() bool`

HasLimitDateRangeTo returns a boolean if a field has been set.

### GetLimitNumClients

`func (o *EntitlementsPartialUpdate) GetLimitNumClients() int64`

GetLimitNumClients returns the LimitNumClients field if non-nil, zero value otherwise.

### GetLimitNumClientsOk

`func (o *EntitlementsPartialUpdate) GetLimitNumClientsOk() (*int64, bool)`

GetLimitNumClientsOk returns a tuple with the LimitNumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNumClients

`func (o *EntitlementsPartialUpdate) SetLimitNumClients(v int64)`

SetLimitNumClients sets LimitNumClients field to given value.

### HasLimitNumClients

`func (o *EntitlementsPartialUpdate) HasLimitNumClients() bool`

HasLimitNumClients returns a boolean if a field has been set.

### GetLimitNumDownloads

`func (o *EntitlementsPartialUpdate) GetLimitNumDownloads() int64`

GetLimitNumDownloads returns the LimitNumDownloads field if non-nil, zero value otherwise.

### GetLimitNumDownloadsOk

`func (o *EntitlementsPartialUpdate) GetLimitNumDownloadsOk() (*int64, bool)`

GetLimitNumDownloadsOk returns a tuple with the LimitNumDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNumDownloads

`func (o *EntitlementsPartialUpdate) SetLimitNumDownloads(v int64)`

SetLimitNumDownloads sets LimitNumDownloads field to given value.

### HasLimitNumDownloads

`func (o *EntitlementsPartialUpdate) HasLimitNumDownloads() bool`

HasLimitNumDownloads returns a boolean if a field has been set.

### GetLimitPackageQuery

`func (o *EntitlementsPartialUpdate) GetLimitPackageQuery() string`

GetLimitPackageQuery returns the LimitPackageQuery field if non-nil, zero value otherwise.

### GetLimitPackageQueryOk

`func (o *EntitlementsPartialUpdate) GetLimitPackageQueryOk() (*string, bool)`

GetLimitPackageQueryOk returns a tuple with the LimitPackageQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPackageQuery

`func (o *EntitlementsPartialUpdate) SetLimitPackageQuery(v string)`

SetLimitPackageQuery sets LimitPackageQuery field to given value.

### HasLimitPackageQuery

`func (o *EntitlementsPartialUpdate) HasLimitPackageQuery() bool`

HasLimitPackageQuery returns a boolean if a field has been set.

### GetLimitPathQuery

`func (o *EntitlementsPartialUpdate) GetLimitPathQuery() string`

GetLimitPathQuery returns the LimitPathQuery field if non-nil, zero value otherwise.

### GetLimitPathQueryOk

`func (o *EntitlementsPartialUpdate) GetLimitPathQueryOk() (*string, bool)`

GetLimitPathQueryOk returns a tuple with the LimitPathQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPathQuery

`func (o *EntitlementsPartialUpdate) SetLimitPathQuery(v string)`

SetLimitPathQuery sets LimitPathQuery field to given value.

### HasLimitPathQuery

`func (o *EntitlementsPartialUpdate) HasLimitPathQuery() bool`

HasLimitPathQuery returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementsPartialUpdate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementsPartialUpdate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementsPartialUpdate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementsPartialUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *EntitlementsPartialUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementsPartialUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementsPartialUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntitlementsPartialUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduledResetAt

`func (o *EntitlementsPartialUpdate) GetScheduledResetAt() string`

GetScheduledResetAt returns the ScheduledResetAt field if non-nil, zero value otherwise.

### GetScheduledResetAtOk

`func (o *EntitlementsPartialUpdate) GetScheduledResetAtOk() (*string, bool)`

GetScheduledResetAtOk returns a tuple with the ScheduledResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledResetAt

`func (o *EntitlementsPartialUpdate) SetScheduledResetAt(v string)`

SetScheduledResetAt sets ScheduledResetAt field to given value.

### HasScheduledResetAt

`func (o *EntitlementsPartialUpdate) HasScheduledResetAt() bool`

HasScheduledResetAt returns a boolean if a field has been set.

### GetScheduledResetPeriod

`func (o *EntitlementsPartialUpdate) GetScheduledResetPeriod() string`

GetScheduledResetPeriod returns the ScheduledResetPeriod field if non-nil, zero value otherwise.

### GetScheduledResetPeriodOk

`func (o *EntitlementsPartialUpdate) GetScheduledResetPeriodOk() (*string, bool)`

GetScheduledResetPeriodOk returns a tuple with the ScheduledResetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledResetPeriod

`func (o *EntitlementsPartialUpdate) SetScheduledResetPeriod(v string)`

SetScheduledResetPeriod sets ScheduledResetPeriod field to given value.

### HasScheduledResetPeriod

`func (o *EntitlementsPartialUpdate) HasScheduledResetPeriod() bool`

HasScheduledResetPeriod returns a boolean if a field has been set.

### GetToken

`func (o *EntitlementsPartialUpdate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EntitlementsPartialUpdate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EntitlementsPartialUpdate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EntitlementsPartialUpdate) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


