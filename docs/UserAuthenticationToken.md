# UserAuthenticationToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The time at which the API key was created. | [optional] [readonly] 
**Key** | Pointer to **string** | The unique API key used for authentication. This will be obfuscated on read-only HTTP methods. | [optional] [readonly] 
**SlugPerm** | Pointer to **string** | The slug_perm for token. | [optional] [readonly] 

## Methods

### NewUserAuthenticationToken

`func NewUserAuthenticationToken() *UserAuthenticationToken`

NewUserAuthenticationToken instantiates a new UserAuthenticationToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAuthenticationTokenWithDefaults

`func NewUserAuthenticationTokenWithDefaults() *UserAuthenticationToken`

NewUserAuthenticationTokenWithDefaults instantiates a new UserAuthenticationToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *UserAuthenticationToken) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserAuthenticationToken) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserAuthenticationToken) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UserAuthenticationToken) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetKey

`func (o *UserAuthenticationToken) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UserAuthenticationToken) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UserAuthenticationToken) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UserAuthenticationToken) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSlugPerm

`func (o *UserAuthenticationToken) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *UserAuthenticationToken) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *UserAuthenticationToken) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *UserAuthenticationToken) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


