# OrganizationInviteUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | The role to be assigned to the invited user. | [optional] [default to "Member"]

## Methods

### NewOrganizationInviteUpdate

`func NewOrganizationInviteUpdate() *OrganizationInviteUpdate`

NewOrganizationInviteUpdate instantiates a new OrganizationInviteUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInviteUpdateWithDefaults

`func NewOrganizationInviteUpdateWithDefaults() *OrganizationInviteUpdate`

NewOrganizationInviteUpdateWithDefaults instantiates a new OrganizationInviteUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *OrganizationInviteUpdate) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationInviteUpdate) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationInviteUpdate) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationInviteUpdate) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


