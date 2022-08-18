# ReposRsaCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RsaPassphrase** | Pointer to **string** | The RSA passphrase used for signing. | [optional] 
**RsaPrivateKey** | **string** | The RSA private key. | 

## Methods

### NewReposRsaCreate

`func NewReposRsaCreate(rsaPrivateKey string, ) *ReposRsaCreate`

NewReposRsaCreate instantiates a new ReposRsaCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposRsaCreateWithDefaults

`func NewReposRsaCreateWithDefaults() *ReposRsaCreate`

NewReposRsaCreateWithDefaults instantiates a new ReposRsaCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRsaPassphrase

`func (o *ReposRsaCreate) GetRsaPassphrase() string`

GetRsaPassphrase returns the RsaPassphrase field if non-nil, zero value otherwise.

### GetRsaPassphraseOk

`func (o *ReposRsaCreate) GetRsaPassphraseOk() (*string, bool)`

GetRsaPassphraseOk returns a tuple with the RsaPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaPassphrase

`func (o *ReposRsaCreate) SetRsaPassphrase(v string)`

SetRsaPassphrase sets RsaPassphrase field to given value.

### HasRsaPassphrase

`func (o *ReposRsaCreate) HasRsaPassphrase() bool`

HasRsaPassphrase returns a boolean if a field has been set.

### GetRsaPrivateKey

`func (o *ReposRsaCreate) GetRsaPrivateKey() string`

GetRsaPrivateKey returns the RsaPrivateKey field if non-nil, zero value otherwise.

### GetRsaPrivateKeyOk

`func (o *ReposRsaCreate) GetRsaPrivateKeyOk() (*string, bool)`

GetRsaPrivateKeyOk returns a tuple with the RsaPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaPrivateKey

`func (o *ReposRsaCreate) SetRsaPrivateKey(v string)`

SetRsaPrivateKey sets RsaPrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


