# RepositoryRsaKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | If selected this is the active key for this repository. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Default** | Pointer to **bool** | If selected this is the default key for this repository. | [optional] [readonly] 
**Fingerprint** | Pointer to **string** | The long identifier used by RSA for this key. | [optional] [readonly] 
**FingerprintShort** | Pointer to **string** |  | [optional] [readonly] 
**PublicKey** | Pointer to **string** | The public key given to repository users. | [optional] [readonly] 
**SshFingerprint** | Pointer to **NullableString** | The SSH fingerprint used by RSA for this key. | [optional] [readonly] 

## Methods

### NewRepositoryRsaKey

`func NewRepositoryRsaKey() *RepositoryRsaKey`

NewRepositoryRsaKey instantiates a new RepositoryRsaKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryRsaKeyWithDefaults

`func NewRepositoryRsaKeyWithDefaults() *RepositoryRsaKey`

NewRepositoryRsaKeyWithDefaults instantiates a new RepositoryRsaKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *RepositoryRsaKey) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RepositoryRsaKey) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RepositoryRsaKey) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RepositoryRsaKey) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RepositoryRsaKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryRsaKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryRsaKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepositoryRsaKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *RepositoryRsaKey) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RepositoryRsaKey) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RepositoryRsaKey) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RepositoryRsaKey) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFingerprint

`func (o *RepositoryRsaKey) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *RepositoryRsaKey) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *RepositoryRsaKey) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *RepositoryRsaKey) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetFingerprintShort

`func (o *RepositoryRsaKey) GetFingerprintShort() string`

GetFingerprintShort returns the FingerprintShort field if non-nil, zero value otherwise.

### GetFingerprintShortOk

`func (o *RepositoryRsaKey) GetFingerprintShortOk() (*string, bool)`

GetFingerprintShortOk returns a tuple with the FingerprintShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintShort

`func (o *RepositoryRsaKey) SetFingerprintShort(v string)`

SetFingerprintShort sets FingerprintShort field to given value.

### HasFingerprintShort

`func (o *RepositoryRsaKey) HasFingerprintShort() bool`

HasFingerprintShort returns a boolean if a field has been set.

### GetPublicKey

`func (o *RepositoryRsaKey) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *RepositoryRsaKey) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *RepositoryRsaKey) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *RepositoryRsaKey) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetSshFingerprint

`func (o *RepositoryRsaKey) GetSshFingerprint() string`

GetSshFingerprint returns the SshFingerprint field if non-nil, zero value otherwise.

### GetSshFingerprintOk

`func (o *RepositoryRsaKey) GetSshFingerprintOk() (*string, bool)`

GetSshFingerprintOk returns a tuple with the SshFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshFingerprint

`func (o *RepositoryRsaKey) SetSshFingerprint(v string)`

SetSshFingerprint sets SshFingerprint field to given value.

### HasSshFingerprint

`func (o *RepositoryRsaKey) HasSshFingerprint() bool`

HasSshFingerprint returns a boolean if a field has been set.

### SetSshFingerprintNil

`func (o *RepositoryRsaKey) SetSshFingerprintNil(b bool)`

 SetSshFingerprintNil sets the value for SshFingerprint to be an explicit nil

### UnsetSshFingerprint
`func (o *RepositoryRsaKey) UnsetSshFingerprint()`

UnsetSshFingerprint ensures that no value is present for SshFingerprint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


