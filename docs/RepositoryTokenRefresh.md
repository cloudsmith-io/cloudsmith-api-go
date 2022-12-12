# RepositoryTokenRefresh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | Pointer to **int64** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The datetime the token was updated at. | [optional] [readonly] 
**CreatedBy** | Pointer to **string** |  | [optional] [readonly] 
**CreatedByUrl** | Pointer to **string** |  | [optional] [readonly] 
**Default** | Pointer to **bool** | If selected this is the default token for this repository. | [optional] [readonly] 
**DisableUrl** | Pointer to **string** |  | [optional] [readonly] 
**Downloads** | Pointer to **int64** |  | [optional] [readonly] 
**EnableUrl** | Pointer to **string** |  | [optional] [readonly] 
**EulaAccepted** | Pointer to [**NullableEula**](Eula.md) |  | [optional] 
**EulaAcceptedAt** | Pointer to **time.Time** | The datetime the EULA was accepted at. | [optional] [readonly] 
**EulaAcceptedFrom** | Pointer to **NullableString** |  | [optional] [readonly] 
**EulaRequired** | Pointer to **bool** | If checked, a EULA acceptance is required for this token. | [optional] 
**HasLimits** | Pointer to **bool** |  | [optional] [readonly] 
**Identifier** | Pointer to **int64** |  | [optional] [readonly] 
**IsActive** | Pointer to **bool** | If enabled, the token will allow downloads based on configured restrictions (if any). | [optional] 
**IsLimited** | Pointer to **bool** |  | [optional] [readonly] 
**LimitBandwidth** | Pointer to **NullableInt64** | The maximum download bandwidth allowed for the token. Values are expressed as the selected unit of bandwidth. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.  | [optional] 
**LimitBandwidthUnit** | Pointer to **NullableString** |  | [optional] [default to "Byte"]
**LimitDateRangeFrom** | Pointer to **NullableTime** | The starting date/time the token is allowed to be used from. | [optional] 
**LimitDateRangeTo** | Pointer to **NullableTime** | The ending date/time the token is allowed to be used until. | [optional] 
**LimitNumClients** | Pointer to **NullableInt64** | The maximum number of unique clients allowed for the token. Please note that since clients are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point. | [optional] 
**LimitNumDownloads** | Pointer to **NullableInt64** | The maximum number of downloads allowed for the token. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point. | [optional] 
**LimitPackageQuery** | Pointer to **NullableString** | The package-based search query to apply to restrict downloads to. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. This will still allow access to non-package files, such as metadata. | [optional] 
**LimitPathQuery** | Pointer to **NullableString** | The path-based search query to apply to restrict downloads to. This supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. The path evaluated does not include the domain name, the namespace, the entitlement code used, the package format, etc. and it always starts with a forward slash. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**RefreshUrl** | Pointer to **string** |  | [optional] [readonly] 
**ResetUrl** | Pointer to **string** |  | [optional] [readonly] 
**ScheduledResetAt** | Pointer to **NullableTime** | The time at which the scheduled reset period has elapsed and the token limits were automatically reset to zero. | [optional] 
**ScheduledResetPeriod** | Pointer to **NullableString** |  | [optional] [default to "Never Reset"]
**SelfUrl** | Pointer to **string** |  | [optional] [readonly] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**Token** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The datetime the token was updated at. | [optional] [readonly] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] [readonly] 
**UpdatedByUrl** | Pointer to **NullableString** |  | [optional] [readonly] 
**Usage** | Pointer to **string** |  | [optional] [readonly] 
**User** | Pointer to **NullableString** |  | [optional] [readonly] 
**UserUrl** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewRepositoryTokenRefresh

`func NewRepositoryTokenRefresh() *RepositoryTokenRefresh`

NewRepositoryTokenRefresh instantiates a new RepositoryTokenRefresh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryTokenRefreshWithDefaults

