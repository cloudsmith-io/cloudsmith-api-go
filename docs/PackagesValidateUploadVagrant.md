# PackagesValidateUploadVagrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this package. | 
**PackageFile** | **string** | The primary file for the package. | 
**Provider** | **string** | The virtual machine provider for the box. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Tags** | Pointer to **string** | A comma-separated values list of tags to add to the package. | [optional] 
**Version** | **string** | The raw version for this package. | 

## Methods

### NewPackagesValidateUploadVagrant

`func NewPackagesValidateUploadVagrant(name string, packageFile string, provider string, version string, ) *PackagesValidateUploadVagrant`

NewPackagesValidateUploadVagrant instantiates a new PackagesValidateUploadVagrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesValidateUploadVagrantWithDefaults

`func NewPackagesValidateUploadVagrantWithDefaults() *PackagesValidateUploadVagrant`

NewPackagesValidateUploadVagrantWithDefaults instantiates a new PackagesValidateUploadVagrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PackagesValidateUploadVagrant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackagesValidateUploadVagrant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackagesValidateUploadVagrant) SetName(v string)`

SetName sets Name field to given value.


### GetPackageFile

`func (o *PackagesValidateUploadVagrant) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *PackagesValidateUploadVagrant) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *PackagesValidateUploadVagrant) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetProvider

`func (o *PackagesValidateUploadVagrant) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PackagesValidateUploadVagrant) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PackagesValidateUploadVagrant) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRepublish

`func (o *PackagesValidateUploadVagrant) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *PackagesValidateUploadVagrant) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *PackagesValidateUploadVagrant) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *PackagesValidateUploadVagrant) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetTags

`func (o *PackagesValidateUploadVagrant) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PackagesValidateUploadVagrant) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PackagesValidateUploadVagrant) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PackagesValidateUploadVagrant) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *PackagesValidateUploadVagrant) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PackagesValidateUploadVagrant) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PackagesValidateUploadVagrant) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


