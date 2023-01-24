# PackageFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdnUrl** | Pointer to **NullableString** |  | [optional] [readonly] 
**ChecksumMd5** | Pointer to **NullableString** |  | [optional] [readonly] 
**ChecksumSha1** | Pointer to **NullableString** |  | [optional] [readonly] 
**ChecksumSha256** | Pointer to **NullableString** |  | [optional] [readonly] 
**ChecksumSha512** | Pointer to **NullableString** |  | [optional] [readonly] 
**Downloads** | Pointer to **int64** |  | [optional] [readonly] 
**Filename** | Pointer to **string** |  | [optional] [readonly] 
**IsDownloadable** | Pointer to **bool** |  | [optional] [readonly] 
**IsPrimary** | Pointer to **bool** |  | [optional] [readonly] 
**IsSynchronised** | Pointer to **bool** |  | [optional] [readonly] 
**SignatureUrl** | Pointer to **NullableString** |  | [optional] [readonly] 
**Size** | Pointer to **int64** | The calculated size of the file. | [optional] [readonly] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**Tag** | Pointer to **NullableString** | Freeform descriptor that describes what the file is. | [optional] [readonly] 

## Methods

### NewPackageFile

`func NewPackageFile() *PackageFile`

NewPackageFile instantiates a new PackageFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageFileWithDefaults

`func NewPackageFileWithDefaults() *PackageFile`

NewPackageFileWithDefaults instantiates a new PackageFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdnUrl

`func (o *PackageFile) GetCdnUrl() string`

GetCdnUrl returns the CdnUrl field if non-nil, zero value otherwise.

### GetCdnUrlOk

`func (o *PackageFile) GetCdnUrlOk() (*string, bool)`

GetCdnUrlOk returns a tuple with the CdnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnUrl

`func (o *PackageFile) SetCdnUrl(v string)`

SetCdnUrl sets CdnUrl field to given value.

### HasCdnUrl

`func (o *PackageFile) HasCdnUrl() bool`

HasCdnUrl returns a boolean if a field has been set.

### SetCdnUrlNil

`func (o *PackageFile) SetCdnUrlNil(b bool)`

 SetCdnUrlNil sets the value for CdnUrl to be an explicit nil

### UnsetCdnUrl
`func (o *PackageFile) UnsetCdnUrl()`

UnsetCdnUrl ensures that no value is present for CdnUrl, not even an explicit nil
### GetChecksumMd5

`func (o *PackageFile) GetChecksumMd5() string`

GetChecksumMd5 returns the ChecksumMd5 field if non-nil, zero value otherwise.

### GetChecksumMd5Ok

`func (o *PackageFile) GetChecksumMd5Ok() (*string, bool)`

GetChecksumMd5Ok returns a tuple with the ChecksumMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumMd5

`func (o *PackageFile) SetChecksumMd5(v string)`

SetChecksumMd5 sets ChecksumMd5 field to given value.

### HasChecksumMd5

`func (o *PackageFile) HasChecksumMd5() bool`

HasChecksumMd5 returns a boolean if a field has been set.

### SetChecksumMd5Nil

`func (o *PackageFile) SetChecksumMd5Nil(b bool)`

 SetChecksumMd5Nil sets the value for ChecksumMd5 to be an explicit nil

### UnsetChecksumMd5
`func (o *PackageFile) UnsetChecksumMd5()`

UnsetChecksumMd5 ensures that no value is present for ChecksumMd5, not even an explicit nil
### GetChecksumSha1

`func (o *PackageFile) GetChecksumSha1() string`

GetChecksumSha1 returns the ChecksumSha1 field if non-nil, zero value otherwise.

### GetChecksumSha1Ok

`func (o *PackageFile) GetChecksumSha1Ok() (*string, bool)`

GetChecksumSha1Ok returns a tuple with the ChecksumSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha1

`func (o *PackageFile) SetChecksumSha1(v string)`

SetChecksumSha1 sets ChecksumSha1 field to given value.

### HasChecksumSha1

`func (o *PackageFile) HasChecksumSha1() bool`

HasChecksumSha1 returns a boolean if a field has been set.

### SetChecksumSha1Nil

`func (o *PackageFile) SetChecksumSha1Nil(b bool)`

 SetChecksumSha1Nil sets the value for ChecksumSha1 to be an explicit nil

### UnsetChecksumSha1
`func (o *PackageFile) UnsetChecksumSha1()`

UnsetChecksumSha1 ensures that no value is present for ChecksumSha1, not even an explicit nil
### GetChecksumSha256

`func (o *PackageFile) GetChecksumSha256() string`

GetChecksumSha256 returns the ChecksumSha256 field if non-nil, zero value otherwise.

### GetChecksumSha256Ok

`func (o *PackageFile) GetChecksumSha256Ok() (*string, bool)`

