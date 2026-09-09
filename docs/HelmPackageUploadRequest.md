# HelmPackageUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageFile** | **string** | The primary file for the package. | 
**ProvenanceFile** | Pointer to **NullableString** | The provenance file containing the signature for the chart. If one is not provided, it will be generated automatically. | [optional] 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Tags** | Pointer to **NullableString** | A comma-separated values list of tags to add to the package. | [optional] 

## Methods

### NewHelmPackageUploadRequest

`func NewHelmPackageUploadRequest(packageFile string, ) *HelmPackageUploadRequest`

NewHelmPackageUploadRequest instantiates a new HelmPackageUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmPackageUploadRequestWithDefaults

`func NewHelmPackageUploadRequestWithDefaults() *HelmPackageUploadRequest`

NewHelmPackageUploadRequestWithDefaults instantiates a new HelmPackageUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageFile

`func (o *HelmPackageUploadRequest) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *HelmPackageUploadRequest) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *HelmPackageUploadRequest) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetProvenanceFile

`func (o *HelmPackageUploadRequest) GetProvenanceFile() string`

GetProvenanceFile returns the ProvenanceFile field if non-nil, zero value otherwise.

### GetProvenanceFileOk

`func (o *HelmPackageUploadRequest) GetProvenanceFileOk() (*string, bool)`

GetProvenanceFileOk returns a tuple with the ProvenanceFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenanceFile

`func (o *HelmPackageUploadRequest) SetProvenanceFile(v string)`

SetProvenanceFile sets ProvenanceFile field to given value.

### HasProvenanceFile

`func (o *HelmPackageUploadRequest) HasProvenanceFile() bool`

HasProvenanceFile returns a boolean if a field has been set.

### SetProvenanceFileNil

`func (o *HelmPackageUploadRequest) SetProvenanceFileNil(b bool)`

 SetProvenanceFileNil sets the value for ProvenanceFile to be an explicit nil

### UnsetProvenanceFile
`func (o *HelmPackageUploadRequest) UnsetProvenanceFile()`

UnsetProvenanceFile ensures that no value is present for ProvenanceFile, not even an explicit nil
### GetRepublish

`func (o *HelmPackageUploadRequest) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *HelmPackageUploadRequest) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *HelmPackageUploadRequest) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *HelmPackageUploadRequest) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetTags

`func (o *HelmPackageUploadRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HelmPackageUploadRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HelmPackageUploadRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HelmPackageUploadRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *HelmPackageUploadRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *HelmPackageUploadRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


