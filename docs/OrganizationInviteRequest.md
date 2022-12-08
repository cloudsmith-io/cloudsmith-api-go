# OrganizationInviteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email of the user to be invited. | [optional] 
**Role** | Pointer to **string** | The role to be assigned to the invited user. | [optional] [default to "Member"]
**User** | Pointer to **string** | The slug of the user to be invited. | [optional] 

## Methods

### NewOrganizationInviteRequest

`func NewOrganizationInviteRequest() *OrganizationInviteRequest`

NewOrganizationInviteRequest instantiates a new OrganizationInviteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInviteRequestWithDefaults

`func NewOrganizationInviteRequestWithDefaults() *OrganizationInviteRequest`

NewOrganizationInviteRequestWithDefaults instantiates a new OrganizationInviteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OrganizationInviteRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationInviteRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationInviteRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationInviteRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRole

`func (o *OrganizationInviteRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationInviteRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationInviteRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationInviteRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUser

`func (o *OrganizationInviteRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrganizationInviteRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrganizationInviteRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *OrganizationInviteRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


