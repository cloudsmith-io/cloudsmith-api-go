# OrganizationInviteUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | The role to be assigned to the invited user. | [optional] [default to "Member"]

## Methods

### NewOrganizationInviteUpdateResponse

`func NewOrganizationInviteUpdateResponse() *OrganizationInviteUpdateResponse`

NewOrganizationInviteUpdateResponse instantiates a new OrganizationInviteUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInviteUpdateResponseWithDefaults

`func NewOrganizationInviteUpdateResponseWithDefaults() *OrganizationInviteUpdateResponse`

NewOrganizationInviteUpdateResponseWithDefaults instantiates a new OrganizationInviteUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *OrganizationInviteUpdateResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationInviteUpdateResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationInviteUpdateResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationInviteUpdateResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


