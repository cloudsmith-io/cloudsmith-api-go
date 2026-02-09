# BroadcastToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Token** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewBroadcastToken

`func NewBroadcastToken() *BroadcastToken`

NewBroadcastToken instantiates a new BroadcastToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastTokenWithDefaults

`func NewBroadcastTokenWithDefaults() *BroadcastToken`

NewBroadcastTokenWithDefaults instantiates a new BroadcastToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *BroadcastToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *BroadcastToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *BroadcastToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *BroadcastToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetToken

`func (o *BroadcastToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BroadcastToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BroadcastToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BroadcastToken) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


