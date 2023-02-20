# OrganizationGroupSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpKey** | **string** |  | 
**IdpValue** | **string** |  | 
**Role** | Pointer to **string** |  | [optional] [default to "Member"]
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**Team** | **string** |  | 

## Methods

### NewOrganizationGroupSync

`func NewOrganizationGroupSync(idpKey string, idpValue string, team string, ) *OrganizationGroupSync`

NewOrganizationGroupSync instantiates a new OrganizationGroupSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationGroupSyncWithDefaults

`func NewOrganizationGroupSyncWithDefaults() *OrganizationGroupSync`

NewOrganizationGroupSyncWithDefaults instantiates a new OrganizationGroupSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpKey

`func (o *OrganizationGroupSync) GetIdpKey() string`

GetIdpKey returns the IdpKey field if non-nil, zero value otherwise.

### GetIdpKeyOk

`func (o *OrganizationGroupSync) GetIdpKeyOk() (*string, bool)`

GetIdpKeyOk returns a tuple with the IdpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpKey

`func (o *OrganizationGroupSync) SetIdpKey(v string)`

SetIdpKey sets IdpKey field to given value.


### GetIdpValue

`func (o *OrganizationGroupSync) GetIdpValue() string`

GetIdpValue returns the IdpValue field if non-nil, zero value otherwise.

### GetIdpValueOk

`func (o *OrganizationGroupSync) GetIdpValueOk() (*string, bool)`

GetIdpValueOk returns a tuple with the IdpValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpValue

`func (o *OrganizationGroupSync) SetIdpValue(v string)`

SetIdpValue sets IdpValue field to given value.


### GetRole

`func (o *OrganizationGroupSync) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationGroupSync) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationGroupSync) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationGroupSync) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSlugPerm

`func (o *OrganizationGroupSync) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *OrganizationGroupSync) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *OrganizationGroupSync) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *OrganizationGroupSync) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetTeam

`func (o *OrganizationGroupSync) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *OrganizationGroupSync) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *OrganizationGroupSync) SetTeam(v string)`

SetTeam sets Team field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


