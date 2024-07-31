# OrganizationTeamInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | The role to be assigned to the invited user in the team. | [optional] [default to "Member"]
**Team** | **string** | The team identifier (slug). | 

## Methods

### NewOrganizationTeamInvite

`func NewOrganizationTeamInvite(team string, ) *OrganizationTeamInvite`

NewOrganizationTeamInvite instantiates a new OrganizationTeamInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationTeamInviteWithDefaults

`func NewOrganizationTeamInviteWithDefaults() *OrganizationTeamInvite`

NewOrganizationTeamInviteWithDefaults instantiates a new OrganizationTeamInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *OrganizationTeamInvite) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationTeamInvite) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationTeamInvite) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationTeamInvite) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTeam

`func (o *OrganizationTeamInvite) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *OrganizationTeamInvite) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *OrganizationTeamInvite) SetTeam(v string)`

SetTeam sets Team field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


