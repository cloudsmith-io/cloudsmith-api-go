# PackageQuarantineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Release** | Pointer to **bool** | If true, the package is released from quarantine. | [optional] 
**Restore** | Pointer to **bool** | If true, the package is released from quarantine. Note: This field is deprecated, please use &#39;release&#39; instead. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


