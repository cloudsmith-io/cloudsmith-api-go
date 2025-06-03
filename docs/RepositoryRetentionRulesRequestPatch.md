# RepositoryRetentionRulesRequestPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetentionCountLimit** | Pointer to **int64** | The maximum X number of packages to retain. | [optional] 
**RetentionDaysLimit** | Pointer to **int64** | The X number of days of packages to retain. | [optional] 
**RetentionEnabled** | Pointer to **bool** | If checked, the retention lifecycle rules will be activated for the repository. Any packages that don&#39;t match will be deleted automatically, and the rest are retained. | [optional] 
**RetentionGroupByFormat** | Pointer to **bool** | If checked, retention will apply to packages by package formats rather than across all package formats.For example, when retaining by a limit of 1 and you upload PythonPkg 1.0 and RubyPkg 1.0, no packages are deleted because they are different formats. | [optional] 
**RetentionGroupByName** | Pointer to **bool** | If checked, retention will apply to groups of packages by name rather than all packages.&lt;br&gt;For example, when retaining by a limit of 1 and you upload PkgA 1.0, PkgB 1.0 and PkgB 1.1; only PkgB 1.0 is deleted because there are two (2) PkgBs and one (1) PkgA. | [optional] 
**RetentionGroupByPackageType** | Pointer to **bool** | If checked, retention will apply to packages by package type (e.g. by binary, by source, etc.), rather than across all package types for one or more formats. &lt;br&gt;For example, when retaining by a limit of 1 and you upload DebPackage 1.0 and DebSourcePackage 1.0, no packages are deleted because they are different package types, binary and source respectively. | [optional] 
**RetentionPackageQueryString** | Pointer to **NullableString** | A package search expression which, if provided, filters the packages to be deleted.&lt;br&gt;For example, a search expression of &#x60;name:foo&#x60; will result in only packages called &#39;foo&#39; being deleted, or a search expression of &#x60;tag:~latest&#x60; will prevent any packages tagged &#39;latest&#39; from being deleted.&lt;br&gt;Refer to the Cloudsmith documentation for package query syntax. | [optional] 
**RetentionSizeLimit** | Pointer to **int64** | The maximum X total size (in bytes) of packages to retain. | [optional] 

## Methods

### NewRepositoryRetentionRulesRequestPatch

`func NewRepositoryRetentionRulesRequestPatch() *RepositoryRetentionRulesRequestPatch`

NewRepositoryRetentionRulesRequestPatch instantiates a new RepositoryRetentionRulesRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryRetentionRulesRequestPatchWithDefaults

`func NewRepositoryRetentionRulesRequestPatchWithDefaults() *RepositoryRetentionRulesRequestPatch`

NewRepositoryRetentionRulesRequestPatchWithDefaults instantiates a new RepositoryRetentionRulesRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetentionCountLimit

`func (o *RepositoryRetentionRulesRequestPatch) GetRetentionCountLimit() int64`

GetRetentionCountLimit returns the RetentionCountLimit field if non-nil, zero value otherwise.

### GetRetentionCountLimitOk

`func (o *RepositoryRetentionRulesRequestPatch) GetRetentionCountLimitOk() (*int64, bool)`

GetRetentionCountLimitOk returns a tuple with the RetentionCountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCountLimit

`func (o *RepositoryRetentionRulesRequestPatch) SetRetentionCountLimit(v int64)`

SetRetentionCountLimit sets RetentionCountLimit field to given value.

### HasRetentionCountLimit

`func (o *RepositoryRetentionRulesRequestPatch) HasRetentionCountLimit() bool`

HasRetentionCountLimit returns a boolean if a field has been set.

### GetRetentionDaysLimit

`func (o *RepositoryRetentionRulesRequestPatch) GetRetentionDaysLimit() int64`

GetRetentionDaysLimit returns the RetentionDaysLimit field if non-nil, zero value otherwise.

### GetRetentionDaysLimitOk

`func (o *RepositoryRetentionRulesRequestPatch) GetRetentionDaysLimitOk() (*int64, bool)`

GetRetentionDaysLimitOk returns a tuple with the RetentionDaysLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDaysLimit

`func (o *RepositoryRetentionRulesRequestPatch) SetRetentionDaysLimit(v int64)`

SetRetentionDaysLimit sets RetentionDaysLimit field to given value.

### HasRetentionDaysLimit

`func (o *RepositoryRetentionRulesRequestPatch) HasRetentionDaysLimit() bool`

HasRetentionDaysLimit returns a boolean if a field has been set.

### GetRetentionEnabled

`func (o *RepositoryRetentionRulesRequestPatch) GetRetentionEnabled() bool`

GetRetentionEnabled returns the RetentionEnabled field if non-nil, zero value otherwise.

### GetRetentionEnabledOk

`func (o *RepositoryRetentionRulesRequestPatch) GetRetentionEnabledOk() (*bool, bool)`

GetRetentionEnabledOk returns a tuple with the RetentionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionEnabled

`func (o *RepositoryRetentionRulesRequestPatch) SetRetentionEnabled(v bool)`

SetRetentionEnabled sets RetentionEnabled field to given value.

### HasRetentionEnabled

