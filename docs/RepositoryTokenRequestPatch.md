# RepositoryTokenRequestPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessPrivateBroadcasts** | Pointer to **bool** | If enabled, this token can be used for private broadcasts | [optional] 
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
**Name** | Pointer to **string** |  | [optional] 
**ScheduledResetAt** | Pointer to **NullableTime** | The time at which the scheduled reset period has elapsed and the token limits were automatically reset to zero. | [optional] 
**ScheduledResetPeriod** | Pointer to **NullableString** |  | [optional] [default to "Never Reset"]
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewRepositoryTokenRequestPatch

`func NewRepositoryTokenRequestPatch() *RepositoryTokenRequestPatch`

NewRepositoryTokenRequestPatch instantiates a new RepositoryTokenRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryTokenRequestPatchWithDefaults

`func NewRepositoryTokenRequestPatchWithDefaults() *RepositoryTokenRequestPatch`

NewRepositoryTokenRequestPatchWithDefaults instantiates a new RepositoryTokenRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessPrivateBroadcasts

`func (o *RepositoryTokenRequestPatch) GetAccessPrivateBroadcasts() bool`

GetAccessPrivateBroadcasts returns the AccessPrivateBroadcasts field if non-nil, zero value otherwise.

### GetAccessPrivateBroadcastsOk

`func (o *RepositoryTokenRequestPatch) GetAccessPrivateBroadcastsOk() (*bool, bool)`

GetAccessPrivateBroadcastsOk returns a tuple with the AccessPrivateBroadcasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPrivateBroadcasts

`func (o *RepositoryTokenRequestPatch) SetAccessPrivateBroadcasts(v bool)`

SetAccessPrivateBroadcasts sets AccessPrivateBroadcasts field to given value.

### HasAccessPrivateBroadcasts

`func (o *RepositoryTokenRequestPatch) HasAccessPrivateBroadcasts() bool`

HasAccessPrivateBroadcasts returns a boolean if a field has been set.

### GetEulaRequired

`func (o *RepositoryTokenRequestPatch) GetEulaRequired() bool`

GetEulaRequired returns the EulaRequired field if non-nil, zero value otherwise.

### GetEulaRequiredOk

`func (o *RepositoryTokenRequestPatch) GetEulaRequiredOk() (*bool, bool)`

GetEulaRequiredOk returns a tuple with the EulaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaRequired

`func (o *RepositoryTokenRequestPatch) SetEulaRequired(v bool)`

SetEulaRequired sets EulaRequired field to given value.

### HasEulaRequired

`func (o *RepositoryTokenRequestPatch) HasEulaRequired() bool`

HasEulaRequired returns a boolean if a field has been set.

### GetIsActive

