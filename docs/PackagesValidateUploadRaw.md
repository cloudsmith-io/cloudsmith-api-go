# PackagesValidateUploadRaw

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | A custom content/media (also known as MIME) type to be sent when downloading this file. By default Cloudsmith will attempt to detect the type, but if you need to override it, you can specify it here. | [optional] 
**Description** | Pointer to **string** | A textual description of this package. | [optional] 
**Name** | Pointer to **string** | The name of this package. | [optional] 
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Summary** | Pointer to **string** | A one-liner synopsis of this package. | [optional] 
**Tags** | Pointer to **string** | A comma-separated values list of tags to add to the package. | [optional] 
**Version** | Pointer to **string** | The raw version for this package. | [optional] 

## Methods

### NewPackagesValidateUploadRaw

`func NewPackagesValidateUploadRaw(packageFile string, ) *PackagesValidateUploadRaw`

NewPackagesValidateUploadRaw instantiates a new PackagesValidateUploadRaw object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesValidateUploadRawWithDefaults

`func NewPackagesValidateUploadRawWithDefaults() *PackagesValidateUploadRaw`

NewPackagesValidateUploadRawWithDefaults instantiates a new PackagesValidateUploadRaw object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *PackagesValidateUploadRaw) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PackagesValidateUploadRaw) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PackagesValidateUploadRaw) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PackagesValidateUploadRaw) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetDescription

`func (o *PackagesValidateUploadRaw) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PackagesValidateUploadRaw) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PackagesValidateUploadRaw) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PackagesValidateUploadRaw) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PackagesValidateUploadRaw) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackagesValidateUploadRaw) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackagesValidateUploadRaw) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PackagesValidateUploadRaw) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackageFile

`func (o *PackagesValidateUploadRaw) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *PackagesValidateUploadRaw) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *PackagesValidateUploadRaw) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetRepublish

`func (o *PackagesValidateUploadRaw) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *PackagesValidateUploadRaw) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *PackagesValidateUploadRaw) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *PackagesValidateUploadRaw) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetSummary

`func (o *PackagesValidateUploadRaw) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PackagesValidateUploadRaw) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PackagesValidateUploadRaw) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PackagesValidateUploadRaw) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTags

`func (o *PackagesValidateUploadRaw) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PackagesValidateUploadRaw) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PackagesValidateUploadRaw) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PackagesValidateUploadRaw) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *PackagesValidateUploadRaw) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PackagesValidateUploadRaw) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PackagesValidateUploadRaw) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PackagesValidateUploadRaw) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


