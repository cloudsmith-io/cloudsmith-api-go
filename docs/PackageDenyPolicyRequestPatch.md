# PackageDenyPolicyRequestPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** | Whether this rule is enabled or disabled. | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**PackageQueryString** | Pointer to **string** | Packages that match this query will trigger this deny rule. | [optional] 

## Methods

### NewPackageDenyPolicyRequestPatch

`func NewPackageDenyPolicyRequestPatch() *PackageDenyPolicyRequestPatch`

NewPackageDenyPolicyRequestPatch instantiates a new PackageDenyPolicyRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageDenyPolicyRequestPatchWithDefaults

`func NewPackageDenyPolicyRequestPatchWithDefaults() *PackageDenyPolicyRequestPatch`

NewPackageDenyPolicyRequestPatchWithDefaults instantiates a new PackageDenyPolicyRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PackageDenyPolicyRequestPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PackageDenyPolicyRequestPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PackageDenyPolicyRequestPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PackageDenyPolicyRequestPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PackageDenyPolicyRequestPatch) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PackageDenyPolicyRequestPatch) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *PackageDenyPolicyRequestPatch) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PackageDenyPolicyRequestPatch) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PackageDenyPolicyRequestPatch) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PackageDenyPolicyRequestPatch) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *PackageDenyPolicyRequestPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackageDenyPolicyRequestPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackageDenyPolicyRequestPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PackageDenyPolicyRequestPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PackageDenyPolicyRequestPatch) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PackageDenyPolicyRequestPatch) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPackageQueryString

`func (o *PackageDenyPolicyRequestPatch) GetPackageQueryString() string`

GetPackageQueryString returns the PackageQueryString field if non-nil, zero value otherwise.

### GetPackageQueryStringOk

`func (o *PackageDenyPolicyRequestPatch) GetPackageQueryStringOk() (*string, bool)`

GetPackageQueryStringOk returns a tuple with the PackageQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageQueryString

`func (o *PackageDenyPolicyRequestPatch) SetPackageQueryString(v string)`

SetPackageQueryString sets PackageQueryString field to given value.

### HasPackageQueryString

`func (o *PackageDenyPolicyRequestPatch) HasPackageQueryString() bool`

HasPackageQueryString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