GetChecksumSha256Ok returns a tuple with the ChecksumSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha256

`func (o *PackageFile) SetChecksumSha256(v string)`

SetChecksumSha256 sets ChecksumSha256 field to given value.

### HasChecksumSha256

`func (o *PackageFile) HasChecksumSha256() bool`

HasChecksumSha256 returns a boolean if a field has been set.

### SetChecksumSha256Nil

`func (o *PackageFile) SetChecksumSha256Nil(b bool)`

 SetChecksumSha256Nil sets the value for ChecksumSha256 to be an explicit nil

### UnsetChecksumSha256
`func (o *PackageFile) UnsetChecksumSha256()`

UnsetChecksumSha256 ensures that no value is present for ChecksumSha256, not even an explicit nil
### GetChecksumSha512

`func (o *PackageFile) GetChecksumSha512() string`

GetChecksumSha512 returns the ChecksumSha512 field if non-nil, zero value otherwise.

### GetChecksumSha512Ok

`func (o *PackageFile) GetChecksumSha512Ok() (*string, bool)`

GetChecksumSha512Ok returns a tuple with the ChecksumSha512 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha512

`func (o *PackageFile) SetChecksumSha512(v string)`

SetChecksumSha512 sets ChecksumSha512 field to given value.

### HasChecksumSha512

`func (o *PackageFile) HasChecksumSha512() bool`

HasChecksumSha512 returns a boolean if a field has been set.

### SetChecksumSha512Nil

`func (o *PackageFile) SetChecksumSha512Nil(b bool)`

 SetChecksumSha512Nil sets the value for ChecksumSha512 to be an explicit nil

### UnsetChecksumSha512
`func (o *PackageFile) UnsetChecksumSha512()`

UnsetChecksumSha512 ensures that no value is present for ChecksumSha512, not even an explicit nil
### GetDownloads

`func (o *PackageFile) GetDownloads() int64`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *PackageFile) GetDownloadsOk() (*int64, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *PackageFile) SetDownloads(v int64)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *PackageFile) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetFilename

`func (o *PackageFile) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PackageFile) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PackageFile) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *PackageFile) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetIsDownloadable

`func (o *PackageFile) GetIsDownloadable() bool`

GetIsDownloadable returns the IsDownloadable field if non-nil, zero value otherwise.

### GetIsDownloadableOk

`func (o *PackageFile) GetIsDownloadableOk() (*bool, bool)`

GetIsDownloadableOk returns a tuple with the IsDownloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDownloadable

`func (o *PackageFile) SetIsDownloadable(v bool)`

SetIsDownloadable sets IsDownloadable field to given value.

### HasIsDownloadable

`func (o *PackageFile) HasIsDownloadable() bool`

HasIsDownloadable returns a boolean if a field has been set.

### GetIsPrimary

`func (o *PackageFile) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *PackageFile) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *PackageFile) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *PackageFile) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetIsSynchronised

`func (o *PackageFile) GetIsSynchronised() bool`

GetIsSynchronised returns the IsSynchronised field if non-nil, zero value otherwise.

### GetIsSynchronisedOk

`func (o *PackageFile) GetIsSynchronisedOk() (*bool, bool)`

GetIsSynchronisedOk returns a tuple with the IsSynchronised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSynchronised

`func (o *PackageFile) SetIsSynchronised(v bool)`

SetIsSynchronised sets IsSynchronised field to given value.

### HasIsSynchronised

`func (o *PackageFile) HasIsSynchronised() bool`

HasIsSynchronised returns a boolean if a field has been set.

### GetSignatureUrl

`func (o *PackageFile) GetSignatureUrl() string`

GetSignatureUrl returns the SignatureUrl field if non-nil, zero value otherwise.

### GetSignatureUrlOk

`func (o *PackageFile) GetSignatureUrlOk() (*string, bool)`

GetSignatureUrlOk returns a tuple with the SignatureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureUrl

`func (o *PackageFile) SetSignatureUrl(v string)`

SetSignatureUrl sets SignatureUrl field to given value.

### HasSignatureUrl

`func (o *PackageFile) HasSignatureUrl() bool`

HasSignatureUrl returns a boolean if a field has been set.

### SetSignatureUrlNil

`func (o *PackageFile) SetSignatureUrlNil(b bool)`

 SetSignatureUrlNil sets the value for SignatureUrl to be an explicit nil

### UnsetSignatureUrl
`func (o *PackageFile) UnsetSignatureUrl()`

UnsetSignatureUrl ensures that no value is present for SignatureUrl, not even an explicit nil
### GetSize

`func (o *PackageFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PackageFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PackageFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *PackageFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSlugPerm

`func (o *PackageFile) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *PackageFile) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *PackageFile) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *PackageFile) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetTag

`func (o *PackageFile) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PackageFile) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PackageFile) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PackageFile) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *PackageFile) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *PackageFile) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


