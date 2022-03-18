# ResourcesRateCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to **map[string]interface{}** | Rate limit values per resource | [optional] 

## Methods

### NewResourcesRateCheck

`func NewResourcesRateCheck() *ResourcesRateCheck`

NewResourcesRateCheck instantiates a new ResourcesRateCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesRateCheckWithDefaults

`func NewResourcesRateCheckWithDefaults() *ResourcesRateCheck`

NewResourcesRateCheckWithDefaults instantiates a new ResourcesRateCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *ResourcesRateCheck) GetResources() map[string]interface{}`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ResourcesRateCheck) GetResourcesOk() (*map[string]interface{}, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ResourcesRateCheck) SetResources(v map[string]interface{})`

SetResources sets Resources field to given value.

### HasResources

`func (o *ResourcesRateCheck) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


