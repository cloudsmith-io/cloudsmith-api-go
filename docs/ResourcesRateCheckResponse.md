# ResourcesRateCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**map[string]RateCheck**](RateCheck.md) | Rate limit values per resource | [optional] [readonly] 

## Methods

### NewResourcesRateCheckResponse

`func NewResourcesRateCheckResponse() *ResourcesRateCheckResponse`

NewResourcesRateCheckResponse instantiates a new ResourcesRateCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesRateCheckResponseWithDefaults

`func NewResourcesRateCheckResponseWithDefaults() *ResourcesRateCheckResponse`

NewResourcesRateCheckResponseWithDefaults instantiates a new ResourcesRateCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *ResourcesRateCheckResponse) GetResources() map[string]RateCheck`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ResourcesRateCheckResponse) GetResourcesOk() (*map[string]RateCheck, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ResourcesRateCheckResponse) SetResources(v map[string]RateCheck)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ResourcesRateCheckResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


