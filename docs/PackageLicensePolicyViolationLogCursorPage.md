# PackageLicensePolicyViolationLogCursorPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]PackageLicensePolicyViolationLog**](PackageLicensePolicyViolationLog.md) |  | 

## Methods

### NewPackageLicensePolicyViolationLogCursorPage

`func NewPackageLicensePolicyViolationLogCursorPage(results []PackageLicensePolicyViolationLog, ) *PackageLicensePolicyViolationLogCursorPage`

NewPackageLicensePolicyViolationLogCursorPage instantiates a new PackageLicensePolicyViolationLogCursorPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageLicensePolicyViolationLogCursorPageWithDefaults

`func NewPackageLicensePolicyViolationLogCursorPageWithDefaults() *PackageLicensePolicyViolationLogCursorPage`

NewPackageLicensePolicyViolationLogCursorPageWithDefaults instantiates a new PackageLicensePolicyViolationLogCursorPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PackageLicensePolicyViolationLogCursorPage) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PackageLicensePolicyViolationLogCursorPage) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PackageLicensePolicyViolationLogCursorPage) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PackageLicensePolicyViolationLogCursorPage) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PackageLicensePolicyViolationLogCursorPage) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PackageLicensePolicyViolationLogCursorPage) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PackageLicensePolicyViolationLogCursorPage) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PackageLicensePolicyViolationLogCursorPage) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PackageLicensePolicyViolationLogCursorPage) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PackageLicensePolicyViolationLogCursorPage) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PackageLicensePolicyViolationLogCursorPage) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PackageLicensePolicyViolationLogCursorPage) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PackageLicensePolicyViolationLogCursorPage) GetResults() []PackageLicensePolicyViolationLog`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PackageLicensePolicyViolationLogCursorPage) GetResultsOk() (*[]PackageLicensePolicyViolationLog, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PackageLicensePolicyViolationLogCursorPage) SetResults(v []PackageLicensePolicyViolationLog)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


