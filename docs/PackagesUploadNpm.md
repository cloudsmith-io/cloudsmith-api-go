# PackagesUploadNpm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NpmDistTag** | Pointer to **string** | The default npm dist-tag for this package/version - This will replace any other package/version if they are using the same tag. | [optional] 
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Tags** | Pointer to **string** | A comma-separated values list of tags to add to the package. | [optional] 

## Methods

### NewPackagesUploadNpm

`func NewPackagesUploadNpm(packageFile string, ) *PackagesUploadNpm`

NewPackagesUploadNpm instantiates a new PackagesUploadNpm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesUploadNpmWithDefaults

`func NewPackagesUploadNpmWithDefaults() *PackagesUploadNpm`

NewPackagesUploadNpmWithDefaults instantiates a new PackagesUploadNpm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNpmDistTag

`func (o *PackagesUploadNpm) GetNpmDistTag() string`

GetNpmDistTag returns the NpmDistTag field if non-nil, zero value otherwise.

### GetNpmDistTagOk

`func (o *PackagesUploadNpm) GetNpmDistTagOk() (*string, bool)`

GetNpmDistTagOk returns a tuple with the NpmDistTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmDistTag

`func (o *PackagesUploadNpm) SetNpmDistTag(v string)`

SetNpmDistTag sets NpmDistTag field to given value.

### HasNpmDistTag

`func (o *PackagesUploadNpm) HasNpmDistTag() bool`

HasNpmDistTag returns a boolean if a field has been set.

### GetPackageFile

`func (o *PackagesUploadNpm) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *PackagesUploadNpm) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *PackagesUploadNpm) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRepublish

`func (o *PackagesUploadNpm) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *PackagesUploadNpm) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *PackagesUploadNpm) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *PackagesUploadNpm) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetTags

`func (o *PackagesUploadNpm) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PackagesUploadNpm) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PackagesUploadNpm) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PackagesUploadNpm) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


