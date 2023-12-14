# PackageDenyPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] [readonly] [default to "Block downloads"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** | Whether this rule is enabled or disabled. | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**PackageQueryString** | Pointer to **NullableString** | Packages that match this query will trigger this deny rule. | [optional] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] [default to "In Progress"]
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPackageDenyPolicy

`func NewPackageDenyPolicy() *PackageDenyPolicy`

NewPackageDenyPolicy instantiates a new PackageDenyPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageDenyPolicyWithDefaults

`func NewPackageDenyPolicyWithDefaults() *PackageDenyPolicy`

NewPackageDenyPolicyWithDefaults instantiates a new PackageDenyPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PackageDenyPolicy) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PackageDenyPolicy) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PackageDenyPolicy) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PackageDenyPolicy) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PackageDenyPolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PackageDenyPolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PackageDenyPolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PackageDenyPolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *PackageDenyPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PackageDenyPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PackageDenyPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PackageDenyPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PackageDenyPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PackageDenyPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *PackageDenyPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PackageDenyPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PackageDenyPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PackageDenyPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *PackageDenyPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackageDenyPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackageDenyPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PackageDenyPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PackageDenyPolicy) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PackageDenyPolicy) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPackageQueryString

`func (o *PackageDenyPolicy) GetPackageQueryString() string`

GetPackageQueryString returns the PackageQueryString field if non-nil, zero value otherwise.

### GetPackageQueryStringOk

`func (o *PackageDenyPolicy) GetPackageQueryStringOk() (*string, bool)`

GetPackageQueryStringOk returns a tuple with the PackageQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageQueryString

`func (o *PackageDenyPolicy) SetPackageQueryString(v string)`

SetPackageQueryString sets PackageQueryString field to given value.

### HasPackageQueryString

`func (o *PackageDenyPolicy) HasPackageQueryString() bool`

HasPackageQueryString returns a boolean if a field has been set.

### SetPackageQueryStringNil

`func (o *PackageDenyPolicy) SetPackageQueryStringNil(b bool)`

 SetPackageQueryStringNil sets the value for PackageQueryString to be an explicit nil

### UnsetPackageQueryString
`func (o *PackageDenyPolicy) UnsetPackageQueryString()`

UnsetPackageQueryString ensures that no value is present for PackageQueryString, not even an explicit nil
### GetSlugPerm

`func (o *PackageDenyPolicy) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *PackageDenyPolicy) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *PackageDenyPolicy) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *PackageDenyPolicy) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetStatus

`func (o *PackageDenyPolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PackageDenyPolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PackageDenyPolicy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PackageDenyPolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PackageDenyPolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PackageDenyPolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PackageDenyPolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PackageDenyPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


