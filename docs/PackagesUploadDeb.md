# PackagesUploadDeb

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

### NewPackagesUploadDeb

`func NewPackagesUploadDeb(distribution string, packageFile string, ) *PackagesUploadDeb`

NewPackagesUploadDeb instantiates a new PackagesUploadDeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesUploadDebWithDefaults

`func NewPackagesUploadDebWithDefaults() *PackagesUploadDeb`

NewPackagesUploadDebWithDefaults instantiates a new PackagesUploadDeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangesFile

`func (o *PackagesUploadDeb) GetChangesFile() string`

GetChangesFile returns the ChangesFile field if non-nil, zero value otherwise.

### GetChangesFileOk

`func (o *PackagesUploadDeb) GetChangesFileOk() (*string, bool)`

GetChangesFileOk returns a tuple with the ChangesFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangesFile

`func (o *PackagesUploadDeb) SetChangesFile(v string)`

SetChangesFile sets ChangesFile field to given value.

### HasChangesFile

`func (o *PackagesUploadDeb) HasChangesFile() bool`

HasChangesFile returns a boolean if a field has been set.

### GetDistribution

`func (o *PackagesUploadDeb) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *PackagesUploadDeb) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *PackagesUploadDeb) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.


### GetPackageFile

`func (o *PackagesUploadDeb) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *PackagesUploadDeb) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *PackagesUploadDeb) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRepublish

`func (o *PackagesUploadDeb) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *PackagesUploadDeb) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *PackagesUploadDeb) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *PackagesUploadDeb) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetSourcesFile

`func (o *PackagesUploadDeb) GetSourcesFile() string`

GetSourcesFile returns the SourcesFile field if non-nil, zero value otherwise.

### GetSourcesFileOk

`func (o *PackagesUploadDeb) GetSourcesFileOk() (*string, bool)`

GetSourcesFileOk returns a tuple with the SourcesFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesFile

`func (o *PackagesUploadDeb) SetSourcesFile(v string)`

SetSourcesFile sets SourcesFile field to given value.

### HasSourcesFile

`func (o *PackagesUploadDeb) HasSourcesFile() bool`

HasSourcesFile returns a boolean if a field has been set.

### GetTags

`func (o *PackagesUploadDeb) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PackagesUploadDeb) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PackagesUploadDeb) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PackagesUploadDeb) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


