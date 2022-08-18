# OrgsSamlGroupSyncCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpKey** | **string** | None | 
**IdpValue** | **string** | None | 
**Organization** | **string** | None | 
**Role** | Pointer to **string** |         User role within the team.          A &#x60;manager&#x60; is capable of adding/removing others to/from the team, and         can set the role of other users and other settings pertaining to the         team.          A &#39;member&#39; is a normal user that inherits the settings and privileges         assigned to the team.          | [optional] 
**Team** | **string** | None | 

## Methods

### NewOrgsSamlGroupSyncCreate

`func NewOrgsSamlGroupSyncCreate(idpKey string, idpValue string, organization string, team string, ) *OrgsSamlGroupSyncCreate`

NewOrgsSamlGroupSyncCreate instantiates a new OrgsSamlGroupSyncCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsSamlGroupSyncCreateWithDefaults

`func NewOrgsSamlGroupSyncCreateWithDefaults() *OrgsSamlGroupSyncCreate`

NewOrgsSamlGroupSyncCreateWithDefaults instantiates a new OrgsSamlGroupSyncCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpKey

`func (o *OrgsSamlGroupSyncCreate) GetIdpKey() string`

GetIdpKey returns the IdpKey field if non-nil, zero value otherwise.

### GetIdpKeyOk

`func (o *OrgsSamlGroupSyncCreate) GetIdpKeyOk() (*string, bool)`

GetIdpKeyOk returns a tuple with the IdpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpKey

`func (o *OrgsSamlGroupSyncCreate) SetIdpKey(v string)`

SetIdpKey sets IdpKey field to given value.


### GetIdpValue

`func (o *OrgsSamlGroupSyncCreate) GetIdpValue() string`

GetIdpValue returns the IdpValue field if non-nil, zero value otherwise.

### GetIdpValueOk

`func (o *OrgsSamlGroupSyncCreate) GetIdpValueOk() (*string, bool)`

GetIdpValueOk returns a tuple with the IdpValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpValue

`func (o *OrgsSamlGroupSyncCreate) SetIdpValue(v string)`

SetIdpValue sets IdpValue field to given value.


### GetOrganization

`func (o *OrgsSamlGroupSyncCreate) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrgsSamlGroupSyncCreate) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrgsSamlGroupSyncCreate) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetRole

`func (o *OrgsSamlGroupSyncCreate) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrgsSamlGroupSyncCreate) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrgsSamlGroupSyncCreate) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrgsSamlGroupSyncCreate) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTeam

`func (o *OrgsSamlGroupSyncCreate) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *OrgsSamlGroupSyncCreate) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *OrgsSamlGroupSyncCreate) SetTeam(v string)`

SetTeam sets Team field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


