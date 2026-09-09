# OrganizationUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandwidthMaximum** | Pointer to **NullableInt64** | The maximum allowed package delivery overage in GB for this product. | [optional] [readonly] 
**BandwidthOverageLimit** | Pointer to **NullableInt64** | The on-demand limit for package delivery usage (in GBs). | [optional] [readonly] 
**StorageMaximum** | Pointer to **NullableInt64** | The maximum allowed artifact data overage in GB for this product. | [optional] [readonly] 
**StorageOverageLimit** | Pointer to **NullableInt64** | The on-demand limit for artifact data usage (in GBs). | [optional] [readonly] 

## Methods

### NewOrganizationUsage

`func NewOrganizationUsage() *OrganizationUsage`

NewOrganizationUsage instantiates a new OrganizationUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationUsageWithDefaults

`func NewOrganizationUsageWithDefaults() *OrganizationUsage`

NewOrganizationUsageWithDefaults instantiates a new OrganizationUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidthMaximum

`func (o *OrganizationUsage) GetBandwidthMaximum() int64`

GetBandwidthMaximum returns the BandwidthMaximum field if non-nil, zero value otherwise.

### GetBandwidthMaximumOk

`func (o *OrganizationUsage) GetBandwidthMaximumOk() (*int64, bool)`

GetBandwidthMaximumOk returns a tuple with the BandwidthMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthMaximum

`func (o *OrganizationUsage) SetBandwidthMaximum(v int64)`

SetBandwidthMaximum sets BandwidthMaximum field to given value.

### HasBandwidthMaximum

`func (o *OrganizationUsage) HasBandwidthMaximum() bool`

HasBandwidthMaximum returns a boolean if a field has been set.

### SetBandwidthMaximumNil

`func (o *OrganizationUsage) SetBandwidthMaximumNil(b bool)`

 SetBandwidthMaximumNil sets the value for BandwidthMaximum to be an explicit nil

### UnsetBandwidthMaximum
`func (o *OrganizationUsage) UnsetBandwidthMaximum()`

UnsetBandwidthMaximum ensures that no value is present for BandwidthMaximum, not even an explicit nil
### GetBandwidthOverageLimit

`func (o *OrganizationUsage) GetBandwidthOverageLimit() int64`

GetBandwidthOverageLimit returns the BandwidthOverageLimit field if non-nil, zero value otherwise.

### GetBandwidthOverageLimitOk

`func (o *OrganizationUsage) GetBandwidthOverageLimitOk() (*int64, bool)`

GetBandwidthOverageLimitOk returns a tuple with the BandwidthOverageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthOverageLimit

`func (o *OrganizationUsage) SetBandwidthOverageLimit(v int64)`

SetBandwidthOverageLimit sets BandwidthOverageLimit field to given value.

### HasBandwidthOverageLimit

`func (o *OrganizationUsage) HasBandwidthOverageLimit() bool`

HasBandwidthOverageLimit returns a boolean if a field has been set.

### SetBandwidthOverageLimitNil

`func (o *OrganizationUsage) SetBandwidthOverageLimitNil(b bool)`

 SetBandwidthOverageLimitNil sets the value for BandwidthOverageLimit to be an explicit nil

### UnsetBandwidthOverageLimit
`func (o *OrganizationUsage) UnsetBandwidthOverageLimit()`

UnsetBandwidthOverageLimit ensures that no value is present for BandwidthOverageLimit, not even an explicit nil
### GetStorageMaximum

`func (o *OrganizationUsage) GetStorageMaximum() int64`

GetStorageMaximum returns the StorageMaximum field if non-nil, zero value otherwise.

### GetStorageMaximumOk

`func (o *OrganizationUsage) GetStorageMaximumOk() (*int64, bool)`

GetStorageMaximumOk returns a tuple with the StorageMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageMaximum

`func (o *OrganizationUsage) SetStorageMaximum(v int64)`

SetStorageMaximum sets StorageMaximum field to given value.

### HasStorageMaximum

`func (o *OrganizationUsage) HasStorageMaximum() bool`

HasStorageMaximum returns a boolean if a field has been set.

### SetStorageMaximumNil

`func (o *OrganizationUsage) SetStorageMaximumNil(b bool)`

 SetStorageMaximumNil sets the value for StorageMaximum to be an explicit nil

### UnsetStorageMaximum
`func (o *OrganizationUsage) UnsetStorageMaximum()`

UnsetStorageMaximum ensures that no value is present for StorageMaximum, not even an explicit nil
### GetStorageOverageLimit

`func (o *OrganizationUsage) GetStorageOverageLimit() int64`

GetStorageOverageLimit returns the StorageOverageLimit field if non-nil, zero value otherwise.

### GetStorageOverageLimitOk

`func (o *OrganizationUsage) GetStorageOverageLimitOk() (*int64, bool)`

GetStorageOverageLimitOk returns a tuple with the StorageOverageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageOverageLimit

`func (o *OrganizationUsage) SetStorageOverageLimit(v int64)`

SetStorageOverageLimit sets StorageOverageLimit field to given value.

### HasStorageOverageLimit

`func (o *OrganizationUsage) HasStorageOverageLimit() bool`

HasStorageOverageLimit returns a boolean if a field has been set.

### SetStorageOverageLimitNil

`func (o *OrganizationUsage) SetStorageOverageLimitNil(b bool)`

 SetStorageOverageLimitNil sets the value for StorageOverageLimit to be an explicit nil

### UnsetStorageOverageLimit
`func (o *OrganizationUsage) UnsetStorageOverageLimit()`

UnsetStorageOverageLimit ensures that no value is present for StorageOverageLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


