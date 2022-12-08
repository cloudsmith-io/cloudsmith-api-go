# RateCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | Pointer to **float64** | The time in seconds that you are suggested to wait until the next request in order to avoid consuming too much within the rate limit window. | [optional] [readonly] 
**Limit** | Pointer to **int64** | The maximum number of requests that you are permitted to send per hour | [optional] [readonly] 
**Remaining** | Pointer to **int64** | The number of requests that are remaining in the current rate limit window | [optional] [readonly] 
**Reset** | Pointer to **int64** | The UTC epoch timestamp at which the current rate limit window will reset | [optional] [readonly] 
**ResetIso8601** | Pointer to **string** | The ISO 8601 datetime at which the current rate limit window will reset | [optional] [readonly] 
**Throttled** | Pointer to **bool** | If true, throttling is currently being enforced. | [optional] [readonly] 

## Methods

### NewRateCheck

`func NewRateCheck() *RateCheck`

NewRateCheck instantiates a new RateCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateCheckWithDefaults

`func NewRateCheckWithDefaults() *RateCheck`

NewRateCheckWithDefaults instantiates a new RateCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *RateCheck) GetInterval() float64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RateCheck) GetIntervalOk() (*float64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RateCheck) SetInterval(v float64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *RateCheck) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLimit

`func (o *RateCheck) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RateCheck) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RateCheck) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RateCheck) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetRemaining

`func (o *RateCheck) GetRemaining() int64`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *RateCheck) GetRemainingOk() (*int64, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *RateCheck) SetRemaining(v int64)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *RateCheck) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.

### GetReset

`func (o *RateCheck) GetReset() int64`

GetReset returns the Reset field if non-nil, zero value otherwise.

### GetResetOk

`func (o *RateCheck) GetResetOk() (*int64, bool)`

GetResetOk returns a tuple with the Reset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReset

`func (o *RateCheck) SetReset(v int64)`

SetReset sets Reset field to given value.

### HasReset

`func (o *RateCheck) HasReset() bool`

HasReset returns a boolean if a field has been set.

### GetResetIso8601

`func (o *RateCheck) GetResetIso8601() string`

GetResetIso8601 returns the ResetIso8601 field if non-nil, zero value otherwise.

### GetResetIso8601Ok

`func (o *RateCheck) GetResetIso8601Ok() (*string, bool)`

GetResetIso8601Ok returns a tuple with the ResetIso8601 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIso8601

`func (o *RateCheck) SetResetIso8601(v string)`

SetResetIso8601 sets ResetIso8601 field to given value.

### HasResetIso8601

`func (o *RateCheck) HasResetIso8601() bool`

HasResetIso8601 returns a boolean if a field has been set.

### GetThrottled

`func (o *RateCheck) GetThrottled() bool`

GetThrottled returns the Throttled field if non-nil, zero value otherwise.

### GetThrottledOk

`func (o *RateCheck) GetThrottledOk() (*bool, bool)`

GetThrottledOk returns a tuple with the Throttled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottled

`func (o *RateCheck) SetThrottled(v bool)`

SetThrottled sets Throttled field to given value.

### HasThrottled

`func (o *RateCheck) HasThrottled() bool`

HasThrottled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


