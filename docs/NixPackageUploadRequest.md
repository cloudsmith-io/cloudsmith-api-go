# NixPackageUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NarinfoFile** | Pointer to **NullableString** | The narinfo sidecar metadata file (&lt;storeHash&gt;.narinfo). Optional — the package can be completed later by publishing the matching narinfo via the native HTTP PUT path. | [optional] 
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Tags** | Pointer to **NullableString** | A comma-separated values list of tags to add to the package. | [optional] 
**Version** | Pointer to **NullableString** | The raw version for this package. | [optional] 

## Methods

### NewNixPackageUploadRequest

`func NewNixPackageUploadRequest(packageFile string, ) *NixPackageUploadRequest`

NewNixPackageUploadRequest instantiates a new NixPackageUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNixPackageUploadRequestWithDefaults

`func NewNixPackageUploadRequestWithDefaults() *NixPackageUploadRequest`

NewNixPackageUploadRequestWithDefaults instantiates a new NixPackageUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNarinfoFile

`func (o *NixPackageUploadRequest) GetNarinfoFile() string`

GetNarinfoFile returns the NarinfoFile field if non-nil, zero value otherwise.

### GetNarinfoFileOk

`func (o *NixPackageUploadRequest) GetNarinfoFileOk() (*string, bool)`

GetNarinfoFileOk returns a tuple with the NarinfoFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarinfoFile

`func (o *NixPackageUploadRequest) SetNarinfoFile(v string)`

SetNarinfoFile sets NarinfoFile field to given value.

### HasNarinfoFile

`func (o *NixPackageUploadRequest) HasNarinfoFile() bool`

HasNarinfoFile returns a boolean if a field has been set.

### SetNarinfoFileNil

`func (o *NixPackageUploadRequest) SetNarinfoFileNil(b bool)`

 SetNarinfoFileNil sets the value for NarinfoFile to be an explicit nil

### UnsetNarinfoFile
`func (o *NixPackageUploadRequest) UnsetNarinfoFile()`

UnsetNarinfoFile ensures that no value is present for NarinfoFile, not even an explicit nil
### GetPackageFile

`func (o *NixPackageUploadRequest) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *NixPackageUploadRequest) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *NixPackageUploadRequest) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRepublish

`func (o *NixPackageUploadRequest) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *NixPackageUploadRequest) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *NixPackageUploadRequest) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *NixPackageUploadRequest) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetTags

`func (o *NixPackageUploadRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NixPackageUploadRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NixPackageUploadRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NixPackageUploadRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *NixPackageUploadRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *NixPackageUploadRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetVersion

`func (o *NixPackageUploadRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NixPackageUploadRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NixPackageUploadRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NixPackageUploadRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *NixPackageUploadRequest) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *NixPackageUploadRequest) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


