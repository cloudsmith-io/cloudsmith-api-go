# CommonBandwidthMetricsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Display** | **string** | Bandwidth usage value | 
**Units** | Pointer to **string** | Unit of measurement e.g. bytes | [optional] [default to "bytes"]
**Value** | **int64** | Human readable version of display value | 

## Methods

### NewCommonBandwidthMetricsValue

`func NewCommonBandwidthMetricsValue(display string, value int64, ) *CommonBandwidthMetricsValue`

NewCommonBandwidthMetricsValue instantiates a new CommonBandwidthMetricsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonBandwidthMetricsValueWithDefaults

`func NewCommonBandwidthMetricsValueWithDefaults() *CommonBandwidthMetricsValue`

NewCommonBandwidthMetricsValueWithDefaults instantiates a new CommonBandwidthMetricsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplay

`func (o *CommonBandwidthMetricsValue) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CommonBandwidthMetricsValue) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CommonBandwidthMetricsValue) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUnits

`func (o *CommonBandwidthMetricsValue) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *CommonBandwidthMetricsValue) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *CommonBandwidthMetricsValue) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *CommonBandwidthMetricsValue) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetValue

`func (o *CommonBandwidthMetricsValue) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CommonBandwidthMetricsValue) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CommonBandwidthMetricsValue) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


