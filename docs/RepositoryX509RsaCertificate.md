# RepositoryX509RsaCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | If selected this is the active key for this repository. | [optional] [readonly] 
**Certificate** | Pointer to **NullableString** | The issued certificate. | [optional] [readonly] 
**CertificateChain** | Pointer to **NullableString** | Base64 encoded CA certificate chain. | [optional] [readonly] 
**CertificateChainFingerprint** | Pointer to **string** |  | [optional] [readonly] 
**CertificateChainFingerprintShort** | Pointer to **string** |  | [optional] [readonly] 
**CertificateFingerprint** | Pointer to **NullableString** | The SHA-256 long identifier used | [optional] [readonly] 
**CertificateFingerprintShort** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Default** | Pointer to **bool** | If selected this is the default key for this repository. | [optional] [readonly] 
**IssuingStatus** | Pointer to **string** |  | [optional] [default to "Certificate is pending to be issued"]

## Methods

### NewRepositoryX509RsaCertificate

`func NewRepositoryX509RsaCertificate() *RepositoryX509RsaCertificate`

NewRepositoryX509RsaCertificate instantiates a new RepositoryX509RsaCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryX509RsaCertificateWithDefaults

`func NewRepositoryX509RsaCertificateWithDefaults() *RepositoryX509RsaCertificate`

NewRepositoryX509RsaCertificateWithDefaults instantiates a new RepositoryX509RsaCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *RepositoryX509RsaCertificate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RepositoryX509RsaCertificate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RepositoryX509RsaCertificate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RepositoryX509RsaCertificate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCertificate

