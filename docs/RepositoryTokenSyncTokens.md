# RepositoryTokenSyncTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **string** | The datetime the token was updated at. | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedByUrl** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **bool** | If selected this is the default token for this repository. | [optional] 
**DisableUrl** | Pointer to **string** |  | [optional] 
**Downloads** | Pointer to **int64** |  | [optional] 
**EnableUrl** | Pointer to **string** |  | [optional] 
**EulaAccepted** | Pointer to **map[string]interface{}** |  | [optional] 
**EulaAcceptedAt** | Pointer to **string** | The datetime the EULA was accepted at. | [optional] 
**EulaAcceptedFrom** | Pointer to **string** |  | [optional] 
**EulaRequired** | Pointer to **bool** | If checked, a EULA acceptance is required for this token. | [optional] 
**HasLimits** | Pointer to **bool** |  | [optional] 
**Identifier** | Pointer to **int64** |  | [optional] 
**IsActive** | Pointer to **bool** | If enabled, the token will allow downloads based on configured restrictions (if any). | [optional] 
**IsLimited** | Pointer to **bool** |  | [optional] 
**LimitBandwidth** | Pointer to **int64** | The maximum download bandwidth allowed for the token. Values are expressed as the selected unit of bandwidth. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.  | [optional] 
**LimitBandwidthUnit** | Pointer to **string** |  | [optional] 
**LimitDateRangeFrom** | Pointer to **string** | The starting date/time the token is allowed to be used from. | [optional] 
**LimitDateRangeTo** | Pointer to **string** | The ending date/time the token is allowed to be used until. | [optional] 
**LimitNumClients** | Pointer to **int64** | The maximum number of unique clients allowed for the token. Please note that since clients are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point. | [optional] 
**LimitNumDownloads** | Pointer to **int64** | The maximum number of downloads allowed for the token. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point. | [optional] 
**LimitPackageQuery** | Pointer to **string** | The package-based search query to apply to restrict downloads to. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. This will still allow access to non-package files, such as metadata. | [optional] 
**LimitPathQuery** | Pointer to **string** | The path-based search query to apply to restrict downloads to. This supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. The path evaluated does not include the domain name, the namespace, the entitlement code used, the package format, etc. and it always starts with a forward slash. | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RefreshUrl** | Pointer to **string** |  | [optional] 
**ResetUrl** | Pointer to **string** |  | [optional] 
**ScheduledResetAt** | Pointer to **string** | The time at which the scheduled reset period has elapsed and the token limits were automatically reset to zero. | [optional] 
**ScheduledResetPeriod** | Pointer to **string** |  | [optional] 
**SelfUrl** | Pointer to **string** |  | [optional] 
**SlugPerm** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** | The datetime the token was updated at. | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**UpdatedByUrl** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**UserUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewRepositoryTokenSyncTokens

`func NewRepositoryTokenSyncTokens() *RepositoryTokenSyncTokens`

NewRepositoryTokenSyncTokens instantiates a new RepositoryTokenSyncTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryTokenSyncTokensWithDefaults

`func NewRepositoryTokenSyncTokensWithDefaults() *RepositoryTokenSyncTokens`

NewRepositoryTokenSyncTokensWithDefaults instantiates a new RepositoryTokenSyncTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *RepositoryTokenSyncTokens) GetClients() int64`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *RepositoryTokenSyncTokens) GetClientsOk() (*int64, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *RepositoryTokenSyncTokens) SetClients(v int64)`

SetClients sets Clients field to given value.

### HasClients

