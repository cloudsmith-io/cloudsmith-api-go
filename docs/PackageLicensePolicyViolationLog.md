# PackageLicensePolicyViolationLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Package** | [**PackageVulnerability**](PackageVulnerability.md) |  | 
**Policy** | [**NestedLicensePolicy**](NestedLicensePolicy.md) |  | 
**Reasons** | **[]string** |  | 

## Methods

### NewPackageLicensePolicyViolationLog

`func NewPackageLicensePolicyViolationLog(package_ PackageVulnerability, policy NestedLicensePolicy, reasons []string, ) *PackageLicensePolicyViolationLog`

NewPackageLicensePolicyViolationLog instantiates a new PackageLicensePolicyViolationLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageLicensePolicyViolationLogWithDefaults

`func NewPackageLicensePolicyViolationLogWithDefaults() *PackageLicensePolicyViolationLog`

NewPackageLicensePolicyViolationLogWithDefaults instantiates a new PackageLicensePolicyViolationLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventAt

`func (o *PackageLicensePolicyViolationLog) GetEventAt() time.Time`

GetEventAt returns the EventAt field if non-nil, zero value otherwise.

### GetEventAtOk

`func (o *PackageLicensePolicyViolationLog) GetEventAtOk() (*time.Time, bool)`

GetEventAtOk returns a tuple with the EventAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAt

`func (o *PackageLicensePolicyViolationLog) SetEventAt(v time.Time)`

SetEventAt sets EventAt field to given value.

### HasEventAt

`func (o *PackageLicensePolicyViolationLog) HasEventAt() bool`

HasEventAt returns a boolean if a field has been set.

### GetPackage

`func (o *PackageLicensePolicyViolationLog) GetPackage() PackageVulnerability`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *PackageLicensePolicyViolationLog) GetPackageOk() (*PackageVulnerability, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *PackageLicensePolicyViolationLog) SetPackage(v PackageVulnerability)`

SetPackage sets Package field to given value.


### GetPolicy

`func (o *PackageLicensePolicyViolationLog) GetPolicy() NestedLicensePolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PackageLicensePolicyViolationLog) GetPolicyOk() (*NestedLicensePolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PackageLicensePolicyViolationLog) SetPolicy(v NestedLicensePolicy)`

SetPolicy sets Policy field to given value.


### GetReasons

`func (o *PackageLicensePolicyViolationLog) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *PackageLicensePolicyViolationLog) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *PackageLicensePolicyViolationLog) SetReasons(v []string)`

SetReasons sets Reasons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


