# ConanPackageUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConanChannel** | Pointer to **NullableString** | Conan channel. | [optional] 
**ConanPrefix** | Pointer to **NullableString** | Conan prefix (User). | [optional] 
**InfoFile** | **string** | The info file is an python file containing the package metadata. | 
**ManifestFile** | **string** | The info file is an python file containing the package metadata. | 
**MetadataFile** | **string** | The conan file is an python file containing the package metadata. | 
**Name** | Pointer to **NullableString** | The name of this package. | [optional] 
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Tags** | Pointer to **NullableString** | A comma-separated values list of tags to add to the package. | [optional] 
**Version** | Pointer to **NullableString** | The raw version for this package. | [optional] 

## Methods

### NewConanPackageUploadRequest

`func NewConanPackageUploadRequest(infoFile string, manifestFile string, metadataFile string, packageFile string, ) *ConanPackageUploadRequest`

NewConanPackageUploadRequest instantiates a new ConanPackageUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConanPackageUploadRequestWithDefaults

`func NewConanPackageUploadRequestWithDefaults() *ConanPackageUploadRequest`

NewConanPackageUploadRequestWithDefaults instantiates a new ConanPackageUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConanChannel

`func (o *ConanPackageUploadRequest) GetConanChannel() string`

GetConanChannel returns the ConanChannel field if non-nil, zero value otherwise.

### GetConanChannelOk

`func (o *ConanPackageUploadRequest) GetConanChannelOk() (*string, bool)`

GetConanChannelOk returns a tuple with the ConanChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConanChannel

`func (o *ConanPackageUploadRequest) SetConanChannel(v string)`

SetConanChannel sets ConanChannel field to given value.

### HasConanChannel

`func (o *ConanPackageUploadRequest) HasConanChannel() bool`

HasConanChannel returns a boolean if a field has been set.

### SetConanChannelNil

`func (o *ConanPackageUploadRequest) SetConanChannelNil(b bool)`

 SetConanChannelNil sets the value for ConanChannel to be an explicit nil

### UnsetConanChannel
`func (o *ConanPackageUploadRequest) UnsetConanChannel()`

UnsetConanChannel ensures that no value is present for ConanChannel, not even an explicit nil
### GetConanPrefix

`func (o *ConanPackageUploadRequest) GetConanPrefix() string`

GetConanPrefix returns the ConanPrefix field if non-nil, zero value otherwise.

### GetConanPrefixOk

`func (o *ConanPackageUploadRequest) GetConanPrefixOk() (*string, bool)`

GetConanPrefixOk returns a tuple with the ConanPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConanPrefix

`func (o *ConanPackageUploadRequest) SetConanPrefix(v string)`

SetConanPrefix sets ConanPrefix field to given value.

### HasConanPrefix

`func (o *ConanPackageUploadRequest) HasConanPrefix() bool`

HasConanPrefix returns a boolean if a field has been set.

### SetConanPrefixNil

`func (o *ConanPackageUploadRequest) SetConanPrefixNil(b bool)`

 SetConanPrefixNil sets the value for ConanPrefix to be an explicit nil

### UnsetConanPrefix
`func (o *ConanPackageUploadRequest) UnsetConanPrefix()`

UnsetConanPrefix ensures that no value is present for ConanPrefix, not even an explicit nil
### GetInfoFile

`func (o *ConanPackageUploadRequest) GetInfoFile() string`

GetInfoFile returns the InfoFile field if non-nil, zero value otherwise.

### GetInfoFileOk

`func (o *ConanPackageUploadRequest) GetInfoFileOk() (*string, bool)`

GetInfoFileOk returns a tuple with the InfoFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoFile

`func (o *ConanPackageUploadRequest) SetInfoFile(v string)`

SetInfoFile sets InfoFile field to given value.


### GetManifestFile

`func (o *ConanPackageUploadRequest) GetManifestFile() string`

GetManifestFile returns the ManifestFile field if non-nil, zero value otherwise.

### GetManifestFileOk

`func (o *ConanPackageUploadRequest) GetManifestFileOk() (*string, bool)`

GetManifestFileOk returns a tuple with the ManifestFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestFile

`func (o *ConanPackageUploadRequest) SetManifestFile(v string)`

SetManifestFile sets ManifestFile field to given value.


### GetMetadataFile

`func (o *ConanPackageUploadRequest) GetMetadataFile() string`

GetMetadataFile returns the MetadataFile field if non-nil, zero value otherwise.

### GetMetadataFileOk

`func (o *ConanPackageUploadRequest) GetMetadataFileOk() (*string, bool)`

GetMetadataFileOk returns a tuple with the MetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFile

`func (o *ConanPackageUploadRequest) SetMetadataFile(v string)`

SetMetadataFile sets MetadataFile field to given value.


### GetName

`func (o *ConanPackageUploadRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConanPackageUploadRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConanPackageUploadRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConanPackageUploadRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConanPackageUploadRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConanPackageUploadRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPackageFile

`func (o *ConanPackageUploadRequest) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *ConanPackageUploadRequest) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *ConanPackageUploadRequest) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRepublish

`func (o *ConanPackageUploadRequest) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *ConanPackageUploadRequest) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *ConanPackageUploadRequest) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *ConanPackageUploadRequest) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetTags

`func (o *ConanPackageUploadRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConanPackageUploadRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConanPackageUploadRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConanPackageUploadRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ConanPackageUploadRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ConanPackageUploadRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetVersion

`func (o *ConanPackageUploadRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConanPackageUploadRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConanPackageUploadRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConanPackageUploadRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ConanPackageUploadRequest) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ConanPackageUploadRequest) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


