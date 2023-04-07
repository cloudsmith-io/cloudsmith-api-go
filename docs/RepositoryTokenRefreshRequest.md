# RepositoryTokenRefreshRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EulaRequired** | Pointer to **bool** | If checked, a EULA acceptance is required for this token. | [optional] 
**IsActive** | Pointer to **bool** | If enabled, the token will allow downloads based on configured restrictions (if any). | [optional] 
**LimitBandwidth** | Pointer to **NullableInt64** | The maximum download bandwidth allowed for the token. Values are expressed as the selected unit of bandwidth. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.  | [optional] 
**LimitBandwidthUnit** | Pointer to **NullableString** |  | [optional] [default to "Byte"]
**LimitDateRangeFrom** | Pointer to **NullableTime** | The starting date/time the token is allowed to be used from. | [optional] 
**LimitDateRangeTo** | Pointer to **NullableTime** | The ending date/time the token is allowed to be used until. | [optional] 
**LimitNumClients** | Pointer to **NullableInt64** | The maximum number of unique clients allowed for the token. Please note that since clients are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point. | [optional] 
**LimitNumDownloads** | Pointer to **NullableInt64** | The maximum number of downloads allowed for the token. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point. | [optional] 
**LimitPackageQuery** | Pointer to **NullableString** | The package-based search query to apply to restrict downloads to. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. This will still allow access to non-package files, such as metadata. | [optional] 
**LimitPathQuery** | Pointer to **NullableString** | THIS WILL SOON BE DEPRECATED, please use limit_package_query instead. The path-based search query to apply to restrict downloads to. This supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. The path evaluated does not include the domain name, the namespace, the entitlement code used, the package format, etc. and it always starts with a forward slash. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**ScheduledResetAt** | Pointer to **NullableTime** | The time at which the scheduled reset period has elapsed and the token limits were automatically reset to zero. | [optional] 
**ScheduledResetPeriod** | Pointer to **NullableString** |  | [optional] [default to "Never Reset"]
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewRepositoryTokenRefreshRequest

`func NewRepositoryTokenRefreshRequest() *RepositoryTokenRefreshRequest`

NewRepositoryTokenRefreshRequest instantiates a new RepositoryTokenRefreshRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryTokenRefreshRequestWithDefaults

`func NewRepositoryTokenRefreshRequestWithDefaults() *RepositoryTokenRefreshRequest`

NewRepositoryTokenRefreshRequestWithDefaults instantiates a new RepositoryTokenRefreshRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEulaRequired

`func (o *RepositoryTokenRefreshRequest) GetEulaRequired() bool`

GetEulaRequired returns the EulaRequired field if non-nil, zero value otherwise.

### GetEulaRequiredOk

`func (o *RepositoryTokenRefreshRequest) GetEulaRequiredOk() (*bool, bool)`

GetEulaRequiredOk returns a tuple with the EulaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaRequired

`func (o *RepositoryTokenRefreshRequest) SetEulaRequired(v bool)`

SetEulaRequired sets EulaRequired field to given value.

### HasEulaRequired

`func (o *RepositoryTokenRefreshRequest) HasEulaRequired() bool`

HasEulaRequired returns a boolean if a field has been set.

### GetIsActive

`func (o *RepositoryTokenRefreshRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RepositoryTokenRefreshRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RepositoryTokenRefreshRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RepositoryTokenRefreshRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLimitBandwidth

`func (o *RepositoryTokenRefreshRequest) GetLimitBandwidth() int64`

GetLimitBandwidth returns the LimitBandwidth field if non-nil, zero value otherwise.

### GetLimitBandwidthOk

`func (o *RepositoryTokenRefreshRequest) GetLimitBandwidthOk() (*int64, bool)`

GetLimitBandwidthOk returns a tuple with the LimitBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBandwidth

`func (o *RepositoryTokenRefreshRequest) SetLimitBandwidth(v int64)`

SetLimitBandwidth sets LimitBandwidth field to given value.

### HasLimitBandwidth

`func (o *RepositoryTokenRefreshRequest) HasLimitBandwidth() bool`

HasLimitBandwidth returns a boolean if a field has been set.

### SetLimitBandwidthNil

`func (o *RepositoryTokenRefreshRequest) SetLimitBandwidthNil(b bool)`

 SetLimitBandwidthNil sets the value for LimitBandwidth to be an explicit nil

### UnsetLimitBandwidth
`func (o *RepositoryTokenRefreshRequest) UnsetLimitBandwidth()`

UnsetLimitBandwidth ensures that no value is present for LimitBandwidth, not even an explicit nil
### GetLimitBandwidthUnit

`func (o *RepositoryTokenRefreshRequest) GetLimitBandwidthUnit() string`

GetLimitBandwidthUnit returns the LimitBandwidthUnit field if non-nil, zero value otherwise.

### GetLimitBandwidthUnitOk

`func (o *RepositoryTokenRefreshRequest) GetLimitBandwidthUnitOk() (*string, bool)`

GetLimitBandwidthUnitOk returns a tuple with the LimitBandwidthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBandwidthUnit

`func (o *RepositoryTokenRefreshRequest) SetLimitBandwidthUnit(v string)`

SetLimitBandwidthUnit sets LimitBandwidthUnit field to given value.

### HasLimitBandwidthUnit

`func (o *RepositoryTokenRefreshRequest) HasLimitBandwidthUnit() bool`

HasLimitBandwidthUnit returns a boolean if a field has been set.

### SetLimitBandwidthUnitNil

`func (o *RepositoryTokenRefreshRequest) SetLimitBandwidthUnitNil(b bool)`

 SetLimitBandwidthUnitNil sets the value for LimitBandwidthUnit to be an explicit nil

### UnsetLimitBandwidthUnit
`func (o *RepositoryTokenRefreshRequest) UnsetLimitBandwidthUnit()`

UnsetLimitBandwidthUnit ensures that no value is present for LimitBandwidthUnit, not even an explicit nil
### GetLimitDateRangeFrom

`func (o *RepositoryTokenRefreshRequest) GetLimitDateRangeFrom() time.Time`

GetLimitDateRangeFrom returns the LimitDateRangeFrom field if non-nil, zero value otherwise.

### GetLimitDateRangeFromOk

`func (o *RepositoryTokenRefreshRequest) GetLimitDateRangeFromOk() (*time.Time, bool)`

GetLimitDateRangeFromOk returns a tuple with the LimitDateRangeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitDateRangeFrom

`func (o *RepositoryTokenRefreshRequest) SetLimitDateRangeFrom(v time.Time)`

SetLimitDateRangeFrom sets LimitDateRangeFrom field to given value.

### HasLimitDateRangeFrom

`func (o *RepositoryTokenRefreshRequest) HasLimitDateRangeFrom() bool`

HasLimitDateRangeFrom returns a boolean if a field has been set.

### SetLimitDateRangeFromNil

`func (o *RepositoryTokenRefreshRequest) SetLimitDateRangeFromNil(b bool)`

 SetLimitDateRangeFromNil sets the value for LimitDateRangeFrom to be an explicit nil

### UnsetLimitDateRangeFrom
`func (o *RepositoryTokenRefreshRequest) UnsetLimitDateRangeFrom()`

UnsetLimitDateRangeFrom ensures that no value is present for LimitDateRangeFrom, not even an explicit nil
### GetLimitDateRangeTo

`func (o *RepositoryTokenRefreshRequest) GetLimitDateRangeTo() time.Time`

GetLimitDateRangeTo returns the LimitDateRangeTo field if non-nil, zero value otherwise.

### GetLimitDateRangeToOk

`func (o *RepositoryTokenRefreshRequest) GetLimitDateRangeToOk() (*time.Time, bool)`

GetLimitDateRangeToOk returns a tuple with the LimitDateRangeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitDateRangeTo

`func (o *RepositoryTokenRefreshRequest) SetLimitDateRangeTo(v time.Time)`

SetLimitDateRangeTo sets LimitDateRangeTo field to given value.

### HasLimitDateRangeTo

`func (o *RepositoryTokenRefreshRequest) HasLimitDateRangeTo() bool`

HasLimitDateRangeTo returns a boolean if a field has been set.

### SetLimitDateRangeToNil

`func (o *RepositoryTokenRefreshRequest) SetLimitDateRangeToNil(b bool)`

 SetLimitDateRangeToNil sets the value for LimitDateRangeTo to be an explicit nil

### UnsetLimitDateRangeTo
`func (o *RepositoryTokenRefreshRequest) UnsetLimitDateRangeTo()`

UnsetLimitDateRangeTo ensures that no value is present for LimitDateRangeTo, not even an explicit nil
### GetLimitNumClients

`func (o *RepositoryTokenRefreshRequest) GetLimitNumClients() int64`

GetLimitNumClients returns the LimitNumClients field if non-nil, zero value otherwise.

### GetLimitNumClientsOk

`func (o *RepositoryTokenRefreshRequest) GetLimitNumClientsOk() (*int64, bool)`

GetLimitNumClientsOk returns a tuple with the LimitNumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNumClients

`func (o *RepositoryTokenRefreshRequest) SetLimitNumClients(v int64)`

SetLimitNumClients sets LimitNumClients field to given value.

### HasLimitNumClients

`func (o *RepositoryTokenRefreshRequest) HasLimitNumClients() bool`

HasLimitNumClients returns a boolean if a field has been set.

### SetLimitNumClientsNil

`func (o *RepositoryTokenRefreshRequest) SetLimitNumClientsNil(b bool)`

 SetLimitNumClientsNil sets the value for LimitNumClients to be an explicit nil

### UnsetLimitNumClients
`func (o *RepositoryTokenRefreshRequest) UnsetLimitNumClients()`

UnsetLimitNumClients ensures that no value is present for LimitNumClients, not even an explicit nil
### GetLimitNumDownloads

`func (o *RepositoryTokenRefreshRequest) GetLimitNumDownloads() int64`

GetLimitNumDownloads returns the LimitNumDownloads field if non-nil, zero value otherwise.

### GetLimitNumDownloadsOk

`func (o *RepositoryTokenRefreshRequest) GetLimitNumDownloadsOk() (*int64, bool)`

GetLimitNumDownloadsOk returns a tuple with the LimitNumDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNumDownloads

`func (o *RepositoryTokenRefreshRequest) SetLimitNumDownloads(v int64)`

SetLimitNumDownloads sets LimitNumDownloads field to given value.

### HasLimitNumDownloads

`func (o *RepositoryTokenRefreshRequest) HasLimitNumDownloads() bool`

HasLimitNumDownloads returns a boolean if a field has been set.

### SetLimitNumDownloadsNil

`func (o *RepositoryTokenRefreshRequest) SetLimitNumDownloadsNil(b bool)`

 SetLimitNumDownloadsNil sets the value for LimitNumDownloads to be an explicit nil

### UnsetLimitNumDownloads
`func (o *RepositoryTokenRefreshRequest) UnsetLimitNumDownloads()`

UnsetLimitNumDownloads ensures that no value is present for LimitNumDownloads, not even an explicit nil
### GetLimitPackageQuery

`func (o *RepositoryTokenRefreshRequest) GetLimitPackageQuery() string`

GetLimitPackageQuery returns the LimitPackageQuery field if non-nil, zero value otherwise.

### GetLimitPackageQueryOk

`func (o *RepositoryTokenRefreshRequest) GetLimitPackageQueryOk() (*string, bool)`

GetLimitPackageQueryOk returns a tuple with the LimitPackageQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPackageQuery

`func (o *RepositoryTokenRefreshRequest) SetLimitPackageQuery(v string)`

SetLimitPackageQuery sets LimitPackageQuery field to given value.

### HasLimitPackageQuery

`func (o *RepositoryTokenRefreshRequest) HasLimitPackageQuery() bool`

HasLimitPackageQuery returns a boolean if a field has been set.

### SetLimitPackageQueryNil

`func (o *RepositoryTokenRefreshRequest) SetLimitPackageQueryNil(b bool)`

 SetLimitPackageQueryNil sets the value for LimitPackageQuery to be an explicit nil

### UnsetLimitPackageQuery
`func (o *RepositoryTokenRefreshRequest) UnsetLimitPackageQuery()`

UnsetLimitPackageQuery ensures that no value is present for LimitPackageQuery, not even an explicit nil
### GetLimitPathQuery

`func (o *RepositoryTokenRefreshRequest) GetLimitPathQuery() string`

GetLimitPathQuery returns the LimitPathQuery field if non-nil, zero value otherwise.

### GetLimitPathQueryOk

`func (o *RepositoryTokenRefreshRequest) GetLimitPathQueryOk() (*string, bool)`

GetLimitPathQueryOk returns a tuple with the LimitPathQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPathQuery

`func (o *RepositoryTokenRefreshRequest) SetLimitPathQuery(v string)`

SetLimitPathQuery sets LimitPathQuery field to given value.

### HasLimitPathQuery

`func (o *RepositoryTokenRefreshRequest) HasLimitPathQuery() bool`

HasLimitPathQuery returns a boolean if a field has been set.

### SetLimitPathQueryNil

`func (o *RepositoryTokenRefreshRequest) SetLimitPathQueryNil(b bool)`

 SetLimitPathQueryNil sets the value for LimitPathQuery to be an explicit nil

### UnsetLimitPathQuery
`func (o *RepositoryTokenRefreshRequest) UnsetLimitPathQuery()`

UnsetLimitPathQuery ensures that no value is present for LimitPathQuery, not even an explicit nil
### GetMetadata

`func (o *RepositoryTokenRefreshRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RepositoryTokenRefreshRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RepositoryTokenRefreshRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RepositoryTokenRefreshRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *RepositoryTokenRefreshRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *RepositoryTokenRefreshRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetScheduledResetAt

`func (o *RepositoryTokenRefreshRequest) GetScheduledResetAt() time.Time`

GetScheduledResetAt returns the ScheduledResetAt field if non-nil, zero value otherwise.

### GetScheduledResetAtOk

`func (o *RepositoryTokenRefreshRequest) GetScheduledResetAtOk() (*time.Time, bool)`

GetScheduledResetAtOk returns a tuple with the ScheduledResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledResetAt

`func (o *RepositoryTokenRefreshRequest) SetScheduledResetAt(v time.Time)`

SetScheduledResetAt sets ScheduledResetAt field to given value.

### HasScheduledResetAt

`func (o *RepositoryTokenRefreshRequest) HasScheduledResetAt() bool`

HasScheduledResetAt returns a boolean if a field has been set.

### SetScheduledResetAtNil

`func (o *RepositoryTokenRefreshRequest) SetScheduledResetAtNil(b bool)`

 SetScheduledResetAtNil sets the value for ScheduledResetAt to be an explicit nil

### UnsetScheduledResetAt
`func (o *RepositoryTokenRefreshRequest) UnsetScheduledResetAt()`

UnsetScheduledResetAt ensures that no value is present for ScheduledResetAt, not even an explicit nil
### GetScheduledResetPeriod

`func (o *RepositoryTokenRefreshRequest) GetScheduledResetPeriod() string`

GetScheduledResetPeriod returns the ScheduledResetPeriod field if non-nil, zero value otherwise.

### GetScheduledResetPeriodOk

`func (o *RepositoryTokenRefreshRequest) GetScheduledResetPeriodOk() (*string, bool)`

GetScheduledResetPeriodOk returns a tuple with the ScheduledResetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledResetPeriod

`func (o *RepositoryTokenRefreshRequest) SetScheduledResetPeriod(v string)`

SetScheduledResetPeriod sets ScheduledResetPeriod field to given value.

### HasScheduledResetPeriod

`func (o *RepositoryTokenRefreshRequest) HasScheduledResetPeriod() bool`

HasScheduledResetPeriod returns a boolean if a field has been set.

### SetScheduledResetPeriodNil

`func (o *RepositoryTokenRefreshRequest) SetScheduledResetPeriodNil(b bool)`

 SetScheduledResetPeriodNil sets the value for ScheduledResetPeriod to be an explicit nil

### UnsetScheduledResetPeriod
`func (o *RepositoryTokenRefreshRequest) UnsetScheduledResetPeriod()`

UnsetScheduledResetPeriod ensures that no value is present for ScheduledResetPeriod, not even an explicit nil
### GetToken

`func (o *RepositoryTokenRefreshRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RepositoryTokenRefreshRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RepositoryTokenRefreshRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RepositoryTokenRefreshRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


