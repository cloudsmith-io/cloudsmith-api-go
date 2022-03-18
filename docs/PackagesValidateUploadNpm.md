# PackagesValidateUploadNpm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NpmDistTag** | Pointer to **string** | The default npm dist-tag for this package/version - This will replace any other package/version if they are using the same tag. | [optional] 
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Tags** | Pointer to **string** | A comma-separated values list of tags to add to the package. | [optional] 

## Methods

### NewPackagesValidateUploadNpm

`func NewPackagesValidateUploadNpm(packageFile string, ) *PackagesValidateUploadNpm`

NewPackagesValidateUploadNpm instantiates a new PackagesValidateUploadNpm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesValidateUploadNpmWithDefaults

`func NewPackagesValidateUploadNpmWithDefaults() *PackagesValidateUploadNpm`

NewPackagesValidateUploadNpmWithDefaults instantiates a new PackagesValidateUploadNpm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNpmDistTag

`func (o *PackagesValidateUploadNpm) GetNpmDistTag() string`

GetNpmDistTag returns the NpmDistTag field if non-nil, zero value otherwise.

### GetNpmDistTagOk

`func (o *PackagesValidateUploadNpm) GetNpmDistTagOk() (*string, bool)`

GetNpmDistTagOk returns a tuple with the NpmDistTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmDistTag

`func (o *PackagesValidateUploadNpm) SetNpmDistTag(v string)`

SetNpmDistTag sets NpmDistTag field to given value.

### HasNpmDistTag

`func (o *PackagesValidateUploadNpm) HasNpmDistTag() bool`

HasNpmDistTag returns a boolean if a field has been set.

### GetPackageFile

`func (o *PackagesValidateUploadNpm) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *PackagesValidateUploadNpm) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *PackagesValidateUploadNpm) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRepublish

`func (o *PackagesValidateUploadNpm) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *PackagesValidateUploadNpm) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *PackagesValidateUploadNpm) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *PackagesValidateUploadNpm) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetTags

`func (o *PackagesValidateUploadNpm) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PackagesValidateUploadNpm) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PackagesValidateUploadNpm) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PackagesValidateUploadNpm) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