`func (o *RepositoryTokenRequestPatch) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RepositoryTokenRequestPatch) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RepositoryTokenRequestPatch) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RepositoryTokenRequestPatch) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLimitBandwidth

`func (o *RepositoryTokenRequestPatch) GetLimitBandwidth() int64`

GetLimitBandwidth returns the LimitBandwidth field if non-nil, zero value otherwise.

### GetLimitBandwidthOk

`func (o *RepositoryTokenRequestPatch) GetLimitBandwidthOk() (*int64, bool)`

GetLimitBandwidthOk returns a tuple with the LimitBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBandwidth

`func (o *RepositoryTokenRequestPatch) SetLimitBandwidth(v int64)`

SetLimitBandwidth sets LimitBandwidth field to given value.

### HasLimitBandwidth

`func (o *RepositoryTokenRequestPatch) HasLimitBandwidth() bool`

HasLimitBandwidth returns a boolean if a field has been set.

### SetLimitBandwidthNil

`func (o *RepositoryTokenRequestPatch) SetLimitBandwidthNil(b bool)`

 SetLimitBandwidthNil sets the value for LimitBandwidth to be an explicit nil

### UnsetLimitBandwidth
`func (o *RepositoryTokenRequestPatch) UnsetLimitBandwidth()`

UnsetLimitBandwidth ensures that no value is present for LimitBandwidth, not even an explicit nil
### GetLimitBandwidthUnit

`func (o *RepositoryTokenRequestPatch) GetLimitBandwidthUnit() string`

GetLimitBandwidthUnit returns the LimitBandwidthUnit field if non-nil, zero value otherwise.

### GetLimitBandwidthUnitOk

`func (o *RepositoryTokenRequestPatch) GetLimitBandwidthUnitOk() (*string, bool)`

GetLimitBandwidthUnitOk returns a tuple with the LimitBandwidthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBandwidthUnit

`func (o *RepositoryTokenRequestPatch) SetLimitBandwidthUnit(v string)`

SetLimitBandwidthUnit sets LimitBandwidthUnit field to given value.

### HasLimitBandwidthUnit

`func (o *RepositoryTokenRequestPatch) HasLimitBandwidthUnit() bool`

HasLimitBandwidthUnit returns a boolean if a field has been set.

### SetLimitBandwidthUnitNil

`func (o *RepositoryTokenRequestPatch) SetLimitBandwidthUnitNil(b bool)`

 SetLimitBandwidthUnitNil sets the value for LimitBandwidthUnit to be an explicit nil

### UnsetLimitBandwidthUnit
`func (o *RepositoryTokenRequestPatch) UnsetLimitBandwidthUnit()`

UnsetLimitBandwidthUnit ensures that no value is present for LimitBandwidthUnit, not even an explicit nil
### GetLimitDateRangeFrom

`func (o *RepositoryTokenRequestPatch) GetLimitDateRangeFrom() time.Time`

GetLimitDateRangeFrom returns the LimitDateRangeFrom field if non-nil, zero value otherwise.

### GetLimitDateRangeFromOk

`func (o *RepositoryTokenRequestPatch) GetLimitDateRangeFromOk() (*time.Time, bool)`

GetLimitDateRangeFromOk returns a tuple with the LimitDateRangeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitDateRangeFrom

`func (o *RepositoryTokenRequestPatch) SetLimitDateRangeFrom(v time.Time)`

SetLimitDateRangeFrom sets LimitDateRangeFrom field to given value.

### HasLimitDateRangeFrom

`func (o *RepositoryTokenRequestPatch) HasLimitDateRangeFrom() bool`

HasLimitDateRangeFrom returns a boolean if a field has been set.

### SetLimitDateRangeFromNil

`func (o *RepositoryTokenRequestPatch) SetLimitDateRangeFromNil(b bool)`

 SetLimitDateRangeFromNil sets the value for LimitDateRangeFrom to be an explicit nil

### UnsetLimitDateRangeFrom
`func (o *RepositoryTokenRequestPatch) UnsetLimitDateRangeFrom()`

UnsetLimitDateRangeFrom ensures that no value is present for LimitDateRangeFrom, not even an explicit nil
### GetLimitDateRangeTo

`func (o *RepositoryTokenRequestPatch) GetLimitDateRangeTo() time.Time`

GetLimitDateRangeTo returns the LimitDateRangeTo field if non-nil, zero value otherwise.

### GetLimitDateRangeToOk

`func (o *RepositoryTokenRequestPatch) GetLimitDateRangeToOk() (*time.Time, bool)`

GetLimitDateRangeToOk returns a tuple with the LimitDateRangeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitDateRangeTo

`func (o *RepositoryTokenRequestPatch) SetLimitDateRangeTo(v time.Time)`

SetLimitDateRangeTo sets LimitDateRangeTo field to given value.

### HasLimitDateRangeTo

`func (o *RepositoryTokenRequestPatch) HasLimitDateRangeTo() bool`

HasLimitDateRangeTo returns a boolean if a field has been set.

### SetLimitDateRangeToNil

`func (o *RepositoryTokenRequestPatch) SetLimitDateRangeToNil(b bool)`

 SetLimitDateRangeToNil sets the value for LimitDateRangeTo to be an explicit nil

### UnsetLimitDateRangeTo
`func (o *RepositoryTokenRequestPatch) UnsetLimitDateRangeTo()`

UnsetLimitDateRangeTo ensures that no value is present for LimitDateRangeTo, not even an explicit nil
### GetLimitNumClients

`func (o *RepositoryTokenRequestPatch) GetLimitNumClients() int64`

GetLimitNumClients returns the LimitNumClients field if non-nil, zero value otherwise.

### GetLimitNumClientsOk

`func (o *RepositoryTokenRequestPatch) GetLimitNumClientsOk() (*int64, bool)`

GetLimitNumClientsOk returns a tuple with the LimitNumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNumClients

`func (o *RepositoryTokenRequestPatch) SetLimitNumClients(v int64)`

SetLimitNumClients sets LimitNumClients field to given value.

### HasLimitNumClients

`func (o *RepositoryTokenRequestPatch) HasLimitNumClients() bool`

HasLimitNumClients returns a boolean if a field has been set.

### SetLimitNumClientsNil

`func (o *RepositoryTokenRequestPatch) SetLimitNumClientsNil(b bool)`

 SetLimitNumClientsNil sets the value for LimitNumClients to be an explicit nil

### UnsetLimitNumClients
`func (o *RepositoryTokenRequestPatch) UnsetLimitNumClients()`

UnsetLimitNumClients ensures that no value is present for LimitNumClients, not even an explicit nil
### GetLimitNumDownloads

`func (o *RepositoryTokenRequestPatch) GetLimitNumDownloads() int64`

GetLimitNumDownloads returns the LimitNumDownloads field if non-nil, zero value otherwise.

### GetLimitNumDownloadsOk

`func (o *RepositoryTokenRequestPatch) GetLimitNumDownloadsOk() (*int64, bool)`

GetLimitNumDownloadsOk returns a tuple with the LimitNumDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNumDownloads

`func (o *RepositoryTokenRequestPatch) SetLimitNumDownloads(v int64)`

SetLimitNumDownloads sets LimitNumDownloads field to given value.

### HasLimitNumDownloads

`func (o *RepositoryTokenRequestPatch) HasLimitNumDownloads() bool`

HasLimitNumDownloads returns a boolean if a field has been set.

### SetLimitNumDownloadsNil

`func (o *RepositoryTokenRequestPatch) SetLimitNumDownloadsNil(b bool)`

 SetLimitNumDownloadsNil sets the value for LimitNumDownloads to be an explicit nil

### UnsetLimitNumDownloads
`func (o *RepositoryTokenRequestPatch) UnsetLimitNumDownloads()`

UnsetLimitNumDownloads ensures that no value is present for LimitNumDownloads, not even an explicit nil
### GetLimitPackageQuery

`func (o *RepositoryTokenRequestPatch) GetLimitPackageQuery() string`

GetLimitPackageQuery returns the LimitPackageQuery field if non-nil, zero value otherwise.

### GetLimitPackageQueryOk

`func (o *RepositoryTokenRequestPatch) GetLimitPackageQueryOk() (*string, bool)`

GetLimitPackageQueryOk returns a tuple with the LimitPackageQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPackageQuery

`func (o *RepositoryTokenRequestPatch) SetLimitPackageQuery(v string)`

SetLimitPackageQuery sets LimitPackageQuery field to given value.

### HasLimitPackageQuery

`func (o *RepositoryTokenRequestPatch) HasLimitPackageQuery() bool`

HasLimitPackageQuery returns a boolean if a field has been set.

### SetLimitPackageQueryNil

`func (o *RepositoryTokenRequestPatch) SetLimitPackageQueryNil(b bool)`

 SetLimitPackageQueryNil sets the value for LimitPackageQuery to be an explicit nil

### UnsetLimitPackageQuery
`func (o *RepositoryTokenRequestPatch) UnsetLimitPackageQuery()`

UnsetLimitPackageQuery ensures that no value is present for LimitPackageQuery, not even an explicit nil
### GetLimitPathQuery

`func (o *RepositoryTokenRequestPatch) GetLimitPathQuery() string`

GetLimitPathQuery returns the LimitPathQuery field if non-nil, zero value otherwise.

### GetLimitPathQueryOk

`func (o *RepositoryTokenRequestPatch) GetLimitPathQueryOk() (*string, bool)`

GetLimitPathQueryOk returns a tuple with the LimitPathQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPathQuery

`func (o *RepositoryTokenRequestPatch) SetLimitPathQuery(v string)`

SetLimitPathQuery sets LimitPathQuery field to given value.

### HasLimitPathQuery

`func (o *RepositoryTokenRequestPatch) HasLimitPathQuery() bool`

HasLimitPathQuery returns a boolean if a field has been set.

### SetLimitPathQueryNil

`func (o *RepositoryTokenRequestPatch) SetLimitPathQueryNil(b bool)`

 SetLimitPathQueryNil sets the value for LimitPathQuery to be an explicit nil

### UnsetLimitPathQuery
`func (o *RepositoryTokenRequestPatch) UnsetLimitPathQuery()`

UnsetLimitPathQuery ensures that no value is present for LimitPathQuery, not even an explicit nil
### GetMetadata

`func (o *RepositoryTokenRequestPatch) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RepositoryTokenRequestPatch) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RepositoryTokenRequestPatch) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RepositoryTokenRequestPatch) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *RepositoryTokenRequestPatch) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *RepositoryTokenRequestPatch) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *RepositoryTokenRequestPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryTokenRequestPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryTokenRequestPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RepositoryTokenRequestPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduledResetAt

