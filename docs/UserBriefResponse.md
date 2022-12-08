# UserBriefResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authenticated** | Pointer to **bool** | If true then you&#39;re logged in as a user. | [optional] [readonly] 
**Email** | Pointer to **NullableString** | Your email address that we use to contact you. This is only visible to you. | [optional] 
**Name** | Pointer to **NullableString** | The full name of the user (if any). | [optional] [readonly] 
**ProfileUrl** | Pointer to **NullableString** | The URL for the full profile of the user. | [optional] [readonly] 
**SelfUrl** | Pointer to **string** |  | [optional] [readonly] 
**Slug** | Pointer to **NullableString** |  | [optional] [readonly] 
**SlugPerm** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewUserBriefResponse

`func NewUserBriefResponse() *UserBriefResponse`

NewUserBriefResponse instantiates a new UserBriefResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBriefResponseWithDefaults

`func NewUserBriefResponseWithDefaults() *UserBriefResponse`

NewUserBriefResponseWithDefaults instantiates a new UserBriefResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticated

`func (o *UserBriefResponse) GetAuthenticated() bool`

GetAuthenticated returns the Authenticated field if non-nil, zero value otherwise.

### GetAuthenticatedOk

`func (o *UserBriefResponse) GetAuthenticatedOk() (*bool, bool)`

GetAuthenticatedOk returns a tuple with the Authenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticated

`func (o *UserBriefResponse) SetAuthenticated(v bool)`

SetAuthenticated sets Authenticated field to given value.

### HasAuthenticated

`func (o *UserBriefResponse) HasAuthenticated() bool`

HasAuthenticated returns a boolean if a field has been set.

### GetEmail

`func (o *UserBriefResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserBriefResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserBriefResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserBriefResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserBriefResponse) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserBriefResponse) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetName

`func (o *UserBriefResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserBriefResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserBriefResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserBriefResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UserBriefResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserBriefResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProfileUrl

`func (o *UserBriefResponse) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *UserBriefResponse) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *UserBriefResponse) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *UserBriefResponse) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### SetProfileUrlNil

`func (o *UserBriefResponse) SetProfileUrlNil(b bool)`

 SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil

### UnsetProfileUrl
`func (o *UserBriefResponse) UnsetProfileUrl()`

UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
### GetSelfUrl

`func (o *UserBriefResponse) GetSelfUrl() string`

GetSelfUrl returns the SelfUrl field if non-nil, zero value otherwise.

### GetSelfUrlOk

`func (o *UserBriefResponse) GetSelfUrlOk() (*string, bool)`

GetSelfUrlOk returns a tuple with the SelfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfUrl

`func (o *UserBriefResponse) SetSelfUrl(v string)`

SetSelfUrl sets SelfUrl field to given value.

### HasSelfUrl

`func (o *UserBriefResponse) HasSelfUrl() bool`

HasSelfUrl returns a boolean if a field has been set.

### GetSlug

`func (o *UserBriefResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *UserBriefResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *UserBriefResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *UserBriefResponse) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *UserBriefResponse) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *UserBriefResponse) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetSlugPerm

`func (o *UserBriefResponse) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *UserBriefResponse) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *UserBriefResponse) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *UserBriefResponse) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### SetSlugPermNil

`func (o *UserBriefResponse) SetSlugPermNil(b bool)`

 SetSlugPermNil sets the value for SlugPerm to be an explicit nil

### UnsetSlugPerm
`func (o *UserBriefResponse) UnsetSlugPerm()`

UnsetSlugPerm ensures that no value is present for SlugPerm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


