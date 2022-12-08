# OrganizationGroupSyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpKey** | **string** |  | 
**IdpValue** | **string** |  | 
**Role** | Pointer to **string** |  User role within the team.   A &#x60;manager&#x60; is capable of adding/removing others to/from the team, and  can set the role of other users and other settings pertaining to the  team.   A &#39;member&#39; is a normal user that inherits the settings and privileges  assigned to the team.  | [optional] [default to "Member"]
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**Team** | **string** |  | 

## Methods

### NewOrganizationGroupSyncResponse

`func NewOrganizationGroupSyncResponse(idpKey string, idpValue string, team string, ) *OrganizationGroupSyncResponse`

NewOrganizationGroupSyncResponse instantiates a new OrganizationGroupSyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationGroupSyncResponseWithDefaults

`func NewOrganizationGroupSyncResponseWithDefaults() *OrganizationGroupSyncResponse`

NewOrganizationGroupSyncResponseWithDefaults instantiates a new OrganizationGroupSyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpKey

`func (o *OrganizationGroupSyncResponse) GetIdpKey() string`

GetIdpKey returns the IdpKey field if non-nil, zero value otherwise.

### GetIdpKeyOk

`func (o *OrganizationGroupSyncResponse) GetIdpKeyOk() (*string, bool)`

GetIdpKeyOk returns a tuple with the IdpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpKey

`func (o *OrganizationGroupSyncResponse) SetIdpKey(v string)`

SetIdpKey sets IdpKey field to given value.


### GetIdpValue

`func (o *OrganizationGroupSyncResponse) GetIdpValue() string`

GetIdpValue returns the IdpValue field if non-nil, zero value otherwise.

### GetIdpValueOk

`func (o *OrganizationGroupSyncResponse) GetIdpValueOk() (*string, bool)`

GetIdpValueOk returns a tuple with the IdpValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpValue

`func (o *OrganizationGroupSyncResponse) SetIdpValue(v string)`

SetIdpValue sets IdpValue field to given value.


### GetRole

`func (o *OrganizationGroupSyncResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationGroupSyncResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationGroupSyncResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationGroupSyncResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSlugPerm

`func (o *OrganizationGroupSyncResponse) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *OrganizationGroupSyncResponse) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *OrganizationGroupSyncResponse) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *OrganizationGroupSyncResponse) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetTeam

`func (o *OrganizationGroupSyncResponse) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *OrganizationGroupSyncResponse) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *OrganizationGroupSyncResponse) SetTeam(v string)`

SetTeam sets Team field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


