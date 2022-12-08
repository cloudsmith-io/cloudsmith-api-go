# QuotaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usage** | [**UsageFieldset**](UsageFieldset.md) |  | 

## Methods

### NewQuotaResponse

`func NewQuotaResponse(usage UsageFieldset, ) *QuotaResponse`

NewQuotaResponse instantiates a new QuotaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaResponseWithDefaults

`func NewQuotaResponseWithDefaults() *QuotaResponse`

NewQuotaResponseWithDefaults instantiates a new QuotaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsage

`func (o *QuotaResponse) GetUsage() UsageFieldset`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *QuotaResponse) GetUsageOk() (*UsageFieldset, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *QuotaResponse) SetUsage(v UsageFieldset)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


