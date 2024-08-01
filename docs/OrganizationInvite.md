# OrganizationInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email of the user to be invited. | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Inviter** | Pointer to **string** |  | [optional] [readonly] 
**InviterUrl** | Pointer to **NullableString** |  | [optional] [readonly] 
**Org** | Pointer to **string** |  | [optional] [readonly] 
**Role** | Pointer to **string** | The role to be assigned to the invited user. | [optional] [default to "Member"]
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**Teams** | Pointer to [**[]OrganizationTeamInvite**](OrganizationTeamInvite.md) |  | [optional] 
**User** | Pointer to **string** | The slug of the user to be invited. | [optional] 
**UserUrl** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewOrganizationInvite

`func NewOrganizationInvite() *OrganizationInvite`

NewOrganizationInvite instantiates a new OrganizationInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInviteWithDefaults

`func NewOrganizationInviteWithDefaults() *OrganizationInvite`

NewOrganizationInviteWithDefaults instantiates a new OrganizationInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OrganizationInvite) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationInvite) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationInvite) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationInvite) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OrganizationInvite) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrganizationInvite) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrganizationInvite) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrganizationInvite) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetInviter

`func (o *OrganizationInvite) GetInviter() string`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *OrganizationInvite) GetInviterOk() (*string, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *OrganizationInvite) SetInviter(v string)`

SetInviter sets Inviter field to given value.

### HasInviter

`func (o *OrganizationInvite) HasInviter() bool`

HasInviter returns a boolean if a field has been set.

### GetInviterUrl

`func (o *OrganizationInvite) GetInviterUrl() string`

GetInviterUrl returns the InviterUrl field if non-nil, zero value otherwise.

### GetInviterUrlOk

`func (o *OrganizationInvite) GetInviterUrlOk() (*string, bool)`

GetInviterUrlOk returns a tuple with the InviterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviterUrl

`func (o *OrganizationInvite) SetInviterUrl(v string)`

SetInviterUrl sets InviterUrl field to given value.

### HasInviterUrl

`func (o *OrganizationInvite) HasInviterUrl() bool`

HasInviterUrl returns a boolean if a field has been set.

### SetInviterUrlNil

`func (o *OrganizationInvite) SetInviterUrlNil(b bool)`

 SetInviterUrlNil sets the value for InviterUrl to be an explicit nil

### UnsetInviterUrl
`func (o *OrganizationInvite) UnsetInviterUrl()`

UnsetInviterUrl ensures that no value is present for InviterUrl, not even an explicit nil
### GetOrg

`func (o *OrganizationInvite) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *OrganizationInvite) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *OrganizationInvite) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *OrganizationInvite) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRole

`func (o *OrganizationInvite) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationInvite) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationInvite) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationInvite) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSlugPerm

`func (o *OrganizationInvite) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *OrganizationInvite) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *OrganizationInvite) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *OrganizationInvite) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetTeams

`func (o *OrganizationInvite) GetTeams() []OrganizationTeamInvite`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *OrganizationInvite) GetTeamsOk() (*[]OrganizationTeamInvite, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *OrganizationInvite) SetTeams(v []OrganizationTeamInvite)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *OrganizationInvite) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetUser

`func (o *OrganizationInvite) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrganizationInvite) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrganizationInvite) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *OrganizationInvite) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserUrl

`func (o *OrganizationInvite) GetUserUrl() string`

GetUserUrl returns the UserUrl field if non-nil, zero value otherwise.

### GetUserUrlOk

`func (o *OrganizationInvite) GetUserUrlOk() (*string, bool)`

GetUserUrlOk returns a tuple with the UserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUrl

`func (o *OrganizationInvite) SetUserUrl(v string)`

SetUserUrl sets UserUrl field to given value.

### HasUserUrl

`func (o *OrganizationInvite) HasUserUrl() bool`

HasUserUrl returns a boolean if a field has been set.

### SetUserUrlNil

`func (o *OrganizationInvite) SetUserUrlNil(b bool)`

 SetUserUrlNil sets the value for UserUrl to be an explicit nil

### UnsetUserUrl
`func (o *OrganizationInvite) UnsetUserUrl()`

UnsetUserUrl ensures that no value is present for UserUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


