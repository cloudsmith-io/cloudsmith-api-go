# OrganizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **NullableString** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Location** | Pointer to **string** | The city/town/area your organization is based in. | [optional] [readonly] 
**Name** | Pointer to **string** | A descriptive name for your organization. | [optional] [readonly] 
**Slug** | Pointer to **string** |  | [optional] [readonly] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**Tagline** | Pointer to **string** | A short public descriptive for your organization. | [optional] [readonly] 

## Methods

### NewOrganizationResponse

`func NewOrganizationResponse() *OrganizationResponse`

NewOrganizationResponse instantiates a new OrganizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationResponseWithDefaults

`func NewOrganizationResponseWithDefaults() *OrganizationResponse`

NewOrganizationResponseWithDefaults instantiates a new OrganizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *OrganizationResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrganizationResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrganizationResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OrganizationResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *OrganizationResponse) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *OrganizationResponse) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCreatedAt

`func (o *OrganizationResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganizationResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLocation

`func (o *OrganizationResponse) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OrganizationResponse) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OrganizationResponse) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OrganizationResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *OrganizationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *OrganizationResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OrganizationResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OrganizationResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *OrganizationResponse) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSlugPerm

`func (o *OrganizationResponse) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *OrganizationResponse) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *OrganizationResponse) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *OrganizationResponse) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetTagline

`func (o *OrganizationResponse) GetTagline() string`

GetTagline returns the Tagline field if non-nil, zero value otherwise.

### GetTaglineOk

`func (o *OrganizationResponse) GetTaglineOk() (*string, bool)`

GetTaglineOk returns a tuple with the Tagline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagline

`func (o *OrganizationResponse) SetTagline(v string)`

SetTagline sets Tagline field to given value.

### HasTagline

`func (o *OrganizationResponse) HasTagline() bool`

HasTagline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


