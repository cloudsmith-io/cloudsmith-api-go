# RepositoryToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessPrivateBroadcasts** | Pointer to **bool** | If enabled, this token can be used for private broadcasts | [optional] 
**Clients** | Pointer to **int64** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The datetime the token was updated at. | [optional] [readonly] 
**CreatedBy** | Pointer to **string** |  | [optional] [readonly] 
**CreatedByUrl** | Pointer to **string** |  | [optional] [readonly] 
**Default** | Pointer to **bool** | If selected this is the default token for this repository. | [optional] [readonly] 
**DisableUrl** | Pointer to **string** |  | [optional] [readonly] 
**Downloads** | Pointer to **int64** |  | [optional] [readonly] 
**EnableUrl** | Pointer to **string** |  | [optional] [readonly] 
**EulaAccepted** | Pointer to [**NullableEula**](Eula.md) |  | [optional] 
**EulaAcceptedAt** | Pointer to **NullableTime** | The datetime the EULA was accepted at. | [optional] [readonly] 
**EulaAcceptedFrom** | Pointer to **NullableString** |  | [optional] [readonly] 
**EulaRequired** | Pointer to **bool** | If checked, a EULA acceptance is required for this token. | [optional] 
**HasLimits** | Pointer to **bool** |  | [optional] [readonly] 
**Identifier** | Pointer to **NullableInt64** | Deprecated (23-05-15): Please use &#39;slug_perm&#39; instead. Previously: A monotonically increasing number that identified an entitlement within a repository. | [optional] [readonly] 
**IsActive** | Pointer to **bool** | If enabled, the token will allow downloads based on configured restrictions (if any). | [optional] 
**IsLimited** | Pointer to **bool** |  | [optional] [readonly] 
**LimitBandwidth** | Pointer to **NullableInt64** | The maximum download bandwidth allowed for the token. Values are expressed as the selected unit of bandwidth. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.  | [optional] 
**LimitBandwidthUnit** | Pointer to **NullableString** |  | [optional] [default to "Byte"]
**LimitDateRangeFrom** | Pointer to **NullableTime** | The starting date/time the token is allowed to be used from. | [optional] 
**LimitDateRangeTo** | Pointer to **NullableTime** | The ending date/time the token is allowed to be used until. | [optional] 
**LimitNumClients** | Pointer to **NullableInt64** | The maximum number of unique clients allowed for the token. Please note that since clients are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point. | [optional] 
**LimitNumDownloads** | Pointer to **NullableInt64** | The maximum number of downloads allowed for the token. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point. | [optional] 
**LimitPackageQuery** | Pointer to **NullableString** | The package-based search query to apply to restrict downloads to. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. This will still allow access to non-package files, such as metadata. | [optional] 
**LimitPathQuery** | Pointer to **NullableString** | THIS WILL SOON BE DEPRECATED, please use limit_package_query instead. The path-based search query to apply to restrict downloads to. This supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. The path evaluated does not include the domain name, the namespace, the entitlement code used, the package format, etc. and it always starts with a forward slash. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**RefreshUrl** | Pointer to **string** |  | [optional] [readonly] 
**ResetUrl** | Pointer to **string** |  | [optional] [readonly] 
**ScheduledResetAt** | Pointer to **NullableTime** | The time at which the scheduled reset period has elapsed and the token limits were automatically reset to zero. | [optional] 
**ScheduledResetPeriod** | Pointer to **NullableString** |  | [optional] [default to "Never Reset"]
**SelfUrl** | Pointer to **string** |  | [optional] [readonly] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**Token** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | The datetime the token was updated at. | [optional] [readonly] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] [readonly] 
**UpdatedByUrl** | Pointer to **NullableString** |  | [optional] [readonly] 
**Usage** | Pointer to **string** |  | [optional] [readonly] 
**User** | Pointer to **NullableString** |  | [optional] [readonly] 
**UserUrl** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewRepositoryToken

`func NewRepositoryToken(name string, ) *RepositoryToken`

NewRepositoryToken instantiates a new RepositoryToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryTokenWithDefaults

`func NewRepositoryTokenWithDefaults() *RepositoryToken`

NewRepositoryTokenWithDefaults instantiates a new RepositoryToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessPrivateBroadcasts

`func (o *RepositoryToken) GetAccessPrivateBroadcasts() bool`

