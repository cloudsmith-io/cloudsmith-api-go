# RepositoryPrivilegeInputRequestPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Privileges** | Pointer to [**[]RepositoryPrivilegeDict**](RepositoryPrivilegeDict.md) | List of objects with explicit privileges to the repository. | [optional] 

## Methods

### NewRepositoryPrivilegeInputRequestPatch

`func NewRepositoryPrivilegeInputRequestPatch() *RepositoryPrivilegeInputRequestPatch`

NewRepositoryPrivilegeInputRequestPatch instantiates a new RepositoryPrivilegeInputRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryPrivilegeInputRequestPatchWithDefaults

`func NewRepositoryPrivilegeInputRequestPatchWithDefaults() *RepositoryPrivilegeInputRequestPatch`

NewRepositoryPrivilegeInputRequestPatchWithDefaults instantiates a new RepositoryPrivilegeInputRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivileges

`func (o *RepositoryPrivilegeInputRequestPatch) GetPrivileges() []RepositoryPrivilegeDict`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *RepositoryPrivilegeInputRequestPatch) GetPrivilegesOk() (*[]RepositoryPrivilegeDict, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *RepositoryPrivilegeInputRequestPatch) SetPrivileges(v []RepositoryPrivilegeDict)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *RepositoryPrivilegeInputRequestPatch) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


