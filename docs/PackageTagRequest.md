# PackageTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** |  | [optional] [default to "Add"]
**IsImmutable** | Pointer to **bool** | If true, created tags will be immutable. An immutable flag is a tag that cannot be removed from a package. | [optional] [default to false]
**IsMalwareDetected** | Pointer to **bool** | Whether the package has been detected as containing malware. Requires Ultra plan. | [optional] [readonly] 
**Tags** | Pointer to **[]string** | A list of tags to apply the action to. Not required for clears. | [optional] 
**VulnerabilityCounts** | Pointer to [**NullableWebOSVSeverityCounts**](WebOSVSeverityCounts.md) |  | [optional] 

## Methods

### NewPackageTagRequest

`func NewPackageTagRequest() *PackageTagRequest`

NewPackageTagRequest instantiates a new PackageTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageTagRequestWithDefaults

`func NewPackageTagRequestWithDefaults() *PackageTagRequest`

NewPackageTagRequestWithDefaults instantiates a new PackageTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PackageTagRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PackageTagRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PackageTagRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PackageTagRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *PackageTagRequest) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *PackageTagRequest) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetIsImmutable

`func (o *PackageTagRequest) GetIsImmutable() bool`

GetIsImmutable returns the IsImmutable field if non-nil, zero value otherwise.

### GetIsImmutableOk

`func (o *PackageTagRequest) GetIsImmutableOk() (*bool, bool)`

GetIsImmutableOk returns a tuple with the IsImmutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImmutable

`func (o *PackageTagRequest) SetIsImmutable(v bool)`

SetIsImmutable sets IsImmutable field to given value.

### HasIsImmutable

`func (o *PackageTagRequest) HasIsImmutable() bool`

HasIsImmutable returns a boolean if a field has been set.

### GetIsMalwareDetected

`func (o *PackageTagRequest) GetIsMalwareDetected() bool`

GetIsMalwareDetected returns the IsMalwareDetected field if non-nil, zero value otherwise.

### GetIsMalwareDetectedOk

`func (o *PackageTagRequest) GetIsMalwareDetectedOk() (*bool, bool)`

GetIsMalwareDetectedOk returns a tuple with the IsMalwareDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMalwareDetected

`func (o *PackageTagRequest) SetIsMalwareDetected(v bool)`

SetIsMalwareDetected sets IsMalwareDetected field to given value.

### HasIsMalwareDetected

`func (o *PackageTagRequest) HasIsMalwareDetected() bool`

HasIsMalwareDetected returns a boolean if a field has been set.

### GetTags

`func (o *PackageTagRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PackageTagRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PackageTagRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PackageTagRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *PackageTagRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *PackageTagRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetVulnerabilityCounts

`func (o *PackageTagRequest) GetVulnerabilityCounts() WebOSVSeverityCounts`

GetVulnerabilityCounts returns the VulnerabilityCounts field if non-nil, zero value otherwise.

### GetVulnerabilityCountsOk

`func (o *PackageTagRequest) GetVulnerabilityCountsOk() (*WebOSVSeverityCounts, bool)`

GetVulnerabilityCountsOk returns a tuple with the VulnerabilityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityCounts

`func (o *PackageTagRequest) SetVulnerabilityCounts(v WebOSVSeverityCounts)`

SetVulnerabilityCounts sets VulnerabilityCounts field to given value.

### HasVulnerabilityCounts

`func (o *PackageTagRequest) HasVulnerabilityCounts() bool`

HasVulnerabilityCounts returns a boolean if a field has been set.

### SetVulnerabilityCountsNil

`func (o *PackageTagRequest) SetVulnerabilityCountsNil(b bool)`

 SetVulnerabilityCountsNil sets the value for VulnerabilityCounts to be an explicit nil

### UnsetVulnerabilityCounts
`func (o *PackageTagRequest) UnsetVulnerabilityCounts()`

UnsetVulnerabilityCounts ensures that no value is present for VulnerabilityCounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