`func NewRepositoryTokenRefreshWithDefaults() *RepositoryTokenRefresh`

NewRepositoryTokenRefreshWithDefaults instantiates a new RepositoryTokenRefresh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *RepositoryTokenRefresh) GetClients() int64`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *RepositoryTokenRefresh) GetClientsOk() (*int64, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *RepositoryTokenRefresh) SetClients(v int64)`

SetClients sets Clients field to given value.

### HasClients

`func (o *RepositoryTokenRefresh) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RepositoryTokenRefresh) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryTokenRefresh) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryTokenRefresh) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepositoryTokenRefresh) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RepositoryTokenRefresh) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RepositoryTokenRefresh) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RepositoryTokenRefresh) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RepositoryTokenRefresh) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUrl

`func (o *RepositoryTokenRefresh) GetCreatedByUrl() string`

GetCreatedByUrl returns the CreatedByUrl field if non-nil, zero value otherwise.

### GetCreatedByUrlOk

`func (o *RepositoryTokenRefresh) GetCreatedByUrlOk() (*string, bool)`

GetCreatedByUrlOk returns a tuple with the CreatedByUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUrl

`func (o *RepositoryTokenRefresh) SetCreatedByUrl(v string)`

SetCreatedByUrl sets CreatedByUrl field to given value.

### HasCreatedByUrl

`func (o *RepositoryTokenRefresh) HasCreatedByUrl() bool`

HasCreatedByUrl returns a boolean if a field has been set.

### GetDefault

`func (o *RepositoryTokenRefresh) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RepositoryTokenRefresh) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RepositoryTokenRefresh) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RepositoryTokenRefresh) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDisableUrl

`func (o *RepositoryTokenRefresh) GetDisableUrl() string`

GetDisableUrl returns the DisableUrl field if non-nil, zero value otherwise.

### GetDisableUrlOk

`func (o *RepositoryTokenRefresh) GetDisableUrlOk() (*string, bool)`

GetDisableUrlOk returns a tuple with the DisableUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUrl

`func (o *RepositoryTokenRefresh) SetDisableUrl(v string)`

SetDisableUrl sets DisableUrl field to given value.

### HasDisableUrl

`func (o *RepositoryTokenRefresh) HasDisableUrl() bool`

HasDisableUrl returns a boolean if a field has been set.

### GetDownloads

`func (o *RepositoryTokenRefresh) GetDownloads() int64`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *RepositoryTokenRefresh) GetDownloadsOk() (*int64, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *RepositoryTokenRefresh) SetDownloads(v int64)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *RepositoryTokenRefresh) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetEnableUrl

`func (o *RepositoryTokenRefresh) GetEnableUrl() string`

GetEnableUrl returns the EnableUrl field if non-nil, zero value otherwise.

### GetEnableUrlOk

`func (o *RepositoryTokenRefresh) GetEnableUrlOk() (*string, bool)`

GetEnableUrlOk returns a tuple with the EnableUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUrl

`func (o *RepositoryTokenRefresh) SetEnableUrl(v string)`

SetEnableUrl sets EnableUrl field to given value.

### HasEnableUrl

`func (o *RepositoryTokenRefresh) HasEnableUrl() bool`

HasEnableUrl returns a boolean if a field has been set.

### GetEulaAccepted

`func (o *RepositoryTokenRefresh) GetEulaAccepted() Eula`

GetEulaAccepted returns the EulaAccepted field if non-nil, zero value otherwise.

### GetEulaAcceptedOk

`func (o *RepositoryTokenRefresh) GetEulaAcceptedOk() (*Eula, bool)`

GetEulaAcceptedOk returns a tuple with the EulaAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaAccepted

`func (o *RepositoryTokenRefresh) SetEulaAccepted(v Eula)`

SetEulaAccepted sets EulaAccepted field to given value.

### HasEulaAccepted

`func (o *RepositoryTokenRefresh) HasEulaAccepted() bool`