GetAccessPrivateBroadcasts returns the AccessPrivateBroadcasts field if non-nil, zero value otherwise.

### GetAccessPrivateBroadcastsOk

`func (o *RepositoryToken) GetAccessPrivateBroadcastsOk() (*bool, bool)`

GetAccessPrivateBroadcastsOk returns a tuple with the AccessPrivateBroadcasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPrivateBroadcasts

`func (o *RepositoryToken) SetAccessPrivateBroadcasts(v bool)`

SetAccessPrivateBroadcasts sets AccessPrivateBroadcasts field to given value.

### HasAccessPrivateBroadcasts

`func (o *RepositoryToken) HasAccessPrivateBroadcasts() bool`

HasAccessPrivateBroadcasts returns a boolean if a field has been set.

### GetClients

`func (o *RepositoryToken) GetClients() int64`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *RepositoryToken) GetClientsOk() (*int64, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *RepositoryToken) SetClients(v int64)`

SetClients sets Clients field to given value.

### HasClients

`func (o *RepositoryToken) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RepositoryToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepositoryToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RepositoryToken) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RepositoryToken) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RepositoryToken) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RepositoryToken) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUrl

`func (o *RepositoryToken) GetCreatedByUrl() string`

GetCreatedByUrl returns the CreatedByUrl field if non-nil, zero value otherwise.

### GetCreatedByUrlOk

`func (o *RepositoryToken) GetCreatedByUrlOk() (*string, bool)`

GetCreatedByUrlOk returns a tuple with the CreatedByUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUrl

`func (o *RepositoryToken) SetCreatedByUrl(v string)`

SetCreatedByUrl sets CreatedByUrl field to given value.

### HasCreatedByUrl

`func (o *RepositoryToken) HasCreatedByUrl() bool`

HasCreatedByUrl returns a boolean if a field has been set.

### GetDefault

