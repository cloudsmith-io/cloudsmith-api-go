# PackageBulkAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action to perform on the packages. | 
**Identifiers** | **[]string** | A list of package identifiers to apply the action to. | 
**Repository** | Pointer to **string** | The repository name to filter packages to. If not provided, the action will be performed across all accessible repositories in the workspace. | [optional] 
**TargetRepository** | Pointer to **string** | The slug of the target repository | [optional] 

## Methods

### NewPackageBulkAction

`func NewPackageBulkAction(action string, identifiers []string, ) *PackageBulkAction`

NewPackageBulkAction instantiates a new PackageBulkAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageBulkActionWithDefaults

`func NewPackageBulkActionWithDefaults() *PackageBulkAction`

NewPackageBulkActionWithDefaults instantiates a new PackageBulkAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PackageBulkAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PackageBulkAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PackageBulkAction) SetAction(v string)`

SetAction sets Action field to given value.


### GetIdentifiers

`func (o *PackageBulkAction) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *PackageBulkAction) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *PackageBulkAction) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.


### GetRepository

`func (o *PackageBulkAction) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PackageBulkAction) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PackageBulkAction) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PackageBulkAction) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetTargetRepository

`func (o *PackageBulkAction) GetTargetRepository() string`

GetTargetRepository returns the TargetRepository field if non-nil, zero value otherwise.

### GetTargetRepositoryOk

`func (o *PackageBulkAction) GetTargetRepositoryOk() (*string, bool)`

GetTargetRepositoryOk returns a tuple with the TargetRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRepository

`func (o *PackageBulkAction) SetTargetRepository(v string)`

SetTargetRepository sets TargetRepository field to given value.

### HasTargetRepository

`func (o *PackageBulkAction) HasTargetRepository() bool`

HasTargetRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


