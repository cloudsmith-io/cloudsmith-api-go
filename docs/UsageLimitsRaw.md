# UsageLimitsRaw

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bandwidth** | [**AllocatedLimitRaw**](AllocatedLimitRaw.md) |  | 
**Storage** | [**StorageAllocatedLimitRaw**](StorageAllocatedLimitRaw.md) |  | 

## Methods

### NewUsageLimitsRaw

`func NewUsageLimitsRaw(bandwidth AllocatedLimitRaw, storage StorageAllocatedLimitRaw, ) *UsageLimitsRaw`

NewUsageLimitsRaw instantiates a new UsageLimitsRaw object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageLimitsRawWithDefaults

`func NewUsageLimitsRawWithDefaults() *UsageLimitsRaw`

NewUsageLimitsRawWithDefaults instantiates a new UsageLimitsRaw object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidth

`func (o *UsageLimitsRaw) GetBandwidth() AllocatedLimitRaw`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *UsageLimitsRaw) GetBandwidthOk() (*AllocatedLimitRaw, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *UsageLimitsRaw) SetBandwidth(v AllocatedLimitRaw)`

SetBandwidth sets Bandwidth field to given value.


### GetStorage

`func (o *UsageLimitsRaw) GetStorage() StorageAllocatedLimitRaw`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *UsageLimitsRaw) GetStorageOk() (*StorageAllocatedLimitRaw, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *UsageLimitsRaw) SetStorage(v StorageAllocatedLimitRaw)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


