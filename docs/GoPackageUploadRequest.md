# GoPackageUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMalwareDetected** | Pointer to **bool** | Whether the package has been detected as containing malware. Requires Ultra plan. | [optional] [readonly] 
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Tags** | Pointer to **NullableString** | A comma-separated values list of tags to add to the package. | [optional] 
**VulnerabilityCounts** | Pointer to [**NullableWebOSVSeverityCounts**](WebOSVSeverityCounts.md) |  | [optional] 

## Methods

### NewGoPackageUploadRequest

`func NewGoPackageUploadRequest(packageFile string, ) *GoPackageUploadRequest`

NewGoPackageUploadRequest instantiates a new GoPackageUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoPackageUploadRequestWithDefaults

`func NewGoPackageUploadRequestWithDefaults() *GoPackageUploadRequest`

NewGoPackageUploadRequestWithDefaults instantiates a new GoPackageUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsMalwareDetected

`func (o *GoPackageUploadRequest) GetIsMalwareDetected() bool`

GetIsMalwareDetected returns the IsMalwareDetected field if non-nil, zero value otherwise.

### GetIsMalwareDetectedOk

`func (o *GoPackageUploadRequest) GetIsMalwareDetectedOk() (*bool, bool)`

GetIsMalwareDetectedOk returns a tuple with the IsMalwareDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMalwareDetected

`func (o *GoPackageUploadRequest) SetIsMalwareDetected(v bool)`

SetIsMalwareDetected sets IsMalwareDetected field to given value.

### HasIsMalwareDetected

`func (o *GoPackageUploadRequest) HasIsMalwareDetected() bool`

HasIsMalwareDetected returns a boolean if a field has been set.

### GetPackageFile

`func (o *GoPackageUploadRequest) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *GoPackageUploadRequest) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *GoPackageUploadRequest) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRepublish

`func (o *GoPackageUploadRequest) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *GoPackageUploadRequest) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *GoPackageUploadRequest) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *GoPackageUploadRequest) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetTags

`func (o *GoPackageUploadRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GoPackageUploadRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GoPackageUploadRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GoPackageUploadRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *GoPackageUploadRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *GoPackageUploadRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetVulnerabilityCounts

`func (o *GoPackageUploadRequest) GetVulnerabilityCounts() WebOSVSeverityCounts`

GetVulnerabilityCounts returns the VulnerabilityCounts field if non-nil, zero value otherwise.

### GetVulnerabilityCountsOk

`func (o *GoPackageUploadRequest) GetVulnerabilityCountsOk() (*WebOSVSeverityCounts, bool)`

GetVulnerabilityCountsOk returns a tuple with the VulnerabilityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityCounts

`func (o *GoPackageUploadRequest) SetVulnerabilityCounts(v WebOSVSeverityCounts)`

SetVulnerabilityCounts sets VulnerabilityCounts field to given value.

### HasVulnerabilityCounts

`func (o *GoPackageUploadRequest) HasVulnerabilityCounts() bool`

HasVulnerabilityCounts returns a boolean if a field has been set.

### SetVulnerabilityCountsNil

`func (o *GoPackageUploadRequest) SetVulnerabilityCountsNil(b bool)`

 SetVulnerabilityCountsNil sets the value for VulnerabilityCounts to be an explicit nil

### UnsetVulnerabilityCounts
`func (o *GoPackageUploadRequest) UnsetVulnerabilityCounts()`

UnsetVulnerabilityCounts ensures that no value is present for VulnerabilityCounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