HasEulaAccepted returns a boolean if a field has been set.

### SetEulaAcceptedNil

`func (o *RepositoryTokenRefresh) SetEulaAcceptedNil(b bool)`

 SetEulaAcceptedNil sets the value for EulaAccepted to be an explicit nil

### UnsetEulaAccepted
`func (o *RepositoryTokenRefresh) UnsetEulaAccepted()`

UnsetEulaAccepted ensures that no value is present for EulaAccepted, not even an explicit nil
### GetEulaAcceptedAt

`func (o *RepositoryTokenRefresh) GetEulaAcceptedAt() time.Time`

GetEulaAcceptedAt returns the EulaAcceptedAt field if non-nil, zero value otherwise.

### GetEulaAcceptedAtOk

`func (o *RepositoryTokenRefresh) GetEulaAcceptedAtOk() (*time.Time, bool)`

GetEulaAcceptedAtOk returns a tuple with the EulaAcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaAcceptedAt

`func (o *RepositoryTokenRefresh) SetEulaAcceptedAt(v time.Time)`

SetEulaAcceptedAt sets EulaAcceptedAt field to given value.

### HasEulaAcceptedAt

`func (o *RepositoryTokenRefresh) HasEulaAcceptedAt() bool`

HasEulaAcceptedAt returns a boolean if a field has been set.

### GetEulaAcceptedFrom

`func (o *RepositoryTokenRefresh) GetEulaAcceptedFrom() string`

GetEulaAcceptedFrom returns the EulaAcceptedFrom field if non-nil, zero value otherwise.

### GetEulaAcceptedFromOk

`func (o *RepositoryTokenRefresh) GetEulaAcceptedFromOk() (*string, bool)`

GetEulaAcceptedFromOk returns a tuple with the EulaAcceptedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaAcceptedFrom

`func (o *RepositoryTokenRefresh) SetEulaAcceptedFrom(v string)`

SetEulaAcceptedFrom sets EulaAcceptedFrom field to given value.

### HasEulaAcceptedFrom

`func (o *RepositoryTokenRefresh) HasEulaAcceptedFrom() bool`

HasEulaAcceptedFrom returns a boolean if a field has been set.

### SetEulaAcceptedFromNil

`func (o *RepositoryTokenRefresh) SetEulaAcceptedFromNil(b bool)`

 SetEulaAcceptedFromNil sets the value for EulaAcceptedFrom to be an explicit nil

### UnsetEulaAcceptedFrom
`func (o *RepositoryTokenRefresh) UnsetEulaAcceptedFrom()`

UnsetEulaAcceptedFrom ensures that no value is present for EulaAcceptedFrom, not even an explicit nil
### GetEulaRequired

`func (o *RepositoryTokenRefresh) GetEulaRequired() bool`

GetEulaRequired returns the EulaRequired field if non-nil, zero value otherwise.

### GetEulaRequiredOk

`func (o *RepositoryTokenRefresh) GetEulaRequiredOk() (*bool, bool)`

GetEulaRequiredOk returns a tuple with the EulaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaRequired

`func (o *RepositoryTokenRefresh) SetEulaRequired(v bool)`

SetEulaRequired sets EulaRequired field to given value.

### HasEulaRequired

`func (o *RepositoryTokenRefresh) HasEulaRequired() bool`

HasEulaRequired returns a boolean if a field has been set.

### GetHasLimits

`func (o *RepositoryTokenRefresh) GetHasLimits() bool`

GetHasLimits returns the HasLimits field if non-nil, zero value otherwise.

### GetHasLimitsOk

`func (o *RepositoryTokenRefresh) GetHasLimitsOk() (*bool, bool)`

GetHasLimitsOk returns a tuple with the HasLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLimits

`func (o *RepositoryTokenRefresh) SetHasLimits(v bool)`

SetHasLimits sets HasLimits field to given value.

### HasHasLimits

