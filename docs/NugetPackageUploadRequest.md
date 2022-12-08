# NugetPackageUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**SymbolsFile** | Pointer to **NullableString** | Attaches a symbols file to the package. | [optional] 
**Tags** | Pointer to **NullableString** | A comma-separated values list of tags to add to the package. | [optional] 

## Methods

### NewNugetPackageUploadRequest

`func NewNugetPackageUploadRequest(packageFile string, ) *NugetPackageUploadRequest`

NewNugetPackageUploadRequest instantiates a new NugetPackageUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNugetPackageUploadRequestWithDefaults

`func NewNugetPackageUploadRequestWithDefaults() *NugetPackageUploadRequest`

NewNugetPackageUploadRequestWithDefaults instantiates a new NugetPackageUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageFile

`func (o *NugetPackageUploadRequest) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *NugetPackageUploadRequest) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *NugetPackageUploadRequest) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRepublish

`func (o *NugetPackageUploadRequest) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *NugetPackageUploadRequest) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *NugetPackageUploadRequest) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *NugetPackageUploadRequest) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetSymbolsFile

`func (o *NugetPackageUploadRequest) GetSymbolsFile() string`

GetSymbolsFile returns the SymbolsFile field if non-nil, zero value otherwise.

### GetSymbolsFileOk

`func (o *NugetPackageUploadRequest) GetSymbolsFileOk() (*string, bool)`

GetSymbolsFileOk returns a tuple with the SymbolsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolsFile

`func (o *NugetPackageUploadRequest) SetSymbolsFile(v string)`

SetSymbolsFile sets SymbolsFile field to given value.

### HasSymbolsFile

`func (o *NugetPackageUploadRequest) HasSymbolsFile() bool`

HasSymbolsFile returns a boolean if a field has been set.

### SetSymbolsFileNil

`func (o *NugetPackageUploadRequest) SetSymbolsFileNil(b bool)`

 SetSymbolsFileNil sets the value for SymbolsFile to be an explicit nil

### UnsetSymbolsFile
`func (o *NugetPackageUploadRequest) UnsetSymbolsFile()`

UnsetSymbolsFile ensures that no value is present for SymbolsFile, not even an explicit nil
### GetTags

`func (o *NugetPackageUploadRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NugetPackageUploadRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NugetPackageUploadRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NugetPackageUploadRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *NugetPackageUploadRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *NugetPackageUploadRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


