# RepositoryGpgKeyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpgPassphrase** | Pointer to **string** | The GPG passphrase used for signing. | [optional] 
**GpgPrivateKey** | **string** | The GPG private key. | 

## Methods

### NewRepositoryGpgKeyCreate

`func NewRepositoryGpgKeyCreate(gpgPrivateKey string, ) *RepositoryGpgKeyCreate`

NewRepositoryGpgKeyCreate instantiates a new RepositoryGpgKeyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryGpgKeyCreateWithDefaults

`func NewRepositoryGpgKeyCreateWithDefaults() *RepositoryGpgKeyCreate`

NewRepositoryGpgKeyCreateWithDefaults instantiates a new RepositoryGpgKeyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpgPassphrase

`func (o *RepositoryGpgKeyCreate) GetGpgPassphrase() string`

GetGpgPassphrase returns the GpgPassphrase field if non-nil, zero value otherwise.

### GetGpgPassphraseOk

`func (o *RepositoryGpgKeyCreate) GetGpgPassphraseOk() (*string, bool)`

GetGpgPassphraseOk returns a tuple with the GpgPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgPassphrase

`func (o *RepositoryGpgKeyCreate) SetGpgPassphrase(v string)`

SetGpgPassphrase sets GpgPassphrase field to given value.

### HasGpgPassphrase

`func (o *RepositoryGpgKeyCreate) HasGpgPassphrase() bool`

HasGpgPassphrase returns a boolean if a field has been set.

### GetGpgPrivateKey

`func (o *RepositoryGpgKeyCreate) GetGpgPrivateKey() string`

GetGpgPrivateKey returns the GpgPrivateKey field if non-nil, zero value otherwise.

### GetGpgPrivateKeyOk

`func (o *RepositoryGpgKeyCreate) GetGpgPrivateKeyOk() (*string, bool)`

GetGpgPrivateKeyOk returns a tuple with the GpgPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgPrivateKey

`func (o *RepositoryGpgKeyCreate) SetGpgPrivateKey(v string)`

SetGpgPrivateKey sets GpgPrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