`func (o *RepositoryTokenRefresh) HasHasLimits() bool`

HasHasLimits returns a boolean if a field has been set.

### GetIdentifier

`func (o *RepositoryTokenRefresh) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *RepositoryTokenRefresh) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *RepositoryTokenRefresh) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *RepositoryTokenRefresh) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIsActive

`func (o *RepositoryTokenRefresh) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RepositoryTokenRefresh) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RepositoryTokenRefresh) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RepositoryTokenRefresh) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsLimited

`func (o *RepositoryTokenRefresh) GetIsLimited() bool`

GetIsLimited returns the IsLimited field if non-nil, zero value otherwise.

### GetIsLimitedOk

`func (o *RepositoryTokenRefresh) GetIsLimitedOk() (*bool, bool)`

GetIsLimitedOk returns a tuple with the IsLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLimited

`func (o *RepositoryTokenRefresh) SetIsLimited(v bool)`

SetIsLimited sets IsLimited field to given value.

### HasIsLimited

`func (o *RepositoryTokenRefresh) HasIsLimited() bool`

HasIsLimited returns a boolean if a field has been set.

### GetLimitBandwidth

`func (o *RepositoryTokenRefresh) GetLimitBandwidth() int64`

GetLimitBandwidth returns the LimitBandwidth field if non-nil, zero value otherwise.

### GetLimitBandwidthOk

`func (o *RepositoryTokenRefresh) GetLimitBandwidthOk() (*int64, bool)`

GetLimitBandwidthOk returns a tuple with the LimitBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBandwidth

`func (o *RepositoryTokenRefresh) SetLimitBandwidth(v int64)`

SetLimitBandwidth sets LimitBandwidth field to given value.

### HasLimitBandwidth

`func (o *RepositoryTokenRefresh) HasLimitBandwidth() bool`

HasLimitBandwidth returns a boolean if a field has been set.

### SetLimitBandwidthNil

`func (o *RepositoryTokenRefresh) SetLimitBandwidthNil(b bool)`

 SetLimitBandwidthNil sets the value for LimitBandwidth to be an explicit nil

### UnsetLimitBandwidth
`func (o *RepositoryTokenRefresh) UnsetLimitBandwidth()`

UnsetLimitBandwidth ensures that no value is present for LimitBandwidth, not even an explicit nil
### GetLimitBandwidthUnit

`func (o *RepositoryTokenRefresh) GetLimitBandwidthUnit() string`

GetLimitBandwidthUnit returns the LimitBandwidthUnit field if non-nil, zero value otherwise.

### GetLimitBandwidthUnitOk

`func (o *RepositoryTokenRefresh) GetLimitBandwidthUnitOk() (*string, bool)`

GetLimitBandwidthUnitOk returns a tuple with the LimitBandwidthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBandwidthUnit

`func (o *RepositoryTokenRefresh) SetLimitBandwidthUnit(v string)`

SetLimitBandwidthUnit sets LimitBandwidthUnit field to given value.

### HasLimitBandwidthUnit

`func (o *RepositoryTokenRefresh) HasLimitBandwidthUnit() bool`

HasLimitBandwidthUnit returns a boolean if a field has been set.

### SetLimitBandwidthUnitNil

`func (o *RepositoryTokenRefresh) SetLimitBandwidthUnitNil(b bool)`

 SetLimitBandwidthUnitNil sets the value for LimitBandwidthUnit to be an explicit nil

### UnsetLimitBandwidthUnit
`func (o *RepositoryTokenRefresh) UnsetLimitBandwidthUnit()`

UnsetLimitBandwidthUnit ensures that no value is present for LimitBandwidthUnit, not even an explicit nil
### GetLimitDateRangeFrom

`func (o *RepositoryTokenRefresh) GetLimitDateRangeFrom() time.Time`

GetLimitDateRangeFrom returns the LimitDateRangeFrom field if non-nil, zero value otherwise.

