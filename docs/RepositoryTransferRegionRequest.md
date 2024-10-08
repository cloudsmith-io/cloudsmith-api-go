# RepositoryTransferRegionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageRegion** | Pointer to **string** | The Cloudsmith region in which package files are stored. | [optional] [default to "default"]

## Methods

### NewRepositoryTransferRegionRequest

`func NewRepositoryTransferRegionRequest() *RepositoryTransferRegionRequest`

NewRepositoryTransferRegionRequest instantiates a new RepositoryTransferRegionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryTransferRegionRequestWithDefaults

`func NewRepositoryTransferRegionRequestWithDefaults() *RepositoryTransferRegionRequest`

NewRepositoryTransferRegionRequestWithDefaults instantiates a new RepositoryTransferRegionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageRegion

`func (o *RepositoryTransferRegionRequest) GetStorageRegion() string`

GetStorageRegion returns the StorageRegion field if non-nil, zero value otherwise.

### GetStorageRegionOk

`func (o *RepositoryTransferRegionRequest) GetStorageRegionOk() (*string, bool)`

GetStorageRegionOk returns a tuple with the StorageRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRegion

`func (o *RepositoryTransferRegionRequest) SetStorageRegion(v string)`

SetStorageRegion sets StorageRegion field to given value.

### HasStorageRegion

`func (o *RepositoryTransferRegionRequest) HasStorageRegion() bool`

HasStorageRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


