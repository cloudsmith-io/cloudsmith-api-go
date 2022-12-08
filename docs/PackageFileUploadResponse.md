# PackageFileUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The identifier for the file to use when creating packages | [optional] [readonly] 
**UploadFields** | Pointer to **map[string]interface{}** | The dictionary of fields that must be sent with POST uploads | [optional] [readonly] 
**UploadHeaders** | Pointer to **map[string]interface{}** | The dictionary of headers that must be sent with uploads | [optional] [readonly] 
**UploadQuerystring** | Pointer to **NullableString** | The querystring to use for the next-step POST or PUT upload | [optional] [readonly] 
**UploadUrl** | Pointer to **string** | The URL to use for the next-step POST or PUT upload | [optional] [readonly] 

## Methods

### NewPackageFileUploadResponse

`func NewPackageFileUploadResponse() *PackageFileUploadResponse`

NewPackageFileUploadResponse instantiates a new PackageFileUploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageFileUploadResponseWithDefaults

`func NewPackageFileUploadResponseWithDefaults() *PackageFileUploadResponse`

NewPackageFileUploadResponseWithDefaults instantiates a new PackageFileUploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *PackageFileUploadResponse) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PackageFileUploadResponse) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PackageFileUploadResponse) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PackageFileUploadResponse) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetUploadFields

`func (o *PackageFileUploadResponse) GetUploadFields() map[string]interface{}`

GetUploadFields returns the UploadFields field if non-nil, zero value otherwise.

### GetUploadFieldsOk

`func (o *PackageFileUploadResponse) GetUploadFieldsOk() (*map[string]interface{}, bool)`

GetUploadFieldsOk returns a tuple with the UploadFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadFields

`func (o *PackageFileUploadResponse) SetUploadFields(v map[string]interface{})`

SetUploadFields sets UploadFields field to given value.

### HasUploadFields

`func (o *PackageFileUploadResponse) HasUploadFields() bool`

HasUploadFields returns a boolean if a field has been set.

### SetUploadFieldsNil

`func (o *PackageFileUploadResponse) SetUploadFieldsNil(b bool)`

 SetUploadFieldsNil sets the value for UploadFields to be an explicit nil

### UnsetUploadFields
`func (o *PackageFileUploadResponse) UnsetUploadFields()`

UnsetUploadFields ensures that no value is present for UploadFields, not even an explicit nil
### GetUploadHeaders

`func (o *PackageFileUploadResponse) GetUploadHeaders() map[string]interface{}`

GetUploadHeaders returns the UploadHeaders field if non-nil, zero value otherwise.

### GetUploadHeadersOk

`func (o *PackageFileUploadResponse) GetUploadHeadersOk() (*map[string]interface{}, bool)`

GetUploadHeadersOk returns a tuple with the UploadHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadHeaders

`func (o *PackageFileUploadResponse) SetUploadHeaders(v map[string]interface{})`

SetUploadHeaders sets UploadHeaders field to given value.

### HasUploadHeaders

`func (o *PackageFileUploadResponse) HasUploadHeaders() bool`

HasUploadHeaders returns a boolean if a field has been set.

### SetUploadHeadersNil

`func (o *PackageFileUploadResponse) SetUploadHeadersNil(b bool)`

 SetUploadHeadersNil sets the value for UploadHeaders to be an explicit nil

### UnsetUploadHeaders
`func (o *PackageFileUploadResponse) UnsetUploadHeaders()`

UnsetUploadHeaders ensures that no value is present for UploadHeaders, not even an explicit nil
### GetUploadQuerystring

`func (o *PackageFileUploadResponse) GetUploadQuerystring() string`

GetUploadQuerystring returns the UploadQuerystring field if non-nil, zero value otherwise.

### GetUploadQuerystringOk

`func (o *PackageFileUploadResponse) GetUploadQuerystringOk() (*string, bool)`

GetUploadQuerystringOk returns a tuple with the UploadQuerystring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadQuerystring

`func (o *PackageFileUploadResponse) SetUploadQuerystring(v string)`

SetUploadQuerystring sets UploadQuerystring field to given value.

### HasUploadQuerystring

`func (o *PackageFileUploadResponse) HasUploadQuerystring() bool`

HasUploadQuerystring returns a boolean if a field has been set.

### SetUploadQuerystringNil

`func (o *PackageFileUploadResponse) SetUploadQuerystringNil(b bool)`

 SetUploadQuerystringNil sets the value for UploadQuerystring to be an explicit nil

### UnsetUploadQuerystring
`func (o *PackageFileUploadResponse) UnsetUploadQuerystring()`

UnsetUploadQuerystring ensures that no value is present for UploadQuerystring, not even an explicit nil
### GetUploadUrl

`func (o *PackageFileUploadResponse) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *PackageFileUploadResponse) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *PackageFileUploadResponse) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *PackageFileUploadResponse) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