### GetLimitDateRangeFromOk

`func (o *RepositoryTokenRefresh) GetLimitDateRangeFromOk() (*time.Time, bool)`

GetLimitDateRangeFromOk returns a tuple with the LimitDateRangeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitDateRangeFrom

`func (o *RepositoryTokenRefresh) SetLimitDateRangeFrom(v time.Time)`

SetLimitDateRangeFrom sets LimitDateRangeFrom field to given value.

### HasLimitDateRangeFrom

`func (o *RepositoryTokenRefresh) HasLimitDateRangeFrom() bool`

HasLimitDateRangeFrom returns a boolean if a field has been set.

### SetLimitDateRangeFromNil

`func (o *RepositoryTokenRefresh) SetLimitDateRangeFromNil(b bool)`

 SetLimitDateRangeFromNil sets the value for LimitDateRangeFrom to be an explicit nil

### UnsetLimitDateRangeFrom
`func (o *RepositoryTokenRefresh) UnsetLimitDateRangeFrom()`

UnsetLimitDateRangeFrom ensures that no value is present for LimitDateRangeFrom, not even an explicit nil
### GetLimitDateRangeTo

`func (o *RepositoryTokenRefresh) GetLimitDateRangeTo() time.Time`

GetLimitDateRangeTo returns the LimitDateRangeTo field if non-nil, zero value otherwise.

### GetLimitDateRangeToOk

`func (o *RepositoryTokenRefresh) GetLimitDateRangeToOk() (*time.Time, bool)`

GetLimitDateRangeToOk returns a tuple with the LimitDateRangeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitDateRangeTo

`func (o *RepositoryTokenRefresh) SetLimitDateRangeTo(v time.Time)`

SetLimitDateRangeTo sets LimitDateRangeTo field to given value.

### HasLimitDateRangeTo

`func (o *RepositoryTokenRefresh) HasLimitDateRangeTo() bool`

HasLimitDateRangeTo returns a boolean if a field has been set.

### SetLimitDateRangeToNil

`func (o *RepositoryTokenRefresh) SetLimitDateRangeToNil(b bool)`

 SetLimitDateRangeToNil sets the value for LimitDateRangeTo to be an explicit nil

### UnsetLimitDateRangeTo
`func (o *RepositoryTokenRefresh) UnsetLimitDateRangeTo()`

UnsetLimitDateRangeTo ensures that no value is present for LimitDateRangeTo, not even an explicit nil
### GetLimitNumClients

`func (o *RepositoryTokenRefresh) GetLimitNumClients() int64`

GetLimitNumClients returns the LimitNumClients field if non-nil, zero value otherwise.

### GetLimitNumClientsOk

`func (o *RepositoryTokenRefresh) GetLimitNumClientsOk() (*int64, bool)`

GetLimitNumClientsOk returns a tuple with the LimitNumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNumClients

`func (o *RepositoryTokenRefresh) SetLimitNumClients(v int64)`

SetLimitNumClients sets LimitNumClients field to given value.

### HasLimitNumClients

`func (o *RepositoryTokenRefresh) HasLimitNumClients() bool`

HasLimitNumClients returns a boolean if a field has been set.

### SetLimitNumClientsNil

`func (o *RepositoryTokenRefresh) SetLimitNumClientsNil(b bool)`

 SetLimitNumClientsNil sets the value for LimitNumClients to be an explicit nil

### UnsetLimitNumClients
`func (o *RepositoryTokenRefresh) UnsetLimitNumClients()`

UnsetLimitNumClients ensures that no value is present for LimitNumClients, not even an explicit nil
### GetLimitNumDownloads

`func (o *RepositoryTokenRefresh) GetLimitNumDownloads() int64`

GetLimitNumDownloads returns the LimitNumDownloads field if non-nil, zero value otherwise.

### GetLimitNumDownloadsOk