`func (o *RepositoryToken) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RepositoryToken) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RepositoryToken) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RepositoryToken) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDisableUrl

`func (o *RepositoryToken) GetDisableUrl() string`

GetDisableUrl returns the DisableUrl field if non-nil, zero value otherwise.

### GetDisableUrlOk

`func (o *RepositoryToken) GetDisableUrlOk() (*string, bool)`

GetDisableUrlOk returns a tuple with the DisableUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUrl

`func (o *RepositoryToken) SetDisableUrl(v string)`

SetDisableUrl sets DisableUrl field to given value.

### HasDisableUrl

`func (o *RepositoryToken) HasDisableUrl() bool`

HasDisableUrl returns a boolean if a field has been set.

### GetDownloads

`func (o *RepositoryToken) GetDownloads() int64`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *RepositoryToken) GetDownloadsOk() (*int64, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *RepositoryToken) SetDownloads(v int64)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *RepositoryToken) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetEnableUrl

`func (o *RepositoryToken) GetEnableUrl() string`

GetEnableUrl returns the EnableUrl field if non-nil, zero value otherwise.

### GetEnableUrlOk

`func (o *RepositoryToken) GetEnableUrlOk() (*string, bool)`

GetEnableUrlOk returns a tuple with the EnableUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUrl

`func (o *RepositoryToken) SetEnableUrl(v string)`

SetEnableUrl sets EnableUrl field to given value.

### HasEnableUrl

`func (o *RepositoryToken) HasEnableUrl() bool`

HasEnableUrl returns a boolean if a field has been set.

### GetEulaAccepted

`func (o *RepositoryToken) GetEulaAccepted() Eula`

GetEulaAccepted returns the EulaAccepted field if non-nil, zero value otherwise.

### GetEulaAcceptedOk

`func (o *RepositoryToken) GetEulaAcceptedOk() (*Eula, bool)`

GetEulaAcceptedOk returns a tuple with the EulaAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaAccepted

`func (o *RepositoryToken) SetEulaAccepted(v Eula)`

SetEulaAccepted sets EulaAccepted field to given value.

### HasEulaAccepted

`func (o *RepositoryToken) HasEulaAccepted() bool`

HasEulaAccepted returns a boolean if a field has been set.

### SetEulaAcceptedNil

`func (o *RepositoryToken) SetEulaAcceptedNil(b bool)`

 SetEulaAcceptedNil sets the value for EulaAccepted to be an explicit nil

### UnsetEulaAccepted
`func (o *RepositoryToken) UnsetEulaAccepted()`

UnsetEulaAccepted ensures that no value is present for EulaAccepted, not even an explicit nil
### GetEulaAcceptedAt

`func (o *RepositoryToken) GetEulaAcceptedAt() time.Time`

GetEulaAcceptedAt returns the EulaAcceptedAt field if non-nil, zero value otherwise.

### GetEulaAcceptedAtOk

`func (o *RepositoryToken) GetEulaAcceptedAtOk() (*time.Time, bool)`

GetEulaAcceptedAtOk returns a tuple with the EulaAcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaAcceptedAt

`func (o *RepositoryToken) SetEulaAcceptedAt(v time.Time)`

SetEulaAcceptedAt sets EulaAcceptedAt field to given value.

### HasEulaAcceptedAt

`func (o *RepositoryToken) HasEulaAcceptedAt() bool`

HasEulaAcceptedAt returns a boolean if a field has been set.

### SetEulaAcceptedAtNil

`func (o *RepositoryToken) SetEulaAcceptedAtNil(b bool)`

 SetEulaAcceptedAtNil sets the value for EulaAcceptedAt to be an explicit nil

### UnsetEulaAcceptedAt
`func (o *RepositoryToken) UnsetEulaAcceptedAt()`

UnsetEulaAcceptedAt ensures that no value is present for EulaAcceptedAt, not even an explicit nil
### GetEulaAcceptedFrom

`func (o *RepositoryToken) GetEulaAcceptedFrom() string`

GetEulaAcceptedFrom returns the EulaAcceptedFrom field if non-nil, zero value otherwise.

### GetEulaAcceptedFromOk

`func (o *RepositoryToken) GetEulaAcceptedFromOk() (*string, bool)`

GetEulaAcceptedFromOk returns a tuple with the EulaAcceptedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaAcceptedFrom

`func (o *RepositoryToken) SetEulaAcceptedFrom(v string)`

SetEulaAcceptedFrom sets EulaAcceptedFrom field to given value.

### HasEulaAcceptedFrom

`func (o *RepositoryToken) HasEulaAcceptedFrom() bool`

HasEulaAcceptedFrom returns a boolean if a field has been set.

### SetEulaAcceptedFromNil

`func (o *RepositoryToken) SetEulaAcceptedFromNil(b bool)`

 SetEulaAcceptedFromNil sets the value for EulaAcceptedFrom to be an explicit nil

### UnsetEulaAcceptedFrom
`func (o *RepositoryToken) UnsetEulaAcceptedFrom()`

UnsetEulaAcceptedFrom ensures that no value is present for EulaAcceptedFrom, not even an explicit nil
### GetEulaRequired

`func (o *RepositoryToken) GetEulaRequired() bool`

GetEulaRequired returns the EulaRequired field if non-nil, zero value otherwise.

### GetEulaRequiredOk

`func (o *RepositoryToken) GetEulaRequiredOk() (*bool, bool)`

GetEulaRequiredOk returns a tuple with the EulaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaRequired

`func (o *RepositoryToken) SetEulaRequired(v bool)`

SetEulaRequired sets EulaRequired field to given value.

### HasEulaRequired

`func (o *RepositoryToken) HasEulaRequired() bool`

HasEulaRequired returns a boolean if a field has been set.

### GetHasLimits

`func (o *RepositoryToken) GetHasLimits() bool`

GetHasLimits returns the HasLimits field if non-nil, zero value otherwise.

### GetHasLimitsOk

`func (o *RepositoryToken) GetHasLimitsOk() (*bool, bool)`

GetHasLimitsOk returns a tuple with the HasLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLimits

`func (o *RepositoryToken) SetHasLimits(v bool)`

SetHasLimits sets HasLimits field to given value.

### HasHasLimits

`func (o *RepositoryToken) HasHasLimits() bool`

HasHasLimits returns a boolean if a field has been set.

### GetIdentifier

`func (o *RepositoryToken) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *RepositoryToken) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *RepositoryToken) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *RepositoryToken) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *RepositoryToken) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *RepositoryToken) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetIsActive

