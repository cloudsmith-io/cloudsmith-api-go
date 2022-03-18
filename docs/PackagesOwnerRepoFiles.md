# PackagesOwnerRepoFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdnUrl** | Pointer to **string** |  | [optional] 
**ChecksumMd5** | Pointer to **string** |  | [optional] 
**ChecksumSha1** | Pointer to **string** |  | [optional] 
**ChecksumSha256** | Pointer to **string** |  | [optional] 
**ChecksumSha512** | Pointer to **string** |  | [optional] 
**Downloads** | Pointer to **int64** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**IsDownloadable** | Pointer to **bool** |  | [optional] 
**IsPrimary** | Pointer to **bool** |  | [optional] 
**IsSynchronised** | Pointer to **bool** |  | [optional] 
**SignatureUrl** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** | The calculated size of the file. | [optional] 
**SlugPerm** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** | Freeform descriptor that describes what the file is. | [optional] 

## Methods

### NewPackagesOwnerRepoFiles

`func NewPackagesOwnerRepoFiles() *PackagesOwnerRepoFiles`

NewPackagesOwnerRepoFiles instantiates a new PackagesOwnerRepoFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesOwnerRepoFilesWithDefaults

`func NewPackagesOwnerRepoFilesWithDefaults() *PackagesOwnerRepoFiles`

NewPackagesOwnerRepoFilesWithDefaults instantiates a new PackagesOwnerRepoFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdnUrl

`func (o *PackagesOwnerRepoFiles) GetCdnUrl() string`

GetCdnUrl returns the CdnUrl field if non-nil, zero value otherwise.

### GetCdnUrlOk

`func (o *PackagesOwnerRepoFiles) GetCdnUrlOk() (*string, bool)`

GetCdnUrlOk returns a tuple with the CdnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnUrl

`func (o *PackagesOwnerRepoFiles) SetCdnUrl(v string)`

SetCdnUrl sets CdnUrl field to given value.

### HasCdnUrl

`func (o *PackagesOwnerRepoFiles) HasCdnUrl() bool`

HasCdnUrl returns a boolean if a field has been set.

### GetChecksumMd5

`func (o *PackagesOwnerRepoFiles) GetChecksumMd5() string`

GetChecksumMd5 returns the ChecksumMd5 field if non-nil, zero value otherwise.

### GetChecksumMd5Ok

`func (o *PackagesOwnerRepoFiles) GetChecksumMd5Ok() (*string, bool)`

GetChecksumMd5Ok returns a tuple with the ChecksumMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumMd5

`func (o *PackagesOwnerRepoFiles) SetChecksumMd5(v string)`

SetChecksumMd5 sets ChecksumMd5 field to given value.

### HasChecksumMd5

`func (o *PackagesOwnerRepoFiles) HasChecksumMd5() bool`

HasChecksumMd5 returns a boolean if a field has been set.

### GetChecksumSha1

`func (o *PackagesOwnerRepoFiles) GetChecksumSha1() string`

GetChecksumSha1 returns the ChecksumSha1 field if non-nil, zero value otherwise.

### GetChecksumSha1Ok

`func (o *PackagesOwnerRepoFiles) GetChecksumSha1Ok() (*string, bool)`

GetChecksumSha1Ok returns a tuple with the ChecksumSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha1

`func (o *PackagesOwnerRepoFiles) SetChecksumSha1(v string)`

SetChecksumSha1 sets ChecksumSha1 field to given value.

### HasChecksumSha1

`func (o *PackagesOwnerRepoFiles) HasChecksumSha1() bool`

HasChecksumSha1 returns a boolean if a field has been set.

### GetChecksumSha256

`func (o *PackagesOwnerRepoFiles) GetChecksumSha256() string`

GetChecksumSha256 returns the ChecksumSha256 field if non-nil, zero value otherwise.

### GetChecksumSha256Ok

`func (o *PackagesOwnerRepoFiles) GetChecksumSha256Ok() (*string, bool)`

GetChecksumSha256Ok returns a tuple with the ChecksumSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha256

`func (o *PackagesOwnerRepoFiles) SetChecksumSha256(v string)`

SetChecksumSha256 sets ChecksumSha256 field to given value.

### HasChecksumSha256

`func (o *PackagesOwnerRepoFiles) HasChecksumSha256() bool`

