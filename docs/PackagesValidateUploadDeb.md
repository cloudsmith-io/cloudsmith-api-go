# PackagesValidateUploadDeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangesFile** | Pointer to **string** | The changes archive containing the changes made to the source and debian packaging files | [optional] 
**Distribution** | **string** | The distribution to store the package for. | 
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**SourcesFile** | Pointer to **string** | The sources archive containing the source code for the binary | [optional] 
**Tags** | Pointer to **string** | A comma-separated values list of tags to add to the package. | [optional] 

## Methods

### NewPackagesValidateUploadDeb

`func NewPackagesValidateUploadDeb(distribution string, packageFile string, ) *PackagesValidateUploadDeb`

NewPackagesValidateUploadDeb instantiates a new PackagesValidateUploadDeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesValidateUploadDebWithDefaults

`func NewPackagesValidateUploadDebWithDefaults() *PackagesValidateUploadDeb`

NewPackagesValidateUploadDebWithDefaults instantiates a new PackagesValidateUploadDeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangesFile

`func (o *PackagesValidateUploadDeb) GetChangesFile() string`

GetChangesFile returns the ChangesFile field if non-nil, zero value otherwise.

### GetChangesFileOk

`func (o *PackagesValidateUploadDeb) GetChangesFileOk() (*string, bool)`

GetChangesFileOk returns a tuple with the ChangesFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangesFile

`func (o *PackagesValidateUploadDeb) SetChangesFile(v string)`

SetChangesFile sets ChangesFile field to given value.

### HasChangesFile

`func (o *PackagesValidateUploadDeb) HasChangesFile() bool`

HasChangesFile returns a boolean if a field has been set.

### GetDistribution

`func (o *PackagesValidateUploadDeb) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *PackagesValidateUploadDeb) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *PackagesValidateUploadDeb) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.


### GetPackageFile

`func (o *PackagesValidateUploadDeb) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *PackagesValidateUploadDeb) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *PackagesValidateUploadDeb) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRepublish

`func (o *PackagesValidateUploadDeb) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *PackagesValidateUploadDeb) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *PackagesValidateUploadDeb) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *PackagesValidateUploadDeb) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetSourcesFile

`func (o *PackagesValidateUploadDeb) GetSourcesFile() string`

GetSourcesFile returns the SourcesFile field if non-nil, zero value otherwise.

### GetSourcesFileOk

`func (o *PackagesValidateUploadDeb) GetSourcesFileOk() (*string, bool)`

GetSourcesFileOk returns a tuple with the SourcesFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesFile

`func (o *PackagesValidateUploadDeb) SetSourcesFile(v string)`

SetSourcesFile sets SourcesFile field to given value.

### HasSourcesFile

`func (o *PackagesValidateUploadDeb) HasSourcesFile() bool`

HasSourcesFile returns a boolean if a field has been set.

### GetTags

`func (o *PackagesValidateUploadDeb) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PackagesValidateUploadDeb) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PackagesValidateUploadDeb) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PackagesValidateUploadDeb) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


