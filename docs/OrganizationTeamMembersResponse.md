# OrganizationTeamMembersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**[]OrganizationTeamMembership**](OrganizationTeamMembership.md) | The team members | 

## Methods

### NewOrganizationTeamMembersResponse

`func NewOrganizationTeamMembersResponse(members []OrganizationTeamMembership, ) *OrganizationTeamMembersResponse`

NewOrganizationTeamMembersResponse instantiates a new OrganizationTeamMembersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationTeamMembersResponseWithDefaults

`func NewOrganizationTeamMembersResponseWithDefaults() *OrganizationTeamMembersResponse`

NewOrganizationTeamMembersResponseWithDefaults instantiates a new OrganizationTeamMembersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *OrganizationTeamMembersResponse) GetMembers() []OrganizationTeamMembership`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *OrganizationTeamMembersResponse) GetMembersOk() (*[]OrganizationTeamMembership, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *OrganizationTeamMembersResponse) SetMembers(v []OrganizationTeamMembership)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


