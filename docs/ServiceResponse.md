# ServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the service | [optional] 
**Key** | Pointer to **string** | The API key of the service | [optional] [readonly] 
**Name** | **string** | The name of the service | 
**Role** | Pointer to **string** | The role of the service. | [optional] [default to "Member"]
**Slug** | Pointer to **string** | The slug of the service | [optional] [readonly] 
**Teams** | Pointer to [**[]ServiceTeams**](ServiceTeams.md) |  | [optional] 

## Methods

### NewServiceResponse

`func NewServiceResponse(name string, ) *ServiceResponse`

NewServiceResponse instantiates a new ServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceResponseWithDefaults

`func NewServiceResponseWithDefaults() *ServiceResponse`

NewServiceResponseWithDefaults instantiates a new ServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ServiceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *ServiceResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ServiceResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ServiceResponse) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ServiceResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *ServiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *ServiceResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ServiceResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ServiceResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ServiceResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSlug

`func (o *ServiceResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ServiceResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ServiceResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ServiceResponse) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTeams

`func (o *ServiceResponse) GetTeams() []ServiceTeams`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ServiceResponse) GetTeamsOk() (*[]ServiceTeams, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ServiceResponse) SetTeams(v []ServiceTeams)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *ServiceResponse) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


