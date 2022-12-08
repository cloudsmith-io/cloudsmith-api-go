# OrganizationTeamResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Slug** | Pointer to **string** |  | [optional] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**Visibility** | Pointer to **string** |  | [optional] [default to "Visible"]

## Methods

### NewOrganizationTeamResponse

`func NewOrganizationTeamResponse(name string, ) *OrganizationTeamResponse`

NewOrganizationTeamResponse instantiates a new OrganizationTeamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationTeamResponseWithDefaults

`func NewOrganizationTeamResponseWithDefaults() *OrganizationTeamResponse`

NewOrganizationTeamResponseWithDefaults instantiates a new OrganizationTeamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OrganizationTeamResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationTeamResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationTeamResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationTeamResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *OrganizationTeamResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationTeamResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationTeamResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *OrganizationTeamResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OrganizationTeamResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OrganizationTeamResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *OrganizationTeamResponse) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSlugPerm

`func (o *OrganizationTeamResponse) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *OrganizationTeamResponse) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *OrganizationTeamResponse) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *OrganizationTeamResponse) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetVisibility

`func (o *OrganizationTeamResponse) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OrganizationTeamResponse) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OrganizationTeamResponse) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *OrganizationTeamResponse) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