`func (o *RepositoryTokenSyncTokens) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RepositoryTokenSyncTokens) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryTokenSyncTokens) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryTokenSyncTokens) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepositoryTokenSyncTokens) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RepositoryTokenSyncTokens) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RepositoryTokenSyncTokens) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RepositoryTokenSyncTokens) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RepositoryTokenSyncTokens) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUrl

`func (o *RepositoryTokenSyncTokens) GetCreatedByUrl() string`

GetCreatedByUrl returns the CreatedByUrl field if non-nil, zero value otherwise.

### GetCreatedByUrlOk

`func (o *RepositoryTokenSyncTokens) GetCreatedByUrlOk() (*string, bool)`

GetCreatedByUrlOk returns a tuple with the CreatedByUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUrl

`func (o *RepositoryTokenSyncTokens) SetCreatedByUrl(v string)`

SetCreatedByUrl sets CreatedByUrl field to given value.

### HasCreatedByUrl

`func (o *RepositoryTokenSyncTokens) HasCreatedByUrl() bool`

HasCreatedByUrl returns a boolean if a field has been set.

### GetDefault

`func (o *RepositoryTokenSyncTokens) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RepositoryTokenSyncTokens) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RepositoryTokenSyncTokens) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RepositoryTokenSyncTokens) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDisableUrl

`func (o *RepositoryTokenSyncTokens) GetDisableUrl() string`

GetDisableUrl returns the DisableUrl field if non-nil, zero value otherwise.

### GetDisableUrlOk

`func (o *RepositoryTokenSyncTokens) GetDisableUrlOk() (*string, bool)`

GetDisableUrlOk returns a tuple with the DisableUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUrl

`func (o *RepositoryTokenSyncTokens) SetDisableUrl(v string)`

SetDisableUrl sets DisableUrl field to given value.

### HasDisableUrl

`func (o *RepositoryTokenSyncTokens) HasDisableUrl() bool`

HasDisableUrl returns a boolean if a field has been set.

### GetDownloads

`func (o *RepositoryTokenSyncTokens) GetDownloads() int64`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *RepositoryTokenSyncTokens) GetDownloadsOk() (*int64, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *RepositoryTokenSyncTokens) SetDownloads(v int64)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *RepositoryTokenSyncTokens) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetEnableUrl

`func (o *RepositoryTokenSyncTokens) GetEnableUrl() string`

GetEnableUrl returns the EnableUrl field if non-nil, zero value otherwise.

### GetEnableUrlOk

`func (o *RepositoryTokenSyncTokens) GetEnableUrlOk() (*string, bool)`

GetEnableUrlOk returns a tuple with the EnableUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUrl

`func (o *RepositoryTokenSyncTokens) SetEnableUrl(v string)`

SetEnableUrl sets EnableUrl field to given value.

### HasEnableUrl

`func (o *RepositoryTokenSyncTokens) HasEnableUrl() bool`

HasEnableUrl returns a boolean if a field has been set.

### GetEulaAccepted

`func (o *RepositoryTokenSyncTokens) GetEulaAccepted() map[string]interface{}`

GetEulaAccepted returns the EulaAccepted field if non-nil, zero value otherwise.

### GetEulaAcceptedOk

`func (o *RepositoryTokenSyncTokens) GetEulaAcceptedOk() (*map[string]interface{}, bool)`

GetEulaAcceptedOk returns a tuple with the EulaAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaAccepted

`func (o *RepositoryTokenSyncTokens) SetEulaAccepted(v map[string]interface{})`

SetEulaAccepted sets EulaAccepted field to given value.

### HasEulaAccepted

`func (o *RepositoryTokenSyncTokens) HasEulaAccepted() bool`

HasEulaAccepted returns a boolean if a field has been set.

### GetEulaAcceptedAt

`func (o *RepositoryTokenSyncTokens) GetEulaAcceptedAt() string`

GetEulaAcceptedAt returns the EulaAcceptedAt field if non-nil, zero value otherwise.

### GetEulaAcceptedAtOk

`func (o *RepositoryTokenSyncTokens) GetEulaAcceptedAtOk() (*string, bool)`

GetEulaAcceptedAtOk returns a tuple with the EulaAcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaAcceptedAt

`func (o *RepositoryTokenSyncTokens) SetEulaAcceptedAt(v string)`

SetEulaAcceptedAt sets EulaAcceptedAt field to given value.

### HasEulaAcceptedAt

`func (o *RepositoryTokenSyncTokens) HasEulaAcceptedAt() bool`

HasEulaAcceptedAt returns a boolean if a field has been set.

### GetEulaAcceptedFrom

