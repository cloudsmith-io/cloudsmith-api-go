# UserTokenCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email address to authenticate with | [optional] 
**Password** | Pointer to **string** | Password to authenticate with | [optional] 

## Methods

### NewUserTokenCreate

`func NewUserTokenCreate() *UserTokenCreate`

NewUserTokenCreate instantiates a new UserTokenCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTokenCreateWithDefaults

`func NewUserTokenCreateWithDefaults() *UserTokenCreate`

NewUserTokenCreateWithDefaults instantiates a new UserTokenCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserTokenCreate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserTokenCreate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserTokenCreate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserTokenCreate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *UserTokenCreate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserTokenCreate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserTokenCreate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserTokenCreate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