`func (o *RepositoryTokenRequestPatch) GetScheduledResetAt() time.Time`

GetScheduledResetAt returns the ScheduledResetAt field if non-nil, zero value otherwise.

### GetScheduledResetAtOk

`func (o *RepositoryTokenRequestPatch) GetScheduledResetAtOk() (*time.Time, bool)`

GetScheduledResetAtOk returns a tuple with the ScheduledResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledResetAt

`func (o *RepositoryTokenRequestPatch) SetScheduledResetAt(v time.Time)`

SetScheduledResetAt sets ScheduledResetAt field to given value.

### HasScheduledResetAt

`func (o *RepositoryTokenRequestPatch) HasScheduledResetAt() bool`

HasScheduledResetAt returns a boolean if a field has been set.

### SetScheduledResetAtNil

`func (o *RepositoryTokenRequestPatch) SetScheduledResetAtNil(b bool)`

 SetScheduledResetAtNil sets the value for ScheduledResetAt to be an explicit nil

### UnsetScheduledResetAt
`func (o *RepositoryTokenRequestPatch) UnsetScheduledResetAt()`

UnsetScheduledResetAt ensures that no value is present for ScheduledResetAt, not even an explicit nil
### GetScheduledResetPeriod

`func (o *RepositoryTokenRequestPatch) GetScheduledResetPeriod() string`

