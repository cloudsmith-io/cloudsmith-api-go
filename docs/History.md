# History

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **int64** |  | [optional] 
**Display** | [**HistoryFieldset**](HistoryFieldset.md) |  | 
**End** | **time.Time** |  | 
**Plan** | **string** |  | 
**Raw** | [**HistoryFieldset**](HistoryFieldset.md) |  | 
**Start** | **time.Time** |  | 

## Methods

### NewHistory

`func NewHistory(display HistoryFieldset, end time.Time, plan string, raw HistoryFieldset, start time.Time, ) *History`

NewHistory instantiates a new History object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryWithDefaults

`func NewHistoryWithDefaults() *History`

NewHistoryWithDefaults instantiates a new History object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *History) GetDays() int64`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *History) GetDaysOk() (*int64, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *History) SetDays(v int64)`

SetDays sets Days field to given value.

### HasDays

`func (o *History) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetDisplay

`func (o *History) GetDisplay() HistoryFieldset`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *History) GetDisplayOk() (*HistoryFieldset, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *History) SetDisplay(v HistoryFieldset)`

SetDisplay sets Display field to given value.


### GetEnd

`func (o *History) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *History) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *History) SetEnd(v time.Time)`

SetEnd sets End field to given value.


### GetPlan

`func (o *History) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *History) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *History) SetPlan(v string)`

SetPlan sets Plan field to given value.


### GetRaw

`func (o *History) GetRaw() HistoryFieldset`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *History) GetRawOk() (*HistoryFieldset, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *History) SetRaw(v HistoryFieldset)`

SetRaw sets Raw field to given value.


### GetStart

`func (o *History) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *History) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *History) SetStart(v time.Time)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


