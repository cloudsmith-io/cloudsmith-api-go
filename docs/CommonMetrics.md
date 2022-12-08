# CommonMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int64** |  | [optional] 
**Bandwidth** | [**CommonBandwidthMetrics**](CommonBandwidthMetrics.md) |  | 
**Downloads** | [**CommonDownloadsMetrics**](CommonDownloadsMetrics.md) |  | 
**Inactive** | Pointer to **int64** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewCommonMetrics

`func NewCommonMetrics(bandwidth CommonBandwidthMetrics, downloads CommonDownloadsMetrics, ) *CommonMetrics`

NewCommonMetrics instantiates a new CommonMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonMetricsWithDefaults

`func NewCommonMetricsWithDefaults() *CommonMetrics`

NewCommonMetricsWithDefaults instantiates a new CommonMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CommonMetrics) GetActive() int64`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CommonMetrics) GetActiveOk() (*int64, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CommonMetrics) SetActive(v int64)`

SetActive sets Active field to given value.

### HasActive

`func (o *CommonMetrics) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBandwidth

`func (o *CommonMetrics) GetBandwidth() CommonBandwidthMetrics`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *CommonMetrics) GetBandwidthOk() (*CommonBandwidthMetrics, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *CommonMetrics) SetBandwidth(v CommonBandwidthMetrics)`

SetBandwidth sets Bandwidth field to given value.


### GetDownloads

`func (o *CommonMetrics) GetDownloads() CommonDownloadsMetrics`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *CommonMetrics) GetDownloadsOk() (*CommonDownloadsMetrics, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *CommonMetrics) SetDownloads(v CommonDownloadsMetrics)`

SetDownloads sets Downloads field to given value.


### GetInactive

`func (o *CommonMetrics) GetInactive() int64`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *CommonMetrics) GetInactiveOk() (*int64, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *CommonMetrics) SetInactive(v int64)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *CommonMetrics) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### GetTotal

`func (o *CommonMetrics) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CommonMetrics) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CommonMetrics) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CommonMetrics) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