`func (o *RepositoryTokenSyncTokens) GetEulaAcceptedFrom() string`

GetEulaAcceptedFrom returns the EulaAcceptedFrom field if non-nil, zero value otherwise.

### GetEulaAcceptedFromOk

`func (o *RepositoryTokenSyncTokens) GetEulaAcceptedFromOk() (*string, bool)`

GetEulaAcceptedFromOk returns a tuple with the EulaAcceptedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaAcceptedFrom

`func (o *RepositoryTokenSyncTokens) SetEulaAcceptedFrom(v string)`

SetEulaAcceptedFrom sets EulaAcceptedFrom field to given value.

### HasEulaAcceptedFrom

`func (o *RepositoryTokenSyncTokens) HasEulaAcceptedFrom() bool`

HasEulaAcceptedFrom returns a boolean if a field has been set.

### GetEulaRequired

`func (o *RepositoryTokenSyncTokens) GetEulaRequired() bool`

GetEulaRequired returns the EulaRequired field if non-nil, zero value otherwise.

### GetEulaRequiredOk

`func (o *RepositoryTokenSyncTokens) GetEulaRequiredOk() (*bool, bool)`

GetEulaRequiredOk returns a tuple with the EulaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaRequired

`func (o *RepositoryTokenSyncTokens) SetEulaRequired(v bool)`

SetEulaRequired sets EulaRequired field to given value.

### HasEulaRequired

`func (o *RepositoryTokenSyncTokens) HasEulaRequired() bool`

HasEulaRequired returns a boolean if a field has been set.

### GetHasLimits

`func (o *RepositoryTokenSyncTokens) GetHasLimits() bool`

GetHasLimits returns the HasLimits field if non-nil, zero value otherwise.

### GetHasLimitsOk

`func (o *RepositoryTokenSyncTokens) GetHasLimitsOk() (*bool, bool)`

GetHasLimitsOk returns a tuple with the HasLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLimits

`func (o *RepositoryTokenSyncTokens) SetHasLimits(v bool)`

SetHasLimits sets HasLimits field to given value.

### HasHasLimits

`func (o *RepositoryTokenSyncTokens) HasHasLimits() bool`

HasHasLimits returns a boolean if a field has been set.

### GetIdentifier

`func (o *RepositoryTokenSyncTokens) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *RepositoryTokenSyncTokens) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *RepositoryTokenSyncTokens) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *RepositoryTokenSyncTokens) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIsActive

`func (o *RepositoryTokenSyncTokens) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RepositoryTokenSyncTokens) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RepositoryTokenSyncTokens) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RepositoryTokenSyncTokens) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsLimited

`func (o *RepositoryTokenSyncTokens) GetIsLimited() bool`

GetIsLimited returns the IsLimited field if non-nil, zero value otherwise.

### GetIsLimitedOk

`func (o *RepositoryTokenSyncTokens) GetIsLimitedOk() (*bool, bool)`

GetIsLimitedOk returns a tuple with the IsLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLimited

`func (o *RepositoryTokenSyncTokens) SetIsLimited(v bool)`

SetIsLimited sets IsLimited field to given value.

### HasIsLimited

`func (o *RepositoryTokenSyncTokens) HasIsLimited() bool`

HasIsLimited returns a boolean if a field has been set.

### GetLimitBandwidth

`func (o *RepositoryTokenSyncTokens) GetLimitBandwidth() int64`

GetLimitBandwidth returns the LimitBandwidth field if non-nil, zero value otherwise.

### GetLimitBandwidthOk

`func (o *RepositoryTokenSyncTokens) GetLimitBandwidthOk() (*int64, bool)`

GetLimitBandwidthOk returns a tuple with the LimitBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBandwidth

`func (o *RepositoryTokenSyncTokens) SetLimitBandwidth(v int64)`

SetLimitBandwidth sets LimitBandwidth field to given value.

### HasLimitBandwidth

`func (o *RepositoryTokenSyncTokens) HasLimitBandwidth() bool`

HasLimitBandwidth returns a boolean if a field has been set.

### GetLimitBandwidthUnit