`func (o *RepositoryToken) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RepositoryToken) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RepositoryToken) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RepositoryToken) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsLimited

`func (o *RepositoryToken) GetIsLimited() bool`

GetIsLimited returns the IsLimited field if non-nil, zero value otherwise.

### GetIsLimitedOk

`func (o *RepositoryToken) GetIsLimitedOk() (*bool, bool)`

GetIsLimitedOk returns a tuple with the IsLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLimited

`func (o *RepositoryToken) SetIsLimited(v bool)`

SetIsLimited sets IsLimited field to given value.

### HasIsLimited

`func (o *RepositoryToken) HasIsLimited() bool`

HasIsLimited returns a boolean if a field has been set.

### GetLimitBandwidth

`func (o *RepositoryToken) GetLimitBandwidth() int64`

GetLimitBandwidth returns the LimitBandwidth field if non-nil, zero value otherwise.

### GetLimitBandwidthOk

`func (o *RepositoryToken) GetLimitBandwidthOk() (*int64, bool)`

GetLimitBandwidthOk returns a tuple with the LimitBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBandwidth

`func (o *RepositoryToken) SetLimitBandwidth(v int64)`

SetLimitBandwidth sets LimitBandwidth field to given value.

### HasLimitBandwidth

`func (o *RepositoryToken) HasLimitBandwidth() bool`

HasLimitBandwidth returns a boolean if a field has been set.

### SetLimitBandwidthNil

`func (o *RepositoryToken) SetLimitBandwidthNil(b bool)`

 SetLimitBandwidthNil sets the value for LimitBandwidth to be an explicit nil

### UnsetLimitBandwidth
`func (o *RepositoryToken) UnsetLimitBandwidth()`

UnsetLimitBandwidth ensures that no value is present for LimitBandwidth, not even an explicit nil
### GetLimitBandwidthUnit

`func (o *RepositoryToken) GetLimitBandwidthUnit() string`

GetLimitBandwidthUnit returns the LimitBandwidthUnit field if non-nil, zero value otherwise.

### GetLimitBandwidthUnitOk

`func (o *RepositoryToken) GetLimitBandwidthUnitOk() (*string, bool)`

GetLimitBandwidthUnitOk returns a tuple with the LimitBandwidthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBandwidthUnit

`func (o *RepositoryToken) SetLimitBandwidthUnit(v string)`

SetLimitBandwidthUnit sets LimitBandwidthUnit field to given value.

### HasLimitBandwidthUnit

`func (o *RepositoryToken) HasLimitBandwidthUnit() bool`

HasLimitBandwidthUnit returns a boolean if a field has been set.

### SetLimitBandwidthUnitNil

`func (o *RepositoryToken) SetLimitBandwidthUnitNil(b bool)`

 SetLimitBandwidthUnitNil sets the value for LimitBandwidthUnit to be an explicit nil

### UnsetLimitBandwidthUnit
`func (o *RepositoryToken) UnsetLimitBandwidthUnit()`

UnsetLimitBandwidthUnit ensures that no value is present for LimitBandwidthUnit, not even an explicit nil
### GetLimitDateRangeFrom

`func (o *RepositoryToken) GetLimitDateRangeFrom() time.Time`

GetLimitDateRangeFrom returns the LimitDateRangeFrom field if non-nil, zero value otherwise.

### GetLimitDateRangeFromOk

`func (o *RepositoryToken) GetLimitDateRangeFromOk() (*time.Time, bool)`

GetLimitDateRangeFromOk returns a tuple with the LimitDateRangeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitDateRangeFrom

`func (o *RepositoryToken) SetLimitDateRangeFrom(v time.Time)`

SetLimitDateRangeFrom sets LimitDateRangeFrom field to given value.

### HasLimitDateRangeFrom

`func (o *RepositoryToken) HasLimitDateRangeFrom() bool`

HasLimitDateRangeFrom returns a boolean if a field has been set.

### SetLimitDateRangeFromNil

`func (o *RepositoryToken) SetLimitDateRangeFromNil(b bool)`

 SetLimitDateRangeFromNil sets the value for LimitDateRangeFrom to be an explicit nil

### UnsetLimitDateRangeFrom
`func (o *RepositoryToken) UnsetLimitDateRangeFrom()`

UnsetLimitDateRangeFrom ensures that no value is present for LimitDateRangeFrom, not even an explicit nil
### GetLimitDateRangeTo

`func (o *RepositoryToken) GetLimitDateRangeTo() time.Time`

GetLimitDateRangeTo returns the LimitDateRangeTo field if non-nil, zero value otherwise.

### GetLimitDateRangeToOk

`func (o *RepositoryToken) GetLimitDateRangeToOk() (*time.Time, bool)`

GetLimitDateRangeToOk returns a tuple with the LimitDateRangeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitDateRangeTo

`func (o *RepositoryToken) SetLimitDateRangeTo(v time.Time)`

SetLimitDateRangeTo sets LimitDateRangeTo field to given value.

### HasLimitDateRangeTo

`func (o *RepositoryToken) HasLimitDateRangeTo() bool`

HasLimitDateRangeTo returns a boolean if a field has been set.

### SetLimitDateRangeToNil

`func (o *RepositoryToken) SetLimitDateRangeToNil(b bool)`

 SetLimitDateRangeToNil sets the value for LimitDateRangeTo to be an explicit nil

### UnsetLimitDateRangeTo
`func (o *RepositoryToken) UnsetLimitDateRangeTo()`

UnsetLimitDateRangeTo ensures that no value is present for LimitDateRangeTo, not even an explicit nil
### GetLimitNumClients

`func (o *RepositoryToken) GetLimitNumClients() int64`

GetLimitNumClients returns the LimitNumClients field if non-nil, zero value otherwise.

### GetLimitNumClientsOk

`func (o *RepositoryToken) GetLimitNumClientsOk() (*int64, bool)`

GetLimitNumClientsOk returns a tuple with the LimitNumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNumClients

`func (o *RepositoryToken) SetLimitNumClients(v int64)`

SetLimitNumClients sets LimitNumClients field to given value.

### HasLimitNumClients

`func (o *RepositoryToken) HasLimitNumClients() bool`

HasLimitNumClients returns a boolean if a field has been set.

### SetLimitNumClientsNil

`func (o *RepositoryToken) SetLimitNumClientsNil(b bool)`

 SetLimitNumClientsNil sets the value for LimitNumClients to be an explicit nil

### UnsetLimitNumClients
`func (o *RepositoryToken) UnsetLimitNumClients()`

UnsetLimitNumClients ensures that no value is present for LimitNumClients, not even an explicit nil
### GetLimitNumDownloads

`func (o *RepositoryToken) GetLimitNumDownloads() int64`

GetLimitNumDownloads returns the LimitNumDownloads field if non-nil, zero value otherwise.

### GetLimitNumDownloadsOk

`func (o *RepositoryToken) GetLimitNumDownloadsOk() (*int64, bool)`

GetLimitNumDownloadsOk returns a tuple with the LimitNumDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNumDownloads

`func (o *RepositoryToken) SetLimitNumDownloads(v int64)`

SetLimitNumDownloads sets LimitNumDownloads field to given value.

### HasLimitNumDownloads

`func (o *RepositoryToken) HasLimitNumDownloads() bool`

HasLimitNumDownloads returns a boolean if a field has been set.

### SetLimitNumDownloadsNil

`func (o *RepositoryToken) SetLimitNumDownloadsNil(b bool)`

 SetLimitNumDownloadsNil sets the value for LimitNumDownloads to be an explicit nil

### UnsetLimitNumDownloads
`func (o *RepositoryToken) UnsetLimitNumDownloads()`

UnsetLimitNumDownloads ensures that no value is present for LimitNumDownloads, not even an explicit nil
### GetLimitPackageQuery

`func (o *RepositoryToken) GetLimitPackageQuery() string`

GetLimitPackageQuery returns the LimitPackageQuery field if non-nil, zero value otherwise.

### GetLimitPackageQueryOk

`func (o *RepositoryToken) GetLimitPackageQueryOk() (*string, bool)`

GetLimitPackageQueryOk returns a tuple with the LimitPackageQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPackageQuery

`func (o *RepositoryToken) SetLimitPackageQuery(v string)`

SetLimitPackageQuery sets LimitPackageQuery field to given value.

### HasLimitPackageQuery

`func (o *RepositoryToken) HasLimitPackageQuery() bool`

HasLimitPackageQuery returns a boolean if a field has been set.

### SetLimitPackageQueryNil

`func (o *RepositoryToken) SetLimitPackageQueryNil(b bool)`

 SetLimitPackageQueryNil sets the value for LimitPackageQuery to be an explicit nil

### UnsetLimitPackageQuery
`func (o *RepositoryToken) UnsetLimitPackageQuery()`

UnsetLimitPackageQuery ensures that no value is present for LimitPackageQuery, not even an explicit nil
### GetLimitPathQuery

`func (o *RepositoryToken) GetLimitPathQuery() string`

GetLimitPathQuery returns the LimitPathQuery field if non-nil, zero value otherwise.

### GetLimitPathQueryOk

`func (o *RepositoryToken) GetLimitPathQueryOk() (*string, bool)`

GetLimitPathQueryOk returns a tuple with the LimitPathQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPathQuery

`func (o *RepositoryToken) SetLimitPathQuery(v string)`

SetLimitPathQuery sets LimitPathQuery field to given value.

### HasLimitPathQuery

`func (o *RepositoryToken) HasLimitPathQuery() bool`

HasLimitPathQuery returns a boolean if a field has been set.

### SetLimitPathQueryNil

`func (o *RepositoryToken) SetLimitPathQueryNil(b bool)`

 SetLimitPathQueryNil sets the value for LimitPathQuery to be an explicit nil

### UnsetLimitPathQuery
`func (o *RepositoryToken) UnsetLimitPathQuery()`

UnsetLimitPathQuery ensures that no value is present for LimitPathQuery, not even an explicit nil
### GetMetadata

`func (o *RepositoryToken) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RepositoryToken) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RepositoryToken) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RepositoryToken) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *RepositoryToken) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *RepositoryToken) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *RepositoryToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryToken) SetName(v string)`

SetName sets Name field to given value.


### GetRefreshUrl

`func (o *RepositoryToken) GetRefreshUrl() string`

GetRefreshUrl returns the RefreshUrl field if non-nil, zero value otherwise.

### GetRefreshUrlOk

`func (o *RepositoryToken) GetRefreshUrlOk() (*string, bool)`

GetRefreshUrlOk returns a tuple with the RefreshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshUrl

`func (o *RepositoryToken) SetRefreshUrl(v string)`

SetRefreshUrl sets RefreshUrl field to given value.

### HasRefreshUrl

`func (o *RepositoryToken) HasRefreshUrl() bool`

HasRefreshUrl returns a boolean if a field has been set.

### GetResetUrl

`func (o *RepositoryToken) GetResetUrl() string`

GetResetUrl returns the ResetUrl field if non-nil, zero value otherwise.

### GetResetUrlOk

`func (o *RepositoryToken) GetResetUrlOk() (*string, bool)`

GetResetUrlOk returns a tuple with the ResetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetUrl

`func (o *RepositoryToken) SetResetUrl(v string)`

SetResetUrl sets ResetUrl field to given value.

### HasResetUrl

`func (o *RepositoryToken) HasResetUrl() bool`

HasResetUrl returns a boolean if a field has been set.

### GetScheduledResetAt

`func (o *RepositoryToken) GetScheduledResetAt() time.Time`

GetScheduledResetAt returns the ScheduledResetAt field if non-nil, zero value otherwise.

### GetScheduledResetAtOk

`func (o *RepositoryToken) GetScheduledResetAtOk() (*time.Time, bool)`

GetScheduledResetAtOk returns a tuple with the ScheduledResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledResetAt

`func (o *RepositoryToken) SetScheduledResetAt(v time.Time)`

SetScheduledResetAt sets ScheduledResetAt field to given value.

### HasScheduledResetAt

`func (o *RepositoryToken) HasScheduledResetAt() bool`

HasScheduledResetAt returns a boolean if a field has been set.

### SetScheduledResetAtNil

`func (o *RepositoryToken) SetScheduledResetAtNil(b bool)`

 SetScheduledResetAtNil sets the value for ScheduledResetAt to be an explicit nil

### UnsetScheduledResetAt
`func (o *RepositoryToken) UnsetScheduledResetAt()`

UnsetScheduledResetAt ensures that no value is present for ScheduledResetAt, not even an explicit nil
### GetScheduledResetPeriod

`func (o *RepositoryToken) GetScheduledResetPeriod() string`

GetScheduledResetPeriod returns the ScheduledResetPeriod field if non-nil, zero value otherwise.

### GetScheduledResetPeriodOk

`func (o *RepositoryToken) GetScheduledResetPeriodOk() (*string, bool)`

GetScheduledResetPeriodOk returns a tuple with the ScheduledResetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledResetPeriod

`func (o *RepositoryToken) SetScheduledResetPeriod(v string)`

SetScheduledResetPeriod sets ScheduledResetPeriod field to given value.

### HasScheduledResetPeriod

`func (o *RepositoryToken) HasScheduledResetPeriod() bool`

HasScheduledResetPeriod returns a boolean if a field has been set.

### SetScheduledResetPeriodNil

`func (o *RepositoryToken) SetScheduledResetPeriodNil(b bool)`

 SetScheduledResetPeriodNil sets the value for ScheduledResetPeriod to be an explicit nil

### UnsetScheduledResetPeriod
`func (o *RepositoryToken) UnsetScheduledResetPeriod()`

UnsetScheduledResetPeriod ensures that no value is present for ScheduledResetPeriod, not even an explicit nil
### GetSelfUrl

`func (o *RepositoryToken) GetSelfUrl() string`

GetSelfUrl returns the SelfUrl field if non-nil, zero value otherwise.

### GetSelfUrlOk

`func (o *RepositoryToken) GetSelfUrlOk() (*string, bool)`

GetSelfUrlOk returns a tuple with the SelfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfUrl

`func (o *RepositoryToken) SetSelfUrl(v string)`

SetSelfUrl sets SelfUrl field to given value.

### HasSelfUrl

`func (o *RepositoryToken) HasSelfUrl() bool`

HasSelfUrl returns a boolean if a field has been set.

### GetSlugPerm

`func (o *RepositoryToken) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *RepositoryToken) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *RepositoryToken) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *RepositoryToken) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetToken

