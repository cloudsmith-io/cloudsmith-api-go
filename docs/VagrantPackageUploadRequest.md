# VagrantPackageUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMalwareDetected** | Pointer to **bool** | Whether the package has been detected as containing malware. Requires Ultra plan. | [optional] [readonly] 
**Name** | **string** | The name of this package. | 
**PackageFile** | **string** | The primary file for the package. | 
**Provider** | **string** | The virtual machine provider for the box. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Tags** | Pointer to **NullableString** | A comma-separated values list of tags to add to the package. | [optional] 
**Version** | **string** | The raw version for this package. | 
**VulnerabilityCounts** | Pointer to [**NullableWebOSVSeverityCounts**](WebOSVSeverityCounts.md) |  | [optional] 

## Methods

### NewVagrantPackageUploadRequest

`func NewVagrantPackageUploadRequest(name string, packageFile string, provider string, version string, ) *VagrantPackageUploadRequest`

NewVagrantPackageUploadRequest instantiates a new VagrantPackageUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVagrantPackageUploadRequestWithDefaults

`func NewVagrantPackageUploadRequestWithDefaults() *VagrantPackageUploadRequest`

NewVagrantPackageUploadRequestWithDefaults instantiates a new VagrantPackageUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsMalwareDetected

`func (o *VagrantPackageUploadRequest) GetIsMalwareDetected() bool`

GetIsMalwareDetected returns the IsMalwareDetected field if non-nil, zero value otherwise.

### GetIsMalwareDetectedOk

`func (o *VagrantPackageUploadRequest) GetIsMalwareDetectedOk() (*bool, bool)`

GetIsMalwareDetectedOk returns a tuple with the IsMalwareDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMalwareDetected

`func (o *VagrantPackageUploadRequest) SetIsMalwareDetected(v bool)`

SetIsMalwareDetected sets IsMalwareDetected field to given value.

### HasIsMalwareDetected

`func (o *VagrantPackageUploadRequest) HasIsMalwareDetected() bool`

HasIsMalwareDetected returns a boolean if a field has been set.

### GetName

`func (o *VagrantPackageUploadRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VagrantPackageUploadRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VagrantPackageUploadRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPackageFile

`func (o *VagrantPackageUploadRequest) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *VagrantPackageUploadRequest) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *VagrantPackageUploadRequest) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetProvider

`func (o *VagrantPackageUploadRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VagrantPackageUploadRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VagrantPackageUploadRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRepublish

`func (o *VagrantPackageUploadRequest) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *VagrantPackageUploadRequest) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *VagrantPackageUploadRequest) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *VagrantPackageUploadRequest) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetTags

`func (o *VagrantPackageUploadRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VagrantPackageUploadRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VagrantPackageUploadRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VagrantPackageUploadRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *VagrantPackageUploadRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *VagrantPackageUploadRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetVersion

`func (o *VagrantPackageUploadRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VagrantPackageUploadRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VagrantPackageUploadRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVulnerabilityCounts

`func (o *VagrantPackageUploadRequest) GetVulnerabilityCounts() WebOSVSeverityCounts`

GetVulnerabilityCounts returns the VulnerabilityCounts field if non-nil, zero value otherwise.

### GetVulnerabilityCountsOk

`func (o *VagrantPackageUploadRequest) GetVulnerabilityCountsOk() (*WebOSVSeverityCounts, bool)`

GetVulnerabilityCountsOk returns a tuple with the VulnerabilityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityCounts

`func (o *VagrantPackageUploadRequest) SetVulnerabilityCounts(v WebOSVSeverityCounts)`

SetVulnerabilityCounts sets VulnerabilityCounts field to given value.

### HasVulnerabilityCounts

`func (o *VagrantPackageUploadRequest) HasVulnerabilityCounts() bool`

HasVulnerabilityCounts returns a boolean if a field has been set.

### SetVulnerabilityCountsNil

`func (o *VagrantPackageUploadRequest) SetVulnerabilityCountsNil(b bool)`

 SetVulnerabilityCountsNil sets the value for VulnerabilityCounts to be an explicit nil

### UnsetVulnerabilityCounts
`func (o *VagrantPackageUploadRequest) UnsetVulnerabilityCounts()`

UnsetVulnerabilityCounts ensures that no value is present for VulnerabilityCounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


