# OrganizationGroupSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpKey** | **string** |  | 
**IdpValue** | **string** |  | 
**Organization** | **string** |  | 
**Role** | Pointer to **string** |  | [optional] [default to "Member"]
**Team** | **string** |  | 

## Methods

### NewOrganizationGroupSyncRequest

`func NewOrganizationGroupSyncRequest(idpKey string, idpValue string, organization string, team string, ) *OrganizationGroupSyncRequest`

NewOrganizationGroupSyncRequest instantiates a new OrganizationGroupSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationGroupSyncRequestWithDefaults

`func NewOrganizationGroupSyncRequestWithDefaults() *OrganizationGroupSyncRequest`

NewOrganizationGroupSyncRequestWithDefaults instantiates a new OrganizationGroupSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpKey

`func (o *OrganizationGroupSyncRequest) GetIdpKey() string`

GetIdpKey returns the IdpKey field if non-nil, zero value otherwise.

### GetIdpKeyOk

`func (o *OrganizationGroupSyncRequest) GetIdpKeyOk() (*string, bool)`

GetIdpKeyOk returns a tuple with the IdpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpKey

`func (o *OrganizationGroupSyncRequest) SetIdpKey(v string)`

SetIdpKey sets IdpKey field to given value.


### GetIdpValue

`func (o *OrganizationGroupSyncRequest) GetIdpValue() string`

GetIdpValue returns the IdpValue field if non-nil, zero value otherwise.

### GetIdpValueOk

`func (o *OrganizationGroupSyncRequest) GetIdpValueOk() (*string, bool)`

GetIdpValueOk returns a tuple with the IdpValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpValue

`func (o *OrganizationGroupSyncRequest) SetIdpValue(v string)`

SetIdpValue sets IdpValue field to given value.


### GetOrganization

`func (o *OrganizationGroupSyncRequest) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrganizationGroupSyncRequest) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrganizationGroupSyncRequest) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetRole

`func (o *OrganizationGroupSyncRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationGroupSyncRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationGroupSyncRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationGroupSyncRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTeam

`func (o *OrganizationGroupSyncRequest) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *OrganizationGroupSyncRequest) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *OrganizationGroupSyncRequest) SetTeam(v string)`

SetTeam sets Team field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


