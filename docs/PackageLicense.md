# PackageLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseNotes** | Pointer to **NullableString** |  | [optional] 
**LicenseOverride** | Pointer to **NullableString** |  | [optional] [default to "None"]
**LicenseUrl** | Pointer to **NullableString** |  | [optional] 
**SpdxLicense** | **string** |  | 

## Methods

### NewPackageLicense

`func NewPackageLicense(spdxLicense string, ) *PackageLicense`

NewPackageLicense instantiates a new PackageLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageLicenseWithDefaults

`func NewPackageLicenseWithDefaults() *PackageLicense`

NewPackageLicenseWithDefaults instantiates a new PackageLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseNotes

`func (o *PackageLicense) GetLicenseNotes() string`

GetLicenseNotes returns the LicenseNotes field if non-nil, zero value otherwise.

### GetLicenseNotesOk

`func (o *PackageLicense) GetLicenseNotesOk() (*string, bool)`

GetLicenseNotesOk returns a tuple with the LicenseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNotes

`func (o *PackageLicense) SetLicenseNotes(v string)`

SetLicenseNotes sets LicenseNotes field to given value.

### HasLicenseNotes

`func (o *PackageLicense) HasLicenseNotes() bool`

HasLicenseNotes returns a boolean if a field has been set.

### SetLicenseNotesNil

`func (o *PackageLicense) SetLicenseNotesNil(b bool)`

 SetLicenseNotesNil sets the value for LicenseNotes to be an explicit nil

### UnsetLicenseNotes
`func (o *PackageLicense) UnsetLicenseNotes()`

UnsetLicenseNotes ensures that no value is present for LicenseNotes, not even an explicit nil
### GetLicenseOverride

`func (o *PackageLicense) GetLicenseOverride() string`

GetLicenseOverride returns the LicenseOverride field if non-nil, zero value otherwise.

### GetLicenseOverrideOk

`func (o *PackageLicense) GetLicenseOverrideOk() (*string, bool)`

GetLicenseOverrideOk returns a tuple with the LicenseOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseOverride

`func (o *PackageLicense) SetLicenseOverride(v string)`

SetLicenseOverride sets LicenseOverride field to given value.

### HasLicenseOverride

`func (o *PackageLicense) HasLicenseOverride() bool`

HasLicenseOverride returns a boolean if a field has been set.

### SetLicenseOverrideNil

`func (o *PackageLicense) SetLicenseOverrideNil(b bool)`

 SetLicenseOverrideNil sets the value for LicenseOverride to be an explicit nil

### UnsetLicenseOverride
`func (o *PackageLicense) UnsetLicenseOverride()`

UnsetLicenseOverride ensures that no value is present for LicenseOverride, not even an explicit nil
### GetLicenseUrl

`func (o *PackageLicense) GetLicenseUrl() string`

GetLicenseUrl returns the LicenseUrl field if non-nil, zero value otherwise.

### GetLicenseUrlOk

`func (o *PackageLicense) GetLicenseUrlOk() (*string, bool)`

GetLicenseUrlOk returns a tuple with the LicenseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUrl

`func (o *PackageLicense) SetLicenseUrl(v string)`

SetLicenseUrl sets LicenseUrl field to given value.

### HasLicenseUrl

`func (o *PackageLicense) HasLicenseUrl() bool`

HasLicenseUrl returns a boolean if a field has been set.

### SetLicenseUrlNil

`func (o *PackageLicense) SetLicenseUrlNil(b bool)`

 SetLicenseUrlNil sets the value for LicenseUrl to be an explicit nil

### UnsetLicenseUrl
`func (o *PackageLicense) UnsetLicenseUrl()`

UnsetLicenseUrl ensures that no value is present for LicenseUrl, not even an explicit nil
### GetSpdxLicense

`func (o *PackageLicense) GetSpdxLicense() string`

GetSpdxLicense returns the SpdxLicense field if non-nil, zero value otherwise.

### GetSpdxLicenseOk

`func (o *PackageLicense) GetSpdxLicenseOk() (*string, bool)`

GetSpdxLicenseOk returns a tuple with the SpdxLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpdxLicense

`func (o *PackageLicense) SetSpdxLicense(v string)`

SetSpdxLicense sets SpdxLicense field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


