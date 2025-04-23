# UserAuthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | API token for the authenticated user | [optional] [readonly] 
**TwoFactorRequired** | Pointer to **bool** | Flag indicating whether a 2FA code is required to complete authentication | [optional] [readonly] 
**TwoFactorToken** | Pointer to **string** | Token to use when providing 2FA code | [optional] [readonly] 

## Methods

### NewUserAuthToken

`func NewUserAuthToken() *UserAuthToken`

NewUserAuthToken instantiates a new UserAuthToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAuthTokenWithDefaults

`func NewUserAuthTokenWithDefaults() *UserAuthToken`

NewUserAuthTokenWithDefaults instantiates a new UserAuthToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *UserAuthToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UserAuthToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UserAuthToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UserAuthToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTwoFactorRequired

`func (o *UserAuthToken) GetTwoFactorRequired() bool`

GetTwoFactorRequired returns the TwoFactorRequired field if non-nil, zero value otherwise.

### GetTwoFactorRequiredOk

`func (o *UserAuthToken) GetTwoFactorRequiredOk() (*bool, bool)`

GetTwoFactorRequiredOk returns a tuple with the TwoFactorRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorRequired

`func (o *UserAuthToken) SetTwoFactorRequired(v bool)`

SetTwoFactorRequired sets TwoFactorRequired field to given value.

### HasTwoFactorRequired

`func (o *UserAuthToken) HasTwoFactorRequired() bool`

HasTwoFactorRequired returns a boolean if a field has been set.

### GetTwoFactorToken

`func (o *UserAuthToken) GetTwoFactorToken() string`

GetTwoFactorToken returns the TwoFactorToken field if non-nil, zero value otherwise.

### GetTwoFactorTokenOk

`func (o *UserAuthToken) GetTwoFactorTokenOk() (*string, bool)`

GetTwoFactorTokenOk returns a tuple with the TwoFactorToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorToken

`func (o *UserAuthToken) SetTwoFactorToken(v string)`

SetTwoFactorToken sets TwoFactorToken field to given value.

### HasTwoFactorToken

`func (o *UserAuthToken) HasTwoFactorToken() bool`

HasTwoFactorToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


