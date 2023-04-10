# OrgsLicensePolicyViolationList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]PackageLicensePolicyViolationLog**](PackageLicensePolicyViolationLog.md) |  | 

## Methods

### NewOrgsLicensePolicyViolationList200Response

`func NewOrgsLicensePolicyViolationList200Response(results []PackageLicensePolicyViolationLog, ) *OrgsLicensePolicyViolationList200Response`

NewOrgsLicensePolicyViolationList200Response instantiates a new OrgsLicensePolicyViolationList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsLicensePolicyViolationList200ResponseWithDefaults

`func NewOrgsLicensePolicyViolationList200ResponseWithDefaults() *OrgsLicensePolicyViolationList200Response`

NewOrgsLicensePolicyViolationList200ResponseWithDefaults instantiates a new OrgsLicensePolicyViolationList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *OrgsLicensePolicyViolationList200Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *OrgsLicensePolicyViolationList200Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *OrgsLicensePolicyViolationList200Response) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *OrgsLicensePolicyViolationList200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *OrgsLicensePolicyViolationList200Response) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *OrgsLicensePolicyViolationList200Response) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *OrgsLicensePolicyViolationList200Response) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *OrgsLicensePolicyViolationList200Response) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *OrgsLicensePolicyViolationList200Response) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *OrgsLicensePolicyViolationList200Response) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *OrgsLicensePolicyViolationList200Response) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *OrgsLicensePolicyViolationList200Response) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *OrgsLicensePolicyViolationList200Response) GetResults() []PackageLicensePolicyViolationLog`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *OrgsLicensePolicyViolationList200Response) GetResultsOk() (*[]PackageLicensePolicyViolationLog, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *OrgsLicensePolicyViolationList200Response) SetResults(v []PackageLicensePolicyViolationLog)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


