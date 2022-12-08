# ServiceTeams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | The team role associated with the service | [optional] [default to "Manager"]
**Slug** | Pointer to **string** | The teams associated with the service | [optional] 

## Methods

### NewServiceTeams

`func NewServiceTeams() *ServiceTeams`

NewServiceTeams instantiates a new ServiceTeams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTeamsWithDefaults

`func NewServiceTeamsWithDefaults() *ServiceTeams`

NewServiceTeamsWithDefaults instantiates a new ServiceTeams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ServiceTeams) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ServiceTeams) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ServiceTeams) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ServiceTeams) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSlug

`func (o *ServiceTeams) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ServiceTeams) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ServiceTeams) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ServiceTeams) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


