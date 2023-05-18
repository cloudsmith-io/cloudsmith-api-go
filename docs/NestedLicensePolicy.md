# NestedLicensePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUnknownLicenses** | Pointer to **bool** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Description** | Pointer to **NullableString** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**OnViolationQuarantine** | Pointer to **bool** |  | [optional] [readonly] 
**PackageQueryString** | Pointer to **NullableString** |  | [optional] [readonly] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**SpdxIdentifiers** | **[]string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewNestedLicensePolicy

`func NewNestedLicensePolicy(spdxIdentifiers []string, ) *NestedLicensePolicy`

NewNestedLicensePolicy instantiates a new NestedLicensePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedLicensePolicyWithDefaults

`func NewNestedLicensePolicyWithDefaults() *NestedLicensePolicy`

NewNestedLicensePolicyWithDefaults instantiates a new NestedLicensePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUnknownLicenses

`func (o *NestedLicensePolicy) GetAllowUnknownLicenses() bool`

GetAllowUnknownLicenses returns the AllowUnknownLicenses field if non-nil, zero value otherwise.

### GetAllowUnknownLicensesOk

`func (o *NestedLicensePolicy) GetAllowUnknownLicensesOk() (*bool, bool)`

GetAllowUnknownLicensesOk returns a tuple with the AllowUnknownLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnknownLicenses

`func (o *NestedLicensePolicy) SetAllowUnknownLicenses(v bool)`

SetAllowUnknownLicenses sets AllowUnknownLicenses field to given value.

### HasAllowUnknownLicenses

`func (o *NestedLicensePolicy) HasAllowUnknownLicenses() bool`

HasAllowUnknownLicenses returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NestedLicensePolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NestedLicensePolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NestedLicensePolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NestedLicensePolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *NestedLicensePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NestedLicensePolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NestedLicensePolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NestedLicensePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NestedLicensePolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NestedLicensePolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *NestedLicensePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedLicensePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedLicensePolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NestedLicensePolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnViolationQuarantine

`func (o *NestedLicensePolicy) GetOnViolationQuarantine() bool`

GetOnViolationQuarantine returns the OnViolationQuarantine field if non-nil, zero value otherwise.

### GetOnViolationQuarantineOk

`func (o *NestedLicensePolicy) GetOnViolationQuarantineOk() (*bool, bool)`

GetOnViolationQuarantineOk returns a tuple with the OnViolationQuarantine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnViolationQuarantine

`func (o *NestedLicensePolicy) SetOnViolationQuarantine(v bool)`

SetOnViolationQuarantine sets OnViolationQuarantine field to given value.

### HasOnViolationQuarantine

`func (o *NestedLicensePolicy) HasOnViolationQuarantine() bool`

HasOnViolationQuarantine returns a boolean if a field has been set.

### GetPackageQueryString

`func (o *NestedLicensePolicy) GetPackageQueryString() string`

GetPackageQueryString returns the PackageQueryString field if non-nil, zero value otherwise.

### GetPackageQueryStringOk

`func (o *NestedLicensePolicy) GetPackageQueryStringOk() (*string, bool)`

GetPackageQueryStringOk returns a tuple with the PackageQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageQueryString

`func (o *NestedLicensePolicy) SetPackageQueryString(v string)`

SetPackageQueryString sets PackageQueryString field to given value.

### HasPackageQueryString

`func (o *NestedLicensePolicy) HasPackageQueryString() bool`

HasPackageQueryString returns a boolean if a field has been set.

### SetPackageQueryStringNil

`func (o *NestedLicensePolicy) SetPackageQueryStringNil(b bool)`

 SetPackageQueryStringNil sets the value for PackageQueryString to be an explicit nil

### UnsetPackageQueryString
`func (o *NestedLicensePolicy) UnsetPackageQueryString()`

UnsetPackageQueryString ensures that no value is present for PackageQueryString, not even an explicit nil
### GetSlugPerm

`func (o *NestedLicensePolicy) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *NestedLicensePolicy) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *NestedLicensePolicy) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *NestedLicensePolicy) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetSpdxIdentifiers

`func (o *NestedLicensePolicy) GetSpdxIdentifiers() []string`

GetSpdxIdentifiers returns the SpdxIdentifiers field if non-nil, zero value otherwise.

### GetSpdxIdentifiersOk

`func (o *NestedLicensePolicy) GetSpdxIdentifiersOk() (*[]string, bool)`

GetSpdxIdentifiersOk returns a tuple with the SpdxIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpdxIdentifiers

`func (o *NestedLicensePolicy) SetSpdxIdentifiers(v []string)`

SetSpdxIdentifiers sets SpdxIdentifiers field to given value.


### GetUpdatedAt

`func (o *NestedLicensePolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NestedLicensePolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NestedLicensePolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NestedLicensePolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *NestedLicensePolicy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedLicensePolicy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedLicensePolicy) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NestedLicensePolicy) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