`func (o *RepositoryToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RepositoryToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RepositoryToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RepositoryToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RepositoryToken) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RepositoryToken) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RepositoryToken) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RepositoryToken) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *RepositoryToken) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *RepositoryToken) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUpdatedBy

`func (o *RepositoryToken) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RepositoryToken) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RepositoryToken) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RepositoryToken) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *RepositoryToken) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *RepositoryToken) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetUpdatedByUrl

`func (o *RepositoryToken) GetUpdatedByUrl() string`

GetUpdatedByUrl returns the UpdatedByUrl field if non-nil, zero value otherwise.

### GetUpdatedByUrlOk

`func (o *RepositoryToken) GetUpdatedByUrlOk() (*string, bool)`

GetUpdatedByUrlOk returns a tuple with the UpdatedByUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUrl

`func (o *RepositoryToken) SetUpdatedByUrl(v string)`

SetUpdatedByUrl sets UpdatedByUrl field to given value.

### HasUpdatedByUrl

`func (o *RepositoryToken) HasUpdatedByUrl() bool`

HasUpdatedByUrl returns a boolean if a field has been set.

### SetUpdatedByUrlNil

`func (o *RepositoryToken) SetUpdatedByUrlNil(b bool)`

 SetUpdatedByUrlNil sets the value for UpdatedByUrl to be an explicit nil

