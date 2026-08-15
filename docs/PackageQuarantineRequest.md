# PackageQuarantineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMalwareDetected** | Pointer to **bool** | Whether the package has been detected as containing malware. Requires Ultra plan. | [optional] [readonly] 
**Release** | Pointer to **bool** | If true, the package is released from quarantine. | [optional] 
**Restore** | Pointer to **bool** | If true, the package is released from quarantine. Note: This field is deprecated, please use &#39;release&#39; instead. | [optional] 
**VulnerabilityCounts** | Pointer to [**NullableWebOSVSeverityCounts**](WebOSVSeverityCounts.md) |  | [optional] 

## Methods

### NewPackageQuarantineRequest

`func NewPackageQuarantineRequest() *PackageQuarantineRequest`

NewPackageQuarantineRequest instantiates a new PackageQuarantineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageQuarantineRequestWithDefaults

`func NewPackageQuarantineRequestWithDefaults() *PackageQuarantineRequest`

NewPackageQuarantineRequestWithDefaults instantiates a new PackageQuarantineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsMalwareDetected

`func (o *PackageQuarantineRequest) GetIsMalwareDetected() bool`

GetIsMalwareDetected returns the IsMalwareDetected field if non-nil, zero value otherwise.

### GetIsMalwareDetectedOk

`func (o *PackageQuarantineRequest) GetIsMalwareDetectedOk() (*bool, bool)`

GetIsMalwareDetectedOk returns a tuple with the IsMalwareDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMalwareDetected

`func (o *PackageQuarantineRequest) SetIsMalwareDetected(v bool)`

SetIsMalwareDetected sets IsMalwareDetected field to given value.

### HasIsMalwareDetected

`func (o *PackageQuarantineRequest) HasIsMalwareDetected() bool`

HasIsMalwareDetected returns a boolean if a field has been set.

### GetRelease

`func (o *PackageQuarantineRequest) GetRelease() bool`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *PackageQuarantineRequest) GetReleaseOk() (*bool, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *PackageQuarantineRequest) SetRelease(v bool)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *PackageQuarantineRequest) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetRestore

`func (o *PackageQuarantineRequest) GetRestore() bool`

GetRestore returns the Restore field if non-nil, zero value otherwise.

### GetRestoreOk

`func (o *PackageQuarantineRequest) GetRestoreOk() (*bool, bool)`

GetRestoreOk returns a tuple with the Restore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestore

`func (o *PackageQuarantineRequest) SetRestore(v bool)`

SetRestore sets Restore field to given value.

### HasRestore

`func (o *PackageQuarantineRequest) HasRestore() bool`

HasRestore returns a boolean if a field has been set.

### GetVulnerabilityCounts

`func (o *PackageQuarantineRequest) GetVulnerabilityCounts() WebOSVSeverityCounts`

GetVulnerabilityCounts returns the VulnerabilityCounts field if non-nil, zero value otherwise.

### GetVulnerabilityCountsOk

`func (o *PackageQuarantineRequest) GetVulnerabilityCountsOk() (*WebOSVSeverityCounts, bool)`

GetVulnerabilityCountsOk returns a tuple with the VulnerabilityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityCounts

`func (o *PackageQuarantineRequest) SetVulnerabilityCounts(v WebOSVSeverityCounts)`

SetVulnerabilityCounts sets VulnerabilityCounts field to given value.

### HasVulnerabilityCounts

`func (o *PackageQuarantineRequest) HasVulnerabilityCounts() bool`

HasVulnerabilityCounts returns a boolean if a field has been set.

### SetVulnerabilityCountsNil

`func (o *PackageQuarantineRequest) SetVulnerabilityCountsNil(b bool)`

 SetVulnerabilityCountsNil sets the value for VulnerabilityCounts to be an explicit nil

### UnsetVulnerabilityCounts
`func (o *PackageQuarantineRequest) UnsetVulnerabilityCounts()`

UnsetVulnerabilityCounts ensures that no value is present for VulnerabilityCounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