`func (o *RepositoryTokenRefresh) GetLimitNumDownloadsOk() (*int64, bool)`

GetLimitNumDownloadsOk returns a tuple with the LimitNumDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNumDownloads

`func (o *RepositoryTokenRefresh) SetLimitNumDownloads(v int64)`

SetLimitNumDownloads sets LimitNumDownloads field to given value.

### HasLimitNumDownloads

`func (o *RepositoryTokenRefresh) HasLimitNumDownloads() bool`

HasLimitNumDownloads returns a boolean if a field has been set.

### SetLimitNumDownloadsNil

`func (o *RepositoryTokenRefresh) SetLimitNumDownloadsNil(b bool)`

 SetLimitNumDownloadsNil sets the value for LimitNumDownloads to be an explicit nil

### UnsetLimitNumDownloads
`func (o *RepositoryTokenRefresh) UnsetLimitNumDownloads()`

UnsetLimitNumDownloads ensures that no value is present for LimitNumDownloads, not even an explicit nil
### GetLimitPackageQuery

`func (o *RepositoryTokenRefresh) GetLimitPackageQuery() string`

GetLimitPackageQuery returns the LimitPackageQuery field if non-nil, zero value otherwise.

### GetLimitPackageQueryOk

`func (o *RepositoryTokenRefresh) GetLimitPackageQueryOk() (*string, bool)`

GetLimitPackageQueryOk returns a tuple with the LimitPackageQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPackageQuery

`func (o *RepositoryTokenRefresh) SetLimitPackageQuery(v string)`

SetLimitPackageQuery sets LimitPackageQuery field to given value.

### HasLimitPackageQuery

`func (o *RepositoryTokenRefresh) HasLimitPackageQuery() bool`

HasLimitPackageQuery returns a boolean if a field has been set.

### SetLimitPackageQueryNil

`func (o *RepositoryTokenRefresh) SetLimitPackageQueryNil(b bool)`

 SetLimitPackageQueryNil sets the value for LimitPackageQuery to be an explicit nil

### UnsetLimitPackageQuery
`func (o *RepositoryTokenRefresh) UnsetLimitPackageQuery()`

UnsetLimitPackageQuery ensures that no value is present for LimitPackageQuery, not even an explicit nil
### GetLimitPathQuery

`func (o *RepositoryTokenRefresh) GetLimitPathQuery() string`

GetLimitPathQuery returns the LimitPathQuery field if non-nil, zero value otherwise.

### GetLimitPathQueryOk

`func (o *RepositoryTokenRefresh) GetLimitPathQueryOk() (*string, bool)`

GetLimitPathQueryOk returns a tuple with the LimitPathQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPathQuery

`func (o *RepositoryTokenRefresh) SetLimitPathQuery(v string)`

SetLimitPathQuery sets LimitPathQuery field to given value.

### HasLimitPathQuery

`func (o *RepositoryTokenRefresh) HasLimitPathQuery() bool`

HasLimitPathQuery returns a boolean if a field has been set.

### SetLimitPathQueryNil

`func (o *RepositoryTokenRefresh) SetLimitPathQueryNil(b bool)`

 SetLimitPathQueryNil sets the value for LimitPathQuery to be an explicit nil

### UnsetLimitPathQuery
`func (o *RepositoryTokenRefresh) UnsetLimitPathQuery()`

UnsetLimitPathQuery ensures that no value is present for LimitPathQuery, not even an explicit nil
### GetMetadata

