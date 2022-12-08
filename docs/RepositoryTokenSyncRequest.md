# RepositoryTokenSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | The source repository slug (in the same owner namespace). | 

## Methods

### NewRepositoryTokenSyncRequest

`func NewRepositoryTokenSyncRequest(source string, ) *RepositoryTokenSyncRequest`

NewRepositoryTokenSyncRequest instantiates a new RepositoryTokenSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryTokenSyncRequestWithDefaults

`func NewRepositoryTokenSyncRequestWithDefaults() *RepositoryTokenSyncRequest`

NewRepositoryTokenSyncRequestWithDefaults instantiates a new RepositoryTokenSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RepositoryTokenSyncRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RepositoryTokenSyncRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RepositoryTokenSyncRequest) SetSource(v string)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


