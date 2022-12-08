# StatusBasicResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** | The message describing the state of the API. | [optional] [readonly] [default to "Cloudsmith API is operational."]
**Version** | Pointer to **string** | The current version for the Cloudsmith service. | [optional] [readonly] [default to "1.181.6"]

## Methods

### NewStatusBasicResponse

`func NewStatusBasicResponse() *StatusBasicResponse`

NewStatusBasicResponse instantiates a new StatusBasicResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusBasicResponseWithDefaults

`func NewStatusBasicResponseWithDefaults() *StatusBasicResponse`

NewStatusBasicResponseWithDefaults instantiates a new StatusBasicResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *StatusBasicResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *StatusBasicResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *StatusBasicResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *StatusBasicResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetVersion

`func (o *StatusBasicResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StatusBasicResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StatusBasicResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StatusBasicResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


