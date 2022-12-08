# RepositoryGpgKeyResponse

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

### NewRepositoryGpgKeyResponse

`func NewRepositoryGpgKeyResponse(comment string, ) *RepositoryGpgKeyResponse`

NewRepositoryGpgKeyResponse instantiates a new RepositoryGpgKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryGpgKeyResponseWithDefaults

`func NewRepositoryGpgKeyResponseWithDefaults() *RepositoryGpgKeyResponse`

NewRepositoryGpgKeyResponseWithDefaults instantiates a new RepositoryGpgKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *RepositoryGpgKeyResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RepositoryGpgKeyResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RepositoryGpgKeyResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RepositoryGpgKeyResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetComment

`func (o *RepositoryGpgKeyResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RepositoryGpgKeyResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RepositoryGpgKeyResponse) SetComment(v string)`

SetComment sets Comment field to given value.


### GetCreatedAt

`func (o *RepositoryGpgKeyResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryGpgKeyResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryGpgKeyResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepositoryGpgKeyResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *RepositoryGpgKeyResponse) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RepositoryGpgKeyResponse) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RepositoryGpgKeyResponse) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RepositoryGpgKeyResponse) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFingerprint

`func (o *RepositoryGpgKeyResponse) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *RepositoryGpgKeyResponse) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *RepositoryGpgKeyResponse) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *RepositoryGpgKeyResponse) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetFingerprintShort

`func (o *RepositoryGpgKeyResponse) GetFingerprintShort() string`

GetFingerprintShort returns the FingerprintShort field if non-nil, zero value otherwise.

### GetFingerprintShortOk

`func (o *RepositoryGpgKeyResponse) GetFingerprintShortOk() (*string, bool)`

GetFingerprintShortOk returns a tuple with the FingerprintShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintShort

`func (o *RepositoryGpgKeyResponse) SetFingerprintShort(v string)`

SetFingerprintShort sets FingerprintShort field to given value.

### HasFingerprintShort

`func (o *RepositoryGpgKeyResponse) HasFingerprintShort() bool`

HasFingerprintShort returns a boolean if a field has been set.

### GetPublicKey

`func (o *RepositoryGpgKeyResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *RepositoryGpgKeyResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *RepositoryGpgKeyResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *RepositoryGpgKeyResponse) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


