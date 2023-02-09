# UsageLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bandwidth** | [**AllocatedLimit**](AllocatedLimit.md) |  | 
**Storage** | [**StorageAllocatedLimit**](StorageAllocatedLimit.md) |  | 

## Methods

### NewUsageLimits

`func NewUsageLimits(bandwidth AllocatedLimit, storage StorageAllocatedLimit, ) *UsageLimits`

NewUsageLimits instantiates a new UsageLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageLimitsWithDefaults

`func NewUsageLimitsWithDefaults() *UsageLimits`

NewUsageLimitsWithDefaults instantiates a new UsageLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidth

`func (o *UsageLimits) GetBandwidth() AllocatedLimit`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *UsageLimits) GetBandwidthOk() (*AllocatedLimit, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *UsageLimits) SetBandwidth(v AllocatedLimit)`

SetBandwidth sets Bandwidth field to given value.


### GetStorage

`func (o *UsageLimits) GetStorage() StorageAllocatedLimit`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *UsageLimits) GetStorageOk() (*StorageAllocatedLimit, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *UsageLimits) SetStorage(v StorageAllocatedLimit)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


