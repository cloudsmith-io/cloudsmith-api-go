# PackagesUploadVagrant

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

### NewPackagesUploadVagrant

`func NewPackagesUploadVagrant(name string, packageFile string, provider string, version string, ) *PackagesUploadVagrant`

NewPackagesUploadVagrant instantiates a new PackagesUploadVagrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesUploadVagrantWithDefaults

`func NewPackagesUploadVagrantWithDefaults() *PackagesUploadVagrant`

NewPackagesUploadVagrantWithDefaults instantiates a new PackagesUploadVagrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PackagesUploadVagrant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackagesUploadVagrant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackagesUploadVagrant) SetName(v string)`

SetName sets Name field to given value.


### GetPackageFile

`func (o *PackagesUploadVagrant) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *PackagesUploadVagrant) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *PackagesUploadVagrant) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetProvider

`func (o *PackagesUploadVagrant) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PackagesUploadVagrant) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PackagesUploadVagrant) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRepublish

`func (o *PackagesUploadVagrant) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *PackagesUploadVagrant) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *PackagesUploadVagrant) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *PackagesUploadVagrant) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetTags

`func (o *PackagesUploadVagrant) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PackagesUploadVagrant) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PackagesUploadVagrant) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PackagesUploadVagrant) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *PackagesUploadVagrant) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PackagesUploadVagrant) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PackagesUploadVagrant) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


