# PackageRecycleBin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action to perform on soft-deleted packages. | 
**Identifiers** | **[]string** | A list of soft-deleted package identifiers to action. | 
**Repository** | Pointer to **string** | The repository name to filter packages to. If not provided, the action will be performed across all accessible repositories in the workspace. | [optional] 

## Methods

### NewPackageRecycleBin

`func NewPackageRecycleBin(action string, identifiers []string, ) *PackageRecycleBin`

NewPackageRecycleBin instantiates a new PackageRecycleBin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageRecycleBinWithDefaults

`func NewPackageRecycleBinWithDefaults() *PackageRecycleBin`

NewPackageRecycleBinWithDefaults instantiates a new PackageRecycleBin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PackageRecycleBin) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PackageRecycleBin) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PackageRecycleBin) SetAction(v string)`

SetAction sets Action field to given value.


### GetIdentifiers

`func (o *PackageRecycleBin) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *PackageRecycleBin) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *PackageRecycleBin) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.


### GetRepository

`func (o *PackageRecycleBin) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PackageRecycleBin) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PackageRecycleBin) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PackageRecycleBin) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


