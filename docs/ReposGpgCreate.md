# ReposGpgCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpgPassphrase** | Pointer to **string** | The GPG passphrase used for signing. | [optional] 
**GpgPrivateKey** | **string** | The GPG private key. | 

## Methods

### NewReposGpgCreate

`func NewReposGpgCreate(gpgPrivateKey string, ) *ReposGpgCreate`

NewReposGpgCreate instantiates a new ReposGpgCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposGpgCreateWithDefaults

`func NewReposGpgCreateWithDefaults() *ReposGpgCreate`

NewReposGpgCreateWithDefaults instantiates a new ReposGpgCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpgPassphrase

`func (o *ReposGpgCreate) GetGpgPassphrase() string`

GetGpgPassphrase returns the GpgPassphrase field if non-nil, zero value otherwise.

### GetGpgPassphraseOk

`func (o *ReposGpgCreate) GetGpgPassphraseOk() (*string, bool)`

GetGpgPassphraseOk returns a tuple with the GpgPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgPassphrase

`func (o *ReposGpgCreate) SetGpgPassphrase(v string)`

SetGpgPassphrase sets GpgPassphrase field to given value.

### HasGpgPassphrase

`func (o *ReposGpgCreate) HasGpgPassphrase() bool`

HasGpgPassphrase returns a boolean if a field has been set.

### GetGpgPrivateKey

`func (o *ReposGpgCreate) GetGpgPrivateKey() string`

GetGpgPrivateKey returns the GpgPrivateKey field if non-nil, zero value otherwise.

### GetGpgPrivateKeyOk

`func (o *ReposGpgCreate) GetGpgPrivateKeyOk() (*string, bool)`

GetGpgPrivateKeyOk returns a tuple with the GpgPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgPrivateKey

`func (o *ReposGpgCreate) SetGpgPrivateKey(v string)`

SetGpgPrivateKey sets GpgPrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


