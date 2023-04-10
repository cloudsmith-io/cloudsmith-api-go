# PackageDependenciesSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dependencies** | [**[]PackageDependency**](PackageDependency.md) |  | 

## Methods

### NewPackageDependenciesSchema

`func NewPackageDependenciesSchema(dependencies []PackageDependency, ) *PackageDependenciesSchema`

NewPackageDependenciesSchema instantiates a new PackageDependenciesSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageDependenciesSchemaWithDefaults

`func NewPackageDependenciesSchemaWithDefaults() *PackageDependenciesSchema`

NewPackageDependenciesSchemaWithDefaults instantiates a new PackageDependenciesSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependencies

`func (o *PackageDependenciesSchema) GetDependencies() []PackageDependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *PackageDependenciesSchema) GetDependenciesOk() (*[]PackageDependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *PackageDependenciesSchema) SetDependencies(v []PackageDependency)`

SetDependencies sets Dependencies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