`func (o *RepositoryTokenSyncTokens) GetLimitBandwidthUnit() string`

GetLimitBandwidthUnit returns the LimitBandwidthUnit field if non-nil, zero value otherwise.

### GetLimitBandwidthUnitOk

`func (o *RepositoryTokenSyncTokens) GetLimitBandwidthUnitOk() (*string, bool)`

GetLimitBandwidthUnitOk returns a tuple with the LimitBandwidthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBandwidthUnit

`func (o *RepositoryTokenSyncTokens) SetLimitBandwidthUnit(v string)`

SetLimitBandwidthUnit sets LimitBandwidthUnit field to given value.

### HasLimitBandwidthUnit

`func (o *RepositoryTokenSyncTokens) HasLimitBandwidthUnit() bool`

HasLimitBandwidthUnit returns a boolean if a field has been set.

### GetLimitDateRangeFrom

`func (o *RepositoryTokenSyncTokens) GetLimitDateRangeFrom() string`

GetLimitDateRangeFrom returns the LimitDateRangeFrom field if non-nil, zero value otherwise.

### GetLimitDateRangeFromOk

`func (o *RepositoryTokenSyncTokens) GetLimitDateRangeFromOk() (*string, bool)`

GetLimitDateRangeFromOk returns a tuple with the LimitDateRangeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitDateRangeFrom

`func (o *RepositoryTokenSyncTokens) SetLimitDateRangeFrom(v string)`

SetLimitDateRangeFrom sets LimitDateRangeFrom field to given value.

### HasLimitDateRangeFrom

`func (o *RepositoryTokenSyncTokens) HasLimitDateRangeFrom() bool`

HasLimitDateRangeFrom returns a boolean if a field has been set.

### GetLimitDateRangeTo

`func (o *RepositoryTokenSyncTokens) GetLimitDateRangeTo() string`

GetLimitDateRangeTo returns the LimitDateRangeTo field if non-nil, zero value otherwise.

### GetLimitDateRangeToOk

`func (o *RepositoryTokenSyncTokens) GetLimitDateRangeToOk() (*string, bool)`

GetLimitDateRangeToOk returns a tuple with the LimitDateRangeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitDateRangeTo

`func (o *RepositoryTokenSyncTokens) SetLimitDateRangeTo(v string)`

SetLimitDateRangeTo sets LimitDateRangeTo field to given value.

### HasLimitDateRangeTo

`func (o *RepositoryTokenSyncTokens) HasLimitDateRangeTo() bool`

HasLimitDateRangeTo returns a boolean if a field has been set.

### GetLimitNumClients

`func (o *RepositoryTokenSyncTokens) GetLimitNumClients() int64`

GetLimitNumClients returns the LimitNumClients field if non-nil, zero value otherwise.

### GetLimitNumClientsOk

`func (o *RepositoryTokenSyncTokens) GetLimitNumClientsOk() (*int64, bool)`

GetLimitNumClientsOk returns a tuple with the LimitNumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNumClients

`func (o *RepositoryTokenSyncTokens) SetLimitNumClients(v int64)`

SetLimitNumClients sets LimitNumClients field to given value.

### HasLimitNumClients

`func (o *RepositoryTokenSyncTokens) HasLimitNumClients() bool`

HasLimitNumClients returns a boolean if a field has been set.

### GetLimitNumDownloads

`func (o *RepositoryTokenSyncTokens) GetLimitNumDownloads() int64`

GetLimitNumDownloads returns the LimitNumDownloads field if non-nil, zero value otherwise.

### GetLimitNumDownloadsOk

`func (o *RepositoryTokenSyncTokens) GetLimitNumDownloadsOk() (*int64, bool)`

GetLimitNumDownloadsOk returns a tuple with the LimitNumDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNumDownloads

`func (o *RepositoryTokenSyncTokens) SetLimitNumDownloads(v int64)`

SetLimitNumDownloads sets LimitNumDownloads field to given value.

### HasLimitNumDownloads

`func (o *RepositoryTokenSyncTokens) HasLimitNumDownloads() bool`

