# RepositoryX509EcdsaCertificate

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

### NewRepositoryX509EcdsaCertificate

`func NewRepositoryX509EcdsaCertificate() *RepositoryX509EcdsaCertificate`

NewRepositoryX509EcdsaCertificate instantiates a new RepositoryX509EcdsaCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryX509EcdsaCertificateWithDefaults

`func NewRepositoryX509EcdsaCertificateWithDefaults() *RepositoryX509EcdsaCertificate`

NewRepositoryX509EcdsaCertificateWithDefaults instantiates a new RepositoryX509EcdsaCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *RepositoryX509EcdsaCertificate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RepositoryX509EcdsaCertificate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RepositoryX509EcdsaCertificate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RepositoryX509EcdsaCertificate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCertificate

`func (o *RepositoryX509EcdsaCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *RepositoryX509EcdsaCertificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *RepositoryX509EcdsaCertificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *RepositoryX509EcdsaCertificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *RepositoryX509EcdsaCertificate) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *RepositoryX509EcdsaCertificate) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCertificateChain

`func (o *RepositoryX509EcdsaCertificate) GetCertificateChain() string`

GetCertificateChain returns the CertificateChain field if non-nil, zero value otherwise.

### GetCertificateChainOk

`func (o *RepositoryX509EcdsaCertificate) GetCertificateChainOk() (*string, bool)`

GetCertificateChainOk returns a tuple with the CertificateChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChain

`func (o *RepositoryX509EcdsaCertificate) SetCertificateChain(v string)`

SetCertificateChain sets CertificateChain field to given value.

### HasCertificateChain

`func (o *RepositoryX509EcdsaCertificate) HasCertificateChain() bool`

HasCertificateChain returns a boolean if a field has been set.

### SetCertificateChainNil

`func (o *RepositoryX509EcdsaCertificate) SetCertificateChainNil(b bool)`

 SetCertificateChainNil sets the value for CertificateChain to be an explicit nil

### UnsetCertificateChain
`func (o *RepositoryX509EcdsaCertificate) UnsetCertificateChain()`

UnsetCertificateChain ensures that no value is present for CertificateChain, not even an explicit nil
### GetCertificateChainFingerprint

`func (o *RepositoryX509EcdsaCertificate) GetCertificateChainFingerprint() string`

GetCertificateChainFingerprint returns the CertificateChainFingerprint field if non-nil, zero value otherwise.

### GetCertificateChainFingerprintOk

`func (o *RepositoryX509EcdsaCertificate) GetCertificateChainFingerprintOk() (*string, bool)`

GetCertificateChainFingerprintOk returns a tuple with the CertificateChainFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChainFingerprint

`func (o *RepositoryX509EcdsaCertificate) SetCertificateChainFingerprint(v string)`

SetCertificateChainFingerprint sets CertificateChainFingerprint field to given value.

### HasCertificateChainFingerprint

`func (o *RepositoryX509EcdsaCertificate) HasCertificateChainFingerprint() bool`

HasCertificateChainFingerprint returns a boolean if a field has been set.

### GetCertificateChainFingerprintShort

`func (o *RepositoryX509EcdsaCertificate) GetCertificateChainFingerprintShort() string`

GetCertificateChainFingerprintShort returns the CertificateChainFingerprintShort field if non-nil, zero value otherwise.

### GetCertificateChainFingerprintShortOk

`func (o *RepositoryX509EcdsaCertificate) GetCertificateChainFingerprintShortOk() (*string, bool)`

GetCertificateChainFingerprintShortOk returns a tuple with the CertificateChainFingerprintShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChainFingerprintShort

`func (o *RepositoryX509EcdsaCertificate) SetCertificateChainFingerprintShort(v string)`

SetCertificateChainFingerprintShort sets CertificateChainFingerprintShort field to given value.

### HasCertificateChainFingerprintShort

`func (o *RepositoryX509EcdsaCertificate) HasCertificateChainFingerprintShort() bool`

HasCertificateChainFingerprintShort returns a boolean if a field has been set.

### GetCertificateFingerprint

`func (o *RepositoryX509EcdsaCertificate) GetCertificateFingerprint() string`

GetCertificateFingerprint returns the CertificateFingerprint field if non-nil, zero value otherwise.

### GetCertificateFingerprintOk

`func (o *RepositoryX509EcdsaCertificate) GetCertificateFingerprintOk() (*string, bool)`

GetCertificateFingerprintOk returns a tuple with the CertificateFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFingerprint

`func (o *RepositoryX509EcdsaCertificate) SetCertificateFingerprint(v string)`

SetCertificateFingerprint sets CertificateFingerprint field to given value.

### HasCertificateFingerprint

`func (o *RepositoryX509EcdsaCertificate) HasCertificateFingerprint() bool`

HasCertificateFingerprint returns a boolean if a field has been set.

### SetCertificateFingerprintNil

`func (o *RepositoryX509EcdsaCertificate) SetCertificateFingerprintNil(b bool)`

 SetCertificateFingerprintNil sets the value for CertificateFingerprint to be an explicit nil

### UnsetCertificateFingerprint
`func (o *RepositoryX509EcdsaCertificate) UnsetCertificateFingerprint()`

UnsetCertificateFingerprint ensures that no value is present for CertificateFingerprint, not even an explicit nil
### GetCertificateFingerprintShort

`func (o *RepositoryX509EcdsaCertificate) GetCertificateFingerprintShort() string`

GetCertificateFingerprintShort returns the CertificateFingerprintShort field if non-nil, zero value otherwise.

### GetCertificateFingerprintShortOk

`func (o *RepositoryX509EcdsaCertificate) GetCertificateFingerprintShortOk() (*string, bool)`

GetCertificateFingerprintShortOk returns a tuple with the CertificateFingerprintShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFingerprintShort

`func (o *RepositoryX509EcdsaCertificate) SetCertificateFingerprintShort(v string)`

SetCertificateFingerprintShort sets CertificateFingerprintShort field to given value.

### HasCertificateFingerprintShort

`func (o *RepositoryX509EcdsaCertificate) HasCertificateFingerprintShort() bool`

HasCertificateFingerprintShort returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RepositoryX509EcdsaCertificate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryX509EcdsaCertificate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryX509EcdsaCertificate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepositoryX509EcdsaCertificate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *RepositoryX509EcdsaCertificate) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RepositoryX509EcdsaCertificate) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RepositoryX509EcdsaCertificate) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RepositoryX509EcdsaCertificate) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetIssuingStatus

`func (o *RepositoryX509EcdsaCertificate) GetIssuingStatus() string`

GetIssuingStatus returns the IssuingStatus field if non-nil, zero value otherwise.

### GetIssuingStatusOk

`func (o *RepositoryX509EcdsaCertificate) GetIssuingStatusOk() (*string, bool)`

GetIssuingStatusOk returns a tuple with the IssuingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingStatus

`func (o *RepositoryX509EcdsaCertificate) SetIssuingStatus(v string)`

SetIssuingStatus sets IssuingStatus field to given value.

### HasIssuingStatus

`func (o *RepositoryX509EcdsaCertificate) HasIssuingStatus() bool`

HasIssuingStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


