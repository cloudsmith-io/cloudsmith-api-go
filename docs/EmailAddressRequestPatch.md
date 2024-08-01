# EmailAddressRequestPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Primary** | Pointer to **bool** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 

## Methods

### NewEmailAddressRequestPatch

`func NewEmailAddressRequestPatch() *EmailAddressRequestPatch`

NewEmailAddressRequestPatch instantiates a new EmailAddressRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAddressRequestPatchWithDefaults

`func NewEmailAddressRequestPatchWithDefaults() *EmailAddressRequestPatch`

NewEmailAddressRequestPatchWithDefaults instantiates a new EmailAddressRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EmailAddressRequestPatch) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailAddressRequestPatch) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailAddressRequestPatch) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmailAddressRequestPatch) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPrimary

`func (o *EmailAddressRequestPatch) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *EmailAddressRequestPatch) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *EmailAddressRequestPatch) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *EmailAddressRequestPatch) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetVerified

`func (o *EmailAddressRequestPatch) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *EmailAddressRequestPatch) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *EmailAddressRequestPatch) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *EmailAddressRequestPatch) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