HasLimitNumDownloads returns a boolean if a field has been set.

### GetLimitPackageQuery

`func (o *RepositoryTokenSyncTokens) GetLimitPackageQuery() string`

GetLimitPackageQuery returns the LimitPackageQuery field if non-nil, zero value otherwise.

### GetLimitPackageQueryOk

`func (o *RepositoryTokenSyncTokens) GetLimitPackageQueryOk() (*string, bool)`

GetLimitPackageQueryOk returns a tuple with the LimitPackageQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPackageQuery

`func (o *RepositoryTokenSyncTokens) SetLimitPackageQuery(v string)`

SetLimitPackageQuery sets LimitPackageQuery field to given value.

### HasLimitPackageQuery

`func (o *RepositoryTokenSyncTokens) HasLimitPackageQuery() bool`

HasLimitPackageQuery returns a boolean if a field has been set.

### GetLimitPathQuery

`func (o *RepositoryTokenSyncTokens) GetLimitPathQuery() string`

GetLimitPathQuery returns the LimitPathQuery field if non-nil, zero value otherwise.

### GetLimitPathQueryOk

`func (o *RepositoryTokenSyncTokens) GetLimitPathQueryOk() (*string, bool)`

GetLimitPathQueryOk returns a tuple with the LimitPathQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPathQuery

`func (o *RepositoryTokenSyncTokens) SetLimitPathQuery(v string)`

SetLimitPathQuery sets LimitPathQuery field to given value.

### HasLimitPathQuery

`func (o *RepositoryTokenSyncTokens) HasLimitPathQuery() bool`

HasLimitPathQuery returns a boolean if a field has been set.

### GetMetadata

`func (o *RepositoryTokenSyncTokens) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RepositoryTokenSyncTokens) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RepositoryTokenSyncTokens) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RepositoryTokenSyncTokens) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *RepositoryTokenSyncTokens) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryTokenSyncTokens) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryTokenSyncTokens) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RepositoryTokenSyncTokens) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRefreshUrl

`func (o *RepositoryTokenSyncTokens) GetRefreshUrl() string`

GetRefreshUrl returns the RefreshUrl field if non-nil, zero value otherwise.

### GetRefreshUrlOk

`func (o *RepositoryTokenSyncTokens) GetRefreshUrlOk() (*string, bool)`

GetRefreshUrlOk returns a tuple with the RefreshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshUrl

`func (o *RepositoryTokenSyncTokens) SetRefreshUrl(v string)`

SetRefreshUrl sets RefreshUrl field to given value.

### HasRefreshUrl

`func (o *RepositoryTokenSyncTokens) HasRefreshUrl() bool`

HasRefreshUrl returns a boolean if a field has been set.

### GetResetUrl

`func (o *RepositoryTokenSyncTokens) GetResetUrl() string`

GetResetUrl returns the ResetUrl field if non-nil, zero value otherwise.

### GetResetUrlOk

`func (o *RepositoryTokenSyncTokens) GetResetUrlOk() (*string, bool)`

GetResetUrlOk returns a tuple with the ResetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetUrl

`func (o *RepositoryTokenSyncTokens) SetResetUrl(v string)`

SetResetUrl sets ResetUrl field to given value.

### HasResetUrl

`func (o *RepositoryTokenSyncTokens) HasResetUrl() bool`

HasResetUrl returns a boolean if a field has been set.

### GetScheduledResetAt

`func (o *RepositoryTokenSyncTokens) GetScheduledResetAt() string`

GetScheduledResetAt returns the ScheduledResetAt field if non-nil, zero value otherwise.

### GetScheduledResetAtOk

`func (o *RepositoryTokenSyncTokens) GetScheduledResetAtOk() (*string, bool)`

GetScheduledResetAtOk returns a tuple with the ScheduledResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledResetAt

`func (o *RepositoryTokenSyncTokens) SetScheduledResetAt(v string)`

SetScheduledResetAt sets ScheduledResetAt field to given value.

### HasScheduledResetAt

`func (o *RepositoryTokenSyncTokens) HasScheduledResetAt() bool`

HasScheduledResetAt returns a boolean if a field has been set.

### GetScheduledResetPeriod

`func (o *RepositoryTokenSyncTokens) GetScheduledResetPeriod() string`

GetScheduledResetPeriod returns the ScheduledResetPeriod field if non-nil, zero value otherwise.

### GetScheduledResetPeriodOk

`func (o *RepositoryTokenSyncTokens) GetScheduledResetPeriodOk() (*string, bool)`

GetScheduledResetPeriodOk returns a tuple with the ScheduledResetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledResetPeriod

`func (o *RepositoryTokenSyncTokens) SetScheduledResetPeriod(v string)`

SetScheduledResetPeriod sets ScheduledResetPeriod field to given value.

### HasScheduledResetPeriod

`func (o *RepositoryTokenSyncTokens) HasScheduledResetPeriod() bool`

HasScheduledResetPeriod returns a boolean if a field has been set.

### GetSelfUrl

`func (o *RepositoryTokenSyncTokens) GetSelfUrl() string`

GetSelfUrl returns the SelfUrl field if non-nil, zero value otherwise.

### GetSelfUrlOk

`func (o *RepositoryTokenSyncTokens) GetSelfUrlOk() (*string, bool)`

GetSelfUrlOk returns a tuple with the SelfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfUrl

`func (o *RepositoryTokenSyncTokens) SetSelfUrl(v string)`

SetSelfUrl sets SelfUrl field to given value.

### HasSelfUrl

`func (o *RepositoryTokenSyncTokens) HasSelfUrl() bool`

HasSelfUrl returns a boolean if a field has been set.

### GetSlugPerm

`func (o *RepositoryTokenSyncTokens) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *RepositoryTokenSyncTokens) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *RepositoryTokenSyncTokens) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *RepositoryTokenSyncTokens) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetToken

