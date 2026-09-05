# PackageMoveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | **string** | The name of the destination repository without the namespace. | 

## Methods

### NewPackageMoveRequest

`func NewPackageMoveRequest(destination string, ) *PackageMoveRequest`

NewPackageMoveRequest instantiates a new PackageMoveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageMoveRequestWithDefaults

`func NewPackageMoveRequestWithDefaults() *PackageMoveRequest`

NewPackageMoveRequestWithDefaults instantiates a new PackageMoveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *PackageMoveRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PackageMoveRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PackageMoveRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


