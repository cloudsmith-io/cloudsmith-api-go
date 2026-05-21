# MemberTeams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Role** | Pointer to **string** |  | [optional] [readonly] [default to "Manager"]
**Slug** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewMemberTeams

`func NewMemberTeams() *MemberTeams`

NewMemberTeams instantiates a new MemberTeams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberTeamsWithDefaults

`func NewMemberTeamsWithDefaults() *MemberTeams`

NewMemberTeamsWithDefaults instantiates a new MemberTeams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MemberTeams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberTeams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberTeams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MemberTeams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *MemberTeams) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MemberTeams) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MemberTeams) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *MemberTeams) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSlug

`func (o *MemberTeams) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *MemberTeams) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *MemberTeams) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *MemberTeams) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


