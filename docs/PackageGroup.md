# PackageGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendKind** | Pointer to **int64** |  | [optional] 
**Count** | **NullableInt64** |  | 
**LastPush** | **NullableTime** |  | 
**Name** | Pointer to **string** |  | [optional] 
**NumDownloads** | **NullableInt64** |  | 
**Size** | **NullableInt64** |  | 

## Methods

### NewPackageGroup

`func NewPackageGroup(count NullableInt64, lastPush NullableTime, numDownloads NullableInt64, size NullableInt64, ) *PackageGroup`

NewPackageGroup instantiates a new PackageGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageGroupWithDefaults

`func NewPackageGroupWithDefaults() *PackageGroup`

NewPackageGroupWithDefaults instantiates a new PackageGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackendKind

`func (o *PackageGroup) GetBackendKind() int64`

GetBackendKind returns the BackendKind field if non-nil, zero value otherwise.

### GetBackendKindOk

`func (o *PackageGroup) GetBackendKindOk() (*int64, bool)`

GetBackendKindOk returns a tuple with the BackendKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendKind

`func (o *PackageGroup) SetBackendKind(v int64)`

SetBackendKind sets BackendKind field to given value.

### HasBackendKind

`func (o *PackageGroup) HasBackendKind() bool`

HasBackendKind returns a boolean if a field has been set.

### GetCount

`func (o *PackageGroup) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PackageGroup) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PackageGroup) SetCount(v int64)`

SetCount sets Count field to given value.


### SetCountNil

`func (o *PackageGroup) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *PackageGroup) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLastPush

`func (o *PackageGroup) GetLastPush() time.Time`

GetLastPush returns the LastPush field if non-nil, zero value otherwise.

### GetLastPushOk

`func (o *PackageGroup) GetLastPushOk() (*time.Time, bool)`

GetLastPushOk returns a tuple with the LastPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPush

`func (o *PackageGroup) SetLastPush(v time.Time)`

SetLastPush sets LastPush field to given value.


### SetLastPushNil

`func (o *PackageGroup) SetLastPushNil(b bool)`

 SetLastPushNil sets the value for LastPush to be an explicit nil

### UnsetLastPush
`func (o *PackageGroup) UnsetLastPush()`

UnsetLastPush ensures that no value is present for LastPush, not even an explicit nil
### GetName

`func (o *PackageGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackageGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackageGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PackageGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumDownloads

`func (o *PackageGroup) GetNumDownloads() int64`

GetNumDownloads returns the NumDownloads field if non-nil, zero value otherwise.

### GetNumDownloadsOk

`func (o *PackageGroup) GetNumDownloadsOk() (*int64, bool)`

GetNumDownloadsOk returns a tuple with the NumDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDownloads

`func (o *PackageGroup) SetNumDownloads(v int64)`

SetNumDownloads sets NumDownloads field to given value.


### SetNumDownloadsNil

`func (o *PackageGroup) SetNumDownloadsNil(b bool)`

 SetNumDownloadsNil sets the value for NumDownloads to be an explicit nil

### UnsetNumDownloads
`func (o *PackageGroup) UnsetNumDownloads()`

UnsetNumDownloads ensures that no value is present for NumDownloads, not even an explicit nil
### GetSize

`func (o *PackageGroup) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PackageGroup) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PackageGroup) SetSize(v int64)`

SetSize sets Size field to given value.


### SetSizeNil

`func (o *PackageGroup) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *PackageGroup) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


