# RepositoryRsaKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | If selected this is the active key for this repository. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Default** | Pointer to **bool** | If selected this is the default key for this repository. | [optional] [readonly] 
**Fingerprint** | Pointer to **string** | The long identifier used by RSA for this key. | [optional] [readonly] 
**FingerprintShort** | Pointer to **string** |  | [optional] [readonly] 
**PublicKey** | Pointer to **string** | The public key given to repository users. | [optional] [readonly] 

## Methods

### NewRepositoryRsaKeyResponse

`func NewRepositoryRsaKeyResponse() *RepositoryRsaKeyResponse`

NewRepositoryRsaKeyResponse instantiates a new RepositoryRsaKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryRsaKeyResponseWithDefaults

`func NewRepositoryRsaKeyResponseWithDefaults() *RepositoryRsaKeyResponse`

NewRepositoryRsaKeyResponseWithDefaults instantiates a new RepositoryRsaKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *RepositoryRsaKeyResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RepositoryRsaKeyResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RepositoryRsaKeyResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RepositoryRsaKeyResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RepositoryRsaKeyResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryRsaKeyResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryRsaKeyResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepositoryRsaKeyResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *RepositoryRsaKeyResponse) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RepositoryRsaKeyResponse) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RepositoryRsaKeyResponse) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RepositoryRsaKeyResponse) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFingerprint

`func (o *RepositoryRsaKeyResponse) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *RepositoryRsaKeyResponse) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *RepositoryRsaKeyResponse) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *RepositoryRsaKeyResponse) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetFingerprintShort

`func (o *RepositoryRsaKeyResponse) GetFingerprintShort() string`

GetFingerprintShort returns the FingerprintShort field if non-nil, zero value otherwise.

### GetFingerprintShortOk

`func (o *RepositoryRsaKeyResponse) GetFingerprintShortOk() (*string, bool)`

GetFingerprintShortOk returns a tuple with the FingerprintShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintShort

`func (o *RepositoryRsaKeyResponse) SetFingerprintShort(v string)`

SetFingerprintShort sets FingerprintShort field to given value.

### HasFingerprintShort

`func (o *RepositoryRsaKeyResponse) HasFingerprintShort() bool`

HasFingerprintShort returns a boolean if a field has been set.

### GetPublicKey

`func (o *RepositoryRsaKeyResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *RepositoryRsaKeyResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *RepositoryRsaKeyResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *RepositoryRsaKeyResponse) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


