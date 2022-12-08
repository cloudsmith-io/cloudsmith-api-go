# RepositoryGpgKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | If selected this is the active key for this repository. | [optional] [readonly] 
**Comment** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Default** | Pointer to **bool** | If selected this is the default key for this repository. | [optional] [readonly] 
**Fingerprint** | Pointer to **string** | The long identifier used by GPG for this key. | [optional] [readonly] 
**FingerprintShort** | Pointer to **string** |  | [optional] [readonly] 
**PublicKey** | Pointer to **string** | The public key given to repository users. | [optional] [readonly] 

## Methods

### NewRepositoryGpgKey

`func NewRepositoryGpgKey(comment string, ) *RepositoryGpgKey`

NewRepositoryGpgKey instantiates a new RepositoryGpgKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryGpgKeyWithDefaults

`func NewRepositoryGpgKeyWithDefaults() *RepositoryGpgKey`

NewRepositoryGpgKeyWithDefaults instantiates a new RepositoryGpgKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *RepositoryGpgKey) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RepositoryGpgKey) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RepositoryGpgKey) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RepositoryGpgKey) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetComment

`func (o *RepositoryGpgKey) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RepositoryGpgKey) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RepositoryGpgKey) SetComment(v string)`

SetComment sets Comment field to given value.


### GetCreatedAt

`func (o *RepositoryGpgKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryGpgKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryGpgKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepositoryGpgKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *RepositoryGpgKey) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RepositoryGpgKey) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RepositoryGpgKey) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RepositoryGpgKey) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFingerprint

`func (o *RepositoryGpgKey) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *RepositoryGpgKey) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *RepositoryGpgKey) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *RepositoryGpgKey) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetFingerprintShort

`func (o *RepositoryGpgKey) GetFingerprintShort() string`

GetFingerprintShort returns the FingerprintShort field if non-nil, zero value otherwise.

### GetFingerprintShortOk

`func (o *RepositoryGpgKey) GetFingerprintShortOk() (*string, bool)`

GetFingerprintShortOk returns a tuple with the FingerprintShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintShort

`func (o *RepositoryGpgKey) SetFingerprintShort(v string)`

SetFingerprintShort sets FingerprintShort field to given value.

### HasFingerprintShort

`func (o *RepositoryGpgKey) HasFingerprintShort() bool`

HasFingerprintShort returns a boolean if a field has been set.

### GetPublicKey

`func (o *RepositoryGpgKey) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *RepositoryGpgKey) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *RepositoryGpgKey) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *RepositoryGpgKey) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


