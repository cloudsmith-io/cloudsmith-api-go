# PackagesUploadAlpine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distribution** | **string** | The distribution to store the package for. | 
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Tags** | Pointer to **string** | A comma-separated values list of tags to add to the package. | [optional] 

## Methods

### NewPackagesUploadAlpine

`func NewPackagesUploadAlpine(distribution string, packageFile string, ) *PackagesUploadAlpine`

NewPackagesUploadAlpine instantiates a new PackagesUploadAlpine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesUploadAlpineWithDefaults

`func NewPackagesUploadAlpineWithDefaults() *PackagesUploadAlpine`

NewPackagesUploadAlpineWithDefaults instantiates a new PackagesUploadAlpine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistribution

`func (o *PackagesUploadAlpine) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *PackagesUploadAlpine) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *PackagesUploadAlpine) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.


### GetPackageFile

`func (o *PackagesUploadAlpine) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *PackagesUploadAlpine) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *PackagesUploadAlpine) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRepublish

`func (o *PackagesUploadAlpine) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *PackagesUploadAlpine) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *PackagesUploadAlpine) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *PackagesUploadAlpine) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetTags

`func (o *PackagesUploadAlpine) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PackagesUploadAlpine) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PackagesUploadAlpine) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PackagesUploadAlpine) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