`func (o *RepositoryTokenRefresh) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RepositoryTokenRefresh) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RepositoryTokenRefresh) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RepositoryTokenRefresh) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *RepositoryTokenRefresh) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *RepositoryTokenRefresh) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *RepositoryTokenRefresh) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryTokenRefresh) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryTokenRefresh) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RepositoryTokenRefresh) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRefreshUrl

`func (o *RepositoryTokenRefresh) GetRefreshUrl() string`

GetRefreshUrl returns the RefreshUrl field if non-nil, zero value otherwise.

### GetRefreshUrlOk

`func (o *RepositoryTokenRefresh) GetRefreshUrlOk() (*string, bool)`

GetRefreshUrlOk returns a tuple with the RefreshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshUrl

`func (o *RepositoryTokenRefresh) SetRefreshUrl(v string)`

SetRefreshUrl sets RefreshUrl field to given value.

### HasRefreshUrl

`func (o *RepositoryTokenRefresh) HasRefreshUrl() bool`

HasRefreshUrl returns a boolean if a field has been set.

### GetResetUrl

`func (o *RepositoryTokenRefresh) GetResetUrl() string`

GetResetUrl returns the ResetUrl field if non-nil, zero value otherwise.

### GetResetUrlOk

`func (o *RepositoryTokenRefresh) GetResetUrlOk() (*string, bool)`

GetResetUrlOk returns a tuple with the ResetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetUrl

`func (o *RepositoryTokenRefresh) SetResetUrl(v string)`

SetResetUrl sets ResetUrl field to given value.

### HasResetUrl

`func (o *RepositoryTokenRefresh) HasResetUrl() bool`

HasResetUrl returns a boolean if a field has been set.

### GetScheduledResetAt

`func (o *RepositoryTokenRefresh) GetScheduledResetAt() time.Time`

GetScheduledResetAt returns the ScheduledResetAt field if non-nil, zero value otherwise.

### GetScheduledResetAtOk

`func (o *RepositoryTokenRefresh) GetScheduledResetAtOk() (*time.Time, bool)`

GetScheduledResetAtOk returns a tuple with the ScheduledResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledResetAt

`func (o *RepositoryTokenRefresh) SetScheduledResetAt(v time.Time)`

SetScheduledResetAt sets ScheduledResetAt field to given value.

### HasScheduledResetAt

`func (o *RepositoryTokenRefresh) HasScheduledResetAt() bool`

HasScheduledResetAt returns a boolean if a field has been set.

### SetScheduledResetAtNil

`func (o *RepositoryTokenRefresh) SetScheduledResetAtNil(b bool)`

 SetScheduledResetAtNil sets the value for ScheduledResetAt to be an explicit nil

### UnsetScheduledResetAt
`func (o *RepositoryTokenRefresh) UnsetScheduledResetAt()`

UnsetScheduledResetAt ensures that no value is present for ScheduledResetAt, not even an explicit nil
### GetScheduledResetPeriod

`func (o *RepositoryTokenRefresh) GetScheduledResetPeriod() string`

GetScheduledResetPeriod returns the ScheduledResetPeriod field if non-nil, zero value otherwise.

### GetScheduledResetPeriodOk

`func (o *RepositoryTokenRefresh) GetScheduledResetPeriodOk() (*string, bool)`

GetScheduledResetPeriodOk returns a tuple with the ScheduledResetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledResetPeriod

`func (o *RepositoryTokenRefresh) SetScheduledResetPeriod(v string)`

SetScheduledResetPeriod sets ScheduledResetPeriod field to given value.

### HasScheduledResetPeriod

`func (o *RepositoryTokenRefresh) HasScheduledResetPeriod() bool`

HasScheduledResetPeriod returns a boolean if a field has been set.

### SetScheduledResetPeriodNil

`func (o *RepositoryTokenRefresh) SetScheduledResetPeriodNil(b bool)`

 SetScheduledResetPeriodNil sets the value for ScheduledResetPeriod to be an explicit nil

### UnsetScheduledResetPeriod
`func (o *RepositoryTokenRefresh) UnsetScheduledResetPeriod()`

UnsetScheduledResetPeriod ensures that no value is present for ScheduledResetPeriod, not even an explicit nil
### GetSelfUrl

`func (o *RepositoryTokenRefresh) GetSelfUrl() string`

GetSelfUrl returns the SelfUrl field if non-nil, zero value otherwise.

### GetSelfUrlOk

`func (o *RepositoryTokenRefresh) GetSelfUrlOk() (*string, bool)`

GetSelfUrlOk returns a tuple with the SelfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfUrl

`func (o *RepositoryTokenRefresh) SetSelfUrl(v string)`

SetSelfUrl sets SelfUrl field to given value.

### HasSelfUrl

`func (o *RepositoryTokenRefresh) HasSelfUrl() bool`

HasSelfUrl returns a boolean if a field has been set.

### GetSlugPerm

`func (o *RepositoryTokenRefresh) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *RepositoryTokenRefresh) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *RepositoryTokenRefresh) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *RepositoryTokenRefresh) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetToken

