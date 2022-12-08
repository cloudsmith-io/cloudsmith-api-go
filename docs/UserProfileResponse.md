# UserProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **NullableString** |  | [optional] 
**FirstName** | **string** |  | 
**JobTitle** | Pointer to **NullableString** |  | [optional] 
**JoinedAt** | Pointer to **time.Time** |  | [optional] 
**LastName** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Slug** | Pointer to **string** |  | [optional] [readonly] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**Tagline** | Pointer to **NullableString** | Your tagline is a sentence about you. Make it funny. Make it professional. Either way, it&#39;s public and it represents who you are. | [optional] 
**Url** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewUserProfileResponse

`func NewUserProfileResponse(firstName string, lastName string, ) *UserProfileResponse`

NewUserProfileResponse instantiates a new UserProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileResponseWithDefaults

`func NewUserProfileResponseWithDefaults() *UserProfileResponse`

NewUserProfileResponseWithDefaults instantiates a new UserProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *UserProfileResponse) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UserProfileResponse) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UserProfileResponse) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UserProfileResponse) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *UserProfileResponse) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *UserProfileResponse) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetFirstName

`func (o *UserProfileResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserProfileResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserProfileResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetJobTitle

`func (o *UserProfileResponse) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *UserProfileResponse) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *UserProfileResponse) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *UserProfileResponse) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *UserProfileResponse) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *UserProfileResponse) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetJoinedAt

`func (o *UserProfileResponse) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *UserProfileResponse) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *UserProfileResponse) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *UserProfileResponse) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetLastName

`func (o *UserProfileResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserProfileResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserProfileResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetName

`func (o *UserProfileResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserProfileResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserProfileResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserProfileResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *UserProfileResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *UserProfileResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *UserProfileResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *UserProfileResponse) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSlugPerm

`func (o *UserProfileResponse) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *UserProfileResponse) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *UserProfileResponse) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *UserProfileResponse) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetTagline

`func (o *UserProfileResponse) GetTagline() string`

GetTagline returns the Tagline field if non-nil, zero value otherwise.

### GetTaglineOk

`func (o *UserProfileResponse) GetTaglineOk() (*string, bool)`

GetTaglineOk returns a tuple with the Tagline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagline

`func (o *UserProfileResponse) SetTagline(v string)`

SetTagline sets Tagline field to given value.

### HasTagline

`func (o *UserProfileResponse) HasTagline() bool`

HasTagline returns a boolean if a field has been set.

### SetTaglineNil

`func (o *UserProfileResponse) SetTaglineNil(b bool)`

 SetTaglineNil sets the value for Tagline to be an explicit nil

### UnsetTagline
`func (o *UserProfileResponse) UnsetTagline()`

UnsetTagline ensures that no value is present for Tagline, not even an explicit nil
### GetUrl

`func (o *UserProfileResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserProfileResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserProfileResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UserProfileResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


