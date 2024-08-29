# RepositoryRetentionRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetentionCountLimit** | Pointer to **int64** | The maximum X number of packages to retain. | [optional] 
**RetentionDaysLimit** | Pointer to **int64** | The X number of days of packages to retain. | [optional] 
**RetentionEnabled** | Pointer to **bool** | If checked, the retention lifecycle rules will be activated for the repository. Any packages that don&#39;t match will be deleted automatically, and the rest are retained. | [optional] 
**RetentionGroupByFormat** | Pointer to **bool** | If checked, retention will apply to packages by package formats rather than across all package formats.For example, when retaining by a limit of 1 and you upload PythonPkg 1.0 and RubyPkg 1.0, no packages are deleted because they are different formats. | [optional] 
**RetentionGroupByName** | Pointer to **bool** | If checked, retention will apply to groups of packages by name rather than all packages.&lt;br&gt;For example, when retaining by a limit of 1 and you upload PkgA 1.0, PkgB 1.0 and PkgB 1.1; only PkgB 1.0 is deleted because there are two (2) PkgBs and one (1) PkgA. | [optional] 
**RetentionGroupByPackageType** | Pointer to **bool** | If checked, retention will apply to packages by package type (e.g. by binary, by source, etc.), rather than across all package types for one or more formats. &lt;br&gt;For example, when retaining by a limit of 1 and you upload DebPackage 1.0 and DebSourcePackage 1.0, no packages are deleted because they are different package types, binary and source respectively. | [optional] 
**RetentionSizeLimit** | Pointer to **int64** | The maximum X total size (in bytes) of packages to retain. | [optional] 

## Methods

### NewRepositoryRetentionRules

`func NewRepositoryRetentionRules() *RepositoryRetentionRules`

NewRepositoryRetentionRules instantiates a new RepositoryRetentionRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryRetentionRulesWithDefaults

`func NewRepositoryRetentionRulesWithDefaults() *RepositoryRetentionRules`

NewRepositoryRetentionRulesWithDefaults instantiates a new RepositoryRetentionRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetentionCountLimit

`func (o *RepositoryRetentionRules) GetRetentionCountLimit() int64`

GetRetentionCountLimit returns the RetentionCountLimit field if non-nil, zero value otherwise.

### GetRetentionCountLimitOk

`func (o *RepositoryRetentionRules) GetRetentionCountLimitOk() (*int64, bool)`

GetRetentionCountLimitOk returns a tuple with the RetentionCountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCountLimit

`func (o *RepositoryRetentionRules) SetRetentionCountLimit(v int64)`

SetRetentionCountLimit sets RetentionCountLimit field to given value.

### HasRetentionCountLimit

`func (o *RepositoryRetentionRules) HasRetentionCountLimit() bool`

HasRetentionCountLimit returns a boolean if a field has been set.

### GetRetentionDaysLimit

`func (o *RepositoryRetentionRules) GetRetentionDaysLimit() int64`

GetRetentionDaysLimit returns the RetentionDaysLimit field if non-nil, zero value otherwise.

### GetRetentionDaysLimitOk

`func (o *RepositoryRetentionRules) GetRetentionDaysLimitOk() (*int64, bool)`

GetRetentionDaysLimitOk returns a tuple with the RetentionDaysLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDaysLimit

`func (o *RepositoryRetentionRules) SetRetentionDaysLimit(v int64)`

SetRetentionDaysLimit sets RetentionDaysLimit field to given value.

### HasRetentionDaysLimit

`func (o *RepositoryRetentionRules) HasRetentionDaysLimit() bool`

HasRetentionDaysLimit returns a boolean if a field has been set.

### GetRetentionEnabled

`func (o *RepositoryRetentionRules) GetRetentionEnabled() bool`

GetRetentionEnabled returns the RetentionEnabled field if non-nil, zero value otherwise.

### GetRetentionEnabledOk

`func (o *RepositoryRetentionRules) GetRetentionEnabledOk() (*bool, bool)`

GetRetentionEnabledOk returns a tuple with the RetentionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionEnabled

`func (o *RepositoryRetentionRules) SetRetentionEnabled(v bool)`

SetRetentionEnabled sets RetentionEnabled field to given value.

### HasRetentionEnabled

`func (o *RepositoryRetentionRules) HasRetentionEnabled() bool`

HasRetentionEnabled returns a boolean if a field has been set.

### GetRetentionGroupByFormat

`func (o *RepositoryRetentionRules) GetRetentionGroupByFormat() bool`

GetRetentionGroupByFormat returns the RetentionGroupByFormat field if non-nil, zero value otherwise.

### GetRetentionGroupByFormatOk

`func (o *RepositoryRetentionRules) GetRetentionGroupByFormatOk() (*bool, bool)`

GetRetentionGroupByFormatOk returns a tuple with the RetentionGroupByFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionGroupByFormat

`func (o *RepositoryRetentionRules) SetRetentionGroupByFormat(v bool)`

SetRetentionGroupByFormat sets RetentionGroupByFormat field to given value.

### HasRetentionGroupByFormat

`func (o *RepositoryRetentionRules) HasRetentionGroupByFormat() bool`

HasRetentionGroupByFormat returns a boolean if a field has been set.

### GetRetentionGroupByName

`func (o *RepositoryRetentionRules) GetRetentionGroupByName() bool`

GetRetentionGroupByName returns the RetentionGroupByName field if non-nil, zero value otherwise.

### GetRetentionGroupByNameOk

`func (o *RepositoryRetentionRules) GetRetentionGroupByNameOk() (*bool, bool)`

GetRetentionGroupByNameOk returns a tuple with the RetentionGroupByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionGroupByName

`func (o *RepositoryRetentionRules) SetRetentionGroupByName(v bool)`

SetRetentionGroupByName sets RetentionGroupByName field to given value.

### HasRetentionGroupByName

`func (o *RepositoryRetentionRules) HasRetentionGroupByName() bool`

HasRetentionGroupByName returns a boolean if a field has been set.

### GetRetentionGroupByPackageType

`func (o *RepositoryRetentionRules) GetRetentionGroupByPackageType() bool`

GetRetentionGroupByPackageType returns the RetentionGroupByPackageType field if non-nil, zero value otherwise.

### GetRetentionGroupByPackageTypeOk

`func (o *RepositoryRetentionRules) GetRetentionGroupByPackageTypeOk() (*bool, bool)`

GetRetentionGroupByPackageTypeOk returns a tuple with the RetentionGroupByPackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionGroupByPackageType

`func (o *RepositoryRetentionRules) SetRetentionGroupByPackageType(v bool)`

SetRetentionGroupByPackageType sets RetentionGroupByPackageType field to given value.

### HasRetentionGroupByPackageType

`func (o *RepositoryRetentionRules) HasRetentionGroupByPackageType() bool`

HasRetentionGroupByPackageType returns a boolean if a field has been set.

### GetRetentionSizeLimit

`func (o *RepositoryRetentionRules) GetRetentionSizeLimit() int64`

GetRetentionSizeLimit returns the RetentionSizeLimit field if non-nil, zero value otherwise.

### GetRetentionSizeLimitOk

`func (o *RepositoryRetentionRules) GetRetentionSizeLimitOk() (*int64, bool)`

GetRetentionSizeLimitOk returns a tuple with the RetentionSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionSizeLimit

`func (o *RepositoryRetentionRules) SetRetentionSizeLimit(v int64)`

SetRetentionSizeLimit sets RetentionSizeLimit field to given value.

### HasRetentionSizeLimit

`func (o *RepositoryRetentionRules) HasRetentionSizeLimit() bool`

HasRetentionSizeLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


