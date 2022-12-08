# PackageDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepType** | Pointer to **string** |  | [optional] [readonly] [default to "Depends"]
**Name** | **string** | The name of the package dependency. | 
**Operator** | Pointer to **NullableString** |  | [optional] [readonly] [default to "="]
**Version** | Pointer to **NullableString** | The raw dependency version (if any). | [optional] 

## Methods

### NewPackageDependency

`func NewPackageDependency(name string, ) *PackageDependency`

NewPackageDependency instantiates a new PackageDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageDependencyWithDefaults

`func NewPackageDependencyWithDefaults() *PackageDependency`

NewPackageDependencyWithDefaults instantiates a new PackageDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepType

`func (o *PackageDependency) GetDepType() string`

GetDepType returns the DepType field if non-nil, zero value otherwise.

### GetDepTypeOk

`func (o *PackageDependency) GetDepTypeOk() (*string, bool)`

GetDepTypeOk returns a tuple with the DepType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepType

`func (o *PackageDependency) SetDepType(v string)`

SetDepType sets DepType field to given value.

### HasDepType

`func (o *PackageDependency) HasDepType() bool`

HasDepType returns a boolean if a field has been set.

### GetName

`func (o *PackageDependency) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackageDependency) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackageDependency) SetName(v string)`

SetName sets Name field to given value.


### GetOperator

`func (o *PackageDependency) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PackageDependency) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PackageDependency) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *PackageDependency) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *PackageDependency) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *PackageDependency) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetVersion

`func (o *PackageDependency) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PackageDependency) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PackageDependency) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PackageDependency) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *PackageDependency) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *PackageDependency) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


