# RepositoryPrivilegeDict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Privilege** | **string** | The level of privilege that the user or team should be granted to the specified repository. | 
**Service** | Pointer to **string** | The service identifier (slug). | [optional] 
**Team** | Pointer to **string** | The team identifier (slug). | [optional] 
**User** | Pointer to **string** | The user identifier (slug). | [optional] 

## Methods

### NewRepositoryPrivilegeDict

`func NewRepositoryPrivilegeDict(privilege string, ) *RepositoryPrivilegeDict`

NewRepositoryPrivilegeDict instantiates a new RepositoryPrivilegeDict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryPrivilegeDictWithDefaults

`func NewRepositoryPrivilegeDictWithDefaults() *RepositoryPrivilegeDict`

NewRepositoryPrivilegeDictWithDefaults instantiates a new RepositoryPrivilegeDict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivilege

`func (o *RepositoryPrivilegeDict) GetPrivilege() string`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *RepositoryPrivilegeDict) GetPrivilegeOk() (*string, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *RepositoryPrivilegeDict) SetPrivilege(v string)`

SetPrivilege sets Privilege field to given value.


### GetService

`func (o *RepositoryPrivilegeDict) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *RepositoryPrivilegeDict) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *RepositoryPrivilegeDict) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *RepositoryPrivilegeDict) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTeam

`func (o *RepositoryPrivilegeDict) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *RepositoryPrivilegeDict) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *RepositoryPrivilegeDict) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *RepositoryPrivilegeDict) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetUser

`func (o *RepositoryPrivilegeDict) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RepositoryPrivilegeDict) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RepositoryPrivilegeDict) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RepositoryPrivilegeDict) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


