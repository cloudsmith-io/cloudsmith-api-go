# NpmPackageUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMalwareDetected** | Pointer to **bool** | Whether the package has been detected as containing malware. Requires Ultra plan. | [optional] [readonly] 
**NpmDistTag** | Pointer to **NullableString** | The default npm dist-tag for this package/version - This will replace any other package/version if they are using the same tag. | [optional] [default to "latest"]
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Tags** | Pointer to **NullableString** | A comma-separated values list of tags to add to the package. | [optional] 
**VulnerabilityCounts** | Pointer to [**NullableWebOSVSeverityCounts**](WebOSVSeverityCounts.md) |  | [optional] 

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

### GetIsMalwareDetected

`func (o *NpmPackageUploadRequest) GetIsMalwareDetected() bool`

GetIsMalwareDetected returns the IsMalwareDetected field if non-nil, zero value otherwise.

### GetIsMalwareDetectedOk

`func (o *NpmPackageUploadRequest) GetIsMalwareDetectedOk() (*bool, bool)`

GetIsMalwareDetectedOk returns a tuple with the IsMalwareDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMalwareDetected

`func (o *NpmPackageUploadRequest) SetIsMalwareDetected(v bool)`

SetIsMalwareDetected sets IsMalwareDetected field to given value.

### HasIsMalwareDetected

`func (o *NpmPackageUploadRequest) HasIsMalwareDetected() bool`

HasIsMalwareDetected returns a boolean if a field has been set.

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

### SetNpmDistTagNil

`func (o *NpmPackageUploadRequest) SetNpmDistTagNil(b bool)`

 SetNpmDistTagNil sets the value for NpmDistTag to be an explicit nil

### UnsetNpmDistTag
`func (o *NpmPackageUploadRequest) UnsetNpmDistTag()`

UnsetNpmDistTag ensures that no value is present for NpmDistTag, not even an explicit nil
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
### GetVulnerabilityCounts

`func (o *NpmPackageUploadRequest) GetVulnerabilityCounts() WebOSVSeverityCounts`

GetVulnerabilityCounts returns the VulnerabilityCounts field if non-nil, zero value otherwise.

### GetVulnerabilityCountsOk

`func (o *NpmPackageUploadRequest) GetVulnerabilityCountsOk() (*WebOSVSeverityCounts, bool)`

GetVulnerabilityCountsOk returns a tuple with the VulnerabilityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityCounts

`func (o *NpmPackageUploadRequest) SetVulnerabilityCounts(v WebOSVSeverityCounts)`

SetVulnerabilityCounts sets VulnerabilityCounts field to given value.

### HasVulnerabilityCounts

`func (o *NpmPackageUploadRequest) HasVulnerabilityCounts() bool`

HasVulnerabilityCounts returns a boolean if a field has been set.

### SetVulnerabilityCountsNil

`func (o *NpmPackageUploadRequest) SetVulnerabilityCountsNil(b bool)`

 SetVulnerabilityCountsNil sets the value for VulnerabilityCounts to be an explicit nil

### UnsetVulnerabilityCounts
`func (o *NpmPackageUploadRequest) UnsetVulnerabilityCounts()`

UnsetVulnerabilityCounts ensures that no value is present for VulnerabilityCounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


