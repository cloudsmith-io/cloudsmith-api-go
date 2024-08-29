# RepositoryEcdsaKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | If selected this is the active key for this repository. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Default** | Pointer to **bool** | If selected this is the default key for this repository. | [optional] [readonly] 
**Fingerprint** | Pointer to **string** | The long identifier used by ECDSA for this key. | [optional] [readonly] 
**FingerprintShort** | Pointer to **string** |  | [optional] [readonly] 
**PublicKey** | Pointer to **string** | The public key given to repository users. | [optional] [readonly] 
**SshFingerprint** | Pointer to **NullableString** | The SSH fingerprint used by ECDSA for this key. | [optional] [readonly] 

## Methods

### NewRepositoryEcdsaKey

`func NewRepositoryEcdsaKey() *RepositoryEcdsaKey`

NewRepositoryEcdsaKey instantiates a new RepositoryEcdsaKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryEcdsaKeyWithDefaults

`func NewRepositoryEcdsaKeyWithDefaults() *RepositoryEcdsaKey`

NewRepositoryEcdsaKeyWithDefaults instantiates a new RepositoryEcdsaKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *RepositoryEcdsaKey) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RepositoryEcdsaKey) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RepositoryEcdsaKey) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RepositoryEcdsaKey) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RepositoryEcdsaKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryEcdsaKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryEcdsaKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepositoryEcdsaKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *RepositoryEcdsaKey) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RepositoryEcdsaKey) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RepositoryEcdsaKey) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RepositoryEcdsaKey) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFingerprint

`func (o *RepositoryEcdsaKey) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *RepositoryEcdsaKey) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *RepositoryEcdsaKey) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *RepositoryEcdsaKey) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetFingerprintShort

`func (o *RepositoryEcdsaKey) GetFingerprintShort() string`

GetFingerprintShort returns the FingerprintShort field if non-nil, zero value otherwise.

### GetFingerprintShortOk

`func (o *RepositoryEcdsaKey) GetFingerprintShortOk() (*string, bool)`

GetFingerprintShortOk returns a tuple with the FingerprintShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintShort

`func (o *RepositoryEcdsaKey) SetFingerprintShort(v string)`

SetFingerprintShort sets FingerprintShort field to given value.

### HasFingerprintShort

`func (o *RepositoryEcdsaKey) HasFingerprintShort() bool`

HasFingerprintShort returns a boolean if a field has been set.

### GetPublicKey

`func (o *RepositoryEcdsaKey) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *RepositoryEcdsaKey) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *RepositoryEcdsaKey) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *RepositoryEcdsaKey) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetSshFingerprint

`func (o *RepositoryEcdsaKey) GetSshFingerprint() string`

GetSshFingerprint returns the SshFingerprint field if non-nil, zero value otherwise.

### GetSshFingerprintOk

`func (o *RepositoryEcdsaKey) GetSshFingerprintOk() (*string, bool)`

GetSshFingerprintOk returns a tuple with the SshFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshFingerprint

`func (o *RepositoryEcdsaKey) SetSshFingerprint(v string)`

SetSshFingerprint sets SshFingerprint field to given value.

### HasSshFingerprint

`func (o *RepositoryEcdsaKey) HasSshFingerprint() bool`

HasSshFingerprint returns a boolean if a field has been set.

### SetSshFingerprintNil

`func (o *RepositoryEcdsaKey) SetSshFingerprintNil(b bool)`

 SetSshFingerprintNil sets the value for SshFingerprint to be an explicit nil

### UnsetSshFingerprint
`func (o *RepositoryEcdsaKey) UnsetSshFingerprint()`

UnsetSshFingerprint ensures that no value is present for SshFingerprint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


