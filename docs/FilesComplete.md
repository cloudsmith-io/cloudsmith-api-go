# FilesComplete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** | Filename for the package file upload. | 
**Md5Checksum** | Pointer to **string** | MD5 checksum for a POST-based package file upload. | [optional] 
**Method** | Pointer to **string** | The method to use for package file upload. | [optional] 
**Sha256Checksum** | Pointer to **string** | SHA256 checksum for a PUT-based package file upload. | [optional] 

## Methods

### NewFilesComplete

`func NewFilesComplete(filename string, ) *FilesComplete`

NewFilesComplete instantiates a new FilesComplete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesCompleteWithDefaults

`func NewFilesCompleteWithDefaults() *FilesComplete`

NewFilesCompleteWithDefaults instantiates a new FilesComplete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *FilesComplete) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *FilesComplete) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *FilesComplete) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetMd5Checksum

`func (o *FilesComplete) GetMd5Checksum() string`

GetMd5Checksum returns the Md5Checksum field if non-nil, zero value otherwise.

### GetMd5ChecksumOk

`func (o *FilesComplete) GetMd5ChecksumOk() (*string, bool)`

GetMd5ChecksumOk returns a tuple with the Md5Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Checksum

`func (o *FilesComplete) SetMd5Checksum(v string)`

SetMd5Checksum sets Md5Checksum field to given value.

### HasMd5Checksum

`func (o *FilesComplete) HasMd5Checksum() bool`

HasMd5Checksum returns a boolean if a field has been set.

### GetMethod

`func (o *FilesComplete) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *FilesComplete) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *FilesComplete) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *FilesComplete) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetSha256Checksum

`func (o *FilesComplete) GetSha256Checksum() string`

GetSha256Checksum returns the Sha256Checksum field if non-nil, zero value otherwise.

### GetSha256ChecksumOk

`func (o *FilesComplete) GetSha256ChecksumOk() (*string, bool)`

GetSha256ChecksumOk returns a tuple with the Sha256Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Checksum

`func (o *FilesComplete) SetSha256Checksum(v string)`

SetSha256Checksum sets Sha256Checksum field to given value.

### HasSha256Checksum

`func (o *FilesComplete) HasSha256Checksum() bool`

HasSha256Checksum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


