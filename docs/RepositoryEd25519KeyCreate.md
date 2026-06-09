# RepositoryEd25519KeyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ed25519Passphrase** | Pointer to **string** | The Ed25519 passphrase used for signing. | [optional] 
**Ed25519PrivateKey** | **string** | The Ed25519 private key. | 

## Methods

### NewRepositoryEd25519KeyCreate

`func NewRepositoryEd25519KeyCreate(ed25519PrivateKey string, ) *RepositoryEd25519KeyCreate`

NewRepositoryEd25519KeyCreate instantiates a new RepositoryEd25519KeyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryEd25519KeyCreateWithDefaults

`func NewRepositoryEd25519KeyCreateWithDefaults() *RepositoryEd25519KeyCreate`

NewRepositoryEd25519KeyCreateWithDefaults instantiates a new RepositoryEd25519KeyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEd25519Passphrase

`func (o *RepositoryEd25519KeyCreate) GetEd25519Passphrase() string`

GetEd25519Passphrase returns the Ed25519Passphrase field if non-nil, zero value otherwise.

### GetEd25519PassphraseOk

`func (o *RepositoryEd25519KeyCreate) GetEd25519PassphraseOk() (*string, bool)`

GetEd25519PassphraseOk returns a tuple with the Ed25519Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEd25519Passphrase

`func (o *RepositoryEd25519KeyCreate) SetEd25519Passphrase(v string)`

SetEd25519Passphrase sets Ed25519Passphrase field to given value.

### HasEd25519Passphrase

`func (o *RepositoryEd25519KeyCreate) HasEd25519Passphrase() bool`

HasEd25519Passphrase returns a boolean if a field has been set.

### GetEd25519PrivateKey

`func (o *RepositoryEd25519KeyCreate) GetEd25519PrivateKey() string`

GetEd25519PrivateKey returns the Ed25519PrivateKey field if non-nil, zero value otherwise.

### GetEd25519PrivateKeyOk

`func (o *RepositoryEd25519KeyCreate) GetEd25519PrivateKeyOk() (*string, bool)`

GetEd25519PrivateKeyOk returns a tuple with the Ed25519PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEd25519PrivateKey

`func (o *RepositoryEd25519KeyCreate) SetEd25519PrivateKey(v string)`

SetEd25519PrivateKey sets Ed25519PrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