`func (o *RepositoryRetentionRulesRequestPatch) HasRetentionEnabled() bool`

HasRetentionEnabled returns a boolean if a field has been set.

### GetRetentionGroupByFormat

`func (o *RepositoryRetentionRulesRequestPatch) GetRetentionGroupByFormat() bool`

GetRetentionGroupByFormat returns the RetentionGroupByFormat field if non-nil, zero value otherwise.

### GetRetentionGroupByFormatOk

`func (o *RepositoryRetentionRulesRequestPatch) GetRetentionGroupByFormatOk() (*bool, bool)`

GetRetentionGroupByFormatOk returns a tuple with the RetentionGroupByFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionGroupByFormat

`func (o *RepositoryRetentionRulesRequestPatch) SetRetentionGroupByFormat(v bool)`

SetRetentionGroupByFormat sets RetentionGroupByFormat field to given value.

### HasRetentionGroupByFormat

`func (o *RepositoryRetentionRulesRequestPatch) HasRetentionGroupByFormat() bool`

HasRetentionGroupByFormat returns a boolean if a field has been set.

### GetRetentionGroupByName

`func (o *RepositoryRetentionRulesRequestPatch) GetRetentionGroupByName() bool`

GetRetentionGroupByName returns the RetentionGroupByName field if non-nil, zero value otherwise.

### GetRetentionGroupByNameOk

`func (o *RepositoryRetentionRulesRequestPatch) GetRetentionGroupByNameOk() (*bool, bool)`

GetRetentionGroupByNameOk returns a tuple with the RetentionGroupByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionGroupByName

`func (o *RepositoryRetentionRulesRequestPatch) SetRetentionGroupByName(v bool)`

SetRetentionGroupByName sets RetentionGroupByName field to given value.

### HasRetentionGroupByName

`func (o *RepositoryRetentionRulesRequestPatch) HasRetentionGroupByName() bool`

HasRetentionGroupByName returns a boolean if a field has been set.

### GetRetentionGroupByPackageType

`func (o *RepositoryRetentionRulesRequestPatch) GetRetentionGroupByPackageType() bool`

GetRetentionGroupByPackageType returns the RetentionGroupByPackageType field if non-nil, zero value otherwise.

### GetRetentionGroupByPackageTypeOk

`func (o *RepositoryRetentionRulesRequestPatch) GetRetentionGroupByPackageTypeOk() (*bool, bool)`

GetRetentionGroupByPackageTypeOk returns a tuple with the RetentionGroupByPackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionGroupByPackageType

`func (o *RepositoryRetentionRulesRequestPatch) SetRetentionGroupByPackageType(v bool)`

SetRetentionGroupByPackageType sets RetentionGroupByPackageType field to given value.

### HasRetentionGroupByPackageType

`func (o *RepositoryRetentionRulesRequestPatch) HasRetentionGroupByPackageType() bool`

HasRetentionGroupByPackageType returns a boolean if a field has been set.

### GetRetentionPackageQueryString

`func (o *RepositoryRetentionRulesRequestPatch) GetRetentionPackageQueryString() string`

GetRetentionPackageQueryString returns the RetentionPackageQueryString field if non-nil, zero value otherwise.

### GetRetentionPackageQueryStringOk

`func (o *RepositoryRetentionRulesRequestPatch) GetRetentionPackageQueryStringOk() (*string, bool)`

GetRetentionPackageQueryStringOk returns a tuple with the RetentionPackageQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPackageQueryString

`func (o *RepositoryRetentionRulesRequestPatch) SetRetentionPackageQueryString(v string)`

SetRetentionPackageQueryString sets RetentionPackageQueryString field to given value.

### HasRetentionPackageQueryString

`func (o *RepositoryRetentionRulesRequestPatch) HasRetentionPackageQueryString() bool`

HasRetentionPackageQueryString returns a boolean if a field has been set.

### SetRetentionPackageQueryStringNil

`func (o *RepositoryRetentionRulesRequestPatch) SetRetentionPackageQueryStringNil(b bool)`

 SetRetentionPackageQueryStringNil sets the value for RetentionPackageQueryString to be an explicit nil

### UnsetRetentionPackageQueryString
`func (o *RepositoryRetentionRulesRequestPatch) UnsetRetentionPackageQueryString()`

UnsetRetentionPackageQueryString ensures that no value is present for RetentionPackageQueryString, not even an explicit nil
### GetRetentionSizeLimit

`func (o *RepositoryRetentionRulesRequestPatch) GetRetentionSizeLimit() int64`

GetRetentionSizeLimit returns the RetentionSizeLimit field if non-nil, zero value otherwise.

### GetRetentionSizeLimitOk

`func (o *RepositoryRetentionRulesRequestPatch) GetRetentionSizeLimitOk() (*int64, bool)`

GetRetentionSizeLimitOk returns a tuple with the RetentionSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionSizeLimit

`func (o *RepositoryRetentionRulesRequestPatch) SetRetentionSizeLimit(v int64)`

SetRetentionSizeLimit sets RetentionSizeLimit field to given value.

### HasRetentionSizeLimit

`func (o *RepositoryRetentionRulesRequestPatch) HasRetentionSizeLimit() bool`

HasRetentionSizeLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


