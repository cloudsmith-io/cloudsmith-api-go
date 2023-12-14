# PackageDenyPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** | Whether this rule is enabled or disabled. | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**PackageQueryString** | Pointer to **NullableString** | Packages that match this query will trigger this deny rule. | [optional] 

## Methods

### NewPackageDenyPolicyRequest

`func NewPackageDenyPolicyRequest() *PackageDenyPolicyRequest`

NewPackageDenyPolicyRequest instantiates a new PackageDenyPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageDenyPolicyRequestWithDefaults

`func NewPackageDenyPolicyRequestWithDefaults() *PackageDenyPolicyRequest`

NewPackageDenyPolicyRequestWithDefaults instantiates a new PackageDenyPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PackageDenyPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PackageDenyPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PackageDenyPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PackageDenyPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PackageDenyPolicyRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PackageDenyPolicyRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *PackageDenyPolicyRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PackageDenyPolicyRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PackageDenyPolicyRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PackageDenyPolicyRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *PackageDenyPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackageDenyPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackageDenyPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PackageDenyPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PackageDenyPolicyRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PackageDenyPolicyRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPackageQueryString

`func (o *PackageDenyPolicyRequest) GetPackageQueryString() string`

GetPackageQueryString returns the PackageQueryString field if non-nil, zero value otherwise.

### GetPackageQueryStringOk

`func (o *PackageDenyPolicyRequest) GetPackageQueryStringOk() (*string, bool)`

GetPackageQueryStringOk returns a tuple with the PackageQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageQueryString

`func (o *PackageDenyPolicyRequest) SetPackageQueryString(v string)`

SetPackageQueryString sets PackageQueryString field to given value.

### HasPackageQueryString

`func (o *PackageDenyPolicyRequest) HasPackageQueryString() bool`

HasPackageQueryString returns a boolean if a field has been set.

### SetPackageQueryStringNil

`func (o *PackageDenyPolicyRequest) SetPackageQueryStringNil(b bool)`

 SetPackageQueryStringNil sets the value for PackageQueryString to be an explicit nil

### UnsetPackageQueryString
`func (o *PackageDenyPolicyRequest) UnsetPackageQueryString()`

UnsetPackageQueryString ensures that no value is present for PackageQueryString, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


