# EmailAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Identifier** | Pointer to **string** |  | [optional] [readonly] 
**Primary** | Pointer to **bool** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 

## Methods

### NewEmailAddress

`func NewEmailAddress(email string, ) *EmailAddress`

NewEmailAddress instantiates a new EmailAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAddressWithDefaults

`func NewEmailAddressWithDefaults() *EmailAddress`

NewEmailAddressWithDefaults instantiates a new EmailAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EmailAddress) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailAddress) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailAddress) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetIdentifier

`func (o *EmailAddress) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *EmailAddress) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *EmailAddress) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *EmailAddress) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetPrimary

`func (o *EmailAddress) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *EmailAddress) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *EmailAddress) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *EmailAddress) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetVerified

`func (o *EmailAddress) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *EmailAddress) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *EmailAddress) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *EmailAddress) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


