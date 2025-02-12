# ErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | **string** | An extended message for the response. | 
**Fields** | Pointer to **map[string][]string** | A Dictionary of related errors where key: Field and value: Array of Errors related to that field | [optional] 

## Methods

### NewErrorDetail

`func NewErrorDetail(detail string, ) *ErrorDetail`

NewErrorDetail instantiates a new ErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDetailWithDefaults

`func NewErrorDetailWithDefaults() *ErrorDetail`

NewErrorDetailWithDefaults instantiates a new ErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *ErrorDetail) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorDetail) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorDetail) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetFields

`func (o *ErrorDetail) GetFields() map[string][]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ErrorDetail) GetFieldsOk() (*map[string][]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ErrorDetail) SetFields(v map[string][]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ErrorDetail) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


