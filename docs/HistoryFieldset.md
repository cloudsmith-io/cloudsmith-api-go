# HistoryFieldset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Downloaded** | [**Usage**](Usage.md) |  | 
**StorageUsed** | [**StorageUsage**](StorageUsage.md) |  | 
**Uploaded** | [**Usage**](Usage.md) |  | 

## Methods

### NewHistoryFieldset

`func NewHistoryFieldset(downloaded Usage, storageUsed StorageUsage, uploaded Usage, ) *HistoryFieldset`

NewHistoryFieldset instantiates a new HistoryFieldset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryFieldsetWithDefaults

`func NewHistoryFieldsetWithDefaults() *HistoryFieldset`

NewHistoryFieldsetWithDefaults instantiates a new HistoryFieldset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloaded

`func (o *HistoryFieldset) GetDownloaded() Usage`

GetDownloaded returns the Downloaded field if non-nil, zero value otherwise.

### GetDownloadedOk

`func (o *HistoryFieldset) GetDownloadedOk() (*Usage, bool)`

GetDownloadedOk returns a tuple with the Downloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloaded

`func (o *HistoryFieldset) SetDownloaded(v Usage)`

SetDownloaded sets Downloaded field to given value.


### GetStorageUsed

`func (o *HistoryFieldset) GetStorageUsed() StorageUsage`

GetStorageUsed returns the StorageUsed field if non-nil, zero value otherwise.

### GetStorageUsedOk

`func (o *HistoryFieldset) GetStorageUsedOk() (*StorageUsage, bool)`

GetStorageUsedOk returns a tuple with the StorageUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUsed

`func (o *HistoryFieldset) SetStorageUsed(v StorageUsage)`

SetStorageUsed sets StorageUsed field to given value.


### GetUploaded

`func (o *HistoryFieldset) GetUploaded() Usage`

GetUploaded returns the Uploaded field if non-nil, zero value otherwise.

### GetUploadedOk

`func (o *HistoryFieldset) GetUploadedOk() (*Usage, bool)`

GetUploadedOk returns a tuple with the Uploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaded

`func (o *HistoryFieldset) SetUploaded(v Usage)`

SetUploaded sets Uploaded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


