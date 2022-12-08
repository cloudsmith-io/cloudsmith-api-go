# OrganizationInviteResponse

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
**User** | Pointer to **string** | The slug of the user to be invited. | [optional] 
**UserUrl** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewOrganizationInviteResponse

`func NewOrganizationInviteResponse() *OrganizationInviteResponse`

NewOrganizationInviteResponse instantiates a new OrganizationInviteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInviteResponseWithDefaults

`func NewOrganizationInviteResponseWithDefaults() *OrganizationInviteResponse`

NewOrganizationInviteResponseWithDefaults instantiates a new OrganizationInviteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OrganizationInviteResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationInviteResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationInviteResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationInviteResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OrganizationInviteResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrganizationInviteResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrganizationInviteResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrganizationInviteResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetInviter

`func (o *OrganizationInviteResponse) GetInviter() string`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *OrganizationInviteResponse) GetInviterOk() (*string, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *OrganizationInviteResponse) SetInviter(v string)`

SetInviter sets Inviter field to given value.

### HasInviter

`func (o *OrganizationInviteResponse) HasInviter() bool`

HasInviter returns a boolean if a field has been set.

### GetInviterUrl

`func (o *OrganizationInviteResponse) GetInviterUrl() string`

GetInviterUrl returns the InviterUrl field if non-nil, zero value otherwise.

### GetInviterUrlOk

`func (o *OrganizationInviteResponse) GetInviterUrlOk() (*string, bool)`

GetInviterUrlOk returns a tuple with the InviterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviterUrl

`func (o *OrganizationInviteResponse) SetInviterUrl(v string)`

SetInviterUrl sets InviterUrl field to given value.

### HasInviterUrl

`func (o *OrganizationInviteResponse) HasInviterUrl() bool`

HasInviterUrl returns a boolean if a field has been set.

### SetInviterUrlNil

`func (o *OrganizationInviteResponse) SetInviterUrlNil(b bool)`

 SetInviterUrlNil sets the value for InviterUrl to be an explicit nil

### UnsetInviterUrl
`func (o *OrganizationInviteResponse) UnsetInviterUrl()`

UnsetInviterUrl ensures that no value is present for InviterUrl, not even an explicit nil
### GetOrg

`func (o *OrganizationInviteResponse) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *OrganizationInviteResponse) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *OrganizationInviteResponse) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *OrganizationInviteResponse) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRole

`func (o *OrganizationInviteResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationInviteResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationInviteResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationInviteResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSlugPerm

`func (o *OrganizationInviteResponse) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *OrganizationInviteResponse) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *OrganizationInviteResponse) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *OrganizationInviteResponse) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetUser

`func (o *OrganizationInviteResponse) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrganizationInviteResponse) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrganizationInviteResponse) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *OrganizationInviteResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserUrl

`func (o *OrganizationInviteResponse) GetUserUrl() string`

GetUserUrl returns the UserUrl field if non-nil, zero value otherwise.

### GetUserUrlOk

`func (o *OrganizationInviteResponse) GetUserUrlOk() (*string, bool)`

GetUserUrlOk returns a tuple with the UserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUrl

`func (o *OrganizationInviteResponse) SetUserUrl(v string)`

SetUserUrl sets UserUrl field to given value.

### HasUserUrl

`func (o *OrganizationInviteResponse) HasUserUrl() bool`

HasUserUrl returns a boolean if a field has been set.

### SetUserUrlNil

`func (o *OrganizationInviteResponse) SetUserUrlNil(b bool)`

 SetUserUrlNil sets the value for UserUrl to be an explicit nil

### UnsetUserUrl
`func (o *OrganizationInviteResponse) UnsetUserUrl()`

UnsetUserUrl ensures that no value is present for UserUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