HasChecksumSha256 returns a boolean if a field has been set.

### GetChecksumSha512

`func (o *PackagesOwnerRepoFiles) GetChecksumSha512() string`

GetChecksumSha512 returns the ChecksumSha512 field if non-nil, zero value otherwise.

### GetChecksumSha512Ok

`func (o *PackagesOwnerRepoFiles) GetChecksumSha512Ok() (*string, bool)`

GetChecksumSha512Ok returns a tuple with the ChecksumSha512 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha512

`func (o *PackagesOwnerRepoFiles) SetChecksumSha512(v string)`

SetChecksumSha512 sets ChecksumSha512 field to given value.

### HasChecksumSha512

`func (o *PackagesOwnerRepoFiles) HasChecksumSha512() bool`

HasChecksumSha512 returns a boolean if a field has been set.

### GetDownloads

`func (o *PackagesOwnerRepoFiles) GetDownloads() int64`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *PackagesOwnerRepoFiles) GetDownloadsOk() (*int64, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *PackagesOwnerRepoFiles) SetDownloads(v int64)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *PackagesOwnerRepoFiles) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetFilename

`func (o *PackagesOwnerRepoFiles) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PackagesOwnerRepoFiles) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PackagesOwnerRepoFiles) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *PackagesOwnerRepoFiles) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetIsDownloadable

`func (o *PackagesOwnerRepoFiles) GetIsDownloadable() bool`

GetIsDownloadable returns the IsDownloadable field if non-nil, zero value otherwise.

### GetIsDownloadableOk

`func (o *PackagesOwnerRepoFiles) GetIsDownloadableOk() (*bool, bool)`

GetIsDownloadableOk returns a tuple with the IsDownloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDownloadable

`func (o *PackagesOwnerRepoFiles) SetIsDownloadable(v bool)`

SetIsDownloadable sets IsDownloadable field to given value.

### HasIsDownloadable

`func (o *PackagesOwnerRepoFiles) HasIsDownloadable() bool`

HasIsDownloadable returns a boolean if a field has been set.

### GetIsPrimary

`func (o *PackagesOwnerRepoFiles) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *PackagesOwnerRepoFiles) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *PackagesOwnerRepoFiles) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *PackagesOwnerRepoFiles) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetIsSynchronised

`func (o *PackagesOwnerRepoFiles) GetIsSynchronised() bool`

GetIsSynchronised returns the IsSynchronised field if non-nil, zero value otherwise.

### GetIsSynchronisedOk

`func (o *PackagesOwnerRepoFiles) GetIsSynchronisedOk() (*bool, bool)`

GetIsSynchronisedOk returns a tuple with the IsSynchronised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSynchronised

`func (o *PackagesOwnerRepoFiles) SetIsSynchronised(v bool)`

SetIsSynchronised sets IsSynchronised field to given value.

### HasIsSynchronised

`func (o *PackagesOwnerRepoFiles) HasIsSynchronised() bool`

HasIsSynchronised returns a boolean if a field has been set.

### GetSignatureUrl

`func (o *PackagesOwnerRepoFiles) GetSignatureUrl() string`

GetSignatureUrl returns the SignatureUrl field if non-nil, zero value otherwise.

### GetSignatureUrlOk

`func (o *PackagesOwnerRepoFiles) GetSignatureUrlOk() (*string, bool)`

GetSignatureUrlOk returns a tuple with the SignatureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureUrl

`func (o *PackagesOwnerRepoFiles) SetSignatureUrl(v string)`

SetSignatureUrl sets SignatureUrl field to given value.

### HasSignatureUrl

`func (o *PackagesOwnerRepoFiles) HasSignatureUrl() bool`

HasSignatureUrl returns a boolean if a field has been set.

### GetSize

`func (o *PackagesOwnerRepoFiles) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PackagesOwnerRepoFiles) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PackagesOwnerRepoFiles) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *PackagesOwnerRepoFiles) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSlugPerm

`func (o *PackagesOwnerRepoFiles) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *PackagesOwnerRepoFiles) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *PackagesOwnerRepoFiles) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *PackagesOwnerRepoFiles) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetTag

`func (o *PackagesOwnerRepoFiles) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PackagesOwnerRepoFiles) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PackagesOwnerRepoFiles) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PackagesOwnerRepoFiles) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


