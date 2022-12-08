# DebPackageUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangesFile** | Pointer to **NullableString** | The changes archive containing the changes made to the source and debian packaging files | [optional] 
**Distribution** | **string** | The distribution to store the package for. | 
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**SourcesFile** | Pointer to **NullableString** | The sources archive containing the source code for the binary | [optional] 
**Tags** | Pointer to **NullableString** | A comma-separated values list of tags to add to the package. | [optional] 

## Methods

### NewDebPackageUploadRequest

`func NewDebPackageUploadRequest(distribution string, packageFile string, ) *DebPackageUploadRequest`

NewDebPackageUploadRequest instantiates a new DebPackageUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebPackageUploadRequestWithDefaults

`func NewDebPackageUploadRequestWithDefaults() *DebPackageUploadRequest`

NewDebPackageUploadRequestWithDefaults instantiates a new DebPackageUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangesFile

`func (o *DebPackageUploadRequest) GetChangesFile() string`

GetChangesFile returns the ChangesFile field if non-nil, zero value otherwise.

### GetChangesFileOk

`func (o *DebPackageUploadRequest) GetChangesFileOk() (*string, bool)`

GetChangesFileOk returns a tuple with the ChangesFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangesFile

`func (o *DebPackageUploadRequest) SetChangesFile(v string)`

SetChangesFile sets ChangesFile field to given value.

### HasChangesFile

`func (o *DebPackageUploadRequest) HasChangesFile() bool`

HasChangesFile returns a boolean if a field has been set.

### SetChangesFileNil

`func (o *DebPackageUploadRequest) SetChangesFileNil(b bool)`

 SetChangesFileNil sets the value for ChangesFile to be an explicit nil

### UnsetChangesFile
`func (o *DebPackageUploadRequest) UnsetChangesFile()`

UnsetChangesFile ensures that no value is present for ChangesFile, not even an explicit nil
### GetDistribution

`func (o *DebPackageUploadRequest) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *DebPackageUploadRequest) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *DebPackageUploadRequest) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.


### GetPackageFile

`func (o *DebPackageUploadRequest) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *DebPackageUploadRequest) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *DebPackageUploadRequest) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRepublish

`func (o *DebPackageUploadRequest) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *DebPackageUploadRequest) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *DebPackageUploadRequest) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *DebPackageUploadRequest) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetSourcesFile

`func (o *DebPackageUploadRequest) GetSourcesFile() string`

GetSourcesFile returns the SourcesFile field if non-nil, zero value otherwise.

### GetSourcesFileOk

`func (o *DebPackageUploadRequest) GetSourcesFileOk() (*string, bool)`

GetSourcesFileOk returns a tuple with the SourcesFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesFile

`func (o *DebPackageUploadRequest) SetSourcesFile(v string)`

SetSourcesFile sets SourcesFile field to given value.

### HasSourcesFile

`func (o *DebPackageUploadRequest) HasSourcesFile() bool`

HasSourcesFile returns a boolean if a field has been set.

### SetSourcesFileNil

`func (o *DebPackageUploadRequest) SetSourcesFileNil(b bool)`

 SetSourcesFileNil sets the value for SourcesFile to be an explicit nil

### UnsetSourcesFile
`func (o *DebPackageUploadRequest) UnsetSourcesFile()`

UnsetSourcesFile ensures that no value is present for SourcesFile, not even an explicit nil
### GetTags

`func (o *DebPackageUploadRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DebPackageUploadRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DebPackageUploadRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DebPackageUploadRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DebPackageUploadRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DebPackageUploadRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


