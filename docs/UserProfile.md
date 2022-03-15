# UserProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **string** |  | [optional] 
**FirstName** | **string** |  | 
**JobTitle** | Pointer to **string** |  | [optional] 
**JoinedAt** | Pointer to **string** |  | [optional] 
**LastName** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**SlugPerm** | Pointer to **string** |  | [optional] 
**Tagline** | Pointer to **string** | Your tagline is a sentence about you. Make it funny. Make it professional. Either way, it&#39;s public and it represents who you are. | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewUserProfile

`func NewUserProfile(firstName string, lastName string, ) *UserProfile`

NewUserProfile instantiates a new UserProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileWithDefaults

`func NewUserProfileWithDefaults() *UserProfile`

NewUserProfileWithDefaults instantiates a new UserProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *UserProfile) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UserProfile) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UserProfile) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UserProfile) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetFirstName

`func (o *UserProfile) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserProfile) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserProfile) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetJobTitle

`func (o *UserProfile) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *UserProfile) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *UserProfile) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *UserProfile) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetJoinedAt

`func (o *UserProfile) GetJoinedAt() string`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *UserProfile) GetJoinedAtOk() (*string, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *UserProfile) SetJoinedAt(v string)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *UserProfile) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetLastName

`func (o *UserProfile) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserProfile) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserProfile) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetName

`func (o *UserProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *UserProfile) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *UserProfile) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *UserProfile) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *UserProfile) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSlugPerm

`func (o *UserProfile) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *UserProfile) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *UserProfile) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *UserProfile) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetTagline

`func (o *UserProfile) GetTagline() string`

GetTagline returns the Tagline field if non-nil, zero value otherwise.

### GetTaglineOk

`func (o *UserProfile) GetTaglineOk() (*string, bool)`

GetTaglineOk returns a tuple with the Tagline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagline

`func (o *UserProfile) SetTagline(v string)`

SetTagline sets Tagline field to given value.

### HasTagline

`func (o *UserProfile) HasTagline() bool`

HasTagline returns a boolean if a field has been set.

### GetUrl

`func (o *UserProfile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserProfile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserProfile) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UserProfile) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


