# OrgsUpdateUsageLimits200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowOpenSourceOverage** | Pointer to **bool** | Whether on-demand open source overage is allowed. | [optional] 
**BandwidthOverageLimit** | Pointer to **int64** | Effective bandwidth overage limit in GB. | [optional] 
**StorageOverageLimit** | Pointer to **int64** | Effective storage overage limit in GB. | [optional] 

## Methods

### NewOrgsUpdateUsageLimits200Response

`func NewOrgsUpdateUsageLimits200Response() *OrgsUpdateUsageLimits200Response`

NewOrgsUpdateUsageLimits200Response instantiates a new OrgsUpdateUsageLimits200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsUpdateUsageLimits200ResponseWithDefaults

`func NewOrgsUpdateUsageLimits200ResponseWithDefaults() *OrgsUpdateUsageLimits200Response`

NewOrgsUpdateUsageLimits200ResponseWithDefaults instantiates a new OrgsUpdateUsageLimits200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowOpenSourceOverage

`func (o *OrgsUpdateUsageLimits200Response) GetAllowOpenSourceOverage() bool`

GetAllowOpenSourceOverage returns the AllowOpenSourceOverage field if non-nil, zero value otherwise.

### GetAllowOpenSourceOverageOk

`func (o *OrgsUpdateUsageLimits200Response) GetAllowOpenSourceOverageOk() (*bool, bool)`

GetAllowOpenSourceOverageOk returns a tuple with the AllowOpenSourceOverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOpenSourceOverage

`func (o *OrgsUpdateUsageLimits200Response) SetAllowOpenSourceOverage(v bool)`

SetAllowOpenSourceOverage sets AllowOpenSourceOverage field to given value.

### HasAllowOpenSourceOverage

`func (o *OrgsUpdateUsageLimits200Response) HasAllowOpenSourceOverage() bool`

HasAllowOpenSourceOverage returns a boolean if a field has been set.

### GetBandwidthOverageLimit

`func (o *OrgsUpdateUsageLimits200Response) GetBandwidthOverageLimit() int64`

GetBandwidthOverageLimit returns the BandwidthOverageLimit field if non-nil, zero value otherwise.

### GetBandwidthOverageLimitOk

`func (o *OrgsUpdateUsageLimits200Response) GetBandwidthOverageLimitOk() (*int64, bool)`

GetBandwidthOverageLimitOk returns a tuple with the BandwidthOverageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthOverageLimit

`func (o *OrgsUpdateUsageLimits200Response) SetBandwidthOverageLimit(v int64)`

SetBandwidthOverageLimit sets BandwidthOverageLimit field to given value.

### HasBandwidthOverageLimit

`func (o *OrgsUpdateUsageLimits200Response) HasBandwidthOverageLimit() bool`

HasBandwidthOverageLimit returns a boolean if a field has been set.

### GetStorageOverageLimit

`func (o *OrgsUpdateUsageLimits200Response) GetStorageOverageLimit() int64`

GetStorageOverageLimit returns the StorageOverageLimit field if non-nil, zero value otherwise.

### GetStorageOverageLimitOk

`func (o *OrgsUpdateUsageLimits200Response) GetStorageOverageLimitOk() (*int64, bool)`

GetStorageOverageLimitOk returns a tuple with the StorageOverageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageOverageLimit

`func (o *OrgsUpdateUsageLimits200Response) SetStorageOverageLimit(v int64)`

SetStorageOverageLimit sets StorageOverageLimit field to given value.

### HasStorageOverageLimit

`func (o *OrgsUpdateUsageLimits200Response) HasStorageOverageLimit() bool`

HasStorageOverageLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