`func (o *RepositoryTokenSyncTokens) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RepositoryTokenSyncTokens) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RepositoryTokenSyncTokens) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RepositoryTokenSyncTokens) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RepositoryTokenSyncTokens) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RepositoryTokenSyncTokens) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RepositoryTokenSyncTokens) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RepositoryTokenSyncTokens) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RepositoryTokenSyncTokens) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RepositoryTokenSyncTokens) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RepositoryTokenSyncTokens) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RepositoryTokenSyncTokens) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedByUrl

`func (o *RepositoryTokenSyncTokens) GetUpdatedByUrl() string`

GetUpdatedByUrl returns the UpdatedByUrl field if non-nil, zero value otherwise.

### GetUpdatedByUrlOk

`func (o *RepositoryTokenSyncTokens) GetUpdatedByUrlOk() (*string, bool)`

GetUpdatedByUrlOk returns a tuple with the UpdatedByUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUrl

`func (o *RepositoryTokenSyncTokens) SetUpdatedByUrl(v string)`

SetUpdatedByUrl sets UpdatedByUrl field to given value.

### HasUpdatedByUrl

`func (o *RepositoryTokenSyncTokens) HasUpdatedByUrl() bool`

HasUpdatedByUrl returns a boolean if a field has been set.

### GetUsage

`func (o *RepositoryTokenSyncTokens) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *RepositoryTokenSyncTokens) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *RepositoryTokenSyncTokens) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *RepositoryTokenSyncTokens) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUser

`func (o *RepositoryTokenSyncTokens) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RepositoryTokenSyncTokens) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RepositoryTokenSyncTokens) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RepositoryTokenSyncTokens) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserUrl

`func (o *RepositoryTokenSyncTokens) GetUserUrl() string`

GetUserUrl returns the UserUrl field if non-nil, zero value otherwise.

### GetUserUrlOk

`func (o *RepositoryTokenSyncTokens) GetUserUrlOk() (*string, bool)`

GetUserUrlOk returns a tuple with the UserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUrl

`func (o *RepositoryTokenSyncTokens) SetUserUrl(v string)`

SetUserUrl sets UserUrl field to given value.

### HasUserUrl

`func (o *RepositoryTokenSyncTokens) HasUserUrl() bool`

HasUserUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