`func (o *RepositoryTokenRefresh) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RepositoryTokenRefresh) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RepositoryTokenRefresh) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RepositoryTokenRefresh) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RepositoryTokenRefresh) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RepositoryTokenRefresh) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RepositoryTokenRefresh) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RepositoryTokenRefresh) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RepositoryTokenRefresh) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RepositoryTokenRefresh) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RepositoryTokenRefresh) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RepositoryTokenRefresh) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *RepositoryTokenRefresh) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *RepositoryTokenRefresh) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetUpdatedByUrl

`func (o *RepositoryTokenRefresh) GetUpdatedByUrl() string`

GetUpdatedByUrl returns the UpdatedByUrl field if non-nil, zero value otherwise.

### GetUpdatedByUrlOk

`func (o *RepositoryTokenRefresh) GetUpdatedByUrlOk() (*string, bool)`

GetUpdatedByUrlOk returns a tuple with the UpdatedByUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUrl

`func (o *RepositoryTokenRefresh) SetUpdatedByUrl(v string)`

SetUpdatedByUrl sets UpdatedByUrl field to given value.

### HasUpdatedByUrl

`func (o *RepositoryTokenRefresh) HasUpdatedByUrl() bool`

HasUpdatedByUrl returns a boolean if a field has been set.

### SetUpdatedByUrlNil

`func (o *RepositoryTokenRefresh) SetUpdatedByUrlNil(b bool)`

 SetUpdatedByUrlNil sets the value for UpdatedByUrl to be an explicit nil

### UnsetUpdatedByUrl
`func (o *RepositoryTokenRefresh) UnsetUpdatedByUrl()`

UnsetUpdatedByUrl ensures that no value is present for UpdatedByUrl, not even an explicit nil
### GetUsage

`func (o *RepositoryTokenRefresh) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *RepositoryTokenRefresh) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *RepositoryTokenRefresh) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *RepositoryTokenRefresh) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUser

`func (o *RepositoryTokenRefresh) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RepositoryTokenRefresh) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RepositoryTokenRefresh) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RepositoryTokenRefresh) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *RepositoryTokenRefresh) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *RepositoryTokenRefresh) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUserUrl

`func (o *RepositoryTokenRefresh) GetUserUrl() string`

GetUserUrl returns the UserUrl field if non-nil, zero value otherwise.

### GetUserUrlOk

`func (o *RepositoryTokenRefresh) GetUserUrlOk() (*string, bool)`

GetUserUrlOk returns a tuple with the UserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUrl

`func (o *RepositoryTokenRefresh) SetUserUrl(v string)`

SetUserUrl sets UserUrl field to given value.

### HasUserUrl

`func (o *RepositoryTokenRefresh) HasUserUrl() bool`

HasUserUrl returns a boolean if a field has been set.

### SetUserUrlNil

`func (o *RepositoryTokenRefresh) SetUserUrlNil(b bool)`

 SetUserUrlNil sets the value for UserUrl to be an explicit nil

### UnsetUserUrl
`func (o *RepositoryTokenRefresh) UnsetUserUrl()`

UnsetUserUrl ensures that no value is present for UserUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


