# OrganizationTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | A detailed description of the team. | [optional] 
**Name** | **string** | A descriptive name for the team. | 
**Slug** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] [default to "Visible"]

## Methods

### NewOrganizationTeamRequest

`func NewOrganizationTeamRequest(name string, ) *OrganizationTeamRequest`

NewOrganizationTeamRequest instantiates a new OrganizationTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationTeamRequestWithDefaults

`func NewOrganizationTeamRequestWithDefaults() *OrganizationTeamRequest`

NewOrganizationTeamRequestWithDefaults instantiates a new OrganizationTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OrganizationTeamRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationTeamRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationTeamRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationTeamRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OrganizationTeamRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OrganizationTeamRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *OrganizationTeamRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationTeamRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationTeamRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *OrganizationTeamRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OrganizationTeamRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OrganizationTeamRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *OrganizationTeamRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetVisibility

`func (o *OrganizationTeamRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OrganizationTeamRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OrganizationTeamRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *OrganizationTeamRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