`func (o *RepositoryX509RsaCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *RepositoryX509RsaCertificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *RepositoryX509RsaCertificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *RepositoryX509RsaCertificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *RepositoryX509RsaCertificate) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *RepositoryX509RsaCertificate) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCertificateChain

`func (o *RepositoryX509RsaCertificate) GetCertificateChain() string`

GetCertificateChain returns the CertificateChain field if non-nil, zero value otherwise.

### GetCertificateChainOk

`func (o *RepositoryX509RsaCertificate) GetCertificateChainOk() (*string, bool)`

GetCertificateChainOk returns a tuple with the CertificateChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChain

`func (o *RepositoryX509RsaCertificate) SetCertificateChain(v string)`

SetCertificateChain sets CertificateChain field to given value.

### HasCertificateChain

`func (o *RepositoryX509RsaCertificate) HasCertificateChain() bool`

HasCertificateChain returns a boolean if a field has been set.

### SetCertificateChainNil

`func (o *RepositoryX509RsaCertificate) SetCertificateChainNil(b bool)`

 SetCertificateChainNil sets the value for CertificateChain to be an explicit nil

### UnsetCertificateChain
`func (o *RepositoryX509RsaCertificate) UnsetCertificateChain()`

UnsetCertificateChain ensures that no value is present for CertificateChain, not even an explicit nil
### GetCertificateChainFingerprint

`func (o *RepositoryX509RsaCertificate) GetCertificateChainFingerprint() string`

GetCertificateChainFingerprint returns the CertificateChainFingerprint field if non-nil, zero value otherwise.

### GetCertificateChainFingerprintOk

`func (o *RepositoryX509RsaCertificate) GetCertificateChainFingerprintOk() (*string, bool)`

GetCertificateChainFingerprintOk returns a tuple with the CertificateChainFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChainFingerprint

`func (o *RepositoryX509RsaCertificate) SetCertificateChainFingerprint(v string)`

SetCertificateChainFingerprint sets CertificateChainFingerprint field to given value.

### HasCertificateChainFingerprint

`func (o *RepositoryX509RsaCertificate) HasCertificateChainFingerprint() bool`

HasCertificateChainFingerprint returns a boolean if a field has been set.

### GetCertificateChainFingerprintShort

`func (o *RepositoryX509RsaCertificate) GetCertificateChainFingerprintShort() string`

GetCertificateChainFingerprintShort returns the CertificateChainFingerprintShort field if non-nil, zero value otherwise.

### GetCertificateChainFingerprintShortOk

`func (o *RepositoryX509RsaCertificate) GetCertificateChainFingerprintShortOk() (*string, bool)`

GetCertificateChainFingerprintShortOk returns a tuple with the CertificateChainFingerprintShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChainFingerprintShort

`func (o *RepositoryX509RsaCertificate) SetCertificateChainFingerprintShort(v string)`

SetCertificateChainFingerprintShort sets CertificateChainFingerprintShort field to given value.

### HasCertificateChainFingerprintShort

`func (o *RepositoryX509RsaCertificate) HasCertificateChainFingerprintShort() bool`

HasCertificateChainFingerprintShort returns a boolean if a field has been set.

### GetCertificateFingerprint

`func (o *RepositoryX509RsaCertificate) GetCertificateFingerprint() string`

GetCertificateFingerprint returns the CertificateFingerprint field if non-nil, zero value otherwise.

### GetCertificateFingerprintOk

`func (o *RepositoryX509RsaCertificate) GetCertificateFingerprintOk() (*string, bool)`

GetCertificateFingerprintOk returns a tuple with the CertificateFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFingerprint

`func (o *RepositoryX509RsaCertificate) SetCertificateFingerprint(v string)`

SetCertificateFingerprint sets CertificateFingerprint field to given value.

### HasCertificateFingerprint

`func (o *RepositoryX509RsaCertificate) HasCertificateFingerprint() bool`

HasCertificateFingerprint returns a boolean if a field has been set.

### SetCertificateFingerprintNil

`func (o *RepositoryX509RsaCertificate) SetCertificateFingerprintNil(b bool)`

 SetCertificateFingerprintNil sets the value for CertificateFingerprint to be an explicit nil

### UnsetCertificateFingerprint
`func (o *RepositoryX509RsaCertificate) UnsetCertificateFingerprint()`

UnsetCertificateFingerprint ensures that no value is present for CertificateFingerprint, not even an explicit nil
### GetCertificateFingerprintShort

`func (o *RepositoryX509RsaCertificate) GetCertificateFingerprintShort() string`

GetCertificateFingerprintShort returns the CertificateFingerprintShort field if non-nil, zero value otherwise.

### GetCertificateFingerprintShortOk

`func (o *RepositoryX509RsaCertificate) GetCertificateFingerprintShortOk() (*string, bool)`

GetCertificateFingerprintShortOk returns a tuple with the CertificateFingerprintShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFingerprintShort

`func (o *RepositoryX509RsaCertificate) SetCertificateFingerprintShort(v string)`

SetCertificateFingerprintShort sets CertificateFingerprintShort field to given value.

### HasCertificateFingerprintShort

`func (o *RepositoryX509RsaCertificate) HasCertificateFingerprintShort() bool`

HasCertificateFingerprintShort returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RepositoryX509RsaCertificate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryX509RsaCertificate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryX509RsaCertificate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepositoryX509RsaCertificate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *RepositoryX509RsaCertificate) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RepositoryX509RsaCertificate) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RepositoryX509RsaCertificate) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RepositoryX509RsaCertificate) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetIssuingStatus

`func (o *RepositoryX509RsaCertificate) GetIssuingStatus() string`

GetIssuingStatus returns the IssuingStatus field if non-nil, zero value otherwise.

### GetIssuingStatusOk

`func (o *RepositoryX509RsaCertificate) GetIssuingStatusOk() (*string, bool)`

GetIssuingStatusOk returns a tuple with the IssuingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingStatus

`func (o *RepositoryX509RsaCertificate) SetIssuingStatus(v string)`

SetIssuingStatus sets IssuingStatus field to given value.

### HasIssuingStatus

`func (o *RepositoryX509RsaCertificate) HasIssuingStatus() bool`

HasIssuingStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


