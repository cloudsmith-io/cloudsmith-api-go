# RepositoryX509Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | If selected this is the active key for this repository. | [optional] [readonly] 
**Certificate** | Pointer to **NullableString** | The issued certificate. | [optional] [readonly] 
**CertificateFingerprint** | Pointer to **NullableString** | The SHA-256 long identifier used | [optional] [readonly] 
**CertificateFingerprintShort** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Default** | Pointer to **bool** | If selected this is the default key for this repository. | [optional] [readonly] 
**IssuingStatus** | Pointer to **string** |  | [optional] [default to "Certificate is pending to be issued"]

## Methods

### NewRepositoryX509Certificate

`func NewRepositoryX509Certificate() *RepositoryX509Certificate`

NewRepositoryX509Certificate instantiates a new RepositoryX509Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryX509CertificateWithDefaults

`func NewRepositoryX509CertificateWithDefaults() *RepositoryX509Certificate`

NewRepositoryX509CertificateWithDefaults instantiates a new RepositoryX509Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *RepositoryX509Certificate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RepositoryX509Certificate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RepositoryX509Certificate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RepositoryX509Certificate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCertificate

`func (o *RepositoryX509Certificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *RepositoryX509Certificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *RepositoryX509Certificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *RepositoryX509Certificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *RepositoryX509Certificate) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *RepositoryX509Certificate) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCertificateFingerprint

`func (o *RepositoryX509Certificate) GetCertificateFingerprint() string`

GetCertificateFingerprint returns the CertificateFingerprint field if non-nil, zero value otherwise.

### GetCertificateFingerprintOk

`func (o *RepositoryX509Certificate) GetCertificateFingerprintOk() (*string, bool)`

GetCertificateFingerprintOk returns a tuple with the CertificateFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFingerprint

`func (o *RepositoryX509Certificate) SetCertificateFingerprint(v string)`

SetCertificateFingerprint sets CertificateFingerprint field to given value.

### HasCertificateFingerprint

`func (o *RepositoryX509Certificate) HasCertificateFingerprint() bool`

HasCertificateFingerprint returns a boolean if a field has been set.

### SetCertificateFingerprintNil

`func (o *RepositoryX509Certificate) SetCertificateFingerprintNil(b bool)`

 SetCertificateFingerprintNil sets the value for CertificateFingerprint to be an explicit nil

### UnsetCertificateFingerprint
`func (o *RepositoryX509Certificate) UnsetCertificateFingerprint()`

UnsetCertificateFingerprint ensures that no value is present for CertificateFingerprint, not even an explicit nil
### GetCertificateFingerprintShort

`func (o *RepositoryX509Certificate) GetCertificateFingerprintShort() string`

GetCertificateFingerprintShort returns the CertificateFingerprintShort field if non-nil, zero value otherwise.

### GetCertificateFingerprintShortOk

`func (o *RepositoryX509Certificate) GetCertificateFingerprintShortOk() (*string, bool)`

GetCertificateFingerprintShortOk returns a tuple with the CertificateFingerprintShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFingerprintShort

`func (o *RepositoryX509Certificate) SetCertificateFingerprintShort(v string)`

SetCertificateFingerprintShort sets CertificateFingerprintShort field to given value.

### HasCertificateFingerprintShort

`func (o *RepositoryX509Certificate) HasCertificateFingerprintShort() bool`

HasCertificateFingerprintShort returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RepositoryX509Certificate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryX509Certificate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryX509Certificate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepositoryX509Certificate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *RepositoryX509Certificate) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RepositoryX509Certificate) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RepositoryX509Certificate) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RepositoryX509Certificate) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetIssuingStatus

`func (o *RepositoryX509Certificate) GetIssuingStatus() string`

GetIssuingStatus returns the IssuingStatus field if non-nil, zero value otherwise.

### GetIssuingStatusOk

`func (o *RepositoryX509Certificate) GetIssuingStatusOk() (*string, bool)`

GetIssuingStatusOk returns a tuple with the IssuingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingStatus

`func (o *RepositoryX509Certificate) SetIssuingStatus(v string)`

SetIssuingStatus sets IssuingStatus field to given value.

### HasIssuingStatus

`func (o *RepositoryX509Certificate) HasIssuingStatus() bool`

HasIssuingStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


