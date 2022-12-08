# NpmPackageUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NpmDistTag** | Pointer to **string** | The default npm dist-tag for this package/version - This will replace any other package/version if they are using the same tag. | [optional] [default to "latest"]
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Tags** | Pointer to **NullableString** | A comma-separated values list of tags to add to the package. | [optional] 

## Methods

### NewNpmPackageUploadRequest

`func NewNpmPackageUploadRequest(packageFile string, ) *NpmPackageUploadRequest`

NewNpmPackageUploadRequest instantiates a new NpmPackageUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNpmPackageUploadRequestWithDefaults

`func NewNpmPackageUploadRequestWithDefaults() *NpmPackageUploadRequest`

NewNpmPackageUploadRequestWithDefaults instantiates a new NpmPackageUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNpmDistTag

`func (o *NpmPackageUploadRequest) GetNpmDistTag() string`

GetNpmDistTag returns the NpmDistTag field if non-nil, zero value otherwise.

### GetNpmDistTagOk

`func (o *NpmPackageUploadRequest) GetNpmDistTagOk() (*string, bool)`

GetNpmDistTagOk returns a tuple with the NpmDistTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmDistTag

`func (o *NpmPackageUploadRequest) SetNpmDistTag(v string)`

SetNpmDistTag sets NpmDistTag field to given value.

### HasNpmDistTag

`func (o *NpmPackageUploadRequest) HasNpmDistTag() bool`

HasNpmDistTag returns a boolean if a field has been set.

### GetPackageFile

`func (o *NpmPackageUploadRequest) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *NpmPackageUploadRequest) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *NpmPackageUploadRequest) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRepublish

`func (o *NpmPackageUploadRequest) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *NpmPackageUploadRequest) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *NpmPackageUploadRequest) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *NpmPackageUploadRequest) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetTags

`func (o *NpmPackageUploadRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NpmPackageUploadRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NpmPackageUploadRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NpmPackageUploadRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *NpmPackageUploadRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *NpmPackageUploadRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


