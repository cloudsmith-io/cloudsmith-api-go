# OrganizationPackageLicensePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUnknownLicenses** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**OnViolationQuarantine** | Pointer to **bool** |  | [optional] 
**PackageQueryString** | Pointer to **NullableString** |  | [optional] 
**SpdxIdentifiers** | **[]string** |  | 

## Methods

### NewOrganizationPackageLicensePolicyRequest

`func NewOrganizationPackageLicensePolicyRequest(name string, spdxIdentifiers []string, ) *OrganizationPackageLicensePolicyRequest`

NewOrganizationPackageLicensePolicyRequest instantiates a new OrganizationPackageLicensePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationPackageLicensePolicyRequestWithDefaults

`func NewOrganizationPackageLicensePolicyRequestWithDefaults() *OrganizationPackageLicensePolicyRequest`

NewOrganizationPackageLicensePolicyRequestWithDefaults instantiates a new OrganizationPackageLicensePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUnknownLicenses

`func (o *OrganizationPackageLicensePolicyRequest) GetAllowUnknownLicenses() bool`

GetAllowUnknownLicenses returns the AllowUnknownLicenses field if non-nil, zero value otherwise.

### GetAllowUnknownLicensesOk

`func (o *OrganizationPackageLicensePolicyRequest) GetAllowUnknownLicensesOk() (*bool, bool)`

GetAllowUnknownLicensesOk returns a tuple with the AllowUnknownLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnknownLicenses

`func (o *OrganizationPackageLicensePolicyRequest) SetAllowUnknownLicenses(v bool)`

SetAllowUnknownLicenses sets AllowUnknownLicenses field to given value.

### HasAllowUnknownLicenses

`func (o *OrganizationPackageLicensePolicyRequest) HasAllowUnknownLicenses() bool`

HasAllowUnknownLicenses returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationPackageLicensePolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationPackageLicensePolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationPackageLicensePolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationPackageLicensePolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OrganizationPackageLicensePolicyRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OrganizationPackageLicensePolicyRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *OrganizationPackageLicensePolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationPackageLicensePolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationPackageLicensePolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOnViolationQuarantine

`func (o *OrganizationPackageLicensePolicyRequest) GetOnViolationQuarantine() bool`

GetOnViolationQuarantine returns the OnViolationQuarantine field if non-nil, zero value otherwise.

### GetOnViolationQuarantineOk

`func (o *OrganizationPackageLicensePolicyRequest) GetOnViolationQuarantineOk() (*bool, bool)`

GetOnViolationQuarantineOk returns a tuple with the OnViolationQuarantine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnViolationQuarantine

`func (o *OrganizationPackageLicensePolicyRequest) SetOnViolationQuarantine(v bool)`

SetOnViolationQuarantine sets OnViolationQuarantine field to given value.

### HasOnViolationQuarantine

`func (o *OrganizationPackageLicensePolicyRequest) HasOnViolationQuarantine() bool`

HasOnViolationQuarantine returns a boolean if a field has been set.

### GetPackageQueryString

`func (o *OrganizationPackageLicensePolicyRequest) GetPackageQueryString() string`

GetPackageQueryString returns the PackageQueryString field if non-nil, zero value otherwise.

### GetPackageQueryStringOk

`func (o *OrganizationPackageLicensePolicyRequest) GetPackageQueryStringOk() (*string, bool)`

GetPackageQueryStringOk returns a tuple with the PackageQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageQueryString

`func (o *OrganizationPackageLicensePolicyRequest) SetPackageQueryString(v string)`

SetPackageQueryString sets PackageQueryString field to given value.

### HasPackageQueryString

`func (o *OrganizationPackageLicensePolicyRequest) HasPackageQueryString() bool`

HasPackageQueryString returns a boolean if a field has been set.

### SetPackageQueryStringNil

`func (o *OrganizationPackageLicensePolicyRequest) SetPackageQueryStringNil(b bool)`

 SetPackageQueryStringNil sets the value for PackageQueryString to be an explicit nil

### UnsetPackageQueryString
`func (o *OrganizationPackageLicensePolicyRequest) UnsetPackageQueryString()`

UnsetPackageQueryString ensures that no value is present for PackageQueryString, not even an explicit nil
### GetSpdxIdentifiers

`func (o *OrganizationPackageLicensePolicyRequest) GetSpdxIdentifiers() []string`

GetSpdxIdentifiers returns the SpdxIdentifiers field if non-nil, zero value otherwise.

### GetSpdxIdentifiersOk

`func (o *OrganizationPackageLicensePolicyRequest) GetSpdxIdentifiersOk() (*[]string, bool)`

GetSpdxIdentifiersOk returns a tuple with the SpdxIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpdxIdentifiers

`func (o *OrganizationPackageLicensePolicyRequest) SetSpdxIdentifiers(v []string)`

SetSpdxIdentifiers sets SpdxIdentifiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


