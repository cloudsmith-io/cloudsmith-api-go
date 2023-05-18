# OrganizationPackageLicensePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUnknownLicenses** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**OnViolationQuarantine** | Pointer to **bool** |  | [optional] 
**PackageQueryString** | Pointer to **NullableString** |  | [optional] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**SpdxIdentifiers** | **[]string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewOrganizationPackageLicensePolicy

`func NewOrganizationPackageLicensePolicy(name string, spdxIdentifiers []string, ) *OrganizationPackageLicensePolicy`

NewOrganizationPackageLicensePolicy instantiates a new OrganizationPackageLicensePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationPackageLicensePolicyWithDefaults

`func NewOrganizationPackageLicensePolicyWithDefaults() *OrganizationPackageLicensePolicy`

NewOrganizationPackageLicensePolicyWithDefaults instantiates a new OrganizationPackageLicensePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUnknownLicenses

`func (o *OrganizationPackageLicensePolicy) GetAllowUnknownLicenses() bool`

GetAllowUnknownLicenses returns the AllowUnknownLicenses field if non-nil, zero value otherwise.

### GetAllowUnknownLicensesOk

`func (o *OrganizationPackageLicensePolicy) GetAllowUnknownLicensesOk() (*bool, bool)`

GetAllowUnknownLicensesOk returns a tuple with the AllowUnknownLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnknownLicenses

`func (o *OrganizationPackageLicensePolicy) SetAllowUnknownLicenses(v bool)`

SetAllowUnknownLicenses sets AllowUnknownLicenses field to given value.

### HasAllowUnknownLicenses

`func (o *OrganizationPackageLicensePolicy) HasAllowUnknownLicenses() bool`

HasAllowUnknownLicenses returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrganizationPackageLicensePolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationPackageLicensePolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationPackageLicensePolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganizationPackageLicensePolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationPackageLicensePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationPackageLicensePolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationPackageLicensePolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationPackageLicensePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OrganizationPackageLicensePolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OrganizationPackageLicensePolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *OrganizationPackageLicensePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationPackageLicensePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationPackageLicensePolicy) SetName(v string)`

SetName sets Name field to given value.


### GetOnViolationQuarantine

`func (o *OrganizationPackageLicensePolicy) GetOnViolationQuarantine() bool`

GetOnViolationQuarantine returns the OnViolationQuarantine field if non-nil, zero value otherwise.

### GetOnViolationQuarantineOk

`func (o *OrganizationPackageLicensePolicy) GetOnViolationQuarantineOk() (*bool, bool)`

GetOnViolationQuarantineOk returns a tuple with the OnViolationQuarantine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnViolationQuarantine

`func (o *OrganizationPackageLicensePolicy) SetOnViolationQuarantine(v bool)`

SetOnViolationQuarantine sets OnViolationQuarantine field to given value.

### HasOnViolationQuarantine

`func (o *OrganizationPackageLicensePolicy) HasOnViolationQuarantine() bool`

HasOnViolationQuarantine returns a boolean if a field has been set.

### GetPackageQueryString

`func (o *OrganizationPackageLicensePolicy) GetPackageQueryString() string`

GetPackageQueryString returns the PackageQueryString field if non-nil, zero value otherwise.

### GetPackageQueryStringOk

`func (o *OrganizationPackageLicensePolicy) GetPackageQueryStringOk() (*string, bool)`

GetPackageQueryStringOk returns a tuple with the PackageQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageQueryString

`func (o *OrganizationPackageLicensePolicy) SetPackageQueryString(v string)`

SetPackageQueryString sets PackageQueryString field to given value.

### HasPackageQueryString

`func (o *OrganizationPackageLicensePolicy) HasPackageQueryString() bool`

HasPackageQueryString returns a boolean if a field has been set.

### SetPackageQueryStringNil

`func (o *OrganizationPackageLicensePolicy) SetPackageQueryStringNil(b bool)`

 SetPackageQueryStringNil sets the value for PackageQueryString to be an explicit nil

### UnsetPackageQueryString
`func (o *OrganizationPackageLicensePolicy) UnsetPackageQueryString()`

UnsetPackageQueryString ensures that no value is present for PackageQueryString, not even an explicit nil
### GetSlugPerm

`func (o *OrganizationPackageLicensePolicy) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *OrganizationPackageLicensePolicy) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *OrganizationPackageLicensePolicy) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *OrganizationPackageLicensePolicy) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetSpdxIdentifiers

`func (o *OrganizationPackageLicensePolicy) GetSpdxIdentifiers() []string`

GetSpdxIdentifiers returns the SpdxIdentifiers field if non-nil, zero value otherwise.

### GetSpdxIdentifiersOk

`func (o *OrganizationPackageLicensePolicy) GetSpdxIdentifiersOk() (*[]string, bool)`

GetSpdxIdentifiersOk returns a tuple with the SpdxIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpdxIdentifiers

`func (o *OrganizationPackageLicensePolicy) SetSpdxIdentifiers(v []string)`

SetSpdxIdentifiers sets SpdxIdentifiers field to given value.


### GetUpdatedAt

`func (o *OrganizationPackageLicensePolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationPackageLicensePolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationPackageLicensePolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationPackageLicensePolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


