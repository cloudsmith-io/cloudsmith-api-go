# BroadcastTokenInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementToken** | **string** | Repository entitlement token used to authorize the creation of a broadcast token | 
**ExpiresIn** | Pointer to **int64** | Token expiry time in seconds (optional, defaults to 3600 seconds) | [optional] 

## Methods

### NewBroadcastTokenInput

`func NewBroadcastTokenInput(entitlementToken string, ) *BroadcastTokenInput`

NewBroadcastTokenInput instantiates a new BroadcastTokenInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastTokenInputWithDefaults

`func NewBroadcastTokenInputWithDefaults() *BroadcastTokenInput`

NewBroadcastTokenInputWithDefaults instantiates a new BroadcastTokenInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementToken

`func (o *BroadcastTokenInput) GetEntitlementToken() string`

GetEntitlementToken returns the EntitlementToken field if non-nil, zero value otherwise.

### GetEntitlementTokenOk

`func (o *BroadcastTokenInput) GetEntitlementTokenOk() (*string, bool)`

GetEntitlementTokenOk returns a tuple with the EntitlementToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementToken

`func (o *BroadcastTokenInput) SetEntitlementToken(v string)`

SetEntitlementToken sets EntitlementToken field to given value.


### GetExpiresIn

`func (o *BroadcastTokenInput) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *BroadcastTokenInput) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *BroadcastTokenInput) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *BroadcastTokenInput) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


