# PackageFilePartsUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The identifier for the file to use uploading parts. | [optional] 
**UploadQuerystring** | Pointer to **string** | The querystring to use for the next-step PUT upload. | [optional] 
**UploadUrl** | Pointer to **string** | The URL to use for the next-step PUT upload | [optional] 

## Methods

### NewPackageFilePartsUpload

`func NewPackageFilePartsUpload() *PackageFilePartsUpload`

NewPackageFilePartsUpload instantiates a new PackageFilePartsUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageFilePartsUploadWithDefaults

`func NewPackageFilePartsUploadWithDefaults() *PackageFilePartsUpload`

NewPackageFilePartsUploadWithDefaults instantiates a new PackageFilePartsUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *PackageFilePartsUpload) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PackageFilePartsUpload) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PackageFilePartsUpload) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PackageFilePartsUpload) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetUploadQuerystring

`func (o *PackageFilePartsUpload) GetUploadQuerystring() string`

GetUploadQuerystring returns the UploadQuerystring field if non-nil, zero value otherwise.

### GetUploadQuerystringOk

`func (o *PackageFilePartsUpload) GetUploadQuerystringOk() (*string, bool)`

GetUploadQuerystringOk returns a tuple with the UploadQuerystring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadQuerystring

`func (o *PackageFilePartsUpload) SetUploadQuerystring(v string)`

SetUploadQuerystring sets UploadQuerystring field to given value.

### HasUploadQuerystring

`func (o *PackageFilePartsUpload) HasUploadQuerystring() bool`

HasUploadQuerystring returns a boolean if a field has been set.

### GetUploadUrl

`func (o *PackageFilePartsUpload) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *PackageFilePartsUpload) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *PackageFilePartsUpload) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *PackageFilePartsUpload) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


