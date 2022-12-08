# PackageFilePartsUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The identifier for the file to use uploading parts. | [optional] [readonly] 
**UploadQuerystring** | Pointer to **NullableString** | The querystring to use for the next-step PUT upload. | [optional] [readonly] 
**UploadUrl** | Pointer to **string** | The URL to use for the next-step PUT upload | [optional] [readonly] 

## Methods

### NewPackageFilePartsUploadResponse

`func NewPackageFilePartsUploadResponse() *PackageFilePartsUploadResponse`

NewPackageFilePartsUploadResponse instantiates a new PackageFilePartsUploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageFilePartsUploadResponseWithDefaults

`func NewPackageFilePartsUploadResponseWithDefaults() *PackageFilePartsUploadResponse`

NewPackageFilePartsUploadResponseWithDefaults instantiates a new PackageFilePartsUploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *PackageFilePartsUploadResponse) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PackageFilePartsUploadResponse) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PackageFilePartsUploadResponse) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PackageFilePartsUploadResponse) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetUploadQuerystring

`func (o *PackageFilePartsUploadResponse) GetUploadQuerystring() string`

GetUploadQuerystring returns the UploadQuerystring field if non-nil, zero value otherwise.

### GetUploadQuerystringOk

`func (o *PackageFilePartsUploadResponse) GetUploadQuerystringOk() (*string, bool)`

GetUploadQuerystringOk returns a tuple with the UploadQuerystring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadQuerystring

`func (o *PackageFilePartsUploadResponse) SetUploadQuerystring(v string)`

SetUploadQuerystring sets UploadQuerystring field to given value.

### HasUploadQuerystring

`func (o *PackageFilePartsUploadResponse) HasUploadQuerystring() bool`

HasUploadQuerystring returns a boolean if a field has been set.

### SetUploadQuerystringNil

`func (o *PackageFilePartsUploadResponse) SetUploadQuerystringNil(b bool)`

 SetUploadQuerystringNil sets the value for UploadQuerystring to be an explicit nil

### UnsetUploadQuerystring
`func (o *PackageFilePartsUploadResponse) UnsetUploadQuerystring()`

UnsetUploadQuerystring ensures that no value is present for UploadQuerystring, not even an explicit nil
### GetUploadUrl

`func (o *PackageFilePartsUploadResponse) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *PackageFilePartsUploadResponse) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *PackageFilePartsUploadResponse) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *PackageFilePartsUploadResponse) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


