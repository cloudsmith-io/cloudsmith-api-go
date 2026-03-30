# OrganizationTeamServiceMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** |  | 
**User** | **string** |  | 
**UserKind** | Pointer to **string** |  | [optional] [default to "User"]

## Methods

### NewOrganizationTeamServiceMember

`func NewOrganizationTeamServiceMember(role string, user string, ) *OrganizationTeamServiceMember`

NewOrganizationTeamServiceMember instantiates a new OrganizationTeamServiceMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationTeamServiceMemberWithDefaults

`func NewOrganizationTeamServiceMemberWithDefaults() *OrganizationTeamServiceMember`

NewOrganizationTeamServiceMemberWithDefaults instantiates a new OrganizationTeamServiceMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *OrganizationTeamServiceMember) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationTeamServiceMember) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationTeamServiceMember) SetRole(v string)`

SetRole sets Role field to given value.


### GetUser

`func (o *OrganizationTeamServiceMember) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrganizationTeamServiceMember) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrganizationTeamServiceMember) SetUser(v string)`

SetUser sets User field to given value.


### GetUserKind

`func (o *OrganizationTeamServiceMember) GetUserKind() string`

GetUserKind returns the UserKind field if non-nil, zero value otherwise.

### GetUserKindOk

`func (o *OrganizationTeamServiceMember) GetUserKindOk() (*string, bool)`

GetUserKindOk returns a tuple with the UserKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKind

`func (o *OrganizationTeamServiceMember) SetUserKind(v string)`

SetUserKind sets UserKind field to given value.

### HasUserKind

`func (o *OrganizationTeamServiceMember) HasUserKind() bool`

HasUserKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


