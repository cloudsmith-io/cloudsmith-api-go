# HistoryFieldsetRaw

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Downloaded** | [**UsageRaw**](UsageRaw.md) |  | 
**StorageUsed** | [**StorageUsageRaw**](StorageUsageRaw.md) |  | 
**Uploaded** | [**UsageRaw**](UsageRaw.md) |  | 

## Methods

### NewHistoryFieldsetRaw

`func NewHistoryFieldsetRaw(downloaded UsageRaw, storageUsed StorageUsageRaw, uploaded UsageRaw, ) *HistoryFieldsetRaw`

NewHistoryFieldsetRaw instantiates a new HistoryFieldsetRaw object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryFieldsetRawWithDefaults

`func NewHistoryFieldsetRawWithDefaults() *HistoryFieldsetRaw`

NewHistoryFieldsetRawWithDefaults instantiates a new HistoryFieldsetRaw object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloaded

`func (o *HistoryFieldsetRaw) GetDownloaded() UsageRaw`

GetDownloaded returns the Downloaded field if non-nil, zero value otherwise.

### GetDownloadedOk

`func (o *HistoryFieldsetRaw) GetDownloadedOk() (*UsageRaw, bool)`

GetDownloadedOk returns a tuple with the Downloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloaded

`func (o *HistoryFieldsetRaw) SetDownloaded(v UsageRaw)`

SetDownloaded sets Downloaded field to given value.


### GetStorageUsed

`func (o *HistoryFieldsetRaw) GetStorageUsed() StorageUsageRaw`

GetStorageUsed returns the StorageUsed field if non-nil, zero value otherwise.

### GetStorageUsedOk

`func (o *HistoryFieldsetRaw) GetStorageUsedOk() (*StorageUsageRaw, bool)`

GetStorageUsedOk returns a tuple with the StorageUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUsed

`func (o *HistoryFieldsetRaw) SetStorageUsed(v StorageUsageRaw)`

SetStorageUsed sets StorageUsed field to given value.


### GetUploaded

`func (o *HistoryFieldsetRaw) GetUploaded() UsageRaw`

GetUploaded returns the Uploaded field if non-nil, zero value otherwise.

### GetUploadedOk

`func (o *HistoryFieldsetRaw) GetUploadedOk() (*UsageRaw, bool)`

GetUploadedOk returns a tuple with the Uploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaded

`func (o *HistoryFieldsetRaw) SetUploaded(v UsageRaw)`

SetUploaded sets Uploaded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


