# PackagesUploadLuarocks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Tags** | Pointer to **string** | A comma-separated values list of tags to add to the package. | [optional] 

## Methods

### NewPackagesUploadLuarocks

`func NewPackagesUploadLuarocks(packageFile string, ) *PackagesUploadLuarocks`

NewPackagesUploadLuarocks instantiates a new PackagesUploadLuarocks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesUploadLuarocksWithDefaults

`func NewPackagesUploadLuarocksWithDefaults() *PackagesUploadLuarocks`

NewPackagesUploadLuarocksWithDefaults instantiates a new PackagesUploadLuarocks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageFile

`func (o *PackagesUploadLuarocks) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *PackagesUploadLuarocks) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *PackagesUploadLuarocks) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRepublish

`func (o *PackagesUploadLuarocks) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *PackagesUploadLuarocks) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *PackagesUploadLuarocks) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *PackagesUploadLuarocks) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetTags

`func (o *PackagesUploadLuarocks) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PackagesUploadLuarocks) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PackagesUploadLuarocks) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PackagesUploadLuarocks) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


