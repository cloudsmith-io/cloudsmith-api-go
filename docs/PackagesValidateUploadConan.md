# PackagesValidateUploadConan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConanChannel** | Pointer to **string** | Conan channel. | [optional] 
**ConanPrefix** | Pointer to **string** | Conan prefix (User). | [optional] 
**InfoFile** | **string** | The info file is an python file containing the package metadata. | 
**ManifestFile** | **string** | The info file is an python file containing the package metadata. | 
**MetadataFile** | **string** | The conan file is an python file containing the package metadata. | 
**Name** | Pointer to **string** | The name of this package. | [optional] 
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Tags** | Pointer to **string** | A comma-separated values list of tags to add to the package. | [optional] 
**Version** | Pointer to **string** | The raw version for this package. | [optional] 

## Methods

### NewPackagesValidateUploadConan

`func NewPackagesValidateUploadConan(infoFile string, manifestFile string, metadataFile string, packageFile string, ) *PackagesValidateUploadConan`

NewPackagesValidateUploadConan instantiates a new PackagesValidateUploadConan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesValidateUploadConanWithDefaults

`func NewPackagesValidateUploadConanWithDefaults() *PackagesValidateUploadConan`

NewPackagesValidateUploadConanWithDefaults instantiates a new PackagesValidateUploadConan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConanChannel

`func (o *PackagesValidateUploadConan) GetConanChannel() string`

GetConanChannel returns the ConanChannel field if non-nil, zero value otherwise.

### GetConanChannelOk

`func (o *PackagesValidateUploadConan) GetConanChannelOk() (*string, bool)`

GetConanChannelOk returns a tuple with the ConanChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConanChannel

`func (o *PackagesValidateUploadConan) SetConanChannel(v string)`

SetConanChannel sets ConanChannel field to given value.

### HasConanChannel

`func (o *PackagesValidateUploadConan) HasConanChannel() bool`

HasConanChannel returns a boolean if a field has been set.

### GetConanPrefix

`func (o *PackagesValidateUploadConan) GetConanPrefix() string`

GetConanPrefix returns the ConanPrefix field if non-nil, zero value otherwise.

### GetConanPrefixOk

`func (o *PackagesValidateUploadConan) GetConanPrefixOk() (*string, bool)`

GetConanPrefixOk returns a tuple with the ConanPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConanPrefix

`func (o *PackagesValidateUploadConan) SetConanPrefix(v string)`

SetConanPrefix sets ConanPrefix field to given value.

### HasConanPrefix

`func (o *PackagesValidateUploadConan) HasConanPrefix() bool`

HasConanPrefix returns a boolean if a field has been set.

### GetInfoFile

`func (o *PackagesValidateUploadConan) GetInfoFile() string`

GetInfoFile returns the InfoFile field if non-nil, zero value otherwise.

### GetInfoFileOk

`func (o *PackagesValidateUploadConan) GetInfoFileOk() (*string, bool)`

GetInfoFileOk returns a tuple with the InfoFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoFile

`func (o *PackagesValidateUploadConan) SetInfoFile(v string)`

SetInfoFile sets InfoFile field to given value.


### GetManifestFile

`func (o *PackagesValidateUploadConan) GetManifestFile() string`

GetManifestFile returns the ManifestFile field if non-nil, zero value otherwise.

### GetManifestFileOk

`func (o *PackagesValidateUploadConan) GetManifestFileOk() (*string, bool)`

GetManifestFileOk returns a tuple with the ManifestFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestFile

`func (o *PackagesValidateUploadConan) SetManifestFile(v string)`

SetManifestFile sets ManifestFile field to given value.


### GetMetadataFile

`func (o *PackagesValidateUploadConan) GetMetadataFile() string`

GetMetadataFile returns the MetadataFile field if non-nil, zero value otherwise.

### GetMetadataFileOk

`func (o *PackagesValidateUploadConan) GetMetadataFileOk() (*string, bool)`

GetMetadataFileOk returns a tuple with the MetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFile

`func (o *PackagesValidateUploadConan) SetMetadataFile(v string)`

SetMetadataFile sets MetadataFile field to given value.


### GetName

`func (o *PackagesValidateUploadConan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackagesValidateUploadConan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackagesValidateUploadConan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PackagesValidateUploadConan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackageFile

`func (o *PackagesValidateUploadConan) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *PackagesValidateUploadConan) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *PackagesValidateUploadConan) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRepublish

`func (o *PackagesValidateUploadConan) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *PackagesValidateUploadConan) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *PackagesValidateUploadConan) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *PackagesValidateUploadConan) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetTags

`func (o *PackagesValidateUploadConan) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PackagesValidateUploadConan) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PackagesValidateUploadConan) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PackagesValidateUploadConan) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *PackagesValidateUploadConan) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PackagesValidateUploadConan) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PackagesValidateUploadConan) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PackagesValidateUploadConan) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