### UnsetUpdatedByUrl
`func (o *RepositoryToken) UnsetUpdatedByUrl()`

UnsetUpdatedByUrl ensures that no value is present for UpdatedByUrl, not even an explicit nil
### GetUsage

`func (o *RepositoryToken) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *RepositoryToken) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *RepositoryToken) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *RepositoryToken) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUser

`func (o *RepositoryToken) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RepositoryToken) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RepositoryToken) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RepositoryToken) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *RepositoryToken) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *RepositoryToken) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUserUrl

`func (o *RepositoryToken) GetUserUrl() string`

GetUserUrl returns the UserUrl field if non-nil, zero value otherwise.

### GetUserUrlOk

`func (o *RepositoryToken) GetUserUrlOk() (*string, bool)`

GetUserUrlOk returns a tuple with the UserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUrl

`func (o *RepositoryToken) SetUserUrl(v string)`

SetUserUrl sets UserUrl field to given value.

### HasUserUrl

`func (o *RepositoryToken) HasUserUrl() bool`

HasUserUrl returns a boolean if a field has been set.

### SetUserUrlNil

`func (o *RepositoryToken) SetUserUrlNil(b bool)`

 SetUserUrlNil sets the value for UserUrl to be an explicit nil

### UnsetUserUrl
`func (o *RepositoryToken) UnsetUserUrl()`

UnsetUserUrl ensures that no value is present for UserUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