GetScheduledResetPeriod returns the ScheduledResetPeriod field if non-nil, zero value otherwise.

### GetScheduledResetPeriodOk

`func (o *RepositoryTokenRequestPatch) GetScheduledResetPeriodOk() (*string, bool)`

GetScheduledResetPeriodOk returns a tuple with the ScheduledResetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledResetPeriod

`func (o *RepositoryTokenRequestPatch) SetScheduledResetPeriod(v string)`

SetScheduledResetPeriod sets ScheduledResetPeriod field to given value.

### HasScheduledResetPeriod

`func (o *RepositoryTokenRequestPatch) HasScheduledResetPeriod() bool`

HasScheduledResetPeriod returns a boolean if a field has been set.

### SetScheduledResetPeriodNil

`func (o *RepositoryTokenRequestPatch) SetScheduledResetPeriodNil(b bool)`

 SetScheduledResetPeriodNil sets the value for ScheduledResetPeriod to be an explicit nil

### UnsetScheduledResetPeriod
`func (o *RepositoryTokenRequestPatch) UnsetScheduledResetPeriod()`

UnsetScheduledResetPeriod ensures that no value is present for ScheduledResetPeriod, not even an explicit nil
### GetToken

`func (o *RepositoryTokenRequestPatch) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RepositoryTokenRequestPatch) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RepositoryTokenRequestPatch) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RepositoryTokenRequestPatch) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


