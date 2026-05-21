# ConnectedRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The date and time when the connection was created. | [optional] [readonly] 
**IsActive** | Pointer to **bool** |  | [optional] [default to true]
**Priority** | Pointer to **int64** | Repositories are checked in ascending order (starting at 1). If multiple repositories have the same priority, the oldest one is used first. | [optional] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**TargetRepository** | **string** | The slug of the target repository to connect to. | 

## Methods

### NewConnectedRepository

`func NewConnectedRepository(targetRepository string, ) *ConnectedRepository`

NewConnectedRepository instantiates a new ConnectedRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedRepositoryWithDefaults

`func NewConnectedRepositoryWithDefaults() *ConnectedRepository`

NewConnectedRepositoryWithDefaults instantiates a new ConnectedRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ConnectedRepository) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConnectedRepository) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConnectedRepository) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConnectedRepository) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIsActive

`func (o *ConnectedRepository) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ConnectedRepository) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ConnectedRepository) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ConnectedRepository) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetPriority

`func (o *ConnectedRepository) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ConnectedRepository) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ConnectedRepository) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ConnectedRepository) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSlugPerm

`func (o *ConnectedRepository) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *ConnectedRepository) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *ConnectedRepository) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *ConnectedRepository) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetTargetRepository

`func (o *ConnectedRepository) GetTargetRepository() string`

GetTargetRepository returns the TargetRepository field if non-nil, zero value otherwise.

### GetTargetRepositoryOk

`func (o *ConnectedRepository) GetTargetRepositoryOk() (*string, bool)`

GetTargetRepositoryOk returns a tuple with the TargetRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRepository

`func (o *ConnectedRepository) SetTargetRepository(v string)`

SetTargetRepository sets TargetRepository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


