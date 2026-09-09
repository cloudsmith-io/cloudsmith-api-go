# PackageLicenseRequestPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** |  | [optional] [default to "Update"]
**IsMalwareDetected** | Pointer to **bool** | Whether the package has been detected as containing malware. Requires Ultra plan. | [optional] [readonly] 
**LicenseNotes** | Pointer to **NullableString** |  | [optional] 
**LicenseOverride** | Pointer to **NullableString** |  | [optional] [default to "None"]
**LicenseUrl** | Pointer to **NullableString** |  | [optional] 
**SpdxLicense** | **string** |  | 
**VulnerabilityCounts** | Pointer to [**NullableWebOSVSeverityCounts**](WebOSVSeverityCounts.md) |  | [optional] 

## Methods

### NewPackageLicenseRequestPatch

`func NewPackageLicenseRequestPatch(spdxLicense string, ) *PackageLicenseRequestPatch`

NewPackageLicenseRequestPatch instantiates a new PackageLicenseRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageLicenseRequestPatchWithDefaults

`func NewPackageLicenseRequestPatchWithDefaults() *PackageLicenseRequestPatch`

NewPackageLicenseRequestPatchWithDefaults instantiates a new PackageLicenseRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PackageLicenseRequestPatch) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PackageLicenseRequestPatch) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PackageLicenseRequestPatch) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PackageLicenseRequestPatch) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *PackageLicenseRequestPatch) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *PackageLicenseRequestPatch) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetIsMalwareDetected

`func (o *PackageLicenseRequestPatch) GetIsMalwareDetected() bool`

GetIsMalwareDetected returns the IsMalwareDetected field if non-nil, zero value otherwise.

### GetIsMalwareDetectedOk

`func (o *PackageLicenseRequestPatch) GetIsMalwareDetectedOk() (*bool, bool)`

GetIsMalwareDetectedOk returns a tuple with the IsMalwareDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMalwareDetected

`func (o *PackageLicenseRequestPatch) SetIsMalwareDetected(v bool)`

SetIsMalwareDetected sets IsMalwareDetected field to given value.

### HasIsMalwareDetected

`func (o *PackageLicenseRequestPatch) HasIsMalwareDetected() bool`

HasIsMalwareDetected returns a boolean if a field has been set.

### GetLicenseNotes

`func (o *PackageLicenseRequestPatch) GetLicenseNotes() string`

GetLicenseNotes returns the LicenseNotes field if non-nil, zero value otherwise.

### GetLicenseNotesOk

`func (o *PackageLicenseRequestPatch) GetLicenseNotesOk() (*string, bool)`

GetLicenseNotesOk returns a tuple with the LicenseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNotes

`func (o *PackageLicenseRequestPatch) SetLicenseNotes(v string)`

SetLicenseNotes sets LicenseNotes field to given value.

### HasLicenseNotes

`func (o *PackageLicenseRequestPatch) HasLicenseNotes() bool`

HasLicenseNotes returns a boolean if a field has been set.

### SetLicenseNotesNil

`func (o *PackageLicenseRequestPatch) SetLicenseNotesNil(b bool)`

 SetLicenseNotesNil sets the value for LicenseNotes to be an explicit nil

### UnsetLicenseNotes
`func (o *PackageLicenseRequestPatch) UnsetLicenseNotes()`

UnsetLicenseNotes ensures that no value is present for LicenseNotes, not even an explicit nil
### GetLicenseOverride

`func (o *PackageLicenseRequestPatch) GetLicenseOverride() string`

GetLicenseOverride returns the LicenseOverride field if non-nil, zero value otherwise.

### GetLicenseOverrideOk

`func (o *PackageLicenseRequestPatch) GetLicenseOverrideOk() (*string, bool)`

GetLicenseOverrideOk returns a tuple with the LicenseOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseOverride

`func (o *PackageLicenseRequestPatch) SetLicenseOverride(v string)`

SetLicenseOverride sets LicenseOverride field to given value.

### HasLicenseOverride

`func (o *PackageLicenseRequestPatch) HasLicenseOverride() bool`

HasLicenseOverride returns a boolean if a field has been set.

### SetLicenseOverrideNil

`func (o *PackageLicenseRequestPatch) SetLicenseOverrideNil(b bool)`

 SetLicenseOverrideNil sets the value for LicenseOverride to be an explicit nil

### UnsetLicenseOverride
`func (o *PackageLicenseRequestPatch) UnsetLicenseOverride()`

UnsetLicenseOverride ensures that no value is present for LicenseOverride, not even an explicit nil
### GetLicenseUrl

`func (o *PackageLicenseRequestPatch) GetLicenseUrl() string`

GetLicenseUrl returns the LicenseUrl field if non-nil, zero value otherwise.

### GetLicenseUrlOk

`func (o *PackageLicenseRequestPatch) GetLicenseUrlOk() (*string, bool)`

GetLicenseUrlOk returns a tuple with the LicenseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUrl

`func (o *PackageLicenseRequestPatch) SetLicenseUrl(v string)`

SetLicenseUrl sets LicenseUrl field to given value.

### HasLicenseUrl

`func (o *PackageLicenseRequestPatch) HasLicenseUrl() bool`

HasLicenseUrl returns a boolean if a field has been set.

### SetLicenseUrlNil

`func (o *PackageLicenseRequestPatch) SetLicenseUrlNil(b bool)`

 SetLicenseUrlNil sets the value for LicenseUrl to be an explicit nil

### UnsetLicenseUrl
`func (o *PackageLicenseRequestPatch) UnsetLicenseUrl()`

UnsetLicenseUrl ensures that no value is present for LicenseUrl, not even an explicit nil
### GetSpdxLicense

`func (o *PackageLicenseRequestPatch) GetSpdxLicense() string`

GetSpdxLicense returns the SpdxLicense field if non-nil, zero value otherwise.

### GetSpdxLicenseOk

`func (o *PackageLicenseRequestPatch) GetSpdxLicenseOk() (*string, bool)`

GetSpdxLicenseOk returns a tuple with the SpdxLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpdxLicense

`func (o *PackageLicenseRequestPatch) SetSpdxLicense(v string)`

SetSpdxLicense sets SpdxLicense field to given value.


### GetVulnerabilityCounts

`func (o *PackageLicenseRequestPatch) GetVulnerabilityCounts() WebOSVSeverityCounts`

GetVulnerabilityCounts returns the VulnerabilityCounts field if non-nil, zero value otherwise.

### GetVulnerabilityCountsOk

`func (o *PackageLicenseRequestPatch) GetVulnerabilityCountsOk() (*WebOSVSeverityCounts, bool)`

GetVulnerabilityCountsOk returns a tuple with the VulnerabilityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityCounts

`func (o *PackageLicenseRequestPatch) SetVulnerabilityCounts(v WebOSVSeverityCounts)`

SetVulnerabilityCounts sets VulnerabilityCounts field to given value.

### HasVulnerabilityCounts

`func (o *PackageLicenseRequestPatch) HasVulnerabilityCounts() bool`

HasVulnerabilityCounts returns a boolean if a field has been set.

### SetVulnerabilityCountsNil

`func (o *PackageLicenseRequestPatch) SetVulnerabilityCountsNil(b bool)`

 SetVulnerabilityCountsNil sets the value for VulnerabilityCounts to be an explicit nil

### UnsetVulnerabilityCounts
`func (o *PackageLicenseRequestPatch) UnsetVulnerabilityCounts()`

UnsetVulnerabilityCounts ensures that no value is present for VulnerabilityCounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


