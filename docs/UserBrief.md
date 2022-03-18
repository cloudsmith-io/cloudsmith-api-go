# UserBrief

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authenticated** | Pointer to **bool** | If true then you&#39;re logged in as a user. | [optional] 
**Email** | Pointer to **string** | Your email address that we use to contact you. This is only visible to you. | [optional] 
**Name** | Pointer to **string** | The full name of the user (if any). | [optional] 
**ProfileUrl** | Pointer to **string** | The URL for the full profile of the user. | [optional] 
**SelfUrl** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**SlugPerm** | Pointer to **string** |  | [optional] 

## Methods

### NewUserBrief

`func NewUserBrief() *UserBrief`

NewUserBrief instantiates a new UserBrief object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBriefWithDefaults

`func NewUserBriefWithDefaults() *UserBrief`

NewUserBriefWithDefaults instantiates a new UserBrief object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticated

`func (o *UserBrief) GetAuthenticated() bool`

GetAuthenticated returns the Authenticated field if non-nil, zero value otherwise.

### GetAuthenticatedOk

`func (o *UserBrief) GetAuthenticatedOk() (*bool, bool)`

GetAuthenticatedOk returns a tuple with the Authenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticated

`func (o *UserBrief) SetAuthenticated(v bool)`

SetAuthenticated sets Authenticated field to given value.

### HasAuthenticated

`func (o *UserBrief) HasAuthenticated() bool`

HasAuthenticated returns a boolean if a field has been set.

### GetEmail

`func (o *UserBrief) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserBrief) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserBrief) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserBrief) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UserBrief) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserBrief) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserBrief) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserBrief) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfileUrl

`func (o *UserBrief) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *UserBrief) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *UserBrief) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *UserBrief) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### GetSelfUrl

`func (o *UserBrief) GetSelfUrl() string`

GetSelfUrl returns the SelfUrl field if non-nil, zero value otherwise.

### GetSelfUrlOk

`func (o *UserBrief) GetSelfUrlOk() (*string, bool)`

GetSelfUrlOk returns a tuple with the SelfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfUrl

`func (o *UserBrief) SetSelfUrl(v string)`

SetSelfUrl sets SelfUrl field to given value.

### HasSelfUrl

`func (o *UserBrief) HasSelfUrl() bool`

HasSelfUrl returns a boolean if a field has been set.

### GetSlug

`func (o *UserBrief) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *UserBrief) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *UserBrief) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *UserBrief) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSlugPerm

`func (o *UserBrief) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *UserBrief) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *UserBrief) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *UserBrief) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


