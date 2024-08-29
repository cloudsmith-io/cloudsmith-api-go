# RepositoryEcdsaKeyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EcdsaPassphrase** | Pointer to **string** | The ECDSA passphrase used for signing. | [optional] 
**EcdsaPrivateKey** | **string** | The ECDSA private key. | 

## Methods

### NewRepositoryEcdsaKeyCreate

`func NewRepositoryEcdsaKeyCreate(ecdsaPrivateKey string, ) *RepositoryEcdsaKeyCreate`

NewRepositoryEcdsaKeyCreate instantiates a new RepositoryEcdsaKeyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryEcdsaKeyCreateWithDefaults

`func NewRepositoryEcdsaKeyCreateWithDefaults() *RepositoryEcdsaKeyCreate`

NewRepositoryEcdsaKeyCreateWithDefaults instantiates a new RepositoryEcdsaKeyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEcdsaPassphrase

`func (o *RepositoryEcdsaKeyCreate) GetEcdsaPassphrase() string`

GetEcdsaPassphrase returns the EcdsaPassphrase field if non-nil, zero value otherwise.

### GetEcdsaPassphraseOk

`func (o *RepositoryEcdsaKeyCreate) GetEcdsaPassphraseOk() (*string, bool)`

GetEcdsaPassphraseOk returns a tuple with the EcdsaPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdsaPassphrase

`func (o *RepositoryEcdsaKeyCreate) SetEcdsaPassphrase(v string)`

SetEcdsaPassphrase sets EcdsaPassphrase field to given value.

### HasEcdsaPassphrase

`func (o *RepositoryEcdsaKeyCreate) HasEcdsaPassphrase() bool`

HasEcdsaPassphrase returns a boolean if a field has been set.

### GetEcdsaPrivateKey

`func (o *RepositoryEcdsaKeyCreate) GetEcdsaPrivateKey() string`

GetEcdsaPrivateKey returns the EcdsaPrivateKey field if non-nil, zero value otherwise.

### GetEcdsaPrivateKeyOk

`func (o *RepositoryEcdsaKeyCreate) GetEcdsaPrivateKeyOk() (*string, bool)`

GetEcdsaPrivateKeyOk returns a tuple with the EcdsaPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdsaPrivateKey

`func (o *RepositoryEcdsaKeyCreate) SetEcdsaPrivateKey(v string)`

SetEcdsaPrivateKey sets EcdsaPrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


