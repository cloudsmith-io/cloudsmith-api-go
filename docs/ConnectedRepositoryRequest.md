# ConnectedRepositoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **bool** |  | [optional] [default to true]
**Priority** | Pointer to **int64** | Repositories are checked in ascending order (starting at 1). If multiple repositories have the same priority, the oldest one is used first. | [optional] 
**TargetRepository** | **string** | The slug of the target repository to connect to. | 

## Methods

### NewConnectedRepositoryRequest

`func NewConnectedRepositoryRequest(targetRepository string, ) *ConnectedRepositoryRequest`

NewConnectedRepositoryRequest instantiates a new ConnectedRepositoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedRepositoryRequestWithDefaults

`func NewConnectedRepositoryRequestWithDefaults() *ConnectedRepositoryRequest`

NewConnectedRepositoryRequestWithDefaults instantiates a new ConnectedRepositoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *ConnectedRepositoryRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ConnectedRepositoryRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ConnectedRepositoryRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ConnectedRepositoryRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetPriority

`func (o *ConnectedRepositoryRequest) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ConnectedRepositoryRequest) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ConnectedRepositoryRequest) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ConnectedRepositoryRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTargetRepository

`func (o *ConnectedRepositoryRequest) GetTargetRepository() string`

GetTargetRepository returns the TargetRepository field if non-nil, zero value otherwise.

### GetTargetRepositoryOk

`func (o *ConnectedRepositoryRequest) GetTargetRepositoryOk() (*string, bool)`

GetTargetRepositoryOk returns a tuple with the TargetRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRepository

`func (o *ConnectedRepositoryRequest) SetTargetRepository(v string)`

SetTargetRepository sets TargetRepository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


