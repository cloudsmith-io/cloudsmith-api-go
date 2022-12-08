# ServiceRequestPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the service | [optional] 
**Name** | Pointer to **string** | The name of the service | [optional] 
**Role** | Pointer to **string** | The role of the service. | [optional] [default to "Member"]
**Teams** | Pointer to [**[]ServiceTeams**](ServiceTeams.md) |  | [optional] 

## Methods

### NewServiceRequestPatch

`func NewServiceRequestPatch() *ServiceRequestPatch`

NewServiceRequestPatch instantiates a new ServiceRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRequestPatchWithDefaults

`func NewServiceRequestPatchWithDefaults() *ServiceRequestPatch`

NewServiceRequestPatchWithDefaults instantiates a new ServiceRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ServiceRequestPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceRequestPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceRequestPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceRequestPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ServiceRequestPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceRequestPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceRequestPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceRequestPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *ServiceRequestPatch) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ServiceRequestPatch) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ServiceRequestPatch) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ServiceRequestPatch) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTeams

`func (o *ServiceRequestPatch) GetTeams() []ServiceTeams`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ServiceRequestPatch) GetTeamsOk() (*[]ServiceTeams, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ServiceRequestPatch) SetTeams(v []ServiceTeams)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *ServiceRequestPatch) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


