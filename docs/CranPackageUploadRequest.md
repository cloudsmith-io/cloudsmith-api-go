# CranPackageUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** | Binary package uploads for macOS should specify the architecture they were built for. | [optional] 
**PackageFile** | **string** | The primary file for the package. | 
**RVersion** | Pointer to **NullableString** | Binary package uploads should specify the version of R they were built for. | [optional] 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Tags** | Pointer to **NullableString** | A comma-separated values list of tags to add to the package. | [optional] 

## Methods

### NewCranPackageUploadRequest

`func NewCranPackageUploadRequest(packageFile string, ) *CranPackageUploadRequest`

NewCranPackageUploadRequest instantiates a new CranPackageUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCranPackageUploadRequestWithDefaults

`func NewCranPackageUploadRequestWithDefaults() *CranPackageUploadRequest`

NewCranPackageUploadRequestWithDefaults instantiates a new CranPackageUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *CranPackageUploadRequest) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *CranPackageUploadRequest) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *CranPackageUploadRequest) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *CranPackageUploadRequest) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetPackageFile

`func (o *CranPackageUploadRequest) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *CranPackageUploadRequest) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *CranPackageUploadRequest) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRVersion

`func (o *CranPackageUploadRequest) GetRVersion() string`

GetRVersion returns the RVersion field if non-nil, zero value otherwise.

### GetRVersionOk

`func (o *CranPackageUploadRequest) GetRVersionOk() (*string, bool)`

GetRVersionOk returns a tuple with the RVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRVersion

`func (o *CranPackageUploadRequest) SetRVersion(v string)`

SetRVersion sets RVersion field to given value.

### HasRVersion

`func (o *CranPackageUploadRequest) HasRVersion() bool`

HasRVersion returns a boolean if a field has been set.

### SetRVersionNil

`func (o *CranPackageUploadRequest) SetRVersionNil(b bool)`

 SetRVersionNil sets the value for RVersion to be an explicit nil

### UnsetRVersion
`func (o *CranPackageUploadRequest) UnsetRVersion()`

UnsetRVersion ensures that no value is present for RVersion, not even an explicit nil
### GetRepublish

`func (o *CranPackageUploadRequest) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *CranPackageUploadRequest) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *CranPackageUploadRequest) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *CranPackageUploadRequest) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetTags

`func (o *CranPackageUploadRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CranPackageUploadRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CranPackageUploadRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CranPackageUploadRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CranPackageUploadRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CranPackageUploadRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


